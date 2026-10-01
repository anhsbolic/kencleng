# Review findings — TP-S2-002-014

> Phase: Independent Techplan Review  
> Author: Codex Reviewer  
> Participant ID: `P-S2-002-RV-009-1`  
> Profile: `KC-REVIEWER` (`profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`)  
> Role: Reviewer — Independent Complex review of material O1/O8/O11 reconciliation  
> Model / reasoning: `gpt-6-luna` / `high` (Invocation configuration; runtime model not independently exposed)  
> Session: Fresh Reviewer Session; Session ID not exposed  
> Created: 2026-10-01  
> Target revision: `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus current working-tree orchestration update; assigned Techplan read from the live checkout  
> Workflow revision: `33b03a3f62cc3aacba6534b8a011465613c64b09`

## Review findings — TP-S2-002-014

**Gate:** Complex. The plan has 14 Rules & Validation entries, crosses Donation/Campaign and authored API contracts, and covers high-stakes money, payment, concurrency, security, and PII boundaries.

**Sections resolved:** Background §1; Scope §2; Requirements §3; Rules & Validation §4; Decision Log §5; Backward Compatibility §6; Edge Cases & Risks §7; Interface Contract §8; Architecture / Plan §9; Implementation Details §10; Files Changed / Files NOT Changed §11; Testing Checklist + Test Focus Pointer §12; Open Items §13.

### Blocking

- **MATERIAL / BLOCKING — O4 credential decisions are reopened as unresolved — §3 Q7; §8 “Persistence/data shape” status-credential bullet; §8 “API/event/external interface” paragraph; §13 Active O4.** TP-014 says token carrier and storage remain undecided and routes them back to Security/API. The current OIR-002 resolution records Human choices: fragment URL with frontend handoff and URL cleanup, and a one-way HMAC verifier; it routes these choices into contract reconciliation while leaving implementation controls and residual-risk acceptance open. The plan therefore drops settled interface/security decisions and would require a later contract author to choose them again or invent a different interpretation. Material because this is the private status credential’s interface and storage boundary. Carry the recorded choices forward without claiming runtime proof or residual-risk acceptance. Evidence: `OIR-S2-002-002/resolution-brief.md`, O4 “Observed Human decision” and “Next route”; TP-014 lines 67, 82, 148, 281.

- **MATERIAL / BLOCKING — O5’s selected public failure code is presented as still open — §3 Q7; §8 “API/event/external interface”; §13 Active O5.** TP-014 says API/Security still need to define the contract-time generic behavior, including the status code, and to reconcile the historical `401`/`404` conflict. OIR-002 records the Human selection of uniform `404` for absent Donation, missing/wrong token, and expired token, with the same public result and body/headers/cache behavior. Exact Problem Details coordinates and the `Cache-Control: private, no-store` contract detail still need reconciliation, and empirical timing/abuse proof remains downstream; the chosen code itself is not undecided. This reopens a resolved interface decision and could lead to a conflicting contract. Evidence: `OIR-S2-002-002/resolution-brief.md`, O5 “Observed Human decision” and “Next route”; TP-014 lines 67, 82, 154, 283.

### Non-blocking

- None identified.

### Clean

- Every R1–R14 has at least one Testing Checklist row. Ownership/rationale generally distinguishes Human authority and rendered acceptance, authored contract validation, and independent runtime Testing; protected ledger/locking work is explicitly fenced.
- O1 carries the approved major-unit decimal string plus explicit currency code direction, exact-decimal calculation/persistence, Product-owned active currency set, and deliberate deferral of range, fraction rules, global scale, and migration detail. O8 is bounded to replacing the two historical operations in this repository/current Slice-2 scope. O11 supersedes route B while leaving numeric bound, architecture, timeout meaning, and residual-risk acceptance undecided.
- D1 preserves eligibility ordering against Campaign close, accepted-pending full settlement, exact-once atomic funding, stable close reason, and possible threshold overshoot without selecting locking/isolation or importing Slice 3. O2–O3 terminal email lifecycle and recovery obligations are carried forward as open delivery/security evidence.
- R14 gives a usable phase boundary: contract-time policy/owner decisions, spec/API consistency, and source validation precede `CONTRACT_READY`; empirical simulator, lifecycle, exposure, and response-parity proof remains for Build/Testing. It does not claim the milestone.
- Task 01 Human acceptance remains parallel. O1 is explicitly Active; O8 and O11 remain Resolved history. No report, approval, Build, or residual-risk acceptance is claimed.
- Test Focus Pointer covers money/settlement concurrency, submission retries, status-credential/security, guest-email PII, and Campaign ordering with exact durable evidence anchors from Exploration/OIR evidence; each remains relevant after synthesis.
- Technical spot-checks agree with current authority: the monetary standard prohibits float and defers universal precision/scale; `api/README.md` defines split OpenAPI files as authored and the bundle/generated types as derived; root and `backend/AGENTS.md` preserve the relevant Tier-0 boundaries. Live source search found historical Donation operations in OpenAPI but no Donation route/domain implementation, consistent with the plan’s evidence-only treatment.

## Phase handoff

- Completed: independent Complex review of the whole TP-014 spine against canonical Techplan guidance, current Product/MVP and monetary authority, durable Exploration/OIR evidence, orchestration decisions, and targeted live-source facts.
- Artifacts: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-009/review-findings.md`.
- Human decision: revise through a fresh Planner resolution Run; do not approve this spine until the blocking findings are resolved. Re-review the material changes.
- Open / deferred: blocking O4 and O5 decision-fidelity findings above; implementation controls, runtime evidence, and Security/PII residual-risk acceptance remain with their stated owners and later gates.
- Recommended next step: one fresh Planner resolution pass to restore the already-selected O4/O5 directions and retain unresolved implementation/evidence work at the correct boundary, followed by fresh independent review.
- Session transition: fresh Planner Participant/Run for resolution, because this Reviewer Run is complete and the findings concern material contract/security decisions.
- Context pointers: TP-014 §3 Q7, §8 status credential and public failure paragraphs, §13 Active O4/O5; `OIR-S2-002-002/resolution-brief.md` O4/O5 decisions.
