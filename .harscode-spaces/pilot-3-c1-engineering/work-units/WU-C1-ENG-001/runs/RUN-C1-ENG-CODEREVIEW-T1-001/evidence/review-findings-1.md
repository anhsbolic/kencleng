# Review findings — RUN-C1-ENG-CODEREVIEW-T1-001

> Phase: Independent Code Review — T1 shared contract; Author / Participant: `PARTICIPANT-C1-ENG-CONTRACT-REVIEWER-001`; Role: Reviewer; Specialization: Shared API contract; Created/Updated: 2026-10-09. Work Unit: `WU-C1-ENG-001`; Run: `RUN-C1-ENG-CODEREVIEW-T1-001`; Session: `SESSION-C1-ENG-CONTRACT-REVIEWER-001` (binding invocation, bukan ID runtime). Model / Reasoning: `gpt-6-sol` / `medium` menurut dispatch invocation; runtime model/effort tidak diverifikasi independen. Target baseline dan observed HEAD: `112c35f8322b53bac39d511ed4f22982cea619af`; branch: `pilot/3-c1-engineering`; workflow: `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`, tree bersih saat entry.

## Scope dan provenance

Empat pass dijalankan berurutan terhadap target yang sama: tracked diff dari baseline di atas untuk `api/openapi/index.yaml`, `api/openapi.yaml`, `api/package.json`, `api/README.md`, serta isi penuh dua file baru `api/openapi.d.ts` dan `api/.gitignore`. Source, bundle, dan seluruh generated types diperiksa langsung. Build report hanya orientation dan evidence tooling terdahulu. Tidak membaca raw Exploration, tidak menulis implementation atau owning artifacts, dan tidak dispatch downstream.

Hash berikut cocok dengan invocation pada entry dan kembali cocok pada terminal check; tidak ada material drift. Working-tree WU/events/Space/Control Tower/invocation changes yang sudah ada berada di luar implementation review dan dipertahankan.

| Target / input | SHA-256 entry = terminal |
|---|---|
| `api/openapi/index.yaml` | `8ff663e975c9d72e0d6b636994533f420fe77639398197e6137fae49e59fb62d` |
| `api/openapi.yaml` | `9e20887011ad30a3c774edd0fb725659bccd28446fe5e4aa70c0216fbcb7a4e9` |
| `api/openapi.d.ts` | `a0ae33bccdf194190f3dbb49d69435d983acb8cc28fd712b3a2869442dbe6e6e` |
| `api/package.json` | `b3ef5560aa7d58f5320a74dbbdbcbe9eadc0d8431b5459667e8c5605b4a1e954` |
| `api/README.md` | `fd769b04dacdbd3340b95f83e5df3d3908315812cd585ef68cbcfd78c07b5cf5` |
| `api/.gitignore` | `4d56952b0fb13bf8f9b6c13a6d4c34a075bac3af447636a1df4335d7576e2f97` |
| `techplan/techplan.md` | `bd732748b486f0327f6cba2cdf14cab32f1ece3c5be8ebb24f9659444947fa47` |
| `techplan/tasks/T1-shared-contract.md` | `bbd5ab13c882ca3534ab3397ca908f70157ad41c9a81ee7cafd2827dd94946e3` |
| `techplan/tasks/manifest.md` | `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d` |
| `runs/RUN-C1-ENG-BUILD-T1-001/evidence/build-report.md` | `1be1b5ad607e3ad7f1a031a65def52d60ba95d33f6d2c1d9959c8619f0b71137` |

Path Techplan/task/manifest/Build pada tabel relatif terhadap WU directory. WU `events.md` dibaca untuk exact Techplan approval, penerimaan split, G1–G3, dan model approval Run ini. Keputusan settled tetap efektif; review tidak meminta approval ulang.

## 1. Safety

### F1 — Blocking / P1: cookie authentication contract tidak cocok dengan production cookie yang approved

