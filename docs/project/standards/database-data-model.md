# Kencleng — Database & Data Model Standards (T05)

> **Status:** Project-wide Engineering Standard T05 — Human-settled, persisted on research branch `pilot/station5-c1-space-only`. **This is not an implemented schema, migration, or runtime verification.**
> **Scope:** PostgreSQL relational modeling, schema authority, types, integrity, data ownership, transactions, migrations, indexing, logical ERD, feature-level design scaffold, dan high-volume partitioning.
> **Authority:** [Engineering Standards index](README.md). Product/scoped approved requirements own business meaning; [T01 Backend Architecture](backend.md) owns module boundaries; [T02 Identity](identity-authentication.md), [T03 Authorization](authorization.md), dan [T04 Privacy](privacy-data-protection.md) own their respective concerns.

## Baseline and decisions

- **SQL-first, PostgreSQL 16** as selected database baseline; `pgx` for Go access, `goqu` when parameterized query building justifies it, `golang-migrate` for versioned explicit SQL migrations. Tidak mewajibkan ORM atau generic repository abstraction.
- **One PostgreSQL logical schema `public` for MVP by default**, shared by the Modular Monolith. Ownership ditentukan oleh module/domain operations, **bukan** kewajiban membuat separate PostgreSQL schema untuk tiap domain. Isolation tambahan memerlukan justified solution decision.
- **Regular (non-partitioned) tables by default**, terutama transactional core yang membutuhkan stable identity, global uniqueness dan referential integrity. Partitioning hanya dengan justification pada T05-J.

## T05-A — Schema authority and actual state

- Approved Product/feature contract defines required durable state and invariants; **versioned SQL migration files are the authoritative mechanism for intended physical schema evolution**; **introspected live PostgreSQL** is evidence of actual applied state, not automatically valid desired policy.
- `backend/migrations/` menjadi owning directory ketika schema delivery pertama diotorisasi; baseline saat ini hanya `.gitkeep` dan **belum ada active schema migrations**.
- Live schema, migration history/version, dan repo migration source harus direkonsiliasi sebelum mengklaim deployment siap. Jangan menganggap ERD, historical SQL, code models, atau migration files semata sebagai bukti database telah sesuai.
- API dan Worker memakai shared PostgreSQL schema/migration authority; app startup tidak otomatis menjalankan migrations atau melakukan repair schema.

## T05-B — Naming, stable IDs and data types

- Default nama tables **plural `snake_case`**, columns/indexes/constraints **`snake_case`**, PK column `id`, FK reference `<entity>_id` bila tidak ambigu; pengecualian harus jelas dan konsisten. Constraint/index names sebaiknya human-readable dan stabil, bukan menggantikan semantic owner.
- **UUID is the default for new opaque business entity IDs**, dengan explicit exception bila domain/query/operability memerlukannya. ID internal tidak berasal dari user-supplied business name/email atau external provider subject; T02 tetap menetapkan stable internal `account_id`.
- Gunakan PostgreSQL native types sesuai semantics: `TIMESTAMPTZ` untuk time instants dengan interpretasi UTC di application boundary; `DATE` untuk calendar date tanpa time-of-day; `BOOLEAN` untuk binary state; numeric/integer types yang tepat untuk amounts/counts. Simpan unit/currency secara eksplisit bila bermakna. `created_at`/`updated_at` dibuat hanya jika dipakai oleh lifecycle/audit, bukan template wajib seluruh tabel.
- **Money uses exact arithmetic**: `BIGINT` minor units atau constrained `NUMERIC(precision,scale)` berdasarkan approved currency/rounding contract; jangan gunakan binary floating point. Jangan menciptakan settlement precision/rounding policy dari database convention.
- Gunakan `TEXT` plus appropriate `CHECK` atau constrained representation lainnya ketika status values memang settled; **PostgreSQL ENUM bukan universal default**. `JSONB` hanya untuk justified flexible payload, bukan pengganti explicit relational model untuk state yang memiliki invariants/query penting.
- `NULL` vs `NOT NULL` mencerminkan true optionality; hindari invented defaults, semantically overloaded sentinel values, dan status columns tanpa Product meanings.

## T05-C — Relational constraints and invariant enforcement

- Gunakan DB-level `PRIMARY KEY`, `FOREIGN KEY`, `UNIQUE`, `NOT NULL`, dan `CHECK` untuk relational/local invariants yang sungguh berlaku; implement application/domain validation tetap diperlukan untuk context/business workflow.
- Constraints harus berasal dari approved requirements. Contoh: T02 verified `(issuer,subject)` mapping perlu uniqueness yang sesuai, tetapi **kemiripan/nama organisasi pada Product ORG-01 tidak mengizinkan `UNIQUE(organization_name)`** tanpa Product decision baru.
- Tentukan `ON DELETE`/`ON UPDATE` effects dengan sengaja; jangan memilih cascade untuk relasi sensitif hanya demi nyaman. Soft delete, historical snapshots dan auditability bukan default universal.
- Multi-row, cross-table, cross-domain, time-dependent, atau authorization invariants mungkin memerlukan bounded transaction coordination. PostgreSQL `CHECK` tidak menggantikan domain orchestration. Jangan mengekspresikan correctness finansial sebagai trigger/retry gimmick tanpa model transaksi yang jelas.

