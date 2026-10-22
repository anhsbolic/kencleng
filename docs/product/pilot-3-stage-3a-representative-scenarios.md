# Kencleng Pilot #3 — Stage 3A Representative Scenarios

> **CURRENT EXPERIMENT ROUTING NOTICE (2026-10-22; pilot branch only):** This file preserves its dated Pilot #3 reasoning, scoped approvals and historical evidence **without rewriting history**. It is **not current input authority** for selecting O1 or generating O1 Product Contract/PRD. Do not carry forward historical commitment identifiers, business rules, acceptance or stage sequence into the current Interconnected Journeys research by inheritance. Use [current Product routing](README.md), [Level 2](interconnected-journeys/level-2-network.md), and [current selection/handoff decisions](interconnected-journeys/selection-and-handoff-research.md). Historical approvals retain only their original scope; operational `main` is unchanged.

> **Status:** WORKING ARTIFACT — Pilot #3 Stage 3A complete  
> **Human status:** Anhar-approved working state through final cross-scenario audit  
> **Authority:** Not canonical Product Authority  
> **Working baseline:** `pilot/3-clean-delivery-baseline`  
> **Related method:** `docs/product/pilot-3-stage-3-working-model.md`  
> **Purpose:** Preserve the current minimal representative scenario set and coverage audit before deeper scenario elaboration.

## 1. Representative scenario set v0

### R1 — Organization legitimacy + review

**Core question:** How does an Organization become legitimately represented in Kencleng, and what bounded meaning does Organization review create?

**Primary actors:** Owner, Staff, Reviewer.

**Pressure-tests:**
- Owner vs Staff authority;
- legitimate Organization representation;
- Organization-provided information and provenance;
- Organization review;
- positive meaning of review;
- reviewer independence / conflict of interest.

Boundary:

```text
Organization review
≠ Campaign curation
```

### R2 — Campaign proposition → curation → actual public/donation state

**Core question:** How does a Campaign move from Organization-provided proposition through bounded curation into actual visibility and donation-eligibility states?

**Primary actors:** Owner / Staff, Reviewer, Public Visitor as downstream dependency.

**Pressure-tests:**
- Campaign preparation;
- curation submission;
- requested changes;
- bounded review history;
- revision / re-review;
- rejection semantics;
- positive meaning of curation;
- relationship among curation, public visibility, and donation eligibility.

Boundary:

```text
Campaign curated
≠ automatically public
≠ automatically donation-eligible
```

### R3 — Public understanding → considered decision → Donatur relationship

**Core question:** How does a Public Visitor understand a Campaign well enough to decide, and if they contribute, how does a legitimate donation relationship form?

**Primary actors:** Public Visitor, Donatur, Organization as information source.

```text
understand Campaign
├─ no donation
│  → valid considered outcome
└─ contribution
   → legitimate donation fact?
   → Donatur relationship
```

**Pressure-tests:**
- Confidence Before Conversion;
- guest vs registered relationship;
- donation eligibility;
- donation fact;
- donation fact ≠ automatic funding consequence;
- money meaning;
- pending / failure states.

### R4 — Fundraising phase change → reporting → continuing accountability

**Core question:** After fundraising materially changes or ends, how does the Organization remain accountable and how can Donatur continue understanding what actually happened?

**Primary actors:** Owner / Staff, Donatur, Reviewer when applicable.

**Pressure-tests:**
- Campaign closure semantics;
- Organization reporting obligation;
- incomplete / pending / unavailable information;
- operational progress;
- reported outcome;
- fund-usage verification;
- accountability complete / pending / incomplete / unmet;
- public vs Donatur-specific accountability visibility.

Boundaries:

```text
fundraising closure
≠ accountability closure

reported outcome
≠ verified outcome
≠ real-world impact
```

Fund-usage verification remains a branch of R4 unless later evidence shows it requires a materially distinct representative scenario.

### X1 — Privileged correction without semantic override

