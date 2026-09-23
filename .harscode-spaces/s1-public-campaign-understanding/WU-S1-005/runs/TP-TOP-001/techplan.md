# Tech Plan: Slice 1 Topology & Controlled Media Enablement

> Phase             : Techplan
> Ticket            : WU-S1-005 / TP-TOP-001
> Author            : Planner
> Model             : gpt-5.6-terra
> Reasoning         : medium
> Session           : Fresh session
> Created           : 2026-09-23
> Target revision   : a053b48ef3fef35a073c5fe57e5ee4581f848e07
> Workflow revision : not exposed by the invocation
> Status            : Draft
> Approach          : Perbaikan root topology yang sempit: strip prefix /api di Caddy dan tegaskan policy anonymous none hanya untuk bucket private, tanpa mengambil alih delivery Campaign.
> Refs              : EXP-TOP-001 gap analysis + solutioning; WU-S1-002/TP-001; WU-S1-002/TST-001; Caddyfile; docker-compose.yml; docs/spec/4-campaign/features/03-campaign-media.md; docs/spec/4-campaign/invariants.md (INV-campaign-14); docs/spec/4-campaign/threat-model.md; api/openapi/index.yaml; api/openapi/campaign.yaml; docs/project/kencleng-integration-map.md

---

## 1. Background

