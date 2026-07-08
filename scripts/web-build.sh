#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${ROOT}/web"
bun install --frozen-lockfile
bun run build
cd "${ROOT}"
bash scripts/sync-web-static.sh