X1 is a cross-cutting exception probe, not currently a core representative scenario.

**Core question:** Can Platform Operator intervention occur without the operator acquiring unsupported Organization, Reviewer, funding, verification, settlement, or impact authority?

```text
operational privilege
≠ semantic authority
```

## 2. Whole-product coverage audit

| Product area | Covered by |
| --- | --- |
| Organization establishment / representation | R1 |
| Organization review | R1 |
| Owner / Staff authority | R1, R2, R4 |
| Campaign preparation | R2 |
| Campaign curation | R2 |
| Review history / revision | R2 |
| Public visibility | R2 |
| Donation eligibility | R2 → R3 |
| Public understanding | R3 |
| Considered decision | R3 |
| No-donation outcome | R3 |
| Guest / registered distinction | R3 |
| Donation fact | R3 |
| Funding consequence boundary | R3 |
| Campaign closure | R4 |
| Organization reporting | R4 |
| Continuing accountability | R4 |
| Pending / incomplete information | R4 |
| Fund-usage verification | R4 |
| Donatur continued understanding | R4 |
| Platform Operator exception | X1 |

**Coverage verdict:** no obvious whole-product hole is currently visible.

## 3. Actor coverage audit

| Actor | Representative coverage |
| --- | --- |
| Public Visitor | R3 |
| Donatur | R3, R4 |
| Organization Owner | R1, R2, R4 |
| Organization Staff | R1, R2, R4 |
| Reviewer | R1, R2, R4 |
| Platform Operator | X1 when materially required |

**Actor verdict:** all materially relevant current actors are covered without forcing Platform Operator into the core value journey.

## 4. Overlap / handoff audit

The overlap among R1–R4 is intentional handoff rather than duplicate workflow coverage.

```text
R1 Organization can participate legitimately
        ↓
R2 Campaign can reach an actual bounded product state
        ↓
R3 Visitor can legitimately understand / decide / donate
        ↓
R4 Accountability continues after fundraising / donation
```

Handoff boundaries:
- R1 → R2: Organization legitimacy does not automatically create Campaign legitimacy.
- R2 → R3: actual Campaign visibility / eligibility is a dependency for truthful donation behavior.
- R3 → R4: donation relationship creates continuing accountability dependency.

## 5. Branches not promoted to standalone scenarios

For now, these remain branches/failure cases:
- Campaign rejected;
- multiple review revisions;
- pending / failed donation;
- visitor chooses not to donate;
- incomplete reporting;
- verification unavailable;
- Staff lacks authority.

Promote one only if later scenario work shows materially distinct product behavior or actor responsibility.

## 6. Checkpoint verdict

**Candidate generation + coverage audit: PASS.**

Current minimal representative scenario set:

```text
R1 — Organization legitimacy + review
R2 — Campaign proposition + curation + actual visibility / eligibility
R3 — Public understanding + considered decision + donation / Donatur relationship
R4 — Closure + reporting + continuing accountability / verification
X1 — Platform Operator cross-cutting exception probe
```

This is **not yet Stage 3A complete** and is **not a delivery sequence**.

Next Stage 3A work is to define for each R1–R4:
- product hypothesis / purpose;
- starting condition;
- end condition;
- major semantic transitions;
- material actor dependencies;
- OPEN PRODUCT DECISION exposed.

After that, Stage 3A feeds candidate commitments into Stage 3B for dependency closure, eligibility, partial ordering, and sequencing.


## 7. R1–R4 scenario challenge checkpoint

The representative scenarios were challenged at the same semantic level before any interaction-detail work. The goal was to expose hidden assumptions, refine scenario boundaries, and identify product decisions each scenario must force.

### R1 — Organization legitimacy + review

**Purpose**

Test how an Organization becomes legitimately represented and how Organization review creates bounded meaning without becoming a universal trust claim.

**Starting condition**

The Organization does not yet have sufficient legitimate Kencleng representation / review context.

