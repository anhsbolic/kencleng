# Kencleng — Current Authority Map (Pilot #2)

> Status: COMPLETE for the six Slice-2 authority areas material to `WU-S2-002` and the explicitly attributed project-wide currency-standard area; later areas remain event-driven.
> Scope: current `WU-S2-002` authority areas plus the named project-wide shared currency standard. Paths below are relative to the Kencleng repository root.
> This map records current ownership; it does not replace canonical Product/Design/spec/API authority or historical Decision provenance.

| Authority area | Decision scope at current frontier | Named current owner | Effective context | Evidence |
|---|---|---|---|---|
| Donation delivery/domain | Amount wire/storage detail, simulator timing/scenario, guest notification delivery contract, retry-record detail, settlement/funding invariant translation | Anhar Solehudin | Current Slice 2 only; effective 2026-09-28 | Human owner attribution, 2026-09-28; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` §13 |
| Campaign delivery/domain | Eligibility, threshold crossing, accepted pending donations, concurrent close/settlement contract ordering | Anhar Solehudin | Current Slice 2 only; effective 2026-09-28 | Human owner attribution, 2026-09-28; same Techplan §13, Campaign/Donation ordering item |
| API/contract | Authored request/response and credential shape, idempotency/status parity, conditional consumer/distribution decision | Anhar Solehudin | Current Slice 2 only; effective 2026-09-28 | Human owner attribution, 2026-09-28; same Techplan §13 O1/O3–O5/O8/O9 |
| Security/PII | Email verification/retention windows, status credential exposure, abuse/anti-enumeration controls, residual security/privacy risk decision | Anhar Solehudin | Current Slice 2 only; effective 2026-09-28 | Human owner attribution, 2026-09-28; same Techplan §13 O3–O5 |
| Product Design | Truthful state/method labels, notification and recovery presentation, material interaction/visual decisions | Anhar Solehudin | Current Slice 2 only; effective 2026-09-28 | Human owner attribution, 2026-09-28; same Techplan §13 O7; `docs/ui-ux/README.md` |
| Product/MVP — conditional O3 notification meaning | Whether an independent pending-email retention cap may end terminal-email eligibility before Donation reaches terminal state, and any resulting current-Slice-2 Product/MVP wording change | Anhar Solehudin | Current Slice 2 and this material Product/MVP decision only; effective 2026-09-28 | Explicit Human owner attribution, 2026-09-28; `TP-S2-002-009/o2-delivery-proposal.md` O3 retention consequences; `docs/product/mvp-scope.md` Stage B/C |
| Project-wide Shared Currency Standard | Cross-feature monetary representation: currency identifier/code, wire encoding, exact-decimal representation, precision/range, persistence/storage scale, and consistency across domains, APIs, and storage. Excludes feature business policies (per-feature amount limits, tax, FX/exchange rates, pricing) unless separately attributed. | Anhar Solehudin | Project-wide across Kencleng; effective 2026-09-30; reassignable | Explicit Human attribution and direction approval, 2026-09-30; `docs/project/kencleng-monetary-data-standard.md`; `.harscode-spaces/s2-guest-donation-truthful-state/events.md`. |

## Conditional Product/MVP area now mapped

Product/MVP policy for Slice 2 is already approved in `docs/product/mvp-scope.md` and `docs/product/mvp-delivery-slices.md`. The Product/MVP row above was added only after the independent pending-email cap route made a possible terminal-notification exception material. The current-slice exception question is now resolved in Events; this row does not establish permanent project-wide Product ownership.

## Known decision provenance, distinct from this map

`.harscode-spaces/s2-guest-donation-truthful-state/events.md` records Anhar Solehudin as the Human approver of current-effective Techplan `TP-S2-002-007` on 2026-09-27. That approval did not itself establish ownership of the six Slice-2 areas; separate explicit Human attributions established those scopes on 2026-09-28 and 2026-09-30. The project-wide currency-standard attribution is a separate explicit Human decision on 2026-09-30. Prior Decision provenance remains distinct from current ownership.

## Use and update rule

At an authority decision boundary, confirm the area has a named current owner and applicable scope/effective context. Record a new owner attribution from explicit Human/owner evidence; do not infer it from previous Run participation, generic approval, or this table. Prior Decisions retain their original provenance if ownership later changes. `UNMAPPED` routes to Authority Sync; a known owner with a bounded unresolved question routes to that owner or to useful specialist analysis first.

The six Product, Design, Donation, Campaign, API, and Security/PII attributions above remain limited to current Slice 2; the Product/MVP row is further limited to the identified O3 notification-meaning question, which Human resolved by preserving the terminal-notice obligation. The currency-standard owner attribution is explicitly project-wide and reassignable; the separate direction decision is recorded in the linked standard. It does not set the deferred currency-specific precision/range/fraction/storage values or extend to feature business policies excluded in the row. None of these attributions decide unrelated request/response shapes, numeric terminal/retention bounds, token or anti-enumeration controls, simulator architecture or timeout semantics, or residual-risk acceptance.