## T05-D — Logical domain data ownership

- Setiap relevant domain module bertanggung jawab terhadap data dan invariants-nya sesuai T01. Domain lain tidak boleh mutation langsung melewati owner operations; shared DB bukan shared business ownership.
- Read-only cross-domain composition boleh jika bounded dan authorization/PII constraints dihormati. Physical FK lintas-domain boleh jika mendukung integrity; tidak otomatis menjadikan logical ownership kabur.
- Satu database dan schema `public` untuk MVP tidak berarti query dari seluruh module bebas memodifikasi tabel domain lain. Database schema-per-domain, multi-DB, dan event-driven replication memerlukan actual requirement.

## T05-E — Transactions, concurrency and idempotency

- **Transaction boundary mengikuti invariant, bukan package boundaries.** Cross-domain transaction dalam satu PostgreSQL instance diperbolehkan saat atomicity diperlukan; siapa yang mengkoordinasikan transaksi harus eksplisit. Efek yang harus terjadi bersama tidak boleh dipersist sebagai hasil parsial.
- Tentukan transaction outcome/commit certainty sebelum mengklaim sukses. Jika result unknown, error, atau timeout, jangan buat klaim sukses palsu dan jangan ulang efek finansial secara membabi buta.
- Pilih isolation level, row/advisory locks, unique constraints, conflict handling, dan retries berdasarkan measured/proven race cases; **`SERIALIZABLE` bukan baseline semua transaksi**.
- Idempotency, duplicate delivery, retry dan multi-actor concurrency untuk consequential writes harus dirancang berdasarkan approved business effect. Exactly-once effect bukan jaminan otomatis hanya dari queue, request ID, atau DB transaction; butuh correct uniqueness/atomic decision boundary.
- Security-sensitive effective state (membership revocation, key lifecycle, privacy access) harus memakai appropriate trusted state dan locking/serialization bila diperlukan sesuai T02–T04.

## T05-F — Controlled migration lifecycle and recovery

- Gunakan versioned `golang-migrate` SQL files di `backend/migrations/`, dengan paired `.up.sql` dan `.down.sql` naming convention **UTC timestamp + descriptive snake_case**. Naming/ordering harus unique, monotonically ordered, dan stable setelah diterapkan di shared environment.
- **Migrations adalah explicit controlled deployment/admin action**, bukan implicit side effect saat API atau Worker startup. Baca current schema/version dan assess safety sebelum menjalankan.
- Jangan edit/reorder previously applied migrations pada shared environments. Perubahan memakai migration baru; exception untuk disposable local databases hanya setelah memastikan tidak menjadi shared baseline.
- Untuk migration yang mengubah data/constraint, catat compatibility window, data backfill, lock/downtime risk, forward strategy dan realistic rollback/recovery. Untuk large live table dapat gunakan expand → backfill → contract jika lebih aman; tidak perlu memaksa ceremony ini pada empty MVP schema.
- `.down.sql` **tidak menjamin data dapat dipulihkan**. Jika reversal destructive/unsafe, tulis batasan dan recovery plan secara jujur; jangan mengemas drop data sebagai safe rollback.
- **Critical existing hazard:** `backend/Makefile` memiliki `migrate-down` yang memanggil `down -all`. Target tersebut **tidak boleh dijalankan sebagai generic rollback/reset/verification command** pada database berdata. Explicit Human approval dan scope + backup/recovery inspection diperlukan sebelum destructive actions. Standar ini **tidak mengubah target Makefile** atau mengotorisasi migrasi apa pun.
- Jika migration gagal atau history berstatus dirty, investigasi applied SQL dan actual DB state sebelum perbaikan; `force` bukan repair script yang boleh dipakai sembarangan. Jaga perbedaan antara verified schema, assumed schema dan failed attempt.
- Database backup, volume deletion dan major collation changes tetap mengikuti existing [repository setup safeguards](../kencleng-repo-setup.md).

## T05-G — Indexes, query plans and operability

- Index berdasar actual correctness constraints dan query patterns; jangan tambahkan index universal untuk seluruh FK/status/timestamp tanpa justified needs. PK/unique constraints membuat index relevan, sedangkan FK referencing columns tidak otomatis selalu indexed oleh PostgreSQL.
- Pertimbangkan selectivity, write overhead, cardinality, pagination/filter/order, and tenant/resource scoping. Security constraints T03/T04 tetap berlaku pada database read paths.
- Untuk performance-driven index/persistence change, verifikasi `EXPLAIN (ANALYZE, BUFFERS)` dengan representative dataset dan inspect observed plan. Jangan menyebut optimisasi terbukti dari ukuran data imajiner.
- Untuk index pada database aktif, pertimbangkan locks, deployment availability, dan PostgreSQL `CREATE INDEX CONCURRENTLY` semantics: **tidak dapat berjalan dalam transaction block**, dapat meninggalkan invalid index bila gagal, dan memiliki batasan khusus partitioned parents. Migration runner behavior perlu diperiksa sebelum menggunakannya.