**End condition**

The Organization has an actual representation context and an understandable Organization-review state with bounded meaning.

**Major semantic transitions**

```text
no established representation
→ legitimate Organization representation exists

Organization assertions unavailable / insufficient
→ Organization-provided context exists with provenance

no review judgment
→ bounded Organization-review state / outcome exists

review outcome exists
→ downstream consequence remains limited to its actual meaning
```

**Material actor dependencies**

Organization Representative supplies Organization assertions/evidence; Reviewer creates bounded judgment; Kencleng preserves authority, provenance, state, and meaning.

**Important refinements**

- Organization establishment / representation ≠ Organization review.
- Owner does not need to be assumed as a mandatory first workflow step before Staff.
- Reviewed Organization ≠ universally trusted Organization.
- Organization review completion is not assumed to be a prerequisite for all Campaign preparation.

**OPEN PRODUCT DECISION exposed**

- positive meaning and legitimate outcomes of Organization review;
- minimum Organization context required for meaningful review;
- consequence of pending / incomplete / negative review;
- relationship between Organization review and downstream Campaign activity;
- when material Organization change requires review reconsideration;
- Owner vs Staff Organization-level authority that is materially consequential.

**Verdict:** R1 PASS as a representative scenario. It is not an "Organization onboarding flow."

---

### R2 — Campaign proposition + curation + actual public/donation state

**Purpose**

Test how a Campaign obtains bounded curation meaning and actual visibility / donation-eligibility state without turning any one of those into a proxy for the others.

**Starting condition**

The Organization has sufficient legitimate representative context to prepare a Campaign. Organization-review completion is not assumed.

**End condition**

The Campaign has actual, distinguishable curation, visibility, and donation-eligibility states whose relationships are understandable.

**Major semantic transitions**

```text
no Campaign proposition
→ Organization-provided Campaign proposition exists

proposition exists
→ Campaign enters bounded curation context

no Reviewer judgment
→ bounded curation judgment / feedback exists

feedback requires response
→ relevant Campaign material changes

changed Campaign
→ Reviewer can understand relevant change against prior judgment

curation cycle reaches an outcome
→ bounded curation meaning exists

Campaign state changes
→ actual visibility state exists

Campaign state changes
→ actual donation-eligibility state exists
```

Visibility and donation eligibility are deliberately not modeled as automatic sequential consequences of curation.

**Material actor dependencies**

Organization Representative provides proposition/evidence/revisions; Reviewer provides bounded curation judgment; Kencleng preserves relevant review history and actual state; Public Visitor depends downstream on truthful visibility/eligibility.

**Important refinements**

- curated ≠ automatically public;
- public ≠ automatically donation-eligible;
- Campaign curation ≠ verification of all Campaign claims;
- Reviewer requires bounded relevant history, not only the latest snapshot;
- every edit is not assumed to require re-review;
- rejection is not assumed to be terminal;
- Organization review outcome is not assumed to automatically govern Campaign preparation, curation, visibility, or eligibility.

**OPEN PRODUCT DECISION exposed**

- positive meaning and outcomes of Campaign curation;
- distinction among changes requested, negative outcome, pending / insufficient;
- re-submission / re-review semantics;
- material change and superseded-outcome semantics;
- relationship between Organization review and Campaign curation;
- relationship among curation, public visibility, and donation eligibility;
- whether visibility and eligibility can change independently;
- materially consequential Owner vs Staff Campaign authority.

**Verdict:** R2 PASS. It is currently the most decision-heavy scenario and may later yield multiple delivery commitments.

---

### R3 — Public understanding + considered decision + Donatur relationship

**Purpose**

Test whether a person can make a considered, non-coercive decision from ordered product truth and become Donatur only through a legitimate donation outcome.

**Starting condition**

The Campaign is in an actual public context. Contribution is available only when actual donation-eligibility state permits it.

**End condition**

