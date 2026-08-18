# docs/spec — Structure & Templates

> File: `docs/spec/README.md`
> This guide explains the four types of spec documents that live in
> `docs/spec/`, when each one is created/updated, and the blank
> template for each type. See `docs/kencleng-agentic-workflow.md` for
> the reasoning/principles behind this structure.

## 1. Four document types

| Type | Location | Lifespan | Written by |
|---|---|---|---|
| Domain invariant | `docs/spec/<domain>/invariants.md` | Once per domain, stable | Anhar (with drafting help from an agent), must be human-reviewed |
| Threat model | `docs/spec/<domain>/threat-model.md` | Once per domain, revised on domain-level changes | Same as above |
| Task list | `docs/spec/<domain>/tasks.md` | Once per domain, status updated as tasks complete | Same as above — see §3.1 |
| Feature spec | `docs/spec/<domain>/features/<NN>-<fitur>.md` | New for each vertical slice/endpoint | Same as above, per feature |

Layout is **domain-first**: every document about one domain lives
under a single `docs/spec/<domain>/` folder (`account`, `notification`,
`organization`, `campaign`, `donation`, `disbursement`), mirroring
`backend/internal/domain/<domain>/`. This file (`docs/spec/README.md`)
is the one exception — it's shared/cross-domain reference material, so
it stays at `spec/` root. See `kencleng-repo-setup.md` §2.1 for the
full rationale.

Order of creation for a new domain: **domain invariant → domain
threat model → task list → feature spec per endpoint**, because
feature specs always reference back to the documents above rather
than redefining anything, and tiering in the task list depends on
what the invariants/threat model already surfaced.

