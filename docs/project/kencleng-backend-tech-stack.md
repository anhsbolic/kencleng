# Kencleng — Backend Tech Stack

> Status: Active clean-baseline architecture
> Purpose: Record settled backend engineering choices without preselecting product capability or delivery order.

## Baseline choices

| Concern | Choice | Baseline meaning |
|---|---|---|
| Language | Go | Core backend implementation language |
| HTTP | `net/http` with Go pattern routing | Default until a demonstrated need justifies another router |
| Database | PostgreSQL | Relational persistence capability |
| Driver | `pgx` | PostgreSQL access |
| Query construction | `goqu` when query-building needs it | Parameterized query builder, not an ORM |
| Migrations | `golang-migrate` | Explicit SQL migration mechanism |
| Architecture | Domain-oriented monolith | One deployable backend; product domains emerge from real delivery needs |
| Testing | Go test + `httptest`, proportional integration/race/security checks | Evidence depth follows the active risk |
| Object storage | MinIO/S3-compatible capability available locally | Use only when an active requirement needs object storage |

A selected capability being available does not require every feature to use it. Do not create empty domain/platform layers merely to make the architecture look complete.

## Neutral baseline

The clean delivery baseline contains infrastructure wiring only: configuration, PostgreSQL connectivity, a health endpoint, and graceful shutdown. No product domain, authentication model, API operation, migration, or background job is considered delivered at baseline.

Historical backend implementation remains available in Git history as evidence. It must not silently establish current release scope, business semantics, or sequencing.

## Authority boundary

Product meaning starts at `docs/product/product-intent.md`. Product-design concerns start at `docs/ui-ux/`. Detailed backend behavior, threat decisions, persistence shape, API contracts, and implementation choices are derived when the selected delivery work needs them.

Do not promote a technical convenience into product truth. If engineering exposes a material missing product decision, route it back to the Human product owner rather than inventing it in code.

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
