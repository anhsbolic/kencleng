# Run Invocation — RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001

## Identity and posture

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001`.
- **Phase / route:** Build/Patch re-entry from independent Review finding F1.
- **Role / specialization:** Implementer / backend auth harness.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-IMPLEMENTER-001` / `SESSION-C1-ENG-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-IMPLEMENTER-001` (fresh identities; not launched).
- **Dispatch:** Human-Assisted; `READY_FOR_HUMAN_DISPATCH`. Human mechanical dispatch is pending; no Run is active.
- **Selected model / effort:** `gpt-6-luna` / `medium`; current Human-owned runtime registry supports this bounded repository patch and marks it `approval_required: false`.
- **Model routing rationale:** A small Python subprocess-environment boundary patch plus mocked regression is within Luna's coding/repository-work/reasoning capabilities. No security-semantics or production auth behavior is assigned.
- **Target revision / branch:** `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering`; exact current content is bound in `input-bindings.json`.
- **Workflow revision:** `651ce57f044bff23fba99082322cee6c066d25a6`.
- **Harscode root:** `../harscode-workspace`.
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001`.
- **Artifact target:** None; patch report belongs to this Run.
- **Working directory:** `/home/anhar-solehudin/kencleng-workspace/kencleng`; focused command executes from `backend/`.
- **Communication language:** Bahasa Indonesia.
- **Scope:** `scope.md`; exact Review finding/plan and current source identities are bound in `input-bindings.json`. Mutable WU/Events/Control Tower lifecycle projections are not assignment bindings.

## Claim

Review F1 found the fixture-bearing environment is passed to `go version`, which does not need the fixture credentials. Patch only the metadata subprocess environment and add a focused mocked regression; preserve the fixture environment for the actual Go test child. No fixture or real runtime evidence is needed for this patch.

## Execution envelope

- **PREAUTHORIZED after Human mechanical dispatch:** edit the two scoped test-infrastructure files and run the single focused mocked Python unit suite specified in `scope.md`; write this Run's patch report.
- **ORCHESTRATOR_DECISION:** reconcile exact delta/result and decide targeted Review applicability.
- **HUMAN_REQUIRED:** mechanical dispatch; any fixture, DB/runtime, Provider, shared/persistent DB, `.env`, privilege/resource expansion, production/auth behavior, dependency/policy/authority change, or acceptance/commit decision.

Re-ground on the exact Review report and patch plan, current live runner/test code, root/backend instructions, and current Build guidance before editing. Do not run any other test, gate, scanner, fixture, or runtime operation. Preserve the prior Build and Review terminal evidence unchanged.

## Human-Assisted kickoff

```text
Jalankan RUN-C1-ENG-BUILD-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001 sesuai Invocation dan scope yang terikat. Perbaiki hanya F1: cegah fixture DSN/credential masuk ke environment proses `go version`, sambil mempertahankan environment fixture untuk Go test child. Tambahkan regression mock dalam file test yang diizinkan dan jalankan hanya `python3 -m unittest discover -s scripts -p test_auth_harness.py` dari backend/. Tidak perlu fixture. Jangan ubah file lain atau menjalankan gate, scanner, DB/runtime, Provider, atau Google login. Tulis patch-report-1.md dengan tepat satu Phase handoff.
```

## Terminal closure

This occurrence is prepared and not launched. Human mechanical dispatch remains pending. It can only establish this narrow F1 patch/mocked-regression claim. Further Review/Testing uses fresh applicable phase routing; the previously staged Testing Run is not dispatchable against a changed target.
