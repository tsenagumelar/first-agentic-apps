# ADR-002 - Stack Decision for MVP (Accepted)

## Status
Accepted

## Date
2026-05-22

## Context
MVP membutuhkan kecepatan delivery, namun flow autentikasi dan RBAC tetap harus aman dan maintainable.

## Decision
Gunakan stack berikut:
- Frontend: Next.js
- Backend critical domain: Golang + Fiber (REST API)
- Non-critical general CRUD acceleration: Hasura
- Database: PostgreSQL

Boundary keputusan:
- Auth + RBAC enforcement berada di Go Fiber.
- CRUD non-kritikal boleh di Hasura selama tidak memindahkan security-critical decisions ke layer yang salah.

## Consequences
- Pro:
  - Security-critical logic tetap terkontrol dalam service backend.
  - Delivery CRUD non-kritikal bisa dipercepat.
  - Arsitektur tetap pragmatis untuk MVP.
- Cons:
  - Operasional dua backend surface (Go API + Hasura).
  - Perlu governance schema/permission agar tidak overlap liar.

## Guardrails
- Satu sumber kebenaran otorisasi: backend policy + permission model konsisten.
- Endpoint auth/token tidak di-handle Hasura.
- Semua perubahan schema lewat migration terkontrol.

## Next Action
Turunkan ke task implementasi per slice pada backlog engineer.
