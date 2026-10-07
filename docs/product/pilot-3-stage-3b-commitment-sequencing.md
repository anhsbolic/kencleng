# Kencleng Pilot #3 — Stage 3B Commitment Sequencing

> **Status:** WORKING ARTIFACT — Pilot #3 Stage 3B in progress
> **Human status:** Anhar-approved through C1 blocker #1
> **Authority:** Not canonical Product Authority
> **Working baseline:** `pilot/3-clean-delivery-baseline`
> **Related method:** `docs/product/pilot-3-stage-3-working-model.md`
> **Scenario source:** `docs/product/pilot-3-stage-3a-representative-scenarios.md`

## C1 — Legitimate Organization representation

### Blocker #1 — minimum truth for legitimate representation

**Working decision — Anhar-approved**

A person may be treated as a legitimate Organization Representative only when there is an **explicit, attributable relationship between that person and the Organization** that carries:

- a role / authority context sufficient to explain what they may legitimately do for the Organization; and
- provenance sufficient for Kencleng to explain why that representation is treated as legitimate.

```text
self-declared affiliation
≠ legitimate representative authority

explicit attributable person ↔ Organization relationship
+ role / authority context
+ understandable provenance
→ legitimate representation context
```

Boundaries:

```text
legitimate representation
≠ Organization review

Organization review
≠ representative authority grant

authority to act
≠ Reviewer judgment
≠ provenance itself
```

This decision intentionally does **not** define invitation / onboarding flow, authentication mechanism, legal-document requirements, KYC-like verification, exact Owner / Staff permission matrix, or Organization-review procedure.

### C1 status after blocker #1

C1 remains **BLOCKED**.

Next blocker:

> How does the initial explicit authority relationship become legitimate, especially for the first representative / initial Owner, without relying on Organization review or unsupported self-assertion?


### Blocker #2 — how initial authority becomes legitimate

**Decision question**

> How does the initial explicit authority relationship become legitimate, especially for the first representative / initial Owner, without relying on Organization review or unsupported self-assertion?

**Working decision — Anhar-approved**

Use an **evidence-backed bootstrap + derived authority** model.

Initial / bootstrap representation:

```text
person claims Organization relationship
+ independently attributable basis
  that does not rely only on the person's own assertion
→ bounded representation-establishment judgment
→ initial legitimate authority relationship
```

Derived representation:

```text
existing legitimate authority-bearing representative
→ explicitly establishes another
   person ↔ Organization authority relationship
   within their own authority
→ new legitimate representative relationship
```

Therefore legitimate representative authority may originate through two distinct paths:

1. **Bootstrap authority** — requires non-self-referential provenance.
2. **Derived authority** — derives from an already legitimate authority relationship whose scope permits that grant / delegation.

Important boundaries:

```text
bootstrap authority
≠ purely self-declared authority

representation establishment
≠ Organization review

Platform Operator assistance
≠ source of representative semantic authority

independently attributable basis
≠ automatically independent identity verification / KYC
```

This decision intentionally does **not** yet define the exact bootstrap evidence or operational procedure. Examples such as legal documents, Organization-controlled channels, external registries, manual review, or combinations remain open until materially required.

### C1 re-gate after blocker #2

C1 is closer to semantic readiness but remains **BLOCKED**.

The next material question is:

> What minimum role / authority distinction must be real and understandable so that Owner and Staff are meaningfully different business roles rather than labels, without prematurely defining a full permission matrix?
