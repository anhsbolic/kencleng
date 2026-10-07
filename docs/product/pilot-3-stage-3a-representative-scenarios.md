# Kencleng Pilot #3 — Stage 3A Representative Scenarios

> **Status:** WORKING ARTIFACT — Pilot #3 Stage 3A checkpoint  
> **Human status:** Anhar-approved working state through candidate-scenario coverage audit  
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