- **Location:** `api/openapi/index.yaml:220`–225 (`sessionCookie`), callback `Set-Cookie` pada :56–58, logout :91–93; bundle membawa deklarasi yang sama.
- **Problem:** Satu security scheme global menyatakan cookie bernama `kencleng_session` tanpa membatasi ke development. Approved R10 menetapkan production host-only `__Host-kencleng_session` dengan HttpOnly, Secure, SameSite=Lax, Path=/, dan cookie insecure yang berbeda hanya untuk explicit development localhost:8080. Tidak ada penjelasan environment/flags ini di kontrak atau header issuance/expiry.
- **Impact:** Implementasi/tooling yang mengikuti security scheme mengirim/menerima nama cookie yang salah di production, atau memakai cookie tanpa perlindungan host-prefix/flags yang disetujui. Deklarasi auth T1 menjadi bertentangan dengan owning security boundary; ini bukan bukti bahwa runtime saat ini sudah vulnerable.
- **Resolution:** Deklarasikan production cookie dengan nama/flags/host-only semantics R10, dokumentasikan pengecualian development yang terpisah dan terbatas, dan selaraskan issuance serta expiry/logout. Jangan menjadikan kedua cookie authority yang dapat dipakai bersamaan di production.
- **Authority:** Approved Techplan §4 R10, §8 HTTP / external boundary; T1 Scope dan authority R10/R13; root `AGENTS.md` Security scrutiny / authority separation. Matching stack guidance: `../harscode-workspace/best-practices/restapi/csrf-and-cookie-security.md`.

### F2 — Blocking / P1: ordinary mutation browser boundary belum lengkap, logout menghasilkan body `never`

- **Location:** `api/openapi/index.yaml:78`–95 (logout), :96–121 (prepare), :140–182 (confirm), :230–235 (`CsrfToken`); `api/openapi.d.ts`, `operations.logout.requestBody?: never`.
- **Problem:** R10 mewajibkan exact configured Origin + JSON + session-bound CSRF untuk mutations dan tanpa credentialed CORS. Logout hanya mendeklarasikan same-origin/CSRF dan tidak mempunyai JSON request contract; tipe generated menyatakan request body `never`. Prepare/confirm mendeklarasikan JSON dan header CSRF, tetapi kontrak tidak menyatakan exact Origin policy, session binding token, atau larangan credentialed CORS. Jadi consumer yang mengikuti generated logout shape dapat mengirim bodyless POST tanpa JSON sementara backend R10 wajib menegakkan JSON boundary.
- **Impact:** T2 dan T5 dapat menghasilkan protocol yang berbeda untuk logout dan berbeda dalam penegakan browser boundary. Header CSRF saja tidak menyatakan seluruh authorized rule. Tidak ada klaim bypass runtime telah direproduksi.
- **Resolution:** Nyatakan shared rule untuk ketiga ordinary mutations: exact configured Origin, application/json, token yang terikat sesi, penolakan sebelum processing, tanpa credentialed CORS. Buat wire declaration logout dapat dipakai consumer JSON secara konsisten (misalnya closed empty JSON object, tanpa authority fields); cantumkan invalid-request response yang applicable. Jelaskan start/callback memakai browser-bound state/nonce/PKCE, bukan ordinary mutation Origin/CSRF rule. Jangan meminta frontend menyetel browser-owned Origin secara manual.
- **Authority:** Approved Techplan §4 R10, §8 Cross-layer/security boundary dan logout row; D9; T1 Scope. Matching guidance: `restapi/csrf-and-cookie-security.md`, `restapi/cors-configuration.md` di current Harscode best-practices.

**Safety yang tetap sesuai:** Strict closed prepare/confirm bodies tidak menerima person/role/target/override; global session auth dengan explicit anonymous start/callback; same-person preparation/confirm dan Owner object-scoped read dinyatakan; unknown/outside-person sama 404. Preparation tidak memberi Organization/Owner. Guard closed/missing/unreadable menolak new grant; historical replay/read tetap tersedia setelah hold/expiry. Frozen name/version dan distinct unknown-outcome instructions mempertahankan receipt ID tanpa blind replacement. OrganizationView tidak mengubah provenance atau mengklaim review/verification/curation.

Tidak ada executable concurrent/resource/network core dalam target ini, sehingga nil/race/resource-release/cancellation checks tidak applicable sebagai runtime review. Test Focus Pointer §12 sudah mencakup session/browser, IDOR, guard race, atomicity dan unknown outcome; tidak ditemukan specialized security concern baru yang membutuhkan Techplan drift finding.

## 2. Quality

