# Go Backend Template

A single-user Go backend scaffold for quickly bootstrapping Go API projects.

## Integrated components

- [Gin](https://github.com/gin-gonic/gin)
- [Viper](https://github.com/spf13/viper)
- [Cobra](https://github.com/spf13/cobra)
- [Charmbracelet/log](https://github.com/charmbracelet/log)
- [SQLite](https://github.com/go-gorm/sqlite)
- [CORS](https://github.com/gin-contrib/cors)
- Cookie Auth
- Static
- Router
- Shutdown

## Web (TanStack Start)

The UI is a TanStack Start SPA in `web/`. Production assets are built into repo-root `static/` and served by the Go binary. JavaScript tooling uses [Bun](https://bun.sh/).

- **Dev:** Terminal 1 — `go-backend-template start`. Terminal 2 — `cd web && bun run dev` (Vite proxies `/api` to `http://127.0.0.1:8080`).
- **Prod build:** `bash scripts/web-build.sh` (`bun install --frozen-lockfile` + `bun run build` in `web/`)
- **Auth:** Cookie login uses `fetch` with `credentials: 'include'`; use the Vite proxy in dev so the browser stays same-origin for `/api`.

## Container images (GHCR)

GitHub Actions push Alpine images to [GHCR](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry) as `ghcr.io/<owner>/<repo>:<tag>` (`IMAGE_NAME` = `${{ github.repository }}` in workflows):

| Tag | When |
|-----|------|
| `:develop` | Push to `develop` |
| `:latest` | Release |
| `:<git-tag>` | Release |

Example: `docker pull ghcr.io/312022151125/go-backend-template:develop`

Uses `GITHUB_TOKEN` with `packages: write`. For private packages: `docker login ghcr.io` with a PAT (`read:packages`).