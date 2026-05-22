# MVP Web Admin Monorepo

## Apps
- `apps/web`: Next.js + Tailwind + Atomic Design baseline
- `apps/api`: Go Fiber REST API baseline (auth/RBAC ready)
- `infra/hasura`: notes dan boundary Hasura

## Quick Start
1. Start infra containers (PostgreSQL + Hasura)
   - `docker compose up -d`
2. Run API locally (hot iteration)
   - `cd apps/api && go mod tidy && go run ./cmd/api`
3. Run Web locally (hot iteration)
   - `cd apps/web && npm install && npm run dev`

## Local Endpoints
- API: `http://localhost:8080`
- Hasura Console: `http://localhost:8081`
- Web: `http://localhost:3000`
