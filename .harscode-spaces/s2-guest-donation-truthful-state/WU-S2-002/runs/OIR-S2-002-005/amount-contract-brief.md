# O1 Amount Contract Brief — `OIR-S2-002-005`

## Provenance

- Work Unit / Run: `WU-S2-002` / `OIR-S2-002-005`
- Phase / Role: Exploration / Explorer — O1 amount-contract evidence and owner-resolution facilitation
- Participant: `P-S2-002-OIR-005-1`
- Created: 2026-09-28
- Model / reasoning: `gpt-6-luna` / `high` per Invocation; runtime model not independently exposed
- Session: not exposed
- Target revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535` (Invocation; current working-tree decisions effective)
- Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`
- Evidence: [Stage 2 gap analysis](evidence/stage-2-gap-analysis.md), [Stage 3 solutioning](evidence/stage-3-solutioning.md)

## Outcome

**O1 remains partially resolved and requires Orchestrator routing.** The current Slice 2 Product amount rules are settled. Human direction is that a future currency representation should become a shared standard across currencies/tables/features, while the active Kencleng user input remains whole Rupiah. The evidence and discussion do not settle the API wire encoding, global database precision/scale, supported currencies, or durable owner for a project-wide standard.

## Active requirements retained

- Current Slice 2 input is IDR whole Rupiah, minimum Rp5.000, Rp1 increments; Rp5.001 is valid.
- Stored/calculated monetary values use exact decimal representation and no `float64`.
- Successful Donation funding contribution must be calculated exactly; no fractional Donation-derived amount or tax rule is established for Slice 2.
- Tax rules and derived-value rounding are out of Slice 2 unless Product authority separately changes scope.
- Historical specs/contracts/code are evidence and do not override current Product/MVP authority.

Sources: `docs/product/mvp-scope.md` §5 Stage B; `docs/product/mvp-delivery-slices.md` §5; Approved `TP-S2-002-007` §§4, 5, 8, 13; root `AGENTS.md` §2; `../harscode-workspace/best-practices/go/decimal-and-money.md`.

## Historical detail disposition

| Detail | Finding | Disposition |
|---|---|---|
| Product whole-IDR/minimum Rp5.000/step Rp1/Rp5.001 valid | Current approved boundary | **KEEP** |
| Exact decimal / no `float64` | Current money rule and applicable technical guidance | **KEEP** |
| Historical Donation API string example `"50000.00"` and feature-spec “decimal string” | Historical shape conflicts with whole-Rupiah semantics unless constrained; not machine-readable integer validation | **ADAPT** during authorized contract reconciliation; do not copy as settled |
| Historical `amount ≥ 5000` only | Does not state whole-Rupiah constraint | **ADAPT** when active contract is authored |
| Campaign `NUMERIC(19,2)` target/collected columns | Existing Slice 1 Campaign implementation precedent only; not a global currency standard or Donation storage decision | **KEEP** as current Campaign evidence; **DEFER** reuse as a global standard |
| Campaign `Round(2)` percentage projection | Slice 1 funding ratio behavior; not an O1 Donation amount rounding rule | **KEEP** in its existing context; **DEFER** as authority for Donation derived money |
| Tax calculation/rounding | Not required by current Slice 2 | **DEFER** |

These labels classify historical detail for the active reconciliation; they do not authorize edits to Product, spec, OpenAPI, or implementation.

## Direction and options

Anhar Solehudin stated on 2026-09-28 that the currency representation should be a standard usable across currencies, tables, and features that need it, while current user input is whole Rupiah. After the scope boundary was explained, Human agreed to proceed with Orchestrator routing.

Explorer recommends that the separate standard decision start with major-unit decimal strings plus explicit currency code in APIs and exact-decimal database storage, validated against an explicitly supported currency set and per-currency fraction rules. Illustrations only:

```json
{"amount":"5001","currency":"IDR"}
{"amount":"12.34","currency":"USD"}
```

The alternative of integer minor units plus currency metadata remains viable, with additional conversion, exponent/versioning, and range decisions. The direction is **not an approved project-wide standard**. The current Authority Map scopes Anhar's Donation/API assignment to Slice 2 (`.harscode-spaces/authority-map.md` lines 7–11, 24–28), while the Invocation routes material expansion beyond O1 to `ORCHESTRATOR_DECISION`.

## Findings / Decisions / Blockers

### Finding F1 — Historical contract does not encode whole-Rupiah validation

The historical Donation spec says decimal string and minimum 5000; OpenAPI uses `type: string` and example `"50000.00"` without machine-readable whole-Rupiah restrictions. Current Product semantics are narrower. Reconcile only through the authorized future contract route.

### Finding F2 — No live Donation persistence precedent exists

Current backend has no Donation package or Donation migration. Campaign `NUMERIC(19,2)` is not evidence of a Donation schema or globally sufficient scale.

### Decision D1 — Human direction for shared currency standard

**Direction recorded, implementation decision open.** Human wants one currency standard across currencies/tables/features that need money, with current IDR input remaining whole Rupiah. This Run does not establish a permanent authority, global schema, or API contract from that direction.

### Blocker B1 — Orchestrator decision required for scope/authority

The global standard expands beyond current O1/Slice 2 Run scope. Orchestrator must decide whether to create or expand a Work Unit for the shared currency standard and identify its durable decision owner before a project-wide API/database standard is authored. No Orchestrator state was edited by this Run.

## Smallest next authoring route

1. Orchestrator routes a bounded cross-feature money/currency-standard Work Unit and names the authority that can settle shared API/domain/storage policy.
2. That owner decides supported currency scope, amount representation (major decimal or minor units), exact-decimal persistence type/range/scale, currency metadata and historical interpretation, and whether current Slice 2 shares the new standard immediately.
3. Only then, an authorized Planner/Build route reconciles Donation spec and split `api/openapi/donation.yaml` plus implementation scope. Keep tax behavior outside Slice 2 unless Product authority changes it.

## Verification

- Performed: targeted read of active Product/MVP sources, approved Techplan, O1 prior resolution, split Donation spec/OpenAPI and shared components, current backend money conventions, routed decimal best practice; verified pinned input hashes.
- Not performed: tests, API lint/bundle validation, runtime checks, migration validation, persistence proof, generated-client checks, or implementation changes.
- No `CONTRACT_READY`, migration approval, money implementation proof, or residual-risk acceptance is claimed.

