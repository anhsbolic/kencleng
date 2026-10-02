## Provenance

- Phase: Independent Techplan Review
- Author: P-S2-003-RV-001-1 (Reviewer)
- Created: 2026-10-01
- Model / reasoning: Invocation configures `gpt-6-luna` / `high`; active runtime values were not independently exposed
- Session: not exposed
- Target revision: `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree
- Workflow revision: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`

## Review findings — WU-S2-003
**Gate:** Complex — the plan has 13 Rules & Validation entries and crosses Campaign–Donation transaction ownership, financial concurrency, bearer credential access, guest PII/email lifecycle, and API anti-enumeration boundaries.
**Sections resolved:** Rules & Validation §4; Decision Log §5; Edge Cases & Risks §7; Interface Contract §8; Architecture / Plan §9; Implementation Details §10; Testing Checklist and Test Focus Pointer §12; Open Items §13.

### Blocking
- **[MATERIAL / interface-data contract] Idempotent retries can be rejected after Campaign closure.** Location: §9 admission transaction, also §8 API contract and R3/R8. The proposed sequence locks Campaign, rechecks eligibility, rejects if closed, and only then “insert-or-resolve[s]” same-key Donation. A retry of an already accepted intent after the Campaign closes therefore reaches the closed rejection before resolving the existing Donation. This contradicts §4 R3 and the accepted `api/openapi/donation.yaml` POST contract (“same `Idempotency-Key` and same payload returns the original donation”); Exploration Stage 2 Area 2 also records that exact behavior. Define the transaction ordering so an existing matching idempotency record returns its original result while a genuinely new key still participates in atomic eligibility ordering. The current plan leaves Build to invent which rule wins.
- **[MATERIAL / monetary contract] Unconstrained Donation values cannot all be reflected by the existing Campaign funding column.** Location: §§8–9 and §11 migration scope; D12; R2/R9. D12 selects unconstrained PostgreSQL `NUMERIC` and explicitly adds no maximum, while the live `backend/migrations/000011_create_public_campaigns.up.sql` defines `campaigns.collected_amount NUMERIC(19,2)`. The accepted `SubmitDonationRequest.amount` schema has no maximum, and the monetary standard explicitly leaves concrete range unresolved. A sufficiently large otherwise-valid whole-IDR Donation can be persisted but overflow Campaign funding during settlement, defeating the requirement to settle an accepted Donation in full and reflect it atomically. The migration plan only adds `max_amount` and `closed_reason`, and expressly avoids adding an amount bound. Resolve the representable-range mismatch through the owning monetary/product authority (for example, compatible Campaign funding capacity or an approved bound), then record the chosen behavior and verification; do not silently infer a cap or claim the existing column preserves all valid values.

### Non-blocking
- None identified.

### Clean
- Rule sources and coverage: R1–R13 map to current Product/MVP, accepted Donation invariants/features, contract, and approved TP-S2-002-015; each has §12 coverage. R3's stated behavior is correctly sourced, but the architecture sequence conflicts with it as noted above.
- Decision fidelity: D1–D14 preserve the Stage 3 solutioning directions and recorded Human planning decisions. The plan keeps the D1 mechanism approval separate from file-specific Tier-0 authorization, and preserves external/runtime/security gates rather than treating them as resolved.
- Diagram validation: no diagram is present; not applicable.
- Open Items lifecycle: Active and Resolved are explicit; resolved entries retain actual resolutions and consequences. Operational controls, residual-risk gates, and the WU-S2-005 producer dependency remain appropriately Active/scoped.
- Technical facts: spot-checked route registration and Campaign read-only repository, current Campaign migration columns/types, Donation amount schema, D1 invariant and threshold semantics against live source. No checked claim selected the public `published` read as transaction authority; D13 is recorded as a planning decision, not runtime evidence.
- Test Focus Pointer: all five surviving sensitive Exploration areas are represented with evidence anchors to the Stage 2 area headings and remain relevant. The entries do not promote ordinary rule-level cases into specialized Testing without a stated concurrency/security rationale.

## Phase handoff
- Completed: independent review gate + review; Complex criteria apply.
- Artifacts: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-001/review-findings.md`
- Human decision: resolve the two material findings in a fresh Planner resolution pass; then review/approve or revise the converged Techplan at the applicable Human gate.
- Open / deferred: idempotency lookup versus closed-Campaign rejection ordering; exact amount representability across Donation persistence and Campaign funding. Existing O3/O4/O5 controls, runtime evidence, and separate Tier-0 authorization remain as recorded in the Techplan.
- Recommended next step: one resolution pass, then Human gate; re-review if the resolution changes material semantics, architecture, data contract, or verification strategy.
- Session transition: start a fresh Planner Run/Participant Session for resolution so the Reviewer stays independent and the prior Draft remains history.
- Context pointers: [Techplan §9 admission ordering and §§8, 4 R3](/home/anhar-solehudin/kencleng-workspace/kencleng/.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-002/techplan.md); [Techplan D12, §§8–9, §11](/home/anhar-solehudin/kencleng-workspace/kencleng/.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-002/techplan.md); `api/openapi/donation.yaml` `SubmitDonationRequest`; `docs/project/kencleng-monetary-data-standard.md`; `backend/migrations/000011_create_public_campaigns.up.sql`; `docs/spec/5-donation/invariants.md#inv-donation-08`.
