package openmcpkit

import (
	"context"
	"fmt"
)

func ReviewOnlyPolicy() Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req Request) (Response, error) {
			resp, err := next(ctx, req)
			if err != nil {
				return Response{}, err
			}
			resp.DryRun = true
			if req.Params["risk"] == "high" || req.Params["action"] == "delete" || req.Params["action"] == "submit" {
				resp.Warnings = append(resp.Warnings, "review required before any high-risk or mutating action")
				resp.NextSteps = append(resp.NextSteps, "keep this sample dry-run only")
			}
			return resp, nil
		}
	}
}

func ResponseBudget(maxPayloadValues int) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, req Request) (Response, error) {
			resp, err := next(ctx, req)
			if err != nil {
				return Response{}, err
			}
			if maxPayloadValues <= 0 || len(resp.Payload) <= maxPayloadValues {
				return resp, nil
			}
			trimmed := make(map[string]any, maxPayloadValues)
			kept := 0
			for key, value := range resp.Payload {
				if kept == maxPayloadValues {
					break
				}
				trimmed[key] = value
				kept++
			}
			resp.Payload = trimmed
			resp.Warnings = append(resp.Warnings, fmt.Sprintf("payload trimmed to %d fields", maxPayloadValues))
			return resp, nil
		}
	}
}
