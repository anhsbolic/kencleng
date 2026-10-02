# Stage 2 — Gap Analysis

> Phase/Stage: Exploration / Stage 2  
> Author: Explorer (`P-S2-005-EXP-001-1`)  
> Created: 2026-10-01 08:02 UTC  
> Model / Reasoning: Invocation-selected `gpt-6-luna` / `high` (runtime model metadata not independently exposed)  
> Session: interactive dispatch session; session identifier not exposed  
> Target revision: Kencleng `7fd8b473b239b20bda3990ab29c51440d321a796`; working tree already contained other Slice-2 orchestration artifacts at dispatch; no production-source changes were reported  
> Workflow revision: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`  
> Work Unit / Run: `WU-S2-005` / `EXP-S2-005-001`

## Ringkasan

Product/MVP Slice 2 secara eksplisit memerlukan jalur guest donation dari Campaign publik yang memenuhi syarat. Campaign feature spec dan authored API yang berlaku saat ini masih merupakan kontrak Slice 1: `donation_action` wajib tetapi hanya dapat menyatakan unavailable, tanpa activation target. Backend Campaign service, generated contract surfaces, fixture, dan UI saat ini mencerminkan arah Slice 1 tersebut. Ini adalah gap Campaign delivery/API yang jelas; belum ditemukan kekosongan makna Product yang perlu diangkat.

`donation_action` pada detail publik tidak menjadi bukti otorisasi/kelayakan submit. Kontrak Donation menyatakan backend menerima atau menolak submission sesuai eligibility saat submit dan dapat mengembalikan `409`; sumber kode Donation untuk runtime submit belum menjadi bagian dari implementasi yang terverifikasi pada frontier ini. Reconciliation harus tetap menjaga pemisahan tersebut.

## Area 1 — Product/MVP Slice 2

**Sumber dan requirement**

- `docs/product/mvp-delivery-slices.md` §5 menyebut outcome Slice 2: pengunjung publik dapat berkontribusi tanpa membuat akun dan memahami hasil sandbox yang sebenarnya.
- §5 Product scope mensyaratkan donation eligibility untuk Campaign publik aktif dan satu guest donation flow yang koheren.
- `docs/product/mvp-scope.md` §5 Stage B menyatakan pengunjung dapat submit guest donation ke Campaign yang eligible. Stage D menyatakan donation unavailable setelah Campaign closed, identitas public Campaign tetap dapat dijangkau bila memang memasuki lifecycle publik, dan hasil akhir tetap truthful.
- `WU-S2-005/manifest.md` membatasi Run pada rekonsiliasi Campaign Detail donation-action contract dan acceptance criteria Campaign. Product tetap dirujuk, bukan diubah.

**Gap**

Tidak ada gap Product yang teridentifikasi untuk kewajiban entry guest Donation. Product tidak menentukan wire shape `donation_action`; manifest dan authority map merutekan bentuk kontrak tersebut ke Campaign/API owner.

**Sniffing**

| Lensa | Temuan |
|---|---|
| Risk | Jika halaman Campaign Detail terus menyatakan aksi unavailable, loop MVP tidak dapat dimulai dari entry point yang diwajibkan. Jika UI menganggap respons detail sebagai jaminan eligibility saat POST, state yang stale dapat menghasilkan klaim keliru atau submission yang sudah tidak eligible. |
| Edge cases | Detail yang telah dimuat dapat menjadi stale sebelum submit; Campaign dapat ditutup sesudah pembacaan. Product juga membedakan Campaign closed dari detail yang tidak pernah public. |
| Miscontext | Task ini tidak sekadar mengaktifkan donation CTA di UI: otoritas Product mewajibkan alur guest, sedangkan eligibility aktual tetap bergantung pada keadaan Campaign saat submission. |
| Misleading signals | `public_state: fundraising` menggambarkan proyeksi detail saat dibaca; field itu sendiri bukan izin submit dan tidak membuktikan state belum berubah. |
| Inconsistency | Requirement Product Slice 2 dan kontrak Campaign Slice 1 berbeda pada aksi donasi. Sesuai product-first precedence, kontrak/acceptance Campaign perlu direkonsiliasi; requirement Product tidak perlu ditafsir ulang. |

## Area 2 — Campaign acceptance criteria dan invariants

**Current state**

- `docs/spec/4-campaign/features/02-campaign-detail-listing.md` berstatus reconciled untuk Slice 1. Pada acceptance behavior detail, `PublicCampaignDetail` memuat required `donation_action` unavailable dan tanpa link/activation target.
- `docs/spec/4-campaign/invariants.md#inv-campaign-14` membatasi public detail ke Campaign berstatus internal `published` pada Slice 1 dan menetapkan action unavailable pada Slice 1. Invariant yang sama melindungi closed public projection, anti-enumeration, serta no-store behavior.
- `docs/spec/4-campaign/invariants.md#inv-campaign-13` dan `docs/spec/5-donation/invariants.md#inv-donation-02` menjaga batas D1: eligibility submit berurutan terhadap close; Campaign owns lifecycle/close reason dan Donation owns submission/settlement contribution.
- `docs/spec/5-donation/features/01-submit-donation-settlement.md` telah diterima untuk Slice 2 dan secara eksplisit tidak menetapkan bentuk API submission atau Campaign Detail action.

