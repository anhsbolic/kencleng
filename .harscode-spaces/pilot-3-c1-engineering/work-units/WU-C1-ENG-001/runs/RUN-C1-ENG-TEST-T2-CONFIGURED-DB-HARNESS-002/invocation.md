# Run Invocation — RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-002

## Identity and posture

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-002`.
- **Phase / route:** Independent Testing — configured-pool R14 and post-lock-expiry B-T2-003 evidence, assessed separately.
- **Role / specialization:** Tester / backend auth harness runtime.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-CONFIGURED-DB-HARNESS-TESTER-002` / `SESSION-C1-ENG-T2-CONFIGURED-DB-HARNESS-TESTER-002` (fresh; not launched).
- **Dispatch:** Human-Assisted; prepared after terminal Review reconciliation. The Human instruction to prepare independent Testing with a fixture owned by the new Run authorizes the exact bounded fixture in `scope.md`. Human mechanical dispatch remains pending; no Testing Run is active.
- **Selected model / effort:** `gpt-6-luna` / `medium`; current Human-owned registry marks it `approval_required: false` and supports repository work, coding, and reasoning.
- **Model routing rationale:** A bounded existing harness, one focused mocked suite, and one fixture command fit Luna at medium. Escalate only if evidence materially expands beyond this scope.
- **Fixture authorization:** AUTHORIZED by the Human instruction to prepare independent Testing with a new Run-owned fixture; limited to the single disposable PostgreSQL 16 fixture and boundaries in `scope.md`. No prior Run approval is reused.
- **Target revision / branch:** `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering`; exact material inputs and current target hashes are in `input-bindings.json`.
- **Workflow revision:** `651ce57f044bff23fba99082322cee6c066d25a6`.
- **Harscode root:** `../harscode-workspace`.
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-002`.
- **Working directory:** `/home/anhar-solehudin/kencleng-workspace/kencleng`; commands execute from `backend/`.
- **Communication language:** Bahasa Indonesia.
- **Scope:** `scope.md`; assignment-defining inputs and exact target identities are in `input-bindings.json`. Mutable WU/Events/Control Tower lifecycle fields are context, not assignment bindings.

## Claim and applicability

Independent Testing remains applicable: Techplan §12 assigns configured PostgreSQL pool/SCRAM/session-query evidence to Testing, and the R14 obligation is not met by the previous direct-pool harness. B-T2-003's actual expiry-under-lock oracle is a separate open condition. The six selected `t2` cases cover the declared runtime path; an additional mocked unittest checks the F1 regression. Terminal Review `RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001` approved F1 for the exact two-file delta by static inspection; this verdict does not itself establish runtime behavior.

## Required guidance and execution envelope

Read the current canonical Testing prompt/guidelines/checklist and orchestrated overlay; Approved Techplan and T2 task; exact relevant Test Focus evidence anchor; configured-pool and credential-patch Build reports; both relevant Review reports and F1 patch plan; current backend README/Makefile and auth verification environment; root/backend `AGENTS.md`. Do not read `.env` or access a shared DB.

- **PREAUTHORIZED only after exact fixture authorization and Human dispatch:** one fresh Run-owned disposable fixture; the two commands in `scope.md`; bounded evidence/reporting.
- **ORCHESTRATOR_DECISION:** reconcile separate R14 and B-T2-003 evidence and any findings against owning conditions.
- **HUMAN_REQUIRED:** exact fixture authorization, mechanical dispatch, privilege/resource expansion, shared/persistent DB, `.env` or provider secrets, real Google login, source/dependency/policy/authority change, owner acceptance/commit, or residual-risk acceptance.

## Human-Assisted kickoff

```text
Jalankan RUN-C1-ENG-TEST-T2-CONFIGURED-DB-HARNESS-002 sesuai Scope setelah Human mechanical dispatch. Gunakan hanya fixture disposable PostgreSQL baru milik Run ini dengan batas yang sudah diotorisasi di Scope. Dari backend/, jalankan hanya dua perintah di Scope: focused unittest dan runner `t2` dengan Run ID ini. Laporkan R14 configured-pool/session queries dan B-T2-003 expiry-under-lock sebagai dua klaim terpisah. Klasifikasikan PostgreSQL/TCP/SCRAM nyata, HTTP `httptest`, Provider simulasi, dan cleanup runner. Jangan akses shared DB/`.env`, login Google, atau perluas fixture/privilege; jangan jalankan gate/scanner/`make verify`. Tulis `testing-report-1.md` dengan tepat satu Phase handoff.
```

## Terminal closure

This Run is prepared but not launched. Fixture authorization and mechanical dispatch are Run-specific. Any terminal result supports only its declared claims; B-T2-002 still requires R14 reconciliation plus owning acceptance/commit, and no T2/T3 readiness/acceptance or T3 dispatch follows automatically.
