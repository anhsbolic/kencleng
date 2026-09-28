# Kencleng — Current Authority Map (Pilot #2)

> Status: INCOMPLETE — named ownership for currently needed Slice 2 areas remains to be recorded.
> Scope: only authority areas material to the current `WU-S2-002` frontier. Paths below are relative to the Kencleng repository root.
> This map records current ownership; it does not replace canonical Product/Design/spec/API authority or historical Decision provenance.

| Authority area | Decision scope at current frontier | Named current owner | Effective context | Evidence |
|---|---|---|---|---|
| Donation delivery/domain | Amount wire/storage detail, simulator timing/scenario, guest notification delivery contract, retry-record detail, settlement/funding invariant translation | UNMAPPED | Slice 2 / `WU-S2-002` | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` §13 |
| Campaign delivery/domain | Eligibility, threshold crossing, accepted pending donations, concurrent close/settlement contract ordering | UNMAPPED | Slice 2 / `WU-S2-002` | Same Techplan §13, Campaign/Donation ordering item |
| API/contract | Authored request/response and credential shape, idempotency/status parity, conditional consumer/distribution decision | UNMAPPED | Slice 2 / `WU-S2-002` | Same Techplan §13 O1/O3–O5/O8/O9 |
| Security/PII | Email verification/retention windows, status credential exposure, abuse/anti-enumeration controls, residual security/privacy risk decision | UNMAPPED | Slice 2 / `WU-S2-002` | Same Techplan §13 O3–O5 |
| Product Design | Truthful state/method labels, notification and recovery presentation, material interaction/visual decisions | UNMAPPED | Slice 2 / `WU-S2-002` | Same Techplan §13 O7; `docs/ui-ux/README.md` |

## Conditional area

Product/MVP policy for Slice 2 is already approved in `docs/product/mvp-scope.md` and `docs/product/mvp-delivery-slices.md`. Map its named current owner if a new or changed product-semantic decision actually becomes necessary. The current technical/owner gaps do not authorize reopening settled O1–O6/O9 policy.

## Known decision provenance, distinct from this map

`.harscode-spaces/s2-guest-donation-truthful-state/events.md` records Anhar Solehudin as the Human approver of current-effective Techplan `TP-S2-002-007` on 2026-09-27. That scoped approval does not establish his ownership of any UNMAPPED area above. A named person may own multiple areas only when the relevant scopes are attributed explicitly.

## Use and update rule

At an authority decision boundary, confirm the area has a named current owner and applicable scope/effective context. Record a new owner attribution from explicit Human/owner evidence; do not infer it from previous Run participation, generic approval, or this table. Prior Decisions retain their original provenance if ownership later changes. `UNMAPPED` routes to Authority Sync; a known owner with a bounded unresolved question routes to that owner or to useful specialist analysis first.
