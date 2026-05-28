#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$repo_root"

go_cmd="${GO_CMD:-go}"
forbidden_output_pattern='gmail|linkedin|oauth|cookie|tenant|mitch|mitchell|/home/hg|hairglasses-studio|jobb|secret|password|token|api[_-]?key|@[A-Za-z0-9._%+-]+\.[A-Za-z]{2,}'
tmp_dir="$(mktemp -d)"
trap 'rm -rf "$tmp_dir"' EXIT

echo "== sample output boundary =="
$go_cmd run ./cmd/open-mcpkit manifest > "$tmp_dir/manifest.json"
$go_cmd run ./cmd/open-mcpkit call sample_echo --param message=hello > "$tmp_dir/echo.json"
$go_cmd run ./cmd/open-mcpkit call sample_launch_plan --param goal='inspect a repo' --param provider=codex > "$tmp_dir/launch.json"
$go_cmd run ./cmd/open-mcpkit call sample_policy_check --param action=delete --param risk=high > "$tmp_dir/policy.json"

if rg --ignore-case --line-number "$forbidden_output_pattern" "$tmp_dir"; then
  echo "sample output contains private-boundary markers" >&2
  exit 1
fi

echo "== gitleaks =="
if command -v gitleaks >/dev/null 2>&1; then
  gitleaks detect --source . --no-git --redact
else
  echo "gitleaks not installed; skipping local secret scan"
fi

echo "== actionlint =="
if command -v actionlint >/dev/null 2>&1; then
  actionlint .github/workflows/*.yml
else
  echo "actionlint not installed; skipping workflow lint"
fi

echo "public boundary checks passed"
