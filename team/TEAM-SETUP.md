# AI Product Team Setup

Dokumen ini adalah sumber kebenaran tim agent untuk membangun produk.

## Tujuan

Membangun aplikasi production-ready dengan kolaborasi 3 peran:
- Architect Agent
- Product Owner / Business Analyst Agent
- Fullstack Software Engineer Agent

## Struktur Tim

1. Architect Agent
- Menentukan arsitektur, standar teknis, dan guardrail non-fungsional.
- Menjaga keputusan teknis konsisten dan scalable.

2. PO/BA Agent
- Mendefinisikan kebutuhan bisnis, user flow, acceptance criteria, prioritas.
- Menjaga scope tetap fokus ke outcome.

3. Fullstack Engineer Agent
- Implementasi end-to-end berdasarkan spesifikasi.
- Menulis test, dokumentasi teknis, dan hardening dasar.

## Working Mode (Default)

1. PO/BA buat PRD ringkas + backlog prioritas.
2. Architect turunkan ke blueprint teknis + milestone.
3. Engineer implementasi per tiket (vertical slice).
4. Architect review teknis + risk.
5. PO/BA validasi acceptance criteria.
6. Rilis iteratif.

## Artefak Wajib

- `team/roles/*.md`
- `team/templates/*.md`
- `team/backlog.md`
- `team/decisions/` (ADR ringan)

## Definition of Ready (DoR)

Item backlog siap dikerjakan jika punya:
- Problem statement
- Scope (in/out)
- Acceptance criteria terukur
- Dependency jelas
- Risiko utama disebutkan

## Definition of Done (DoD)

Sebuah item dianggap selesai jika:
- Fitur sesuai acceptance criteria
- Test minimal berjalan (unit/integration sesuai kebutuhan)
- Tidak ada secret hardcoded
- Logging/error handling dasar ada
- Catatan deploy/rollback tersedia


## Kickoff Saat Ini (MVP v1)

Requirement fase awal yang sudah disepakati:
- Auth: login, register, forgot password, reset password, verifikasi akun
- Dashboard: home dashboard placeholder
- Master Data: user management, role management, module/menu management
- Akses: RBAC role access

Dokumen kerja aktif:
- PRD awal: `team/prd/mvp-web-admin-v1.md`
- Blueprint teknis: `team/architecture/mvp-blueprint-auth-rbac.md`
- Backlog implementasi: `team/backlog.md`
- ADR awal: `team/decisions/ADR-001-auth-session-strategy.md`
