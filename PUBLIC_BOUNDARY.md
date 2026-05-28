# Public Boundary

`open-mcpkit` is a public-safe MCP-style server-pattern sample. Fixtures,
examples, tests, and command output must remain synthetic, local-only, and
non-mutating.

## Included

- Typed tool contracts and JSON-schema-like manifests.
- Gateway dispatch and middleware chaining.
- Dry-run policy gates and response budgeting.
- Deterministic local demo calls.

## Excluded

- Real account connectors, browser sessions, OAuth state, cookies, credentials,
  tenant databases, recruiter/application data, private repo manifests, local
  workstation paths, and generated private artifacts.
- Live send, live submit, calendar changes, destructive admin actions, or any
  network side effect.

## Release Gate

Before each release or major example expansion:

```bash
make ci
gitleaks detect --source . --redact
```
