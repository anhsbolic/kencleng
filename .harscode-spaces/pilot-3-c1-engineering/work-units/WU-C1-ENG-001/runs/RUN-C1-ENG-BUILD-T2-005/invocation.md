# Run Invocation — RUN-C1-ENG-BUILD-T2-005

## Identity and dispatch

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T2-005`.
- **Phase route / Role:** Build re-entry — T2 backend identity/session (G1) / Implementer.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-IMPLEMENTER-005` / `SESSION-C1-ENG-T2-IMPLEMENTER-005` (fresh).
- **Dispatch posture:** Human-Assisted; Anhar approved the selected model for this exact Run and it was dispatched. Terminal report received; Orchestrator reconciliation is recorded separately.
- **Language:** Bahasa Indonesia; preserve canonical technical identifiers.
- **Escalation owner:** Anhar via Orchestrator.
- **Target / workflow revision:** `70f0ca000054001cbc7b1d1fd23e1e8e984e8395` / `91a97af8ba64f98b0a692a83fbff866a210c45c0`.
- **Target worktree:** Preserve all existing Orchestrator and diagnostic documentation edits. Do not commit.
- **Harness / working directory:** `codex-cli` / `/home/anhar-solehudin/kencleng-workspace/kencleng`.
- **Work Unit / Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001` / this Run directory.
- **ARTIFACT_TARGET:** T2/G1 backend only. Run evidence: this Run's `evidence/build-report.md`.

## Purpose and bounded re-entry delta

Fresh Build re-entry after T2-004 terminated BLOCKED at the isolated D8 scan because the sandbox could not load Go build-cache entries. The post-Run scanner diagnosis (preserved in T2-004 evidence) established that the selected Go scanner can complete when cache writes and vulnerability-database network access are available. This diagnosis reduces tooling uncertainty but did not complete the D8 gate, clear actual-graph advisories, or reopen the terminal T2-004 Run.

Assignment: complete the isolated approved-D8 gate with evidence for the selected graph, module origins/tagged source/checksums, transitive delta and Go minimums, integrity, full source findings/traces, compatibility/source paths, and per-ID disposition. Only if every isolated gate passes with no unresolved material choice may the Run verify the actual backend graph, apply approved pins, re-establish the required actual-graph evidence, correct `LockSession` to evaluate expiry using time sampled after acquiring the row lock, and run focused T2 checks. Stop at any failed or incomplete gate.

## Current-effective inputs

- Approved Techplan: `.../techplan/techplan.md`, SHA-256 `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79`.
- Current T2 snapshot: `.../techplan/tasks/T2-backend-identity-session.md`, SHA-256 `dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37`.
- Accepted manifest: `.../techplan/tasks/manifest.md`, SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.
- Approved Techplan report: `.../techplan/report-techplan.md`, SHA-256 `18562231cc39a6c1f85ce7c7c20cc8e161d00a203a2b612f2f60117e02b2e611`.
- T2-004 terminal report: `.../runs/RUN-C1-ENG-BUILD-T2-004/evidence/build-report.md`, SHA-256 `eff0a3b9c2f0392bbbcaf0f2aa33cf0b93d872afe3eb8b8573f6cef2a81ed051`.
- T2-004 Orchestrator reconciliation and preserved diagnosis: `.../runs/RUN-C1-ENG-BUILD-T2-004/evidence/orchestrator-reconciliation.md`; diagnosis SHA-256 `07c0330674fb1b43f1d15596e77c2af4a60e3a3260158221ce4c5b590175b78b`; actual-backend scan SHA-256 `99219c4bb8e965eb8601d86e3ecb13c437019b2497bd8fb733c8c05f9fd471e9`.
- Scanner runbook: `docs/project/go-scanner-runbook.md`, SHA-256 `a463e157e8519adfd1ea63c487a4e7546b55ca80cf2186221a6db1f4aacc12ae`. Use it as target-project execution guidance; it does not alter the Techplan or Run gates.
- T2-003 source finding evidence: `.../runs/RUN-C1-ENG-BUILD-T2-003/evidence/build-report.md`, SHA-256 `f87ea3ae133cb8279008b31dadd3cb501f9e50c31a56817313e3d5f42dd93853`.
- Prior actual graph/scan/integrity evidence: T2-002 report SHA-256 `474e39ddd7358c4bf0f1c94a1cb217406cd81285046b06207a8cb4332e91bf93`; graph `6095e0b6d7d80d83c10411602246c2cf0c02b1d78888f5579f0350ad57823b9f`; scan `15ea30ec5eb124d75299be0109f632cd4884648464525b06d9d5bfa02d4bf4fe`; integrity `b4537ed75f533f993f371954de47e42a793b8e5b0587577de7e27fb3e50696bd`.
- Live baseline from T2-004 preparation: HEAD `70f0ca000054001cbc7b1d1fd23e1e8e984e8395`; `backend/go.mod` SHA-256 `2c089366f9b54e86da2993dadbeffe3cd7d3786a7b5f0535dd9163521e4db5d3`; `backend/go.sum` `979a350418d3d0e4906ea8af162cb69f8cfcfd38507eaa7d8657f6c7c5d7960f`; `backend/internal/platform/auth/store.go` `dfeed8714bc1881b9c94e8b9ac8f29ed2226f7396b1f2c789b0bc7dcfd0032b9`. Recheck live state and all bound inputs before action; these are preparation-time identities, not a substitute for live inspection.
- Required Build guidance: current `workflow/3-build-prompt.md`, `workflow/3-build/guidelines.md`, `workflow/3-build/checklist.md`, and `workflow/orchestrated-run-overlay.md`. Workflow revision is recorded above. Follow root and `backend/AGENTS.md` current-effective instructions.

Approved exact D8 pins remain Go 1.26.9, OIDC v3.16.0, OAuth2 v0.34.0, go-jose v4.1.4, pgx/v5 v5.9.2, and explicit x/text v0.41.0. The actual backend currently retains pgx v5.8.0 and x/text v0.29.0. The post-Run diagnostic reports the approved D8 probe source scan completed with zero findings; this is only one piece of evidence, and is not a completed gate.

F-T2-003 remains source-derived: `LockSession` compares expiry against caller time sampled before `FOR UPDATE` returns. Real PostgreSQL contention, configured runtime exposure, and exploit reproduction remain unverified. Existing B-T2-002 (approved dependency acceptance) and B-T2-003 (session-lock correctness/T2 completion) remain open. B-T2-001 remains unclosed historical context. T3 cannot consume the T2 session interface; T5 remains independent.

## Required sequence and stop conditions

1. Recheck every bound hash and the live worktree. Reopen current `AGENTS.md`, `backend/AGENTS.md`, applicable Build guidance, approved Techplan §§9–10/12, current T2 task, and live auth/session and DB-composition anchors. Preserve unrelated Orchestrator edits. Stop on bound-input drift.
2. Before any tracked backend/dependency write or G1 continuation, run a fresh isolated probe from the current backend module inputs using only approved D8 pins and Go 1.26.9. Establish and record the complete resolved graph; module origin, tagged source and checksums; full transitive delta and minimum Go versions; `go mod verify`; fresh full `govulncheck ./...` findings and traces; focused compile; and pgx API/query-mode/config plus SCRAM/OpaqueString source paths. Reconcile GO-2026-4771/4772/5004/5970/6629 and every new/surviving finding. Follow `docs/project/go-scanner-runbook.md`; record exact cache, scanner, network/escalation context, command, exit, stdout and stderr. Inspect structured JSON findings, not only exit status.
3. Stop before tracked writes, implementation, or commit if any gate is failed or incomplete, an affected range remains without the approved disposition, a new affected path/finding appears, or a material dependency/risk/semantic choice is needed. Do not waive, suppress, substitute versions, or accept risk. Route the smallest Finding/Decision to Orchestrator/Techplan/Human.
4. Only if the complete isolated gate passes without unresolved material choice, update the actual backend to the exact approved pins, then re-establish actual graph/source/origin/integrity/fresh scan/compatibility/focused compile. Probe evidence alone cannot close B-T2-002.
5. Once the actual graph gate passes, make only the approved G1 correction: check expiry against current time sampled after acquiring the session `FOR UPDATE` lock (or equivalent ordering proof), preserving expiry/revocation/person semantics. Author and run the focused expired-during-lock-wait check and required T2 protocol-negative/person-upsert checks. Record that real PostgreSQL contention and configured auth behavior remain Testing evidence unless actually run under their authority.
6. Do not commit until approved §10 gates and focused checks pass. No G2/G3 expansion, T3/T5 dispatch, Product/API/Techplan/task/manifest/Harscode guidance edits, or unrelated work. Write one Build report at `evidence/build-report.md` with exact commands/results, changed-file identities, limitations, explicit verification-scope confirmation, and exactly one structured `## Phase handoff`; state whether the Run stopped at the graph gate or completed G1.

