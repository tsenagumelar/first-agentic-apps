# PRD Lite - MVP Web Admin v1

## 1. Problem
Tim belum memiliki aplikasi web admin terpusat untuk autentikasi dan kontrol akses role-based. Akibatnya, pengelolaan user, role, menu/module, dan izin akses belum konsisten dan sulit diaudit.

## 2. Goal
Membangun MVP web-based admin panel yang menyediakan:
- Auth lengkap: login, register, forgot password, reset password, verifikasi akun.
- Dashboard home placeholder.
- Master data awal: user management, role management, module/menu management.
- RBAC: mapping role ke module/menu/action.

Target hasil MVP:
- User bisa autentikasi end-to-end secara aman.
- Admin bisa mengelola user, role, menu/module, dan akses role.
- Akses halaman dan aksi mengikuti permission RBAC.

## 3. Scope
- In Scope:
  - Web app (frontend + backend API + database).
  - Email verification flow.
  - Forgot/reset password flow berbasis token expire.
  - CRUD User Management.
  - CRUD Role Management.
  - CRUD Module/Menu Management.
  - Role Access RBAC (role-permission mapping).
  - Empty dashboard page setelah login.
- Out of Scope:
  - SSO (Google/Microsoft/OAuth enterprise).
  - Multi-tenant.
  - Audit log advanced dan approval workflow.
  - Fine-grained ABAC/policy engine complex.

## 4. Target User
- Persona utama:
  - Super Admin: setup awal sistem.
  - Admin Operasional: kelola user/role/menu.
- Konteks penggunaan:
  - Operasional internal melalui browser desktop/mobile.

## 5. User Flow
1. User register akun.
2. Sistem kirim email verifikasi.
3. User verifikasi akun lalu login.
4. User masuk ke dashboard kosong.
5. Admin mengelola user/role/module-menu.
6. Admin set role access RBAC.
7. Sistem enforce akses berdasar role.

## 6. Functional Requirements
- FR-1 Authentication
  - Register akun baru.
  - Login dengan kredensial valid.
  - Forgot password kirim email reset.
  - Reset password dengan token valid.
  - Verifikasi akun via email token.
- FR-2 Dashboard
  - Menampilkan halaman home dashboard kosong (placeholder).
- FR-3 User Management
  - List, detail, create, update, deactivate user.
  - Assign role ke user.
- FR-4 Role Management
  - CRUD role.
- FR-5 Module/Menu Management
  - CRUD module/menu.
- FR-6 RBAC
  - Set permission role terhadap module/menu/action (view/create/update/delete).
  - Middleware/guard API dan UI berdasarkan permission.

## 7. Non-Functional Requirements
- Performance:
  - P95 response API read < 500ms pada beban MVP.
- Security:
  - Password hashing (Argon2id/bcrypt).
  - JWT/secure session + refresh strategy.
  - Token reset/verifikasi sekali pakai dan expiry.
  - Validasi input server-side.
  - Baseline security default:
    - Access token TTL 15 menit
    - Refresh token TTL 7 hari
    - Verify token TTL 24 jam
    - Reset password token TTL 30 menit
    - Max gagal login 5 kali per 15 menit, lock 15 menit
- Reliability:
  - Error handling konsisten.
  - Migration database versioned.

## 8. Acceptance Criteria
- AC-1 User dapat register, verifikasi, login, forgot/reset password end-to-end.
- AC-2 Setelah login, user diarahkan ke dashboard placeholder.
- AC-3 Admin dapat CRUD user, role, module/menu.
- AC-4 Admin dapat set role access; akses endpoint/halaman diblok jika tidak punya izin.
- AC-5 Semua flow utama punya test minimal (unit/integration/e2e light sesuai prioritas).

## 9. Risks & Assumptions
- Risiko:
  - Scope auth + RBAC cukup besar untuk MVP jika tanpa prioritas slice.
  - Potensi bug security pada token lifecycle.
- Asumsi:
  - Email service tersedia (SMTP/dev mail catcher) untuk verifikasi & reset.
  - Hanya 1 organisasi/internal tenant untuk fase awal.

## 10. Technology Direction (Locked for MVP)
- Frontend: Next.js
- Backend critical domain: Golang + Fiber (REST API)
- General non-critical CRUD acceleration: Hasura
- Database utama: PostgreSQL
