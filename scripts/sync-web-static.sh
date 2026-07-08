#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CLIENT_DIR="${ROOT}/web/dist/client"
STATIC_DIR="${ROOT}/static"

if [[ ! -d "${CLIENT_DIR}" ]]; then
  echo "error: ${CLIENT_DIR} not found. Run: cd web && npm run build" >&2
  exit 1
fi

mkdir -p "${STATIC_DIR}"
rm -rf "${STATIC_DIR:?}"/*
cp -a "${CLIENT_DIR}/." "${STATIC_DIR}/"

if [[ ! -f "${STATIC_DIR}/index.html" && -f "${STATIC_DIR}/_shell.html" ]]; then
  cp "${STATIC_DIR}/_shell.html" "${STATIC_DIR}/index.html"
fi

if [[ ! -f "${STATIC_DIR}/index.html" ]]; then
  echo "error: static/index.html missing after sync" >&2
  exit 1
fi

echo "synced ${CLIENT_DIR} -> ${STATIC_DIR}"