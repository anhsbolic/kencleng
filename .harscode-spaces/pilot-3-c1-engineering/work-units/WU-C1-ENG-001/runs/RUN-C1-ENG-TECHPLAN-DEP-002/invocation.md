# Run Invocation — RUN-C1-ENG-TECHPLAN-DEP-002

## Identity and routing

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-TECHPLAN-DEP-002`.
- **Phase route:** Techplan synthesis — bounded material successor reconciliation after fresh actual-graph pgx/x/text findings.
- **Role:** Planner; canonical Planner, no special profile.
- **Participant / Session:** `PARTICIPANT-C1-ENG-DEPENDENCY-PLANNER-002` / `SESSION-C1-ENG-DEPENDENCY-PLANNER-002` (fresh context; reserved, not dispatched).
- **Dispatch posture:** Human-Assisted; prepared, not dispatched.
- **Human owner:** Anhar via Orchestrator.
- **Communication:** Bahasa Indonesia; retain canonical technical identifiers.

## Assignment and completion boundary

Reconcile only the current Approved Techplan's dependency/toolchain baseline, security/advisory dispositions, and directly affected Build verification gates using the five open findings on the actual T2-002 graph. Current report IDs: GO-2026-4771, GO-2026-4772, GO-2026-5004, GO-2026-5970, GO-2026-6629. Inspect the full scanner output and the per-ID source traces before deciding applicability. Confirm current advisory pages and module release/compatibility facts from primary sources; record access date, affected/fixed ranges, source paths, application reachability limits, and residual risk. Distinguish scanner report, possible source path, trigger conditions, and demonstrated exploitability.

Because `techplan.md` is Approved, preserve it unchanged as predecessor and write exactly one material successor candidate at `techplan/techplan.candidate.md`. Changes are limited to affected pgx/x/text version pins or baseline constraints, supporting compatibility/security rationale, per-advisory dispositions, and directly affected §10/§12 Build gates / §13 Open Items. Do not assume a version upgrade is automatically correct: check the real module graph, Go version constraints, API compatibility, and fixed-version facts against primary evidence. Do not silently accept residual risk, suppress findings, or instruct Build to commit a selection that has not passed required checks. If source evidence cannot settle a material choice or introduces scope outside these boundaries, state the smallest Open Item/Decision and stop at that decision point.

Preserve unchanged the approved OIDC trust contract, C1 behavior/scope, active API, accepted five-task topology and manifest conditions, T1 prerequisite, G1–G3 permissions, and all non-dependency material requirements. Do not edit production files, `go.mod`/`go.sum`, T2 task, manifest, Solution Contract, Product/pre-engineering authority, current Approved predecessor, report, Control Tower, Space, WU, Events, invocation, or any artifact outside the candidate and this Run's evidence. Do not run builds/tests/runtime probes, commit dependencies, close blockers, promote a candidate, or dispatch another phase.

The candidate remains Draft/In Review. Recommend independent Techplan Review because this is a material security/supply-chain decision across Go and database/auth dependency paths. After review/resolution converges, a separate Planner occurrence generates the exact-candidate Human report; exact Human approval and bounded promotion are still required. A fresh T2 Build must then recheck the actual graph/source/integrity/scan/compatibility/compile and remaining findings before dependency commit. Partial uncommitted T2-002 implementation is not proof of behavior and must not be treated as accepted merely because this planning Run exists.

Complete with candidate hash, source/evidence hashes, inventory of allowed changes, unchanged contract/split/permissions confirmation, remaining Open Items, review recommendation, and one structured Phase handoff. Do not generate a report while synthesis is in churn.

## Bound inputs and provenance

- **Current Approved predecessor:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`, SHA-256 `bae2470f9038521b5d2d51fe4bc8de192fe5a9cf91bc0975831a3046fbf7a8ad`.
- **Current derived Human report:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/report-techplan.md`, SHA-256 `c8b59544f80f74bf64ec4d7aedb4babbd45f232a9d31cdc4663e7b3bbfb78779`.
- **Refreshed T2 task:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/T2-backend-identity-session.md`, SHA-256 `ba34963b725b25f78d908b2e72fceb66b26aa207701110fb0031ffc34affcd8d`.
- **Accepted manifest:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/manifest.md`, SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.
- **Trigger report:** `runs/RUN-C1-ENG-BUILD-T2-002/evidence/build-report.md`, SHA-256 `474e39ddd7358c4bf0f1c94a1cb217406cd81285046b06207a8cb4332e91bf93`.
- **Raw actual source scan:** `runs/RUN-C1-ENG-BUILD-T2-002/evidence/dependency-vuln.json`, SHA-256 `15ea30ec5eb124d75299be0109f632cd4884648464525b06d9d5bfa02d4bf4fe`; actual graph `dependency-graph.txt`, SHA-256 `6095e0b6d7d80d83c10411602246c2cf0c02b1d78888f5579f0350ad57823b9f`; integrity evidence `dependency-integrity.txt`, SHA-256 `b4537ed75f533f993f371954de47e42a793b8e5b0587577de7e27fb3e50696bd`; changed-file identities `changed-file-identities.sha256`, SHA-256 `61f007bf9a6215f4b535ce38c579f95f1b994864e3f8e5f82ceb75989fcc2349`.
- **Relevant advisory facts from Build report (verify current primary sources):** pgx v5.8.0 findings first fixed at v5.9.0 (4771/4772), v5.9.2 (5004); x/text v0.29.0 findings first fixed at v0.39.0 (5970) and v0.41.0 (6629). Build traces suggest potential paths, not exploit proof. Actual environment/query protocol and configured database-auth input conditions remain unverified.
- **Prior blocker evidence:** T2-001 report SHA-256 `130b888011a5c5286b818078410d2b022faf60d0033f20d0e89d628098c0ad50`; retain distinction between its earlier Go/jose findings and these new actual-graph IDs.
- **Target / workflow revision:** target HEAD `59f2b25940fa4baf243b4f1d2ca2ba04c9b17fd4`, branch `pilot/3-c1-engineering`; workflow revision `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.
- **Required authorities:** current root/backend AGENTS; Approved Techplan and refreshed T2 task/manifest; T2-002 Run evidence; `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`, `workflow/2-techplan/{template.md,rules.md,guardrails.md}`, `workflow/orchestrated-run-overlay.md`; `best-practices/AGENTS.md` and only matching supply-chain/Go/PostgreSQL/security primary authorities; current Go/pkg advisory pages and upstream pgx/x-text release/module metadata.

