# Release Checklist

Use this checklist before visibility changes, tagged releases, or major example
expansions.

## Local Gates

```bash
make ci
gitleaks detect --source . --redact
```

`make ci` runs tests, vet, build, smoke commands, output-boundary checks,
gitleaks when available, and actionlint when available.

## Review

- Confirm all tools are local and deterministic.
- Confirm sample output contains no account names, emails, local paths, real
  operational data, tenant data, credentials, or private repo names.
- Confirm docs describe excluded connectors without adding implementation hooks.
- Review full Git history with gitleaks before publishing a release.
- Keep push, pull request, and manual CI triggers enabled.

## Latest Verification Snapshot

Checked on 2026-05-27 20:41 PDT:

- Local `make ci`: passed on initial scaffold before first public push.
- Working-tree secret scan: passed via `gitleaks detect --source . --no-git --redact`
  inside `make ci`.
- Full-history `gitleaks detect --source . --redact`: passed after initial
  local commit.
- GitHub visibility: public repository verified after first push.
- Public unauthenticated read: verified with `git ls-remote` and raw README
  fetch.
- GitHub Actions CI: passed on the first public push.
