# Build report — RUN-C1-ENG-BUILD-T2-003

> Phase: Build; Author / Role: `PARTICIPANT-C1-ENG-T2-IMPLEMENTER-003` / Implementer T2/G1; Session: `SESSION-C1-ENG-T2-IMPLEMENTER-003`, fresh. Work Unit: `WU-C1-ENG-001`. Created/Updated: 2026-10-11. Target HEAD: `70f0ca000054001cbc7b1d1fd23e1e8e984e8395`; observed workflow HEAD: `91a97af8ba64f98b0a692a83fbff866a210c45c0`, workflow working tree clean. Invocation selects `gpt-6-sol` / medium; runtime model/effort not independently established. Invocation still records model approval PENDING; the Human supplied the Run-specific execution kickoff in this session. This report does not amend the dispatch registry or claim a separate model approval record.

## What changed

**STOPPED at graph-first gate step 1 (live-anchor inspection), before the isolated candidate probe. G1 did not continue.** New source-observed expiry revalidation risk F-T2-003 requires owning reconciliation under the Human's explicit stop-on-new-risk instruction. Only this report was written. No backend/dependency changes, dependency acceptance, implementation/dependency commit, upstream edits, or other task dispatch occurred.

All nine durable input hashes match Invocation. Target HEAD and live go.mod/go.sum also match. Existing Orchestrator edits in Space README, WU, events, T2 provenance, and Control Tower were preserved. Current pgx v5.8.0 / x/text v0.29.0 are the recorded starting state, not a new mismatch. Approved D8 remains Go 1.26.9, OIDC v3.16.0, OAuth2 v0.34.0, jose v4.1.4, pgx v5.9.2, explicit x/text v0.41.0; none were applied by this Run.

The previously recorded provenance discrepancy is confirmed: `git log -1 --format='%H %s' -- backend/internal/platform/auth backend/go.mod backend/go.sum backend/migrations backend/cmd/server/main.go` reports `f47ae36af34f0bd57412812bd0b34e18b7eecffc WU-C1-ENG-001 : T2 Plan`; `git merge-base --is-ancestor f47ae36 HEAD` exits 0. Current backend diff is empty. T2-002's bound terminal report describes these files as uncommitted at termination. Tracking now establishes neither authorship, acceptance, nor validation; the earlier report was preserved.

## Finding F-T2-003 — expiry time predates session lock acquisition

Owner: Orchestrator routes to the owning Techplan/G1 security concern and Anhar as needed. Scope: T2 exported authenticated session-lock interface, not the whole Work Unit.

Approved Techplan §9 runtime order requires rechecking session revocation/expiry/person under lock; T2 requires session-lock/revalidation checks before completion. In `backend/internal/platform/auth/store.go:120`, `LockSession` accepts `now time.Time`. Lines 126–128 acquire `FOR UPDATE` and scan the row; line 135 then compares expiry against that same caller-supplied timestamp. The function never samples time after lock acquisition. Current `RevokeSession` passes its pre-query `now` into this interface; T3 consumers are not implemented or validated here.

Concrete source-derived scenario: caller supplies t0, session expires at t1 > t0, row-lock wait finishes at t2 >= t1, and row is not revoked. The comparison `now.Before(session.ExpiresAt)` remains true because `now` is t0, so this interface can return the expired session as valid after the wait. Reading the updated revocation field does not refresh the expiry clock. The comment promising expiry revalidation after acquisition therefore overstates what the interface guarantees. A future confirmation consumer could rely on stale authority; no actual Organization grant or deployed exploit is claimed.

Verified: timestamp flow and comparison in live source. Assumed for the scenario: a valid token, unrevoked row and lock wait spanning expiry. Not tested: executable fake-clock reproduction, real PostgreSQL contention, configured runtime reachability, or exploitability. No repair or alternate clock/interface choice was made. Return the smallest revalidation Finding for owning disposition before a fresh continuation; do not weaken expiry semantics or interpret compile as clearance.

## Input and live identities