The visitor has made a considered decision. If contribution produces a legitimate donation fact, a Donatur relationship exists; pending or failed outcomes remain truthful without unsupported funding consequence.

**Major semantic transitions**

```text
public Campaign context exists
→ Visitor can inspect relevant Campaign truth

insufficient understanding
→ sufficient understanding for considered decision

considered decision
├─ no contribution
│  → valid outcome
└─ intention to contribute
   → contribution context begins

contribution context
→ actual donation outcome becomes knowable

legitimate donation fact exists
→ person becomes Donatur for this Campaign

unresolved / failed donation
→ no unsupported Donatur or funding consequence

legitimate donation fact
→ funding consequence may follow only when product truth supports it
```

**Material actor dependencies**

Organization provides Campaign assertions; applicable Reviewer outcomes contribute bounded judgment; Kencleng preserves provenance, actual state, money meaning, and uncertainty; Public Visitor decides; legitimate donation relationship creates the Donatur role.

**Important refinements**

- public ≠ donation-eligible;
- considered no-donation is a valid product outcome;
- sufficient understanding ≠ exhaustive disclosure;
- intent to contribute ≠ contribution attempt ≠ donation fact ≠ funding consequence;
- guest / registered is an account relationship, not the Donatur business role;
- pending / failed donation is a product state, not merely an error;
- continuing accountability belongs primarily to R4 rather than being absorbed into R3.

**OPEN PRODUCT DECISION exposed**

- minimum understanding before contribution;
- consequential information that must be understood before commitment;
- exact meaning of donation eligibility;
- product-level transition from contribution intent to legitimate donation fact;
- legitimate pending / failed contribution states;
- when donation fact creates funding consequence;
- money-state meaning;
- guest vs registered Donatur capability difference, if any;
- consequence of materially changed Campaign state during contribution;
- minimum product relationship created when someone becomes Donatur.

**Verdict:** R3 PASS as one representative scenario. Considered decision and donation remain coupled because "confidence before conversion" is the coherent product question, even if Stage 3B later yields multiple commitments.

---

### R4 — Continuing accountability after donation / fundraising change

**Purpose**

Test whether post-donation accountability remains truthful over time, with Organization reporting, bounded independent verification when applicable, chronology, provenance, uncertainty, and reported ≠ verified ≠ impact distinctions preserved.

**Starting condition**

One or more legitimate Donatur relationships already exist. Fundraising materially progresses, changes, or ends while Organization accountability continues.

**End condition**

Applicable users can understand what is reported, what remains pending / unavailable, what has been independently judged when applicable, and the actual accountability state without implying unsupported impact.

**Major semantic transitions**

```text
donation relationship exists
→ continuing accountability dependency exists

fundraising phase materially changes
→ Campaign enters a different fundraising / accountability context

no post-fundraising account available
→ Organization report / evidence becomes available
   OR honestly remains pending / unavailable

Organization reports activity / outcome
→ reported operational progress / reported outcome exists

verification is applicable
→ Reviewer may produce bounded verification judgment

relevant information changes over time
→ chronology preserves prior state + new evidence + uncertainty

accountability progresses
→ remains pending / incomplete / unresolved
   OR reaches a legitimate completion condition
```

**Material actor dependencies**

Organization is the substantive accountability source; Owner / Staff act within authority; Reviewer contributes bounded fund-usage judgment when applicable; Kencleng preserves chronology, provenance, state meaning, and uncertainty; Donatur is the primary beneficiary.

**Important refinements**

- accountability does not begin only at Campaign closure;
- fundraising closure / phase change ≠ accountability closure;
- Organization report ≠ independent verification ≠ impact;
- honest negative, delayed, incomplete, or uncertain outcomes remain valid accountability;
- no report / incomplete information ≠ automatic success or failure;
- verification is not assumed mandatory for every Campaign;
- verification remains bounded and does not create universal truth;
- accountability visibility is not assumed entirely public or entirely Donatur-only;
- Owner is not assumed to approve every report;
- accountability ≠ one final report; evolving chronology is a strong Stage 3A hypothesis.

