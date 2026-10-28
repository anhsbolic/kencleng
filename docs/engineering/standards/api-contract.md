# Kencleng — API & Contract Standards (T06)

> **Status:** Project-wide Engineering policy T06, disetujui Anhar dan dipersist pada research branch `pilot/station5-c1-space-only`. **Tidak ada feature endpoint baru, deployed API contract, API client, atau conformance test yang diklaim telah diimplementasikan.**
> **Scope:** REST interface, OpenAPI authority, JSON response/error conventions, API security, pagination, idempotency, compatibility, input/output safety, dan feature-level contract verification.
> **Authority:** [Engineering Standards index](README.md). Approved Product/scoped requirements menentukan behavior dan domain invariants; T06 menentukan reusable technical API rules. [T01 Backend](backend.md), [T02 Authentication](identity-authentication.md), [T03 Authorization](authorization.md), [T04 Privacy](privacy-data-protection.md), dan [T05 Database](database-data-model.md) tetap memiliki concern masing-masing.

## T06-A — Contract authority dan contract-first execution

- **Contract-first saat API operation benar-benar dibutuhkan**, bukan mendefinisikan semua endpoint sebelum ada authorized Product/delivery requirement. Kontrak API adalah turunan dari requirement approved, tidak boleh membuat Product semantics/permissions/state baru.
- **Editable OpenAPI source** saat ini di `api/openapi/index.yaml` (OpenAPI **3.0.3**), dan boleh dipecah menjadi beberapa source files saat complexity aktual membutuhkan. `api/openapi.yaml` adalah **generated bundled view**, bukan second editable authority. Sumber kontrak machine-readable yang sudah accepted menentukan interface yang harus dipatuhi FE/BE.
- `api/` saat ini tooling-only dengan `paths: {}`; jangan menambahkan endpoint placeholder, menyalin historical API, atau menyatakan Product API sudah aktif hanya karena T06 berlaku.
- Schema/operation details harus mencakup method, path, parameters, authentication/security requirements, request/response schemas, status/failure semantics, exposed fields, dan state transition requirements yang approved. OpenAPI tidak mengalahkan upstream Product authority maupun existing T02–T05 security policy.
- Generate/reuse contract-derived types pada frontend ketika active API tersedia; jangan pelihara parallel handwritten DTO schemas sebagai rival authority. Tooling lain hanya ditambahkan jika actual development membutuhkannya.

## T06-B — Resource-oriented REST dan HTTP semantics

- Default resource-oriented HTTP paths memakai plural nouns dan consistent casing untuk segments/parameter names. Contoh `/organizations/{organizationId}` hanya illustrative, **bukan authorized endpoint**. API boleh menyediakan domain action endpoint eksplisit ketika CRUD tidak secara jujur mewakili operation.
- `GET`/HEAD tidak mengubah consequential business state. `POST` untuk create atau explicit domain action; `PUT` untuk replacement bila semantics-nya cocok; `PATCH` untuk partial update yang bentuknya terdefinisi; `DELETE` hanya bila removal/revocation memang diizinkan oleh Product.
- Gunakan HTTP status code sesuai outcome sebenarnya: `200` read/action success dengan representation; `201` created; `202` accepted but **not completed**; `204` success tanpa body. `401` authentication missing/invalid, `403` disallowed, `404` unavailable/not visible, `409` actual conflict, `422` semantically unprocessable content saat appropriate. Status dipilih berdasarkan operation/risk; jangan menganggap list tersebut complete/default universal.
- Tidak boleh membungkus business failure dalam `200 OK` dengan `success:false`. HTTP status dan runtime outcome harus konsisten. `Location`, caching headers, retry hints dan other HTTP semantics digunakan bila relevan.

## T06-C — JSON payload conventions dan minimal success envelope

- **JSON field naming: `camelCase`** untuk publicly exposed application contract, independen dari PostgreSQL `snake_case`. Gunakan stable resource IDs, explicit optional/nullability semantics, dan RFC 3339 date-time strings untuk instants (UTC di system boundary). Money harus memakai exact representation sesuai approved Product contract (mis. string decimal atau integer minor unit yang aman untuk consumer), **bukan binary floating-point** atau silent JS rounding.
- **Success JSON responses yang memang memiliki response body** memakai top-level envelope **`{ "data": <contract-defined value>, "meta"?: <metadata> }`**. `data` merupakan resource/collection/action representation sesuai masing-masing operation; `meta` hanya ada ketika metadata nyata dibutuhkan (mis. cursor/pagination). Jangan mengisi `meta:null` secara default.
- **Tidak wajib top-level `success`, `code`, `message`, atau `errors` dalam success wire payload**. Status HTTP sudah menggambarkan berhasil/gagal; UX copy ditentukan UI atau approved Product semantics. Jangan menerbitkan redundant `success:true`/generic message hanya demi bentuk enam-field.
- **No-body/non-JSON exceptions**: `204 No Content`, `HEAD`, `304 Not Modified`, redirects, file download/streaming, dan protocol-specific representations tidak dipaksa JSON envelope. Existing technical health endpoint `/healthz` boleh mempertahankan response shape baseline; kebijakan ini tidak mengotorisasi perubahan endpoint tersebut. Exceptions lain hanya saat benar-benar diperlukan oleh operation dan dicatat jelas pada contract.
- **Contoh illustrative**, bukan approved API:

