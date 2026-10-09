# Run Invocation — RUN-C1-ENG-TECHREVIEW-001

## Identity and routing

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-TECHREVIEW-001`.
- **Phase route:** Independent Techplan review of the post-shaping pre-Approval execution plan.
- **Role:** Reviewer.
- **Specialization / Participant Profile:** None; canonical Reviewer role with bounded C1 planning-review scope. No reusable project Profile is established.
- **Participant:** `PARTICIPANT-C1-ENG-TECHREVIEWER-001` (reserved; instantiate on approved Human dispatch).
- **Session:** `SESSION-C1-ENG-TECHREVIEWER-001` (reserved; FRESH).
- **Session transition reason:** Independent actor context from Planner `PARTICIPANT-C1-ENG-PLANNER-002` and its synthesis conversation. Reconstruct from durable sources; do not reuse that Participant/Session or present a self-check as independent review.
- **Dispatch posture:** Human-Assisted; prepared, not dispatched.
- **Trigger:** Completed `RUN-C1-ENG-TECHPLAN-002` and its recommendation; Orchestrator routes review before resolution/report/Human approval.
- **Escalation owner:** Anhar via the Orchestrator.

## Assignment and completion boundary

Independently review the exact In Review Techplan against durable source evidence and current applicable authorities using the canonical Draft independent-review prompt. Evaluate Step 0 before substantive review. Routing rationale: high-stakes auth/session/PII, privileged Owner/object scope, shared API, and PostgreSQL guard/transaction boundaries meet its Complex signals, although this plan has 14 Rules & Validation entries. This is invocation of the current Draft correctness tool, not promotion to universally mandatory policy.

Check rule/decision fidelity and verification ownership/economics, Open Items lifecycle, sensitive Test Focus Pointers, and 2–3 non-obvious technical facts against live sources. Consume Solution Contract selections and H1–H3 when reconciling earlier Exploration/Draft recommendations; do not re-litigate accurately recorded settled choices merely from preference. Preserve upstream Product meaning. Pending G1–G3/approval/runtime prerequisites must remain explicitly scoped gates, not hidden implementation permission or fabricated unresolved Product meaning.

Every finding includes location, defect, exact source evidence, materiality, and the affected concern. Classify MATERIAL / BLOCKING or MECHANICAL / NON-BLOCKING using the prompt; do not polish after material checks. If evidence exposes a Product/solution contradiction, preserve and route it to its owning authority. Do not repair upstream artifacts or revise the Techplan.

Completion: Step 0 verdict and, when warranted, independent review of the captured revision, durable findings, and one structured terminal Phase handoff. No correctness verdict is presumed by dispatch. Findings route through Planner resolution and the exact-revision Human gate; material resolution is re-reviewed unless the Human gate waives it. This Reviewer does not self-dispatch that continuation, produce an approval report, approve the plan, or start Build.

## Exact target and effective inputs

**Assignment-defining review target:**

- `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`
- Status: **In Review**, mutable pre-first-Approval spine, not a candidate successor.
- SHA-256: `577673c2b03bc93526362a848636699ebd888b8917519bae65b1adde15860954`.
- This is working-tree synthesis content at observed HEAD `524ef600c7f71246af6b71671d89c3c040fd9d44`, not a claim that HEAD alone contains the revised plan. Verify the hash before review and before terminal handoff. If it changes, stop/reconcile rather than cover a newer mutable target silently; retain only conclusions tied to reconstructable reviewed content. Do not create a per-Run Techplan copy for history.

**Assignment-defining bounded solution evidence:**

- `solution-shaping/solution-contract.md`, SHA-256 `c136e673937c9ac1a583ecfa6a98ee630e1b0f6b42b78d373876243e55b01482`; read in full. It owns selected engineering/design direction and Planning Ready, not Product authority or protected implementation authorization.

**PRIOR_ARTIFACTS / independent source reconstruction:**

- All durable files under `runs/RUN-C1-ENG-EXPLORATION-001/evidence/`, currently `stage-2-gap-analysis.md` and `stage-3-solutioning.md`. Enumerate/read all for independent fidelity checking when the Complex gate applies.
- `runs/RUN-C1-ENG-TECHPLAN-001/evidence/phase-handoff.md` — prior discovery/decision-gap context, interpreted through the later Solution Contract.
- `runs/RUN-C1-ENG-TECHPLAN-002/evidence/phase-handoff.md` and `synthesis-check.md` — exact revised-output provenance and self-check claims, not substitutes for independent review.
- Paths in this section without a full repository prefix are relative to `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/`.

**Binding authority / current-effective guidance:**

- Root `AGENTS.md`; Stage 7 progressive authority route → Product Intent → binding Stage 5 C1 behavior → Stage 6 requirements. Upstream Product/pre-engineering inputs are read-only.
- Applicable backend/frontend `AGENTS.md`, `docs/project/`, `docs/ui-ux/`, `api/`, and live source where facts require spot-checking. Existing neutral code does not establish Product meaning.
- Harscode root/workflow routing; read Techplan `template.md`, `rules.md`, `guardrails.md` and the exact plan in full. Read `diagram-guidelines.md` only if the reviewed plan actually contains a diagram.
- Route targeted security/stack guidance through the current best-practice router as needed; verify referenced authority rather than treating synthesis paraphrases as policy. Reopen current primary dependency documentation/source when a material technical fact depends on it; distinguish metadata inspection from compilation/runtime evidence.
- Proposal `0040` remains experimental background, not adopted policy. Ordinary guidance is current-effective; workflow revision below captures observed provenance without universally pinning it. Compare materially changed guidance before relying on it.

## Orchestrated bindings and runtime route

- **Project root / CWD:** `/home/anhar-solehudin/kencleng-workspace/kencleng`.
- **Harscode workspace root:** `../harscode-workspace`.
- **Work Unit record:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md`.
- **WORK_UNIT_PATH / TASK_PATH compatibility root:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TECHREVIEW-001`.
- **ARTIFACT_TARGET:** None; review findings and handoff are Run-owned evidence. The reviewed Techplan is read-only for this Run.
- **PRIOR_ARTIFACTS:** Explicit sources above replace implicit ordinal workflow paths.
- **Ticket / area:** None established; C1 — Legitimate Organization representation.
- **Communication language / profile:** Bahasa Indonesia / none established; preserve canonical terms/enums and technical identifiers.
- **TARGET_REVISION:** `524ef600c7f71246af6b71671d89c3c040fd9d44`, branch `pilot/3-c1-engineering`, observed at preparation with existing coordination edits and uncommitted synthesis/evidence. Baseline reference only; exact target hash defines the review.
- **WORKFLOW_REVISION:** `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`, observed at preparation; Harscode working tree clean.
- **Runtime harness:** `codex-cli`; Human-Assisted mechanical launch, no automatic subagent substitution.
- **Selected model / reasoning:** `gpt-6-sol` / `medium`, approved for this Run.
- **MODEL_APPROVAL:** APPROVED by Anhar on 2026-10-09 for `RUN-C1-ENG-TECHREVIEW-001` only, in direct response to the exact model/effort request: “approve bro”. This authorizes model dispatch for independent review; it is not Techplan approval, G1–G3 implementation permission, or a review verdict. Prior Planner approvals are not inherited.
- **Model routing rationale:** Independent fidelity review needs architecture/cross-cutting analysis across security, interface, persistence, and verification ownership. The non-gated `gpt-6-luna` registry entry does not declare those capabilities; `gpt-6-sol` does, and `medium` is its lowest supported sufficiently capable effort. Independence comes from a distinct Reviewer/actor context; a different or stronger model is not required by phase policy. If complete inputs reveal capability insufficiency, route escalation to Orchestrator/Human under the registry; do not substitute stronger reasoning for missing authority/context.
- **CONTINUATION_CHECKPOINT:** None; new Run/fresh independent reconstruction.

## Execution envelope and phase entrypoint

- **PREAUTHORIZED:** Read current authorities, source/evidence, and primary technical documentation; write only Run-owned findings and handoff under RUN_PATH. Routine read-only hash/structure/fact checks are permitted.
- **ORCHESTRATOR_DECISION:** Route findings to Planner resolution, material re-review, or Human ownership after handoff. Reviewer cannot perform those phases within this Run.
- **HUMAN_REQUIRED:** Approve gated model dispatch; exact Techplan approval; separate G1 auth/session, G2 Owner/object scope, G3 guard-control implementation permission. No approval is inferred from shaping, model use, clean review, or successful synthesis.
- **Out of scope:** Techplan/report rewrites, upstream authority/projection writes, production/API/migration changes, tests/build/runtime verification, protected implementation, approval/promotion, and delivery claims.

Use `../harscode-workspace/workflow/2-2-techplan-review-prompt.md` with `../harscode-workspace/workflow/orchestrated-run-overlay.md` and these input/path bindings. The canonical prompt owns review checks/findings rendering; the overlay owns orchestrated identity and structured terminal handoff. Resolve current section names/numbers from the current template.

Write findings to `RUN_PATH/evidence/review-findings.md` and the terminal carrier to `RUN_PATH/evidence/phase-handoff.md`, with exactly one `## Phase handoff`, known provenance, exact reviewed content identity, Step 0/result refs, Findings, Decision requests, Blockers, Open / unverified, Recommended continuation, and Context refs. A Step 0 N/A result still records its reason and terminal outcome; do not manufacture a skipped substantive review. Return both evidence pointers, or the exact scoped input discrepancy if execution stops.

## Prepared state

- **Dispatch readiness:** READY_FOR_HUMAN_DISPATCH.
- **Participant dispatched:** No.
- **Run outcome:** Not started.
- **Next action:** Human mechanically dispatches this Run using `gpt-6-sol` / `medium` in a fresh independent Reviewer Session; no further model approval is required for this Run.
