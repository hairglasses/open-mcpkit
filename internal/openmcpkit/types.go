package openmcpkit

import "context"

type Tool struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	InputSchema map[string]any `json:"input_schema"`
	ReadOnly    bool           `json:"read_only"`
	Risk        string         `json:"risk"`
}

type Manifest struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Tools       []Tool `json:"tools"`
}

type Request struct {
	Tool   string            `json:"tool"`
	Params map[string]string `json:"params"`
}

type Response struct {
	OK        bool           `json:"ok"`
	Tool      string         `json:"tool"`
	DryRun    bool           `json:"dry_run"`
	Payload   map[string]any `json:"payload,omitempty"`
	Warnings  []string       `json:"warnings,omitempty"`
	NextSteps []string       `json:"next_steps,omitempty"`
}

type Handler func(context.Context, Request) (Response, error)

type Middleware func(Handler) Handler
