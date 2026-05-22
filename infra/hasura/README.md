# Hasura Notes

Hasura dipakai khusus CRUD non-critical sesuai boundary arsitektur.

## Local Run
- Service berjalan via `docker-compose.yml` (service `hasura` + `postgres`).
- Hasura Console: `http://localhost:8081`

## Scope awal
- Read/list data admin non-sensitive
- CRUD general yang tidak menyentuh auth/token/RBAC decision

## Important
- Jangan expose admin secret ke frontend.
- Enforcement final akses protected tetap di backend Go.
