# Kencleng Pilot #3 — Control Tower

> **CONTROL TOWER / PROGRESS VIEW — NOT PRODUCT AUTHORITY**  
> This dashboard visualizes current position, lineage, dependency, readiness, and progress.  
> Owning Product / working artifacts remain authoritative for meaning and decisions. If this dashboard conflicts with an owning artifact, the owning artifact wins.
>
> **Status:** Experimental living dashboard for Pilot #3  
> **Working baseline:** `pilot/3-clean-delivery-baseline`

## Control Tower

Read left → right as product maturity / progression.

- **Stage areas are stations.**
- **Nodes are tracked product-work pointers.**
- **Commitments are the nodes that can travel from Stage 3B through Stage 4–7.**
- An ELIGIBLE commitment stays physically in Stage 3B until Stage 4 actually begins.
- A **READY** rail means it may depart; it does not mean it already moved.

```mermaid
flowchart LR

  subgraph S1["STAGE 1 · Business / Value Loop"]
    direction TB
    V1(["Org / Campaign ✓"])
    V2(["Person ✓"])
    V3(["Trust / Evidence ✓"])
  end

  subgraph S2["STAGE 2 · Actor Outcomes"]
    direction TB
    A1(["Visitor / Donatur ✓"])
    A2(["Owner / Staff ✓"])
    A3(["Reviewer ✓"])
    A4(["Operator ✓"])
  end

  subgraph S3["STAGE 3 · Scenario + Commitment Frontier"]
    direction TB
    R1(["R1 ✓"])
    R2(["R2 ✓"])
    R3(["R3 ✓"])
    R4(["R4 ✓"])
    X1(["◌ X1"])

    C2(["C2 ⛔"])
    C3(["C3 ⏳"])
  end

  subgraph S4["STAGE 4 · Interaction Exploration"]
    direction TB
    C1S4(["C1 ✓"])
  end

  subgraph S5["STAGE 5 · Confirmed Behavior"]
    direction TB
    C1(["C1 ▶"])
  end

  subgraph S6["STAGE 6 · Requirements"]
    direction TB
    G6[" "]
  end

  subgraph S7["STAGE 7 · Durable Handoff"]
    direction TB
    G7[" "]
  end

  V1 --> A2
  V2 --> A1
  V3 --> A3
  V3 --> A4

  A2 --> R1
  A2 --> R2
  A2 --> R4
  A3 --> R1
  A3 --> R2
  A3 --> R4
  A1 --> R3
  A1 --> R4
  A4 --> X1

  R1 --> C1
  R1 --> C2
  R2 --> C3

  C1 -. "dependency" .-> C3

  classDef complete fill:#dcfce7,stroke:#15803d,color:#14532d,stroke-width:1.5px;
  classDef eligible fill:#dbeafe,stroke:#1d4ed8,color:#1e3a8a,stroke-width:2px;
  classDef active fill:#e0f2fe,stroke:#0369a1,color:#0c4a6e,stroke-width:2.5px;
  classDef blocked fill:#fee2e2,stroke:#b91c1c,color:#7f1d1d,stroke-width:2px;
  classDef waiting fill:#fef3c7,stroke:#b45309,color:#78350f,stroke-width:2px;
  classDef probe fill:#f3f4f6,stroke:#6b7280,color:#374151;
  classDef ghost fill:transparent,stroke:transparent,color:transparent;

  class V1,V2,V3,A1,A2,A3,A4,R1,R2,R3,R4 complete;
  class C1S4 complete;
  class C1 active;
  class C2 blocked;
  class C3 waiting;
  class X1 probe;
  class G6,G7 ghost;

  style S1 fill:#f0fdf4,stroke:#86efac,stroke-width:1px
  style S2 fill:#f0fdf4,stroke:#86efac,stroke-width:1px
  style S3 fill:#ffffff,stroke:#d1d5db,stroke-width:1px
  style S4 fill:#f0fdf4,stroke:#86efac,stroke-width:1px
  style S5 fill:#eff6ff,stroke:#2563eb,stroke-width:2px
  style S6 fill:#ffffff,stroke:#d1d5db,stroke-width:1px
  style S7 fill:#ffffff,stroke:#d1d5db,stroke-width:1px
```

