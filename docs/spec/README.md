# docs/spec — Structure & Templates

> File: `docs/spec/README.md`
>
> This guide defines the four spec document types under `docs/spec/`, when they are created/updated, and their blank templates.
>
> `docs/kencleng-agentic-workflow.md` owns Kencleng-specific domain preparation, risk tiering, orchestration, and project-level coordination. Harscode owns the per-feature development lifecycle.

Throughout this document, `<domain-dir>` means the actual numbered domain directory under `docs/spec/`, for example:

```text
1-account
2-notification
3-organization
4-campaign
5-donation
6-disbursement
```

The numeric prefix expresses project/domain order. Backend package directories remain unnumbered (for example `backend/internal/domain/account/`); do not assume spec and package paths are literal mirrors.

## 1. Four document types

| Type | Location | Lifespan | Written by |
|---|---|---|---|
| Domain invariant | `docs/spec/<domain-dir>/invariants.md` | Once per domain, stable | Human/project authority; an agent may draft, human review required |
| Threat model | `docs/spec/<domain-dir>/threat-model.md` | Once per domain, revised on domain-level changes | Same as above |
| Task list | `docs/spec/<domain-dir>/tasks.md` | Once per domain, task status updated as work progresses | Same as above |
| Feature spec | `docs/spec/<domain-dir>/features/<NN>-<fitur>.md` | New for each coherent feature/work item | Same as above, per feature |

Layout is **domain-first**: every spec about one domain lives under its single numbered domain directory. This file is the exception because it is shared/cross-domain reference material.

See `docs/project/kencleng-repo-setup.md` for repository structure.

Order of creation for a new domain:

```text
domain invariant
→ threat model
→ task list
→ feature specs for coherent work items
```

These documents are executable domain/feature specifications. They own domain behavior, acceptance criteria, invariants, and threat-model expectations. Other documents own architecture, design, workflow, or project status.

If another document conflicts with `docs/spec/*` on domain/feature behavior, the spec is authoritative and the contradiction must be surfaced explicitly. Do not use this as a blanket precedence claim over documents that own different concerns.

## 2. Template: Domain Invariant

`docs/spec/<domain-dir>/invariants.md`

```markdown
# Domain Invariant — <domain name>

> Status: draft / agreed
> Last updated: <date>

## Domain summary

One or two sentences: what this domain is responsible for.

## Invariants

Each invariant is written as a machine-verifiable statement plus when it must hold.

### INV-<domain>-01: <short name>

- **Statement**: <condition that must always hold>
- **Holds after operations**: <operations/endpoints that can affect it>
- **Verification**: <test/property/evidence that proves it>

## State machine (if applicable)

Describe valid states and allowed transitions. Any transition not listed is invalid.

```text
draft -> submitted -> approved -> disbursed
                    -> rejected
```

## References

- Related ERD: `docs/project/kencleng-erd.md#<section>`
- Related business process: `docs/project/kencleng-business-process-overview.md#<section>`
```

## 3. Template: Threat Model

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

## 4. Template: Domain Task List

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

One row per coherent project work item. A task may cover one endpoint, several tightly coupled endpoints, or another meaningful feature unit. Implementation/session boundaries follow the active Harscode workflow; Kencleng project preconditions are described in `docs/kencleng-agentic-workflow.md` §7.

| # | Task | Endpoints | Tier | Rationale | Parallel group |
|---|---|---|---|---|---|
| 1 | ... | `METHOD /path`, ... | 0/1/2/3 | why this tier | A / serial |

For Tier-0 sub-areas, name the protected implementation area explicitly; do not rely on the tier number alone.

## Parallel / serial grouping

State which tasks can run concurrently and which must run serially. Consider overlapping files, shared tables, migration numbering, API schema, domain dependencies, and broad shared frontend components. See `docs/kencleng-agentic-workflow.md` §20.

## Status tracker

Keep domain task status current. Cross-domain backend/frontend/integration state belongs in `docs/project/kencleng-development-tracker.md` as described by `docs/kencleng-agentic-workflow.md` §24.

| # | Status | Notes |
|---|---|---|
| 1 | not started / in progress / blocked / verified / merged | ... |
```

## 5. Template: Feature Spec

`docs/spec/<domain-dir>/features/<NN>-<fitur>.md`

`<NN>` is the 2-digit task number from that domain's `tasks.md`.

```markdown
# Feature Spec — <feature/work item name>

> File: `docs/spec/<domain-dir>/features/<NN>-<fitur>.md`
> Status: draft / agreed / implemented
> Risk tier: 0 / 1 / 2 / 3 (see `docs/kencleng-agentic-workflow.md` §4)
> Domain: <domain name>

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

## Assumptions / open questions

Record material assumptions or unresolved ambiguity. Do not leave an implementation-affecting assumption implicit.
```

## 6. Rules for filling these out

1. **Claims require evidence.** Claims that something is mitigated/tested should point to concrete executable evidence when such evidence exists. See `docs/kencleng-agentic-workflow.md` §22–§23.
2. **Keep status current.** `draft` is not a final implementation basis; `agreed` means reviewed/accepted by the appropriate human/project authority. See `docs/kencleng-agentic-workflow.md` §19.
3. **Implementation must not rewrite requirements to make code pass.** Requirement/spec changes are separate decisions governed by root `AGENTS.md` §4 and the appropriate human/project authority.
4. **Record ambiguity rather than silently resolving it.** Use `Assumptions / open questions` whenever a material assumption is required because the specification is unclear.
