# Run Invocation — RUN-C1-ENG-TECHPLAN-REPORT-REFRESH-002

## Identity and routing

- Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-TECHPLAN-REPORT-REFRESH-002`.
- Phase: Post-approval exact-source report refresh / Planner.
- Participant / Session: `PARTICIPANT-C1-ENG-REPORT-REFRESH-002` / `SESSION-C1-ENG-REPORT-REFRESH-002` (fresh).
- Model / effort: `gpt-6-luna` / `medium`; no registry approval required.
- Dispatch: Human-Assisted; ready, not dispatched.

## Assignment

Regenerate `techplan/report-techplan.md` in full from current-effective Approved `techplan/techplan.md` only, using the report template. Bind exact source SHA below. Summarize approved pgx/v5 v5.9.2 and explicit x/text v0.41.0 planned baseline, five advisory dispositions/evidence, independent Review/resolution history, exact Human approval, remaining Build gates, and approval boundary. Do not imply actual graph compatibility, clean scan, dependency commit, implementation, runtime security, blocker closure, or C1 completion. Do not edit Techplan/task/manifest, code/dependency files, upstream authority, or orchestration projections. No tests, builds, downloads, or runtime probes. Write one terminal Phase handoff with report hash and provenance.

## Bound inputs

- Current Approved source: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`, SHA-256 `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79`.
- Exact approved candidate SHA-256 `26c8948c9a1a8687b173fce5a473314dfc735ec47d2176a4f2fa82fee733a32b`; gate report SHA-256 `de5381eb0976aabd5127ea4f9014024d89404edfacb1a5bd5480e761f0413884`.
- Predecessor Techplan SHA-256 `bae2470f9038521b5d2d51fe4bc8de192fe5a9cf91bc0975831a3046fbf7a8ad`.
- Review findings `RUN-C1-ENG-TECHREVIEW-DEP-002/evidence/review-findings.md`, SHA-256 `eeb82a25b3f566fd16a983a9d1ff2d74d006d3f8a99103cd9dab4eab392385e2`; resolution `RUN-C1-ENG-TECHPLAN-DEP-RESOLVE-002/evidence/resolution.md`.
- Target/workflow revision `59f2b25940fa4baf243b4f1d2ca2ba04c9b17fd4` / `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

## Bindings

- RUN_PATH `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TECHPLAN-REPORT-REFRESH-002`.
- Artifact target: `techplan/report-techplan.md` only.
- Follow `../harscode-workspace/workflow/2-techplan/report-template.md`, `guardrails.md`, and `workflow/orchestrated-run-overlay.md`.
- PREAUTHORIZED: read source/guidance; regenerate report; write Run evidence.
- HUMAN_REQUIRED: none in this bounded refresh; approval already applies to exact candidate hash above.

## Prepared state

- Dispatch readiness READY_FOR_HUMAN_DISPATCH.
- Participant dispatched: No. Run not started.
- Human mechanically dispatches a fresh Planner Session; no model approval required.