Tidak ada finding tambahan. Komponen schemas/responses/headers digunakan bersama, operation IDs jelas, preparation-state union mendiskriminasi state, nullable `next_cursor` menjadi `string | null`, dan seluruh OrganizationView effects menjadi literal false. Duplikasi receipt fields pada prepared state tidak menjadi defect material pada OpenAPI 3.0 closed-object composition.

Tooling reuse tidak menambah dependency: `api/package-lock.json` mengunci Redocly 2.47.0; `frontend/package-lock.json` mengunci openapi-typescript 7.13.0, sesuai dependency frontend yang sudah ada. README menjelaskan provisioning frontend dan lokasi shared generated output. Inspection menunjukkan package script memakai existing generator; reproduksibilitas command berasal dari Build evidence, tidak diuji ulang di Review. `.gitignore` baru hanya `node_modules/`.

## 3. Stack-Specific Best Practices

Routing melalui `best-practices/AGENTS.md`, targeted REST/OpenAPI/idempotency/pagination/cookie/CSRF/CORS entries serta authn/authz Security Concern Map. Matching files dibaca: `restapi/csrf-and-cookie-security.md`, `cors-configuration.md`, `openapi-spec-first-drift.md`, `pagination-and-status-codes.md`, `idempotency-and-versioning.md`, `anti-enumeration.md`, `go/authorization-and-idor.md`, `go/role-and-privilege-separation.md`, `go/jwt-and-token-lifecycle.md`. Local JWT signing-key/refresh-family guidance tidak applicable untuk opaque cookie/no-refresh D9; DB privilege/runtime IDOR evidence tetap downstream-owned.

F1/F2 juga melanggar routed cookie/CSRF/CORS concern; tidak dihitung sebagai finding baru pada pass ini.

### F3 — Blocking / P2: callback failure mendeklarasikan 400 sebagai redirect dan tidak mencakup provider denial input

