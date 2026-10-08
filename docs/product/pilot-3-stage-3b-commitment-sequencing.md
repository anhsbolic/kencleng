# Kencleng Pilot #3 — Stage 3B Commitment Sequencing

> **Status:** WORKING ARTIFACT — Pilot #3 Stage 3B in progress
> **Human status:** Anhar-approved through revised C1 representation model
> **Authority:** Not canonical Product Authority
> **Working baseline:** `pilot/3-clean-delivery-baseline`
> **Related method:** `docs/product/pilot-3-stage-3-working-model.md`
> **Scenario source:** `docs/product/pilot-3-stage-3a-representative-scenarios.md`

## C1 — Legitimate Organization representation

### Blocker #1 — minimum truth for legitimate representation — SUPERSEDED

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


### Blocker #2 — how initial authority becomes legitimate — SUPERSEDED

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


## Revised C1 representation model — superseding decision

**Status:** WORKING DECISION — Anhar-approved.  
This section supersedes the earlier blocker #1 and blocker #2 decisions where they conflict.

### Decision

The first person who establishes an Organization in Kencleng becomes the initial **Kencleng Organization Owner by product rule**.

```text
person establishes Organization in Kencleng
→ initial Kencleng Organization Owner relationship exists
```

The meaning of Owner is internal to Kencleng:

> **Owner is an authority-bearing Organization role within the Kencleng Organization context. It is not proof that the person is the legal owner, chair, authorized signatory, legally verified representative, or otherwise independently validated real-world authority of the Organization.**

Therefore:

```text
Kencleng Owner
≠ legal / administrative ownership of the real-world Organization
≠ independently verified external authority
≠ Organization review outcome
```

The first registrant does **not** need a separate evidence-review process merely to receive the initial Kencleng Owner role.

### Owner / Staff minimum meaning

The previously approved Owner / Staff distinction remains:

```text
Owner
→ Organization-level authority-bearing role
  within the Kencleng Organization context

Staff
→ bounded operational Organization role
  acting within legitimately established authority
```

Owner is not a universal approver or superuser. Staff is not merely "Owner with fewer buttons."

Detailed permission, invitation, delegation, revocation, ownership-transfer, and multi-Owner rules remain intentionally open.

### Organization evidence and truth class

Organization information or evidence submitted by Owner / Staff remains:

> **Organization-provided information / evidence**

until a separate bounded review establishes additional meaning.

```text
evidence uploaded by Organization Representative
≠ evidence independently validated
≠ Organization reviewed
≠ Campaign curated
```

Responsibility for the truthfulness of Organization-provided assertions primarily attaches to the Organization, while Kencleng preserves provenance and does not silently upgrade those assertions into platform-known or reviewed fact.

### Organization review remains distinct

This simplification does **not** remove Organization review.

```text
Organization establishment / internal Kencleng ownership
≠ Organization review
≠ Campaign curation
```

Organization review remains a distinct bounded decision. Its exact positive meaning and its relationship to Campaign preparation, curation, visibility, and donation eligibility remain OPEN and should be resolved only when the affected commitment requires it.

### Public / external interpretation invariant

Kencleng must not present the internal Owner role as evidence that the person is a legally valid or independently verified real-world authority of the Organization.

If the product later needs to make an external-authority claim, that claim requires separately established evidence / review semantics.

## C1 re-gate after revised representation model

**C1 product claim**

> Organization can have a legitimate representation context with clear representative authority and provenance.

Re-gate:

| Gate | Result | Reason |
| --- | --- | --- |
| Meaningful | PASS | Organization can genuinely act through an accountable internal Kencleng authority context |
| Semantically ready enough | PASS | Owner / Staff minimum meaning and initial Owner origin are now sufficiently clear |
| Dependency-closed enough | PASS for product semantics | No Organization review is required merely to establish the internal Kencleng representation context |
| Truthful | PASS | Internal Kencleng authority is explicitly not represented as verified real-world authority |
| Bounded | PASS | Detailed permission / delegation mechanics remain outside this commitment |

**C1: ELIGIBLE for the Stage 4–7 delivery loop.**

At the time of this re-gate, C1 became eligible to leave Stage 3B. It subsequently entered and completed the Stage 4–7 **pre-engineering** route.

That later route completion does **not** mean C1 has been implemented or delivered.

### Frontier reconciliation after C1 departure

C2 / C3 have **not** been re-gated or sequenced in this owning Stage 3B artifact after the revised C1 model.

Therefore, until future Stage 3B work explicitly evaluates them:

- C2 / C3 remain **PARKED / UNRESOLVED frontier placeholders**;
- do not assign them BLOCKED, WAITING, or ELIGIBLE status from the Control Tower alone;
- do not infer a commitment-level dependency graph beyond what the owning scenario / Stage 3B artifacts explicitly establish.

Future frontier work must re-evaluate the relevant commitment meaning and dependencies before assigning a sequencing status.
