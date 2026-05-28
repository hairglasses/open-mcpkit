package openmcpkit

import (
	"context"
	"fmt"
	"strings"
)

func SampleGateway() *Gateway {
	g := NewGateway("open-mcpkit", "0.1.0", "Public-safe MCP-style server pattern sample", ReviewOnlyPolicy(), ResponseBudget(8))
	must(g.Register(Tool{
		Name:        "sample_echo",
		Description: "Echo a synthetic message through the gateway.",
		InputSchema: schema([]string{"message"}, map[string]any{"message": map[string]string{"type": "string"}}),
		ReadOnly:    true,
		Risk:        "low",
	}, echoHandler))
	must(g.Register(Tool{
		Name:        "sample_launch_plan",
		Description: "Build a review-only launch plan for a synthetic provider task.",
		InputSchema: schema([]string{"goal", "provider"}, map[string]any{"goal": map[string]string{"type": "string"}, "provider": map[string]string{"type": "string"}}),
		ReadOnly:    true,
		Risk:        "medium",
	}, launchPlanHandler))
	must(g.Register(Tool{
		Name:        "sample_policy_check",
		Description: "Explain whether a synthetic action needs review before execution.",
		InputSchema: schema([]string{"action", "risk"}, map[string]any{"action": map[string]string{"type": "string"}, "risk": map[string]string{"type": "string"}}),
		ReadOnly:    true,
		Risk:        "medium",
	}, policyCheckHandler))
	return g
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func echoHandler(_ context.Context, req Request) (Response, error) {
	message := strings.TrimSpace(req.Params["message"])
	if message == "" {
		return Response{}, fmt.Errorf("message parameter is required")
	}
	return Response{OK: true, Payload: map[string]any{"message": message, "boundary": "synthetic-local-only"}}, nil
}

func launchPlanHandler(_ context.Context, req Request) (Response, error) {
	goal := strings.TrimSpace(req.Params["goal"])
	provider := strings.TrimSpace(req.Params["provider"])
	if goal == "" || provider == "" {
		return Response{}, fmt.Errorf("goal and provider parameters are required")
	}
	return Response{OK: true, Payload: map[string]any{
		"provider": provider,
		"goal":     goal,
		"plan": []string{
			"resolve tool contract before loading broad schemas",
			"run preflight checks before any connector-backed workflow",
			"emit a review packet instead of mutating external state",
		},
	}}, nil
}

func policyCheckHandler(_ context.Context, req Request) (Response, error) {
	action := strings.TrimSpace(req.Params["action"])
	risk := strings.TrimSpace(req.Params["risk"])
	if action == "" || risk == "" {
		return Response{}, fmt.Errorf("action and risk parameters are required")
	}
	allowed := risk == "low" && action != "delete" && action != "submit"
	return Response{OK: true, Payload: map[string]any{
		"action":                 action,
		"risk":                   risk,
		"allowed_without_review": allowed,
		"decision":               decisionText(allowed),
	}}, nil
}

func decisionText(allowed bool) string {
	if allowed {
		return "read-only sample action may proceed locally"
	}
	return "review required; sample remains dry-run only"
}
