# Run Invocation — RUN-C1-ENG-BUILD-T2-004

## Identity and dispatch

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T2-004`.
- **Phase route / Role:** Build re-entry — T2 backend identity/session (G1) / Implementer.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-IMPLEMENTER-004` / `SESSION-C1-ENG-T2-IMPLEMENTER-004` (fresh).
- **Dispatch posture:** Human-Assisted; dispatched by Anhar after exact Run-specific approval; terminal report received and reconciled as BLOCKED.
- **Language:** Bahasa Indonesia; preserve canonical technical identifiers.
- **Escalation owner:** Anhar via Orchestrator.
- **Target / workflow revision:** `70f0ca000054001cbc7b1d1fd23e1e8e984e8395` / `91a97af8ba64f98b0a692a83fbff866a210c45c0`.
- **Target worktree:** Preserve existing Orchestrator reconciliation edits under `.harscode-spaces/` and `docs/product/pilot-3-control-tower.md`; they are outside this Run's artifact target.
- **Harness / working directory:** `codex-cli` / `/home/anhar-solehudin/kencleng-workspace/kencleng`.
- **Work Unit / Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001` / this Run directory.
- **ARTIFACT_TARGET:** T2/G1 backend only. Run evidence: this Run's `evidence/build-report.md`.

## Purpose and bound inputs

Fresh Build re-entry after `RUN-C1-ENG-BUILD-T2-003` ended BLOCKED before the graph probe and surfaced source-derived F-T2-003. The assignment is: (1) establish the approved D8 graph gate; then, only if it passes, (2) fix session expiry evaluation in `LockSession` to use time sampled after the row lock is acquired; (3) finish required focused T2/G1 checks. The approved Techplan §9 already specifies revalidation of session expiry/revocation/person under lock; this is implementation correction within existing G1 authority, not a new product/security policy choice.

Bound inputs (paths relative to project root):

- Approved Techplan `.../techplan/techplan.md`, SHA-256 `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79`.
- Current T2 snapshot `.../techplan/tasks/T2-backend-identity-session.md`, SHA-256 `dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37`.
- Accepted manifest `.../techplan/tasks/manifest.md`, SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.
- Approved Techplan report `.../techplan/report-techplan.md`, SHA-256 `18562231cc39a6c1f85ce7c7c20cc8e161d00a203a2b612f2f60117e02b2e611`.
- T2-003 terminal report `.../runs/RUN-C1-ENG-BUILD-T2-003/evidence/build-report.md`, SHA-256 `f87ea3ae133cb8279008b31dadd3cb501f9e50c31a56817313e3d5f42dd93853`.
- T2-002 terminal report SHA-256 `474e39ddd7358c4bf0f1c94a1cb217406cd81285046b06207a8cb4332e91bf93`; prior graph `6095e0b6d7d80d83c10411602246c2cf0c02b1d78888f5579f0350ad57823b9f`; prior scan `15ea30ec5eb124d75299be0109f632cd4884648464525b06d9d5bfa02d4bf4fe`; prior integrity `b4537ed75f533f993f371954de47e42a793b8e5b0587577de7e27fb3e50696bd`.
- Live baseline at preparation: target HEAD `70f0ca000054001cbc7b1d1fd23e1e8e984e8395`; `backend/go.mod` SHA-256 `2c089366f9b54e86da2993dadbeffe3cd7d3786a7b5f0535dd9163521e4db5d3`; `backend/go.sum` SHA-256 `979a350418d3d0e4906ea8af162cb69f8cfcfd38507eaa7d8657f6c7c5d7960f`; `backend/internal/platform/auth/store.go` SHA-256 `dfeed8714bc1881b9c94e8b9ac8f29ed2226f7396b1f2c789b0bc7dcfd0032b9`.
- Live `LockSession` anchor: `backend/internal/platform/auth/store.go:118–138`. T2-003 Finding: its `now` argument is compared against `ExpiresAt` only after `SELECT ... FOR UPDATE` returns, so a session may expire during the lock wait but remain accepted against stale pre-lock time. This is source-derived; no fake-clock test, real PostgreSQL contention, configured runtime exposure, or exploit reproduction exists.
- Exact approved D8 pins: Go 1.26.9; OIDC v3.16.0; OAuth2 v0.34.0; go-jose v4.1.4; pgx/v5 v5.9.2; explicit x/text v0.41.0. Current module still has pgx v5.8.0 and x/text v0.29.0.
- T1 contract prerequisite remains satisfied. Five-task manifest and G1–G3 authority are unchanged. T5 remains independent; T3 cannot consume the T2 session interface until its prerequisite is established.

## Required sequence and stop conditions

1. Verify every bound hash and inspect the actual target worktree. Re-open `backend/AGENTS.md`, Build prompt/guidelines/checklist/overlay, approved Techplan §§9–10/12, current T2 task, live auth/session anchors, and relevant DB composition. Preserve unrelated Orchestrator edits. If a bound input changed, stop and report the exact mismatch.
2. **Before tracked backend/dependency changes or G1 continuation**, create an isolated temporary probe from the current backend module inputs. Apply only the approved D8 pins in the probe; under Go 1.26.9 record the resolved graph, origins/tagged source/checksums, transitive delta/minimum Go versions, `go mod verify`, fresh full `govulncheck ./...` findings/traces, focused compile, pgx API/query-mode/config and SCRAM/OpaqueString source path. Reconcile each of GO-2026-4771/4772/5004/5970/6629 and disposition all new/surviving findings.
3. Stop before tracked writes, patch, or dependency commit if any prior affected range remains, a new affected dependency/path appears, scan/integrity/compatibility/compile evidence is incomplete or fails, or any new material dependency/risk/semantic choice is needed. Route the smallest Finding/Decision to Techplan/Human. Do not waive, suppress, substitute, or accept risk.
4. If and only if the isolated gate passes with no unresolved material choice, apply the same approved pins to the tracked backend and re-establish actual backend graph/source/integrity/fresh scan/compatibility/focused compile. Probe evidence alone cannot clear B-T2-002.
5. After the graph gate, make the bounded G1 correction: expiry must be evaluated using a current time sampled after the session `FOR UPDATE` lock is obtained (or an equivalent implementation that proves the same ordering). Do not weaken expiry/revocation/person semantics. Add and run focused session-lock checks that demonstrate a session expiring while waiting is rejected, plus the required T2 protocol-negative/person-upsert checks. Real PostgreSQL contention and configured auth behavior remain independent Testing evidence; do not claim those here unless actually executed.
6. No dependency or implementation commit until approved §10 gates pass and focused checks are complete. No G2/G3 expansion, T3/T5 dispatch, Product/API/Techplan/task/manifest edits, or unrelated work. Write one Build report with exact commands/results, changed-file identities, limitations, and exactly one structured `## Phase handoff`; state whether the Run stopped at graph gate or completed G1.

