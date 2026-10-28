# Kencleng — Authorization Standards (T03)

> **Status:** Project-wide Engineering standard T03 — Human-settled, branch `pilot/station5-c1-space-only`. Belum ada authorization runtime, role tables, atau machine-to-machine API yang dianggap delivered.
> **Scope:** User authorization, system/service authorization, protected operation enforcement dan feature-level permission scaffold. Authentication/session policy dimiliki oleh [Identity & Authentication](identity-authentication.md).
> **Authority:** [Engineering Standards](README.md). Product Authority menentukan **hak siapa terhadap apa**, sementara Engineering menetapkan cara mengevaluasi dan menegakkan hak itu. Tidak menetapkan role matrix atau feature-specific permission di dokumen ini.

## General security rules

- **Deny by default + least privilege**: setiap protected operation harus memiliki dasar allow eksplisit. Public/guest access hanya bila effective Product requirements memang mengizinkannya; no implicit system-wide admin.
- Evaluasi authorization berdasarkan **principal, action, resource, scope/relationship, dan current context**. Role label saja tidak memberikan izin universal, khususnya pada boundary antar-organisasi.
- Enforcement ada di backend pada protected business operation/data boundary yang benar dan digunakan oleh setiap entry point (REST API, Worker, internal callers). Middleware/authentication check boleh mendukung, tetapi bukan satu-satunya authorization pada consequential operations. Frontend gating hanya UX.
- Untuk reads, lists, mutations, dan sensitive fields, wajib menghormati resource/data ownership serta scope yang berlaku. Jangan trust user-supplied role, resource owner, org ID, maupun client-visible claims sebagai sumber grant. Query/list tidak boleh membocorkan data lintas-scope.
- Decisions harus memakai effective server-side state pada titik yang relevan. For consequential mutations pertimbangkan concurrent role revocation, resource-state changes, dan transactional consistency. Log/audit hanya data minimum yang aman; jangan bocorkan alasan internal yang sensitif ke unauthorized clients.
- Privileged authorization core implementation tetap merupakan protected write yang memerlukan explicit Human authorization sesuai root `AGENTS.md`. Approved standard bukan otomatis izin implementasi.

## User authorization — hybrid RBAC + relationship/context

**Kebijakan baseline:** `permission definitions` dan default `role -> permission` policy **didefinisikan dalam version-controlled code**, sedangkan **role/membership/relationship assignments dan resource state** disimpan sebagai durable database state sesuai actual domain.

- Permission identifiers, actions, dan grants ditambahkan berdasarkan approved feature Product semantics. Jangan membuat universal permission list atau menyamakan Product `Owner`, `Staff`, `Reviewer` dengan global unrestricted roles.
- Assignments harus memiliki applicable scope (misalnya organization context) dan dapat berubah tanpa release kode. Revoke/change wajib diperhitungkan pada subsequent effective authorization decisions; jangan jadikan cached JWT permission claims sebagai source of truth.
- Conditional policies seperti review conflict-of-interest, valid resource lifecycle, atau ownership restrictions diterapkan oleh domain authorization logic dengan data yang dipercaya. Permission grant saja tidak melewati contextual denial.
- Dynamic role/permission editor, arbitrary custom roles, generalized policy engine, dan generic polymorphic role schema bukan kewajiban MVP. Pilih physical DB tables dan local authorization components dari kebutuhan nyata, mempertahankan logical relationships dan testability.
- Pengubahan code-defined permission policy melewati code review/verification; pemberian/pencabutan relationship di database mengikuti authorized business operations dan auditability. Jangan mengedit assignments lewat unprotected administrative shortcuts.

## System/service authorization — distinct machine identity

Pisahkan **user principal** (internal `account_id`) dari **service principal** (registered/trusted `service_id`). Authentication aplikasi pengguna tidak otomatis membuktikan pemanggilnya adalah trusted service.

- **Browser/frontend JavaScript bukan confidential client.** API key atau `X-App-Key` yang dikirim dari browser bukan bukti kuat asal request “Frontend Kencleng” dan tidak dapat menjadi secret. Origin/CORS/routing juga bukan pengganti user/resource authorization.
- **Trusted server-to-server clients**, bila diperlukan, harus mempunyai distinct service identity, protected server-side credentials, explicit allowed operations/scopes dan bounded trust. Client credentials dipisahkan per service/client; jangan menggunakan satu shared master key atau memberikan semua system rights secara default.
- **External API clients** baru diaktifkan atas scoped delivery decision: registrasi client identity, issuance/rotation/revocation, audit, rate limits, audience/scope restriction, dan threat-appropriate authentication. Prefer short-lived scoped credentials melalui standard mechanism (misalnya OAuth 2.0 client credentials untuk confidential clients), dengan asymmetric authentication/mTLS jika risikonya membenarkan. Beberapa static header keys tidak secara otomatis menambah jaminan keamanan.
- **Current FE-to-BE development rewrite** adalah routing path, bukan service authentication. Tidak perlu membangun OAuth server, API gateway, atau microservices infrastructure sekarang hanya untuk mengantisipasi future integrations.
- **Worker** menggunakan bounded non-interactive execution identity. Background job tidak boleh mewarisi blanket user privileges. Jika job terkait user action, simpan original actor attribution dan tentukan authorized effective point: mana yang final saat accepted, mana yang harus re-check pada execution; keputusan mengikuti invariants Product dan security, termasuk revocation/retry effects.
- Jika suatu operation memerlukan dua identitas (service client **dan** end-user), keduanya harus divalidasi secara terpisah serta hubungan/delegation context-nya eksplisit. Service authentication tidak menggantikan user authorization.

## Feature authorization scaffold — wajib diisi untuk protected features

Agent/Solution Architect wajib mendokumentasikan **permission inventory dan decision rules** hanya sejauh fitur membutuhkan protected operations, tidak membangun permission management platform generik.

Salin dan isi tabel berikut pada **feature-level solution/contract** (bukan sebagai active permission matrix di dokumen ini):

| Permission ID (stable) | Action | Resource | Applicable scope | Allowed role/relationship | Contextual deny/allow | Enforcement point | Required verification |
| --- | --- | --- | --- | --- | --- | --- | --- |
| *(isi berdasarkan approved feature, tidak ada grant default)* | | | | | | | |

Untuk setiap protected operation, jelaskan: canonical Product authority untuk grants; bagaimana `principal` diverifikasi; bagaimana resource + scope dicari dari trusted state; cara revocation, context change, atau concurrency mempengaruhi decision; dan apakah API/Worker menggunakan protected operation yang sama. Record absence sebagai decision gap, bukan mengarang role atau eligibility.

Saat implementation berlangsung, minimal verifikasi positive allow, denied access, cross-resource/cross-organization access, missing/expired identity, revoked relationship, contextual restriction, dan Worker/service boundary bila digunakan. Evidence berupa executed tests/observations, bukan asumsi. Detail verification gates project-wide ditangani pada T07.

## Batas implementation dan references

Dokumen ini **bukan** commitment untuk generic RBAC tables, runtime permission engine, public API, machine tokens, infrastructure atau deployment. Feature-level architectural choices yang mematuhi boundary ini didelegasikan sesuai [T01](backend.md). Session JWT, refresh rotation, dan cookie safeguards mengikuti [T02](identity-authentication.md).

Supporting security references (guidance, bukan pengganti project authority): [OWASP Authorization Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Authorization_Cheat_Sheet.html), [OAuth 2.0 Security BCP — RFC 9700](https://www.rfc-editor.org/rfc/rfc9700.html).