Kontrak Slice 1 menetapkan API same-origin pada base /api dan membatasi PublicCampaignMediaItem.content_url ke /api/campaigns/{campaignId}/media/{mediaId}/content. Namun Caddyfile meneruskan /api/* apa adanya, sedangkan backend native mendaftarkan rute tanpa prefix tersebut. Maka entrypoint root belum memenuhi path kontrak, walaupun Swagger development bekerja dengan mengakses backend :8090 langsung.

Compose menyediakan bucket MinIO public dan private, tetapi hanya mengeset anonymous download pada bucket public. Karena kencleng_miniodata persisten, source init tanpa assertion eksplisit tidak membuktikan policy private yang sedang berjalan. Campaign media Slice 1 harus private dan dilayani origin terkontrol; origin Campaign itu sendiri belum ada dan dimiliki WU-S1-003.

## 2. Scope

**In scope:**

- Mengubah boundary Caddy /api/* agar /api dihilangkan sebelum reverse_proxy ke backend native, dengan URL same-origin dan fallback frontend tetap.
- Menegaskan ulang policy anonymous none bagi bucket private pada setiap minio-init, termasuk volume MinIO yang dipakai ulang.
- Memperbarui setup documentation dari caveat prefix menjadi invariant executable dan batas bukti runtime.
- Menetapkan evidence source, proxy, policy persisted, dan integration yang diperlukan agar boundary topology siap dipakai Slice 1.
- Menyerahkan MINIO_BUCKET_PRIVATE sebagai handoff konfigurasi eksplisit ke WU-S1-003 tanpa memindahkan ownership Campaign.

**Out of scope (explicit):**

- Campaign persistence, handler/rute, authorization, parent/member recheck, byte streaming, MIME enforcement, Problem Details, response cache-header production, atau retraction implementation (WU-S1-003).
- Campaign UI, image consumer, mock behavior, atau Next.js rewrite (WU-S1-004).
- Perubahan OpenAPI server base, content_url, CORS, backend route-prefix, proxy response synthesis, cache override, atau redirect/signed URL.
- Menghapus/merename bucket, volume, port MinIO, atau mengubah policy kencleng-public; storage non-Campaign tetap milik owner lain.
- Hardening security-header/TLS proxy yang lebih luas. Caddyfile saat ini belum memiliki header rekomendasi best-practice; Work Unit ini tidak mengklaim menutup exposure tersebut.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | URL API browser tetap berawalan /api, tetapi backend menerima remainder path tanpa /api. | api/openapi/index.yaml servers; Caddyfile; EXP gap Area 1 |
| Q2 | Fallback non-API tetap menuju frontend native; tidak ada proxy kedua. | Caddyfile; repo setup §1/§8; EXP solutioning Decision |
| Q3 | Campaign media private; public object URL, redirect, dan long-lived signed URL bukan jalur Slice 1. | Feature 03 Summary; INV-campaign-14; threat model media |
| Q4 | kencleng-private tidak anonymous setelah minio-init, termasuk volume persisten; kencleng-public tetap download. | docker-compose.yml; EXP gap Area 2; EXP solutioning Decision |
| Q5 | MINIO_BUCKET_PRIVATE adalah handoff backend Campaign; semantic delivery tetap WU-S1-003. | .env.example; backend main initMinIO; WU-S1-003 manifest; integration map |
| Q6 | Topology tidak mengubah atau menggantikan Cache-Control: private, no-store yang diproduksi backend untuk media 200/404/503. | Feature 03; INV-campaign-14; getPublicCampaignMediaContent |
| Q7 | Setup document tidak lagi menyatakan prefix mismatch sebagai caveat aktif dan membedakan source change dari bukti runtime. | repo setup §8; EXP solutioning source boundary |
| Q8 | Setelah WU-S1-003 menarik eligibility Campaign atau membership media, fresh request melalui root Caddy pada `content_url` yang sama dan telah diketahui tidak boleh mengirim byte baru; byte yang sudah diunduh/dipegang client berada di luar jaminan ini. | Feature 03; INV-campaign-14; EXP solutioning runtime item 6; WU-S1-003 TP-BE-001 RISK-3/§13 |

## 4. Rules & Validation

- **R1 — API path translation.** Given request melalui root entrypoint yang cocok dengan /api/*, when Caddy memprosesnya, then backend menerima path yang sama tanpa /api dan URL browser tidak berubah.
- **R2 — Frontend preservation.** Given request yang tidak cocok dengan /api/*, when Caddy memprosesnya, then request tetap memakai fallback frontend native, bukan backend.
- **R3 — Private-policy convergence.** Given minio-init berjalan pada bootstrap baru atau volume kencleng_miniodata yang sudah ada, when inisiasi selesai, then anonymous policy kencleng-private adalah none dan policy download kencleng-public tetap ada.
- **R4 — Controlled boundary preservation.** Given backend WU-S1-003 menyediakan operation media, when valid content URL melewati localhost:8080, then Caddy meneruskan respons tanpa redirect, object URL, cache override, atau response synthesis; Cache-Control: private, no-store pada 200, public 404, dan eligible-dependency 503 tetap sampai client.
- **R5 — No topology ownership leak.** Given Work Unit ini selesai di source boundary, then tidak ada perubahan API contract, backend handler/domain, frontend rewrite/UI, bucket public, volume/port, atau policy storage lain; consumer Campaign memakai handoff MINIO_BUCKET_PRIVATE di WU-S1-003.
- **R6 — Documentation truth.** Given setup document dibaca setelah Build, then ia menyatakan path invariant baru serta deferred runtime evidence, bukan claim verified.
- **R7 — Retraction through proxy.** Given `WU-S1-003` menyediakan seeded eligible media dan kemudian menarik public eligibility parent Campaign **atau** membership media, when Testing melakukan fresh request melalui root Caddy (`localhost:8080`) ke `content_url` yang sama yang sebelumnya menghasilkan media, then respons tidak mengirim byte media baru dan tetap mengikuti public non-disclosure semantics backend (`404` serta `Cache-Control: private, no-store`, bukan redirect, object URL, signed URL, atau proxy-synthesized substitute). Jaminan ini tidak mencakup byte yang sudah sebelumnya diunduh atau dipegang client. `WU-S1-003` memiliki mutasi persisted eligibility/membership dan endpoint; Testing bersama WU-S1-005 memiliki evidence proxy/runtime. Ketiadaan fixture atau capability runtime backend menunda eksekusi bukti ini, tetapi tidak menghapus kewajiban R7 maupun mengizinkan klaim retraction/integrasi selesai.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Gunakan handle_path /api/* untuk backend block yang ada. | Chosen | Caddy mendokumentasikan handle_path sebagai handle dengan implicit uri strip_prefix; memetakan URL contract /api/... ke backend unprefixed tanpa mengubah URL browser. |
| D1-alt | Pertahankan handle /api/* lalu mount /api di backend. | Rejected | Menggeser mismatch root topology ke transport backend dan mengubah konvensi rute yang bukan milik Work Unit ini. |
| D2 | Tambahkan mc anonymous set none local/kencleng-private setelah bucket creation; pertahankan download pada public bucket. | Chosen | mc anonymous set mendukung policy none; assertion ini convergent terhadap persistent state dan terbatas pada bucket yang telah ditetapkan private. |
| D2-alt | Hapus/privatkan kencleng-public, buat bucket Campaign ketiga, atau sembunyikan port MinIO. | Rejected | Tidak dibutuhkan oleh Slice 1 dan mengubah deferred storage use cases lain. |
| D3 | Gunakan MINIO_BUCKET_PRIVATE yang sudah ada sebagai handoff ke WU-S1-003. | Chosen | Tidak perlu konfigurasi baru; backend owner tetap membuktikan Campaign object dibaca/ditulis di bucket ini. |
| D4 | Backend menghasilkan Problem Details dan Cache-Control; Caddy hanya meneruskan. | Chosen | Contract mensyaratkan controlled origin dan header success/error. Cache/header rewriting menambah scope dan bisa merusak retraction semantics. |
| D4-alt | Ubah OpenAPI/content URL, buat Next.js proxy, gunakan direct/signed object URL, atau tambah proxy cache rules. | Rejected | Bertentangan dengan same-origin owner atau private-origin/recheck/retraction boundary. |
| D5 | Jangan menambahkan security headers/TLS. | Chosen — scoped out | Ini concern hardening terpisah, bukan prerequisite path/policy fix; perubahan ini tidak boleh dianggap remediation-nya. |

## 6. Backward Compatibility

- Base URL public tetap /api. Perubahan mengoreksi proxy agar cocok dengan contract dan backend route unprefixed; backend native langsung tetap memakai path unprefixed.
- Nama kedua bucket, kencleng_miniodata, exposed ports, dan download policy public bucket tidak berubah. Tidak ada migration atau operasi data destruktif.
- Tidak ada perubahan OpenAPI, generated types, atau response schema. content_url tetap opaque same-origin reference.
- Assertion private policy dapat menghapus anonymous access historis pada hanya kencleng-private. Ini koreksi security requirement, bukan perubahan object data.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| RISK-1 | Prefix tidak di-strip atau fallback menangkap API. | Medium | High | R1 source/config validation dan test root GET /api/healthz; juga test frontend. |
| RISK-2 | Persistent volume menyimpan private policy salah atau initializer belum selesai. | Medium | High | R3 inspeksi running persisted policy kedua bucket, termasuk volume reuse. |
| RISK-3 | Campaign object masuk public bucket, bypass origin checks, atau known `content_url` tetap mengirim byte baru setelah retraction. | Medium | High | D2+D3; WU-S1-003 select/mutates private persisted Campaign state; Testing membuktikan direct anonymous object request ditolak dan R7 fresh fetch melalui root Caddy setelah eligibility/member withdrawal tidak mengirim byte baru. |
| RISK-4 | Proxy error/header menyamarkan 404/503 backend atau membuang no-store. | Medium | High | R4 end-to-end setelah endpoint backend tersedia; pisahkan upstream failure Caddy dari response contract. |
| RISK-5 | Source configuration dianggap runtime/integrated verification. | High | Medium | R6 dan Open Items; jangan claim BACKEND_VERIFIED, FRONTEND_MOCK_VERIFIED, atau INTEGRATED_VERIFIED. |
| RISK-6 | Caddy belum memiliki hardening security headers/TLS umum. | Existing | Medium | Explicitly outside scope (D5); route sebagai work unit terpisah bila diprioritaskan manusia. |

## 8. Interface Contract

**Persistence/data shape:** Tidak ada schema atau migration. Compose tetap mengelola kencleng_miniodata; yang berubah hanya desired anonymous policy kencleng-private setelah init.

**API/event/external interface:** Base URL public tetap /api. Caddy mengubah upstream request path dari /api/<remainder> menjadi /<remainder>. Tidak ada operation atau response baru. Media content_url tetap opaque dan hanya controlled endpoint pattern yang diizinkan.

**Cross-layer/business boundary:** Root topology meneruskan path dan response, bukan authorization point Campaign. WU-S1-003 memakai MINIO_BUCKET_PRIVATE dan memiliki eligibility/member recheck, allowed MIME, 404/503, no-store, dan retraction. WU-S1-004 hanya mengonsumsi content_url tanpa URL MinIO atau production proxy lain.

## 9. Architecture / Plan

1. Re-open Caddyfile dan ganti hanya API handler menjadi handle_path /api/*; upstream host.containers.internal:8090 serta fallback handle frontend tidak berubah.
2. Re-open docker-compose.yml dan, setelah dua mc mb -p yang ada, pertahankan mc anonymous set download local/kencleng-public lalu tambahkan mc anonymous set none local/kencleng-private. Jangan ubah credentials, aliases, bucket/volume/ports, atau lifecycle service.
3. Re-open repo setup §8: ganti caveat dengan root-to-backend path invariant, nyatakan ownership root, dan rujuk bukti runtime/persisted policy yang tersisa.
4. Jalankan source-level checks yang tersedia. Caddy/Compose runtime dan integration evidence hanya dilakukan pada environment dengan Compose runtime, backend native, dan media fixture/operation WU-S1-003; evidence itu mencakup R7 fresh retraction fetch melalui root Caddy, bukan hanya fetch eligible awal.

Tidak ada runbook terpisah: tiga perubahan source dan verifikasi topology adalah satu lifecycle enablement.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| Caddyfile — API handle /api/* block | Satu-satunya root path boundary; fallback frontend sibling handle. | Ganti keyword block menjadi handle_path dengan matcher/upstream sama. Caddy handle_path menjadi precedent prefix stripping. |
| docker-compose.yml — minio-init entrypoint | Membuat dua bucket dan sudah mengeset public download; kencleng_miniodata membuat policy persistent. | Tambahkan mc anonymous set none local/kencleng-private setelah creation; pertahankan public semantics. |
| docs/project/kencleng-repo-setup.md — §1, §8 | Menjelaskan topology dan caveat yang diselesaikan. | Perbarui §8 dengan invariant executable dan deferred runtime evidence; sinkronkan §1 hanya jika prose path perlu. |
| .env.example — MINIO_BUCKET_PRIVATE | Nama configuration native backend yang menjadi storage rendezvous. | Read-only handoff; jangan tambah variable atau nilai credential ke artifact. |
| backend/cmd/server/main.go — initMinIO, root ServeMux | Bucket hanya diverifikasi ada dan path backend tidak memakai /api. | Re-open sebelum final runtime verification; tidak dimodifikasi oleh WU ini. |
| api/openapi/index.yaml — servers; api/openapi/campaign.yaml — media operation, item, cache responses | Contract exact base path, controlled route, dan header yang harus dilewati topology. | Read-only oracle R1/R4; tidak diubah. |
| WU-S1-003/manifest.md — Scope/coordination | Menetapkan backend controlled delivery owner. | Handoff D3; root Build tidak membuat Campaign storage/handler. |

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| Caddyfile | Modify | Strip /api hanya untuk API sebelum proxy backend. |
| docker-compose.yml | Modify | Assert anonymous none pada kencleng-private; pertahankan public policy. |
| docs/project/kencleng-repo-setup.md | Modify | Ganti caveat proxy dengan invariant dan batas runtime evidence. |

| File / area intentionally untouched | Why |
|---|---|
| api/openapi/**, frontend/lib/api/generated/** | Contract dan types sudah menetapkan /api/opaque URL; tidak ada contract change. |
| backend/** | Campaign storage/delivery dan route registration milik WU-S1-003. |
| frontend/** | Tidak ada UI/consumer/rewrite dalam Work Unit topology. |
| .env.example | Existing MINIO_BUCKET_PRIVATE cukup sebagai handoff; tidak menambah configuration. |
| Bucket names, volume declarations, ports, kencleng-public policy | Di luar scope dan diperlukan deferred use cases lain. |
| .harscode-spaces/** manifest/control surface | Status orchestration dimiliki Orchestration Operator. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | Review diff: hanya API block menjadi handle_path /api/*. Di environment Compose, validate Caddy config lalu dengan backend native hidup GET localhost:8080/api/healthz membuktikan backend menerima /healthz. | Build for source review; Testing for runtime | Source review tidak membuktikan parsing/translation actual. |
| R2 | Dengan Caddy/backend/frontend native hidup, request non-API localhost:8080 harus dijawab frontend, bukan backend. | Testing | Perbaikan API tidak boleh merusak root UI entrypoint. |
| R3 | Validate rendered Compose bila tooling tersedia. Pada first boot dan setelah initializer terhadap kencleng_miniodata reuse, gunakan authenticated mc anonymous get untuk observasi kedua policy: private none, public download. | Testing | Policy source tidak membuktikan state persisted; salah policy membuka bypass atau merusak public use case. |
| R4 | Setelah WU-S1-003 menyediakan seeded eligible media operation, fetch contract URL melalui localhost:8080; cek JPEG/PNG 200 dan exact Cache-Control. Uji absent/non-public/non-member identical backend 404 body/header; induce eligible storage/object failure untuk backend 503 + header. Pisahkan Caddy upstream failure dari response contract. | Testing coordinated with WU-S1-003 | Membuktikan proxy tidak mengganti header/error. Tanpanya retraction/anti-enumeration dapat gagal pada browser entrypoint. |
| R5 | Review git diff --check dan changed-file scope; re-open WU-S1-003 manifest dan .env.example. Di integration environment, direct anonymous request ke seeded private Campaign object harus ditolak dan public response tidak memuat object URL, redirect, atau signed URL. | Build for scope; Testing with WU-S1-003 for security observable | Menjaga root work tidak menyerap domain work dan policy menutup bypass origin. |
| R6 | Inspect §8 update terhadap executable Caddyfile/docker-compose.yml; document menyebut path translation dan deferred evidence, bukan claim verified. | Build | Mencegah caveat usang dan false completion signal. |
| R7 | Bersama WU-S1-003, seed satu Campaign eligible dengan media yang `content_url`-nya dicatat. Fetch URL tersebut sekali melalui `localhost:8080` untuk membuktikan precondition. Kemudian WU-S1-003 menarik **salah satu per skenario**: public eligibility parent Campaign atau membership media yang sama pada persisted state. Tanpa memakai respons/cache client sebelumnya, lakukan fresh request baru melalui root Caddy ke `content_url` yang persis sama; assert tidak ada byte JPEG/PNG baru yang delivered, `404` public non-disclosure backend dan `Cache-Control: private, no-store` diteruskan, serta tidak ada redirect/object URL/signed URL. Ulangi untuk jalur withdrawal lain bila fixture mendukungnya. Catat bahwa byte yang telah diunduh/dipegang client bukan objek jaminan. | WU-S1-003 untuk fixture dan mutasi persisted eligibility/membership; Testing + WU-S1-005 untuk eksekusi/evidence Caddy/proxy | Hanya recheck backend pada fresh browser-facing request yang membuktikan known URL tidak bertahan melalui proxy/cache setelah withdrawal; source inspection atau direct-backend test tidak cukup. Bila fixture/capability runtime belum tersedia, tandai eksekusi deferred/not tested tanpa menghapus R7 atau mengklaim retraction/integrasi verified. |

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Private bucket / direct object bypass / retraction | Anonymous read melewati parent/member origin check dan retraction; known URL perlu dibuktikan kembali setelah withdrawal. | EXP-TOP-001/evidence/solutioning.md#runtime-verification-required-not-performed-here items 2, 5–6 | Yes — R3/R5/R7; retraction fresh-fetch melalui Caddy memerlukan endpoint dan mutasi persisted state WU-S1-003. |
| Proxy cache and public error preservation | Header/error substitution dapat meniadakan no-store, 404 anti-enumeration, atau 503 distinction. | EXP-TOP-001/evidence/gap-analysis.md#stage-2-carry-forward-evidence | Yes — R4. |
| Same-origin API path boundary | Salah prefix membuat security/delivery logic backend tidak tercapai. | EXP-TOP-001/evidence/gap-analysis.md#area-1--root-caddy-same-origin-api-boundary | Yes — R1/R2. |

## 13. Open Items

### Active — needs external input or verification

1. **Compose/Caddy/MinIO runtime evidence unavailable in this session.** Environment Compose-capable harus validate rendered Caddy config dan inspeksi policy pada volume kencleng_miniodata sebenarnya setelah init. Owner: Testing/environment operator. Ini tidak menghalangi narrow source Build, tetapi menghalangi claim topology runtime complete.
2. **Controlled-media integration fixture belum ada.** WU-S1-003 harus menyediakan backend operation dan seeded eligible/non-eligible/media states sebelum R4/R5 end-to-end, no-store, direct-bypass, dan retraction checks dapat berjalan. Owner: WU-S1-003 untuk capability; Testing untuk joint evidence. Ini tidak mengizinkan root Build menambah Campaign code.
3. **R7 retraction-through-proxy evidence belum dapat dieksekusi tanpa capability WU-S1-003 dan Compose runtime.** WU-S1-003 harus menyediakan fixture yang dapat menarik persisted public eligibility parent dan membership media setelah `content_url` diketahui; Testing bersama owner root topology harus melakukan fresh request ke URL yang sama via `localhost:8080`. Owner mutasi: WU-S1-003; owner evidence proxy/runtime: Testing dan WU-S1-005. Sampai itu tersedia, R7 adalah deferred/not tested, bukan kewajiban yang hilang dan bukan dasar untuk klaim retraction atau integrasi selesai; already-downloaded/client-held bytes tetap di luar jaminan kontrak.

### Resolved — retained as decision history

1. ~~**Owner /api correction dan private policy.**~~ **RESOLVED — root topology memiliki Caddy path translation dan Compose policy assertion; Campaign delivery semantics tetap WU-S1-003.** Direkam EXP-TOP-001 dan D1–D4.
2. ~~**Apakah bucket/configuration baru diperlukan.**~~ **RESOLVED — MINIO_BUCKET_PRIVATE yang ada adalah handoff; hanya anonymous none policy-nya yang ditegaskan.** Bucket/variable redesign tidak diotorisasi.
