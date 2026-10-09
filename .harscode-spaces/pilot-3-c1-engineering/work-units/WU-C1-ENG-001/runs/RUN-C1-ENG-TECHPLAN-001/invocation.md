# Run Invocation — RUN-C1-ENG-TECHPLAN-001

## Identity and routing

- **Work Unit:** `WU-C1-ENG-001`
- **Run:** `RUN-C1-ENG-TECHPLAN-001`
- **Phase route:** Techplan synthesis — first pre-Approval synthesis of the C1 execution-grade plan.
- **Role:** Planner
- **Specialization:** None; Run-local C1 planning scope is sufficient.
- **Participant:** `PARTICIPANT-C1-ENG-PLANNER-001` (ephemeral; instantiate only after gated-model approval and Human mechanical dispatch)
- **Session:** `SESSION-C1-ENG-PLANNER-001` (fresh context; new Run after Exploration termination)
- **Dispatch posture:** Human-Assisted; prepared, not dispatched.
- **Trigger:** Orchestrator routing after completed C1 Exploration Run `RUN-C1-ENG-EXPLORATION-001` with no Blocker.

## Task

Synthesize the first execution-grade Techplan for C1 — Legitimate Organization representation — from the completed Exploration evidence and current binding C1 behavior/requirements. Treat Stage 3 solutioning as recommendations and evidence, not approved architecture. Reopen live code and current authority when technical facts or code anchors matter.

The Techplan must preserve Stage 5/6 meaning, keep unresolved material technical or authority questions visible, and be usable by a fresh Build agent without inventing product/domain, authority/security, architecture/ownership, interface/data, risk, or verification decisions. Do not implement code, alter upstream Product/pre-engineering authorities, or turn the Exploration transaction recommendation into an approved decision without synthesis and rationale. Preserve the explicit Human authorization boundary before implementation of an initial Owner grant or other privileged authorization rule.

Write the first mutable Techplan at `ARTIFACT_TARGET` with status Draft/In Review. Generate `report-techplan.md` only when the exact Techplan revision has converged to its Human approval gate, after any applicable planning review/resolution. Do not continue to Build or treat the plan as approved.

## Current-effective inputs

Read all durable Exploration evidence for this fresh Techplan Session:

- `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-EXPLORATION-001/evidence/stage-2-gap-analysis.md`
- `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-EXPLORATION-001/evidence/stage-3-solutioning.md`

Binding/current project authorities:

- `AGENTS.md`, `backend/AGENTS.md`, and `frontend/AGENTS.md` as relevant;
- `docs/product/pilot-3-stage-7-c1-engineering-handoff.md` for scope/navigation;
- `docs/product/pilot-3-stage-5-c1-confirmed-behavior.md` for binding C1 behavior;
- `docs/product/pilot-3-stage-6-c1-requirements.md` for binding C1 requirements;
- `docs/product/product-intent.md` for whole-product authority;
- applicable current docs in `docs/project/`, `docs/ui-ux/`, `api/`, and live code only where a plan claim depends on them.

Read in full the required Techplan authorities:

- `../harscode-workspace/workflow/2-techplan/template.md`
- `../harscode-workspace/workflow/2-techplan/rules.md`
- `../harscode-workspace/workflow/2-techplan/guardrails.md`

Use the canonical prompt and overlay below. Route stack/security topics through the current Harscode best-practice index and verify any matching authority cited by Exploration. Do not read unrelated examples or best-practice files.

## Orchestrated inputs

- **Project root / CWD:** `/home/anhar-solehudin/kencleng-workspace/kencleng`
- **Harscode workspace root:** `../harscode-workspace`
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md`
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TECHPLAN-001`
- **Artifact target:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`; conditionally also `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/report-techplan.md` only at the converged Human approval gate.
- **Prior artifacts:** the two completed Exploration evidence files listed above.
- **Codebase context:** Kencleng — Go + PostgreSQL backend and Next.js/React/TypeScript frontend; current product behavior is not inferred from the neutral scaffold.
- **Ticket / Area:** none established; C1 — Legitimate Organization representation.
- **Communication language:** Bahasa Indonesia.
- **Communication profile:** none.
- **Target revision:** `e9cb1f31f031ed3905181eb86ad3c3e3af4e248f` (`pilot/3-c1-engineering` at preparation; reopen current files before reliance).
- **Workflow revision:** `d882b35c88ae7b4adff19abba991e0f2f9d7b7ac` (`../harscode-workspace` at preparation).
- **Runtime harness:** `codex-cli`.
- **Selected model:** `gpt-6-sol`.
- **Reasoning effort:** `medium`.
- **Model approval:** **APPROVED by Anhar on 2026-10-09 for this Run only.** The Human-owned registry marks this model `approval_required: true`; approval does not change the registry or authorize another Run.
- **Model routing rationale:** The Planner must synthesize cross-layer architecture, data/transaction boundaries, verification, and privileged-authorization constraints from the approved C1 requirements and neutral baseline. The registry's non-gated `gpt-6-luna` entry does not declare the architecture/cross-cutting-analysis capabilities needed here. `gpt-6-sol` declares those capabilities; `medium` is its lowest supported effort and is the selected minimum candidate. This is a gated recommendation, not dispatch authorization.
- **Continuation checkpoint:** none; new Run, Participant, and fresh Session after Exploration termination.

## Execution envelope

- **PREAUTHORIZED:** Read current project/workflow authorities and live source; write only the stable pre-Approval Techplan artifact at `ARTIFACT_TARGET` and Run-owned synthesis evidence under `RUN_PATH`; record technical recommendations and unresolved material questions within the Techplan lifecycle.
- **HUMAN_REQUIRED:** Dispatch this gated model; edit Product/pre-engineering authority; implementation changes to source; implementation of privileged Owner/authorization rules without explicit Human authorization; and any material authority change.
- **Out of scope:** Build, test execution, approval/promotion of the Techplan, Product semantic changes, and claiming C1 delivered or verified.

## Canonical prompt and overlay

Run the canonical phase contract from:

`../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`

Apply orchestrated identity/path semantics from:

`../harscode-workspace/workflow/orchestrated-run-overlay.md`

The canonical prompt owns Techplan synthesis, self-check, optional review/decomposition recommendation, and its Human approval boundary. The overlay maps its prior Exploration and output paths to the explicit bindings above.

## Prepared state

- **Dispatch readiness:** READY_FOR_HUMAN_DISPATCH.
- **Participant dispatched:** No.
- **Run outcome:** Not started.
- **Next action:** Human mechanically dispatches the prepared Run using the card; no further model approval is needed for this Run.
- **Expected terminal report-back:** One structured `## Phase handoff`, exact Techplan artifact/revision, findings/open items, independent review/decomposition recommendation, and the exact Human approval decision required. Build must not start from this Run.
