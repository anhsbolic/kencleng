# Manifest — TPD-S2-002-002

> Phase: Techplan decomposition — post-approval task-snapshot reconciliation  
> Author: Codex Planner  
> Created / Updated: 2026-10-01  
> Model / reasoning: Invocation configured `gpt-6-luna` / `high`; active runtime not independently exposed  
> Session: Fresh Planner Session; Session ID not exposed  
> Target revision: `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus current durable decisions/projections; reopen live sources  
> Workflow revision: `33b03a3f62cc3aacba6534b8a011465613c64b09`  
> Work Unit: `WU-S2-002`  
> Run: `TPD-S2-002-002`  
> Participant: `P-S2-002-PLD-002-1`

## Parent Techplan

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` — current-effective, Human-approved `Approved` Techplan. It remains the complete authoritative spine; task files neither supersede nor compress it.

## STEP 0 — gate

**YES — retain the accepted decomposition.** The two outputs have independent execution/review contexts and a genuine contract dependency: Donation domain-spec reconciliation establishes the behavior that authored Donation OpenAPI must express. Keeping that boundary lets spec owners review domain invariants/threat/acceptance separately from API schema, error, header, and generated-artifact work, while keeping the hard sequence explicit. This is execution/review value, not document-length splitting.

## Splitting axis and rationale

**Dependency / sequence.** Reconcile Donation domain specs first, then reconcile authored Donation OpenAPI against the resulting behavior and Approved TP-015. The API contract depends on Task 01's reconciled domain behavior. This preserves the previously Human-accepted Task 01 → Task 02 order, dependency, task purposes, scope axes, and execution map; TP-015 introduces no material topology or dependency change.

## Task files

| File | Purpose | Execution posture |
|---|---|---|
| `01-donation-domain-spec-reconciliation.md` | Reconcile Slice 2 Donation invariants, threat model, task/features, and any necessary Campaign eligibility/threshold reference. | No hard dependency task. Human acceptance of current Donation spec drafts remains a separate parallel item; this task does not claim acceptance. |
| `02-donation-openapi-reconciliation.md` | Reconcile authored Donation OpenAPI and required shared source/index/derived artifacts to the approved Slice 2 contract. | Hard dependency on Task 01. Preserve Task 01 Build/review history as context; current spec acceptance and `CONTRACT_READY` remain unclaimed. |

## Dependency graph / order

```text
Task 01 — Donation domain spec reconciliation
  └── Task 02 — Authored Donation OpenAPI reconciliation
```

- Task 01: no hard dependency task.
- Task 02: hard dependency on Task 01; no other sibling task is required.

## Shared coordination references

- Parent spine: TP-015 Q1–Q13; R1–R14; D1–D19 and D16-alt; RISK-1–RISK-10; §12 verification; §13 Open Items.
- Task 01 specifically carries O1's settled shared amount representation plus still-open concrete parameters, O11/D19 terminal-notice direction, and current spec acceptance status.
- Task 02 additionally carries bounded O8/D18 replacement clearance; settled O4 fragment/handoff/URL-cleanup/one-way-HMAC direction; and O5 uniform `404`, identical body/header/cache behavior, and `Cache-Control: private, no-store`. Runtime/security controls and empirical evidence stay downstream as TP-015 assigns.
- Active items remain O1 concrete amount parameters, O2, O3, O4 implementation controls/evidence and risk gate, O5 empirical parity/timing/abuse evidence, and independent Human acceptance of the current Donation drafts. O1 shared representation, O8 bounded compatibility clearance, O11/D19 direction, and O4/O5 contract directions remain resolved decision history, not open alternatives.
- Neither task nor this manifest changes accepted topology, resolves an Open Item, accepts residual risk, records progress status, or claims `CONTRACT_READY`.
