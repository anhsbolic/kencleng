# Kencleng Pilot #3 — Stage 7 C1 Durable Engineering Handoff

> **Status:** WORKING ARTIFACT — C1 Stage 7 complete  
> **Human status:** Anhar-approved durable engineering handoff; C1 pre-engineering route complete  
> **Authority:** Not canonical Product Authority  
> **Working baseline:** `pilot/3-clean-delivery-baseline`  
> **Control Tower:** `docs/product/pilot-3-control-tower.md`

## 1. Handoff purpose

This is the durable pre-engineering handoff for:

**C1 — Legitimate Organization representation**

Engineering should be able to begin C1 work from the durable artifacts referenced here without needing the original product conversation.

This handoff packages existing confirmed behavior and requirements. It does not add new Product Truth or prescribe implementation architecture.

## 2. Recommended read order

1. [Product Intent](./product-intent.md) — upstream Product Authority.
2. [Stage 5 — C1 Confirmed Product Behavior](./pilot-3-stage-5-c1-confirmed-behavior.md) — confirmed C1 behavior and invariants.
3. [Stage 6 — C1 Requirements](./pilot-3-stage-6-c1-requirements.md) — Experience Requirements + Engineering Requirements derived from the same behavior.
4. [Stage 4 — C1 Interaction Exploration](./pilot-3-stage-4-c1-interaction-exploration.md) — interaction evidence and pressure tests when additional context is needed.
5. [Stage 3B — Commitment Sequencing](./pilot-3-stage-3b-commitment-sequencing.md) — why C1 became eligible and its dependency boundary.

