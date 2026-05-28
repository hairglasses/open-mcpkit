# Examples

All examples are local and deterministic.

## Manifest

```bash
go run ./cmd/open-mcpkit manifest
```

## Echo Tool

```bash
go run ./cmd/open-mcpkit call sample_echo --param message=hello
```

## Launch Plan Tool

```bash
go run ./cmd/open-mcpkit call sample_launch_plan --param goal='inspect a repo' --param provider=codex
```

## Policy Check Tool

```bash
go run ./cmd/open-mcpkit call sample_policy_check --param action=delete --param risk=high
```

The policy sample returns a dry-run approval requirement instead of performing a
mutation.
