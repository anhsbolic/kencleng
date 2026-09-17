# docs/spec — Structure & Templates

> File: `docs/spec/README.md`
>
> This guide defines the delivery/domain spec document types under `docs/spec/`, when they are created/updated, and their blank templates.
>
> `docs/product/` owns canonical whole-product truth, current MVP scope, and current MVP delivery sequencing. `docs/kencleng-agentic-workflow.md` owns Kencleng-specific orchestration/risk coordination. Harscode owns the per-feature development lifecycle.

Throughout this document, `<domain-dir>` means the actual numbered business-domain directory under `docs/spec/`, for example:

```text
1-account
2-notification
3-organization
4-campaign
5-donation
6-disbursement
```

The numeric prefix is a historical/domain organization aid. It is **not** the current MVP delivery order. Current MVP sequencing is owned by `docs/product/mvp-delivery-slices.md`.

Backend package directories remain unnumbered (for example `backend/internal/domain/account/`); do not assume spec and package paths are literal mirrors.

`docs/spec/0-foundations/` is a deliberate non-business-domain exception for cross-domain project foundation requirements/tasks that need to exist before or across business-domain feature delivery. It may use `tasks.md` plus `features/` without implying a new business domain or requiring domain invariants/threat models that do not truthfully exist.

## 1. Authority relationship

Delivery specs sit **below Product/MVP authority** and become authoritative for the detailed behavior they own after that detail has been reconciled for the active slice.

Direction:

```text
docs/product/product-overview.md
+ docs/product/mvp-scope.md
+ docs/product/mvp-delivery-slices.md
+ relevant docs/ui-ux authority
        ↓
active delivery slice
        ↓
docs/spec/<domain-dir>/...
        ↓
api/openapi/...
        ↓
implementation
```

Existing `docs/spec/` material created before Product Authority promotion remains valuable evidence. It is not automatically wrong, but it must not silently override canonical Product/MVP truth merely because it is more detailed or already implemented.

For an active slice, historical spec decisions may be classified:

```text
KEEP
ADAPT
REPLACE
DEFER
```

If a product-level contradiction is found, resolve the Product Authority concern first. If Product/MVP truth is clear, reconcile the lower-level spec. If the missing question is purely delivery/security detail, resolve it in the owning spec/architecture/contract authority rather than promoting it unnecessarily into Product Authority.

## 2. Four document types

| Type | Location | Lifespan | Written by |
|---|---|---|---|
| Domain invariant | `docs/spec/<domain-dir>/invariants.md` | Stable while applicable to delivered domain behavior | Human/project authority; an agent may draft, human review required when policy requires |
| Threat model | `docs/spec/<domain-dir>/threat-model.md` | Revised when active risk surface changes | Same as above |
| Task list | `docs/spec/<domain-dir>/tasks.md` | Domain-scoped delivery view; task status updated as relevant | Same as above |
| Feature spec | `docs/spec/<domain-dir>/features/<NN>-<fitur>.md` | Per coherent delivery work item | Same as above, per feature |

Layout remains **domain-first** for business-domain specs: every delivery spec about one business domain lives under its single numbered domain directory. Shared/cross-domain exceptions are limited to this reference file and explicit project-foundation work under `docs/spec/0-foundations/`; do not use the exception as a catch-all for unrelated documentation.

Foundation work follows the same `tasks.md` + `features/<NN>-*.md` task/requirement split where useful. Domain-only fields such as applicable domain invariants or threat-model derivation may be marked non-applicable with a concrete reason rather than inventing false domain artifacts.

See `docs/project/kencleng-repo-setup.md` for repository structure.

For a domain that genuinely needs new/revised delivery specification, a common preparation order is:

```text
reconcile active slice need
→ applicable domain invariants
→ threat model
→ task list / work unit
→ feature spec
→ shared contract
```

Do not pre-spec an entire domain merely because its directory exists or because it appears earlier in the historical numbered order.

These documents are executable delivery/domain specifications. Once reconciled for an active slice, they own applicable domain/feature behavior, acceptance criteria, invariants, and threat-model expectations. Foundation feature specs own the concrete project-foundation requirement/acceptance criteria they define. Other documents own product truth, architecture, design systems, workflow, or project status.

If a lower-level artifact conflicts with a reconciled spec on the delivery concern the spec owns, surface the contradiction. If the spec itself conflicts with canonical Product/MVP authority on product meaning or scope, the spec is the artifact that needs reconciliation.

## 3. Template: Domain Invariant

`docs/spec/<domain-dir>/invariants.md`

```markdown
# Domain Invariant — <domain name>

> Status: draft / agreed
> Last updated: <date>
> Product/MVP basis: <relevant Product Authority / active slice reference>

## Domain summary

One or two sentences: what this domain is responsible for in the active/reconciled product model.

## Invariants

Each invariant is written as a machine-verifiable statement plus when it must hold.

### INV-<domain>-01: <short name>

- **Statement**: <condition that must always hold>
- **Holds after operations**: <operations/endpoints that can affect it>
- **Verification**: <test/property/evidence that proves it>

## State machine (if applicable)

Describe valid states and allowed transitions. Any transition not listed is invalid for the reconciled delivery scope.

```text
draft -> submitted -> approved -> disbursed
                    -> rejected
