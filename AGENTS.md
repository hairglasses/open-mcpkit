# open-mcpkit - Agent Instructions

> Canonical instructions: AGENTS.md

Public-safe Go MCP-style server-pattern sample. Every fixture, example, and
commit must preserve the synthetic/local-only boundary.

## Public Boundary

- Use only synthetic tool names, requests, and outputs.
- Do not copy private framework code, tenant files, OAuth state, browser state,
  account identifiers, local workstation paths, or generated private artifacts.
- Do not add network connectors, browser automation, Gmail, LinkedIn, calendar,
  or live-submit code.

## Verification

- Run `make ci` before committing meaningful changes.
- Run `gitleaks detect --source . --redact` before any tagged release or major
  public example expansion.
- Keep README, examples, and tool outputs aligned with `PUBLIC_BOUNDARY.md`.
