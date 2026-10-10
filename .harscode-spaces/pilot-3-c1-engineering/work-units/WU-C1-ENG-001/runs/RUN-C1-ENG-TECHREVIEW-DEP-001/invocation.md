# Run Invocation — RUN-C1-ENG-TECHREVIEW-DEP-001

## Identity and routing

- **Work Unit:** `WU-C1-ENG-001`
- **Run:** `RUN-C1-ENG-TECHREVIEW-DEP-001`
- **Phase route:** Independent Techplan Review — bounded material successor for T2 dependency/toolchain reconciliation.
- **Role:** Reviewer
- **Specialization / Participant Profile:** Independent Techplan review; no reusable project Profile established.
- **Participant:** `PARTICIPANT-C1-ENG-DEPENDENCY-TECHREVIEWER-001` (reserved; instantiate on mechanical dispatch).
- **Session:** `SESSION-C1-ENG-DEPENDENCY-TECHREVIEWER-001` (fresh, independent context from Planner).
- **Session transition:** FRESH — independent reviewer/actor context required by the review prompt; Planner Run has terminated.
- **Dispatch posture:** Human-Assisted; prepared, not dispatched.
- **Human escalation owner:** Anhar via Orchestrator.

## Review target and completion boundary

Independently review exactly the post-Approval candidate Techplan revision bound below. It is a `techplan.candidate.md` material successor; do not review the predecessor as the target, rewrite either Techplan, create the Human report, approve/promote the candidate, or self-dispatch resolution/Build. Recheck the candidate hash before substantive review. If the candidate changes materially before completion, stop and follow the review prompt's re-grounding rule; do not attribute findings to a different revision.

The synthesis handoff recommends review because a security/toolchain baseline change spans Go standard-library advisories, auth transitive dependency pins, applicability, and preserved R9/R10 OIDC trust. Apply the review prompt's Step 0 independently, identify its current Complex criteria, resolve semantic Techplan section mapping from current template, and perform broad independent fidelity checking against every durable Exploration artifact and current applicable project/workflow authority.

Focus review on whether the candidate's exact Go 1.26.9 baseline and JOSE v4.1.4 / related pins are supported by the evidence and current primary sources; whether all 50 scanner IDs are represented without claiming scanner-reported metadata as independently verified facts; whether advisory applicability, reachability, conditional non-applicability, residual risk, and Build-owned fresh graph/scan obligations are clearly separated; whether fallback/stop conditions are executable; and whether the original OIDC trust semantics, product requirements, architecture/ownership, verification ownership, accepted five-task split, and G1–G3 boundaries remain unchanged. Verify high-risk/non-obvious claims against primary sources/live repo as appropriate. Do not perform a second architecture choice or turn scanner output into an exploitability claim. Finding format/classification is owned by the canonical independent Techplan Review prompt.

Complete only with the canonical review output and one structured terminal `## Phase handoff`, or stop at Step 0 if the prompt says review is not warranted. Recommend exact next route from findings. No Techplan approval or `B-T2-001` closure is established by a clean review; exact Human approval and bounded promotion remain separate gates. No tests, builds, dependency downloads, runtime probes, or production writes are authorized by this Run.

## Bound inputs and provenance

