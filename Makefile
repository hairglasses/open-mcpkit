.PHONY: build test vet smoke public-boundary gitleaks actionlint ci

GO ?= env GOWORK=off go

build:
	tmp_dir="$$(mktemp -d)"; \
	trap 'rm -rf "$$tmp_dir"' EXIT; \
	$(GO) build -o "$$tmp_dir/open-mcpkit" ./cmd/open-mcpkit

test:
	$(GO) test ./...

vet:
	$(GO) vet ./...

smoke:
	$(GO) run ./cmd/open-mcpkit manifest
	$(GO) run ./cmd/open-mcpkit call sample_echo --param message=hello
	$(GO) run ./cmd/open-mcpkit call sample_launch_plan --param goal='inspect a repo' --param provider=codex
	$(GO) run ./cmd/open-mcpkit call sample_policy_check --param action=delete --param risk=high

public-boundary:
	GO_CMD="$(GO)" scripts/check-public-boundary.sh

gitleaks:
	gitleaks detect --source . --no-git --redact

actionlint:
	actionlint .github/workflows/*.yml

ci: test vet build smoke public-boundary
