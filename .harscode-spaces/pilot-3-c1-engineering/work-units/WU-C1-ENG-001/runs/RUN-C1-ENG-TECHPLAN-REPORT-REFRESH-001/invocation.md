# Run Invocation — RUN-C1-ENG-TECHPLAN-REPORT-REFRESH-001

## Identity and routing

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-TECHPLAN-REPORT-REFRESH-001`
- **Phase route / Role:** Approved-successor report refresh / Planner.
- **Specialization / Profile:** None; canonical Planner with bounded derived-report scope.
- **Participant / Session:** `PARTICIPANT-C1-ENG-REPORT-REFRESH-001` / `SESSION-C1-ENG-REPORT-REFRESH-001` (fresh; new Run).
- **Transition reason:** Fresh report-generation occurrence after promotion of an exact Human-approved Techplan successor; guardrails require the current report be regenerated from the approved successor.
- **Dispatch posture:** Human-Assisted; prepared, not dispatched.
- **Human owner:** Anhar via Orchestrator.

## Assignment and boundary

Regenerate `techplan/report-techplan.md` in full from current-effective Approved `techplan/techplan.md` only, following the report template and guardrail lifecycle rule. This refresh is required because the approved predecessor was replaced by its approved successor; the report used at the candidate approval gate identified the candidate path/status and is no longer the current-effective report. Bind the report to the exact current `techplan.md` hash below and identify its approved status/provenance accurately. Summarize the approved Go 1.26.9 / OIDC v3.16.0 / OAuth2 v0.34.0 / go-jose/v4 v4.1.4 plan baseline and the still-open Build evidence gate without claiming actual graph compatibility, clean scan, dependency commit, implementation, or runtime security. Include the independent review + mechanical N1 resolution history, exact Human approval, and approval boundary; no new decision may be introduced by the report.

Do not modify `techplan.md`, task files, manifest, upstream Product/pre-engineering artifacts, implementation, or Run evidence outside this Run. Do not close `B-T2-001`, dispatch Build, run tests/builds/dependency downloads/runtime probes, or claim C1 completion. If a required report fact is absent/ambiguous in the Approved Techplan, stop and route it; do not invent it. Complete with the exact generated report SHA-256 and one structured Phase handoff.

## Bound inputs and provenance

- **Current-effective source:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`, SHA-256 `bae2470f9038521b5d2d51fe4bc8de192fe5a9cf91bc0975831a3046fbf7a8ad`, status Approved.
- **Approval source identity:** exact candidate SHA-256 `2c64128e95a52773da024c2ec856bc5ba7a047539f2197a6d747513d378ae629`, approved by Anhar 2026-10-10; lifecycle promotion changed only status/approval-resolution/Open Items records entailed by that decision.
- **Prior gate report:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/report-techplan.md`, SHA-256 `d1d6dc231fd33af329e25ad57938fa3c06f0e977c0593140a87a29740794a5c5`; it represents the approved candidate at the prior Human gate and must be refreshed for current-effective source/path/status.
- **Review/resolution evidence:** Review findings `e6baad5300466c980151909d1beba54a371d25c3ccc30d9610f41a1cca640f8a`; review handoff `5c7700646b6c62f0fb78b089fe1ce0c937c9cc5ce5de6368133fa84d7db2490c`; Planner resolution handoff under `RUN-C1-ENG-TECHPLAN-DEP-RESOLVE-001/evidence/phase-handoff.md`.
- **Harscode sources:** `../harscode-workspace/workflow/2-techplan/report-template.md`, `guardrails.md`, `orchestrated-run-overlay.md`.
- **Target/workflow revision:** target `59f2b25940fa4baf243b4f1d2ca2ba04c9b17fd4`, branch `pilot/3-c1-engineering`; workflow `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

## Bindings and execution envelope

- **CWD / project root:** `/home/anhar-solehudin/kencleng-workspace/kencleng`
- **Harscode root:** `../harscode-workspace`
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`
- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TECHPLAN-REPORT-REFRESH-001`
- **ARTIFACT_TARGET:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/report-techplan.md` only.
- **Communication:** Bahasa Indonesia; canonical Harscode labels/technical identifiers retained.
- **Harness:** `codex-cli`, Human-Assisted, fresh Session.
- **Model / effort:** `gpt-6-luna` / `medium`; Human registry does not require per-Run approval.
- **PREAUTHORIZED:** Read bound source/report/required report guidance; regenerate the one derived report; write Run-owned evidence.
- **ORCHESTRATOR_DECISION:** Reconcile returned report identity and next route.
- **HUMAN_REQUIRED:** No new Techplan decision in this Run; Techplan approval already recorded for exact candidate. Any new material ambiguity is routed, not decided here.
- **Out of scope:** All production/contract/task/manifest/upstream authority writes, candidate promotion, dependency choice, blocker closure, tests/build/runtime, downstream dispatch.

## Terminal evidence

Write one structured `## Phase handoff` at `RUN_PATH/evidence/phase-handoff.md` with source hash, report hash, provenance, outcome, any finding/open item, current `B-T2-001` status, and recommended next route. Do not copy the full report into Run evidence.

## Prepared state

- **Dispatch readiness:** READY_FOR_HUMAN_DISPATCH.
- **Participant dispatched:** No.
- **Run outcome:** Not started.
- **Next action:** Human mechanically dispatches a fresh Planner Session with `gpt-6-luna` / `medium`; no model approval required.
