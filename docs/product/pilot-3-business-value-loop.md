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
   → person becomes a Donor in the context of that Campaign

Donor has contributed
→
Donor can continue understanding what happened after contribution
```

A Donor is not defined by having a registered account. Current Product Authority allows a visitor to understand a public Campaign without signing in and may allow guest contribution when applicable product rules permit it.

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

### L2-3 — Public Campaign → considered decision → Donor / no donation

Value before donation is not conversion itself.

The supported product value is:

> A visitor can understand a fundraising effort well enough to make a considered, non-coercive decision.

Both outcomes are legitimate:

```text
understand → do not donate
understand → choose to contribute
```

If contribution occurs, a donation fact may exist. That fact must not be silently upgraded into a funding consequence, execution claim, verification claim, or impact claim.

### L2-4 — Donation fact → Donor relationship → continued accountability

Donation is not the end of the product relationship.

Current Product Authority supports:

```text
contribution
→ donation fact
→ continuing donor interest in what happened afterward
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

Donor
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

A contribution creates an ongoing Donor ↔ Campaign relationship that is a useful anchor for understanding continuing accountability.

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
9. Which post-campaign accountability information is public, donor-specific, or otherwise visibility-bounded?

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