## Execution envelope

- **PREAUTHORIZED:** Read the bound inputs and live anchors; run focused graph/source/integrity/scan/compatibility/compile checks in the isolated probe; after gates pass, make the bounded G1 expiry correction and author/run focused tests; update only approved D8 pins; write this Run's report.
- **ORCHESTRATOR_DECISION:** Material live authority/contract drift; reconcile blockers after terminal report and decide subsequent routing.
- **HUMAN_REQUIRED:** Exact model approval before dispatch; any new material dependency/risk/security choice or residual-risk acceptance; unapproved G2/G3 surface; destructive/external action or deployment.
- **Forbidden:** Weaken session expiry, suppress/waive findings, select alternate dependency versions, commit with gates open, change upstream/shared API/product/Techplan/task/manifest/Harscode guidance, or dispatch another task.

## Model and Human-Assisted dispatch package

- **Selected model / effort:** `gpt-6-sol` / `medium`; chosen for security-sensitive session authorization and cross-cutting dependency/compatibility analysis. Registry requires explicit Human approval. T2-003 did not leave a separate durable model-approval record; no prior approval transfers.
- **Model approval:** APPROVED by Anhar in the current conversation for `RUN-C1-ENG-BUILD-T2-004`, `gpt-6-sol` / `medium`; approval was Run-specific. Runtime model/effort remains unverified independently, as stated in the terminal report.
- **Dispatch readiness:** DISPATCHED; terminal `BLOCKED` report reconciled. No approval is inferred for any later Run.
- **Compact kickoff prompt:** “Jalankan `RUN-C1-ENG-BUILD-T2-004` sebagai Implementer T2/G1. Ikuti Invocation: mulai dengan isolated graph gate untuk D8; kalau gagal atau ada finding/risk baru, stop. Jika lolos, verifikasi graph backend aktual, lalu perbaiki expiry `LockSession` agar memakai waktu setelah row lock didapat dan jalankan focused T2 checks. Jangan commit sebelum semua gate lolos. Tulis satu `evidence/build-report.md` dengan satu Phase handoff; jangan dispatch task lain.”
