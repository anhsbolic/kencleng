# Kencleng — Current Authority Map (Pilot #2)

> Status: COMPLETE for the six authority areas now material to Slice 2 `WU-S2-002`; later areas remain event-driven.
> Scope: only authority areas material to the current `WU-S2-002` frontier. Paths below are relative to the Kencleng repository root.
> This map records current ownership; it does not replace canonical Product/Design/spec/API authority or historical Decision provenance.

| Authority area | Decision scope at current frontier | Named current owner | Effective context | Evidence |
|---|---|---|---|---|
| Donation delivery/domain | Amount wire/storage detail, simulator timing/scenario, guest notification delivery contract, retry-record detail, settlement/funding invariant translation | Anhar Solehudin | Current Slice 2 only; effective 2026-09-28 | Human owner attribution, 2026-09-28; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` §13 |
| Campaign delivery/domain | Eligibility, threshold crossing, accepted pending donations, concurrent close/settlement contract ordering | Anhar Solehudin | Current Slice 2 only; effective 2026-09-28 | Human owner attribution, 2026-09-28; same Techplan §13, Campaign/Donation ordering item |
| API/contract | Authored request/response and credential shape, idempotency/status parity, conditional consumer/distribution decision | Anhar Solehudin | Current Slice 2 only; effective 2026-09-28 | Human owner attribution, 2026-09-28; same Techplan §13 O1/O3–O5/O8/O9 |
| Security/PII | Email verification/retention windows, status credential exposure, abuse/anti-enumeration controls, residual security/privacy risk decision | Anhar Solehudin | Current Slice 2 only; effective 2026-09-28 | Human owner attribution, 2026-09-28; same Techplan §13 O3–O5 |
| Product Design | Truthful state/method labels, notification and recovery presentation, material interaction/visual decisions | Anhar Solehudin | Current Slice 2 only; effective 2026-09-28 | Human owner attribution, 2026-09-28; same Techplan §13 O7; `docs/ui-ux/README.md` |
| Product/MVP — conditional O3 notification meaning | Whether an independent pending-email retention cap may end terminal-email eligibility before Donation reaches terminal state, and any resulting current-Slice-2 Product/MVP wording change | Anhar Solehudin | Current Slice 2 and this material Product/MVP decision only; effective 2026-09-28 | Explicit Human owner attribution, 2026-09-28; `TP-S2-002-009/o2-delivery-proposal.md` O3 retention consequences; `docs/product/mvp-scope.md` Stage B/C |

## Conditional Product/MVP area now mapped

Product/MVP policy for Slice 2 is already approved in `docs/product/mvp-scope.md` and `docs/product/mvp-delivery-slices.md`. The Product/MVP row above was added only after the independent pending-email cap route made a possible terminal-notification exception material. It does not reopen settled O1–O6/O9 policy by itself, decide the exception, or establish permanent project-wide Product ownership.

## Known decision provenance, distinct from this map

`.harscode-spaces/s2-guest-donation-truthful-state/events.md` records Anhar Solehudin as the Human approver of current-effective Techplan `TP-S2-002-007` on 2026-09-27. That approval did not itself establish ownership of the five areas; separate explicit Human attributions on 2026-09-28 established each scope above. Prior Decision provenance remains distinct from current ownership.

## Use and update rule

At an authority decision boundary, confirm the area has a named current owner and applicable scope/effective context. Record a new owner attribution from explicit Human/owner evidence; do not infer it from previous Run participation, generic approval, or this table. Prior Decisions retain their original provenance if ownership later changes. `UNMAPPED` routes to Authority Sync; a known owner with a bounded unresolved question routes to that owner or to useful specialist analysis first.

All six named-owner attributions above apply to current Slice 2 only; the Product/MVP row is further limited to the identified O3 notification-meaning question. They do **not** choose request/response shapes, set verification/retention windows, select token or anti-enumeration controls, decide simulator timing or Design expression, decide the new notification exception, accept residual risk, or imply permanent project-wide ownership of any area.
