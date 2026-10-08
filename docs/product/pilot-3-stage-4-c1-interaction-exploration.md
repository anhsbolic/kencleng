# Kencleng Pilot #3 — Stage 4 C1 Interaction Exploration

> **Status:** WORKING ARTIFACT — C1 Stage 4 complete  
> **Human status:** Anhar-approved through SX material exception behavior; Stage 4 exit audit complete  
> **Authority:** Not canonical Product Authority  
> **Working baseline:** `pilot/3-clean-delivery-baseline`  
> **Commitment source:** `docs/product/pilot-3-stage-3b-commitment-sequencing.md`  
> **Control Tower:** `docs/product/pilot-3-control-tower.md`

## 1. Commitment in scope

**C1 — Legitimate Organization representation**

Working product claim:

> A person can establish an Organization context in Kencleng and become its initial Kencleng Organization Owner, with clear internal authority meaning and provenance, without Kencleng implying that the person has independently verified real-world legal authority.

C1 is the only active delivery commitment for this Stage 4 exploration unless materially new evidence proves another commitment is required for C1 to remain truthful.

## 2. Product semantics already settled for C1

The following are inputs to interaction exploration, not questions to reopen without materially new evidence:

- the first person who establishes an Organization in Kencleng becomes the initial Kencleng Organization Owner by product rule;
- Kencleng Owner is an internal authority-bearing Organization role within the Kencleng Organization context;
- Kencleng Owner is not proof of legal ownership, chairmanship, authorized-signatory status, independently verified external authority, or Organization-review outcome;
- a separate evidence-review is not required merely to grant the initial Kencleng Owner role;
- Organization information/evidence submitted by Owner/Staff remains Organization-provided information/evidence until a separate bounded review gives additional meaning;
- Organization establishment / internal Kencleng ownership, Organization review, and Campaign curation remain distinct;
- detailed permission, invitation, delegation, revoke, transfer, and multi-owner rules remain parked unless C1 interaction evidence makes one materially necessary.

## 3. Stage 4 question

> What must a person know, see, do, and understand for C1 to be experienced truthfully — and what state must Kencleng preserve — from initiation through a successfully established Organization + initial Owner relationship?

Do not begin from a predetermined form or screen.

For each scene, explore:

```text
what the person knows
what the person sees
what the person can do
what Kencleng knows
what product state changes
what must remain distinguishable
material failure / recovery
```

## 4. Initial scene map

This map is exploratory. Scene boundaries may change as interaction evidence improves.

```text
S1 — Initiate Organization establishment
        ↓
S2 — Understand consequential meaning before establishment completes
        ↓
S3 — Establish Organization + initial Kencleng Owner relationship
        ↓
S4 — Understand resulting Organization / Owner state
        ↓
SX — Material exception / recovery paths
```

### S1 — Initiate Organization establishment

**WORKING DECISION — Anhar-approved**

> A person may begin Organization establishment without Kencleng yet asserting any real-world representative authority. Before establishment becomes successful and the initial Owner relationship is formed, Kencleng must be able to durably attribute that action and resulting Owner relationship to a person.

Therefore:

- the person may begin with intent to establish an Organization context in Kencleng without first proving real-world legal / representative authority;
- before successful establishment, Kencleng must have sufficient person attribution for the resulting initial Owner relationship to be attributable;
- exact authentication, account, sign-in timing, and identity mechanism remain OPEN EXPERIENCE / DESIGN / later engineering matters as appropriate;
- exact Organization data fields remain outside this decision;
- starting establishment does not itself create an Organization, Owner relationship, Organization-review outcome, or external-authority claim.

### S2 — Consequence before commitment

**WORKING DECISION — Anhar-approved**

> Before Organization establishment becomes effective, Kencleng must clearly present the material consequence that an Organization context will be created and the initiating person will become its initial Kencleng Owner. The interaction must distinguish that internal Kencleng authority from independently verified real-world authority. Any Organization information / evidence involved remains Organization-provided unless separately reviewed.

Minimum consequence that must be made clear before the action becomes effective:

- an Organization context will be established in Kencleng;
- the initiating person will become the initial Kencleng Organization Owner;
- Owner is an authority-bearing role within Kencleng, not proof of legal ownership, authorized external representation, or independently verified authority;
- Organization information / evidence involved remains Organization-provided and is not automatically reviewed or independently verified.

Kencleng is responsible for making these consequences clear; the product cannot guarantee that a person has perfectly understood them.

Exact wording, layout, disclosure pattern, acknowledgement mechanism, and interaction treatment remain OPEN EXPERIENCE / DESIGN DECISIONS. A mandatory checkbox / consent control is not required by this product decision.

### S3 — Establishment + relationship formation

**WORKING DECISION — Anhar-approved**

> A successful Organization establishment is one product consequence in which both the Kencleng Organization context and its attributable initial Kencleng Owner relationship become valid. Kencleng must not represent establishment as successful if either side of that relationship has not successfully formed. Incomplete progress may be recoverable, but must remain distinguishable from successful establishment.

Product-semantic success therefore means:

