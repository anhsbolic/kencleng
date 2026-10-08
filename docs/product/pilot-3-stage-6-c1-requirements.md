# Kencleng Pilot #3 — Stage 6 C1 Requirements

> **Status:** WORKING ARTIFACT — C1 Stage 6 complete  
> **Human status:** Anhar-approved requirements set; Stage 6 audit complete  
> **Authority:** Not canonical Product Authority  
> **Working baseline:** `pilot/3-clean-delivery-baseline`  
> **Confirmed behavior source:** `docs/product/pilot-3-stage-5-c1-confirmed-behavior.md`  
> **Control Tower:** `docs/product/pilot-3-control-tower.md`

## 1. Scope

**C1 — Legitimate Organization representation**

Stage 6 derives requirements from the same confirmed C1 behavior set.

```text
Confirmed Product Behavior
        │
        ├── Experience Requirements
        └── Engineering Requirements
```

The two requirement sets must preserve one meaning. Neither side may invent new Product Truth.

This artifact does not choose technical architecture, API shape, database schema, component structure, or implementation strategy.

## 2. Behavior → requirement traceability

### B1 — Establishment may begin without external-authority proof; attribution exists before success

#### Experience Requirements

**XR1.1** A person must be able to begin Organization establishment without the experience implying that they have already proven real-world legal / representative authority.

**XR1.2** Before establishment can complete successfully, the experience must reach a point where the resulting action and initial Owner relationship can be attributable to a person.

**XR1.3** Exact sign-in timing, authentication flow, and identity UI remain design / implementation freedom provided XR1.1–XR1.2 remain true.

#### Engineering Requirements

**ER1.1** The product must preserve durable attribution between the successful Organization establishment, the establishing person, and the resulting initial Owner relationship.

**ER1.2** A successful establishment state must not be produced when the resulting Owner relationship cannot be attributable to a person.

**ER1.3** C1 must not require a successful external-authority verification outcome merely to begin or complete the internal Kencleng Owner formation defined by C1.

---

### B2 — Material consequence is clear before it becomes effective

#### Experience Requirements

**XR2.1** Before the consequential establishment action becomes effective, the experience must make clear that:
- a Kencleng Organization context will be established;
- the initiating person will become its initial Kencleng Owner;
- Owner is internal Kencleng Organization authority, not independently verified external authority;
- Organization information / evidence involved remains Organization-provided unless separately reviewed.

**XR2.2** The consequence must be presented before, not only after, the successful state transition.

**XR2.3** Exact wording, layout, review step, modal / inline treatment, and acknowledgement mechanism remain design freedom. A forced checkbox is not required by C1.

#### Engineering Requirements

**ER2.1** Successful establishment must not take effect before the product has had the opportunity to present the material consequences required by XR2.1.

**ER2.2** The establishment transition must not itself produce Organization-review, external-authority verification, or Campaign-curation meaning.

---

### B3 — Organization + initial Owner form one successful product consequence

#### Experience Requirements

**XR3.1** The experience must not communicate successful Organization establishment while the initial Owner relationship is not valid.

**XR3.2** The experience must not communicate a successfully formed initial Owner relationship while the corresponding Organization context is not valid.

#### Engineering Requirements

**ER3.1** Observable product state must preserve the invariant:

```text
successful Organization establishment
⇔
valid Organization context
+
valid attributable initial Owner relationship
```

for the C1 success transition.

**ER3.2** Exact transactional / persistence implementation is an engineering decision. The requirement is product-semantic consistency, not a mandated technical transaction mechanism.

---

### B4 — Resulting Organization / Owner relationship remains durably inspectable

#### Experience Requirements

**XR4.1** After successful establishment, the person must be able to identify the Organization context they established and their current Kencleng Owner role.

**XR4.2** The resulting experience must make the Owner role's internal Kencleng meaning inspectable beyond the transient success moment.

**XR4.3** Subsequent C1-related surfaces must not imply that Organization establishment itself proved external legal authority or completed Organization review.

**XR4.4** C1 must not introduce a fabricated Organization-review status taxonomy solely to explain this distinction.

#### Engineering Requirements

**ER4.1** The successful Organization–Owner relationship and its attribution must remain durably available to support later inspection.

**ER4.2** The product must preserve enough semantic distinction to prevent the internal Owner relationship from being interpreted by downstream product behavior as an Organization-review or external-authority verification outcome.

---

### B5 — C1 does not upgrade Organization-provided truth

#### Experience Requirements

**XR5.1** Organization information / evidence shown or collected within C1 must not be presented as independently reviewed merely because establishment succeeds.

