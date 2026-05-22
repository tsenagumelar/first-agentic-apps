# ADR-001 - Auth Session Strategy (Accepted)

## Status
Accepted

## Context
MVP membutuhkan autentikasi aman dengan implementasi yang cepat, maintainable, dan mudah di-scale.

## Decision
Gunakan:
- Access token JWT durasi pendek.
- Refresh token tersimpan hashed di DB (`auth_tokens` dengan type refresh).
- Rotation refresh token saat renew session.
- Akses token default TTL: 15 menit.
- Refresh token default TTL: 7 hari.

## Consequences
- Pro:
  - Scale horizontal lebih mudah.
  - Control sesi lebih baik (revoke refresh token).
- Cons:
  - Implementasi lifecycle token sedikit lebih kompleks dibanding session cookie sederhana.

## Next Action
Turunkan ke implementasi endpoint auth dan token lifecycle.