Paths below are relative to the WU unless prefixed `backend/` or `api/`. SHA-256 computed with Python `hashlib.sha256(Path(...).read_bytes())` and, for API, `sha256sum`.

```text
6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79  techplan/techplan.md
dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37  techplan/tasks/T2-backend-identity-session.md
2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d  techplan/tasks/manifest.md
18562231cc39a6c1f85ce7c7c20cc8e161d00a203a2b612f2f60117e02b2e611  techplan/report-techplan.md
474e39ddd7358c4bf0f1c94a1cb217406cd81285046b06207a8cb4332e91bf93  runs/RUN-C1-ENG-BUILD-T2-002/evidence/build-report.md
6095e0b6d7d80d83c10411602246c2cf0c02b1d78888f5579f0350ad57823b9f  runs/RUN-C1-ENG-BUILD-T2-002/evidence/dependency-graph.txt
15ea30ec5eb124d75299be0109f632cd4884648464525b06d9d5bfa02d4bf4fe  runs/RUN-C1-ENG-BUILD-T2-002/evidence/dependency-vuln.json
b4537ed75f533f993f371954de47e42a793b8e5b0587577de7e27fb3e50696bd  runs/RUN-C1-ENG-BUILD-T2-002/evidence/dependency-integrity.txt
0bd8c28322eac3434aa6720d74188b2b7115d1da62ae6180a3b79a3e5415a059  runs/RUN-C1-ENG-LOOPHEALTH-001/evidence/loop-health-audit.md
b530478422f0f5f73e588f89d9e07c7eeab74e00eeb44eb867c2f723183f8cb1  backend/AGENTS.md
d217dfd21898ce89d3eee96683b7512fd7010ec0e55da04c6d211bb5686e2246  backend/cmd/server/main.go
56492320609c7bc0acbbf0761bd6669a9f79fa31986e92a2ea88138798a63660  backend/internal/platform/db/db.go
2c089366f9b54e86da2993dadbeffe3cd7d3786a7b5f0535dd9163521e4db5d3  backend/go.mod
979a350418d3d0e4906ea8af162cb69f8cfcfd38507eaa7d8657f6c7c5d7960f  backend/go.sum
a9a89117ba682e3963f5a02ec526d566763e1f991994cd5526fb01c60016a812  backend/internal/platform/auth/oidc.go
dfeed8714bc1881b9c94e8b9ac8f29ed2226f7396b1f2c789b0bc7dcfd0032b9  backend/internal/platform/auth/store.go
bd6a7706c4e5b9ad9af69a33a537b9af0e28fa74a852672692a96f3d034b5009  backend/internal/platform/auth/http.go
cbb5b1794a3c5b98d44a7bc13661ee7c1977cd511115da6cda0ea5af7a48817e  backend/migrations/000001_identity_session.up.sql
3d291600e2992afea7d336f17287e7d98a3e0eb4351933b3d402a39be64eb804  backend/migrations/000001_identity_session.down.sql
ddefffbe588459cb0cba9f68dfedaccba3da8a1313f1956511cf628df7199965  api/openapi/index.yaml
38bb6c531804b54e9d5ce56d623c397a590722c56dcb7a3c12e646bb8280fa09  api/openapi.yaml
d21e8aa1eec92a95cc1fed87d976f726f59b48bd334aa8e463844d48e16cb80c  api/openapi.d.ts
```

These are unchanged live/input identities, not changed-production-file identities. Sole authored artifact is this report; no self-referential report checksum is embedded.

## Checks run and graph gate status

Read root/backend AGENTS, Invocation, Build prompt/guidelines/checklist/overlay, T2 snapshot, pertinent Approved Techplan §§9–10 and prior Build report, live composition/db/auth/migrations and auth API anchor. Read-only commands `git rev-parse HEAD`, `git status --short`, workflow `git rev-parse HEAD` / `git status --short`, `git log`, ancestry check, hashing, `cat`, `sed`, `rg`, and `nl` produced the evidence above. `git diff --exit-code -- backend` exits 0. Workflow status output is empty. API identities match the T1 prerequisite, which establishes contract availability only.

