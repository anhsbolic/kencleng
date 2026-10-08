# AGENTS.md — Kencleng

Root instruction map and cross-repository hard rules for AI coding agents.

Read this file before writing code. More-specific `AGENTS.md` files apply within their directory scope.

## Authority routing

```text
whole-product purpose / actors / durable product truth
→ docs/product/product-intent.md

reusable product-design / UX / visual guidance
→ docs/ui-ux/README.md

repository/backend/frontend architecture
→ docs/project/

shared API contract
→ api/ (only after an active contract exists)

stack-specific execution
→ backend/AGENTS.md
→ frontend/AGENTS.md

generic lifecycle / engineering practice
→ current Harscode authority
```

Release scope, delivery order, detailed delivery requirements, API operations, and implementation state are intentionally open at this baseline **unless a current Human-approved scoped artifact explicitly establishes them for a bounded active commitment**. Do not infer them from historical specs, contracts, code, migrations, handoffs, or Git history.

Existing historical material may be useful evidence when an active decision/work item calls for it. It is not automatic authority.

## Golden rules

- Handle expected failures intentionally; do not silently swallow errors.
- Never introduce binary floating-point arithmetic for money.
- Parameterize SQL; never interpolate user-controlled values into SQL.
- Never log secrets, raw tokens, or unnecessary PII payloads.
- Client-facing failures must not leak raw SQL, stack traces, filesystem paths, secrets, or other implementation internals.
- Backend authorization, when introduced, is explicit; frontend role gates are UX, not security authority.
- Simulation, fixtures, or operator assistance must not be presented as real settlement, independent verification, provenance, authorization, or real-world evidence.
- If product meaning is materially missing, surface the decision gap instead of inventing it downstream.

## Authority separation

Product intent owns whole-product meaning. Product-design authority owns reusable experience/visual guidance. Engineering architecture owns technical constraints. Human-approved commitment-specific product behavior and requirements may be **binding within that bounded commitment** without becoming whole-product canonical Product Authority. Detailed contracts and implementation are derived only when active delivery work needs them.

An implementation agent must follow the current scoped owning artifacts identified by the active handoff and must not change higher-level requirements merely to make code pass. If authority appears wrong or incomplete, report the contradiction/gap and route it to the owner of that concern.

## C1 engineering product-authority write boundary

For C1 engineering on `pilot/3-c1-engineering`, approved Pilot #3 product and pre-engineering artifacts are upstream read-only inputs for the Orchestrator and engineering workflow Runs.

Engineering and orchestration MUST NOT directly modify:

- `docs/product/product-intent.md`;
- `docs/product/pilot-3-business-value-loop.md`;
- `docs/product/pilot-3-actor-outcomes.md`;
- `docs/product/pilot-3-stage-3-working-model.md`;
- `docs/product/pilot-3-stage-3a-representative-scenarios.md`;
- `docs/product/pilot-3-stage-3b-commitment-sequencing.md`;
- `docs/product/pilot-3-stage-4-c1-interaction-exploration.md`;
- `docs/product/pilot-3-stage-5-c1-confirmed-behavior.md`;
- `docs/product/pilot-3-stage-6-c1-requirements.md`;
- `docs/product/pilot-3-stage-7-c1-engineering-handoff.md`;
- `docs/product/pilot-3-c1-route-retrospective.md`.

If engineering evidence exposes a contradiction, missing Product decision, or required semantic change:

1. preserve the observed evidence;
2. surface the smallest applicable Finding / Decision / Blocker;
3. route the issue back to the applicable Product / pre-engineering owner;
4. do not repair the upstream artifact from engineering or orchestration authority.

The only current exception is `docs/product/pilot-3-control-tower.md`. It may receive bounded deterministic projection updates from already-established owning engineering evidence. Such an update MUST NOT introduce Product meaning, requirements, readiness, dependency, or completion claims that are not established by their owning artifacts or evidence.

## Scope and write boundaries

Work on one coherent unit at a time. Backend and frontend production writes remain separate by default. Shared `docs/` or `api/` changes require explicit coordination with the concern they own.

A small task unexpectedly touching many unrelated files is a risk signal; investigate before expanding scope.

## High-risk write boundary

Explicit Human authorization is required before an agent introduces or materially changes implementation that owns any of these concerns, unless the current authorized work already records that approval for the same surface:

- balance, ledger, settlement, disbursement, or other money-movement transaction/locking logic;
- encryption, HMAC, signing-key, or other cryptographic key-handling core logic;
- authentication token/session/MFA core security logic;
- privileged authorization rules whose failure could grant or exercise materially elevated access.

This is a semantic safety boundary, not a commitment to historical file paths or architecture. Agents may inspect, critique, test, or propose changes without treating old implementations as current precedent.

## Security scrutiny

Security gets proportional scrutiny. Do not copy historical authentication, encryption, authorization, or data-protection mechanisms into a new active path merely because they existed before. Reuse only when current data semantics and risk justify it.

## Evidence

Distinguish verified, assumed, deferred, and not tested. Never claim a check ran when it did not.

Git history and prior Harscode artifacts are evidence/history. Promote reusable truth into the current source that owns the concern rather than treating process artifacts as project-wide precedent.