```

## References

- Product/MVP owner: `docs/product/...`
- Related ERD/data-model reference if applicable: `docs/project/kencleng-erd.md#<section>`
- Historical/source evidence when relevant: ...
```

## 4. Template: Threat Model

`docs/spec/<domain-dir>/threat-model.md`

```markdown
# Threat Model — <domain name>

> Status: draft / agreed
> Last updated: <date>

## Actors & trust boundaries

| Actor | Authenticated? | Trust boundary crossed |
|---|---|---|

## STRIDE per component/endpoint

### <endpoint/component name>

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | ... | ... | ... |
| Tampering | ... | ... | ... |
| Repudiation | ... | ... | ... |
| Information disclosure | ... | ... | ... |
| Denial of service | ... | ... | ... |
| Elevation of privilege | ... | ... | ... |

Use `N/A — <reason>` when a category genuinely does not apply so it is clear the category was considered.

## Knowingly accepted residual risk

List intentionally accepted residual risks and the reason they remain accepted.
```

## 5. Template: Domain Task List

`docs/spec/<domain-dir>/tasks.md`

```markdown
# Task List — <domain name>

> Status: draft / agreed
> Last updated: <date>

## Delivery KPI / metrics

| Metric | Applies to | Threshold |
|---|---|---|
| ... | ... | ... |

## Tasks

One row per coherent project work item. A task may cover one endpoint, several tightly coupled endpoints, or another meaningful feature unit.

Implementation/session boundaries follow the active Harscode workflow. Kencleng project preconditions are defined under **Per-feature project preconditions** in `docs/kencleng-agentic-workflow.md`.

| # | Task | Endpoints | Tier | Rationale | Parallel group |
|---|---|---|---|---|---|
| 1 | ... | `METHOD /path`, ... | 0/1/2/3 | why this tier | A / serial |

For Tier-0 sub-areas, name the protected implementation area explicitly; do not rely on the tier number alone.

## Parallel / serial grouping

State which tasks can run concurrently and which must run serially. Consider overlapping files, shared tables, migration numbering, API schema, domain dependencies, and broad shared frontend components.

See **Parallelization and write scopes** in `docs/kencleng-agentic-workflow.md`.

## Status tracker

Keep domain task status current where the domain task view remains relevant to active delivery.

Cross-domain/product-slice/integration state belongs in `docs/project/kencleng-development-tracker.md`; see **Project status tracking** in `docs/kencleng-agentic-workflow.md`.

| # | Status | Notes |
|---|---|---|
| 1 | not started / in progress / blocked / verified / merged | ... |
```

## 6. Template: Feature Spec

`docs/spec/<domain-dir>/features/<NN>-<fitur>.md`

`<NN>` is the 2-digit task number from that domain's `tasks.md` when a domain task list is the applicable local index.

```markdown
# Feature Spec — <feature/work item name>

> File: `docs/spec/<domain-dir>/features/<NN>-<fitur>.md`
> Status: draft / agreed / implemented
> Risk tier: 0 / 1 / 2 / 3 (see **Project risk tiering** in `docs/kencleng-agentic-workflow.md`)
> Domain: <domain name>
> Active product slice: <slice/reference>

## Endpoint / feature surface

`<METHOD> <path>`

If the work item is not naturally one endpoint, describe the coherent feature/API surface instead of forcing a false 1:1 mapping.

## Acceptance criteria

- Given <initial condition>, When <action>, Then <expected result>

### Error cases

| Condition | Expected response |
|---|---|
| ... | `4xx`/`5xx` + specific error code |

## Applicable invariants

- `docs/spec/<domain-dir>/invariants.md#INV-<domain>-01`

## Threat breakdown

Derived from `docs/spec/<domain-dir>/threat-model.md` and narrowed to this feature/API surface.

| Threat | Mitigation at this feature's level | Evidence that proves it |
|---|---|---|
| ... | ... | `test_name` / other executable evidence |

## Risk tier & rationale

<chosen tier and why>

## Reconciliation / assumptions / open questions

Record material assumptions, historical decisions retained/rejected/deferred, or unresolved ambiguity. Do not leave an implementation-affecting assumption implicit.
```

## 7. Rules for filling these out

1. **Start from active Product/MVP need.** Do not create or expand a spec solely because historical domain order says the domain is next.
2. **Claims require evidence.** Claims that something is mitigated/tested should point to concrete executable evidence when such evidence exists. See **Evidence and verification ownership** in `docs/kencleng-agentic-workflow.md`.
3. **Keep status current.** `draft` is not a final implementation basis; `agreed` means reviewed/accepted by the appropriate human/project authority when review is required. See **Human authority** in `docs/kencleng-agentic-workflow.md`.
4. **Implementation must not rewrite requirements to make code pass.** Requirement/spec changes are separate decisions governed by root `AGENTS.md` and the authority that owns the concern.
5. **Record ambiguity rather than silently resolving it.** Use reconciliation/open-question sections whenever a material assumption is required because authority is unclear.
6. **Prefer named cross-references over workflow section numbers.** The Kencleng orchestration overlay is intentionally allowed to be compacted/reorganized; topic names are more stable than historical `§NN` references.
7. **Do not treat un-reconciled historical detail as canonical by inertia.** Preserve useful evidence, but resolve active-slice meaning deliberately.