No temporary candidate was created. No `go version`, `go get`, `go list -m all`, `go mod download`, `go mod verify`, `govulncheck`, compile, or focused authored test ran in this Run. Exact candidate/actual resolved graphs, module origins/tagged-source review, transitive delta, new checksums, Go minima, compatibility and fresh scan/traces are **not established**. Existing go.mod is a live manifest, not a newly verified actual graph. Five-ID dispositions remain pending:

| ID | Approved plan's affected/fixed range | This Run disposition |
|---|---|---|
| GO-2026-4771 | pgx <v5.9.0; fixed v5.9.0 | Not re-scanned or closed; planned v5.9.2 not applied |
| GO-2026-4772 | pgx <v5.9.0; fixed v5.9.0 | Not re-scanned or closed; planned v5.9.2 not applied |
| GO-2026-5004 | pgx <v5.9.2; fixed v5.9.2 | Not re-scanned or closed; planned v5.9.2 not applied |
| GO-2026-5970 | x/text <v0.39.0; fixed v0.39.0 | Not re-scanned or closed; planned v0.41.0 not applied |
| GO-2026-6629 | x/text <v0.41.0; fixed v0.41.0 | Not re-scanned or closed; planned v0.41.0 not applied |

Ranges above are bound-plan facts, not fresh primary-source verification. No new dependency finding or residual-risk acceptance is asserted. Query-mode/config exposure and tagged SCRAM/OpaqueString source review were not reached; no database secrets/config values were opened. `db.Open` uses pgxpool.New with DATABASE_URL, so no configured protocol/authentication profile can be inferred solely from that source.

## Verification scope confirmation

No heavyweight Testing sweep ran: no race/concurrency, performance/load, final security-class, broad Testing-owned suite, or supply-chain scan was executed. Source reasoning about a possible lock wait is not a concurrency test. Protocol-negative/person-upsert/session-lock tests were neither authored nor run because the new risk terminated the gate. Real Google, PostgreSQL/SCRAM, migrations, restricted roles, browser/HTTPS and independent Testing evidence remain unverified.

## Contract check

- [x] Bound identities and live anchors inspected; unrelated edits preserved.
- [x] Stop-on-new-risk boundary respected before backend/dependency write.
- [ ] Isolated graph-first gate passed.
- [ ] Actual backend graph verified against D8.
- [ ] G1 implementation and required focused tests completed.

T2/G1 remains incomplete. No blocker closure, T3 prerequisite satisfaction, or delivery claim follows from this Run.

## Phase handoff

- **Outcome:** BLOCKED — stopped at graph-first gate step 1 on source-observed session-expiry revalidation risk; no probe or G1 continuation.
- **Result refs:** This Run's `evidence/build-report.md`, Finding F-T2-003 and input/live identities above.
- **Findings:** F-T2-003 — caller timestamp predates row-lock wait and can admit a session expired by acquisition; configured runtime reproduction not established. Previously recorded provenance discrepancy confirmed without acceptance inference.
- **Decision requests:** Owning Techplan/G1 concern disposition of F-T2-003 and bounded continuation scope through Orchestrator/Anhar; no alternate dependency or residual-risk choice made.
- **Blockers:** T2/G1 continuation and dependency acceptance remain blocked under this Invocation's stop-on-risk gate. Existing B-T2-001/B-T2-002 are not closed. Other tasks are outside this Run.
- **Open / unverified:** Isolated and actual D8 graphs/source/integrity/scan/compatibility, query/SCRAM exposure, all authored focused tests, real provider/DB/browser evidence and G1 completion.
- **Recommended continuation:** Orchestrator reconcile the smallest G1 revalidation Finding, then establish a fresh authorized Build continuation starting with the complete isolated graph-first gate. Do not dispatch dependent tasks from this report.
- **Context refs:** This Invocation; Approved Techplan §§9–10/12; current T2 snapshot; live `backend/internal/platform/auth/store.go:118–138`; bound T2-002 report/graph/scan/integrity.
