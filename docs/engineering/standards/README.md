# Kencleng — Engineering Standards

> **Authority:** Project-wide Engineering constraints yang sudah disetujui Human untuk branch `pilot/station5-c1-space-only`. Ini tidak mempromosikan standar ke `main`.
> **Scope:** Reusable lintas-feature, termasuk standar yang secara khusus berlaku pada backend atau frontend. Jangan menjadikan feature-level Rancang proposal sebagai standard secara otomatis.

## Current standards

| Concern | Owning document | Standing |
| --- | --- | --- |
| T01 — Backend architecture, Modular Monolith, REST API dan Worker boundaries, Engineering decision authority | [Backend stack & architecture](backend.md) | **Settled / persisted** |
| Frontend technology dan state architecture | [Frontend stack & architecture](frontend.md) | **Existing active baseline** |
| T02 — Google-only login, Kencleng-owned identity, HttpOnly JWT + rotating refresh | [Identity & Authentication](identity-authentication.md) | **Settled / persisted** |
| T03 — Hybrid user RBAC/context policy, service identity, permission inventory scaffold | [Authorization](authorization.md) | **Settled / persisted** |
| T04 — Data privacy, PII, encryption, key lifecycle, masking and reveal | [Privacy, PII & Cryptography](privacy-data-protection.md) | **Settled / persisted** |
| T05 — PostgreSQL schema, migrations, integrity, transactions, ERD, conditional partitioning | [Database & Data Model](database-data-model.md) | **Settled / persisted** |
| T06 — REST/OpenAPI contracts, minimal success envelope, RFC 9457 errors, security, compatibility | [API & Contract](api-contract.md) | **Settled / persisted** |
| T07 — Risk-based testing, independent suites, coverage ≥80%, table-driven Go tests | [Testing & Verification](testing-verification.md) | **Settled / persisted** |
| T08 — Minimum MVP runtime configuration, lifecycle, basic logs/health, data safety | [Minimum Runtime Reliability](operational-reliability.md) | **Settled / persisted (MVP minimum; not production-ready)** |
| T09 — Documentation language, code comments, terminology, authority, semantics preservation | [Documentation & Language](documentation-conventions.md) | **Settled / persisted** |

T01–T09 pada episode Engineering Standards ini telah **settled / persisted pada research branch ini**, bukan promosi ke `main` atau bukti seluruh mekanisme sudah diimplementasikan. Tambahkan owning document hanya jika ada keputusan baru dengan concern ownership nyata, bukan untuk melengkapi template folder.

## Final consistency review checkpoint — 28 Oktober 2026

> **Standing:** Assisted **source-level policy consistency review** terhadap T01–T09 di research branch ini, basis commit `6876a8e746ae50ccf4e9f0b2ea0ddb24b5c32f95`. **Bukan independent review, runtime conformance, implementation approval, production-readiness assessment, atau promosi ke `main`.**

- **Observed:** Tidak ditemukan **kontradiksi kebijakan material** yang menghalangi penutupan episode *standards definition* T01–T09 pada source yang diperiksa. Concern ownership tetap terbagi antara Product, Product Design, Engineering Standards, feature-level decisions, dan actual implementation evidence.
- **Structural check:** 73 internal Markdown links dari 14 standards/consumer guidance files yang diperiksa semuanya memiliki target di repository; 2 explicit section anchors juga cocok dengan heading target. Ini bukan klaim bahwa seluruh isi artifact maupun external links telah diverifikasi end-to-end.
- **Known implementation gaps (bukan policy blockers):** `backend/Makefile` masih menggabungkan optional suites dalam default `make verify` (T07); frontend Vitest mengizinkan no-test pass dan belum memiliki configured coverage gate 80% (T07); backend `migrate-down` masih menjalankan `down -all` dan tetap memerlukan explicit safeguard/authorization (T05/T08). API feature contract, runtime auth/privacy/authorization, browser E2E, dan production recovery belum otomatis delivered oleh standards approval.
- **Disposition:** **T01–T09 definition episode complete for this research branch** dengan identified implementation/evidence gaps tetap terbuka. Re-assess concern yang relevan saat actual feature work atau production-readiness dipilih. Tidak ada automatic merge/promotion; `main` tetap di luar scope.

## Batas authority dan konsumsi

- **Product meaning dan kapan login diperlukan:** [Product Authority](../../product/README.md). Engineering Standards tidak menambahkan syarat Product yang belum disetujui.
- **Product Design/UX reusable:** [Product Design Standards](../../design/README.md).
- **Kontrak API spesifik:** [API tooling dan source contract](../../../api/README.md); baseline OpenAPI masih kosong sampai ada authorized delivery requirement.
- **Feature-level Engineering output dan Harscode state:** lokasi fisik working artifacts/assignment state mengikuti workspace-local `.harscode-local.yaml` (override `projects.kencleng.harscode_state_root`) dan exact destination yang disetujui untuk assignment terkait; `kencleng/harscode-state/` adalah lokasi yang direncanakan untuk pilot berikutnya, bukan izin write atau bukti bahwa direktori/config sudah tersedia. Lokasi penyimpanan tidak mengubah project authority: proposal Rancang tidak otomatis menjadi Engineering Standard, accepted handoff, atau canonical API.
- **Executable code, configuration, dan current database state:** sumber evidence implementasi, bukan pengganti kebijakan atau Product authority. Hasil verifikasi harus dibedakan dari proposal.
- **Agent instructions:** root serta stack-scoped `AGENTS.md` mengarahkan implementer ke dokumen ini; [scoped AGENTS](AGENTS.md) berlaku untuk penulis yang mengubah direktori standards ini.

Gunakan owning document per concern. Jika ada gap/konflik material, eskalasikan ke owner yang tepat; jangan membuat aturan baru dengan menduplikasi kebijakan dari domain lain.