- **Location:** `api/openapi/index.yaml:45`–63, khususnya required `code` :48 dan `400` dengan `Location` :59–63; `api/openapi.d.ts`, `operations.completeGoogleAuthentication` query/responses.
- **Problem:** Callback invalid dijelaskan sebagai fixed local error redirect, tetapi HTTP statusnya 400 dengan header Location dan tanpa body. 400 tidak memberikan browser redirection semantics. Selain itu query contract mewajibkan code dan hanya mendeklarasikan code/state; provider denial callback `?error=access_denied&state=...` tidak dapat direpresentasikan sebagai legitimate protocol error branch.
- **Impact:** Implementer yang mengikuti failure response dapat meninggalkan browser pada callback URL/error page, bertentangan dengan fixed local error destination; consumer/tooling tidak memiliki bentuk callback cancellation/denial yang benar. Failure/no-session rule sudah tepat, tetapi transport declaration tidak memenuhi flow yang disetujui.
- **Resolution:** Nyatakan failure melalui redirect status yang benar ke fixed local safe error destination (dapat memakai 302 yang sudah ada), tanpa session issuance dan tanpa merefleksikan code/state/provider error. Modelkan success code versus provider error query branches dan tetap tolak missing/mismatched browser/state dengan aman. Jangan mewajibkan code pada legitimate provider-error branch, dan jangan mengubah failure menjadi sesi atau menambahkan fabricated 2xx.
- **Authority:** Approved Techplan R9/R10, §8 callback row dan fixed local callback error destination; T1 callback constraints. Matching Harscode: `best-practices/restapi/pagination-and-status-codes.md`, `openapi-spec-first-drift.md`. Primary protocol check: [RFC 9110 §15.4.3](https://www.rfc-editor.org/rfc/rfc9110.html#section-15.4.3) / [§15.5.1](https://www.rfc-editor.org/rfc/rfc9110.html#section-15.5.1) membedakan redirect dengan bad request; [RFC 6749 §4.1.2.1](https://www.rfc-editor.org/rfc/rfc6749.html#section-4.1.2.1) mendefinisikan provider denial memakai error/state, bukan code.

### F4 — Blocking / P2: invalid read-input response hilang dari operations dan generated types

- **Location:** `api/openapi/index.yaml:122`–139, :183–200, :201–218; `api/openapi.d.ts`, responses pada `getOrganizationEstablishment`, `listMyOrganizations`, `getOrganization`.
- **Problem:** §8 menetapkan 400 `invalid_request` untuk invalid path, dan source mendeklarasikan UUID path parameters serta integer limit minimum 1. Namun ketiga reads ini tidak mendeklarasikan 400. Misalnya authenticated read dengan malformed UUID atau `limit=0` tidak mempunyai response invalid-input yang dinyatakan operation. Shared `InvalidRequest` component saja tidak menambahkan status ini pada operations/types.
- **Impact:** Backend yang memenuhi approved invalid-path rule tidak cocok dengan operation-specific generated response shapes; typed clients/fixtures kehilangan expected failure handling. Malformed input harus dapat dibedakan dari valid unknown/outside-person UUID yang tetap identical 404.
- **Resolution:** Tambahkan existing private/no-store safe 400 InvalidRequest pada applicable reads dan dokumentasikan invalid cursor/limit handling sesuai R13. Pertahankan valid unknown/outside-person 404. Regenerate bundle/types, jangan memperbaiki dengan handwritten consumer model.
- **Authority:** Approved Techplan §8 HTTP/code table (`400 invalid_request`, invalid path), R13; T1 safe errors/status meanings. Matching Harscode: `best-practices/restapi/openapi-spec-first-drift.md`, `pagination-and-status-codes.md`.

Tidak ada defect baru dalam preparation-ID idempotency, historical replay, stable timestamp/UUID pagination declaration atau anti-enumeration 404 declaration. Numeric page cap/default masih ordinary Build freedom; bounded page-size obligation telah dinyatakan, sehingga tidak diangkat sebagai keputusan baru.

## 4. Consistency

F1/F2/F3/F4 adalah ketidaksesuaian terhadap Approved Techplan/T1 yang disebut di masing-masing finding. Root `AGENTS.md` Authority separation mewajibkan engineering mengikuti scoped owning artifacts; upstream tidak diubah untuk menyesuaikan API. Tidak ada contradiction Product yang perlu direkonstruksi dari sejarah: aturan yang diperlukan sudah eksplisit.

Tidak ada finding tambahan terhadap current repo conventions: `api/README.md` satu editable source; `docs/project/kencleng-backend-tech-stack.md` API relationship mewajibkan derived views; `frontend/AGENTS.md` API and mocks serta `docs/project/kencleng-frontend-tech-stack.md` API boundary mewajibkan generated types, bukan parallel handwritten models. API source dan bundle secara parsed structure sama; types membawa shape/literal/nullable/parameter/response yang dinyatakan source, termasuk defects yang dicatat di atas. Prefix `/api` sesuai §8, backend paths tidak mengandung prefix ganda. Tidak ada backend/frontend tracked changes pada pemeriksaan scope.

Tiga Redocly warnings pada Build report dinilai sebagai context: license belum ditetapkan oleh authority, dua auth operations memang redirect-only. Tidak ada finding karena warning itu sendiri, tidak ada usulan license baru, fabricated 2xx, ignore atau pelonggaran lint. F3 adalah defect error redirect tertentu yang independen dari warning generic 2xx.

## Verification executed during Review

1. **Provenance/drift:** `git rev-parse HEAD`, `git branch --show-current`, `git -C ../harscode-workspace rev-parse HEAD`, `git -C ../harscode-workspace status --short`, dan `sha256sum` seluruh sepuluh target/input pada tabel. Hasil entry/terminal cocok; workflow tree clean. Explicit `git diff 112c35f8322b53bac39d511ed4f22982cea619af -- api/openapi/index.yaml api/openapi.yaml api/package.json api/README.md` dan full reads dua file baru menjadi actual review scope. `git diff --name-only -- backend frontend` tidak mengeluarkan path.
2. **Targeted read-only structure check:** Pertanyaan: apakah source/bundle mempunyai semantic drift, apakah logout punya JSON requestBody, dan apakah read operations mempunyai 400? Dijalankan `python3 - <<'PY'` dengan PyYAML, tanpa writes:

   ```python
   import yaml
   from pathlib import Path
   source = yaml.safe_load(Path('api/openapi/index.yaml').read_text())
   bundle = yaml.safe_load(Path('api/openapi.yaml').read_text())
   print('source_bundle_semantically_equal:', source == bundle)
   for path, method in [('/auth/logout', 'post'),
                        ('/organization-establishments/{id}', 'get'),
                        ('/me/organizations', 'get'),
                        ('/organizations/{id}', 'get')]:
       operation = source['paths'][path][method]
       print(method.upper(), path, 'requestBody=', operation.get('requestBody'),
             'responses=', list(operation['responses']))
   print('cookie_scheme:', source['components']['securitySchemes']['sessionCookie'])
   print('callback_query:', source['paths']['/auth/google/callback']['get']['parameters'])
   ```

   Exit 0; equality `True`; logout requestBody `None`, statuses `204/401/403`; reads masing-masing `200/401/404`, `200/401`, `200/401/404`; cookie `kencleng_session`; callback required code/state. Hasil mendukung F1–F4 dan menolak dugaan independent bundle drift. Generated types diperiksa penuh secara statis, tidak menjalankan generator ulang.
3. **Targeted primary protocol lookup:** Pertanyaan F3: apakah 400 + Location memenuhi redirect semantics dan apakah authorization denial mengandung code? Membaca RFC 9110 §§15.4.3/15.5.1 dan RFC 6749 §4.1.2.1 dari RFC Editor pada 2026-10-09; mendukung F3. Tidak melakukan provider/network callback reproduction.
4. **Terminal artifact/scope checks:** hash/HEAD/workflow/status diperiksa kembali; satu structured Phase handoff pada findings dan hanya dua review evidence files baru ditulis oleh Run ini.

Tidak menjalankan validate/bundle/types, dependency install, unit/integration/browser/full build suites, atau runtime backend/provider/DB/security tests. Build command results tidak direlabel sebagai verification Review. Parsed equality bukan bukti runtime, byte-identical regeneration, atau consumer compilation.

## Verdict

**Request changes.** Empat blocking findings: **F1, F2, F3, F4**. Tidak ada additional non-blocking comment yang memicu patch loop. Patch yang diperlukan dibatasi pada T1 contract source dan derived artifacts; tidak ada perubahan Product/Techplan/security scope baru yang diminta.

## Phase handoff

- **Outcome:** COMPLETED — empat pass independent T1 Code Review selesai; verdict Request changes.
- **Result refs:** `runs/RUN-C1-ENG-CODEREVIEW-T1-001/evidence/review-findings-1.md` (terminal carrier; target/input identities pada tabel provenance); `runs/RUN-C1-ENG-CODEREVIEW-T1-001/evidence/patch-plan-1.md` (F1–F4). Target baseline `112c35f8322b53bac39d511ed4f22982cea619af`; reviewed source/bundle/types SHA-256 masing-masing `8ff663e975c9d72e0d6b636994533f420fe77639398197e6137fae49e59fb62d`, `9e20887011ad30a3c774edd0fb725659bccd28446fe5e4aa70c0216fbcb7a4e9`, `a0ae33bccdf194190f3dbb49d69435d983acb8cc28fd712b3a2869442dbe6e6e`.
- **Findings:** F1 production-cookie mismatch; F2 incomplete ordinary-mutation browser/JSON contract; F3 invalid callback failure redirect/input branch; F4 missing invalid read-input responses. Seluruhnya blocking untuk approval T1 yang direview.
- **Decision requests:** None — koreksi mengikuti approved R9/R10/R13/§8; tidak membuka kembali Techplan/split/G1–G3.
- **Blockers:** Tidak ada execution blocker bagi Review. F1–F4 perlu ditangani sebelum target T1 ini dapat menerima approval; tidak menyatakan seluruh WU blocked.
- **Open / unverified:** Real provider/session/Owner/object scope/guard/DB/commit-replay/browser correctness, real operational prerequisites, generated-output reproduction setelah patch, dan Human semantic acceptance belum established oleh Review ini.
- **Recommended continuation:** Orchestrator rekonsiliasi findings dan route fresh bounded T1 Build/Patch Run/Participant dengan patch plan, lalu independent review karena patch menyentuh auth/browser/HTTP contract semantics. Reviewer tidak dispatch T2/T5/Testing; downstream dependency satisfaction tetap Orchestrator-owned.
- **Context refs:** Approved `techplan/techplan.md` R9/R10/R13 dan §8; `techplan/tasks/T1-shared-contract.md`; invocation Run ini; F1–F4 dan patch plan. Relative refs pada handoff berasal dari WU directory.
