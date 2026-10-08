# Kencleng Pilot #3 — Control Tower

> **CONTROL TOWER / PROGRESS VIEW — NOT PRODUCT AUTHORITY**  
> This dashboard visualizes current progress, lineage, dependency, and readiness.  
> Owning Product / working artifacts remain authoritative for meaning and decisions.  
> If this dashboard conflicts with an owning artifact, the owning artifact wins.
>
> **Status:** Experimental living dashboard for Pilot #3  
> **Working baseline:** `pilot/3-clean-delivery-baseline`

## 1. Visual grammar

- **Stage area / station** — maturity area from Stage 1 through Stage 7.
- **Node / train** — a tracked output or work item currently relevant in that stage.
- **Commitment** — a specific node type that emerges in Stage 3B and can progress through Stage 4–7.
- **Rail** — derivation, progression, or dependency relationship.
- **Status** — current progress/readiness signal for a node.

Not every node is a Commitment. Stage 1 contains value-loop outputs, Stage 2 actor outcomes, Stage 3A representative scenarios, and Stage 3B delivery commitments.

### Status legend

| Signal | Meaning |
| --- | --- |
| ✅ COMPLETE | Stage/output sufficiently complete for progression |
| ▶ ACTIVE | Currently being worked through |
| ◆ ELIGIBLE | Commitment may proceed to Stage 4–7 |
| ⛔ BLOCKED | Material semantic/product blocker remains |
| ⏳ DEPENDENCY WAIT | Product claim depends on another real commitment/capability |
| ○ NOT STARTED | No active work in this stage yet |
| ◌ PROBE | Cross-cutting / exceptional probe, not a core delivery commitment |

### Rail legend

```text
────────▶  derivation / normal progression
- - - - ▶  material product dependency
·······▶  feedback / reopen when materially new evidence appears
```

## 2. Current control tower

```mermaid
flowchart LR

  subgraph S1["STAGE 1 — Business / Value Loop"]
    direction TB
    S1A["Whole-product value loop<br/>✅ COMPLETE"]
    S1B["Org / Campaign thread<br/>✅ COMPLETE"]
    S1C["Person thread<br/>✅ COMPLETE"]
    S1D["Trust / Evidence thread<br/>✅ COMPLETE"]
    S1A --> S1B
    S1A --> S1C
    S1A --> S1D
  end

  subgraph S2["STAGE 2 — Actor Outcomes"]
    direction TB
    S2A["Public Visitor<br/>✅ COMPLETE"]
    S2B["Donatur<br/>✅ COMPLETE"]
    S2C["Owner / Staff<br/>✅ COMPLETE"]
    S2D["Reviewer<br/>✅ COMPLETE"]
    S2E["Platform Operator<br/>✅ COMPLETE"]
  end

  subgraph S3["STAGE 3 — Scenario + Commitment Frontier"]
    direction TB

    subgraph S3A["3A — Representative Scenarios"]
      direction TB
      R1["R1 Organization legitimacy + review<br/>✅ COMPLETE"]
      R2["R2 Campaign curation + actual state<br/>✅ COMPLETE"]
      R3["R3 Public decision + donation<br/>✅ COMPLETE"]
      R4["R4 Continuing accountability<br/>✅ COMPLETE"]
      X1["X1 Operator intervention<br/>◌ PROBE"]
    end

    subgraph S3B["3B — Commitment Frontier"]
      direction TB
      C1["C1 Legitimate Organization representation<br/>◆ ELIGIBLE"]
      C2["C2 Bounded Organization review<br/>⛔ SEMANTIC BLOCKER"]
      C3["C3 Real Campaign proposition<br/>⏳ DEPENDENCY WAIT: C1"]
    end
  end

  subgraph S4["STAGE 4 — Interaction Exploration"]
    direction TB
    S4A["○ No active commitment yet"]
  end

  subgraph S5["STAGE 5 — Confirmed Product Behavior"]
    direction TB
    S5A["○ No active commitment yet"]
  end

  subgraph S6["STAGE 6 — Requirements"]
    direction TB
    S6A["○ No active commitment yet"]
  end

  subgraph S7["STAGE 7 — Durable Handoff"]
    direction TB
    S7A["○ No active commitment yet"]
  end

  S1B --> S2C
  S1C --> S2A
  S1C --> S2B
  S1D --> S2D
  S1D --> S2E

  S2C --> R1
  S2C --> R2
  S2D --> R1
  S2D --> R2
  S2A --> R3
  S2B --> R3
  S2B --> R4
  S2C --> R4
  S2D --> R4
  S2E --> X1

  R1 --> C1
  R1 --> C2
  R2 --> C3

  C1 -. product dependency .-> C3
  C1 -->|eligible progression rail| S4A

  S4A --> S5A --> S6A --> S7A

  classDef complete fill:#dcfce7,stroke:#15803d,color:#14532d;
  classDef eligible fill:#dbeafe,stroke:#1d4ed8,color:#1e3a8a;
  classDef blocked fill:#fee2e2,stroke:#b91c1c,color:#7f1d1d;
  classDef wait fill:#fef3c7,stroke:#b45309,color:#78350f;
  classDef probe fill:#f3f4f6,stroke:#6b7280,color:#374151;
  classDef idle fill:#f9fafb,stroke:#9ca3af,color:#6b7280;

  class S1A,S1B,S1C,S1D,S2A,S2B,S2C,S2D,S2E,R1,R2,R3,R4 complete;
  class C1 eligible;
  class C2 blocked;
  class C3 wait;
  class X1 probe;
  class S4A,S5A,S6A,S7A idle;
```

