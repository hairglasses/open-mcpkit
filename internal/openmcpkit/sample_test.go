package openmcpkit

import (
	"context"
	"testing"
)

func TestSampleGatewayManifest(t *testing.T) {
	manifest := SampleGateway().Manifest()
	if manifest.Name != "open-mcpkit" {
		t.Fatalf("manifest name = %q", manifest.Name)
	}
	if len(manifest.Tools) != 3 {
		t.Fatalf("tool count = %d", len(manifest.Tools))
	}
	for _, tool := range manifest.Tools {
		if tool.Name == "" || tool.InputSchema["type"] != "object" {
			t.Fatalf("invalid tool metadata: %+v", tool)
		}
	}
}

func TestReviewOnlyPolicyMarksHighRiskDryRun(t *testing.T) {
	resp, err := SampleGateway().Call(context.Background(), Request{Tool: "sample_policy_check", Params: map[string]string{"action": "delete", "risk": "high"}})
	if err != nil {
		t.Fatalf("Call() error = %v", err)
	}
	if !resp.DryRun {
		t.Fatal("expected dry-run response")
	}
	if len(resp.Warnings) == 0 || len(resp.NextSteps) == 0 {
		t.Fatalf("expected review warning and next step: %+v", resp)
	}
	if resp.Payload["allowed_without_review"] != false {
		t.Fatalf("expected high-risk action to require review: %+v", resp.Payload)
	}
}

func TestUnknownToolReturnsError(t *testing.T) {
	_, err := SampleGateway().Call(context.Background(), Request{Tool: "missing", Params: map[string]string{}})
	if err == nil {
		t.Fatal("expected unknown tool error")
	}
}
