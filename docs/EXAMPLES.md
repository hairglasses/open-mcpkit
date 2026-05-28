# Examples

All examples are local and deterministic.

## Manifest

```bash
go run ./cmd/open-mcpkit manifest
```

The manifest lists the local sample tools, their input schemas, read-only
status, and risk labels.

Expected shape:

```json
[
  {
    "name": "sample_echo",
    "description": "Echo a synthetic message.",
    "input_schema": {"type": "object"},
    "read_only": true,
    "risk": "low"
  }
]
```

## Echo Tool

```bash
go run ./cmd/open-mcpkit call sample_echo --param message=hello
```

Expected shape:

```json
{
  "ok": true,
  "tool": "sample_echo",
  "dry_run": true,
  "payload": {
    "message": "hello"
  }
}
```

## Launch Plan Tool

```bash
go run ./cmd/open-mcpkit call sample_launch_plan --param goal='inspect a repo' --param provider=codex
```

The launch-plan sample returns reviewable steps. It does not start a provider
process.

## Policy Check Tool

```bash
go run ./cmd/open-mcpkit call sample_policy_check --param action=delete --param risk=high
```

The policy sample returns a dry-run approval requirement instead of performing a
mutation.

Expected high-risk shape:

```json
{
  "ok": true,
  "tool": "sample_policy_check",
  "dry_run": true,
  "payload": {
    "allowed_without_review": false,
    "decision": "review required; sample remains dry-run only"
  }
}
```
