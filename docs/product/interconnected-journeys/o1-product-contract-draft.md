# O1 — Organization Establishment with Legitimate Initial Internal Authority

> **Status:** PRODUCT ELABORATION + PRODUCT CONTRACT **DRAFT v0.1** / **NOT APPROVED AS EFFECTIVE I1**.
> **Owning Product decision:** Anhar; pending fresh scoped decision. **Date:** 2026-10-22. **Repository/branch:** `anhsbolic/kencleng` / `pilot/station5-c1-space-only` (research only, not operational `main`).
> **Research basis (Product-side only):** [current approved Product Intent](../product-intent.md), [approved Level 2 v0.3](level-2-network.md) (O1 + relevant typed edges), and [approved working Product Output Contract v0.1 / selection record](selection-and-handoff-research.md). **No historical commitment-specific Product behavior/requirements were used to derive this draft.**
> **Critical distinction:** The statements labeled **PROPOSED** below are new O1-scoped Product decisions for Anhar; they are **not** binding simply because this file exists. No historical Product decision is inherited. The eventual self-contained, exact-revision approved Product Contract would be the sole effective Product-semantic input for Rancang I1; this draft is **not** that input.
> **Consumer boundary:** No Rancang invocation or delivery authorization. Rancang's black-box definition update is separate and unapproved; technical I2/I3/Human I4 inputs remain necessary.

## 1. One meaningful Product outcome

**Selected research outcome:** O1 — a person can establish an attributable Organization context in Kencleng, resulting in a legitimate initial **internal Organization authority** relationship whose significance is truthful, durable, and inspectable.

**Why this is useful (Product research rationale, not a new requirement):** Kencleng's Campaigns belong to Organizations; someone acting on behalf of an Organization needs intelligible, attributable internal authority. This creates a meaningful Organization-side consequence without requiring an entire Organization review, Campaign flow, or public/donation journey.

**What Product Intent already supports:**
- Organizations, their representatives, Owners and Staff have distinct roles/meaning; Campaigns and Organizations are separate.
- Organization-supplied assertions, platform-known facts and independent review/verification must not be confused.
- Consequential actions and meaningful states require understandable, provenance-aware, truthful communication.
- The Product Intent leaves exact permissions, approval/review mechanics, and implementation decisions open.

**What Level 2 v0.3 provides:** O1 is a *research transition* about Organization establishment and legitimate internal authority; exact grant, coupled success, failure and disputed-claim rules were intentionally left OPEN. O1 is not universal first work for existing valid Organization contexts.

## 2. Current Product decisions needing approval

The following is **one coherent recommended O1 formation policy proposal** for Human review, not five independently approved facts.

| ID | Decision proposed for this O1 scope | Source status / why a new decision is needed |
| --- | --- | --- |
| **PD-O1-01** | A person may initiate establishing a new Kencleng Organization context. **At the effective successful establishment**, the acting, legitimately attributable person is granted that Organization's **initial internal Kencleng Owner** relationship. This grant describes *Kencleng-internal* authority only and does not claim independently verified external representation. | **PROPOSED.** Product Intent distinguishes Owner/Staff but does not decide the initial grant. No pre-established external legal verification is asserted as a mandatory success condition for this limited internal relationship. |
| **PD-O1-02** | A successful establishment is **one coherent consequential result**: valid Organization context **and** valid initial attributable Owner relationship. The product must not treat only one as full establishment success. Person attribution is required **by the time the consequence becomes effective**, not necessarily at entry. | **PROPOSED.** Precise coupled-success and attribution timing are not fixed by Product Intent/Level 2; exact authentication/transaction design is delegated. |
| **PD-O1-03** | Before the consequence becomes effective, the actor must be able to understand which Organization context is being formed, that they will be its initial internal Owner and the limits of that meaning. After success, the current Organization/Owner relationship and its internal meaning must be **durably inspectable**, not merely a transient success message. | **PROPOSED scoped elaboration** of the Product Intent's clarity/provenance/truth safeguards. No UI flow, acknowledgement control or screen count is selected. |
| **PD-O1-04** | Incomplete/unattributable/failed establishment cannot be represented as successful. When an **actually known, relevant unresolved representation dispute** exists, the product cannot silently assert uncontested successful initial authority; the affected claim must remain unresolved or be denied, with truthful user-visible meaning. This does **not** mandate universal duplicate detection, external adjudication, conflict-resolution UX, or a new technical dispute service. | **PROPOSED.** Product Intent requires honest authority/truth but does not settle O1 dispute handling. Condition depends on a real known signal; availability is unverified. |
| **PD-O1-05** | Organization-supplied information remains attributable **Organization-provided** information. Establishment by itself does not perform or imply Organization review, external legal verification, Campaign curation, public visibility, or donation eligibility. | **Product Intent-supported distinction applied to O1**, with the exact O1 consequence proposed above. |

**Recommendation:** Approve or revise this **one bounded formation policy** as a cohesive Product decision. Do **not** authorize Engineering work by approving it. If Anhar rejects part of it, rewrite the affected contract and acceptance before any Rancang handoff.

## 3. Draft person-facing experience (illustrative, not UI flow)

**Entry:** A person wants to establish an Organization presence in Kencleng. They can understand that establishing an internal Organization context is different from external legal/representative verification. They may begin without the product claiming verified external authority.

**Before effect:** The person must be able to understand the internal Organization context that would be formed and the resulting initial Owner authority and its limits. Sufficient person attribution must exist *before the successful consequential effect*; the sign-in/identity interaction mechanism is not fixed.

**Success:** Kencleng truthfully reports a valid Organization context together with the initial attributable internal Owner relationship, and makes the resulting relation understandable and available for later inspection. No Organization review/trust label is automatically implied.

