# Stage 3 — Solutioning

> Phase/Stage: Exploration / Stage 3  
> Author: Explorer (`P-S2-005-EXP-001-1`)  
> Created: 2026-10-01 08:06 UTC  
> Model / Reasoning: Invocation-selected `gpt-6-luna` / `high` (runtime model metadata not independently exposed)  
> Session: interactive dispatch session; session identifier not exposed  
> Target revision: Kencleng `7fd8b473b239b20bda3990ab29c51440d321a796`; working tree already contained other Slice-2 orchestration artifacts at dispatch; no production-source changes were reported  
> Workflow revision: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`  
> Work Unit / Run: `WU-S2-005` / `EXP-S2-005-001`

## Keadaan solutioning

Stage 2 menetapkan bahwa kebutuhan Product Slice 2 sudah jelas dan gap berada pada Campaign Detail contract/spec yang masih mendefinisikan `donation_action` sebagai unavailable. WU-S2-004 sebelumnya tetap frontend-scoped; pilihan koordinasi terpisah pada WU-S2-005 telah disetujui dan tidak dibuka ulang di sini.

Keputusan tersisa untuk owner Campaign/API adalah apakah response `donation_action` hanya menyatakan availability yang ditentukan backend pada saat GET, atau juga membawa target navigasi untuk consumer. Kedua arah harus mempertahankan public-safe projection, membatasi entry untuk Campaign yang backend anggap eligible pada saat pembacaan, dan menyatakan secara eksplisit bahwa hasil GET dapat stale. POST tetap memeriksa eligibility secara otoritatif dan dapat menolak request yang sudah tidak eligible.

## Decision framing — Semantik `donation_action`

**Problem**  
Kontrak saat ini hanya mengizinkan `unavailable` / `donation_flow_not_available`, sehingga Campaign Detail tidak dapat menyatakan entry Slice-2 yang diwajibkan Product. API belum menentukan apakah backend hanya melaporkan availability atau juga mengarahkan consumer ke suatu target.

**Konteks saat ini**

- Product Slice 2 memerlukan guest donation entry dari eligible Public Campaign Detail (`docs/product/mvp-delivery-slices.md` §5; `docs/product/mvp-scope.md` §5 Stage B).
- Campaign API dimiliki Anhar untuk current Slice 2 (`.harscode-spaces/authority-map.md`); bentuk kontrak final harus owner-defined.
- `POST /campaigns/{campaignId}/donations` sudah menjadi submission contract guest di `api/openapi/donation.yaml`; backend dapat menolak eligibility yang berubah dengan `409`.
- Detail response adalah snapshot. Backend harus tetap menjadi eligibility authority saat submit; frontend visibility tidak menjadi authorization.
- Frontend architecture menempatkan route-specific composition pada `frontend/app/<route>/` dan API-owned shape pada generated OpenAPI types (`docs/project/kencleng-frontend-tech-stack.md` §§5, 7; `frontend/AGENTS.md` §§2–3).
- Scope WU-S2-005 tidak mencakup perubahan Product, D1 eligibility invariant, Donation submission contract, Slice 3 closure/result semantics, atau production implementation.

**Opsi**

1. **Availability projection saja (rekomendasi).** `donation_action` menyatakan state backend-authored, misalnya `available` atau `unavailable`; bila unavailable, contract dapat menyatakan reason yang public-safe sesuai keputusan owner. Consumer memakai Campaign ID untuk membuka route yang ia miliki. Makna `available` adalah eligibility yang backend nilai saat GET, bukan jaminan bahwa POST berikutnya akan diterima.
2. **Availability beserta target action.** Selain state, `donation_action` membawa target (misalnya URI/route) yang disediakan backend. Consumer mengikuti target tersebut untuk memulai flow; POST tetap otoritatif dan dapat menolak stale state.

**Rekomendasi**  
Pilih Opsi 1. Kontrak yang telah ada sudah memisahkan resource Campaign dari submit Donation melalui `campaignId`; frontend architecture menempatkan navigasi/komposisi route di frontend. Availability projection cukup untuk menyatakan kondisi yang diketahui backend tanpa menjadikan route frontend bagian dari API Campaign.

**Alasan dan konsekuensi**  
Opsi 1 menjaga batas tanggung jawab API dan route consumer, serta mempertahankan backend sebagai sumber availability snapshot. Konsekuensinya, frontend menentukan route lokal dari Campaign ID dan harus menangani rejection saat state berubah sebelum POST. Opsi 2 membuat kontrak lebih terarah bagi consumer, tetapi mengikat Campaign API ke target navigasi, menambah surface untuk format/validasi target, dan tetap tidak menghilangkan race antara GET dan POST.

**Arah yang ditolak**  
Tidak menampilkan action dengan mengabaikan `donation_action` atau menyimpulkan eligibility dari `lifecycle.public_state`. Itu menghidupkan kembali kontrak yang menyatakan unavailable dan memindahkan makna eligibility ke client. Opsi 2 tetap viable bila API owner menilai target memang bagian dari kontrak public action.

**Keputusan yang diminta**  
Sebagai Campaign/API owner, pilih Opsi 1 (state-only availability projection) atau Opsi 2 (availability beserta target action). Jika memilih Opsi 1, konfirmasi juga bahwa `available` bermakna backend menilai donation eligible pada saat GET, sementara backend submission melakukan pemeriksaan baru; vocabulary `reason` untuk state unavailable dapat ditetapkan dalam reconciliation contract berikutnya.

## Dampak yang harus dibawa ke rekonsiliasi

- Campaign feature acceptance criteria dan authored `api/openapi/campaign.yaml` perlu disejajarkan dengan keputusan owner dan requirement Slice 2. Perubahan generated bundle/type/fixture mengikuti aturan sumber API; generated counterparts bukan authority kedua.
- Closed public projection, anti-enumeration, public-safe fields, dan `Cache-Control: private, no-store` tetap berlaku.
- Jangan mengubah INV-campaign-13 / INV-donation-02 (D1) atau memperluas acceptance ke Slice 3 closure/result semantics.
- Perubahan ini tidak mengimplementasikan Campaign handler/service atau frontend Donation flow. Backend dan frontend production Build tetap terpisah.
- WU-S2-004 tetap menunggu kontrak reconciled, required generated counterparts, applicable independent review, dan owner acceptance sesuai manifest WU-S2-005.

## Batas keputusan dan risiko

- Rekomendasi availability-only adalah rekomendasi Explorer, belum merupakan `Decision` owner.
- Product/MVP tidak perlu diubah berdasarkan bukti saat ini.
- Availability dapat stale setelah GET walaupun response memiliki no-store; submission tetap dapat ditolak. Tidak ada runtime/security proof atau residual-risk acceptance yang dihasilkan oleh Run ini.
- Tidak ada implementation, spec/API edit, generated artifact, atau test yang dilakukan dalam Exploration Run ini.

## Progression

Solutioning menghasilkan pilihan owner yang siap diputuskan. Tidak ada blocker untuk menganalisis kedua opsi; WU-S2-004 tetap tertahan pada dependency kontrak sampai reconciled dan diterima.

**Human decision:** menunggu pilihan Anhar Solehudin sebagai Campaign/API owner: Opsi 1 (direkomendasikan) atau Opsi 2, serta konfirmasi makna `available` bila Opsi 1 dipilih.

## Context untuk downstream

- Product: `docs/product/mvp-delivery-slices.md` §5; `docs/product/mvp-scope.md` §5 Stages B/D.
- Campaign acceptance/invariants: `docs/spec/4-campaign/features/02-campaign-detail-listing.md`; `docs/spec/4-campaign/invariants.md` INV-campaign-13/14.
- Authored API: `api/openapi/campaign.yaml`; generated views: `api/openapi.yaml`, `frontend/lib/api/generated/openapi.ts`, fixture `frontend/mocks/fixtures/public-campaign.ts`.
- Submission boundary: `api/openapi/donation.yaml` `POST /campaigns/{campaignId}/donations`; `docs/spec/5-donation/invariants.md` INV-donation-02.
- Implementation anchors: `backend/internal/domain/campaign/service.go` `toPublicDetail`; `backend/internal/domain/campaign/repository_db.go` `FindPublicDetail`; `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` `CampaignSuccess`.
- Coordination dependency and owner: `.harscode-spaces/authority-map.md`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/manifest.md`.