**OPEN PRODUCT DECISION exposed**

- exact meaning/consequence of Campaign or fundraising closure / phase change;
- minimum continuing accountability obligation;
- when accountability begins and may legitimately complete;
- whether timing/deadline rules create overdue / unmet states;
- meaning of pending / incomplete / unavailable / complete / unmet;
- public vs Donatur-specific accountability visibility;
- Owner vs Staff authority for accountability assertions;
- when fund-usage verification is applicable or mandatory;
- bounded positive meaning of verification outcomes;
- when new evidence supersedes / reopens prior report or verification state.

**Verdict:** R4 PASS and remains one representative scenario. Fund-usage verification stays a material branch rather than a separate scenario unless later evidence proves a distinct product question.

## 8. Scenario-level challenge verdict

```text
R1 — PASS
R2 — PASS
R3 — PASS
R4 — PASS
```

No scenario currently needs to be removed or split.

Two watch items remain:

1. **R2 is decision-heavy** and may later yield more than one delivery commitment.
2. **R4 is broad**, but reporting + verification still belong to one coherent accountability question at representative-scenario level.

Passing this challenge does **not** complete Stage 3A yet.

The next required step is a final cross-scenario audit for:
- contradiction across R1–R4;
- newly visible coverage gaps;
- wrong handoff / dependency assumptions;
- branches that now deserve promotion;
- whether the representative set is sufficient to feed Stage 3B without feature decomposition.


## 9. Final cross-scenario audit

### 9.1 R1 ↔ R2 handoff is not a mandatory linear dependency

The earlier shorthand `R1 → R2` is too strong if read as "Organization review must complete before Campaign work begins."

Current refinement:

```text
legitimate Organization representation
→ relevant prerequisite for real Campaign work

Organization-review completion
→ relationship to Campaign preparation / curation remains OPEN
```

R1 and R2 may overlap in the product lifecycle. Organization review can constrain or inform downstream Campaign behavior without automatically becoming a universal prerequisite.

### 9.2 R2 → R3 dependency is branch-specific

Actual public visibility is required for the public-understanding branch.

Donation eligibility is required only for contribution:

```text
public + not donation-eligible
→ Visitor can still understand Campaign
→ considered decision remains possible
→ contribution unavailable
```

This prevents the public experience from being reduced to a donation funnel.

### 9.3 R4 does not require Donatur existence as a prerequisite

The prior R4 starting condition was too restrictive.

Refined starting condition:

> A Campaign has entered a fundraising context and later materially progresses, changes, or ends while Organization accountability continues. If legitimate Donatur relationships exist, Donatur are primary beneficiaries of continuing accountability.

Therefore:

```text
Campaign accountability
≠ only exists because Donatur exists

legitimate donation relationship
→ creates direct continuing accountability dependency for Donatur
```

### 9.4 R3 explicitly includes actual current funding state

R3 public understanding includes the Campaign's actual current funding state and the meaning / provenance of that state.

```text
new donation fact
→ may change funding state only when legitimately supported
```

This closes the gap between individual donation semantics and later Campaign accountability without creating another representative scenario.

## 10. Branch-promotion audit

No current branch warrants promotion into a new representative scenario.

| Branch | Result |
| --- | --- |
| Organization materially changes after review | Remains R1 reconsideration branch |
| Campaign rejected | Remains R2 curation branch |
| Multiple review revisions | Remains R2 bounded-history branch |
| Campaign changes after positive curation | Remains R2 superseding / re-review branch |
| Public but donation-ineligible | Remains meaningful R2 → R3 branch |
| Visitor chooses no donation | Remains first-class R3 outcome |
| Pending / failed donation | Remains R3 donation-state branch |
| Zero-donation Campaign reaches closure | Handled by refined R4 |
| Reporting incomplete | Remains first-class R4 state |
| Verification unavailable / not applicable | Remains R4 branch |
| Platform correction | Remains X1 cross-cutting probe |

