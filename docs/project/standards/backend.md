# Kencleng — Backend Tech Stack

> Status: Active clean-baseline architecture
> Purpose: Own settled backend stack and Modular Monolith architecture; cross-cutting identity policy belongs to [Identity & Authentication](identity-authentication.md).

## Baseline choices

| Concern | Choice | Baseline meaning |
|---|---|---|
| Language | Go | Core backend implementation language |
| HTTP | `net/http` with Go pattern routing | Default until a demonstrated need justifies another router |
| Database | PostgreSQL | Relational persistence capability |
| Driver | `pgx` | PostgreSQL access |
| Query construction | `goqu` when query-building needs it | Parameterized query builder, not an ORM |
| Migrations | `golang-migrate` | Explicit SQL migration mechanism |
| Architecture | Modular Monolith | One Go backend codebase/module with independently executable REST API and optional Worker processes; product domains emerge from real delivery needs |
| Testing | Go test + `httptest`, proportional integration/race/security checks | Evidence depth follows the active risk |
| Object storage | Optional SeaweedFS through standard S3 APIs | [Setup and verification](../local-object-storage-options.md); local Podman checks passed according to Human's report, deployment readiness remains separate. Use only when an active requirement needs object storage |

A selected capability being available does not require every feature to use it. Do not create empty domain/platform layers merely to make the architecture look complete.

## Modular Monolith: ownership dan boundaries

- **Satu repository dan satu Go backend module** digunakan bersama oleh REST API dan Worker. Domain modules berbagi implementasi business rules dan invariants; modularitas ditentukan oleh ownership dan dependency boundaries, bukan jumlah service, database, directory, atau mandatory interfaces.
- **Domain ownership**: setiap domain module yang nyata memiliki operations, business invariants, dan logical state/data ownership-nya. Module lain tidak boleh melakukan mutation langsung yang melewati operasi/invariants pemiliknya.
- **Explicit module dependencies**: akses lintas-module melalui contracts/operations yang disengaja; hindari import terhadap private implementation serta circular dependencies. Cross-module read-only composition atau reporting boleh menggunakan query terkoordinasi yang jelas dan bounded jika lebih sederhana, tanpa melewati data-access/security constraints.
- **Shared infrastructure** (`internal/platform/`) hanya untuk kemampuan teknis seperti database connectivity dan adapters. Jangan tempatkan cross-domain business policy di sana.
- **Cross-domain consistency**: satu PostgreSQL transaction boleh mencakup lebih dari satu domain bila invariant membutuhkan atomicity. Domain boundaries tidak boleh merusak transactional correctness; penetapan transaction owner dan koordinasinya adalah keputusan solusi yang harus eksplisit.
- **Incremental structure**: buat module, interface, adapter, dan package hanya ketika kebutuhan nyata membenarkannya. Clean/Hexagonal Architecture, repository pattern, event bus, CQRS, atau physical schema isolation bukan kewajiban baseline.

### Separate REST API and Worker entry points

- **REST API**: existing `backend/cmd/server/` dijalankan terpisah dengan `make run`. Runner API menangani HTTP lifecycle; tidak boleh diam-diam menjalankan background Worker.
- **Worker**: `backend/cmd/worker/` dan `make worker-run` adalah **target ketika background work nyata diperlukan**, bukan implementasi yang sudah tersedia. Worker menjadi binary/process terpisah, tidak wajib menjalankan HTTP server, dan mengonsumsi domain operations yang sama.
- API dan Worker dapat hidup, gagal, restart, dan diskalakan secara independen. Keduanya tidak berbagi process memory atau locks; concurrent access, retry, duplicate execution, durable effect, dan graceful shutdown harus ditangani sesuai kebutuhan pekerjaan aktif.
- Keduanya tetap mengikuti domain invariants, data ownership, dan security/authorization boundaries. Worker bukan bypass untuk privileged operations; aturan service identity dan execution authority akan ditetapkan oleh security/authorization standards terkait.
- **Satu schema/migration authority** berlaku untuk shared PostgreSQL database. Pemilihan queue, scheduler, locking/job mechanism, worker concurrency, dan production deployment topology ditunda sampai actual workload membutuhkannya.

## Neutral baseline

The clean delivery baseline contains infrastructure wiring only: configuration, PostgreSQL connectivity, a health endpoint, and graceful shutdown. No product domain, authentication implementation, API operation, migration, or background job is considered delivered at baseline.

Historical backend implementation remains available in Git history as evidence. It must not silently establish current release scope, business semantics, or sequencing.

## Authority boundary

Product meaning starts at `docs/product/product-intent.md`. Product-design concerns start at `docs/ui-ux/`. Detailed backend behavior, threat decisions, persistence shape, API contracts, and implementation choices are derived when the selected delivery work needs them.

Do not promote a technical convenience into product truth. If engineering exposes a material missing product decision, route it back to the Human product owner rather than inventing it in code.

## Engineering decision authority

- **Project-wide standards** menetapkan reusable architecture, module/data ownership principles, shared boundaries dan constraint lintas-feature. Dokumen ini adalah owner untuk backend architecture; kontrak/API dan security standards yang terkait tetap memiliki owner masing-masing.
- **Solution Architect** boleh memilih bounded feature/domain decomposition, local interfaces, read composition, transaction design, dan mekanisme implementasi dalam effective standards dan approved scoped requirements. Pilihan feature tidak otomatis menjadi project-wide standard.
- **Implementer** bebas memilih detail internal yang tidak mengubah accepted contract, invariant, atau effective standard.
- **Human decision/escalation** diperlukan untuk material changes pada project-wide architecture, lintas-module ownership/contract, security/trust boundary, atau keputusan high-impact/sulit dibalik. Protected implementation writes tetap mengikuti explicit Human authorization pada root `AGENTS.md`; keputusan desain bukan izin implementasi.
- Jika Product meaning belum cukup jelas, kembalikan ke Product owner. Jika standard tidak memadai atau bertentangan, eskalasikan kepada owner concern; jangan jadikan satu solusi Agent sebagai project policy secara diam-diam.

## Engineering rules

- Return and handle expected errors intentionally; preserve useful error chains.
- Parameterize SQL. Never construct SQL from user-controlled values through interpolation or concatenation.
- Do not use binary floating point for monetary arithmetic.
- Do not log secrets, raw tokens, or unnecessary PII.
- Authorization, when introduced, is enforced by the backend; frontend visibility is not security authority.
- Security mechanisms must be chosen from the actual risk and data semantics of the active work, not copied from historical implementation by default.
- Simulation or fixtures must not be represented as real settlement, verification, provenance, or external evidence.

## API relationship

`api/` provides OpenAPI tooling. The baseline contract is intentionally empty. Add contract surface only when an active delivery requirement is sufficiently clear to own it. Keep generated/bundled views derived from the source contract rather than hand-maintaining parallel truth.
