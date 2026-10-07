# Kencleng Pilot #3 — Stage 3 Working Model

> **Status:** WORKING ARTIFACT — Pilot #3 Stage 3 method checkpoint  
> **Human status:** Anhar-approved direction; stress-tested working approach  
> **Authority:** Not canonical Product Authority or Harscode methodology  
> **Working baseline:** `pilot/3-clean-delivery-baseline`  
> **Purpose:** Preserve the tested Stage 3A / Stage 3B working model used to discover representative scenarios and sequence delivery commitments without recreating waterfall or hiding unresolved dependencies.

## 1. Stage 3 working split

```text
Stage 3A — Representative Scenario Discovery
→ discover scenarios that pressure-test product truth,
  actor responsibilities, dependencies, and material open semantics

Stage 3B — Commitment Sequencing
→ derive candidate delivery commitments from those scenarios
→ test whether they are ready to enter a Stage 4–7 delivery loop
→ establish dependency-aware ordering
```

A representative scenario is not automatically one delivery commitment.

```text
representative scenario
→ exposes behavior + dependencies + open semantics
→ candidate commitment(s)
→ Stage 3B eligibility
→ selected commitment
→ Stage 4–7 delivery loop
```

Stage 4 is therefore not intended to be completed product-wide before engineering starts.

## 2. Delivery loop

The current Pilot #3 working progression is:

```text
Stage 1 — Business / Value Loop
Stage 2 — Actor Outcomes & Responsibilities
Stage 3A — Representative Scenario Discovery
Stage 3B — Commitment Sequencing
        ↓
select one eligible commitment
        ↓
Stage 4 — Interaction Exploration
Stage 5 — Confirmed Product Behavior
Stage 6 — Sufficient Experience + Engineering-facing Requirements
Stage 7 — Durable Handoff
        ↓
delivery / learning
        ↓
select next commitment
```

This is intended to preserve iterative delivery rather than complete all detailed product interaction upfront.

## 3. Exploration lane vs delivery lane

Later product behavior may be explored before all of its delivery dependencies are real.

```text
EXPLORATION LANE
prototype / simulation / prepared data
→ learn interaction
→ expose product questions
→ evidence only

DELIVERY LANE
real product behavior
→ all product dependencies required for the claimed behavior
  must already be real or be included in the same bounded commitment
```

A downstream prototype does not satisfy missing upstream delivery dependencies.

Simulation, fixtures, or prepared data must remain explicit and must never be presented as real settlement, verification, evidence, or delivered upstream capability.

## 4. Candidate commitment eligibility

A candidate commitment may enter the Stage 4–7 delivery loop only when it passes all material gates below.

### 4.1 Meaningful

The commitment creates a real actor outcome, responsibility capability, understanding, or legitimate product-state consequence.

Internal CRUD/state plumbing alone is not sufficient.

### 4.2 Semantically ready enough

The product meaning required for the commitment is sufficiently understood.

If an `OPEN PRODUCT DECISION` determines the fundamental behavior or consequence, the commitment is not delivery-ready until that meaning is resolved enough.

An `OPEN EXPERIENCE / DESIGN DECISION` is not automatically a blocker; Stage 4 exists to explore those questions.

### 4.3 Dependency-closed enough

Every product dependency required for the commitment's claimed behavior is one of:

- already real and sufficiently established;
- explicitly included in the same commitment; or
- proven not to be required for that claimed behavior.

A simulated or seeded upstream capability does not count as a delivered dependency.

### 4.4 Truthful

The commitment can be demonstrated without implying product truth that does not exist.

Prepared data may support exploration, but cannot be used to claim that missing Organization, Campaign, review, eligibility, donation, settlement, verification, or accountability behavior is delivered.

### 4.5 Bounded

The commitment can reasonably pass through Stage 4–7 as one coherent delivery loop.

Dependency closure must not be achieved by absorbing the whole upstream product into one giant commitment.

## 5. Minimum necessary dependency closure

> **Dependency closure is behavioral, not domain-completeness.**

Do not ask whether the whole upstream domain is complete.

Ask:

> **What is the minimum real upstream product behavior required for this commitment's product claim to be truthful?**

A missing prerequisite may travel inside the same commitment only when:

- it directly supports the same coherent actor outcome;
- it does not introduce a new large product problem;
- the resulting commitment remains bounded.

Otherwise, it remains a predecessor commitment.

Example:

```text
Donation does not require all Organization-management capability to be complete.

It does require enough real Organization + Campaign + actual eligibility behavior
for the donation claim to be legitimate.
```

## 6. Dependency ordering is partial, not automatically linear

Stage 3B should produce a **product dependency graph / partial order**, not automatically one total roadmap.

```text
A → C
B → C

A and B may both be eligible
without A needing to precede B.
```