**Gap**

Campaign feature acceptance criteria dan, bila perlu, INV-campaign-14 belum menyatakan perilaku Slice-2-compatible untuk entry action. Ini adalah perubahan lower-level yang perlu mengikuti Product/MVP. INV-campaign-13/D1 dan INV-donation-02 adalah batas yang harus dipertahankan, bukan objek perubahan WU ini.

**Sniffing**

| Lensa | Temuan |
|---|---|
| Risk | Mengubah closed projection sembarangan dapat membuka field internal atau melemahkan public-safety/anti-enumeration. Menampilkan aksi yang dapat diklik tidak boleh mengubah otoritas lifecycle atau D1. |
| Edge cases | Campaign ditutup atau tidak lagi eligible di antara GET detail dan POST; Campaign publik yang telah closed tetap memiliki halaman, tetapi tidak boleh menerima donation baru. |
| Miscontext | Spec Campaign saat ini adalah Slice-1 authority yang secara eksplisit membekukan action unavailable; keberadaan route/tipe tersebut bukan bukti bahwa Slice-2 action sudah didukung. |
| Misleading signals | INV-campaign-14 memakai kata “eligible” untuk public detail Slice 1 dengan predicate `published`; itu tidak setara dengan acceptance check Donation pada saat submit. |
| Inconsistency | Slice 1 spec mengizinkan pembacaan detail dan mewajibkan unavailable action, sedangkan Product Slice 2 mensyaratkan guest donation entry untuk Campaign publik yang eligible. Perbedaan slice ini adalah reconciliation gap, bukan bukti Product Authority saling bertentangan. |

## Area 3 — Authored dan generated API contract surfaces

**Current state**

- `api/openapi/campaign.yaml`: `GET /campaigns/{campaignId}` (`getPublicCampaignDetail`) merespons `PublicCampaignDetail`; schema tersebut mewajibkan `donation_action`. `PublicCampaignDonationAction` hanya mengizinkan `availability: unavailable` dan `reason: donation_flow_not_available`.
- `api/openapi/common.yaml` tidak memiliki component tambahan yang mengubah makna Campaign action pada path/schema ini. Operasi tetap memakai shared public response components yang didefinisikan di Campaign source.
- `api/openapi.yaml` membawa shape yang sama, dan `frontend/lib/api/generated/openapi.ts` menghasilkan tipe yang sama. `api/README.md` menetapkan split Campaign file sebagai authored source; bundle dan TypeScript adalah generated counterparts.
- `frontend/mocks/fixtures/public-campaign.ts` memakai nilai unavailable yang sama.
- `api/openapi/donation.yaml` pada `POST /campaigns/{campaignId}/donations` menyatakan guest submit tanpa Account; backend menerima hanya jika submission menang atas Campaign close ketika eligible dan menolak kondisi ineligible dengan `409`.

**Gap**

Authored Campaign API belum menggambarkan Slice-2-compatible public donation action. Generated bundle, generated TypeScript, dan fixture konsisten dengan authored source saat ini, sehingga bukan sumber independen atau perbaikan gap. Perubahan kontrak kelak memerlukan generated counterparts yang sesuai menurut `api/README.md`; belum dilakukan di Exploration.

