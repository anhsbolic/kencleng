# Stage A — Current-State Simplification Baseline

> Experiment: Harscode Pilot #2 current-state simplification (`SR-002`, `SR-004`, `SR-005`)
> Captured: 2026-09-29, before OIR-S2-002-006 completion artifacts appeared
> Target snapshot: Kencleng `934b093cea30230743dee951ae1e600762019f30`
> Guidance snapshot: Harscode `9b26db6847d893523838e246b2048aed97dcff67`
> Method: read-only reconstruction from the normal current artifact set; no chat or remembered state used as source.

## Observable reconstruction

- **Parent Outcome / active Slice:** `S2-GUEST-DONATION-TRUTHFUL-STATE`, Slice 2 `IN_PROGRESS`; Slice 1 remains `SLICE_FINALIZED`.
- **Active Work Unit:** `WU-S2-002`, `ACTIVE / QUEUED`, horizon `NOW`. `WU-S2-001` is `DONE`.
- **Current Run / runnable route:** `OIR-S2-002-006` is prepared but not dispatched. It is an independent Campaign/Donation threshold and settlement ordering route; it does not depend on O1 currency-standard resolution.
- **Current-effective plan / milestone:** `TP-S2-002-007` is Human-approved and current-effective. No Slice 2 milestone or `CONTRACT_READY` is earned; implementation has not started.
- **Authority state:** Anhar is named for the relevant Campaign/Donation/API/Security/Design areas within current Slice 2. The project-wide shared currency-standard authority has no named owner/scope. Existing local owner attributions do not imply global authority.
- **Scoped blockers:** O1 `AUTHORITY_SYNC` blocks only final shared-currency amount representation/storage reconciliation. O2/O3 `HUMAN_DECISION` blocks the dependent email-retention/terminal-notification contract. Independent OIR-006 remains runnable. O4/O5 control/contract evidence and O6/O7/O9 spec/API translation remain outstanding; O8 is conditional on operation removal/replacement.
- **Human action now:** mechanically dispatch OIR-006 through Human-Assisted Orchestration. Human confirms the canonical Exploration Stage 1 plan and participates in Stage 3 if the named owner decision is needed.
- **Uncertainty/conflict observed:** detailed state is repeated across the Work Unit manifest, Work Graph, Control Surface, Parent Outcome, and tracker. Their current frontier agrees. The readiness assessment and OIR-006 invocation carry an older Harscode provenance revision than the current guidance snapshot; current guidance treats ordinary workflow revision as provenance unless explicitly assignment-pinned. No material route contradiction was found.

## Comparison use

This is the control snapshot for Stage B readiness-artifact ablation. A later comparison must reconstruct the same underlying target state while excluding reliance on `readiness-reconciliation.md`; if the underlying Work Unit state changes first, capture a new baseline rather than attributing the difference to the ablation.
