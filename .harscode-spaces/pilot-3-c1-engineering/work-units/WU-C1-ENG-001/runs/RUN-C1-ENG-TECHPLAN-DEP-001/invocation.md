# Run Invocation — RUN-C1-ENG-TECHPLAN-DEP-001

## Identity and routing

- **Work Unit:** `WU-C1-ENG-001`
- **Run:** `RUN-C1-ENG-TECHPLAN-DEP-001`
- **Phase route:** Techplan synthesis — bounded material successor reconciliation after `B-T2-001`.
- **Role:** Planner
- **Specialization / Participant Profile:** None; canonical Planner role and Run-local scope suffice.
- **Participant:** `PARTICIPANT-C1-ENG-DEPENDENCY-PLANNER-001` (reserved; instantiate on mechanical dispatch).
- **Session:** `SESSION-C1-ENG-DEPENDENCY-PLANNER-001` (fresh context).
- **Session transition:** FRESH — new planning occurrence after T2 Build terminated BLOCKED; independent Planner reconstruction is required.
- **Dispatch posture:** Human-Assisted; prepared, not dispatched.
- **Human escalation owner:** Anhar via Orchestrator.

## Task and completion boundary

Reconcile only the approved Techplan's Go toolchain/auth dependency baseline and the applicability disposition needed to resolve `F-T2-001` / `B-T2-001`. Preserve the approved OIDC trust contract and all existing C1 semantics, scope, task topology, approval boundaries, and verification ownership. Distinguish scanner findings, dependency/source reachability, advisory applicability, and residual risk using current primary evidence. Do not infer exploitability from a symbol-level scanner hit; do not dismiss an advisory without a source-grounded applicability analysis and a durable disposition.

Read the T2 Build report and its `dependency-vuln.json`, `dependency-summary.json`, `probe.go`, probe `go.mod`/`go.sum`, and `artifact-identities.sha256`; inspect the approved Techplan §10 dependency/security direction, T2 task, accepted manifest, relevant backend module/toolchain anchors, current root/backend instructions, current Harscode synthesis authorities, and matching supply-chain/security best-practice authorities. Browse current primary advisory and Go/module release sources because vulnerability and supported-version facts can change. Record source URLs, access date, affected/fixed ranges, and any reachability limits in Run evidence and candidate.

Allowed candidate changes are limited to dependency/toolchain baseline pins, supporting compatibility/security rationale, explicit advisory applicability/risk disposition, and directly affected verification/Build gates. If evidence indicates changes outside that boundary, or leaves a material decision ambiguous, stop and surface the smallest Decision/Open Item rather than broadening scope or guessing. Candidate remains Draft/In Review and is not effective until applicable independent review/resolution and exact Human approval, followed by bounded promotion.

This Run does not update production dependencies, source, generated API, task snapshots, manifest, Solution Contract, Product/pre-engineering authority, Techplan predecessor, Control Tower, or other orchestration records. It does not dispatch Build/T3/T5, choose or commit a replacement dependency outside the candidate, weaken any trust check, suppress scanner findings, alter task topology, approve/promote the candidate, or claim runtime/security verification. Preserve the five-task split, G1–G3 approvals, and all existing authorization boundaries. Existing T5 dependency readiness remains independently determined by the accepted manifest.

Complete when the single candidate successor has been revised and self-checked against the canonical prompt, exact evidence identities and sources are recorded, remaining Open Items are explicit, review recommendation is stated, and one structured terminal handoff is written. If it converges to the Human approval gate after applicable review/resolution, the Planner owns generating `report-techplan.md` from the canonical template for that exact candidate revision. Do not generate the report during synthesis churn or continue to Build.

## Effective inputs and provenance

- Approved predecessor: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`, SHA-256 `bd732748b486f0327f6cba2cdf14cab32f1ece3c5be8ebb24f9659444947fa47`.
- T2 task snapshot: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/T2-backend-identity-session.md` (verify hash before reliance).
- Accepted split: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/manifest.md`, SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.
- Trigger evidence: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-001/evidence/build-report.md`, SHA-256 `130b888011a5c5286b818078410d2b022faf60d0033f20d0e89d628098c0ad50`; adjacent scanner summary, raw JSON, probe and identity evidence are bound by that report and must be verified before use.
- Prior decisions remain current unless durable contradictory evidence is found: approved Techplan, exact five-task split, and G1–G3 Human authorization. This Run changes none of them by invocation.
- Target repo HEAD observed at preparation: `59f2b25940fa4baf243b4f1d2ca2ba04c9b17fd4`, branch `pilot/3-c1-engineering`.
- Harscode workflow HEAD observed at preparation: `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`, working tree clean.
- Root `AGENTS.md`, `backend/AGENTS.md`, active Techplan, task/manifest, current C1 handoff/read order, and active shared API contract are project authorities to recheck. Product and pre-engineering artifacts are read-only.