## Execution envelope

- **PREAUTHORIZED:** Read bound inputs/live anchors; conduct isolated D8 graph/source/integrity/scan/compatibility/compile checks; if all gates pass, apply only approved D8 pins to backend, perform actual-graph re-verification, make the bounded G1 expiry correction and author/run focused T2 checks; write this Run's report.
- **ORCHESTRATOR_DECISION:** Reconcile terminal findings/blockers and route a new Run after this Run ends; resolve deterministic state projections.
- **HUMAN_REQUIRED:** Exact model approval before dispatch; new material dependency/risk/security choice or residual-risk acceptance; unapproved G2/G3 work; destructive/external action/deployment.
- **Forbidden:** Weaken expiry/session semantics, suppress/waive findings, select alternate dependency versions, commit with gates open, modify upstream/shared API/Product/Techplan/task/manifest/Harscode guidance, or dispatch another task.

## Model and Human-Assisted dispatch package

- **Selected model / effort:** `gpt-6-sol` / `medium`; security-sensitive session authorization and cross-cutting dependency/source compatibility remain in scope. Human-owned runtime registry requires explicit approval.
- **Model approval:** APPROVED by Anhar for `RUN-C1-ENG-BUILD-T2-005`, `gpt-6-sol` / `medium`; the approval is Run-specific. Report states runtime model/effort was not independently exposed.
- **Dispatch readiness:** DISPATCHED; terminal `BLOCKED` report received. No approval transfers to a later Run.
- **Compact kickoff prompt:** “Jalankan `RUN-C1-ENG-BUILD-T2-005` sebagai Implementer T2/G1. Ikuti Invocation dan scanner runbook. Lengkapi seluruh isolated D8 graph/source/integrity/scan/compatibility gate dulu; bila ada gate gagal/tidak lengkap atau finding/risk baru, stop dan laporkan. Hanya jika semua lolos, verifikasi graph backend aktual, terapkan pin D8 yang sudah disetujui, lalu koreksi expiry `LockSession` setelah row lock dan jalankan focused T2 checks. Jangan commit sebelum seluruh gate lolos. Tulis satu `evidence/build-report.md` dengan satu Phase handoff; jangan dispatch task lain.”