Independent eligible commitments may potentially proceed in parallel or be ordered later using product value, uncertainty, and boundedness considerations.

Do not invent a linear order where no product dependency exists.

## 7. Blocker classification

When a candidate is not ready, classify why.

| Blocker | Meaning | Stage 3B response |
| --- | --- | --- |
| **OPEN PRODUCT DECISION** | Fundamental product behavior/consequence not yet defined | Resolve only enough product meaning to make the commitment coherent |
| **Product dependency gap** | Required upstream product behavior is absent | Include minimum prerequisite if bounded; otherwise keep predecessor |
| **OPEN EXPERIENCE / DESIGN DECISION** | Product meaning sufficient, interaction unresolved | Normally valid input to Stage 4 |
| **OPEN ENGINEERING CONSTRAINT** | Material feasibility/technical constraint unresolved | Preserve explicitly; do not invent architecture in Stage 3B |
| **Exploration-only simulation** | Downstream can be learned but real dependency is absent | Explore, but prohibit delivery claim |

If a known engineering constraint makes product feasibility materially uncertain, a commitment may remain conditional until enough evidence exists.

## 8. Ordering eligible commitments

Only compare commitments after they pass material eligibility gates.

Use qualitative considerations, not weighted scoring:

- meaningful actor value;
- upstream-enabling effect;
- useful downstream behavior unblocked;
- material uncertainty reduced;
- truthful end-to-end behavior created;
- boundedness / ability to complete one loop.

**Upstream-enabling is not an absolute priority.**

An upstream commitment that creates no meaningful actor/product outcome should not automatically outrank a more coherent meaningful commitment merely because it sits earlier in a dependency graph.

Prefer the smallest coherent commitment that creates meaningful truthful behavior while unlocking useful downstream progression.

## 9. Stress-test audit

| Failure mode | Verdict | Refinement |
| --- | --- | --- |
| Missing dependency can be included in the same commitment | PASS | Allow only when it supports the same coherent outcome and remains bounded |
| Independent commitments are forced into a sequence | PASS after refinement | Use partial order / dependency graph |
| Semantic blocker is confused with dependency blocker | PASS after refinement | Explicit blocker classification |
| Downstream prototype reveals an upstream problem | PASS | Material new evidence may reopen working decisions |
| Commitment is dependency-closed but meaningless | PASS after refinement | Strengthen the Meaningful gate |
| Commitment becomes huge just to be closed | PASS after refinement | Minimum necessary closure + Bounded gate |
| “Upstream first” recreates waterfall | PASS after refinement | Upstream-enabling is only one ordering consideration |
| Prepared data hides missing readiness | PASS strongly | Exploration allowed; delivery claim prohibited |

## 10. Guest-donation sanity test

Guest donation is a useful downstream representative scenario and exposes the failure this method is intended to prevent.

A simplified product dependency chain is:

```text
legitimate Organization context
→ real Campaign proposition
→ actual Campaign state
→ donation eligibility
→ considered public decision
→ donation
→ Donatur relationship
```

If Organization / Campaign / eligibility are replaced with seeded or prepared data:

- donation interaction may still be explored;
- the prototype may still generate useful evidence;
- donation is **not** dependency-closed as a delivered product commitment.

Seeder/prepared data can support exploration, but cannot make a missing prerequisite count as delivered.

## 11. Stage 3B working method v1

```text
Stage 3A representative scenario
        ↓
derive candidate commitment(s)
        ↓
state actor outcome / product claim
        ↓
map only product dependencies required for that claim
        ↓
classify unresolved blockers
        ↓
eligibility gates:
meaningful?
semantically ready enough?
dependency-closed enough?
truthful?
bounded?
        ↓
NO ──→ scenario / exploration / predecessor / open decision
YES
        ↓
place in product dependency partial order
        ↓
compare eligible frontier:
meaningful value
+ enabling effect
+ uncertainty reduction
+ truthful behavior
+ boundedness
        ↓
select delivery commitment
        ↓
Stage 4–7 loop
```

## 12. Verdict

**Stage 3B method v1: PASS as a Pilot #3 working approach.**

It is considered fit to use during Stage 3A scenario discovery and subsequent commitment selection because it:

- detects downstream commitments whose upstream readiness is fake;
- permits early downstream exploration without false delivery claims;
- allows minimum prerequisite behavior to travel with a commitment;
- avoids requiring complete upstream domains;
- supports parallel/partial ordering rather than forced linear planning;
- distinguishes product, dependency, experience, and engineering questions;
- preserves meaningful actor value as a requirement;
- provides a bounded entry point into Stage 4–7 iteration.

This remains experimental. Revise it if actual Stage 3A usage or later delivery evidence reveals a material failure mode.