## 11. Final semantic map

The representative scenarios are not one mandatory linear pipeline.

```text
                 R1
 Organization legitimacy + review
          │
          │ constrains / informs
          ▼
                 R2
 Campaign proposition + curation
 visibility + donation eligibility
          │
          ├──────── public visibility
          │               ↓
          │              R3
          │     understanding + decision
          │
          └──────── donation eligibility
                          ↓
                 contribution branch
                          ↓
                   donation fact
                          ↓
                       Donatur
                          │
                          ▼
                         R4
                  continuing accountability

R4 also follows Campaign fundraising lifecycle
even when no Donatur exists.
```

## 12. Stage 3A final verdict

**Stage 3A — Representative Scenario Discovery: COMPLETE for Pilot #3 progression.**

Final representative set:

```text
R1 — Organization legitimacy + bounded Organization review
R2 — Campaign proposition + curation + actual visibility / donation eligibility
R3 — Public understanding + considered decision + donation / Donatur relationship
R4 — Fundraising change + reporting + continuing accountability / verification
X1 — Platform Operator cross-cutting exception probe
```

The set is considered sufficient to feed Stage 3B because it now exposes:

- representative product questions;
- starting and end conditions;
- major semantic transitions;
- material actor dependencies;
- meaningful branches and failure states;
- explicit OPEN PRODUCT DECISIONS;
- cross-scenario handoff boundaries.

Remaining OPEN PRODUCT DECISIONS should not all be resolved product-wide before sequencing. Stage 3B should identify which ones actually block a candidate delivery commitment.

Next active work:

> **Stage 3B — derive candidate commitments from R1–R4, test semantic readiness and minimum dependency closure, then establish a product dependency partial order and sequence the eligible frontier.**

## 13. 2026-10-21 Experience-first Stage 3A reconstruction checkpoint

