# open-mcpkit

[![ci](https://github.com/hairglasses/open-mcpkit/actions/workflows/ci.yml/badge.svg)](https://github.com/hairglasses/open-mcpkit/actions/workflows/ci.yml)

Public-safe Go reference implementation for MCP-style server patterns: typed
in-process tools, gateway dispatch, middleware chains, policy gates, response
budgets, and dry-run execution boundaries.

This repository is intentionally small. It demonstrates the engineering shape
of production MCP infrastructure without publishing private account connectors,
tenant data, workstation state, OAuth flows, browser automation, or operational
systems.

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
gitleaks detect --source . --redact
```