## 3. Current commitment frontier

| Commitment | Current station | Status | Current reason |
| --- | --- | --- | --- |
| **C1 — Legitimate Organization representation** | Stage 3B | ◆ ELIGIBLE | Meaningful, semantically ready enough, dependency-closed enough, truthful, and bounded |
| **C2 — Bounded Organization review** | Stage 3B | ⛔ BLOCKED | Positive meaning / legitimate Organization-review outcomes remain unresolved |
| **C3 — Real Campaign proposition** | Stage 3B | ⏳ DEPENDENCY WAIT | Requires C1 to become real; later C3-specific authority semantics should be resolved when material |

**Progression rule:** an ELIGIBLE commitment may enter Stage 4 directly. Do not resolve unrelated blocked future commitments first.

If several independent commitments become ELIGIBLE at the same time, compare only that eligible frontier. Do not force a full-product sequence.

## 4. Owning artifacts

| Dashboard area | Owning artifact |
| --- | --- |
| Product direction | [Product Intent](./product-intent.md) |
| Stage 1 | [Pilot #3 — Business / Value Loop](./pilot-3-business-value-loop.md) |
| Stage 2 | [Pilot #3 — Actor Outcomes & Responsibilities](./pilot-3-actor-outcomes.md) |
| Stage 3 method | [Pilot #3 — Stage 3 Working Model](./pilot-3-stage-3-working-model.md) |
| Stage 3A | [Pilot #3 — Representative Scenarios](./pilot-3-stage-3a-representative-scenarios.md) |
| Stage 3B | [Pilot #3 — Commitment Sequencing](./pilot-3-stage-3b-commitment-sequencing.md) |

## 5. Dashboard update discipline

Update this dashboard when a meaningful progress state changes, for example:

- a stage/output becomes COMPLETE;
- a commitment becomes ELIGIBLE;
- a commitment becomes BLOCKED or unblocked;
- a dependency becomes real / satisfied;
- a commitment enters Stage 4, 5, 6, or 7;
- materially new evidence reopens an earlier working decision.

Do **not** use this dashboard to create new Product Truth.

The correct flow is:

```text
owning discussion / artifact changes
→ decision becomes durable there
→ Control Tower is updated to reflect it
```

The dashboard is allowed to be compact and operational. Detailed decision reasoning belongs in the linked owning artifact.