**XR5.2** Product language and state treatment must preserve the distinction between Organization-provided information and later review / verification meaning.

#### Engineering Requirements

**ER5.1** Successful establishment must not mutate the semantic status of Organization-provided information into reviewed, independently verified, or Campaign-curated truth.

**ER5.2** Provenance / truth-class meaning required by C1 must remain preservable across the successful establishment transition.

---

### B6 — Incomplete state is not successful establishment

#### Experience Requirements

**XR6.1** If attribution is unavailable or Organization + initial Owner formation has not fully succeeded, the experience must not communicate successful establishment.

**XR6.2** If an incomplete state is exposed, it must be distinguishable from successful establishment.

**XR6.3** Save-draft / resume / retry behavior is not required by C1; if introduced, it must preserve the incomplete-vs-success distinction.

#### Engineering Requirements

**ER6.1** Incomplete / failed C1 processing must not produce a successful Organization establishment meaning.

**ER6.2** Recovery persistence is optional for C1. If implemented, recovered state must not be indistinguishable from a valid completed Organization + Owner relationship before success actually occurs.

---

### B7 — Known unresolved representation conflict cannot silently complete as uncontested

#### Experience Requirements

**XR7.1** When Kencleng has a known unresolved representation conflict relevant to the Organization being established, the experience must not silently present establishment as an uncontested success.

**XR7.2** C1 does not require the experience to determine which claimant has valid real-world legal authority.

**XR7.3** Exact conflict presentation and resolution path remain outside C1.

#### Engineering Requirements

**ER7.1** Where a known unresolved representation conflict is present, C1 must not produce the normal uncontested-success meaning or grant a new initial Owner relationship as though representation were uncontested.

**ER7.2** C1 does not define conflict detection, Organization matching, merge, transfer, dispute resolution, or legal-authority adjudication mechanisms.

## 3. Cross-cutting acceptance requirements

C1 is acceptable only if all of the following can remain true together:

1. a person can establish an internal Kencleng Organization + initial Owner relationship without C1 pretending to verify external authority;
2. attribution exists before success;
3. material consequence is available before the action becomes effective;
4. Organization + Owner success meaning is coherent as one product consequence;
5. resulting Organization + Owner meaning survives beyond a transient success moment;
6. Organization-provided truth is not upgraded by establishment;
7. incomplete state cannot masquerade as success;
8. a known unresolved representation conflict cannot silently become uncontested success;
9. C1 remains separate from Organization review, Campaign curation, and fund-use verification semantics.

## 4. Open experience / design freedom

Stage 6 intentionally does not prescribe:

- exact number of screens / steps;
- exact authentication timing;
- exact form structure or Organization field set;
- exact copy;
- checkbox vs non-checkbox acknowledgement;
- modal / inline / summary treatment;
- exact durable role-inspection surface;
- recovery UI;
- conflict UI / resolution journey.

These may be decided during design as long as confirmed behavior and requirements remain true.

## 5. Engineering decision space

Engineering may decide, subject to the requirements above:

- persistence model;
- transaction strategy;
- API / endpoint shape;
- component / service boundaries;
- identity / authentication implementation;
- how durable attribution is technically represented;
- how semantic consistency is technically enforced;
- how incomplete processing is stored, if stored;
- how a known unresolved conflict is technically surfaced to C1 once such a conflict is known.

Engineering may **not** reinterpret the confirmed product meanings to simplify implementation.

## 6. Still outside C1

- Organization-review outcomes and status taxonomy;
- detailed Owner / Staff permissions;
- staff invitation / delegation / revoke / transfer / multi-owner rules;
- Organization matching / duplicate detection policy;
- conflict-resolution / merge / transfer flows;
- Campaign preparation / curation;
- API / schema / architecture specifications.

## 7. Stage 6 audit

| Audit question | Result |
| --- | --- |
| Requirements trace back to confirmed Stage 5 behavior? | PASS |
| Any new Product Truth invented by XR / ER? | NONE FOUND |
| Unnecessary solution-lock? | NONE REQUIRED; UI, auth timing, persistence, transaction, API, and component choices remain open |
| Missing observable requirement that blocks C1 handoff? | NONE FOUND |
| Engineering requirement accidentally prescribes architecture? | NO |
| Experience and engineering requirements disagree semantically? | NO |

**Stage 6 verdict: COMPLETE for C1.**

C1 is ready for Stage 7 durable engineering handoff preparation. Stage 7 should package, not reinterpret, the confirmed behavior and requirements.
