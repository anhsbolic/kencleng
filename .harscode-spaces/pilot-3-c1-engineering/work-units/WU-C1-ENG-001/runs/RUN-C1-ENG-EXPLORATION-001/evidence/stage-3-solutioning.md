# Exploration Stage 3 — C1 Engineering Direction

## Provenance

- **Phase / Stage:** Exploration / Stage 3 — Solutioning
- **Author:** `PARTICIPANT-C1-ENG-EXPLORER-001` (ephemeral Participant)
- **Created:** 2026-10-09
- **Model / Reasoning:** `gpt-6-luna` / `medium` (per Run Invocation)
- **Session:** `SESSION-C1-ENG-EXPLORER-001`
- **Target revision:** `e9cb1f31f031ed3905181eb86ad3c3e3af4e248f`
- **Workflow revision:** `d882b35c88ae7b4adff19abba991e0f2f9d7b7ac`
- **Run:** `RUN-C1-ENG-EXPLORATION-001`
- **Work Unit:** `WU-C1-ENG-001`

## Direction

This is an Exploration recommendation for later Techplan synthesis, not an approved implementation plan or authorization to change privileged authorization behavior.

Treat C1 establishment as one product-semantic success composed of a valid Organization context and a durably attributable initial Kencleng Owner relationship. For the current PostgreSQL monolith baseline, prefer creating those records through one bounded database transaction and expose success only after commit. Require the initiating person to be bound to the action before it can succeed; C1 may allow initiation before identity is established, but must not make external legal-authority verification a prerequisite. Persist enough relationship/provenance meaning to support later inspection without upgrading Organization-provided information to reviewed or externally verified truth.

This direction does not choose API shape, schema, UI flow, exact authentication mechanism, role/permission model, or a technical conflict-detection policy. Those remain for Techplan or a Product owner where applicable. Before implementation of an initial Owner grant or other privileged authorization rule, obtain explicit Human authorization for that implementation surface as required by root `AGENTS.md`.

## Options considered

### Organization + initial Owner consistency

**Recommended: one database transaction for the successful C1 consequence.** A transaction fits the existing single PostgreSQL-backed monolith and directly supports Stage 6 ER3.1's observable invariant. If either record cannot be committed, the operation should not expose successful establishment. This avoids adding pending/recovery semantics that C1 does not require. The recommendation is conditional on both records being owned by the same database transaction boundary; the current baseline has no C1 persistence to verify that assumption.

**Alternative considered: staged/pending establishment with later completion or recovery.** This can support interruption and retry, but introduces additional persisted states and UX/API semantics. Stage 5 B6 and Stage 6 ER6 make recovery optional, so adding it now would create complexity without a current requirement. Keep it available only if a later concrete delivery need justifies it and preserve the incomplete-versus-success distinction.

### Attribution and authentication timing

**Recommended semantic boundary: no successful establishment without a durable person attribution.** A person may start earlier, but at the success boundary the resulting establishment and initial Owner relation must both identify the person. This follows Stage 5 B1 and Stage 6 ER1.1–ER1.2.

**Alternative considered: require sign-in before the person can begin.** This would satisfy attribution but unnecessarily removes the explicitly preserved design freedom in Stage 6 XR1.3 and may imply a more restrictive flow than C1 requires. No exact authentication method or sign-in timing is selected. Do not invent external-authority proof as an authentication substitute or prerequisite.

### Known unresolved representation conflict

**Recommended boundary: do not produce normal uncontested success when a relevant unresolved conflict is already known to the system.** C1 must not add its own Organization matching, duplicate detection, legal-authority adjudication, merge, transfer, or dispute-resolution mechanism. No current conflict signal or source exists in the baseline. If a future implementation needs to discover or classify conflicts to satisfy ER7.1, stop and route the missing product/authority decision to the appropriate owner before defining that mechanism.

### Durable resulting meaning

Preserve the Organization-to-Owner relationship, person attribution, and the distinction between internal Kencleng Owner authority and external/legal verification in durable product state. Keep later inspection available beyond transient success messaging. The exact storage representation and user-facing surface remain open; this is a requirement boundary rather than a schema or screen prescription.

## Material rejected directions and rationale

- **Treat a frontend success message as the success source:** rejected because Stage 5 B4 / Stage 6 XR4.1–XR4.3 require durable resulting meaning beyond the success moment.
- **Create the Owner role after Organization creation in a separate, user-visible success step:** rejected as the default because it can expose one side as successfully established without the other, contrary to Stage 5 B3 / Stage 6 ER3.1. A staged workflow would need an explicit incomplete state and must not claim success until both are valid.
- **Make legal-representative verification or Organization review a C1 precondition:** rejected because it changes confirmed product meaning (Stage 5 B1/B5; Stage 6 ER1.3/ER2.2).
- **Build conflict matching or legal adjudication inside C1:** rejected because Stage 5 B7 and Stage 6 ER7.2 explicitly leave these mechanics outside C1.

## Material questions and constraints

1. **Privileged authorization implementation:** The initial Owner relationship carries internal authority. The product behavior is approved, but this Run does not record Human authorization for implementation of privileged authorization rules. Obtain that authorization before Build changes this surface. This is a safety gate on implementation, not a product behavior gap.
2. **Conflict knowledge source:** No current source of a known unresolved representation conflict exists in the clean baseline. This is not a blocker to the recommended boundary, since C1 does not require conflict discovery. If an implementation design requires a new source/detection policy, route that decision upstream rather than inventing semantics in engineering.
3. **Authentication mechanism:** The requirement fixes attribution at success while leaving exact authentication timing/implementation open. Select an authentication design only in the appropriate authorized technical work; do not infer one from historical code.

No external decision is needed to complete this Exploration. No option here is a Human-approved architecture decision.

## Context pointers

- Stage 2 evidence: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-EXPLORATION-001/evidence/stage-2-gap-analysis.md`
- Binding behavior: `docs/product/pilot-3-stage-5-c1-confirmed-behavior.md` (B1–B7)
- Binding requirements: `docs/product/pilot-3-stage-6-c1-requirements.md` (ER1–ER7, §5)
- Handoff and scope boundary: `docs/product/pilot-3-stage-7-c1-engineering-handoff.md` (§§4–9)
- Repository hard rule: `AGENTS.md` (C1 authority boundary; High-risk write boundary)
- Relevant routed guidance consulted: `../harscode-workspace/best-practices/go/authorization-and-idor.md`; `../harscode-workspace/best-practices/go/role-and-privilege-separation.md`; `../harscode-workspace/best-practices/postgresql/transactions-and-locking.md`

## Verification state

No tests, builds, runtime checks, or implementation changes were performed in this Run. The proposed transaction and attribution direction is analysis only and has not been validated against a C1 implementation because none exists at the target revision.
