# Getting Started

## Requirements

- Go 1.24 or newer
- `make`
- Optional: `gitleaks` and `actionlint` for local parity with release checks

## Five-Minute Review

```bash
git clone https://github.com/hairglasses/open-mcpkit.git
cd open-mcpkit
make ci
go run ./cmd/open-mcpkit manifest
go run ./cmd/open-mcpkit call sample_echo --param message=hello
go run ./cmd/open-mcpkit call sample_launch_plan --param goal='inspect a repo' --param provider=codex
go run ./cmd/open-mcpkit call sample_policy_check --param action=delete --param risk=high
```

All commands are local and deterministic. The sample emits reviewable output and
does not perform network calls or external mutations.

## Review Path

1. Start with `README.md` for positioning.
2. Read `docs/ARCHITECTURE.md` for the request path.
3. Compare output shapes in `docs/EXAMPLES.md`.
4. Read `PUBLIC_BOUNDARY.md` before adding tools or examples.
5. Run `make ci` before opening a pull request or publishing a release.
