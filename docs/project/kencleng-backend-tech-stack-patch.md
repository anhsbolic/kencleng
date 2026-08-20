# Patch — API Contract & Codegen section update

> Target file: `docs/project/kencleng-backend-tech-stack.md`
> Replace the existing "## API Contract & Codegen [RESOLVED — Step 2]"
> section (through the blank line before "## Encryption Key
> Management") with the block below. Two changes from the current
> text: (1) `organisasi` → `organization` in the composite-endpoint
> note, per the 2026-08-19 naming-convention lock; (2) a new
> "File layout" subsection documenting the 2026-08-20 split into
> `api/openapi/*.yaml` + generated bundle.

---

## API Contract & Codegen [RESOLVED — Step 2, amended 2026-08-20]

**Format**: OpenAPI 3.x, spec-first — `api/openapi.yaml` is the single
source of truth for all HTTP endpoints, hand-authored *before*
implementation (not generated from code).

**File layout [RESOLVED — 2026-08-20]**: the spec is authored as
per-domain source files under `api/openapi/` (`account.yaml`,
`organization.yaml`, `campaign.yaml`, `donation.yaml`,
`disbursement.yaml`, `notification.yaml`, plus `common.yaml` for
shared parameters/responses/envelope schemas and `index.yaml` as the
root document), and bundled into the single `api/openapi.yaml` via
`@redocly/cli` (`cd api && npm run bundle`). `api/openapi.yaml` remains
the file frontend codegen and backend developers reference directly —
it's now a **generated artifact**, not hand-edited. Adopted because the
single-file spec grew past ~4,000 lines once all 6 domains were
specced, which was expensive for an agent to re-read in full for a
task scoped to one domain. See `api/README.md` for the full editing
workflow.

**Design philosophy** (agreed 2026-07-20, prior to the format decision
above): domain/resource-driven REST as the default endpoint shape,
with a justified composite-endpoint exception for `/campaigns/{id}` —
this page aggregates campaign + organization + progress data in one
call to avoid client-side N+1 fetches on what's effectively the
platform's primary conversion surface.

**Backend usage — documentation/contract only, no codegen.** Handlers,
request/response structs, and validation in
`internal/transport/http/` are 100% hand-written. The spec is kept in
sync by developer discipline (updating the relevant
`api/openapi/{domain}.yaml` — and re-running the bundle — as part of
the same change as the handler), not by tooling.
- Rejected: `oapi-codegen` (or similar Go server-stub/type generation
  from the spec). Would guarantee spec/implementation parity at
  compile time, but adds a build step and a generated-code layer that
  isn't justified yet for a single-developer sandbox.
- **Accepted trade-off**: no compiler-enforced guarantee that
  `openapi.yaml` matches the actual Go implementation — accuracy
  depends entirely on discipline. Fine for a solo learning project;
  would need revisiting (via `oapi-codegen` or contract testing) if
  this ever became multi-developer or long-lived in production.

**Frontend usage — types only.** Request/response types generated via
`openapi-typescript` from `api/openapi.yaml` (the generated bundle);
fetch functions in `lib/api/` stay hand-written, just typed against
the generated interfaces. See `kencleng-frontend-tech-stack.md`
structural principles for the full FE-side detail.