The [Pilot #3 Control Tower](./pilot-3-control-tower.md) is a progress dashboard only and is **not Product Authority**.

## 3. Product claim being handed off

> A person can establish an Organization context in Kencleng and become its initial Kencleng Organization Owner, with clear internal authority meaning and provenance, without Kencleng implying that the person has independently verified real-world legal authority.

## 4. What must remain true

Engineering and design must preserve these confirmed C1 invariants:

1. **Kencleng Owner is internal product authority, not proof of external legal authority.**
2. **Organization establishment and initial Owner formation are one successful product consequence.**
3. **Attribution exists before the resulting Owner relationship can be valid.**
4. **Material consequence is available before establishment becomes effective.**
5. **Organization-provided information / evidence is not silently upgraded by establishment.**
6. **Incomplete / conflicted state is not presented as successful establishment.**
7. **Resulting Organization / Owner meaning remains durably inspectable after the success moment.**
8. **C1 remains distinct from Organization review, Campaign curation, external-authority verification, and later fund-use verification semantics.**

## 5. Confirmed product behavior

### C1-B1 — Start without external-authority proof

A person may begin Organization establishment without C1 requiring or implying proof of real-world legal / representative authority.

Before successful establishment, the resulting action and initial Owner relationship must be durably attributable to a person.

### C1-B2 — Consequence before effect

Before establishment becomes effective, Kencleng must make clear that:

- a Kencleng Organization context will be established;
- the initiating person will become its initial Kencleng Owner;
- Owner is internal Kencleng Organization authority, not independently verified external authority;
- Organization information / evidence involved remains Organization-provided unless separately reviewed.

### C1-B3 — One successful product consequence

Successful C1 establishment means:

```text
valid Organization context
+
valid attributable initial Kencleng Owner relationship
```

The product must not expose either side as successfully formed while the other side is invalid.

### C1-B4 — Durable resulting meaning

After success, the person must be able to determine later:

- which Organization context exists;
- that they hold its Kencleng Owner role;
- that this is internal Kencleng authority;
- that establishment itself did not complete Organization review or independently verify external authority.

### C1-B5 — Truth class is not upgraded

C1 preserves:

```text
Organization-provided information / evidence
≠ Organization review
≠ independent external-authority verification
≠ Campaign curation
```

C1 does not invent an Organization-review status taxonomy.

### C1-B6 — Incomplete is not success

If required attribution is unavailable, or Organization + initial Owner formation has not fully succeeded:

- establishment must not be represented as successful;
- incomplete state must remain distinguishable from successful establishment;
- recovery may exist, but is not required by C1.

### C1-B7 — Known unresolved representation conflict

C1 does not guarantee real-world Organization uniqueness.

If Kencleng already has a known unresolved representation conflict relevant to the Organization being established:

- normal uncontested-success meaning must not be produced;
- a new initial Owner relationship must not be granted as though representation were uncontested.

C1 does not define conflict detection, Organization matching, legal-authority adjudication, merge, transfer, or dispute-resolution mechanics.

## 6. Requirements handoff

The complete traceable requirements are owned by [Stage 6 — C1 Requirements](./pilot-3-stage-6-c1-requirements.md).

### Experience requirements cover

- truthful start without implied external verification;
- attributable person before success;
- material consequence before effect;
- coherent success meaning;
- durable inspection of Organization + Owner state;
- truthful provenance / truth-class expression;
- incomplete-vs-success distinction;
- known-conflict behavior without adjudicating legal authority.

### Engineering requirements cover

- durable attribution;
- prevention of unattributable success;
- no external-verification prerequisite invented by C1;
- opportunity to present consequence before establishment takes effect;
- product-semantic consistency of Organization + Owner success;
- durable availability of the resulting relationship;
- preservation of truth-class distinction;
- incomplete / failed processing not producing success meaning;
- known unresolved conflict not producing normal uncontested success.

## 7. Experience / design decision space

Design may decide, while preserving the confirmed behavior and requirements:

- number and arrangement of screens / steps;
- exact authentication timing;
- exact form structure and Organization fields;
- exact copy;
- modal, inline, summary, or other consequence treatment;
- checkbox vs no checkbox;
- durable role-inspection surface;
- incomplete / recovery UI if recovery is supported;
- conflict presentation when a known unresolved conflict is present.

These are not invitations to redefine C1 product semantics.

## 8. Engineering decision space

Engineering may decide:

- technical architecture;
- persistence model;
- transaction strategy;
- API / endpoint shape;
- database / schema design;
- service / component boundaries;
- authentication implementation;
- technical representation of attribution;
- how product-semantic consistency is enforced;
- whether and how incomplete progress is stored;
- how a known conflict is technically exposed to C1 once that conflict is known.

Engineering may not reinterpret confirmed product meaning merely to simplify implementation.

## 9. Explicitly open / outside C1

The following are not blockers for C1 handoff and should not be silently solved as part of C1:

- Organization-review outcomes or status taxonomy;
- proof / adjudication of real-world legal representative authority;
- detailed Owner / Staff permission matrix;
- staff invitation, delegation, revoke, transfer, and multi-owner mechanics;
- Organization matching / duplicate-detection policy;
- representation-conflict resolution / merge / transfer flows;
- Campaign preparation / curation;
- donation / funding semantics;
- post-campaign accountability / fund-use verification;
- implementation architecture and contract decomposition.

If implementation reveals that one of these is materially required for C1 to remain truthful, return the question upstream rather than inventing product semantics downstream.

## 10. Material interaction evidence

Stage 4 pressure-tested C1 through:

```text
initiate establishment
→ understand consequence before effect
→ form Organization + initial Owner relationship
→ inspect durable resulting state
→ handle material incomplete / conflict cases
```

Key evidence from that exploration:

- external-authority proof is not required merely to create the internal Kencleng Owner relationship;
- success must be product-semantically coherent across Organization + Owner;
- transient success messaging is insufficient for durable role meaning;
- recovery is optional, but incomplete state cannot masquerade as success;
- a known representation conflict cannot be silently treated as uncontested.

## 11. Engineering handoff acceptance check

Before Stage 7 is marked complete, confirm that a fresh engineering reader can answer from durable artifacts:

| Question | Handoff answer present? |
| --- | --- |
| What is C1 trying to make true for the actor? | YES |
| What product behavior is confirmed? | YES |
| What invariants must not be broken? | YES |
| What Experience Requirements apply? | YES |
| What Engineering Requirements apply? | YES |
| What remains open / out of scope? | YES |
| What may engineering decide? | YES |
| What upstream product distinctions must remain intact? | YES |
| Is original conversation required to begin? | NO |

## 12. Stage 7 exit verdict

Anhar has confirmed that this durable packet is sufficient for engineering to begin without reconstructing the original product discussion.

**Stage 7 verdict: COMPLETE for C1.**

```text
C1
Stage 3B ELIGIBLE
→ Stage 4 Interaction Exploration
→ Stage 5 Confirmed Product Behavior
→ Stage 6 Requirements
→ Stage 7 Durable Handoff
→ PRE-ENGINEERING ROUTE COMPLETE
```

Pre-engineering completion does **not** mean C1 has been implemented or delivered in the real product. Downstream commitments that require C1 to be real remain dependency-blocked until that delivery condition is actually satisfied.

Any later implementation discovery that materially changes C1 product meaning must return upstream through the appropriate Product Authority / pre-engineering decision path rather than silently mutating this handoff.
