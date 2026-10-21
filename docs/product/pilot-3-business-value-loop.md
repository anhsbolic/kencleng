# Kencleng Pilot #3 — Business / Value Loop

> **Status:** WORKING ARTIFACT — Pilot #3 Stage 1 checkpoint  
> **Human status:** Anhar-approved working state for Pilot #3  
> **Authority:** Not canonical Product Authority. Derived from current `docs/product/product-intent.md` and applicable canonical Product Design authority.  
> **Working baseline:** `pilot/3-clean-delivery-baseline`  
> **Purpose:** Preserve the Stage 1 whole-product Business / Value Loop and the material semantic gaps carried into Stage 2.

## 1. Stage 1 question and boundary

Stage 1 asks:

> **How do value, responsibility, trust, and consequence move end-to-end through Kencleng?**

This checkpoint is intentionally broader than a detailed business process and narrower than a feature or delivery specification.

It preserves:

- whole-product state transitions;
- materially relevant actors and responsibility movement;
- value before, during, and after donation;
- trust and semantic boundaries;
- downstream dependencies;
- material product questions that remain open.

It does **not** define:

- feature inventory;
- screen or route structure;
- detailed permissions;
- reviewer procedures;
- delivery slices or release scope;
- API/data contracts;
- technical architecture or implementation.

## 2. Analysis approach used for this checkpoint

The following was the working analysis technique used to make Stage 1 concrete. It is not a new canonical Kencleng methodology.

```text
1. Reconstruct whole-product breadth from current Product Authority
   ↓
2. Model L1 as meaningful before → after states
   ↓
3. Inspect the model through intersecting threads
   - Organization / Campaign
   - Person
   - Trust / Evidence
   ↓
4. Identify L1 transitions hiding material semantic changes
   ↓
5. Expand only those transitions into L2
   ↓
6. Challenge each material transition for:
   - actor / responsibility
   - before → after business meaning
   - value / consequence
   - next dependency
   - trust implication
   - semantic collapse
   - unsupported assumption
   ↓
7. Classify findings as authority, evidence, hypothesis, open decision, or parked
   ↓
8. Audit Stage 1 exit criteria
   ↓
9. Decide GO / NO-GO for Stage 2
```

L2 expansion stops before detailed operational procedure, feature decomposition, or implementation mechanics.

## 3. Business / Value Loop — L1

The whole-product flow is best understood as three intersecting threads rather than one forced linear funnel.

### 3.1 Organization / Campaign thread

```text
Organization does not yet have established Kencleng context
→
Organization becomes established with identifiable responsibility context
and a distinct organization-review state

No Campaign proposition exists
→
A Campaign is prepared by an Organization
and has a distinct campaign-curation state

Campaign is not yet in a public / donation context
→
Campaign has actual public-visibility state
and actual donation-eligibility state

Fundraising is ongoing
→
Campaign reaches a closure state / fundraising phase materially changes

Fundraising phase changes or ends
→
post-campaign reporting and accountability continue
```

Important: organization review, campaign curation, public visibility, and donation eligibility are related product concepts but are not automatically equivalent or sequential prerequisites.

### 3.2 Person thread

```text
Public visitor does not yet understand a Campaign sufficiently
→
visitor gains enough context and evidence to make a considered decision

Considered decision
├─ do not contribute
│  → valid decision outcome; no donation fact is created
│
└─ contribute
   → donation fact exists
   → person becomes a Donatur in the context of that Campaign

Donatur has contributed
→
Donatur can continue understanding what happened after contribution
```

A Donatur is not defined by having a registered account. Current Product Authority allows a visitor to understand a public Campaign without signing in and may allow guest contribution when applicable product rules permit it.

### 3.3 Trust / Evidence thread

```text
Information exists with different truth classes
→
platform-known facts, organizer-provided information, review outcomes,
system states, reports, and unknown / pending information remain distinguishable

Campaign is in public context
→
visitor can understand relevant steward, purpose, actual state,
funding context, uncertainty, provenance, and consequential money meaning

Donation fact exists
→
the contribution is known as a donation fact
≠ automatic funding consequence

Funding state changes when legitimately supported
→
funding progress can be understood
≠ execution

Operational or reporting evidence appears
→
operational progress and/or reported outcome can be understood
≠ verified outcome
≠ real-world impact

Fund-usage verification occurs when applicable
→
verification has bounded meaning
≠ universal proof of impact
```

Unknown, pending, delayed, changed, or unavailable information remains a valid product state throughout the flow.