**Sniffing**

| Lensa | Temuan |
|---|---|
| Risk | Perubahan enum/schema memengaruhi API producer, typed frontend consumer, fixtures, dan acceptance Campaign. Incomplete regeneration dapat membuat consumers membangun terhadap shape yang berbeda. |
| Edge cases | Nilai yang diterima perlu tetap merepresentasikan ketiadaan/ketidaktersediaan aksi untuk Campaign yang bukan donor-eligible, termasuk state closed; makna detail ketika eligible ditutup oleh update berikutnya harus tetap dapat dibedakan atau dipulihkan dari authority. Shape final belum diputuskan di Stage 2. |
| Miscontext | Generated types membuktikan bentuk bundle saat ini, bukan bahwa contract sudah memenuhi requirement Slice 2. |
| Misleading signals | `donation_action` sudah ada dan required di schema, tetapi satu-satunya enum adalah unavailable; sekadar melihat field ada dapat disalahartikan sebagai dukungan action contract. |
| Inconsistency | Campaign GET hanya membawa Slice-1 unavailable action, sementara Donation POST contract sudah memiliki guest submission dengan rejection `409`. Kedua source sama-sama kontraktual tetapi belum tersambung oleh Campaign entry contract. |

## Area 4 — Perilaku backend/frontend yang terkait

**Current state**

- `backend/internal/domain/campaign/repository_db.go` — `RepositoryDB.FindPublicDetail` memfilter `c.status = published` dan memilih allowlist projection.
- `backend/internal/domain/campaign/service.go` — `Service.GetPublicDetail` memetakan projection; `toPublicDetail` saat ini mengisi `DonationAction{Availability: "unavailable", Reason: "donation_flow_not_available"}`.
- `backend/internal/transport/http/campaign_public.go` — `PublicCampaignDetailHandler` memanggil `GetPublicDetail`; `toPublicCampaignDetailResponse` membawa field action tersebut ke JSON. Ini bukti perilaku Campaign GET yang ada, bukan bukti runtime Donation POST.
- `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` — `CampaignSuccess` menampilkan copy bahwa dukungan belum dapat dilakukan dan alur donasi belum tersedia. UI belum memeriksa nilai `campaign.donation_action` untuk menentukan aksi.
- WU-S2-003 (Donation Backend Delivery) masih queued di Work Graph; untuk Run ini perilaku runtime submit belum diimplementasikan/diverifikasi sebagai Slice-2 delivery. Kontrak/spec Donation adalah bukti requirement contract saja.

**Gap**

Backend Campaign GET dan frontend success view sama-sama masih merealisasikan keadaan Slice 1. Frontend tidak akan memperoleh behavior entry hanya dari fakta bahwa generated type mendeklarasikan `donation_action`; consumer/UI belum membaca semantik action. Area frontend/backend production tetap milik Work Unit delivery terpisah; WU-S2-005 memiliki scope contract/spec reconciliation.

**Sniffing**

| Lensa | Temuan |
|---|---|
| Risk | Jika UI kelak menggunakan GET sebagai satu-satunya eligibility check, race dengan close dapat mengarahkan donor ke submit yang harus ditolak. Backend POST tetap titik keputusan aktual. |
| Edge cases | Detail yang disimpan di cache/client lalu digunakan sesudah close; status detail dan submission bisa berbeda karena perubahan lifecycle. Campaign GET saat ini mengirim `Cache-Control: private, no-store`, namun state yang telah diterima client tetap bisa stale. |
| Miscontext | Predicate GET (`status = published`) menjawab apakah detail public dapat diproyeksikan saat dibaca; itu tidak membuktikan backend submit menerima Donation pada waktu berikutnya. |
| Misleading signals | Backend telah memiliki typed `DonationAction` dan response field, tetapi service selalu mengisinya unavailable; frontend hard-coded copy juga tidak bergantung pada field tersebut. Bentuk/keberadaan field belum berarti jalur aktif. |
| Inconsistency | Frontend, backend GET, fixture, generated types, dan Campaign authored API konsisten satu sama lain untuk Slice 1, tetapi semuanya tertinggal dari requirement entry Slice 2. POST contract Donation menggambarkan guest submission namun implementasi runtime terkait berada pada WU-S2-003. |