> **Standing:** **Anhar-approved Stage 3A working reconstruction**, after proportionate final audit on 2026-10-21; sufficient to enter **fresh Stage 3B candidate evaluation**, but **not** newly canonical Kencleng Product Authority, an MVP/delivery decision, or runtime proof.
> **Authority:** [Pilot Product Intent](product-intent.md) is the upstream Human-approved seed. [Accepted Stage 1 Business Value Story](pilot-3-business-value-loop.md#9-2026-10-21-experience-first-stage-1-reconstruction-checkpoint) and [accepted Stage 2 actor outcomes](pilot-3-actor-outcomes.md#12-2026-10-21-experience-first-stage-2-reconstruction-checkpoint) are scoped pilot working baselines. Historical Stage 3A **sections 1–12** and [historical Stage 3B/C1 decisions](pilot-3-stage-3b-commitment-sequencing.md) remain intact and retain their original standing.
> **Experiment:** [Interconnected Journey Blueprint Research v0](https://github.com/anhsbolic/harscode-workspace/blob/74e6fcd45ad52ea3a66960712131b7fa209b836d/product/interconnected-journey-blueprint-research.md), optional research format only. This reconstruction belongs to Kencleng branch `pilot/station5-c1-space-only` and changes neither Kencleng nor Harscode operational `main`.

### 13.1 Scope and representative scenario verdict

The **one value anchor** remains: Organization can contribute attributable Campaign information and continuing accountability; Public Visitor can make a considered, non-coercive choice from truthful available information; legitimate Donatur relationships support continuing understanding, including known unknowns. This is intended product value, **not** observed user research or delivered software.

The historical representative set remains adequate; no new R5, persona inventory or mandatory stage workflow is justified:

| Probe | Human/business outcome pressure-tested | Reconstruction standing |
| --- | --- | --- |
| **R1 — Organization representation and review** | Legitimate, bounded internal action; independent Organization-review judgment stays separate | Covered; scoped C1 history preserved |
| **R2 — Campaign proposition and state** | Attributable proposition, bounded Campaign curation, actual **distinct** public visibility and donation eligibility | Covered; relationships among outcomes partly OPEN |
| **R3 — Visitor choice and legitimate donation** | Understand, choose not to donate, or attempt donation; only a legitimate donation fact makes the Donatur relationship | Covered; eligibility, funding and pending/failure meaning still need commitment-scoped decisions |
| **R4 — Campaign change and continuing accountability** | Organization accountability and truthful continuing information across fundraising change/closure, with Donatur as primary beneficiary **when one exists** | Covered, including **zero-donation Campaign**; reporting/verification completion specifics OPEN |
| **X1 — Platform Operator boundary** | Exceptional privileged intervention cannot claim Organization, Reviewer, donation/funding, settlement, verification or impact authority by privilege alone | Retain as **cross-cutting exception probe**, not a new product journey |

**Final light-audit result:** R1–R4 + X1 cover the material actors, positive and negative paths, authority/truth classes and handoffs sufficiently for Stage 3B exploration. No additional representative scenario was justified by the examined sources. **PASS = semantic reconstruction sufficiency only**, not commitment eligibility or field/runtime validation.

### 13.2 Cross-scenario pressure tests and typed edges

| Link / specific state-to-outcome relationship | Edge standing and pressure-test result |
| --- | --- |
| **R1 ↔ R2:** legitimate internal Organization authority → permitted Campaign-proposition action | **Authority prerequisite** for actions within its scope; completing Organization review is **not established as a universal prerequisite** to all Campaign preparation. Internal Kencleng Owner authority is not verified real-world/legal Organization authority. |
| **R1 ↔ R2:** bounded Organization review → Campaign curation/publication/eligibility | **OPEN / potentially conditional** relationship; do not create an automatic gate or assume positive review proves external legitimacy. |
| **R2 ↔ R3:** actual public Campaign state → accessible Campaign understanding | **State prerequisite** for the public context. A public but donation-ineligible Campaign can still provide Visitor understanding and a legitimate no-donation outcome. |
| **R2 ↔ R3:** actual donation eligibility → possible legitimate contribution | **Conditional prerequisite** for the contribution branch; positive curation **≠** public visibility **≠** donation eligibility. Eligibility may materially change and the effect on pending contribution requires a scoped rule. |
| **R3 → Donatur:** donation attempt → legitimate donation fact → contextual Donatur role | An attempt is **not sufficient**; the **legitimate fact** is the semantic prerequisite. Guest/registered account standing is separate. Donation fact alone does **not** entail funding/settlement. |
| **R2 ↔ R4:** actual Campaign/fundraising lifecycle change → continuing Organization accountability | **Downstream continuity independent of Donatur existence**; closure **≠** accountability completion. Exact substantive obligations and end criteria remain OPEN. |
| **R3 ↔ R4:** Donatur relationship → continuing direct Donatur interest in truthful developments | **Downstream relationship**, not a prerequisite for R4 to exist. Organization is primary substantive accountability source; reports and verification are not assumed universally available. |
| **Reviewer → R1/R2/R4:** scoped judgment / feedback / verification | **Conditional / feedback** when authorized, relevant and accessible. Reviewer cannot manufacture universal trust or impact; relevant conflict of interest bars judgment. |
| **X1 → R1–R4:** exceptional privileged correction | **Not a normal business-value prerequisite.** Intervention requires its own established operational mandate and factual basis; it may not fabricate or replace business/Reviewer truth. |

**Semantic dependency is not total delivery sequence.** The needed state for a public experience is not necessarily the entire upstream Organization journey. Actual Campaign visibility can support understanding **without** current donation eligibility. Organization accountability exists across relevant Campaign lifecycle even where there are zero Donatur. The value story therefore branches; it must not be compressed into a mandatory `R1 → R2 → R3 → R4` pipeline.

### 13.3 Negative/uncertain paths and exception coverage

- **R1:** Organization-review pending, relevant COI, material change and potential reconsideration; unauthorized Staff action does not gain Owner authority.
- **R2:** curation rejection / requested revision / superseding judgment, bounded history, public-but-ineligible, actual visibility change; no presumption that every edit mandates re-review.
- **R3:** considered no donation, eligibility loss during contribution, unresolved/failed attempt, or donation fact without legitimate funding implication; **intent/attempt/result** remain distinct.
- **R4:** zero-donation Campaign closure, no/incomplete/unverified report, limited audience visibility, changing information, and no assumed reporting breach without a valid duty/timing rule.
- **X1:** privileged operator cannot correct bad provenance or monetary presentation by pretending to be Organization, Reviewer, settlement provider, independent verifier or impact source.

These are **pressure-test branches**, not implementation cases declared delivered, a requirement to record a visitor's refusal, or invented product state-machine statuses.

### 13.4 Final-audit safeguards and authority boundaries

1. **R4 starts from Campaign fundraising/lifecycle context, not necessarily from a Donatur.** The first R4 scenario draft in historical section 7 assumed a donation relationship; historical section 9.3 **already corrected this**. This checkpoint adopts the corrected interpretation: Donatur is a beneficiary when a legitimate relationship exists, not the prerequisite for Organization accountability.
2. **Simulation is never actual money truth.** Kencleng is a learning sandbox: simulated payment or operational mechanisms must be explicitly identified as such and must **not** be presented as real external settlement, independent verification or real-world impact. Seeded/fixture behavior may support exploration, **not** claims that upstream capability has been delivered.
3. **Historical C1 is neither automatically chosen nor cancelled.** The [Human-approved scoped C1 representation model](pilot-3-stage-3b-commitment-sequencing.md#revised-c1-representation-model--superseding-decision) states that the first person establishing an Organization receives the initial **internal Kencleng Owner** role by product rule, without implying independent real-world authority. Historical C1 pre-engineering progression is not runtime delivery. Its **current meaningfulness, commitment boundary, relevance and sequencing priority** are open to fresh evaluation at Stage 3B; historical C2/C3 frontier placeholders are not automatically ELIGIBLE/BLOCKED.
4. **Project semantics and authority remain owned by Kencleng.** This pilot-only reconstruction is not a Kencleng `main` modification, new global product rule or Harscode standard. No UI/API, reporting cadence, Reviewer policy, Operator powers, scope or engineering assignments selected.

### 13.5 Stage 3A exit and Harscode experiment assessment

**Human decision:** Stage 3A final audit accepted and persistence approved by Anhar on 2026-10-21. Result: **PASS for Stage 3B exploration** of meaningful product outcome candidates and dependency-closed-enough boundaries. This is not Stage 3B commitment approval.

**Research observation:** Typed interconnections and actor-oriented fragments made it easier to communicate conditional vs true prerequisites, non-linear Campaign/accountability relationships, and authority distinctions. **However, much of the semantic substance was already present in historical Stage 3A and C1 sources.** The experiment has not yet demonstrated improved commitment selection, efficiency, independent cold-start consumption or Rancang handoff. **No reusable Harscode Product guidance change is justified yet.**

**Next Stage 3B task:** Generate fresh, meaningful outcome candidates from this connected value story, assess meaningfulness, semantic sufficiency, truthful real-product prerequisites, dependency closure and boundedness, then compare with *historical* C1/C2/C3 without assuming either that C1 remains first or that it must be discarded. Keep unresolved Product decisions explicit and resolve only those necessary for a chosen bounded commitment.

**Still unverified:** Actual Kencleng runtime behavior; independent consumer interpretation; real user evidence; downstream Design/Rancang effectiveness; any Engineering Space delivery. Reopen this Stage 3A reconstruction only on materially new evidence, not for exhaustive scenario enumeration.
