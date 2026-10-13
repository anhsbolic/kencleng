# Run Invocation — RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-001

## Identity and posture

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-001`.
- **Phase / route:** Independent Testing — configured-pool R14 and post-lock-expiry B-T2-003 evidence, assessed separately.
- **Role / specialization:** Tester / backend auth harness runtime.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-CONFIGURED-DB-HARNESS-TESTER-001` / `SESSION-C1-ENG-T2-CONFIGURED-DB-HARNESS-TESTER-001` (fresh; not launched).
- **Dispatch:** Human-Assisted; prepared behind terminal Review reconciliation, exact Run-specific fixture authorization, and Human mechanical dispatch. Not dispatchable yet; no Testing Run is active.
- **Selected model / effort:** `gpt-6-luna` / `medium`; current Human-owned registry provides repository-work, coding, and reasoning capabilities and marks it `approval_required: false`.
- **Model routing rationale:** This bounded claim uses an existing harness and one predeclared fixture command; Luna at medium is sufficient for independent execution/evidence classification. No model approval is required by the registry.
- **Fixture authorization:** PENDING for this exact Testing Run. Request only the one new fixture and exact boundary in `scope.md`; previous Run authorizations do not transfer.
- **Target revision / branch:** Build baseline `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering`; exact current source hashes are bound. Reconfirm after terminal Review and before dispatch.
- **Workflow revision:** `651ce57f044bff23fba99082322cee6c066d25a6`.
- **Harscode root:** `../harscode-workspace`.
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-001`.
- **Artifact target:** None; Testing evidence belongs to this Run.
- **Working directory:** `/home/anhar-solehudin/kencleng-workspace/kencleng`; command executes from `backend/`.
- **Communication language:** Bahasa Indonesia.
- **Scope:** `scope.md`; assignment-defining sources and exact target identities are in `input-bindings.json`. Add the terminal Review report identity before dispatch after applicability reconciliation. WU/Events/Control Tower lifecycle projections are mutable context, not assignment bindings.

## Claim and applicability

Independent Testing is applicable because Techplan §12 R14 assigns real isolated PostgreSQL configured-pool/SCRAM/session-query evidence to Testing, while the latest independent Testing predates the `DATABASE_URL` → `db.Open` harness route. B-T2-003's real post-lock-expiry oracle is a distinct open condition. This Run uses a new independent Tester/Session and a fresh Run-owned fixture; it must return separate evidence/results for R14 and B-T2-003 even if one execution supports both.

## Required guidance and execution envelope

Read the current canonical Testing prompt/guidelines/checklist and overlay; full current Approved Techplan and T2 task; exact Test Focus Pointer anchors that remain relevant; terminal Build report; terminal targeted Review report; current backend README/Makefile and auth verification environment; applicable scanner/dependency/network runbooks only if the declared assignment expands under owning authority. No raw Exploration corpus. Before dispatch, check only material prerequisites for this bounded fixture claim without reading secrets or accessing `.env`/shared DB.

- **PREAUTHORIZED only after Review reconciliation, exact fixture authorization, and Human dispatch:** one fresh Run-owned fixture; the single scoped `t2` command; bounded evidence/reporting.
- **ORCHESTRATOR_DECISION:** reconcile the separate R14 and B-T2-003 evidence and any findings against their owning conditions.
- **HUMAN_REQUIRED:** exact fixture authorization for this Run, mechanical dispatch, any role/grant/resource expansion, shared/persistent DB, `.env` or provider secrets, real Google login, source/dependency/policy/authority changes, owner acceptance/commit, or residual-risk acceptance.

## Human-Assisted kickoff

```text
Jalankan RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-001 hanya setelah Review predecessor terminal dan direkonsiliasi, fixture disposable PostgreSQL baru untuk Run ini disetujui tepat sesuai scope, dan dispatch Human dilakukan. Jalankan hanya perintah `python3 scripts/auth_harness.py --run-id RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-001 --suite t2` dari backend/. Laporkan R14 configured-pool/session queries dan B-T2-003 expiry-under-lock sebagai dua klaim terpisah. Bedakan PostgreSQL/TCP/SCRAM nyata, HTTP httptest, Provider simulasi, dan cleanup runner. Jangan jalankan scanner/gate/make verify, akses DB bersama/.env, login Google, atau perluas fixture/privilege. Tulis testing-report-1.md dengan tepat satu Phase handoff.
```

## Terminal closure

This Run is prepared but not launched. Do not dispatch while its Review predecessor is unresolved or exact target hashes drift. Fixture authorization and mechanical dispatch are Run-specific. Terminal evidence, if later produced, will support only its declared claims; B-T2-002 still requires owning acceptance/commit, and no T2/T3 acceptance/readiness or T3 dispatch follows automatically.
