# CR-001 — Review Findings 1

> Phase             : Code Review
> Work Unit         : WU-S1-002
> Run               : CR-001
> Author            : Reviewer
> Created/Updated   : 2026-09-22T10:36:27+07:00
> Model             : `gpt-5.6-terra`
> Reasoning         : `high`
> Session           : Fresh independent session
> Build diff reviewed: `e31e23b60ef2a8f994db623f8943a4eac3595e7e..5fb58b2de8f2bf9882fa6c1713744feb899dc8c7`
> Repository head inspected: `2aefd7826a4862cc15a1b2192b926f04a7e7422b`
> Workflow revision : not exposed by the CR-001 invocation

Scope review memakai diff implementasi dari baseline Build hingga `BLD-002`. Commit setelahnya hanya merutekan Run `CR-001`; commit tersebut turut diperiksa sebagai konteks current repository, tetapi bukan perubahan delivery yang dinilai di bawah.

## 1. Safety

### CR-001-F01 — `content_url` belum membatasi direct object-storage URL

- **Location:** `api/openapi/campaign.yaml:821-825`; tercermin di `api/openapi.yaml:3246-3250` dan `frontend/lib/api/generated/openapi.ts:4180-4185`.
- **Problem:** `PublicCampaignMediaItem.content_url` hanya bertipe `string` dengan `format: uri-reference`. Format tersebut masih menerima URI absolut, sehingga URL seperti `https://bucket.example/object.png` tetap valid menurut schema. Deskripsi dan contoh menyatakan same-origin controlled route, tetapi keduanya tidak membatasi nilai kontrak; generated TypeScript akibatnya hanya menjadi `string`.
- **Why it matters:** producer dapat mengembalikan direct bucket/signed URL yang lolos contract validation, melewati parent/member recheck di origin, serta melemahkan jaminan retraction dan `private, no-store`. Ini bertentangan dengan R7–R10 Techplan dan `INV-campaign-14` yang melarang URL object storage/redirect publik.
- **Suggested resolution:** pada authored `content_url`, tambahkan `pattern` yang hanya menerima path same-origin `/api/campaigns/<uuid>/media/<uuid>/content` (tanpa scheme, host, query, atau fragment). OpenAPI 3.0 tidak perlu dan tidak dapat menyatakan bahwa kedua UUID harus identik dengan sibling field/path parameter; pencocokan parent/member tetap merupakan kewajiban runtime. Regenerasi bundle dan frontend types, lalu lakukan targeted negative check bahwa URL absolut dan path di luar controlled route tidak valid.
- **Blocking:** **Blocking.**

Tidak ada Finding Safety lain: diff tidak menambah runtime backend/frontend, state bersama, I/O lifecycle, atau error path baru. Concern runtime anti-enumeration, timing parity, storage recheck/retraction, dan cache preservation sudah tetap tercatat sebagai bukti downstream pada Test Focus Pointer; tidak ada Techplan drift baru dari Review ini.

## 2. Quality

### CR-001-C01 — Label referensi threat-model sudah stale

- **Location:** `docs/spec/4-campaign/features/02-campaign-detail-listing.md:87-88` dan `docs/spec/4-campaign/features/03-campaign-media.md:74`.
- **Problem:** referensi masih menyebut heading lama `"Public campaign listing & detail"` dan `"Campaign media"`, sementara heading yang direkonsiliasi sekarang adalah `"Slice-1 public Campaign detail"` dan `"Slice-1 Campaign media delivery"`.
- **Why it matters:** navigasi evidence lintas feature/threat model menjadi kurang presisi, terutama karena listing dan attachment-list memang sudah `DEFER`.
- **Suggested resolution:** perbarui label referensi saat ada perubahan dokumentasi Campaign berikutnya.
- **Blocking:** Non-blocking. Tidak dimasukkan ke patch loop ini.

Selain itu, perubahan tetap proporsional: tidak ada duplikasi logic/runtime, dead code, atau observability requirement baru yang relevan untuk reconciliation-only diff.

## 3. Stack-Specific Best Practices

Sumber yang diterapkan secara terarah:

