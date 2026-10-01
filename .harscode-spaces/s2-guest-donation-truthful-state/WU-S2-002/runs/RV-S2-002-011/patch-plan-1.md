# Patch Plan — `RV-S2-002-011` / Finding F-001

> Phase: Code Review patch plan  
> Work Unit / Run: `WU-S2-002` / `RV-S2-002-011`  
> Author: Codex Reviewer  
> Created: 2026-10-01  
> Participant: `P-S2-002-RV-011-1`  
> Session: Fresh Reviewer Session; Session ID not exposed  
> Target revision: `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus current working-tree artifacts  
> Workflow revision: `33b03a3f62cc3aacba6534b8a011465613c64b09`

## Finding

**F-001 — Blocking:** Pulihkan persyaratan Product/MVP dan TP-015 bahwa guest status credential sulit ditebak. Biarkan bukti strength/generation yang spesifik implementasi dan kontrol O4 tetap terbuka; jangan mengarang nilai entropy/panjang numerik atau mengklaim bukti.

## Perubahan source yang tepat untuk Build/Patch Run baru

1. `docs/spec/5-donation/invariants.md`, `INV-donation-05`: nyatakan status URL dengan fragment memuat bearer credential yang sulit ditebak, bersama frontend handoff/URL cleanup, one-way HMAC verifier, hard expiry 24 jam, dan hasil status-only. Biarkan exposure/key/comparison/expiry/abuse controls O4 dan bukti residual risk tetap terbuka.
2. `docs/spec/5-donation/features/02-donation-status-check.md`, reconciliation dan acceptance criteria: masukkan “difficult-to-guess token/credential” sebagai persyaratan yang sudah ditetapkan. Pada baris threat Credential guessing/theft, bedakan persyaratan tersebut dari bukti generation/strength konkret, key/comparison controls, expiry enforcement, exposure protections, abuse controls, dan residual-risk decision yang masih terbuka.
3. `docs/spec/5-donation/tasks.md`, Task 02 Acceptance: cantumkan persyaratan credential sulit ditebak bersama arah fragment/HMAC/expiry/status-only. Baris klasifikasi tugas historis sudah memuatnya; selaraskan active acceptance.
4. `docs/spec/5-donation/threat-model.md`, Temporary guest status URL / Spoofing: pertahankan “hard-to-guess” sebagai persyaratan dan jelaskan bahwa pilihan strength/generation konkret serta buktinya masih terbuka di O4. Jangan menetapkan parameter algoritme atau mengklaim properti keamanan telah diverifikasi.

## Batasan dan konfirmasi

- Jangan ubah Product/MVP, TP-015, OpenAPI, code, tests, statuses, projections, atau file yang tidak terkait.
- Pertahankan kelima Donation spec terdampak dalam status `draft`; jangan mengklaim Human acceptance, runtime/security proof, residual-risk acceptance, atau `CONTRACT_READY`.
- Pastikan properti tersebut tercantum di invariant/active feature acceptance dan task acceptance, serta wording threat model tidak lagi membuat persyaratannya tampak opsional/belum diputuskan.
- Setelah patch, baca ulang passage yang berubah terhadap TP-015 Q7 dan Product/MVP §5. Tests/runtime checks tidak diperlukan untuk koreksi dokumentasi ini.
