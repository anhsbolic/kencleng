# Manifest — TPD-S2-002-001

> Phase: Techplan decomposition
> Author: Codex Planner
> Created / Updated: 2026-09-30
> Session: Fresh Planner Session; Session ID not exposed
> Target revision: `6891341a050982e14174ab5af132a200f24e71d9` plus durable working-tree artifacts; reopen current sources
> Workflow revision: `b179360088fbf8a7dec6c285cd8c0b8992cbb4aa`
> Work Unit: `WU-S2-002`
> Run: `TPD-S2-002-001`
> Participant: `P-S2-002-PLD-001-1`

## Parent Techplan

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` — current-effective, Human-approved `Approved` Techplan. It remains the complete authoritative spine; task files do not supersede or compress it.

## STEP 0 — gate

**YES — decomposition is genuinely useful.** TP-011 covers distinct source/spec and authored OpenAPI deliverables, with a meaningful dependency: OpenAPI must express reconciled delivery behavior. The spec and API surfaces have distinct ownership/review contexts. There is also a real execution gate because O1 `AUTHORITY_SYNC` and O11 `HUMAN_DECISION`, alongside O2–O5 owner/security detail, block only dependent contract fields while unaffected Slice 2 spec work can proceed. Separate tasks make that sequence and gate visible without splitting by document length or inventing contract boundaries.

## Splitting axis and rationale

**Dependency / sequence.** Reconcile domain spec first, then reconcile authored Donation OpenAPI against the resulting spec and owner decisions. This is the least-surprising order because `docs/spec/README.md` places domain invariants/threat/feature detail ahead of shared API contract and TP-011 §9 explicitly sequences spec reconciliation before OpenAPI. The OpenAPI task may advance unaffected areas while dependent fields remain gated.

## Task files

| File | Purpose | Execution posture |
|---|---|---|
| `01-donation-domain-spec-reconciliation.md` | Reconcile Slice 2 Donation invariants, threat model, task list/features, and only the necessary Campaign eligibility/threshold reference; preserve scoped Open Items. | Can proceed from the Approved spine and current authorities; no hard task dependency. O1/O11 and other unresolved matters remain gates for only dependent spec detail. |
| `02-donation-openapi-reconciliation.md` | Reconcile authored Donation OpenAPI and required shared source/index/generated artifacts to the approved Slice 2 contract. | Depends on Task 01. O1 `AUTHORITY_SYNC`, O11 `HUMAN_DECISION`, O2–O5 owner/security decisions/evidence, and conditional O8 audit gate affected details as described in the task. |

## Dependency graph / order

```text
Task 01 — Donation domain spec reconciliation
  └── Task 02 — Authored Donation OpenAPI reconciliation
```

- Task 01: no hard dependency task.
- Task 02: hard dependency on Task 01; also subject to external owner/authority gates recorded in its task file. No other sibling task is required.

## Shared coordination IDs

- Scope / requirements: Q1–Q11; rules R1–R11.
- Decisions: D1–D16 (D16-alt is rejected-decision history); preserve D1 exactly and do not reopen settled D1/O7.
- Risks: RISK-1–RISK-10.
- Verification: Techplan §12 obligations for R1–R11; runtime/testing evidence stays downstream.
- Active / conditional Open Items: O1–O5, O8 conditional, O11. The task set does not resolve them, accept residual risk, or claim `CONTRACT_READY`.
