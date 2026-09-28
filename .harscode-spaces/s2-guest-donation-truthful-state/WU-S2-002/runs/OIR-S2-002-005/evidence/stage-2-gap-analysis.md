# Stage 2 — O1 Amount Contract Evidence

## Provenance

- Phase / Stage: Exploration / Stage 2 — Gap Analysis
- Work Unit / Run: `WU-S2-002` / `OIR-S2-002-005`
- Role / specialization: Explorer / O1 amount-contract evidence and owner-resolution facilitation
- Participant: `P-S2-002-OIR-005-1` (Run-local identity)
- Created: 2026-09-28
- Selected model / reasoning: `gpt-6-luna` / `high` per Invocation; active runtime model is not independently exposed
- Session: not exposed
- Target base revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535` (Invocation; current working-tree decisions are effective inputs)
- Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30` (Invocation)
- Profile pin: `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32` — match
- Current-effective input pins rechecked at dispatch:
  - Authority Map: `dd8cd32f65979f766034f6d086018e31ff4c6d584961ef1e296324c47c6eb06c` — match
  - Approved Techplan `TP-S2-002-007`: `441577782f7600fd5507d228c4dd6491009c32d6e8b81607b28a38a5ebb1a1a3` — match
  - `OIR-S2-002-001/resolution-brief.md`: `3d2b22f6fca9bd23024eca41cda79417bcc5cf929513223867d73ddcc6440782` — match
  - `OIR-S2-002-004/design-review-brief.md`: `d9713e35324a0027814bc112237865cbeb5d86e1c3651e590aaca57f61ba29c4` — match
- Verification: targeted source inspection and pinned-input SHA-256 comparisons only. No tests, API validator, migration, runtime behavior, or money implementation correctness was verified.

## Area 1 — Active Product/MVP amount requirements

### Current state and governing requirement

Product/MVP is explicit for amount semantics but intentionally not for authored wire/storage detail:

- `docs/product/mvp-scope.md` §5 Stage B (lines 97–105) requires whole-Rupiah IDR input, minimum Rp5.000, Rp1 increments, and explicitly accepts Rp5.001. It requires exact decimal representation for stored and calculated monetary values, including derived values, while leaving tax rules and derived-value rounding out of scope until separately defined.
- `docs/product/mvp-delivery-slices.md` §5 repeats the same Slice 2 amount direction and identifies it as the current delivery boundary.
- Approved Techplan `TP-S2-002-007` §4 Q2 / §5 D3 and §13 O1 translate that boundary: request/response representation and storage precision/scale remain to be specified by the Donation/API owner; do not invent derived precision/rounding or tax behavior. Product Authority is needed if the choice changes semantics.
- Earlier Human direction is recorded in `OIR-S2-002-001/resolution-brief.md` O1: IDR, integer Rupiah amount, minimum Rp5.000, Rp5.001 valid, and exact-decimal money without `float64`. It leaves authored representation and derived precision unresolved.

The current Techplan also states that threshold/funding contract detail remains in Slice 2, with successful funding increments coupled to Donation success. This implies exact calculation of collected funding from submitted amounts; because submitted donation amounts are whole Rupiah, the product evidence does not establish a fractional monetary result or a new rounding rule for that sum. Tax rules are expressly outside this slice. Existing Campaign funding progress percentage is a different derived ratio and belongs to the existing Slice 1 projection evidence below, not a Product decision about fractional Donation amounts.

### Gap

The product boundaries are known, but no current Product source chooses the request/response encoding, input schema constraints enforcing integer Rupiah, or precision/scale for stored Donation money. No current Product evidence calls for tax calculations in Slice 2. Exact handling of the funding sum is required; a fractional money rounding decision is not established by the amount requirements inspected.

### Sniffing lenses

- **Risk:** A representation that accepts fractional Rupiah or silently rounds request amounts would contradict the explicit integer-Rupiah promise. A representation too narrow for supported stored/calculated values could truncate/corrupt exact money; `float64` is prohibited.
- **Edge cases:** Rp4.999 is below minimum; Rp5.000 is the lower valid boundary; Rp5.001 is valid and proves the increment is Rp1 rather than a larger denomination. Fractional input such as 5000.5 is outside the stated input domain. Zero, negative, and oversized values are not fully specified by the cited Product amount sentence; do not infer a maximum from historical schema examples.
- **Miscontext:** “Exact decimal for stored/calculated money” does not make donor input fractional, and “derived values” does not authorize tax behavior. Conversely, integer input alone does not permit floating-point funding aggregation.
- **Misleading signals:** Historical two-decimal examples or a decimal database type may look like settled Slice 2 semantics; the active Product source sets neither wire encoding nor scale.
- **Inconsistency:** Older detail says decimal-string input while current Product says integer Rupiah input. These can coexist only if the string representation is constrained to integer Rupiah semantics; current authorities do not state that mapping.

