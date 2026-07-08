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

The UI is a TanStack Start SPA in `web/`. Production assets are built into repo-root `static/` and served by the Go binary.

- **Dev:** Terminal 1 — `go-backend-template start`. Terminal 2 — `cd web && npm run dev` (Vite proxies `/api` to `http://127.0.0.1:8080`).
- **Prod build:** `bash scripts/web-build.sh`
- **Auth:** Cookie login uses `fetch` with `credentials: 'include'`; use the Vite proxy in dev so the browser stays same-origin for `/api`.