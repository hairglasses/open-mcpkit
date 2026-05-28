# open-mcpkit

[![ci](https://github.com/hairglasses/open-mcpkit/actions/workflows/ci.yml/badge.svg)](https://github.com/hairglasses/open-mcpkit/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Public-safe Go reference implementation for MCP-style server patterns: typed
in-process tools, gateway dispatch, middleware chains, policy gates, response
budgets, and dry-run execution boundaries.

This repository is intentionally small. It demonstrates the engineering shape
of production MCP infrastructure without publishing private account connectors,
tenant data, workstation state, OAuth flows, browser automation, or operational
systems.

## Start Here

For a quick review path:

1. Run the five-minute commands in [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md).
2. Compare output shapes in [docs/EXAMPLES.md](docs/EXAMPLES.md).
3. Review the request path in [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md).
4. Use [docs/PORTFOLIO_PROOF.md](docs/PORTFOLIO_PROOF.md) for the architecture
   diagram, walkthrough plan, tradeoffs, and interview prompts.
5. Check [PUBLIC_BOUNDARY.md](PUBLIC_BOUNDARY.md) before adding tools or examples.

## What Works Now

- Register typed MCP-style tools with JSON-schema-like input descriptions.
- Dispatch tool calls through a gateway with composable middleware.
- Enforce review-only policy gates for mutating or high-risk actions.
- Apply response-size budgeting so tool output stays context-friendly.
- Emit a deterministic manifest and call sample tools from a local CLI.

## Why It Exists

`open-mcpkit` is a portfolio proof object for Go MCP infrastructure patterns. It
is not a package mirror of any private framework. It is a public reference slice
that shows the reusable ideas: contracts first, dispatch second, policy and
budget checks always in the path.

## Usage

```bash
make ci
go run ./cmd/open-mcpkit manifest
go run ./cmd/open-mcpkit call sample_echo --param message=hello
go run ./cmd/open-mcpkit call sample_launch_plan --param goal='inspect a repo' --param provider=codex
go run ./cmd/open-mcpkit call sample_policy_check --param action=delete --param risk=high
```

See [docs/EXAMPLES.md](docs/EXAMPLES.md) for expected output and validation
commands.

## Not In Scope

- Gmail, LinkedIn, browser automation, OAuth, cookies, cloud credentials, live
  account connectors, tenant databases, or application submission.
- Private repository names, local workstation paths, real operational data, or
  generated private artifacts.
- Network calls. All tools are local, deterministic, and dry-run oriented.

See [PUBLIC_BOUNDARY.md](PUBLIC_BOUNDARY.md) before adding examples or tools.

## Verification

```bash
make ci
gitleaks detect --source . --no-git --redact
```

`make ci` runs tests, vet, a temporary build, deterministic smoke commands,
public-boundary checks, and optional local `gitleaks` / `actionlint` checks when
those tools are installed.
