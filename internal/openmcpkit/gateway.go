package openmcpkit

import (
	"context"
	"errors"
	"fmt"
	"sort"
)

type registeredTool struct {
	tool    Tool
	handler Handler
}

type Gateway struct {
	name        string
	version     string
	description string
	tools       map[string]registeredTool
	middleware  []Middleware
}

func NewGateway(name, version, description string, middleware ...Middleware) *Gateway {
	return &Gateway{
		name:        name,
		version:     version,
		description: description,
		tools:       map[string]registeredTool{},
		middleware:  append([]Middleware(nil), middleware...),
	}
}

func (g *Gateway) Register(tool Tool, handler Handler) error {
	if tool.Name == "" {
		return errors.New("tool name is required")
	}
	if handler == nil {
		return fmt.Errorf("handler for %s is nil", tool.Name)
	}
	if _, exists := g.tools[tool.Name]; exists {
		return fmt.Errorf("tool %s already registered", tool.Name)
	}
	g.tools[tool.Name] = registeredTool{tool: tool, handler: handler}
	return nil
}

func (g *Gateway) Manifest() Manifest {
	tools := make([]Tool, 0, len(g.tools))
	for _, registered := range g.tools {
		tools = append(tools, registered.tool)
	}
	sort.Slice(tools, func(i, j int) bool { return tools[i].Name < tools[j].Name })
	return Manifest{Name: g.name, Version: g.version, Description: g.description, Tools: tools}
}

func (g *Gateway) Call(ctx context.Context, req Request) (Response, error) {
	registered, ok := g.tools[req.Tool]
	if !ok {
		return Response{}, fmt.Errorf("unknown tool %q", req.Tool)
	}
	handler := registered.handler
	for i := len(g.middleware) - 1; i >= 0; i-- {
		handler = g.middleware[i](handler)
	}
	resp, err := handler(ctx, req)
	if err != nil {
		return Response{}, err
	}
	resp.Tool = req.Tool
	return resp, nil
}

func schema(required []string, properties map[string]any) map[string]any {
	return map[string]any{
		"type":       "object",
		"required":   required,
		"properties": properties,
	}
}