## Bindings and execution envelope

- **Project root / CWD:** `/home/anhar-solehudin/kencleng-workspace/kencleng`; Harscode root `../harscode-workspace`.
- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TECHPLAN-DEP-002`.
- **ARTIFACT_TARGET:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.candidate.md` only; one logical successor to the Approved plan.
- **Harness:** `codex-cli`, Human-Assisted, fresh Planner Session.
- **Selected model / effort:** `gpt-6-sol` / `medium`; new Run-specific Human approval required by the registry.
- **Routing rationale:** new actual-graph advisories span database protocol parser/query behavior and x/text normalization/DB startup auth paths, requiring cross-cutting security/applicability, compatibility, and authority analysis. Sol is registry-declared for strong reasoning/architecture/cross-cutting analysis; medium is the lowest supported effort for this bounded task.
- **MODEL_APPROVAL:** APPROVED by Anhar on 2026-10-10 for `RUN-C1-ENG-TECHPLAN-DEP-002` only. Prior `gpt-6-sol` approvals do not transfer.
- **PREAUTHORIZED:** Read bound artifacts/live anchors/current primary sources; inspect actual working graph/source read-only; revise only candidate; write Run-owned evidence.
- **ORCHESTRATOR_DECISION:** Reconcile terminal candidate evidence and route independent Review/resolution/report gate.
- **HUMAN_REQUIRED:** Exact model approval before dispatch; later exact candidate Techplan approval; any changed risk acceptance or material authorization. Existing G1–G3 decisions remain effective and unchanged.
- **Out of scope:** production dependency/code/test writes, dependency commit, current-effective Techplan edits, task/manifest/topology changes, Product/pre-engineering writes, API changes, blocker closure, candidate promotion, Build/Testing/runtime, downstream dispatch.

## Terminal carrier

Write one structured `## Phase handoff` to `RUN_PATH/evidence/phase-handoff.md`. Include exact candidate identity and allowed-delta check; source/evidence identities; per-advisory outcome; unresolved decisions; current B-T2-001/B-T2-002 state; review/report/Human gate route; verification limits; and context pointers. Keep human-readable source findings and model/session provenance precise; do not claim identity beyond what runtime exposes.

## Prepared state

- **Dispatch readiness:** READY_FOR_HUMAN_DISPATCH.
- **Participant dispatched:** No.
- **Run outcome:** Not started.
- **Next action:** Human mechanically dispatches fresh Planner Session using this invocation and `gpt-6-sol` / `medium`. No continuation to another phase from the same Run.