### Signal legend

| Signal | Dashboard meaning |
| --- | --- |
| ✓ COMPLETE | Output is sufficiently complete for progression |
| ▶ ACTIVE | Currently being worked through |
| ◆ ELIGIBLE | Commitment may enter Stage 4–7 now |
| ⛔ BLOCKED | A material semantic / product blocker remains |
| ⏳ WAITING | A material dependency must become real first |
| ◌ PROBE | Cross-cutting exceptional probe; node type, not progress status |

### Rail legend

```text
────────▶  derivation / actual progression
- - - - ▶  product dependency
- READY ▶  eligible to enter the next station; not moved yet
·······▶  feedback / reopen when materially new evidence appears
```

## Current Stage 3 frontier

| Node | Type / source | Current station | Signal | Current meaning |
| --- | --- | --- | --- | --- |
| **R1 — Organization legitimacy + review** | Representative scenario | Stage 3A | ✓ COMPLETE | Scenario pressure-test complete |
| **R2 — Campaign curation + actual state** | Representative scenario | Stage 3A | ✓ COMPLETE | Scenario pressure-test complete |
| **R3 — Public decision + donation** | Representative scenario | Stage 3A | ✓ COMPLETE | Scenario pressure-test complete |
| **R4 — Continuing accountability** | Representative scenario | Stage 3A | ✓ COMPLETE | Scenario pressure-test complete |
| **X1 — Operator intervention** | Cross-cutting probe | Stage 3A | ◌ PROBE | Use only when a core scenario materially requires exceptional operator intervention |
| **C1 — Legitimate Organization representation** | Commitment from R1 | Stage 5 | ▶ ACTIVE | Stage 4 complete; behavior confirmation is active |
| **C2 — Bounded Organization review** | Commitment from R1 | Stage 3B | ⛔ BLOCKED | Positive meaning / legitimate Organization-review outcomes remain unresolved |
| **C3 — Real Campaign proposition** | Commitment from R2 | Stage 3B | ⏳ WAITING | Depends on C1 becoming real; resolve C3-specific semantics when C3 becomes material |

### Position rule

```text
C1 ▶ ACTIVE
now lives in Stage 5

C2 ⛔ and C3 ⏳
remain parked in Stage 3B
```

C1 is now the only active delivery commitment. C2 / C3 stay parked unless interaction evidence reveals a material dependency required for C1 to remain truthful.

If several independent commitments are simultaneously ELIGIBLE, compare only that eligible frontier. Do not force a full-product sequence.

## Owning artifacts

| Dashboard area | Owning artifact |
| --- | --- |
| Product direction | [Product Intent](./product-intent.md) |
| Stage 1 | [Pilot #3 — Business / Value Loop](./pilot-3-business-value-loop.md) |
| Stage 2 | [Pilot #3 — Actor Outcomes & Responsibilities](./pilot-3-actor-outcomes.md) |
| Stage 3 method | [Pilot #3 — Stage 3 Working Model](./pilot-3-stage-3-working-model.md) |
| Stage 3A | [Pilot #3 — Representative Scenarios](./pilot-3-stage-3a-representative-scenarios.md) |
| Stage 3B | [Pilot #3 — Commitment Sequencing](./pilot-3-stage-3b-commitment-sequencing.md) |
| Stage 4 — C1 | [Pilot #3 — C1 Interaction Exploration](./pilot-3-stage-4-c1-interaction-exploration.md) |
| Stage 5 — C1 | [Pilot #3 — C1 Confirmed Product Behavior](./pilot-3-stage-5-c1-confirmed-behavior.md) |

## Dashboard update discipline

Update the Control Tower **after** the owning artifact changes when a meaningful progress state changes, for example:

- a stage/output becomes COMPLETE;
- a commitment becomes ELIGIBLE, BLOCKED, or unblocked;
- a material dependency becomes real / satisfied;
- a commitment actually enters Stage 4, 5, 6, or 7;
- materially new evidence reopens an earlier working decision.

Do **not** create Product Truth here.

```text
discussion / owning artifact
→ durable decision
→ Control Tower reflects the new state
```

The Control Tower should stay compact and operational. Detailed reasoning remains in the linked owning artifact.
