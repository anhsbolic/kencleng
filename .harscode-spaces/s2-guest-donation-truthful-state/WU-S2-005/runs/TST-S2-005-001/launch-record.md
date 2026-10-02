# Launch Record — `TST-S2-005-001`

## Actual launch

- Run / Work Unit: `TST-S2-005-001` / `WU-S2-005`.
- Dispatch date: 2026-10-01.
- Launcher: Human-assisted Codex session after the user's request to run the durable Invocation; session ID not exposed.
- Participant ID / profile: `P-S2-005-TST-001-1` / `KC-VERIFIER`.
- Model / reasoning: Invocation configures `gpt-6-luna` / `medium`; active runtime model and effort were not independently exposed.
- Runtime / working directory: codex-cli / Kencleng repository root.
- Target revision: `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree; six assigned target files inspected.
- Workflow revision: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`.
- Session transition: FRESH independent Verifier after completed Build and Code Review; prior Participant context was not reused.
- Actual verification commands: `cd api && npm run validate`; temporary `redocly lint` against current and HEAD split sources; temporary `redocly bundle` and `openapi-typescript`; `cmp` against current generated files; `git diff --check` scoped to six assigned files; scoped `rg` consumer/value inspection.
- Source preservation: generation outputs were written under `/tmp/kencleng-tst-s2-005-Lm0pdI`; current authored source, aggregate bundle, generated TypeScript, fixture, and other production files were not modified by this Run.

## Outcome

- Verdict: Pass with flagged follow-ups.
- Produced: `testing-report-1.md`; no patch plan because no correction to the assigned artifact diff was required.
- Outstanding: final authored Campaign/API acceptance; downstream predicate source fidelity and runtime GET/POST/security/D1 evidence. Build's warning count 122 was not reproduced; current and HEAD baseline both yielded 124 warnings with no warning pointer/rule additions or removals.
- Next route: Orchestrator facilitates the Human contract acceptance and keeps producer/runtime evidence assigned to Campaign producer delivery and WU-S2-003 Testing.
