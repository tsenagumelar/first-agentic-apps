# MVP Technical Blueprint - Auth + RBAC

## Arsitektur Ringkas
- Frontend: Next.js (SSR/CSR hybrid) dengan route guard berbasis permission.
- Backend critical: Golang Fiber REST API (auth, RBAC enforcement, security-sensitive flow).
- Backend non-critical CRUD: Hasura (terbatas pada resource low-risk).
- Database: Relational DB (PostgreSQL) untuk user/role/menu/permission mapping.
- Mail Service: SMTP provider/dev mail catcher untuk verifikasi dan reset password.

## Local Development Runtime
- Containerized (Docker Compose):
  - PostgreSQL
  - Hasura
- Local process (non-container untuk kecepatan coding):
  - Next.js web app
  - Go Fiber API
- Tujuan:
  - Infra dependency konsisten antar developer
  - Iterasi code FE/BE cepat tanpa rebuild container tiap perubahan

## Frontend Standards (Next.js)
- Core:
  - Next.js (App Router) + TypeScript
  - Tailwind CSS untuk styling utility-first
  - TanStack Query untuk server-state/data fetching
  - React Hook Form + Zod untuk form dan validasi schema
  - Zustand (opsional minimal) untuk client state ringan
- UI Architecture:
  - Atomic Design: `atoms`, `molecules`, `organisms`, `templates`, `pages`
  - Reusable-first components dengan prop API konsisten
  - DRY principle: shared hooks (`useAuth`, `usePermission`), shared util, shared UI primitives
- Folder baseline:
  - `src/components/atoms`
  - `src/components/molecules`
  - `src/components/organisms`
  - `src/features/<domain>`
  - `src/lib` (api client, utils, constants)
  - `src/hooks`

## Backend Standards (Go Fiber)
- Core:
  - Go + Fiber
  - `pgx` atau `sqlx` untuk PostgreSQL access
  - `golang-migrate` untuk database migration
  - `go-playground/validator` untuk request validation
  - `golang-jwt/jwt` untuk access token
  - `golang.org/x/crypto` (argon2/bcrypt) untuk password hashing
  - structured logger (`zap` atau `zerolog`)
- Architecture:
  - Layered: `handler -> service -> repository`
  - Domain package per feature: `auth`, `users`, `roles`, `modules`, `rbac`
  - DRY principle: middleware reusable (authn, authz, request-id, error mapping)
- Security:
  - Token verification/reset disimpan hashed + one-time + expiry
  - Seluruh keputusan permission dilakukan di service/backend

## Boundary Rules (Go vs Hasura)
- Wajib di Go Fiber:
  - Register, login, verify email, forgot/reset password, refresh token.
  - Semua enforcement RBAC di backend.
  - Endpoint yang memproses token/credential/permission decision.
- Boleh di Hasura:
  - Query/mutation CRUD umum yang tidak menyentuh keputusan security kritikal.
  - Read-heavy data list internal admin.
- Larangan:
  - Jangan menaruh business rule security kritikal di client-side atau hanya di Hasura permission tanpa guard backend untuk flow auth inti.

## Resource Ownership Matrix (Go vs Hasura)
- Go-only:
  - `users` (credential lifecycle, activation, verification flags)
  - `auth_tokens`
  - `user_roles` write path
  - Semua endpoint `/auth/*`
  - Semua endpoint decision access (authz middleware)
- Hasura-eligible (non-critical CRUD/query):
  - `roles` read/write administratif
  - `modules` read/write administratif
  - `permissions` read administratif
  - `role_permissions` write administratif (dengan guard policy)
- Catatan:
  - Walau tabel dikelola Hasura, enforcement final tetap di backend Go saat runtime request protected resource.

## Domain Model Awal
- users
  - id, name, email (unique), password_hash, is_active, is_verified, created_at, updated_at
- roles
  - id, code (unique), name, description
- modules
  - id, code (unique), name, path, parent_id (nullable), sort_order
- permissions
  - id, code (unique), module_id, action (view/create/update/delete)
- role_permissions
  - role_id, permission_id
- user_roles
  - user_id, role_id
- auth_tokens
  - id, user_id, token_hash, token_type (verify_email/reset_password/refresh), expires_at, used_at

## API Capability (MVP)
- Auth:
  - POST /auth/register
  - POST /auth/login
  - POST /auth/forgot-password
  - POST /auth/reset-password
  - POST /auth/verify-email
  - POST /auth/refresh (opsional jika pakai refresh token)
- Users:
  - GET /users, GET /users/{id}, POST /users, PUT /users/{id}, PATCH /users/{id}/status
- Roles:
  - GET /roles, POST /roles, PUT /roles/{id}, DELETE /roles/{id}
- Modules/Menus:
  - GET /modules, POST /modules, PUT /modules/{id}, DELETE /modules/{id}
- Role Access:
  - GET /roles/{id}/permissions
  - PUT /roles/{id}/permissions

## Security Baseline
- Password hash: Argon2id (preferred) atau bcrypt cost memadai.
- Token reset/verifikasi:
  - Simpan token dalam bentuk hash.
  - One-time use (`used_at`), expire ketat.
- AuthZ:
  - Middleware cek authn + permission per endpoint.
  - UI guard hanya pelengkap; enforcement utama di backend.

## Security Profile v1 (Operational Defaults)
- Access token TTL: 15 menit.
- Refresh token TTL: 7 hari (rotation saat refresh).
- Verify email token TTL: 24 jam.
- Reset password token TTL: 30 menit.
- Password policy:
  - minimum 8 karakter
  - wajib kombinasi huruf + angka
- Login protection:
  - max 5 percobaan gagal / 15 menit / akun
  - lock sementara 15 menit setelah melewati batas
- API protection:
  - rate limit baseline pada endpoint `/auth/login`, `/auth/forgot-password`, `/auth/reset-password`

## Vertical Slice Delivery
- Slice 1 (P0): Register + Verify + Login + Dashboard placeholder.
- Slice 2 (P0): Forgot/Reset Password secure flow.
- Slice 3 (P0): Role + Permission + Guard API (RBAC core).
- Slice 4 (P1): User Management + Assign role.
- Slice 5 (P1): Module/Menu Management + integrasi RBAC UI.

## Open Decisions (ADR)
- Session strategy: stateless JWT + refresh token vs server session.
- Permission granularity: per module-action fixed vs custom action.
