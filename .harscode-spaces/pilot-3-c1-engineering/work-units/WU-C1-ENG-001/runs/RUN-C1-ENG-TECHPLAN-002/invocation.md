# Run Invocation — RUN-C1-ENG-TECHPLAN-002

## Identity and routing

- **Work Unit:** `WU-C1-ENG-001`
- **Run:** `RUN-C1-ENG-TECHPLAN-002`
- **Phase route:** Techplan synthesis — pre-Approval revision after completed Solution Shaping.
- **Role:** Planner
- **Specialization / Participant Profile:** None; canonical Planner role and Run-local C1 scope remain sufficient, as in Run `001`. No reusable project Profile is established.
- **Participant:** `PARTICIPANT-C1-ENG-PLANNER-002` (reserved; instantiate only on approved Human mechanical dispatch).
- **Session:** `SESSION-C1-ENG-PLANNER-002` (reserved; fresh context).
- **Session transition:** FRESH — prior synthesis has terminated; new solution evidence justifies phase re-entry with a new Run/Participant.
- **Dispatch posture:** Human-Assisted; prepared, not dispatched.
- **Human escalation owner:** Anhar via the Orchestrator.

## Task and completion boundary

Revise the same mutable Draft Techplan into an execution-grade C1 plan using the completed Solution Contract. Preserve binding Stage 5/6 meaning. The Solution Contract owns bounded engineering selections, not new Product authority or protected implementation permission. Treat the earlier Draft and its terminal handoff as prior discovery evidence where the contract resolves their questions; do not re-raise those settled questions as missing Product decisions without contradictory new evidence.

Read the Solution Contract in full, consume H1–H3, its selected ownership/interface/security/consistency/guard directions and verification oracles, and preserve rejected alternatives. Express implementation tasks, scoped backend/frontend write boundaries, shared contract coordination, verification ownership, operator runbook/privileges, real-provider/runtime prerequisites, and protected gates G1–G3. Pin compatible dependency versions from current primary evidence without weakening the selected trust contract. Ordinary implementation detail stays Build freedom.

The Run is complete when the Techplan has been revised and self-checked under the canonical synthesis prompt, with exact artifact identity, remaining Open Items, independent-review/decomposition recommendation, and one durable structured Phase handoff. A new contradiction must be preserved and routed to its owner; do not repair upstream Product/pre-engineering artifacts. Pending protected implementation permission is a Build gate, not an unresolved material solution choice or a planning blocker.

Independent planning review remains recommended by prior evidence because C1 spans auth/session, privileged Owner scope, DB atomicity/guard privileges, API, and rendered experience. The Planner assesses current applicability and recommends routing; this Run does not perform independent review. Generate `report-techplan.md` only when the exact revision reaches the Human approval gate after any invoked review/resolution convergence. Do not approve the Techplan or start Build.

## Effective inputs and provenance

**Assignment-defining solution input:**

- `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/solution-shaping/solution-contract.md`
- Committed at target baseline `3e123bcece561f1d0a181b68c85e2677dcfa37eb`; SHA-256 `c136e673937c9ac1a583ecfa6a98ee630e1b0f6b42b78d373876243e55b01482`.
- Verify identity before reliance. Material change requires Orchestrator reconciliation; do not silently substitute a different solution.

**Prior artifacts / meaningful delta:**

- Read every durable Exploration evidence file under `runs/RUN-C1-ENG-EXPLORATION-001/evidence/`: `stage-2-gap-analysis.md` and `stage-3-solutioning.md`.
- `runs/RUN-C1-ENG-TECHPLAN-001/evidence/phase-handoff.md` — terminal first-synthesis evidence. Its earlier Invocation dispatch snapshot is not current workflow state.
- `techplan/techplan.md` — mutable Draft / In Review, not approved. Preparation identity: SHA-256 `188a55d8a6b951bdbc8d29063689606c082969ca8dca52861d0256d84b54c878`.
- Newly effective evidence is the completed Solution Contract: Google OIDC, conservative durable guard, display-name/consequence/inspection flow, and settled material solution boundaries. This justifies Run `002`; it is not a restart without delta.
- Paths in this prior-artifact list are relative to `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/` unless written in full.

**Binding project authority / current-effective guidance:**

- Root `AGENTS.md`; Stage 7 approved handoff and its progressive read order; Product Intent; binding Stage 5 behavior and Stage 6 requirements.
- Scoped `backend/AGENTS.md` and `frontend/AGENTS.md`; applicable `docs/project/`, `docs/ui-ux/`, `api/`, and reopened live code where a planning claim depends on them.
- Harscode root/workflow routing, canonical synthesis prompt, orchestrated overlay, and required Techplan `template.md`, `rules.md`, `guardrails.md` (read all three in full).
- Targeted stack/security guidance through the current best-practice router. Verify applicable authority cited by earlier evidence; do not copy paraphrases as policy.
- Proposal `0040` is Draft experimental background to completed shaping, not adopted mandatory workflow policy. This Run follows current canonical synthesis/review/approval routing.