## Progression findings

### F-1 — Campaign detail entry contract belum kompatibel dengan Slice 2

- **Status:** decision-relevant untuk WU-S2-005; kontrak/API owner harus menentukan action response semantics dalam Stage 3 berdasarkan Product/MVP dan Campaign authority.
- **Dasar:** Product/MVP Slice 2 mewajibkan entry guest Donation dari eligible Public Campaign Detail; Campaign spec/API/backend/UI saat ini hanya mendukung unavailable action.
- **Yang diblokir:** WU-S2-004 frontend Techplan/Build menunggu HARD dependency WU-S2-005 yang menghasilkan Campaign/API contract reconciled dan owner-accepted.
- **Stage 3:** dapat dilanjutkan dengan aman karena sumber requirement dan gap jelas. Stage 3 dapat membingkai trade-offs/owner decision tanpa mengubah Product atau memilih shape sebelum owner memutuskan.
- **Next owner/action:** setelah Human menyetujui Stage 3, Explorer menyusun solutioning untuk keputusan Campaign/API owner Anhar Solehudin. Sesudah hasilnya, berlaku review/acceptance dan generated contract reconciliation sebagaimana manifest WU-S2-005.

### F-2 — Detail Campaign bukan otoritas submission

- **Status:** informasional / constraint yang harus dipertahankan; tidak menghalangi Stage 3.
- **Dasar:** Product dan WU manifest menetapkan backend-authoritative eligibility; Donation contract dan D1 mensyaratkan submit berurutan terhadap close. Detail GET dapat stale setelah diterima client.
- **Yang diblokir:** tidak menambah blocker WU-S2-005. Ini membatasi makna action projection dan downstream consumer; tidak menetapkan shape final.
- **Stage 3:** dapat dilanjutkan aman dengan constraint ini; jangan memindahkan keputusan eligibility ke frontend atau menyimpulkan Donation POST runtime telah diverifikasi.
- **Next owner/action:** Campaign/API owner menjaga batas di reconciliation; Donation backend delivery/Testing membuktikan submission behavior pada Work Unit dan gate yang memilikinya.

Belum ditemukan Finding yang mengharuskan Product/MVP authority berubah atau bukti tambahan sebelum Stage 3. Bentuk response/action final tetap belum diputuskan.

## Evidence / verification boundary

Analisis ini read-only terhadap production sources. Tidak ada implementation, generated artifact, test, validator, atau runtime verification yang dijalankan. Working tree sudah berisi perubahan orchestration/evidence Work Unit lain saat dispatch; file ini adalah satu-satunya file yang ditambahkan oleh Run ini.

## Anchor untuk Stage 3

- Requirement: `docs/product/mvp-delivery-slices.md` §5; `docs/product/mvp-scope.md` §5 Stages B/D.
- Campaign authority: `docs/spec/4-campaign/features/02-campaign-detail-listing.md` (`Detail`, criteria 3 dan 6); `docs/spec/4-campaign/invariants.md#inv-campaign-13`, `#inv-campaign-14`.
- Authored API: `api/openapi/campaign.yaml` (`getPublicCampaignDetail`, `PublicCampaignDetail`, `PublicCampaignDonationAction`); `api/README.md` generated-source rules.
- Donation contract boundary: `api/openapi/donation.yaml` `POST /campaigns/{campaignId}/donations`; `docs/spec/5-donation/features/01-submit-donation-settlement.md`; `docs/spec/5-donation/invariants.md#inv-donation-02`.
- Live anchors: `backend/internal/domain/campaign/repository_db.go` `RepositoryDB.FindPublicDetail`; `backend/internal/domain/campaign/service.go` `Service.GetPublicDetail` / `toPublicDetail`; `backend/internal/transport/http/campaign_public.go` `PublicCampaignDetailHandler` / `toPublicCampaignDetailResponse`; `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` `CampaignSuccess`.
- Generated/consumer correspondence: `api/openapi.yaml`; `frontend/lib/api/generated/openapi.ts`; `frontend/mocks/fixtures/public-campaign.ts`.
- Coordination/ownership: `.harscode-spaces/authority-map.md`; `WU-S2-005/manifest.md`; prior coordination route `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/EXP-S2-004-001/evidence/stage-3-solutioning.md`.
