# query-builder
A generic query builder UI

```
query-builder/
│── backend/             # Backend (Go + Fiber + GORM)
│   ├── cmd/             # (Optional) CLI commands, migrations
│   ├── db/              # Database connections & setup
│   ├── models/          # Database models
│   ├── routes/          # API routes
│   ├── main.go          # Entrypoint for the Go backend
│   ├── go.mod           # Go dependencies
│   ├── go.sum           # Go dependency hashes
│   ├── .air.toml        # Optional: Hot-reloading config
│   ├── Dockerfile       # (Optional) For containerizing backend
│── frontend/            # Frontend (SvelteKit)
│   ├── src/             # Svelte components, pages
│   ├── package.json     # Frontend dependencies
│   ├── svelte.config.js # Frontend config
│── .gitignore           # Ignore unneeded files
│── README.md            # Project documentation
│── docker-compose.yml   # (Optional) If using Docker for both frontend & backend