### Anchors

- `docs/product/mvp-scope.md` §5 Stage B, lines 97–105 — active amount rule and derived-money boundary.
- `docs/product/mvp-delivery-slices.md` §5 — Slice 2 requirements and truthfulness boundary.
- `TP-S2-002-007/techplan.md` §4 Q2, §5 D3, §8 Interface Contract, §13 O1 — open wire/storage detail and owner routing.
- `OIR-S2-002-001/resolution-brief.md` O1 — Human direction and still-open technical detail.

## Area 2 — Historical Donation feature spec and split OpenAPI

### Current state

The historical feature spec, `docs/spec/5-donation/features/01-submit-donation-settlement.md` §Request (lines 42–51), calls `amount` a required “decimal string” with `≥ 5000`. Its behavior section validates only `amount ≥ 5000` (lines 53–60). `docs/spec/5-donation/invariants.md` INV-donation-01 (lines 31–37) describes application and database minimum checks, but does not state whole-Rupiah/integer validation.

Authored `api/openapi/donation.yaml` has these relevant shapes:

- POST `/campaigns/{campaignId}/donations` references `SubmitDonationRequest` (lines 33–51). `amount` is a `type: string`, with description “Decimal string, minimum 5000.” and example `"50000.00"` (lines 194–201). There is no numeric minimum, pattern, or integer-format constraint in that schema property.
- The immediate Donation response and status response use the `Donation` schema; `Donation.amount` is `type: string` with example `"50000.00"` (lines 225–246; status operation lines 77 onward).
- Other amount-bearing historical projections (`DonationListItem`, `ClaimableDonation`, `MyDonation`) likewise use string and `"50000.00"` examples (Donation lines 260–280, 296–320, 336–350). These are historical API shapes, not current consumer authority for the revised Slice 2 contract.
- `api/openapi/common.yaml` has shared generic Problem/ValidationProblem and Pagination schemas, but no shared money/amount component. The domain amount properties are local to `donation.yaml`.

### Requirement and gap

The active requirement is an integer number of IDR Rupiah with minimum 5000 and step 1. The historical feature spec and OpenAPI describe a decimal string and a two-decimal example, without machine-readable whole-Rupiah constraints. Neither establishes whether decimal-string encoding should be retained with integer-only lexical semantics, replaced by a JSON integer, or adapted another way. Their application minimum does not show enforcement of the new integer constraint. The exact response format is similarly unselected by active Slice 2 authority.

### Sniffing lenses

- **Risk:** Treating the old decimal-string example as sufficient validation could admit fractional Rupiah or leave request/response clients inconsistent. A wire-type change could affect API consumers and generated clients; actual consumer impact is not determined here.
- **Edge cases:** Check 4999/5000/5001 and fractional text `5000.50`; the historical specification only distinguishes below-minimum from accepted values and does not explicitly cover the whole-Rupiah boundary.
- **Miscontext:** Presence of `minimum 5000` in descriptive text is not a machine-readable numeric bound on a string schema. The old spec says decimal string but does not establish scale, canonical form, or whether trailing `.00` is mandatory.
- **Misleading signals:** `"50000.00"` examples in request and response shapes look like a two-decimal contract, but are examples without an associated precision/scale rule. Generic validation responses in `common.yaml` do not establish which values the amount schema admits.
- **Inconsistency:** Historical feature spec/OpenAPI state decimal string; current Product says input is whole Rupiah and the Approved Techplan leaves wire format open. Historical INV-donation-01 mentions only the minimum. Preserve all three as evidence; current Product/Techplan authority governs reconciliation.

### Anchors

- `docs/spec/5-donation/features/01-submit-donation-settlement.md` §Request / §Behavior, lines 42–60 — historical input representation and minimum-only validation.
- `docs/spec/5-donation/invariants.md` INV-donation-01, lines 31–37 — historical minimum invariant.
- `api/openapi/donation.yaml` POST submit operation lines 33–64; `SubmitDonationRequest` lines 194–223; `Donation` lines 225–260; other projection fields lines 260–360 — authored request/response anchors.
- `api/openapi/common.yaml` `components.schemas` lines 117–173 — shared generic schema inventory; no amount schema.
- `api/README.md` — split domain source is authored API source; `api/openapi.yaml` is generated aggregate and not the default editing source.

## Area 3 — Existing exact-decimal and persistence conventions

### Current state

There is no `backend/internal/domain/donation/` package or Donation migration in the current backend tree. Current backend money evidence is the Slice 1 Campaign funding projection:

- `backend/migrations/000011_create_public_campaigns.up.sql` lines 21–28 defines `target_amount` and `collected_amount` as PostgreSQL `NUMERIC(19,2)`, with paired-null and non-negative checks. It is a Campaign schema, not a Donation storage decision.
- `backend/internal/domain/campaign/service.go` `mapFunding` (lines 100–126) parses persisted amount strings using `shopspring/decimal.NewFromString`, formats target/collected output via `StringFixed(2)`, and computes percentage via decimal arithmetic, then `Round(2).StringFixed(2)`.
- `backend/internal/domain/campaign/service_test.go` lines 36–49 exercises a collected value `125.505` and expects `125.51` output (half-up behavior); this is Campaign funding projection evidence, not an approved Donation rounding rule.
- `backend/cmd/campaign-seed/main.go` `validateAmounts` (lines 198–209) parses with `decimal.NewFromString`, rejects negative values, and constrains seed amounts to at most two decimal places using `Equal(Truncate(2))`.
- `backend/go.mod` includes `github.com/shopspring/decimal v1.4.0`.
- Harscode `best-practices/go/decimal-and-money.md` lines 5–28 requires decimal end-to-end, no `float64`, explicit consistent rounding, decimal `.Equal()` comparisons, and database precision/scale consistent with application precision.

### Requirement and gap

The active money principle aligns with the existing decimal library convention. The current Campaign persistence example provides a precedent of `NUMERIC(19,2)`, but that scale is for Slice 1 Campaign target/collected values and is not authority for Donation. There is no Donation table/model/repository yet to establish a live Donation storage convention. The current backend has no amount input parser to show how whole-Rupiah semantics should be enforced.

For current Slice 2, exact arithmetic is needed for successful donations to contribute to collected funding (per Techplan §8 and the retained atomic funding invariant). The evidence does not show a need for tax or another fractional monetary derivation in this slice. Campaign's existing two-place formatting/percentage rounding concerns a separate projection and cannot by itself determine Donation storage scale or API encoding.

### Sniffing lenses

- **Risk:** Reusing `NUMERIC(19,2)` without checking the Donation application range/representation can impose unsupported bounds or semantics. Any conversion through binary float risks loss and violates the money rule. The existing Campaign rounding behavior must not leak into Donation amount semantics by analogy.
- **Edge cases:** Amount at 5000/5001, parsing fractional values, accumulated funding across many integer donations, numeric capacity boundaries, and any future fractional calculation need explicit behavior/evidence. Current sources do not establish the Donation column range or a future fractional precision requirement.
- **Miscontext:** `NUMERIC(19,2)` and `Round(2)` prove a current Campaign convention only. A test using a three-decimal in-memory value does not establish that the database accepts it; the migration has scale 2 and the mapper then formats it to 2 places.
- **Misleading signals:** Existing dependency on `shopspring/decimal` is a reusable technical convention, not proof that Donation persistence/API path already exists. The Campaign table's decimal columns are not a Donation schema.
- **Inconsistency:** Current Product promises whole-Rupiah donation inputs; historical Donation API examples display cents; the existing Campaign data model stores two decimal places. These belong to different authority/evidence layers and are not interchangeable. No direct current Donation code contradicts Product because that package/schema is absent.

### Anchors

- `backend/internal/domain/campaign/service.go` `mapFunding`, lines 100–126 — decimal parsing, output formatting, and percentage rounding.
- `backend/internal/domain/campaign/entity.go` `DetailRecord` and `Funding`, lines 35–101 — Campaign projection uses amount strings.
- `backend/migrations/000011_create_public_campaigns.up.sql`, lines 21–28 — current Campaign `NUMERIC(19,2)` columns and checks.
- `backend/cmd/campaign-seed/main.go` `validateAmounts`, lines 198–209 — operator seed precision restriction.
- `backend/internal/domain/campaign/service_test.go` lines 36–49 — test example demonstrating two-place formatted projection.
- `backend/go.mod` decimal dependency; `../harscode-workspace/best-practices/go/decimal-and-money.md` lines 5–28 — applicable technical money rule.
- Absence coordinate: `backend/internal/domain/` current file inventory contains account and campaign only; no `donation/` package or Donation migration exists in current `backend/migrations/` inventory.

## Cross-area observations for Stage 3

- Decision authority is assigned: Anhar Solehudin owns Donation delivery/domain amount wire/storage detail and API/contract for current Slice 2 (`.harscode-spaces/authority-map.md` rows at lines 7–11; scope limited at lines 24–28). Product Authority is needed only if a proposed choice changes Product semantics.
- The minimum evidence-backed owner question is the representation of whole-Rupiah amount in authored request/response and how the amount is constrained to integer Rupiah. Storage precision/scale remains part of the same O1 item, but current Donation persistence does not exist; the Campaign precedent must not be treated as a decision.
- The sources inspected confirm no Slice 2 tax requirement. They require exact calculation of collected funding from donations, but do not evidence fractional Donation-derived money or prescribe a new rounding decision now. Existing Campaign percentage rounding is not Donation money rounding authority.
- This is evidence only. No solution option is selected here; options, trade-offs, and the owner-facing decision request belong to Stage 3 after Human confirms the gap analysis.