All of these documents are **executable specs for the agent** — as
opposed to the documents in `docs/project/`, which are narrative,
meant for humans to read. If a `docs/project/*` document and a
`docs/spec/*` document ever disagree for the same domain/feature,
`docs/spec/*` wins (since it's the hardened, more precise version) —
but the conflict itself must be resolved explicitly (noted down, not
silently ignored on one side).

## 2. Template: Domain Invariant

`docs/spec/<domain>/invariants.md`

```markdown
# Domain Invariant — <domain name>

> Status: draft / agreed
> Last updated: <date>

## Domain summary

One or two sentences: what this domain is responsible for.

## Invariants

Each invariant is written as a machine-verifiable statement (not
narrative), plus when it must hold.

### INV-<domain>-01: <short name>

- **Statement**: <condition that must always hold, phrased so it can
  become an assertion — e.g. an equation, a bound, a state-transition
  rule>
- **Holds after operations**: <list of operations/endpoints that can
  affect this invariant>
- **Verification**: <name of the test/property test that proves this>

(repeat for each invariant)

## State machine (if applicable)

If the domain has an entity with changing status (e.g. campaign,
disbursement), describe the valid states and allowed transitions. Any
transition not listed is invalid and must be rejected at the code
level.

```
draft -> submitted -> approved -> disbursed
                    -> rejected
```

## References

- Related ERD: `docs/project/kencleng-erd.md#<section>`
- Related business process: `docs/project/kencleng-business-process-overview.md#<section>`
```

## 3. Template: Threat Model

`docs/spec/<domain>/threat-model.md`

```markdown
# Threat Model — <domain name>

> Status: draft / agreed
> Last updated: <date>

## Actors & trust boundaries

List the actors that interact with this domain (e.g. guest donor,
registered donor, org owner/staff, admin, external systems), and
where data crosses a trust boundary (e.g. an unauthenticated guest
can write to the donations table).

| Actor | Authenticated? | Trust boundary crossed |
|---|---|---|

## STRIDE per component/endpoint

For each significant component/endpoint in this domain:

### <endpoint/component name>

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | ... | ... | ... |
| Tampering | ... | ... | ... |
| Repudiation | ... | ... | ... |
| Information disclosure | ... | ... | ... |
| Denial of service | ... | ... | ... |
| Elevation of privilege | ... | ... | ... |

Leave rows that are genuinely not applicable blank (don't force an
entry that doesn't apply), but write "N/A — <reason>" explicitly
rather than deleting the row silently, so it's clear it was
considered rather than missed.

## Knowingly accepted residual risk

Summary of risks that have been identified but are intentionally not
mitigated further in v1, with the reason (e.g. "over-engineering for
a sandbox project," "no demonstrated need yet").
```

## 4. Template: Domain Task List

`docs/spec/<domain>/tasks.md`

```markdown
# Task List — <domain name>

> Status: draft / agreed
> Last updated: <date>

## Delivery KPI / metrics

The concrete, measurable bar every task in this domain must clear
before merge (proposed once, applies to all tasks below unless a task
overrides it with a stricter requirement).

| Metric | Applies to | Threshold |
|---|---|---|
| ... | ... | ... |

## Tasks

One row per vertical slice (a task may group several tightly-coupled
endpoints — see `kencleng-agentic-workflow.md` §11 "scope fencing" for
what counts as one slice).

| # | Task | Endpoints | Tier | Rationale | Parallel group |
|---|---|---|---|---|---|
| 1 | ... | `METHOD /path`, ... | 0/1/2/3 | why this tier | A / serial |

For any task with a Tier 0 sub-area embedded inside an otherwise
higher-tier task (e.g. JWT/TOTP core logic inside an agent-generated
endpoint), name the specific sub-area and note it needs file-path
fencing in `AGENTS.md` — don't leave it implicit in the tier number
alone.

## Parallel / serial grouping

State which tasks can run concurrently (different agent sessions,
non-overlapping files/tables/migration numbers) and which must run
serially (shared tables, shared migration sequence), per
`kencleng-agentic-workflow.md` §12's parallelization note.

## Status tracker

Update as work progresses — this is the lightweight, domain-scoped
substitute for a global tracker (no cross-domain development-phase
tracker exists yet, see `kencleng-agentic-workflow.md` §16).

| # | Status | Notes |
|---|---|---|
| 1 | not started / in progress / gates passed / human-reviewed / merged | ... |
```

## 5. Template: Feature Spec

`docs/spec/<domain>/features/<NN>-<fitur>.md`

`<NN>` is the 2-digit task number from that domain's `tasks.md` (e.g.
`01-register-email-verification.md`) — feature files within a domain
are inherently ordered (dependency order, parallel groups from
`tasks.md`), so the filename should carry that order rather than
leaving it only discoverable by opening `tasks.md`.

```markdown
# Feature Spec — <feature/endpoint name>

> File: `docs/spec/<domain>/features/<NN>-<fitur>.md`
> Status: draft / agreed / implemented
> Risk tier: 0 / 1 / 2 / 3 (see kencleng-agentic-workflow.md §4)
> Domain: <domain name>

## Endpoint

`<METHOD> <path>`

## Acceptance criteria

- Given <initial condition>, When <action>, Then <expected result>
- (repeat for each happy-path scenario)

### Error cases

| Condition | Expected response |
|---|---|
| ... | `4xx`/`5xx` + specific error code |

## Applicable invariants

Reference related invariants, don't redefine them here:

- `docs/spec/<domain>/invariants.md#INV-<domain>-01`

## Threat breakdown

Derived from `docs/spec/<domain>/threat-model.md`, narrowed down to
this specific endpoint:

| Threat | Mitigation at this endpoint's level | Test that proves it |
|---|---|---|
| ... | ... | `test_name` |

## Risk tier & rationale

The chosen tier and why (e.g. "Tier 1 — touches campaign balance,
needs a property test + human review").

## Assumptions / open questions

Anything the agent assumed during implementation if the spec was
ambiguous at some point — must be revisited once a decision is made,
not left blank forever if a gap turns out to exist.
```

## 6. Rules for filling these out (apply to all four document types)

1. **Every claim of "this is mitigated/tested" must point to a
   concrete test name** — not a narrative sentence with no
   reference (see `kencleng-agentic-workflow.md` §9 on agent
   honesty).
2. **The status in the header must be kept up to date** — draft
   means it isn't yet a valid basis for final implementation; agreed
   means it has been human-reviewed per
   `kencleng-agentic-workflow.md` §10.
3. **The implementing agent must not edit this document** to make its
   own code pass (see §11, fencing, in
   `kencleng-agentic-workflow.md`) — changes to spec are a separate
   decision that goes through a human.
4. **Ambiguity is recorded, not silently resolved** — the
   "Assumptions / open questions" section of a feature spec must be
   filled in whenever the agent makes any assumption because the spec
   was unclear.