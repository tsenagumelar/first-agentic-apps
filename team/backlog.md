# Product Backlog

Gunakan format ini untuk semua item.

## Format Item
- ID:
- Title:
- Type: feature / bug / tech-debt / spike
- Priority: P0 / P1 / P2
- Owner Role: PO-BA / Architect / Engineer
- Status: todo / in-progress / blocked / review / done
- Summary:
- Acceptance Criteria:
  - AC-1
  - AC-2
- Dependencies:
- Notes:

---

## Milestone MVP v1 - Web Admin Auth + RBAC

- ID: MVP-ARC-000
- Title: Lock technology stack and service boundary
- Type: feature
- Priority: P0
- Owner Role: Architect
- Status: done
- Summary: Keputusan stack dikunci: Next.js + Go Fiber + PostgreSQL + Hasura untuk non-critical CRUD.
- Acceptance Criteria:
  - AC-1 Batas domain critical vs non-critical terdokumentasi
  - AC-2 Semua role tim pakai stack decision yang sama
- Dependencies: none
- Notes: referensi `team/architecture/mvp-blueprint-auth-rbac.md`

- ID: MVP-PO-001
- Title: Freeze MVP scope and success criteria
- Type: feature
- Priority: P0
- Owner Role: PO-BA
- Status: todo
- Summary: Finalisasi ruang lingkup MVP berdasarkan PRD awal dan target validasi bisnis.
- Acceptance Criteria:
  - AC-1 Scope in/out disepakati
  - AC-2 Success metrics MVP disepakati
- Dependencies: MVP-ARC-000
- Notes: referensi `team/prd/mvp-web-admin-v1.md`

- ID: MVP-ARC-001
- Title: Finalize architecture and data model baseline
- Type: feature
- Priority: P0
- Owner Role: Architect
- Status: todo
- Summary: Menetapkan arsitektur, model data, dan standar security baseline untuk implementasi.
- Acceptance Criteria:
  - AC-1 Blueprint teknis final tersedia
  - AC-2 ERD/domain model siap dipakai engineer
- Dependencies: MVP-PO-001
- Notes: referensi `team/architecture/mvp-blueprint-auth-rbac.md`

- ID: MVP-ENG-001
- Title: Implement auth register + email verification + login
- Type: feature
- Priority: P0
- Owner Role: Engineer
- Status: todo
- Summary: Implementasi alur auth inti hingga user berhasil login setelah verifikasi.
- Acceptance Criteria:
  - AC-1 Register berhasil dan kirim email verifikasi
  - AC-2 Verify token mengaktifkan akun
  - AC-3 Login hanya untuk akun verified aktif
- Dependencies: MVP-ARC-001
- Notes:

- ID: MVP-ENG-000A
- Title: Setup frontend foundation (Next.js + Tailwind + Atomic Design)
- Type: feature
- Priority: P0
- Owner Role: Engineer
- Status: todo
- Summary: Menyiapkan pondasi frontend reusable berbasis Atomic Design dan DRY.
- Acceptance Criteria:
  - AC-1 Tailwind terpasang dan theme token dasar tersedia
  - AC-2 Struktur folder atomic (`atoms/molecules/organisms`) siap pakai
  - AC-3 Shared form components + validation pattern reusable tersedia
- Dependencies: MVP-ARC-001
- Notes: referensi `team/architecture/mvp-blueprint-auth-rbac.md`

- ID: MVP-ENG-000B
- Title: Setup backend foundation (Go Fiber + migration + validation)
- Type: feature
- Priority: P0
- Owner Role: Engineer
- Status: todo
- Summary: Menyiapkan pondasi backend Fiber dengan pattern service/repository dan komponen reusable.
- Acceptance Criteria:
  - AC-1 Skeleton layer `handler/service/repository` tersedia
  - AC-2 Migration pipeline PostgreSQL berjalan
  - AC-3 Middleware reusable auth/error/request-id tersedia
- Dependencies: MVP-ARC-001
- Notes: referensi `team/architecture/mvp-blueprint-auth-rbac.md`

- ID: MVP-ENG-002
- Title: Implement forgot password + reset password
- Type: feature
- Priority: P0
- Owner Role: Engineer
- Status: todo
- Summary: Implementasi alur reset password aman berbasis token expire sekali pakai.
- Acceptance Criteria:
  - AC-1 Forgot password kirim reset token
  - AC-2 Reset password valid ubah password
  - AC-3 Token invalid/expired ditolak
- Dependencies: MVP-ENG-001
- Notes:

- ID: MVP-ENG-003
- Title: Implement dashboard home placeholder
- Type: feature
- Priority: P0
- Owner Role: Engineer
- Status: todo
- Summary: Menyediakan halaman dashboard kosong setelah login berhasil.
- Acceptance Criteria:
  - AC-1 Redirect pasca login ke dashboard
  - AC-2 Halaman hanya bisa diakses user authenticated
- Dependencies: MVP-ENG-001
- Notes:

- ID: MVP-ENG-004
- Title: Implement role and permission (RBAC core)
- Type: feature
- Priority: P0
- Owner Role: Engineer
- Status: todo
- Summary: CRUD role dan mapping role-permission + guard API berbasis permission.
- Acceptance Criteria:
  - AC-1 Role dapat dibuat/diubah/dihapus
  - AC-2 Permission role bisa di-assign
  - AC-3 Endpoint tanpa izin mengembalikan forbidden
- Dependencies: MVP-ARC-001
- Notes:

- ID: MVP-ENG-005
- Title: Implement user management
- Type: feature
- Priority: P1
- Owner Role: Engineer
- Status: todo
- Summary: CRUD user, aktivasi/non-aktif, dan assign role ke user.
- Acceptance Criteria:
  - AC-1 Admin dapat list/create/update/deactivate user
  - AC-2 Admin dapat assign role ke user
- Dependencies: MVP-ENG-004
- Notes:

- ID: MVP-ENG-006
- Title: Implement module/menu management
- Type: feature
- Priority: P1
- Owner Role: Engineer
- Status: todo
- Summary: CRUD module/menu sebagai resource yang dipakai RBAC.
- Acceptance Criteria:
  - AC-1 Admin dapat CRUD module/menu
  - AC-2 Menu terintegrasi dengan permission check
- Dependencies: MVP-ENG-004
- Notes:

- ID: MVP-ARC-002
- Title: Security hardening and review checkpoint
- Type: tech-debt
- Priority: P1
- Owner Role: Architect
- Status: todo
- Summary: Review token lifecycle, password policy, input validation, dan error handling.
- Acceptance Criteria:
  - AC-1 Checklist security baseline terpenuhi
  - AC-2 Temuan kritikal ditutup atau ada mitigasi
- Dependencies: MVP-ENG-001, MVP-ENG-002, MVP-ENG-004
- Notes:

- ID: MVP-PO-002
- Title: UAT checklist and release readiness
- Type: feature
- Priority: P1
- Owner Role: PO-BA
- Status: todo
- Summary: Validasi acceptance kunci dan kesiapan rilis MVP internal.
- Acceptance Criteria:
  - AC-1 UAT flow auth + RBAC lulus
  - AC-2 Daftar known limitations terdokumentasi
- Dependencies: MVP-ENG-003, MVP-ENG-005, MVP-ENG-006, MVP-ARC-002
- Notes:
