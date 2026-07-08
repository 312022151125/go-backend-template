#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${ROOT}/web"
npm ci
npm run build
cd "${ROOT}"
bash scripts/sync-web-static.sh