## 4. Material L2 challenge results

### L2-1 — Organization context → Campaign proposition

Supported:

- Organization and Campaign are distinct.
- A Campaign belongs to an Organization.
- Organization review is distinct from Campaign curation.
- Organization Owners and Staff have meaningfully different authority.

Not assumed:

- organization review must complete before Campaign preparation;
- organization review automatically authorizes curation, publication, or donation eligibility;
- a reviewed Organization is globally “trusted” or “accountable.”

Material open relationship:

> What does an Organization review outcome positively mean, and which downstream Campaign responsibilities legitimately depend on it?

### L2-2 — Campaign proposition → public visibility / donation eligibility

Supported semantic boundary:

```text
Campaign curation
≠ public visibility
≠ donation eligibility
```

The Campaign may have a curation outcome/state, public-visibility state, and donation-eligibility state. Their exact relationship is deliberately not inferred.

The product must not expose visibility or donation capability in a way that contradicts the Campaign’s actual state.

### L2-3 — Public Campaign → considered decision → Donatur / no donation

Value before donation is not conversion itself.

The supported product value is:

> A visitor can understand a fundraising effort well enough to make a considered, non-coercive decision.

Both outcomes are legitimate:

```text
understand → do not donate
understand → choose to contribute
```

If contribution occurs, a donation fact may exist. That fact must not be silently upgraded into a funding consequence, execution claim, verification claim, or impact claim.

### L2-4 — Donation fact → Donatur relationship → continued accountability

Donation is not the end of the product relationship.

Current Product Authority supports:

```text
contribution
→ donation fact
→ continuing donatur interest in what happened afterward
→ continuing post-donation accountability
```

The exact allocation of post-donation obligations among Organization representatives, reviewers, and platform operators remains open for Actor Outcomes.

### L2-5 — Fundraising closure → post-campaign accountability

Supported semantic boundary:

```text
fundraising closure
≠ accountability closure
```

Campaign closure changes the fundraising phase, but post-campaign reporting and accountability continue.

The exact trigger and consequence of closure, its relationship to donation eligibility, and the conditions under which accountability is complete, pending, or unmet remain open.

## 5. Semantic distinctions that must remain preserved

These distinctions are upstream constraints carried into later stages:

```text
Organization
≠ Campaign

Organization review
≠ Campaign curation
≠ fund-usage verification

Campaign curation
≠ public visibility
≠ donation eligibility

Donation fact
≠ automatic funding consequence

Funding
≠ execution
≠ reported outcome
≠ verified outcome
≠ impact

Donatur
≠ necessarily registered account

Campaign / fundraising closure
≠ accountability closure
```

Simulation or sandbox mechanisms must never be represented as real external settlement, independent verification, provenance, authorization, or real-world impact.

## 6. Finding classification

### CANONICAL / SETTLED AUTHORITY carried into this checkpoint

- Kencleng is evidence-led and must support considered decisions before contribution.
- Organization and Campaign are distinct; Campaign belongs to an Organization.
- Organization Owners and Staff have meaningfully different authority.
- Organization review, Campaign curation, and later fund-usage verification are separate decisions.
- Reviewer conflict of interest must be avoided for relevant work.
- Public visitors may understand a public Campaign without signing in.
- Guest contribution may be permitted by applicable product rules.
- Public visibility and donation eligibility must remain tied to actual Campaign state.
- Funding progress is not proof of execution or impact.
- Trust classes, provenance, state, report, and uncertainty must remain distinguishable.
- Accountability continues after contribution / fundraising.
- Unknown or pending information is a valid product state.
- Simulation must never masquerade as real settlement, independent verification, or real-world impact.

### OBSERVATION / EVIDENCE

The three-thread representation — Organization / Campaign, Person, and Trust / Evidence — proved useful for exposing semantic gaps without forcing a false single-track process.

The L1 → selective L2 before/after technique also proved useful for this checkpoint. Neither is automatically reusable Kencleng or Harscode policy.

### HYPOTHESIS

A contribution creates an ongoing Donatur ↔ Campaign relationship that is a useful anchor for understanding continuing accountability.

This is strongly consistent with Product Intent, but the exact relationship semantics remain to be tested in Actor Outcomes and later scenario/interaction work.

### OPEN PRODUCT DECISION — carry forward