- **Exact review target:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.candidate.md`, SHA-256 `a034535b32688bc36fe36ec743d760ecc94c9daca9e751a11e3ad7eeef63c77c` at preparation.
- **Current-effective predecessor:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`, SHA-256 `bd732748b486f0327f6cba2cdf14cab32f1ece3c5be8ebb24f9659444947fa47`; remains Approved/current-effective during review.
- **Planner handoff:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TECHPLAN-DEP-001/evidence/phase-handoff.md`, candidate hash above, predecessor hash above.
- **Planner evidence:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TECHPLAN-DEP-001/evidence/dependency-reconciliation.md`, SHA-256 `10ffeff6f2705bcb49ec6e9029cd60a0054ee09aeb0aac7b77c1bcb5a1548d25`; evidence, not independent authority. Its primary source URLs are leads; verify claims directly.
- **Trigger:** `F-T2-001` / `B-T2-001` from T2 Build report `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-001/evidence/build-report.md`, SHA-256 `130b888011a5c5286b818078410d2b022faf60d0033f20d0e89d628098c0ad50`, plus its scanner JSON/summary/probe and identity manifest. Verify all identities before reliance.
- **Approved execution authorities:** parent Techplan predecessor, T2 task, accepted five-task manifest (SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`), G1–G3 Human approvals, root/backend AGENTS and routed Stage 7/5/6 authorities. They are read-only and unchanged by this Run.
- **Exploration:** read every durable file under `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-EXPLORATION-001/evidence/` as required by review prompt: `stage-2-gap-analysis.md`, `stage-3-solutioning.md`.
- **Target revision / branch:** `59f2b25940fa4baf243b4f1d2ca2ba04c9b17fd4`, `pilot/3-c1-engineering` at preparation.
- **Workflow revision:** `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`, Harscode working tree clean at preparation.

## Orchestrated bindings and runtime route

- **Project root / CWD:** `/home/anhar-solehudin/kencleng-workspace/kencleng`
- **Harscode workspace root:** `../harscode-workspace`
- **Work Unit record:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md`
- **WORK_UNIT_PATH / TASK_PATH compatibility root:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`
- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TECHREVIEW-DEP-001`

- **ARTIFACT_TARGET:** none; write review findings and terminal handoff under RUN_PATH/evidence only.
- **PRIOR_ARTIFACTS:** Exact candidate revision and predecessor above, Planner terminal/evidence, T2 report and bound scanner artifacts, Exploration evidence, accepted Techplan/task/manifest and routed project authorities.
- **Communication language:** Bahasa Indonesia; canonical terms/enums and technical identifiers remain unchanged.
- **Runtime harness:** `codex-cli`; Human-Assisted mechanical dispatch.
- **Selected model / reasoning:** `gpt-6-sol` / `medium` — approved for this Run; not dispatched.
- **MODEL_APPROVAL:** APPROVED by Anhar on 2026-10-10 for `RUN-C1-ENG-TECHREVIEW-DEP-001` only, in direct response to the Orchestrator request: “approve bro”. This is model dispatch approval only; no previous Run approval is inherited.
- **Model routing rationale:** Independent review must challenge a material security/toolchain successor, verify dependency/advisory claims against primary evidence, and preserve the auth contract. Registry-declared non-gated Luna does not declare architecture/cross-cutting-analysis capability. Sol/medium is the lowest supported effort for the needed declared capability. Reviewer context and Participant are fresh and independent from Planner. Missing evidence is reported as a finding/open item; stronger model is not a substitute for evidence.
- **CONTINUATION_CHECKPOINT:** None; fresh independent review Run.

## Execution envelope

- **PREAUTHORIZED:** Read the exact bound candidate and current relevant authorities; independently review using the canonical prompt; read all durable Exploration evidence; verify current claims and primary-source facts; write only review findings and terminal handoff under RUN_PATH/evidence.
- **ORCHESTRATOR_DECISION:** Reconcile review verdict and route resolution/report/Human gate after handoff. Reviewer does not self-dispatch another phase.
- **HUMAN_REQUIRED:** Approval to dispatch this gated model; exact Human approval of the Techplan candidate; any material risk acceptance or change to protected implementation authorization.
- **Out of scope:** Editing candidate/predecessor/report, changing task split, manifest, Product/pre-engineering authority, implementation/dependency writes, tests/builds/runtime probes, candidate promotion, blocker closure, Build/T3/T5 dispatch, or claiming C1 completion.

## Canonical prompt and terminal carrier

Use `../harscode-workspace/workflow/2-2-techplan-review-prompt.md` with `../harscode-workspace/workflow/orchestrated-run-overlay.md`. Read the current Techplan `template.md`, `rules.md`, and `guardrails.md` in full. Open diagram guidelines only if the candidate contains a diagram. The prompt requires broad independent reading of all Exploration evidence and primary/live-source spot checks; it owns the review sequence and finding classification.

Write review findings to `RUN_PATH/evidence/review-findings.md` and one structured `## Phase handoff` to `RUN_PATH/evidence/phase-handoff.md`, with exact reviewed candidate identity, gate, findings, Open / deferred, recommendation, session transition, and context pointers. Keep synthesis claims distinct from reviewer-verified facts. If the review prompt's Step 0 says review is not warranted, record that result and stop substantive review.

## Prepared state

- **Dispatch readiness:** READY_FOR_HUMAN_DISPATCH.
- **Participant dispatched:** No.
- **Run outcome:** Not started.
- **Next action:** Human mechanically dispatches this Run in a fresh independent Reviewer Session using `gpt-6-sol` / `medium`. Approval applies to this Run only; it does not approve the Techplan candidate or resolve `B-T2-001`. Participant has not started.
