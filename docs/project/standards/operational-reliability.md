# Kencleng — Minimum Runtime Reliability (T08)

> **Status:** Project-wide Engineering standard T08 — **settled/persisted only on research branch** `pilot/station5-c1-space-only`. Approval ini hanya untuk **minimum MVP baseline**; **bukan** optimal architecture untuk production, production-readiness approval, bukti uptime/security/recovery, atau promosi ke `main`.
> **Scope:** Minimum configuration safety, runtime lifecycle, basic logging/health, dan non-destructive data/operational safeguards pada tahap MVP Kencleng saat ini.
> **Authority:** [Engineering Standards index](README.md); [Backend architecture](backend.md), [Identity & Authentication](identity-authentication.md), [Authorization](authorization.md), [Privacy & Data Protection](privacy-data-protection.md), [Database & Data Model](database-data-model.md), [API & Contract](api-contract.md), dan [Testing & Verification](testing-verification.md) tetap owning standards untuk concern masing-masing. [Local setup](../kencleng-repo-setup.md) tetap owner developer-environment commands/topology.

## Baseline scope dan production boundary

- **Minimum sufficient untuk MVP saat ini, bukan production-optimal design.** Gunakan kemampuan runtime/tooling yang sudah ada selama cukup dan benar. Tidak wajib menambah third-party logging management, hosted observability/health monitoring, dashboards, tracing collectors, metrics/alerting stack, on-call infrastructure, deployment platform, ataupun operator automation.
- Persetujuan T08 **tidak** mengotorisasi deployment production, memproses uang nyata, memakai data sensitif nyata tanpa safeguards yang sesuai, menambah service/infrastructure, atau mengubah Product delivery scope. Readiness ditentukan oleh evidence dan konsekuensi real usage, bukan label `MVP`.
- **Reassess sebelum real production/external consequential usage**: sebelum Kencleng mengandalkan real financial/identity/sensitive records, real money movement, external users, atau public deployment, identifikasi risks dan safeguards tambahan yang benar-benar diperlukan—termasuk recoverability, access/secrets, appropriate incident detection, deployment compatibility, dan operational responsibility. Tentukan controls dari risk dan actual topology, bukan otomatis menjadikan seluruh observability stack wajib.
- Jika kebutuhan nyata berkembang, keputusan operational tambahan dapat ditetapkan secara scoped dan diverifikasi; tidak perlu sekarang mendefinisikan RPO/RTO numerik, retention schedule, vendor, SLO, atau mandatory 24/7 on-call tanpa evidence/operational requirement.

## T08-A — Configuration safety (minimum)

- Required environment/configuration harus divalidasi pada startup/initialization boundary yang relevan. Invalid required configuration menyebabkan failure yang jelas/fail-closed; jangan silently fallback ke insecure production settings atau misrouted resources.
- Credentials, tokens, signing/encryption keys, dan PII tidak boleh di-hardcode, di-commit, atau dicatat di logs. Gunakan existing local env/Compose mechanism sesuai [repo setup](../kencleng-repo-setup.md) dan secret-handling constraints T02/T04; jangan buat mandatory config/secrets service baru.
- Local-only exceptions (terutama cookie security) **tidak boleh** ditransfer ke public/non-loopback environment. Penerimaan config local secara teknis bukan production-readiness proof.
- Jangan mengubah existing environment files, credentials, port binding, container exposure, atau storage volumes sebagai incidental setup cleanup.

## T08-B — Runtime lifecycle and bounded failures

- Pertahankan/terapkan **graceful shutdown**, resource cleanup, HTTP/request/dependency timeouts dan cancellation **ketika relevant pada active operation**. Jangan memasang timeout/retry boilerplate seragam tanpa memahami side effects dan downstream dependencies.
- Retry hanya untuk operation yang aman diulang atau punya correct idempotency safeguards. **Ambiguous DB commit / financial outcome bukan otomatis failure atau success**; lakukan appropriate state verification/reconciliation sesuai T05/T06 bila business operation membutuhkannya.
- Existing Go server punya startup database connectivity check, HTTP header timeout dan bounded shutdown; itu adalah **implemented baseline behavior**, bukan bukti process/DB selalu sehat atau semua runtime paths sudah reliable.
- Worker, background job scheduling, circuit breakers, queue, elaborate retry engines, autoscaling dan deployment orchestration tidak diwajibkan tanpa active requirements.