1. What does a successful or unsuccessful Organization review positively mean?
2. What does Campaign curation positively mean?
3. What is the exact relationship among curation, public visibility, and donation eligibility?
4. When and how does a donation fact legitimately produce a funding consequence?
5. What exactly changes when a Campaign reaches closure?
6. Who owes which post-donation / post-fundraising accountability responsibility to whom?
7. What product-level responsibility does the platform operator own?
8. When is accountability considered complete, pending, or unmet?
9. Which post-campaign accountability information is public, donatur-specific, or otherwise visibility-bounded?

### PARKED — NOT NEEDED YET

- exact Owner vs Staff permissions;
- detailed review and curation procedures;
- review checklists;
- exact reporting cadence;
- notification mechanisms;
- UI/status labels;
- feature and surface decomposition;
- operational shortcuts;
- release scope and delivery sequencing;
- API/data/implementation mechanics.

## 7. Stage 1 exit audit

| Exit criterion | Result |
| --- | --- |
| Whole-product flow is understandable without implementation detail | PASS |
| Major actors / responsibility transfers are visible enough for this stage | PASS |
| Value before donation can be explained | PASS |
| Value at donation can be explained without semantic upgrade | PASS |
| Value after donation can be explained | PASS |
| No obvious missing link invalidates downstream reasoning | PASS |
| Critical semantic distinctions remain preserved | PASS |
| Remaining gaps can be carried as explicit open product decisions | PASS |

## 8. Stage 1 verdict

**Stage 1 — Business / Value Loop: COMPLETE for Pilot #3 progression.**

There is no known product-level blocker requiring Stage 1 to remain open.

The next active stage is:

> **Stage 2 — Actor Outcomes & Responsibilities**

Stage 2 should use this checkpoint to determine for each materially relevant actor:

- desired outcome;
- responsibility / obligation;
- dependency;
- required understanding;
- legitimate consequence when responsibility is fulfilled, pending, or failed;
- open semantics that still must not be assumed.

This working artifact should be revised only if later stages produce materially new evidence that invalidates or meaningfully changes the Stage 1 model.

## 9. 2026-10-21 Experience-first Stage 1 reconstruction checkpoint

> **Standing:** Anhar-accepted **working reconstruction v0.2** for the Harscode foundation-refactor / Kencleng pilot experiment, not new canonical Kencleng Product Authority.
> **Owner and approval:** Anhar accepted the Stage 1 Business Value Story on 2026-10-21; this checkpoint records the approved *working interpretation*, not an approved MVP, feature sequence, C1 revision, or engineering assignment.
> **Source and precedence:** [Pilot Product Intent](product-intent.md) is the upstream Human-approved seed; [Product Design Principles](../ui-ux/product-design-principles.md) and [Brand/Product UI Brief](../ui-ux/brand-product-ui-brief.md) govern their own experience/design concerns. This section supplements the previous Stage 1 pilot observations above without retroactively rewriting historical approvals. The active Kencleng `main` authority is **not** changed by this pilot-only experiment.

### 9.1 Core framing — one connected story

**Question:** How do human value, responsibility, trust, and consequential business states move *from an initial unmet need to a meaningful, truthful outcome* through Kencleng?

**FROM — initial human and business situation.** A person may want to understand or support a social cause without yet having sufficient, intelligible and provenance-aware information to make a considered decision. An Organization may seek participation in Kencleng to communicate a fundraising effort and its subsequent accountability. These motives are **plausible representative situations**, not independently verified user-research findings.

**VALUE TRANSFORMATION.** Kencleng connects Organization participation, Campaign information, considered public choice, legitimate donation relationships and continuing accountability. The platform preserves the differences among Organization-provided assertions, platform-known facts, bounded independent judgments, actual states, money consequences, reports and unknown/pending information. It makes evidence-aware confidence and realistic optimism possible without manufacturing trust, false certainty, urgency or emotional pressure.

**TO — intended meaningful outcome.** A Public Visitor can decide whether to donate based on intelligible context and evidence; **not donating is a legitimate outcome**. If a legitimate donation fact occurs, the person becomes a Donatur in that Campaign context and can continue understanding what happened afterward. Organization representatives can provide and account for information within bounded legitimate authority. Reviewer judgment adds only its actual scoped meaning. This expresses intended product value, not already-verified usage or delivery.

### 9.2 Business-value continuity (not delivery sequence)