**Failure or uncertainty:** If attribution or coherent formation does not occur, success is not claimed. If the system knows an unresolved relevant representation dispute, no unqualified uncontested authority is claimed. An error or unresolved outcome is not turned into verified/legal authority or an ordinary success.

## 4. Proposed binding requirements once Product accepts PD-O1-01..05

The requirements in this section are **DRAFT/PROPOSED**, and **must not** be consumed as approved Product Authority until a separate exact-revision Human approval.

| ID | Draft Product requirement |
| --- | --- |
| **O1-R01** | Initiation is possible for a person without an assertion of externally verified legal authority; any mandatory actor eligibility conditions beyond this require an explicit scoped Product decision. |
| **O1-R02** | At success, the initiating person is legitimately attributable to the resulting Organization and initial internal Owner relationship, as proposed in PD-O1-01. |
| **O1-R03** | The material consequence and internal limits of initial Owner authority are intelligible before establishment takes effect. |
| **O1-R04** | The product only claims successful establishment when actual valid Organization context and legitimate, attributable initial Owner relationship both hold. |
| **O1-R05** | The resulting Organization and current initial Owner relationship remain durably inspectable with their actual internal authority meaning. |
| **O1-R06** | Establishment does not upgrade Organization-provided assertions into review, external proof, Campaign curation, or other independent truth classes. |
| **O1-R07** | Absent attribution, partial, failed or interrupted establishment is distinguished from genuine success; recovery mechanisms remain a Design/Engineering choice unless a subsequent Product decision makes recovery consequential. |
| **O1-R08** | An actually known relevant unresolved representation dispute prevents false uncontested authority/success; no universal detection/adjudication capability is inferred. |

**Success invariant proposed:** `O1 successful establishment` means `valid Organization context + valid attributable initial internal Owner relationship`, with truthful communication and later inspectability. It **does not** mean legal/external verification or completed Organization review.

**Negative-path invariants proposed:** no unattributable authority grant, no successful claim for partially formed consequence, no uncontested authority claimed across an actually known unresolved relevant dispute, no implied reviewed/verified status.

## 5. Draft observable acceptance

| ID | Testable Product consequence / negative case | Standing |
| --- | --- | --- |
| **O1-AC01** | A person can initiate without the product presenting external representative authority as already proved. | PROPOSED |
| **O1-AC02** | Establishment cannot reach successful effect without real attribution of both action and initial Owner relation to the initiating person. | PROPOSED |
| **O1-AC03** | Before effect, the actor can understand what Organization is established and what the internal Owner relationship means and does not prove. | PROPOSED |
| **O1-AC04** | Valid Organization + initial attributable Owner can be experienced and inspected as **one** successful result; a partially formed pair cannot be ordinary success. | PROPOSED |
| **O1-AC05** | After success, a later inspection exposes the current Organization/Owner relation and relevant limits without relying on the success notification. | PROPOSED |
| **O1-AC06** | An Organization-provided claim remains source-attributed and is not upgraded to independent review, external verification or Campaign curation solely by establishment. | PROPOSED |
| **O1-AC07** | Unattributed, failed, incomplete and uncertain outcomes cannot look like ordinary successful establishment; optional retry/recovery cannot promote partial state into success. | PROPOSED |
| **O1-AC08** | If a relevant unresolved representation conflict is already known, the product does not issue ordinary uncontested initial Owner authority or pretend that claim is resolved. | PROPOSED |
| **O1-AC09** | No O1 completion produces Campaign preparation, public visibility, donation eligibility or Organization review as a side effect. | PROPOSED |

These are product-observable criteria, **not** API tests, database transaction prescriptions, executable proofs, or evidence that current code works.

## 6. Boundaries and delegated decisions

**In-scope Product meaning:** attributable creation of a Kencleng Organization context, initial internal Owner authority, consequential disclosure, coherent result, inspectable truth and material negative outcomes.

**Explicitly excluded:** full Owner/Staff permission matrix, staff invitations/transfers/delegation, independent legal-authority verification, full Organization review and its statuses, Campaign creation/review/publication/eligibility, funding/donation/accountability, universal duplicate detection or representation-dispute adjudication, and a mandated recovery subsystem.

**Delegated Design/Engineering freedom (after valid assignment):** exact form/fields consistent with meaningful identification; sign-in timing, identity UX and authentication technical implementation; disclosure method; interaction layout/copy; persistence/data model; transactional consistency mechanism; API/interface/component boundaries; failure/unknown-state presentation; representation of an actually known dispute signal; verification techniques. These freedoms cannot relax approved semantic outcomes, authorize broader roles, or override project security constraints.

**Required technical/Design/environment evidence is external to Product-semantic I1:** available identity authority and person attribution, data persistence, authorized code/architecture/security constraints, known-conflict signals if relevant, legitimate repository/ref/environment and consumer expectations. No capability is asserted to be delivered. Rancang I2/I3 and Human-direct I4 remain independent requirements.

## 7. Contract-readiness result and Human decision

**Product Output Contract v0.1 logical structure:** Actor/outcome, proposed rules, positive and negative paths, invariants, acceptance, dependencies, exclusions, solution freedom and boundary are **PRESENT AS DRAFT**.

**Cannot claim APPROVED I1:** PD-O1-01..05 require fresh Product owner approval; actual technical readiness and external signals have not been verified. This draft is **not Rancang-consumable binding Product authority**. Anhar may approve a revised exact PRD only after Product reconciliation/quality check.

**Next Product action:** Review the one proposed O1 formation policy; resolve or revise material decisions; then issue and approve the exact standalone Product Contract. Do not pull in historical pilot requirements to make the document appear more complete.
