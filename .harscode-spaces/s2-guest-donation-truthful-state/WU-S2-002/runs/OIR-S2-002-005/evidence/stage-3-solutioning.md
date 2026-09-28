# Stage 3 — O1 Owner Resolution and Routing

## Provenance

- Phase / Stage: Exploration / Stage 3 — Solutioning
- Work Unit / Run: `WU-S2-002` / `OIR-S2-002-005`
- Role / specialization: Explorer / O1 amount-contract evidence and owner-resolution facilitation
- Participant: `P-S2-002-OIR-005-1`
- Created: 2026-09-28
- Model / reasoning: `gpt-6-luna` / `high` per Invocation; active runtime model not independently exposed
- Session: not exposed
- Target revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535`; current working-tree decisions are effective inputs
- Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`
- Profile pin: `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32` — match

## Human direction recorded

On 2026-09-28, Anhar Solehudin stated that the currency representation should become a standard usable across currencies, tables, and features that need money. He also stated that user input in Kencleng should be whole Rupiah because the current application is IDR-specific. After the Explorer explained that the global standard exceeds the current O1/Slice 2 Run boundary and needs an Orchestrator route, Human agreed and instructed continuation.

This records a direction and scope signal. It does **not** select a concrete wire type, database type/scale, currency metadata model, or authorize this Run to establish a project-wide architecture standard. The current Authority Map attributes Anhar's Donation/API authority to current Slice 2 only and explicitly does not imply permanent project-wide authority (`.harscode-spaces/authority-map.md` lines 7–11, 24–28). The Invocation classifies material expansion beyond O1 as `ORCHESTRATOR_DECISION`.

## Decision surface and options

Current Slice 2 remains IDR-only. Product requires the donor to enter a whole number of Rupiah, min Rp5.000, Rp1 increment, and explicitly allows Rp5.001. A future cross-currency standard must also preserve exact money and identify currency; “integer amount” by itself is ambiguous across currencies with different fractional units.

### Option A — major-unit decimal string plus currency code

Illustrative API values:

```json
{"amount":"5001","currency":"IDR"}
{"amount":"12.34","currency":"USD"}
```

Persist exact decimal major-unit amount with an explicit ISO-style currency code, for example a donation row with `amount` and `currency_code`; a campaign with one currency shared by `target_amount` and `collected_amount` can hold one `currency_code` for that row. Parse/compute with the established decimal library. Validate the allowed fraction digits and range for the currency, and use a database precision/scale selected from the supported currency/range requirements.

- **Benefits:** preserves whole-Rupiah IDR semantics (`"5001"` has no fractional digits), expresses major units clearly at the API boundary, avoids float conversion, and supports fractional major units where a supported currency permits them.
- **Costs / unresolved details:** exact accepted currency set, maximum amount, database precision/scale, per-currency fraction-digit authority, and historic representation policy must be specified. A fixed `NUMERIC(19,2)` is not established as globally adequate by current evidence. Database checks and application parsing/validation must agree.

### Option B — minor-unit integer plus currency code

Illustrative API values:

```json
{"amount_minor":5001,"currency":"IDR"}
{"amount_minor":1234,"currency":"USD"}
```

Store integer minor units plus currency code, with currency metadata defining the exponent/unit mapping. The user-facing form still accepts currency-formatted major units; the client converts to minor units before submission.

- **Benefits:** each stored amount is an exact integer count of units and arithmetic avoids fractional database values.
- **Costs / unresolved details:** the API value is less directly recognizable as major units; conversion/validation relies on authoritative currency metadata; currency scale changes and historic interpretation need a stable/versioned rule; maximum integer/range and client-safe encoding must be set. It does not mean all users type whole-number major units (for example, a supported USD entry with cents still needs a fractional major-unit UI value).

### Option C — keep the active Slice 2 API IDR-specific; defer the global pattern

The Slice 2 API can encode whole IDR as an integer-valued amount while an independently owned project-wide standard is designed. It avoids implying multi-currency runtime support now, but does not itself satisfy the expressed requirement for one shared cross-feature standard and leaves the local Donation persistence choice dependent on the global design route.

## Explorer recommendation

Use Option A as the starting candidate for the separate global-standard decision: major-unit decimal strings plus explicit currency code at API boundaries and exact-decimal database storage, with the supported currency set, scale/range, and per-currency validation rules explicitly selected before contract or migration authoring. This is a recommendation for Orchestrator/assigned authority review, not a decision or `CONTRACT_READY` result. It preserves the current product distinction: IDR input is whole Rupiah, while a future currency-enabled feature may accept the precision defined for that currency.

Current evidence does not establish a single global precision/range convention. Therefore the project-wide standard cannot be finalized from existing Campaign `NUMERIC(19,2)` evidence or the historical Donation `"50000.00"` example. No tax or tax-rounding behavior is implied.

## Explicitly unresolved

- Orchestrator decision on whether to create/expand a cross-feature Work Unit for a global money/currency standard, and which durable authority owns it.
- Supported currency set and the authoritative per-currency fraction-digit metadata/source.
- Whether the recommended major-unit decimal-string representation is accepted, or Option B is preferred.
- Database precision/scale and maximum supported amount after the currency/range boundary is approved.
- Treatment of currency conversion, if it is ever in scope; no conversion behavior is requested or assumed here.
- Slice 2 Donation API request/response and persistence remain unselected until the global-standard route is reconciled.