Target and workflow revisions below are observed baseline/provenance, not a universal freeze of ordinary guidance. Record actual guidance provenance and compare material drift before reliance. Unexpected changes to the starting Draft or pinned solution require reconciliation before overwrite.

## Orchestrated bindings and runtime route

- **Project root / CWD:** `/home/anhar-solehudin/kencleng-workspace/kencleng`
- **Harscode workspace root:** `../harscode-workspace`
- **Work Unit record:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md`
- **WORK_UNIT_PATH / TASK_PATH compatibility root:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`
- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TECHPLAN-002`
- **ARTIFACT_TARGET:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`; conditionally sibling `report-techplan.md` at the converged Human approval gate only.
- **PRIOR_ARTIFACTS:** Explicit inputs above; these replace implicit ordinal workflow paths.
- **Ticket / area:** None established; C1 — Legitimate Organization representation.
- **Communication language:** Bahasa Indonesia; canonical terms/enums and code identifiers remain unchanged.
- **Communication profile:** None established.
- **TARGET_REVISION:** `3e123bcece561f1d0a181b68c85e2677dcfa37eb` on `pilot/3-c1-engineering` at preparation, before these coordination edits.
- **WORKFLOW_REVISION:** `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590` observed at preparation; Harscode working tree clean.
- **Runtime harness:** `codex-cli`; Human-Assisted mechanical dispatch, not native subagent substitution.
- **Selected model / reasoning:** `gpt-6-sol` / `medium` — recommendation pending approval.
- **MODEL_APPROVAL:** PENDING for Run `002`. Approval for Run `001` was expressly Run-only and is not inherited.
- **Model routing rationale:** Synthesis needs architecture/cross-cutting analysis across settled auth, Owner scope, guard/transaction/interface, and verification boundaries. The Human registry's non-gated `gpt-6-luna` does not declare architecture/cross-cutting capabilities; `gpt-6-sol` does. `medium` is its lowest supported effort and the minimum sufficiently capable candidate, consistent with prior synthesis. Missing authority/context is not a reason for stronger-model escalation; route those gaps to their owners. If this candidate demonstrates capability insufficiency after complete inputs, ask the Orchestrator to assess escalation under the Human-owned registry.
- **CONTINUATION_CHECKPOINT:** None; fresh Run reconstruction from durable inputs.
- **Configured verification environment:** Podman / podman-compose per `.harscode-spaces/.local-config.yaml`; planning must not assume Docker CLI absence means runtime unavailable. This Run does not execute verification.

## Execution envelope

- **PREAUTHORIZED:** Read current project/workflow authority and live code; perform read-only dependency/documentation checks; revise the mutable pre-Approval Techplan at ARTIFACT_TARGET; write Run-owned synthesis evidence under RUN_PATH; conditionally produce the report at its proper gate.
- **ORCHESTRATOR_DECISION:** Route newly discovered coordination gaps or independent review/decomposition after the Planner handoff. No Participant self-dispatch into another phase.
- **HUMAN_REQUIRED:** Explicit approval before dispatch of the selected gated model; exact-revision Techplan approval; G1 auth/session core, G2 initial Owner and object scope, and G3 guard control implementation authorizations before their Build surfaces. Editing upstream authority requires its owning Product/pre-engineering path, not this engineering Run.
- **Out of scope:** Production code, migrations, executable API/client changes, tests/builds/runtime execution, upstream Product/pre-engineering writes, orchestration projection edits, approval/promotion, independent review, Build, and delivery/verification claims.

## Canonical prompt, overlay, and terminal carrier

Use `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md` with `../harscode-workspace/workflow/orchestrated-run-overlay.md` and these explicit identity/input/output bindings. Canonical phase rules own the work; the overlay owns orchestrated path/identity and structured handoff rendering.

Write the terminal carrier at `RUN_PATH/evidence/phase-handoff.md`, with one `## Phase handoff`, known Run/Participant/Session provenance, exact Techplan content identity, Outcome, Result refs, Findings, Decision requests, Blockers, Open / unverified, Recommended continuation, and Context refs as required by the overlay. Distinguish prepared, executed, reviewed, approved, and verified. Report this pointer to the Human for return to the Orchestrator.

## Prepared state

- **Dispatch readiness:** WAITING_MODEL_APPROVAL.
- **Participant dispatched:** No.
- **Run outcome:** Not started.
- **Next action:** Obtain explicit Anhar approval for `gpt-6-sol` / `medium` for Run `002`; then present the mechanical fresh-session dispatch package. Do not claim RUN READY before this approval.
