# Run Invocation — RUN-C1-ENG-EXPLORATION-001

## Identity and routing

- **Work Unit:** `WU-C1-ENG-001`
- **Run:** `RUN-C1-ENG-EXPLORATION-001`
- **Phase route:** Harscode Exploration, beginning at Stage 1 — Plan Announcement. Stage 1 is a hard stop; Stage 2 requires Human confirmation under the canonical prompt.
- **Role:** Explorer
- **Specialization:** None; Run-local scope is sufficient for this initial handoff-consumption Run.
- **Participant:** `PARTICIPANT-C1-ENG-EXPLORER-001` (ephemeral; instantiated by the Human's mechanical dispatch)
- **Session:** `SESSION-C1-ENG-EXPLORER-001` (fresh context)
- **Dispatch posture:** Human-Assisted; this invocation prepares assignment but does not itself dispatch the Participant.
- **Trigger:** Human approval on 2026-10-09 to prepare the C1 engineering transition.

## Task

Perform the first fresh-reader Engineering Exploration for C1 — Legitimate Organization representation, using the approved Stage 7 handoff as the task entrypoint and the Stage 5/6 artifacts as binding behavior and requirements authority. This first session also begins the handoff's planned cold-start validation: determine from durable artifacts whether the C1 task and its authority can be understood without relying on the original product conversation.

For this dispatch, execute **Stage 1 only**: announce a concise understanding and the ordered areas to explore, then stop for Human confirmation. Do not begin Stage 2 or inspect implementation deeply before that confirmation. If the Run continues after confirmation, use the canonical Exploration stages and preserve evidence under `RUN_PATH`.

Do not steer the exploration toward expected gaps, Work Units, implementation solutions, or a presumed backend/frontend split. Let the current authority and Stage 1 inspection determine relevant areas. Do not modify source code or upstream Product/pre-engineering artifacts. Do not add or run tests in this Run.

## Current-effective inputs

Read in the Stage 7 progressive order, stopping when later context is not materially needed:

1. `AGENTS.md` — project routing and boundaries.
2. `docs/product/pilot-3-stage-7-c1-engineering-handoff.md` — approved task entrypoint and authority navigation.
3. `docs/product/product-intent.md` — whole-product authority.
4. `docs/product/pilot-3-stage-5-c1-confirmed-behavior.md` — binding C1 behavior.
5. `docs/product/pilot-3-stage-6-c1-requirements.md` — binding C1 requirements.
6. Stage 4 and Stage 3B only if materially needed, as directed by Stage 7.

Applicable project architecture sources when relevant: `docs/project/kencleng-backend-tech-stack.md`, `docs/project/kencleng-frontend-tech-stack.md`, and the relevant scoped `AGENTS.md`. These are current at execution time; reopen them before relying on implementation facts.

## Orchestrated inputs

- **Project root / CWD:** `/home/anhar-solehudin/kencleng-workspace/kencleng`
- **Harscode workspace root:** `../harscode-workspace`
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md`
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-EXPLORATION-001`
- **Artifact target:** none; Exploration evidence belongs to this Run.
- **Prior artifacts:** the current-effective inputs listed above; no prior engineering Run artifacts.
- **Codebase context:** Kencleng — Go + PostgreSQL backend and Next.js/React/TypeScript frontend; active product behavior is not inferred from the neutral scaffold.
- **Ticket / Area:** none established; let Stage 1 determine relevant areas.
- **Communication language:** Bahasa Indonesia.
- **Communication profile:** none.
- **Target revision:** `e9cb1f31f031ed3905181eb86ad3c3e3af4e248f` (`pilot/3-c1-engineering` at preparation).
- **Workflow revision:** `d882b35c88ae7b4adff19abba991e0f2f9d7b7ac` (`../harscode-workspace` at preparation).
- **Runtime harness:** `codex-cli`.
- **Selected model:** `gpt-6-luna`.
- **Reasoning effort:** `medium`.
- **Model routing rationale:** The initial fresh-reader Exploration needs repository reading and reasoning across a bounded C1 handoff. The Human-owned registry declares `gpt-6-luna` supports repository work, coding, and reasoning at `medium`, at the lowest listed cost tier, without a model approval gate. No stronger model is justified by current evidence.
- **Model approval:** Not required by the selected registry entry.
- **Continuation checkpoint:** none; this is a new Run with a fresh Participant Session.

## Execution envelope

- **PREAUTHORIZED:** Read/search the target repository and the listed current-effective authorities; write only Run-owned Exploration evidence under `RUN_PATH` when the canonical phase reaches a stage requiring durable evidence.
- **HUMAN_REQUIRED:** Source/production implementation changes; edits to upstream Product/pre-engineering authorities; and any material change in privileged authorization, money movement, cryptographic key handling, or authentication core logic.
- **Out of scope for this Run:** Implementing or testing product behavior, creating API/data semantics, changing authority, or claiming C1 delivery/verification.

## Canonical prompt and overlay

Run the canonical phase contract from:

`../harscode-workspace/workflow/1-exploration-kickoff-prompt.md`

Apply orchestrated identity/path semantics from:

`../harscode-workspace/workflow/orchestrated-run-overlay.md`

Use the wrapper at `../harscode-workspace/orchestration/exploration-kickoff-prompt.md` for dispatch shape. The canonical phase prompt remains authoritative for Exploration behavior and its Stage 1 hard stop.

## Prepared state

- **Dispatch readiness:** READY_FOR_HUMAN_DISPATCH
- **Participant dispatched:** Yes; Stage 1 plan received from the fresh Explorer session.
- **Run state:** Active, paused at the canonical Stage 2 Human checkpoint.
- **Expected immediate response:** Stage 1 plan announcement only, followed by a stop for Human confirmation before Stage 2.
- **Stage 1 checkpoint:** Return the plan announcement to the Human/Orchestrator for review. Do not continue to Stage 2 until the Human explicitly confirms or redirects.
- **Received Stage 1 result:** The Explorer restated the C1 handoff scope and authority, proposed a three-area exploration order, reported no deep implementation inspection or Run evidence written, and stopped for Human direction. The cold-start validation remains incomplete.
- **Stage 2 result:** Complete. The durable handoff was sufficient to reconstruct the C1 task, authority route, scoped behavior, engineering decision space, and exclusions for this fresh reader. The baseline contains no live C1 frontend flow, backend domain/state, persistence/migration, or API operation; this is consistent with the handoff's not-started status and does not block Stage 3. No implementation or verification claim is supported.
- **Stage 2 evidence:** `evidence/stage-2-gap-analysis.md`; Findings F1 (C1 not implemented in the baseline) and F2 (cold-start task/authority comprehension sufficient for this task).
- **Current Human gate:** Confirm Stage 2 and authorize Stage 3, or redirect the analysis.
- **Terminal report-back:** When the Run eventually terminates, return its single structured `## Phase handoff` and any referenced durable Exploration evidence. If a discrepancy prevents safe continuation, return its exact source and affected scope.