## T05-H — Logical ERD boundaries

- Gunakan **feature/domain-scoped logical ERD** saat schema memiliki meaningful relationships, cardinality, optionality, dan constraints yang perlu dikomunikasikan. Mermaid dalam Markdown boleh jika cukup; editable media lain boleh jika memberi value.
- Jangan buat mandatory giant ERD seluruh Kencleng saat features belum dipilih. Project-level consolidated diagram hanya saat cross-domain understanding benar-benar membutuhkannya.
- Diagram/feature design bukan deployed physical-schema authority. Rekonsiliasi dengan SQL migration dan live introspection ketika implementation diperiksa; historical Rancang snapshots tetap historical provenance, bukan canonical schema yang harus diikuti.

## T05-I — Feature-level database design and migration scaffold

Agent/Solution Architect harus mengisi berikut **hanya untuk proposed/approved feature yang butuh persisted state**, lalu menghubungkan ke actual migration/verification evidence ketika implemented.

| Item | Required feature-level description |
| --- | --- |
| Product authority & invariants | Approved requirement IDs, durable outcomes, semantic constraints |
| Entities and ownership | Owning domain, entity ID, lifecycle, relationship/cardinality |
| Logical ERD | Scoped diagram or clear relationship specification |
| Schema design | Tables, fields/types, PK/FK/UNIQUE/CHECK, default/null semantics |
| Cross-concern safeguards | T02 identity, T03 scope/authorization, T04 privacy/encryption/retention |
| Transaction design | Atomic effects, concurrency, idempotency, locking, failure/unknown-result handling |
| Migration/change plan | SQL version, up/down limits, compatibility, backfill, rollback/recovery, approval gate |
| Access/query plan | Important queries, indexes/constraints justified by expected or observed needs |
| Partitioning decision | Default non-partitioned or explicit T05-J justification and tradeoffs |
| Verification | Fresh apply, existing-data upgrades where relevant, constraints, negative/race/rollback checks, runtime evidence |

Do not invent speculative global tables, starter migrations, arbitrary business constraints, or an ORM generator. Missing Product decisions must be escalated.

## T05-J — High-volume data and partitioning

- **Default all new tables to non-partitioned.** In particular, the **transactional Donation core** keeps a regular table by default for stable identity, global uniqueness, FK relationships, and idempotency correctness; this is a design baseline, **not** an approved Donation schema or feature delivery.
- Partitioning is a **conditional implementation choice** based on real/high-confidence volume, query predicates, pruning benefit, retention/archive operations, or demonstrated operational pressure. High row count alone does not prove benefit.
- **Monthly `RANGE` partitioning** is a candidate for appropriately time-scoped **append-only history, activity, webhook, or audit datasets**, **only if** access patterns and data lifecycle support it. This does not require creating any such table or partition now.
- In PostgreSQL 16, a parent partitioned table's `PRIMARY KEY` / `UNIQUE` constraints must include the partition key. A simple globally unique `donation_id` on a table range-partitioned by `created_at` **cannot be enforced by parent `UNIQUE(donation_id)` alone**. Composite keys affect references, idempotency, cross-month settlement and migrations; do not weaken core correctness to make partitioning fit.
- Querying by ID without partition-key predicates can reduce pruning benefits. A partitioned parent must have valid destination partitions for incoming rows or insertion fails (unless a suitable default partition exists); operational planning must cover future partition creation, boundary dates/timezones, lifecycle, monitoring, index strategy, backup and recovery.
- Any partition eviction/archival is governed by T04 retention and financial/legal obligations; calendar rollover is **not** permission to delete donation records. Evaluate partition creation/backfill/migration complexity versus current non-partitioned operation and rollback.
- Before adoption, write a bounded decision record containing measured or credible evidence, query samples, key/constraint/FK and idempotency impacts, operational costs, migration path, fallback, and representative verification (including cross-period effects and `EXPLAIN (ANALYZE, BUFFERS)` when performance is the justification).

## Supporting documentation and authority limits

- [PostgreSQL 16 — Constraints](https://www.postgresql.org/docs/16/ddl-constraints.html)
- [PostgreSQL 16 — Partitioning](https://www.postgresql.org/docs/16/ddl-partitioning.html)
- [PostgreSQL 16 — CREATE TABLE restrictions](https://www.postgresql.org/docs/16/sql-createtable.html)
- [PostgreSQL 16 — CREATE INDEX](https://www.postgresql.org/docs/16/sql-createindex.html)
- [golang-migrate project](https://github.com/golang-migrate/migrate)

T05 is an Engineering policy, **not** authorization to create production tables, modify Makefile/migration runners, change data, enable partition management, or introduce a new database architecture. Implementation requires active scoped Product/Engineering authority and actual tests.