## T08-C — Basic logging and existing health

- **Cukup logging bawaan Go** yang sudah dipakai; tidak wajib migrasi ke structured logger/library baru untuk MVP saat ini. Catat errors pada meaningful boundaries dengan context minimum yang membantu diagnosis, bukan log setiap function/layer atau duplikasi error chain. Jangan log secrets, raw tokens, unnecessary PII, financial credentials atau full confidential payloads (T04).
- Pertahankan existing `GET /healthz` sebagai **basic process HTTP response signal**. Response `{"status":"ok"}` **tidak membuktikan** current DB dependency readiness, full business correctness, payment integration, maupun production availability. Backend connectivity check saat startup juga tidak menjamin ongoing availability.
- Tidak wajib membuat `/readyz`, Prometheus metrics, distributed tracing, log aggregation, dashboards, external uptime probes, alerting service, atau 3rd-party monitoring sekarang. Tambahkan salah satu hanya bila operational/deployment question nyata memerlukannya, dengan scoped decision.
- Development diagnostics boleh menggunakan existing process output, `curl` health check, dan Compose/service logs. Jangan menyatakan log entry sebagai proof committed business outcome; durable business evidence mengikuti Product + T01/T05.

## T08-D — Non-destructive data and operational safety

- Setup/verification tidak boleh diam-diam menghapus/mengganti PostgreSQL data, volumes, object storage, secrets, atau deployed schema. Destructive migrations, `down -all`, DB reset, dan irreversible changes membutuhkan **explicit Human authorization, scoped target, impact review**, dan appropriate backup/recovery plan sesuai T05 dan root AGENTS.
- Local MVP sandbox boleh berjalan dengan existing development setup; itu **tidak** berarti production recoverability sudah dipenuhi. Jangan mengklaim backup-safe, restorable atau production ready hanya karena volume persist atau backup file berhasil dibuat.
- **Sebelum real consequential production data/money dipakai**, tetapkan recovery requirements sesuai data criticality, credentials/key custody, realistic backup/retention dan restore procedure; **uji restore** untuk membuktikan recoverability pada representative environment. Backup procedure saja tidak cukup jika restore belum dibuktikan.
- Deployment/production ingress, TLS termination, secrets distribution, safe migration/rollout compatibility, dan incident response di-evaluate ketika public production deployment benar-benar diputuskan. Tidak ada mandatory production CD pipeline/rollback platform sekarang.
- Kegagalan operasional yang bisa menyebabkan uang/state ganda, kehilangan data, atau unauthorized exposure mengikuti domain invariants serta verification requirements T02–T07. Label `MVP` tidak mengurangi safety gate untuk critical consequences.

## Agent usage — lightweight scope check

Saat scoped runtime/deployment work benar-benar memerlukan operational consideration, cukup jawab singkat:
1. Apa changed operation/dependency dan meaningful failure mode?
2. Existing config, logs, `/healthz`, shutdown, atau local setup mana yang cukup? Adakah risiko unsafe retry, state-loss, secrets, atau exposure?
3. Verification apa yang benar-benar dilakukan; apa yang belum terbukti, dan apakah change tetap aman untuk current scope?
4. Apakah change memasuki production/real consequential usage sehingga **reassessment gate** di atas aktif?

Untuk changes yang unrelated, jangan membuat operational checklist panjang. **T08 bukan izin membuat telemetry stack, production architecture, runtime code changes, atau runbook baru otomatis.** New implementation dan readiness claims memerlukan scoped authority dan evidence.

## Explicit current evidence limits

- Repo baseline has a Go `/healthz` handler and basic startup/shutdown behavior; tidak ada bukti production health monitoring/incident management/backup restore readiness dalam approval T08.
- Document-only standards approval dan source inspection tidak menggantikan executed runtime verification.
- [T07](testing-verification.md) menentukan proportional verification dan truthful evidence reporting.

**Promotion boundary:** T08 hanya settled dalam research branch `pilot/station5-c1-space-only`. Untuk menjadi operational `main` perlu authorized promotion tersendiri.