```json
{"data":{"id":"example-id"}}
```

```json
{"data":[{"id":"item-1"},{"id":"item-2"}],"meta":{"nextCursor":"example-cursor"}}
```

## T06-D — Error format: RFC 9457 Problem Details

- Untuk error HTTP API yang menghasilkan JSON body, gunakan **RFC 9457 `application/problem+json`**, bukan mandatory custom `{success,code,data,errors,message,meta}` envelope. `type` adalah **stable machine-readable problem identifier** milik namespace yang dikendalikan Kencleng ketika diimplementasikan; `title` ringkas, `detail` aman untuk client, `instance` bila relevan.
- HTTP error status tetap authoritative pada wire; jika member `status` dicantumkan, **harus sama dengan actual HTTP status**. Jangan mengeluarkan `200` untuk permission/validation/business failure lalu menyisipkan indikasi gagal dalam JSON body.
- Field-level validation boleh memiliki optional **`errors` extension** yang typed dan terdokumentasi (misalnya field/pointer + safe detail), sesuai RFC 9457. `errors` tidak wajib untuk semua error; gunakan stable `type` untuk machine handling, bukan parsing teks `detail`/`message`.
- Detail field tidak boleh membeberkan SQL, secrets, stack traces, sensitive identity existence, hidden resource data, atau internal policy. Pertimbangkan 404 untuk concealment hanya bila security semantics justified; 401/403/404 distinctions mengikuti T03.
- Contoh ilustratif; URL problem type **placeholder**, bukan published Kencleng identifier:

```json
{
  "type": "https://example.com/problems/validation-error",
  "title": "Validation failed",
  "status": 422,
  "errors": [{"field": "name", "detail": "This field is required"}]
}
```

## T06-E — API authentication, authorization, CSRF dan service clients

- Browser authentication mengikuti **Kencleng-issued short-lived JWT access cookie `HttpOnly` + rotating refresh token**, bukan frontend-readable bearer token. Definisikan cookie security scheme yang merepresentasikan actual cookie transport di OpenAPI 3.0.3 (mis. `type: apiKey`, `in: cookie`; ini **syntax scheme OpenAPI, bukan kewajiban menambahkan separate API key**). Cookie names/lifetimes mengikuti scoped security implementation; T02 tetap owner.
- State-changing cookie-authenticated endpoints (termasuk refresh/logout) tetap memerlukan appropriate backend **CSRF protection** beserta Origin safeguards; nyatakan request headers/requirements di OpenAPI ketika aktif. Security-scheme declaration sendiri tidak otomatis membuktikan CSRF defense.
- Setiap protected operation menjalankan T03 effective user/service authorization untuk resource/action/context; guest access hanya bila approved Product semantics memperbolehkan. CORS, dev rewrite, frontend gating, dan browser-supplied headers bukan trust proof.
- External/machine clients tetap mengikuti separate service identity/scope pada T03, tidak otomatis memakai browser cookie atau mandatory multi-key mechanism.
- FE `/api/*` rewrite saat `next dev` **bukan** deployed public gateway/production API versioning; backend currently unprefixed. Jangan membangun API gateway atau external OAuth service sekarang hanya untuk T06.

## T06-F — Bounded collections, pagination, filter/sort

- List responses harus **bounded**, tidak boleh unbounded bulk export tanpa scoped requirement. Allowed filters, ordering, cursor/offset semantics dan limits dicatat per approved collection operation.
- Gunakan stable deterministic sort/tie-breaker terutama untuk pagination. Cursor sesuai high-volume/changing datasets ketika justified; offset boleh untuk bounded/simple lists. **Tidak wajib cursor untuk seluruh MVP** atau universal magic page-size constant tanpa workload.
- Pagination metadata ditempatkan pada optional `meta` dalam minimal success envelope. Filter/search harus menerapkan scope/authorization/privacy sebelum data keluar; jangan mengembalikan semua raw rows lalu menyaring di browser.
- Pertimbangkan `nextCursor` dan collection consistency semantics bila mutation paralel bisa menimbulkan duplicate/missing rows. Exact behavior dicatat dalam contract.

## T06-G — Idempotency, concurrent retries dan async truth

- Untuk **consequential non-idempotent writes** yang berisiko menciptakan effect ganda (mis. money-related operations setelah authorized), design harus menetapkan retry identity dan idempotency secara eksplisit. Jangan mewajibkan `Idempotency-Key` di semua POST tanpa kebutuhan nyata.
- Jika menggunakan `Idempotency-Key`, definisikan caller/scope binding, key uniqueness, request-fingerprint mismatch behavior, concurrent duplicate handling, replayed result semantics, retention dan backend transaction linkage agar effect tidak duplicated. Application-level in-memory cache atau HTTP client retry sendiri tidak menjamin exactly-once effects.
- Backend tidak boleh mengklaim sukses dari ambiguous DB commit/settlement state. Response `202 Accepted` hanya menyatakan pekerjaan diterima, **bukan sudah selesai**; asynchronous status/result observation diputuskan menurut actual approved Product workflow.
- Worker, callback, webhook, dan retries tetap tunduk pada T01/T03/T05 transactional/domain invariants; jangan membangun generic job API atau workflow engine untuk MVP tanpa kebutuhan.

