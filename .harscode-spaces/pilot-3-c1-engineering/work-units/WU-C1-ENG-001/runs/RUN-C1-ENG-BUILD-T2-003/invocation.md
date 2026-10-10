# Run Invocation — RUN-C1-ENG-BUILD-T2-003

## Identity and dispatch

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T2-003`.
- **Phase route / Role:** Build — T2 backend identity/session boundary (G1), graph-first re-entry / Implementer.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-IMPLEMENTER-003` / `SESSION-C1-ENG-T2-IMPLEMENTER-003` (fresh).
- **Dispatch posture:** Human-Assisted; Run dispatched and terminal BLOCKED, as established by returned `evidence/build-report.md` SHA-256 `f87ea3ae133cb8279008b31dadd3cb501f9e50c31a56817313e3d5f42dd93853`.
- **Communication:** Bahasa Indonesia; preserve canonical Harscode terms and technical identifiers.
- **Escalation owner:** Anhar via Orchestrator.
- **Target / workflow revision observed:** `70f0ca000054001cbc7b1d1fd23e1e8e984e8395` / `91a97af8ba64f98b0a692a83fbff866a210c45c0`.
- **Target tree at preparation:** working tree has Orchestrator reconciliation edits in the WU, Space README, T2 provenance header, and Control Tower. Preserve these files; they are outside `ARTIFACT_TARGET`.
- **Harness / CWD:** `codex-cli` / `/home/anhar-solehudin/kencleng-workspace/kencleng`.
- **Work Unit / Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001` / this Run directory.
- **ARTIFACT_TARGET:** `backend/` T2/G1 implementation and focused tests only, under the exact gates below. Run-owned evidence goes in this Run's `evidence/`.

## Purpose and current-effective inputs

Fresh Build re-entry after `RUN-C1-ENG-BUILD-T2-002` and the reconciled `LOOP_RISK` audit. First establish whether the approved D8 Go/dependency baseline is the actual resolved, compatible, source-reviewed graph. Do not continue G1 implementation or accept/commit dependencies until the graph-first gate is established.

Bound inputs (paths relative to project root):

- Approved Techplan: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`, SHA-256 `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79`.
- Current T2 snapshot: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/T2-backend-identity-session.md`, SHA-256 `dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37`.
- Accepted task manifest: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/manifest.md`, SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.
- Current Techplan report: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/report-techplan.md`, SHA-256 `18562231cc39a6c1f85ce7c7c20cc8e161d00a203a2b612f2f60117e02b2e611`.
- Prior Build report: `.../runs/RUN-C1-ENG-BUILD-T2-002/evidence/build-report.md`, SHA-256 `474e39ddd7358c4bf0f1c94a1cb217406cd81285046b06207a8cb4332e91bf93`.
- Prior actual graph, scan, and integrity evidence: `.../RUN-C1-ENG-BUILD-T2-002/evidence/dependency-graph.txt` SHA-256 `6095e0b6d7d80d83c10411602246c2cf0c02b1d78888f5579f0350ad57823b9f`; `dependency-vuln.json` SHA-256 `15ea30ec5eb124d75299be0109f632cd4884648464525b06d9d5bfa02d4bf4fe`; `dependency-integrity.txt` SHA-256 `b4537ed75f533f993f371954de47e42a793b8e5b0587577de7e27fb3e50696bd`.
- Loop Health verdict: `.../runs/RUN-C1-ENG-LOOPHEALTH-001/evidence/loop-health-audit.md`, SHA-256 `0bd8c28322eac3434aa6720d74188b2b7115d1da62ae6180a3b79a3e5415a059` (`LOOP_RISK`, not `STALLED`).
- Current task-level authority and dependency requirements are the Approved Techplan §10/§12 and T2 snapshot. The exact authorized D8 pins are Go 1.26.9, OIDC v3.16.0, OAuth2 v0.34.0, go-jose v4.1.4, pgx/v5 v5.9.2, and explicit x/text v0.41.0.
- T1 contract prerequisite remains satisfied by the current reviewed source/bundle/types hashes in the T2 task. Accepted five-task topology and G1–G3 authorization remain unchanged.

## Live baseline and drift checks

At preparation, the target repository HEAD is `70f0ca000054001cbc7b1d1fd23e1e8e984e8395`; Harscode HEAD is `91a97af8ba64f98b0a692a83fbff866a210c45c0` and its working tree is clean. Live `backend/go.mod` SHA-256 is `2c089366f9b54e86da2993dadbeffe3cd7d3786a7b5f0535dd9163521e4db5d3`; `backend/go.sum` SHA-256 is `979a350418d3d0e4906ea8af162cb69f8cfcfd38507eaa7d8657f6c7c5d7960f`. Current go.mod has Go 1.26.9 and the approved OIDC/OAuth2/jose pins, but pgx/v5 v5.8.0 and x/text v0.29.0, so it does not yet reflect the approved complete D8 baseline.

The T2-002 report says its partial implementation/dependency changes were uncommitted at Run termination. Current Git evidence shows the auth source, migration, server composition, and current go.mod/go.sum are tracked from commit `f47ae36` (which is an ancestor of current HEAD); the current worktree has no backend/go.mod/go.sum changes. This history/status mismatch does not establish Build acceptance, correctness, or dependency clearance. Reopen and hash the actual live anchors before action; preserve the source, report the discrepancy with exact evidence, and do not infer authorship or validation from Git tracking.

If any bound durable input hash differs, stop before implementation/dependency changes and report the mismatch. Inspect only the current T2/G1 live anchors required by the task and `backend/AGENTS.md`; preserve the unrelated Orchestrator reconciliation edits.

## Graph-first gate (must precede backend writes)

1. Re-ground `backend/AGENTS.md`, `backend/cmd/server/main.go`, `backend/internal/platform/db/db.go`, `backend/go.mod`, `backend/go.sum`, relevant migrations/auth source, and the active API anchors. Record current revisions/hashes and inspect the actual working tree.
2. Establish the exact approved D8 candidate graph under Go 1.26.9 without first modifying tracked backend files. Use an isolated temporary copy/probe based on the live backend module inputs; apply only the exact approved pgx v5.9.2 and x/text v0.41.0 selections plus existing approved pins. Preserve commands, module origins, checksums/source and transitive delta; do not use floating latest, replacements, exclusions, guessed pins, or weaken OIDC trust.
3. In the isolated candidate, record actual `go list -m all` graph, relevant tagged source/release/API compatibility and minimum Go versions, `go mod verify`, fresh full source scan (`govulncheck ./...` with full findings/traces), and focused compile. Reconcile each of the five IDs from T2-002, verify their affected ranges are absent, and disposition every new/surviving finding. Inspect query-mode/config exposure and pgx SCRAM/OpaqueString path without secrets; do not present a symbol trace or fixture as configured runtime evidence.
4. **Stop before any tracked backend/dependency write, G1 continuation, or dependency commit** if the five ranges remain, a new affected dependency/path appears, compile/compatibility fails, evidence is incomplete, or any material risk/dependency choice is needed. Return the smallest Finding/Decision to the owning Techplan/Human gate. Do not accept residual risk or choose another version.
5. If and only if the isolated gate passes with no unresolved material choice, apply the same approved pins to the tracked backend module; re-record the actual backend graph/source/integrity/fresh scan/compatibility and focused compile before continuing existing G1 work. The actual graph, not the probe alone, is the acceptance evidence. Before any dependency commit, complete all approved §10 checks and authored focused protocol-negative, person-upsert, and session-lock/revalidation tests. If any gate fails, stop before commit and preserve precise evidence.

After the graph gate, continue the bounded T2/G1 Build against the approved task: review existing T2 changes before reuse, author/run focused checks, and complete only G1. No T3, T5, G2/G3 expansion, or dependent dispatch. Record exact commands/results and changed-file identities; distinguish compile/fakes from real provider, PostgreSQL, browser, and Testing evidence. Do not run heavyweight race/performance/security-class sweeps or broad Testing-owned suites in this Build unless the Techplan requires the evidence for the specific graph gate, then explain the scope.

## Execution envelope

- **PREAUTHORIZED:** read current/bound inputs; run focused Go toolchain/module graph/source/integrity/scan/compatibility/compile checks; use an isolated temporary probe; after the isolated gate passes, write T2/G1 backend auth/config/person/session/migration code and focused tests under the approved task, update only exact approved dependency pins, and write this Run's evidence. Preserve all unrelated working-tree changes.
- **ORCHESTRATOR_DECISION:** any material live authority/contract/architecture drift; reconcile T3 prerequisite after terminal handoff; choose routing after findings.
- **HUMAN_REQUIRED:** exact Run model approval before dispatch; any new material dependency/risk/security choice or residual-risk acceptance, any unapproved G2/G3 surface, destructive/external action, or deployment.
- **Forbidden:** edit Product/pre-engineering authorities, shared API, frontend, Techplan/task/manifest, another Run, Harscode guidance, or unrelated files; commit dependencies or implementation while any approved gate is incomplete/open; introduce custom crypto, JWT/signing keys, password/MFA, broad role engine, or expand beyond G1.

## Build output and handoff

Follow `../harscode-workspace/workflow/3-build-prompt.md`, `3-build/guidelines.md`, `3-build/checklist.md`, and `orchestrated-run-overlay.md`, plus `backend/AGENTS.md`. Write exactly this Run's `evidence/build-report.md`; report exact graph, source/origins, transitive delta, checksums, scan findings/traces and per-ID dispositions, focused tests, changed-file hashes, limits, and one structured `## Phase handoff`. The terminal report must state explicitly whether the Run stopped at graph-first gate or continued G1, and that no heavyweight Testing sweep ran unless it actually did.

