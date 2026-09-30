# Review findings — `TP-S2-002-010`

## Provenance

- Phase: Independent Techplan review
- Work Unit / Run: `WU-S2-002` / `RV-S2-002-005`
- Author / Participant: Codex Reviewer / `P-S2-002-RV-005-1`
- Created: 2026-09-29
- Model / reasoning: `gpt-6-luna` / `high` requested by current local runtime configuration; active runtime identity/settings not independently exposed
- Session: Fresh Reviewer Session; Session ID not exposed
- Target revision: `dfbcf7b9241dbe73e2dc2b4c7707f0f829b6cff5`, with the explicitly referenced current untracked TP-010 artifacts
- Workflow revision: `06a38c668b66227c3531431471b32f4f7df3699b`

## Review findings — `TP-S2-002-010`

**Gate:** Complex — cross-domain Campaign/Donation contract and concurrency-sensitive money ordering; payment and guest PII/security boundaries. Independent review is required.

**Sections resolved:** 1 Background; 2 Scope; 3 Requirements; 4 Rules & Validation; 5 Decision Log; 6 Backward Compatibility; 7 Edge Cases & Risks; 8 Interface Contract; 9 Architecture / Plan; 10 Implementation Details; 11 Files Changed / Files NOT Changed; 12 Testing Checklist and Test Focus Pointer; 13 Open Items.

### Blocking

- **MATERIAL / BLOCKING — [AUTHORITY / DECISION LIFECYCLE] O7 is incorrectly still Active — §13 O7, with effects in §§5, 8–10 and R10.** The current plan asks Design Authority to decide/review terminal source labels and email wording, but completed `OIR-S2-002-004/design-review-brief.md` records the bounded 2026-09-28 decisions: use “Hasil simulasi donasi: berhasil/gagal” (source first) and the approved near-opt-in verification disclosure; pending copy, unavailable methods, generic link failure, and deliberate failed-state recovery were `GUIDANCE_SUFFICIENT`. The current manifest/events also mark O7 complete. Carry those outcomes into the execution spine and retain only genuinely deferred rendered acceptance and O2/O3 delivery dependencies. No repeat Design decision is needed. This omission leaves a resolved owner item looking like an unresolved approval gate.
- **MATERIAL / BLOCKING — [PRODUCT / SECURITY-PII CONTRACT CONFLICT] O2/O3 incompatibility is absent from current §13 and is not preserved by O3's generic retention wording.** Current durable state records two still-effective directions: Donation delivery owner chose an independent pending-email cap (route B), while the named Product/MVP owner requires a verified opted-in address to remain eligible for the terminal status notice. The Events entry for 2026-09-28 says a finite cap can delete the only verified address before terminal, O2 has no terminal bound, and the directions cannot currently both be met; the tracker and `WU-S2-002/manifest.md` retain this scoped `HUMAN_DECISION`. TP-010 §13 O3 lists generic Security/PII windows and controls but neither records the conflict nor identifies the smallest owner reconciliation needed. Preserve both decisions and this scoped blocker; do not infer a cap, contact-retention alternative, timeout-as-failed meaning, or risk acceptance. Without the conflict, downstream contract work would have to invent how mandatory terminal notice can be delivered.
- **MATERIAL / BLOCKING — [AUTHORITY SYNC / MONEY CONTRACT] Current O1 omits the unmapped cross-feature currency-standard authority — §13 O1 and §8 Persistence/data shape.** OIR-S2-002-005 and current `manifest.md`, `events.md`, and tracker record Human direction for a shared standard across currencies, tables, and features, while wire/storage representation remains open and no project-wide owner is named in `.harscode-spaces/authority-map.md`. TP-010 §13 O1 routes representation only to “Donation/API owner; Product Authority if semantics change” and omits this distinct `AUTHORITY_SYNC`. Keep the Slice-2 Donation owner and the cross-feature authority concern separate; route owner/scope attribution before making the shared representation decision. This gap blocks final O1 representation/storage reconciliation, though it does not block this D1 amendment or unrelated work.

### Non-blocking

- None identified. Do not polish beyond resolving the material lifecycle and scoped-authority gaps.

### Clean

- D1 is faithfully recorded in Q8, R7, D15, §8, the risk row, and R7 verification. It preserves O6, atomic success/funding exact-once behavior, the winning close reason, and Slice 3 boundary without selecting a lock/isolation mechanism or endpoint.
- Rules R1–R10 have corresponding §12 verification coverage; repeated R4/R5/R6/R7/R9 rows assign distinct Human/Testing evidence appropriately. The Test Focus Pointer retains exact Exploration anchors for money/concurrency, retry, guest status exposure, email PII, and the new Campaign/Donation ordering risk. No diagram is present.
- The plan retains the API split-source/generated-artifact discipline and the Tier-0 write fences; no protected implementation write is proposed. O8 remains conditional as intended.
- Technical spot checks confirmed that `campaign.Service.toPublicDetail` sets `DonationAction` to `unavailable` / `donation_flow_not_available` and `PublicCampaignDetailHandler` projects that public-safe field; the live router has no Donation route. Historical `INV-donation-08` / submit-settlement wording conditions funding on Campaign `published`, supporting the stated reconciliation gap rather than contradicting D1. `api/README.md` confirms split domain sources are authored and bundle/generated outputs are derived.

## Phase handoff

- Completed: Complex gate and independent Techplan review.
- Artifacts: this file and `launch-record.md` in `RV-S2-002-005`.
- Human decision: No D1 or O7 decision is requested again. After one Planner resolution pass and independent review convergence, Human reviews the regenerated report and amended Techplan at the required material Techplan approval gate. The O2/O3 conflict needs its bounded Product/Delivery owner reconciliation; O1 cross-feature owner/scope attribution needs Authority Sync before the affected O1 choice.
- Open / deferred: Three material findings above; no Build or `CONTRACT_READY`.
- Recommended next step: Fresh Planner Run to resolve these review findings within existing owner decisions, propagate the completed O7 decisions, and preserve the two scoped O1/O2-O3 gates; then follow the applicable review route for materiality and regenerate the Human report only after convergence.
- Session transition: Fresh Planner Participant/Run, because independent review has ended and material Techplan resolution is a new workflow occurrence.
- Context pointers: TP-010 §§5, 8, 12–13; `OIR-S2-002-004/design-review-brief.md`; `events.md` 2026-09-28 O2/O3 and O1 authority-sync entries; `WU-S2-002/manifest.md`; `docs/project/kencleng-development-tracker.md`.
