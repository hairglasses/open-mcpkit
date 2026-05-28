# Architecture

`open-mcpkit` demonstrates a minimal MCP-style gateway with typed tool
contracts, middleware, and dry-run policy checks.

## Request Path

```text
CLI command
  -> parameter parsing
  -> tool registry lookup
  -> typed handler
  -> middleware chain
  -> response budget
  -> deterministic JSON output
```

## Components

| Area | Purpose |
| --- | --- |
| `internal/openmcpkit` | Tool definitions, handlers, middleware, and gateway dispatch. |
| `cmd/open-mcpkit` | CLI entrypoint for manifest and tool calls. |
| `scripts/check-public-boundary.sh` | Sample-output boundary and local hygiene checks. |

## Design Rules

- Tool inputs are explicit command parameters.
- Handlers return reviewable payloads, not side effects.
- Policy checks stay in the request path for high-risk actions.
- Response budgets keep output context-friendly.
- The sample intentionally avoids transports and connectors so the core
  gateway shape remains easy to inspect.