1. **Organization / Campaign:** participation and attributable responsibility context → Organization-provided Campaign proposition → appropriately distinct curation, public visibility and donation-eligibility states → material fundraising changes and closure → continuing reporting/accountability. Organization review is a separate bounded judgment, **not automatically a prerequisite** of Campaign preparation or publication.
2. **Person:** initial interest, question or possible intention to help → encounters and understands Campaign context → considered decision (donate **or not**) → if contribution is attempted, actual eligibility and result matter → **only a legitimate donation fact** creates the relevant Donatur role → continuing understanding of subsequent developments.
3. **Trust / Evidence:** provenance, relevant chronology, known/unknown states and consequential meaning accompany the relationship throughout. Donation fact **does not automatically imply funding**; funding **does not imply execution**; reported outcome **does not imply independently verified impact**; fundraising closure **does not close accountability**.

These are **intersecting business-value threads**, not a mandatory linear funnel or Engineering plan. They name the value and main handoffs without selecting independent product commitments.

### 9.3 Main actor/value handoffs

| Actor / responsibility | Business-level From → To | Dependency or boundary |
| --- | --- | --- |
| Organization representatives (Owner/Staff authority distinguished) | Organization participation and factual stewardship not yet established → can provide, maintain and account for Organization-sourced context within legitimate authority | Their assertions are not Reviewer judgments or platform-known facts. Exact role permissions remain open. |
| Reviewer (when applicable) | Relevant material requires bounded judgment → scoped, provenance-aware judgment becomes available | Independence and conflict-of-interest boundary; review is not universal verification or automatic downstream eligibility. |
| Public Visitor | Interested or seeking understanding, insufficient context → informed and non-coercive choice | Depends on truthful accessible Campaign information, not a guaranteed favorable judgment or donation. |
| Donatur | Legitimate donation relationship exists → can continue understanding progress, reports and relevant unknowns | Organization accountability and truthful state/evidence remain material after contribution/closure. |
| Platform Operator | Positive intervention role not yet established → any intervention remains within bounded operational authority | Privilege cannot manufacture business evidence, Organization assertions, independent review or money/impact truth. |
| Kencleng Product/System (not a human persona) | Disconnected assertions and events → inspectable, truthful product states and boundaries | Maintains state/provenance semantics; implementation mechanism intentionally open. |

### 9.4 Product Intent and approved Product Design alignment

Approved design principles enrich **the quality of the experience** without independently creating business rules: **Evidence-Led Optimism**, **Confidence Before Conversion**, **Facts First, Story With Them**, **Dignity Over Pity**, **Trust Comes From Structure**, **Unknown Is a Valid State**, and **Progress Is Evidence, Not Gamification**.

The Brand/Product UI sequence *evidence → confidence → optimism → action* expresses an **intended experience posture**, not a mandatory conversion funnel or a claim that evidence guarantees confidence, trust or donation. Experience and visual principles remain owned by their approved design sources; this Stage 1 checkpoint does not mandate a particular UI, Evidence Journal feature, or delivery commitment.

### 9.5 Scope and carried-forward uncertainties

Stage 1 **does not** determine detailed actor journeys, initial Owner mechanics, C1 eligibility/scope, permission tables, review procedures, curation/publication/donation prerequisites, funding transitions, accountability completion criteria, full dependency ordering, feature inventory, UI, API, architecture or release sequence. The open semantic questions in section 6 above remain open for appropriate Human-owned decisions **when a later bounded journey makes them material**.

In particular, this reconstruction **does not** supersede historical approved C1-specific behavior, nor does it claim that the historical Stage 3B outcome is necessarily wrong. Later stages will test whether cross-actor dependencies were correctly carried into commitment selection.

### 9.6 Stage 1 exit and experiment handoff

**Stage 1 verdict: PASS — sufficient for the next Product decision (Anhar-accepted working reconstruction).** The entire Product value can be narrated from human need to meaningful result; actor responsibility transfers, crucial trust boundaries and known open questions are visible without implementing or prescribing detailed journeys.

**Next experiment:** derive actor-specific initial situations, motivations, desired outcomes, responsibilities and connected dependencies from **this same value story** in Stage 2, then test experience continuity through Stage 3A/3B. An Interconnected Journey Blueprint is the experimental visualization approach, **not** a presently mandatory Harscode methodology or a finished Stage 1 deliverable.

**Harscode checkpoint #1:** Product Foundation guidance was sufficient to produce this story and preserve the Product/Design authority distinction. **No change to reusable Harscode Product guidance justified at Stage 1**. Reassess only on material Stage 2–3 evidence. No real Kencleng behavior, Engineering Space execution, fresh-reader validation or runtime effectiveness has been verified by this documentation exercise.