## Model and Human-Assisted dispatch package

- **Model / effort:** Invocation selected `gpt-6-sol` / `medium`; terminal report says runtime model/effort are not independently established. Registry marks Sol approval-required.
- **Model approval:** No separate durable exact Run-specific approval record is present in the Invocation or terminal report. The report records that a Human supplied the Run-specific execution kickoff; this establishes that the Run occurred, but is not silently converted into a model-approval record. Prior Loop Health approval does not transfer.
- **Dispatch / outcome:** Human-Assisted execution is established by the returned terminal report, with fresh Participant `PARTICIPANT-C1-ENG-T2-IMPLEMENTER-003` / Session `SESSION-C1-ENG-T2-IMPLEMENTER-003`; outcome BLOCKED at graph-first step 1. No model identity/effort beyond Invocation selection is independently asserted.
- **Dispatch readiness:** TERMINAL — BLOCKED. Any re-entry is a new Run with a fresh Participant/Session and a new exact model approval before dispatch.
- **Next route:** Orchestrator reconciles F-T2-003 and B-T2-003, then prepares a fresh T2/G1 Build. Preserve the isolated approved-D8 graph gate as the first gate; if it passes, repair the post-lock expiry check under the already approved §9/G1 behavior and run focused authored checks. Real PostgreSQL contention remains independent Testing evidence. B-T2-002 remains open until actual dependency evidence satisfies the approved gate.
- **Compact kickoff prompt (historical):** “Jalankan `RUN-C1-ENG-BUILD-T2-003` sebagai Implementer T2/G1. Baca Invocation dan artefak terikat, mulai dari graph-first gate memakai D8 yang sudah disetujui. Jangan ubah backend/dependency atau lanjutkan implementasi sebelum gate lolos; jika ada finding/risk baru atau mismatch, stop dan laporkan. Setelah lolos, verifikasi graph aktual lalu selesaikan hanya G1 dan tulis satu `evidence/build-report.md` dengan satu Phase handoff. Jangan dispatch task lain.”