```text
Organization context is valid
+
initial Kencleng Owner relationship is valid
+
the relationship is attributable to the establishing person
+
the role meaning remains internal to Kencleng
```

Kencleng must not expose a product meaning equivalent to:

```text
"Organization successfully established"
while
"initial Owner relationship is not yet valid"
```

or the reverse.

Incomplete input / progress may be preserved for recovery, but it must not be presented as a successfully established Organization or successfully formed Owner relationship.

No Organization-review outcome, Campaign-curation outcome, or external-authority verification is implied by this transition.

This is **product-semantic atomicity**, not a technical architecture decision. Exact transactional / persistence implementation remains an engineering concern.

### S4 — Resulting state comprehension

**WORKING DECISION — Anhar-approved**

> After successful establishment, Kencleng must expose a durable and inspectable resulting state that identifies the established Organization context and the person's current Kencleng Owner relationship. The product must preserve the distinction between that internal role and any Organization-review or independently verified external-authority meaning.

At minimum, after successful establishment the person must be able to determine:

- which Organization context now exists in Kencleng;
- that they currently hold the initial Kencleng Owner role;
- that the role is an authority-bearing role inside the Kencleng Organization context;
- that establishment does not itself imply Organization review or independently verified external authority.

This meaning must not exist only in a transient success message. The resulting Organization / Owner relationship must remain durably inspectable later.

The internal-vs-external distinction does not need to be repeated as a warning on every surface, but subsequent surfaces must not create a misleading stronger meaning.

C1 does **not** define an Organization-review status taxonomy such as `unverified`, `verification pending`, or `not reviewed`. Those semantics belong to the separate Organization-review commitment if and when established.

Exact surface, wording, placement, role-details treatment, and navigation remain OPEN EXPERIENCE / DESIGN DECISIONS.

### SX — Material exception / recovery

**WORKING DECISION — Anhar-approved**

Material exception handling for C1:

1. **Attribution cannot be established**
   - Organization establishment must not be represented as successful;
   - the initial Owner relationship must not be represented as valid.

2. **Establishment is interrupted before Organization + Owner become valid**
   - incomplete progress may be recoverable;
   - incomplete progress must remain distinguishable from successful establishment.

3. **Organization-provided information is incomplete or uncertain**
   - incomplete / uncertain information may remain a valid product state where the minimum establishment behavior can still be truthfully satisfied;
   - C1 does not invent stronger claims or silently convert incomplete Organization-provided information into reviewed fact;
   - exact required Organization fields remain outside C1.

4. **An existing Kencleng Organization context is known or materially suspected to represent the same real-world Organization**

   > Kencleng does not guarantee real-world Organization uniqueness at C1. However, when an existing Kencleng Organization context is known or materially suspected to represent the same Organization, Kencleng must not silently complete a second establishment as though no conflict exists.

   Therefore:
   - the conflict remains unresolved rather than silently ignored;
   - a new initial Owner must not be granted as though representation were uncontested;
   - exact matching / detection, claim-existing-Organization flow, dispute resolution, merge, transfer, or Platform Operator intervention are outside C1 unless later evidence makes them necessary.

This conflict behavior is not Organization review and does not determine which person has legally valid external authority.

## 5. Open classifications at Stage 4 entry

### OPEN PRODUCT DECISION

No unresolved product decision currently blocks C1 from leaving Stage 4.

The resolution mechanics for an existing-Organization conflict remain intentionally outside C1; C1 only establishes that a known / materially suspected conflict cannot be silently treated as uncontested establishment.

### OPEN EXPERIENCE / DESIGN DECISION

- how consequential Owner meaning is explained before establishment completes;
- how internal Kencleng authority is distinguished from external real-world authority;
- how resulting Organization / Owner state is communicated;
- how uncertainty or incomplete Organization-provided information is expressed.

### OPEN ENGINEERING CONSTRAINT

None required to begin interaction exploration.

### PARKED — NOT NEEDED YET

- exact account / authentication mechanism;
- detailed Organization field schema;
- detailed Owner / Staff permission matrix;
- staff invitation / delegation / revoke / transfer / multi-owner mechanics;
- Organization-review mechanics and outcomes;
- Campaign preparation / curation behavior.

## 6. Stage 4 exit audit

Stage 4 is sufficiently concrete to feed Stage 5.

| Exit question | Result |
| --- | --- |
| Can the C1 interaction be experienced end-to-end from initiation through resulting state? | PASS |
| Are consequential meanings exposed before the action becomes effective? | PASS |
| Is successful establishment behavior distinguishable from incomplete progress? | PASS |
| Is the Organization + initial Owner relationship product-semantically coherent? | PASS |
| Are internal Kencleng authority and external / reviewed authority kept distinct? | PASS |
| Are material failure / recovery cases sufficiently bounded? | PASS |
| Are design freedoms distinguishable from product semantics? | PASS |
| Is an unresolved product decision still required before behavior can be confirmed? | NO |

**Stage 4 verdict: COMPLETE for C1.**

Stage 4 has produced interaction evidence. Stage 5 should now decide which observed behaviors become confirmed product behavior / invariants for C1. No technical architecture or implementation choice has been made.