- `../harscode-workspace/best-practices/restapi/anti-enumeration.md` — kedua public GET mendeklarasikan `security: []`, memakai response `404` yang sama, dan tidak menambahkan `403`; timing parity tetap tepat didefer ke runtime Testing.
- `../harscode-workspace/best-practices/restapi/openapi-spec-first-drift.md` — split source, bundled `api/openapi.yaml`, dan `frontend/lib/api/generated/openapi.ts` sinkron melalui independent regeneration; success dan Problem error response ikut terdefinisi.
- `../harscode-workspace/best-practices/pwa/xss-and-content-sanitization.md` — kontrak hanya membawa plain text dan tidak menambah renderer/unsafe HTML path.
- `../harscode-workspace/best-practices/pwa/service-worker-caching.md` — tidak ada service-worker change; response success/error public secara konsisten mengontrak `Cache-Control: private, no-store`.

Tidak ada Finding tambahan dari guidance tersebut. CR-001-F01 tetap blocking karena ketidakmampuan schema membatasi direct URL melemahkan boundary media yang guidance cache/security tersebut asumsi-kan.

## 4. Consistency

- Perubahan delivery tidak menyentuh `docs/product/**`, `docs/ui-ux/**`, backend runtime, frontend runtime, atau topology; ini konsisten dengan scope TP-001 dan fencing root `AGENTS.md`.
- `getPublicCampaignDetail` dan `getPublicCampaignMediaContent` memakai `security: []`, response `404`/`503` Problem Details yang sama, dan header no-store pada success/error. Semua exact-projection object yang diwajibkan TP-001 tetap `additionalProperties: false`; tidak ada inheritance `Campaign`/`Organization` di public graph.
- `GET /campaigns/{campaignId}/attachments` sudah dihapus, sementara upload historis tetap eksplisit `DEFER`; tidak ada route runtime baru.
- `frontend/tsconfig.json` hanya menambah `types: ["vitest/globals"]`, tepat sesuai TP-002 dan tidak mengubah runtime/test semantics.
- Tidak ada perubahan tracker pada diff Build yang direview, sehingga tidak ada claim `CONTRACT_READY` atau milestone backend/frontend/topology/integration yang berlebihan.

Dengan pengecualian CR-001-F01 dan komentar referensi non-blocking CR-001-C01, diff konsisten dengan TP-001/TP-002, `AGENTS.md`, `api/README.md`, dan `frontend/AGENTS.md`.

## Verification executed during Review

- `git diff --check e31e23b..5fb58b2` — memastikan hygiene diff Build; **pass**.
- Independent regeneration ke temporary directory dengan `redocly bundle openapi/index.yaml` lalu `openapi-typescript` — menjawab apakah bundle/types committed stale atau hand-edited; kedua `cmp` **pass** sebelum static parser optional gagal karena module `yaml` tidak terpasang. Kegagalan parser tidak mengubah hasil regeneration; inspection langsung pada authored dan bundled YAML dipakai untuk remaining schema review.
- Inspeksi terarah authored source, dereferenced bundle, dan generated declarations — menemukan `content_url` mempunyai `format: uri-reference`, tidak mempunyai `pattern`, dan generated menjadi `string`; mengonfirmasi CR-001-F01.
- Tidak menjalankan full Build/Testing matrix karena Code Review reasoning-first dan command tersebut tidak diperlukan untuk membuktikan Finding.

## Verdict

**Request changes.**

Blocking: **CR-001-F01**. CR-001-C01 adalah komentar non-blocking dan tidak perlu menahan handoff.

## Phase handoff

- **Completed:** empat pass Review terhadap diff Build dan current repository context.
- **Artifacts:** `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/CR-001/review-findings-1.md`; `.harscode-spaces/s1-public-campaign-understanding/WU-S1-002/runs/CR-001/patch-plan-1.md`.
- **Human decision:** none.
- **Open / deferred:** CR-001-F01 blocking; CR-001-C01 non-blocking. Runtime media recheck/retraction, timing parity, cache preservation, decimal calculation, and safe text rendering remain downstream implementation/Testing evidence.
- **Recommended next step:** Build/Patch memakai patch plan spesifik untuk CR-001-F01, lalu lakukan targeted confirmation; full four-pass re-review tidak diperlukan bila patch hanya mengencangkan schema, meregenerasi artifacts, dan tidak memperluas behavior.
- **Session transition:** kembali ke fresh Build/Patch session yang re-grounded pada patch plan karena perubahan membutuhkan authored contract + generated artifacts; tidak ada healthy existing Build session yang harus dipertahankan.
- **Context pointers:** TP-001 §8/R7–R10, TP-002, `api/openapi/campaign.yaml:813-838`, CR-001-F01, dan generated bundle/types anchors di atas.