## T06-H — Versioning dan contract compatibility

- **Internal-first MVP:** tidak mewajibkan `/v1` prefix untuk seluruh internal endpoints. OpenAPI `info.version` mengidentifikasi contract document version, **bukan otomatis URL/public API version**.
- Bila active FE dan BE tidak deploy bersamaan, ubah contract secara compatibility-aware; identifikasi breaking changes pada path/method, validation, status/errors, nullable/required fields, enum semantics, pagination, dan security scheme. Coordinate consumer rollout, tests, dan migration jika diperlukan.
- Public/external API exposure, versioning policy, deprecation window, dan endpoint lifecycle memerlukan scoped decision tersendiri sebelum clients eksternal dijanjikan stabilitas. Jangan menciptakan public `v1` hanya untuk future-proofing.

## T06-I — Input/output and trust boundary

- Semua request harus punya defined content types, bounded payload size, parsing/validation policy, required vs optional fields dan safe error output. Definisi kontrak harus sesuai security/property ownership pada T02–T04.
- **No mass assignment** dari JSON ke privileged internal model. Client-provided `accountId`, owner, role, resource scope, approval/settlement status atau monetary state tidak otomatis menjadi trusted authority; assign/verify server-side sesuai approved workflow.
- **Minimize responses before browser delivery**: bila UI hanya membutuhkan masked PII, API DTO mengirim masked value atau omit raw field (T04), bukan full PII yang dimasking dengan React. Tetapkan cache controls yang sesuai untuk private data, hindari secrets pada logs, URLs, errors dan responses.
- Reject/contain unexpected input secara proporsional; toleransi forward compatibility yang disengaja harus eksplisit, bukan membolehkan unknown fields menaikkan privileges. Hindari oversized body/DoS dan ambiguous JSON duplicate keys sesuai threat review bila relevan.

## T06-J — Contract verification dan Agent scaffold

- Ketika API feature diotorisasi, agent wajib mengisi inventaris bounded berikut pada **feature-level solution/contract**, bukan mengisi OpenAPI `paths` dengan speculative endpoints:

| Item | Required per approved API operation |
| --- | --- |
| Product/scoped authority | Requirement IDs, intended effect dan invariant |
| Operation | Method, path, resource/action, visibility |
| Security | Cookie/session/CSRF, principal, scope, guest/service rules |
| Inputs | Path/query/header/body schema, sizes dan validation |
| Success | HTTP code, `{data,meta?}` jika applicable, nullable/optional fields, no-body exceptions |
| Errors | RFC 9457 problem types, statuses, extensions, privacy-aware details |
| Collections | Limits, filters, sort, pagination/cursor consistency |
| Stateful effects | Atomicity, concurrency, idempotency/replay, async truth and status |
| Outputs | T04 masking, authorized reveal, cache/PII protections |
| Compatibility | Consumer impact, breaking change analysis, external versioning if applicable |
| Evidence | Spec lint/bundle, contract conformance, FE integration, security/negative/retry tests |

- Setelah active contract berubah, jalankan `npm run validate` dan `npm run bundle` dari `api/`; generated bundle harus sama dengan source yang diterima. FE types boleh di-generate, tetapi **generated types dan lint bukan bukti backend runtime memenuhi contract**. Gunakan executed backend/FE/negative tests sesuai risk.
- Untuk FE developer experience, gunakan **small typed API client/result abstraction** ketika active API dibutuhkan: normalisasi `{data,meta?}`, RFC 9457, transport failures, dan typed success/error result agar components tidak mengulang parsing HTTP. Contoh `success` discriminator boleh ada pada **client-side TypeScript result**, **tidak menjadi mandatory field pada wire JSON**. Jangan membangun general SDK atau adapter framework sebelum need.
- OpenAPI source changes tidak boleh memperkenalkan protected auth/security implementations, Product feature semantics, atau migrations tanpa respective authority; project standard is not implementation permission.

## References and boundaries

- [OpenAPI 3.0.3 Specification](https://spec.openapis.org/oas/v3.0.3)
- [RFC 9457 — Problem Details for HTTP APIs](https://www.rfc-editor.org/rfc/rfc9457.html)
- [RFC 9110 — HTTP Semantics](https://www.rfc-editor.org/rfc/rfc9110.html)
- [Current API tooling](../../../api/README.md)

**T06 settles the reusable contract convention.** Belum mengotorisasi actual OpenAPI operations, modifications to `api/openapi/index.yaml`, server/API client code, authentication surface atau public API release. Itu tetap memerlukan active scoped delivery decisions dan verification.