## Orchestrated bindings and runtime route

- **Project root / CWD:** `/home/anhar-solehudin/kencleng-workspace/kencleng`
- **Harscode workspace root:** `../harscode-workspace`
- **Work Unit record:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md`
- **WORK_UNIT_PATH / TASK_PATH compatibility root:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`
- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TECHPLAN-DEP-001`
- **ARTIFACT_TARGET:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.candidate.md`; conditionally the sibling `report-techplan.md` only at the exact converged Human approval gate after applicable review/resolution.
- **PRIOR_ARTIFACTS:** Approved predecessor, T2 task, manifest, T2 Build report and its bound evidence, current project authorities, and canonical Harscode Techplan authorities.
- **Communication language:** Bahasa Indonesia; canonical terms/enums and technical identifiers remain unchanged.
- **Runtime harness:** `codex-cli`; Human-Assisted mechanical dispatch.
- **Selected model / reasoning:** `gpt-6-sol` / `medium` — approved for this Run; not dispatched.
- **MODEL_APPROVAL:** APPROVED by Anhar on 2026-10-10 for `RUN-C1-ENG-TECHPLAN-DEP-001` only, in direct response to the Orchestrator request: “approve bro”. This is model dispatch approval only; no previous Run approval is inherited.
- **Model routing rationale:** The bounded work requires architecture/security cross-cutting analysis of Go toolchain and transitive auth dependency advisories, applicability, and preserved OIDC trust guarantees. Registry-declared non-gated Luna does not declare architecture/cross-cutting-analysis capability. Sol/medium is the lowest supported effort for this declared capability and same class used for prior C1 Techplan synthesis. Escalation is not preauthorized; if the inputs are insufficient, route the evidence gap rather than using a stronger model to replace missing authority.
- **CONTINUATION_CHECKPOINT:** None; fresh Run reconstruction.

## Execution envelope

- **PREAUTHORIZED:** Read current target/workflow authorities and live dependency/source evidence; perform read-only inspection and primary-source research; revise exactly the candidate target under the Techplan lifecycle; write Run-owned evidence under RUN_PATH; create the derived Human report only if this exact candidate reaches its proper approval gate after applicable review/resolution.
- **ORCHESTRATOR_DECISION:** Route newly discovered coordination gaps, review, or downstream readiness after terminal handoff. No Participant self-dispatch into another phase.
- **HUMAN_REQUIRED:** Approval to dispatch the gated model; exact candidate Techplan approval; any changed protected implementation permission or material risk acceptance. Existing G1–G3 approvals remain effective only for their already-approved surfaces and do not approve a Techplan successor.
- **Out of scope:** Production code/dependencies, migrations, tests/builds/runtime execution, upstream Product/pre-engineering writes, current-effective predecessor edits, split/task/manifest edits, orchestration projection writes by Planner, dependency commit, candidate promotion, Build/T3/T5 dispatch, and independent Review.

## Canonical prompt and terminal carrier

Use `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md` with `../harscode-workspace/workflow/orchestrated-run-overlay.md`; read the required Techplan `template.md`, `rules.md`, `guardrails.md` in full. Route stack/security authorities through `../harscode-workspace/best-practices/AGENTS.md`; inspect only matching indexes and primary best-practice files. Use current primary sources for advisories/toolchain/module releases.

Write the terminal carrier at `RUN_PATH/evidence/phase-handoff.md` with one `## Phase handoff`, Run/Participant/Session provenance, exact candidate content identity, Outcome, Result refs, Findings, Decision requests, Blockers, Open / unverified, Recommended continuation, and Context refs per the overlay. Keep source-reported, independently checked, assumed, deferred, and not-tested evidence distinct. Report the pointer to Orchestrator; do not self-route or dispatch another phase.

## Prepared state

- **Dispatch readiness:** READY_FOR_HUMAN_DISPATCH.
- **Participant dispatched:** No.
- **Run outcome:** Not started.
- **Next action:** Human mechanically dispatches this Run in a fresh Planner Session using `gpt-6-sol` / `medium`. This approval is model dispatch approval only; it is not Techplan approval, dependency selection, risk acceptance, or any change to G1–G3. Participant has not started.
