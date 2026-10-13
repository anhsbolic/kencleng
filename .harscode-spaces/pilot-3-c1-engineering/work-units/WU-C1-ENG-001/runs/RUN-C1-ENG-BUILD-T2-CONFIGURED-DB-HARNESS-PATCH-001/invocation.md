# Run Invocation — RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001

## Identity and posture

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001`.
- **Phase / route:** Build/Patch — add the production configured-pool oracle required before independent Testing can establish R14 on a disposable fixture.
- **Role / specialization:** Implementer / backend auth harness.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-CONFIGURED-DB-HARNESS-PATCH-IMPLEMENTER-001` / `SESSION-C1-ENG-T2-CONFIGURED-DB-HARNESS-PATCH-IMPLEMENTER-001` (fresh assignment identities; not launched).
- **Dispatch:** Human-Assisted; `READY_FOR_HUMAN_DISPATCH`. On 2026-10-13, Anhar replied “approve bro” to the exact Run-specific disposable PostgreSQL fixture request. This authorizes only the single fixture and boundary in `scope.md`. Human mechanical dispatch remains pending; no Run execution has started.
- **Selected model / effort:** `gpt-6-luna` / `medium`; supported and `approval_required: false` in the current Human-owned runtime registry. This is a narrow test-infrastructure adaptation; no auth semantics or dependency decisions are assigned.
- **Model routing rationale:** Luna's declared coding, reasoning, and repository-work capabilities at medium are sufficient for this bounded adaptation to the existing harness. Escalate by reporting if live evidence reveals a security-semantic or architecture change beyond the allowed test files.
- **Fixture authorization:** `APPROVED_BY_HUMAN` for this exact Run on 2026-10-13. Approval covers one locally imaged PostgreSQL 16 disposable container, exact existing restricted role/grants, generated private credentials, and ownership-checked cleanup exactly as `scope.md` states. No earlier Run's fixture approval transfers and no broader DB/credential/resource authority is granted.
- **Target revision / branch:** `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering`. Preserve the current dirty workspace; exact live source content is in `input-bindings.json`.
- **Workflow revision:** `651ce57f044bff23fba99082322cee6c066d25a6`.
- **Harscode root:** `../harscode-workspace`.
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001`.
- **Artifact target:** None; patch is implementation output, terminal report is Run-owned evidence.
- **Working directory:** `/home/anhar-solehudin/kencleng-workspace/kencleng`; command runs from `backend/`.
- **Communication language:** Bahasa Indonesia.
- **Scope:** `scope.md`; SHA-256 recorded in `input-bindings.json`. Verify all assignment-defining bindings before edits/fixture setup. WU/Events/Control Tower lifecycle projections and this routing's status are mutable coordination context, not self-invalidating exact bindings.

## Claim and trigger

The latest independent harness Testing passed its bounded claim but explicitly did not establish Techplan §12 R14's configured application pool/session-query condition or the specific B-T2-003 post-lock expiry condition. Current `auth_harness.py` constructs the restricted runtime pool from `C1_HARNESS_RUNTIME_DSN` with a direct pgxpool constructor; production startup obtains its pool through `DATABASE_URL` and `db.Open`. The six existing `t2` tests include real session queries and expiry lock contention, so the smallest prerequisite is to route those tests through the same production pool constructor and a child-only fixture `DATABASE_URL`.

This Run provides Build evidence that the existing suite executes through that configured-pool path. A fresh independent Testing Run must verify it with its own fixture. B-T2-002 also retains the separate owning acceptance/commit decision after evidence reconciliation; B-T2-003 remains separately reconciled. No claim is made that either blocker closes from this Run.

## Required guidance and execution envelope

Read the current Approved Techplan and T2 task, root/backend `AGENTS.md`, `docs/project/c1-auth-verification-environment.md`, canonical Harscode Build/Patch prompt/guidelines/checklist, orchestration overlay/run contract/current runtime registry, and the bound harness Build/Testing reports. Re-ground on live harness, `db.Open`, Store/session lock test, fixture runner and README. Do not load unrelated history or change authority to fit implementation.

- **PREAUTHORIZED only after exact fixture authorization and Human dispatch:** edit the three files in `scope.md`; execute the single Run-owned `t2` command on this Run's disposable fixture; capture safe bounded evidence; write this Run's patch report.
- **ORCHESTRATOR_DECISION:** reconcile exact delta/results, decide whether targeted Review is materially useful, and prepare a fresh independent Testing Run only after terminal handoff.
- **HUMAN_REQUIRED:** this Run's fixture authorization and mechanical dispatch; any production auth/db/server change, role/grant/resource expansion, shared/persistent DB or `.env` access, real provider login, dependency/policy/Product/Techplan change, or acceptance/residual-risk/commit decision.

The current D8 dependency graph/provenance/scan and exact-delta Review evidence is already present and unchanged; do not rerun it here. `make verify` and scanner results from prior terminal Runs remain bounded to their captured targets and are not refreshed by this Build. Do not perform independent Testing in this occurrence.

## Human-Assisted kickoff

```text
Jalankan RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-PATCH-001 sesuai Invocation dan scope yang terikat. Setelah fixture disposable untuk Run ini diotorisasi, ubah hanya tiga file yang diizinkan agar test harness memakai DATABASE_URL child milik fixture dan pool production db.Open; pertahankan enam assertion t2 termasuk expiry saat menunggu lock. Jalankan hanya command t2 yang tercantum. Jangan baca .env atau memakai DB bersama, provider/Google, scanner, gate, atau dependency. Laporkan exact delta, bukti PostgreSQL versus simulasi, cleanup, hasil, dan satu Phase handoff.
```

## Terminal closure

Invocation is prepared and ready for Human mechanical dispatch. No Participant Session has been launched; no fixture, test, scanner, gate, runtime, source edit, dependency operation, or commit has occurred for this Run. After execution begins and a terminal handoff is produced, preserve this occurrence as terminal; any further Build/Review/Testing is a new Run with fresh Participant and Session identities.
