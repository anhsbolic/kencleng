# Run Invocation — RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-001

## Identity and route

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-001`.
- **Phase / Role:** Code Review / independent Reviewer — backend T2 runtime test harness.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-RUNTIME-HARNESS-REVIEWER-001` / `SESSION-C1-ENG-T2-RUNTIME-HARNESS-REVIEWER-001` (fresh; independent from Build Implementer and Orchestrator contexts).
- **Dispatch:** Human-Assisted; terminal `COMPLETED` after Human dispatch and one terminal Review handoff on 2026-10-13. This Run is closed; do not reuse or reopen it.
- **Selected model / effort:** `gpt-6-sol` / `medium`.
- **MODEL_APPROVAL:** `APPROVED_BY_HUMAN` for this exact Review Run, `gpt-6-sol` / `medium`. On 2026-10-13 Asia/Jakarta, Anhar replied “approve bro” to the request naming this Run and model. This approval does not transfer to another Run.
- **Routing rationale:** Four-pass independent review of subprocess-boundary validation, run-owned disposable resource identity/cleanup, test isolation and evidence claims. The registry's strong cross-cutting reasoning capability fits this security/resource-ownership surface; `medium` is the lowest supported sufficient effort.
- **Target revision / branch:** `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering` (observed Review preparation target; exact three-file content identities are assignment-bound below).
- **Workflow revision:** `651ce57f044bff23fba99082322cee6c066d25a6`.
- **Harscode root:** `../harscode-workspace`.
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-001`.
- **Artifact target:** None; Review evidence only.
- **Working directory:** `/home/anhar-solehudin/kencleng-workspace/kencleng`.
- **Communication language:** Bahasa Indonesia.
- **Scope:** `scope.md` (SHA-256 `e886adc4cd09b8ecadf5cd00937fe1790ff6806ded5f738dacc689ae75b10277`). Assignment-defining inputs: `input-bindings.json` (SHA-256 `0d43733690b9f09ffc7dd6300dfe26fcb91a9836b334b28af53f0a0e23ed313b`, 30 bound identities; live recomputation pending Human dispatch).

## Claim and prior evidence

Review applicability is established in this Run's scope: the Build delta affects how a new disposable PostgreSQL fixture is assigned, validated, and deleted, and enables a bounded six-case suite. Independent reasoning about the owner checks, failure cleanup, secret filtering and truth-class boundaries provides material value beyond Build execution evidence.

Bound Build report: `RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001/evidence/build-report.md`, SHA-256 `c43a030248673bd094aef4352758f4d0007bb14bf255c6202dc1708be194f34`.
Bound exact delta identities: `RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-001/evidence/file-identities.json`, SHA-256 `7675a46900fb1ff713f10af784d288bf03156b10a27cbd6594ded92f23c7850a`.
Build reported six selected cases PASS and successful cleanup for that execution. PostgreSQL/TCP/SCRAM and contention were real disposable-fixture observations; handler tests used `httptest`; provider callback was simulated. These are Builder-reported results, not independent Testing acceptance. The initial sandbox Podman failure occurred before fixture creation and was recovered by the Run's approved scoped escalation.

## Required Review workflow

Read current root/backend instructions, current Approved Techplan and T2 task, current OpenAPI contract only as relevant to the test-harness claim, canonical Code Review prompt/guidelines/checklist, orchestrated overlay/run contract, and the exact live three-file diff plus Build report/file identities. Re-ground on current code; do not use the Build report as a substitute for source inspection. Use `best-practices/AGENTS.md` and the relevant routers to read only matching input-validation/injection, resource ownership/privilege separation, or integration-test guidance.

Run all four Review passes in order: Safety, Quality, Stack-Specific Best Practices, Consistency. Do not edit source. No test, harness, scanner, Make, Podman, container, DB, provider, browser or HTTPS operation is authorized. Do not replay prior Build/Testing work. If read-only inspection cannot resolve a concrete question, identify it for Orchestrator routing.

Write `evidence/review-findings-1.md`; include compact provenance, the four passes, verification executed (`none` unless authority changes), scoped verdict, and exactly one `## Phase handoff`. Write `evidence/patch-plan-1.md` only if a source change is necessary. Preserve every terminal report and original failure unchanged.

## Execution envelope

- **PREAUTHORIZED after exact model approval and Human dispatch:** read-only exact-diff four-pass Review and Run-owned findings/patch-plan evidence.
- **ORCHESTRATOR_DECISION:** reconcile Review and determine fresh independent Testing claim, fixture scope and post-delta final command gates.
- **HUMAN_REQUIRED:** exact model approval and mechanical dispatch; any source change, fixture/test/scanner/DB/provider execution, scope/privilege expansion or new product/runtime semantic choice.

Terminal reconciliation: Human dispatch is evidenced by the terminal Review report. Report SHA-256 `40a6018c5cb208cc36e688cfd444f1d7bbd3f1f4c5963ee6eaba20decfacd941`; patch plan SHA-256 `fe85264fcda7857eab979cfd0434fb584715fdda49c559d9f0002fb5ae961239`. The Review Run is closed; preserve both terminal artifacts unchanged.
