# Run Invocation — RUN-C1-ENG-BUILD-T2-006

## Identity and dispatch

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T2-006`.
- **Phase route / Role:** Build re-entry — T2 backend identity/session (G1) / Implementer.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-IMPLEMENTER-006` / `SESSION-C1-ENG-T2-IMPLEMENTER-006` (fresh).
- **Dispatch posture:** Human-Assisted; READY_FOR_HUMAN_DISPATCH. Prepared, not dispatched.
- **Selected model / effort:** `gpt-6-sol` / `medium`.
- **Model approval:** Anhar explicitly approved this exact model/effort for `RUN-C1-ENG-BUILD-T2-006` on 2026-10-11 (“setuju bro”). This approval is Run-specific; it does not establish runtime model identity or authorize another Run.
- **Language / escalation owner:** Bahasa Indonesia; Anhar via Orchestrator.
- **Target / workflow revision:** `059d145d5df64fcda427e5a625486a8e95ce1fd8` / `91a97af8ba64f98b0a692a83fbff866a210c45c0`.
- **Target worktree:** Preserve existing Orchestrator/control-tower edits. Do not commit.
- **Harness / working directory:** `codex-cli` / `/home/anhar-solehudin/kencleng-workspace/kencleng`.
- **Work Unit / Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001` / this Run directory.
- **ARTIFACT_TARGET:** T2/G1 backend only. Run evidence: this Run's `evidence/build-report.md`.

## Purpose and changed premise

Re-enter T2/G1 after terminal T2-005 and its separate, authorized cache-first trial. The terminal T2-005 outcome remains BLOCKED; the trial is evidence, not a reopened Build or acceptance. The independent Loop Health Audit `RUN-C1-ENG-LOOPHEALTH-002` returned `HEALTHY_CONVERGENCE`; Orchestrator reconciled it and selected this smallest normal route.

The changed premise is concrete: cache-first acquisition worked using module inputs recovered in T2-005; the exact approved D8 probe graph resolved; integrity and focused compile passed; and a fresh full source scan found zero finding objects for the selected pins. However, upstream tag-to-commit validation remains incomplete for several unchanged legacy modules. This Run must complete the full isolated D8 evidence, including that provenance gap, before any tracked backend/dependency write. Then, only if every gate passes, verify the actual backend graph and complete the already-authorized G1 correction.

## Current-effective authority and evidence

- Work Unit current state at approved dispatch package, SHA-256 `6202438fdccd9c47c24d9c1949fa737deee8974c5c8bbcb1b2925d1a5afc5fe9`: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md` (use current-state fields and latest reconciliation, not stale chronology).
- Events at approved dispatch package, SHA-256 `b4888945c906b26408e1035257a6bf6fef670cffaa87256bad2b0510872d3ea8`: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/events.md` (use material T2 lifecycle entries).
- Approved Techplan: `techplan/techplan.md`, SHA-256 `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79`.
- Current T2 task: `techplan/tasks/T2-backend-identity-session.md`, SHA-256 `dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37`.
- Accepted manifest: `techplan/tasks/manifest.md`, SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.
- Approved Techplan gate report: `techplan/report-techplan.md`, SHA-256 `18562231cc39a6c1f85ce7c7c20cc8e161d00a203a2b612f2f60117e02b2e611`.
- T2-004 terminal report: `runs/RUN-C1-ENG-BUILD-T2-004/evidence/build-report.md`, SHA-256 `eff0a3b9c2f0392bbbcaf0f2aa33cf0b93d872afe3eb8b8573f6cef2a81ed051`.
- T2-004 scanner diagnosis: `runs/RUN-C1-ENG-BUILD-T2-004/evidence/scanner-diagnosis/diagnosis.md`, SHA-256 `07c0330674fb1b43f1d15596e77c2af4a60e3a3260158221ce4c5b590175b78b`.
- T2-005 terminal report: `runs/RUN-C1-ENG-BUILD-T2-005/evidence/build-report.md`, SHA-256 `c0bb8330fd984afffe96f8eaecb88612511f42ed34b8483bfab53e8f96f4ff37`.
- T2-005 Orchestrator reconciliation: `runs/RUN-C1-ENG-BUILD-T2-005/evidence/orchestrator-reconciliation.md`, SHA-256 `21abe0f80365980275443127399ca3d1ab3e128c3dfa15c6b7806f5d38c4dcef`.
- T2-005 separate cache-first trial: `runs/RUN-C1-ENG-BUILD-T2-005/evidence/targeted-d8-trial/trial-report.md`, SHA-256 `ff6d6cf3e53651699c5d7b3eb25cf53cf902947fad641901076d5d5b359cacd1`.
- Loop Health Audit: `runs/RUN-C1-ENG-LOOPHEALTH-002/evidence/loop-health-audit.md`, SHA-256 `51e1d9c004ce05af37a4bb0c9ce66fa49db55e6424ec0318097c99889ce9051e`.
- Orchestrator's audit reconciliation: `runs/RUN-C1-ENG-LOOPHEALTH-002/evidence/orchestrator-reconciliation.md` (to be written before this Invocation is dispatched).
- Current root `AGENTS.md`, SHA-256 `e18ca8bc35601dad839f23b8728b816f5ad1cb25cdd9775ab55c661623474b6e`; `backend/AGENTS.md`, SHA-256 `297274385cbbe91ccb498b30edabe540acd8c8108a0a3c4ada4f7b336818d8b3`; and scanner runbook `docs/project/go-scanner-runbook.md`, SHA-256 `a463e157e8519adfd1ea63c487a4e7546b55ca80cf2186221a6db1f4aacc12ae` apply.
- Current Build prompt `../harscode-workspace/workflow/3-build-prompt.md`, SHA-256 `8e8a183a839c27d887212b1e51bbd96d2e796febed8d2685534d98fe1b3ad4e5`; Build guidelines `../harscode-workspace/workflow/3-build/guidelines.md`, SHA-256 `1a85acfa5ebc80837737038198154039d0039bb2c4d6d4addcb62ce34dac9546`; Build checklist `../harscode-workspace/workflow/3-build/checklist.md`, SHA-256 `8cb373825edd6e64e2df0e523623bf3b3d6e14fc03580027959e939f9a5ff98a`.
- Orchestrated Run overlay `../harscode-workspace/workflow/orchestrated-run-overlay.md`, SHA-256 `f806673aa5e7d1d82dae9f433ec54c3562e15ce778b709a141c57d3d22dfaf87`; Harscode Run contract `../harscode-workspace/orchestration/run-contract.md`, SHA-256 `4610c20ce2670913870718f9886e22c3c00bdf9da5290406a4a4450a818069aa`; orchestration `AGENTS.md`, SHA-256 `98cdc285f7dcd81284ea6d8a1ea23f938a69b8938bb4bc863a3f59aa534113cc`.
- Human-owned runtime/model registry `.harscode-spaces/.local-config.yaml`, SHA-256 `ffbd0521361eac90e82dbee98ba59f01f3e62641afd79bf8e71c4fc8a0e633ce` (read-only).
- Reopen all bound guidance and live code anchors before action; this Invocation does not replace them.
- Preparation-time baseline: target HEAD `059d145d5df64fcda427e5a625486a8e95ce1fd8`; `backend/go.mod` SHA-256 `2c089366f9b54e86da2993dadbeffe3cd7d3786a7b5f0535dd9163521e4db5d3`; `backend/go.sum` SHA-256 `979a350418d3d0e4906ea8af162cb69f8cfcfd38507eaa7d8657f6c7c5d7960f`; `backend/internal/platform/auth/store.go` SHA-256 `dfeed8714bc1881b9c94e8b9ac8f29ed2226f7396b1f2c789b0bc7dcfd0032b9`. These are preparation-time identities; recheck live state and all bound inputs before action.

Approved exact D8 pins remain Go 1.26.9, OIDC v3.16.0, OAuth2 v0.34.0, go-jose v4.1.4, pgx/v5 v5.9.2, and explicit x/text v0.41.0. The actual backend baseline recorded by T2-005 retains pgx v5.8.0 and x/text v0.29.0. No trial result accepts risk, closes B-T2-002, or proves actual-backend compatibility. F-T2-003/B-T2-003 remains source-derived and uncorrected; real PostgreSQL contention, configured auth behavior, and independent Testing evidence remain unverified.

## Required sequence and stop conditions

1. Recheck every bound hash and the live worktree. Reopen root/backend `AGENTS.md`, applicable Build guidance, approved Techplan §§9–10/12, current T2 task, scanner runbook, and live auth/session and DB-composition anchors. Preserve unrelated Orchestrator edits. Stop on bound-input drift, changed authority, or a material contradiction.
2. Before tracked writes, repeat the isolated approved-D8 gate using the demonstrated cache-first method. Establish the complete resolved graph; module origins, tagged source and checksums, including the outstanding unchanged-legacy provenance; full transitive delta and Go minimums; `go mod verify`; fresh full `govulncheck ./...` findings and traces; focused compile; and pgx API/query-mode/config plus SCRAM/OpaqueString source paths. Inspect structured JSON finding objects, not only scanner exit status. Record precise command, cache/scanner/network/escalation context and outputs per the runbook. Reconcile each required advisory ID and every new/surviving finding.
3. Stop before tracked writes, implementation or commit if any isolated gate is failed or incomplete, any origin/provenance remains unsupported, an affected range lacks the approved disposition, a new finding/path appears, or a material dependency/risk/semantic choice is needed. Do not waive, suppress, substitute versions, or accept risk; report the smallest Finding/Decision to Orchestrator/Techplan/Human.
4. Only if the complete isolated gate passes without unresolved material choice, update the actual backend to the exact approved pins and re-establish actual graph, source/origin, integrity, fresh scan, compatibility and per-ID evidence. The isolated probe/trial cannot close B-T2-002.
5. Once actual-graph acceptance passes, make only the authorized G1 correction: evaluate `LockSession` expiry using time sampled after acquiring the session row lock, preserving expiry/revocation/person semantics. Add and run the focused expired-during-lock-wait check plus required T2 protocol-negative/person-upsert checks. Record PostgreSQL contention and configured auth/runtime as unverified unless actually established under their owner.
6. Do not commit until every applicable approved §10 gate and focused check passes. Keep dependency and G1 changes within approved scope; no unrelated edits. Report changed-file hashes, commands/results, gate-by-gate dispositions, tests, blockers and deferred independent verification in exactly one terminal Build report with one structured Phase handoff.

## Required output and dispatch

- Write exactly this Run's `evidence/build-report.md`, including one terminal `## Phase handoff`.
- Stop and hand off BLOCKED at the first failed/incomplete gate or material decision. Do not dispatch another Run from this Run.
- Keep `B-T2-002` open until actual-backend dependency evidence and required dispositions are complete. Keep `B-T2-003` open until the bounded correction and focused evidence are accepted. T3 remains unavailable until the manifest condition is durably met; T5 remains independent.
- Human may launch the fresh Implementer Participant using this Invocation. No launch has occurred.
