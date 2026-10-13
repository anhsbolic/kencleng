# Run Invocation — RUN-C1-ENG-TEST-T2-AUTH-HARNESS-RUNTIME-001

## Identity and route

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-TEST-T2-AUTH-HARNESS-RUNTIME-001`.
- **Phase / Role / Profile:** Independent Testing / Tester / backend auth harness runtime and cleanup.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-AUTH-HARNESS-RUNTIME-TESTER-001` / `SESSION-C1-ENG-T2-AUTH-HARNESS-RUNTIME-TESTER-001` (fresh; independent from Build, Review, and prior Testing contexts).
- **Dispatch:** Human-Assisted; mechanically dispatched by Human and terminal handoff received on 2026-10-13. Run is closed; terminal status is recorded below.
- **Selected model / effort:** `gpt-6-sol` / `medium`.
- **MODEL_APPROVAL:** APPROVED_BY_HUMAN — on 2026-10-13 Anhar replied “approve bro” to the exact request for this Run / `gpt-6-sol` / `medium`. Approval applies only to this Run; it does not attest runtime model identity.
- **FIXTURE_AUTHORIZATION:** APPROVED_BY_HUMAN — on 2026-10-13 Anhar replied “approve bro” to the exact request for one disposable PostgreSQL fixture owned by this Run, bounded by `scope.md`. Prior Build fixture approvals were not transferred.
- **Routing rationale:** This Run exercises authentication/session persistence, restricted PostgreSQL roles, SCRAM/TCP, database contention, resource deletion ownership, scanner dispositions, and current backend final gates. The registry's security/cross-cutting-capable Sol model at its lowest supported sufficient effort (`medium`) is selected; approval is required. Approval does not attest runtime model identity.
- **Target revision / branch:** `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering`; exact assignment content is captured by the manifest.
- **Workflow revision:** `651ce57f044bff23fba99082322cee6c066d25a6`.
- **Harscode root:** `../harscode-workspace`.
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TEST-T2-AUTH-HARNESS-RUNTIME-001`.
- **Artifact target:** None; Testing evidence belongs to this Run.
- **Working directory:** `/home/anhar-solehudin/kencleng-workspace/kencleng`; declared commands execute from `backend/`.
- **Communication language:** Bahasa Indonesia.
- **Scope:** `scope.md` (SHA-256 `b316d1d3e5faba3b1016045457e0a6b47c9f8c026411e995282b838a4610008c`). Assignment-defining inputs: `input-bindings.json` (SHA-256 `8f4857a145ec565f60083b6e8c929e33a80581d6962f5c009c34e89ec45f2c33`, 48 exact content identities). Recompute before dispatch. Mutable WU/Events/Control Tower lifecycle projections are contextual, not self-invalidating bindings.

## Claim and phase applicability

The claim is limited to independent execution of the current cleanup harness's six-case `t2` suite on this Run's real disposable PostgreSQL fixture, successful-path cleanup of its verified full-ID container and anonymous volume, focused mock cleanup tests, and the current backend `make verify` chain. This Testing adds independent value because the final ID guard currently has Builder-reported mocks and Review's static reasoning only; the prior real fixture belonged to a different Run and used earlier cleanup code.

It does not establish full T2 acceptance, close `B-T2-002` / `B-T2-003`, satisfy T2→T3 dependency conditions, establish configured application-runtime or production HTTPS behavior, or satisfy either deferred Human hold. Builder evidence, real PostgreSQL, `httptest`, simulated Provider, and independent Testing results must remain separately labeled.

## Execution contract

Read the full current Approved Techplan and T2 task; current backend README/Makefile and auth verification environment; terminal full-ID Build and Review reports; original runtime-harness Build report; latest backend-gate report; exact Test Focus anchor in the Techplan; canonical Testing prompt, guidelines/checklist, overlay, run contract, local model policy, scanner/dependency/network runbooks. Check material readiness for Go 1.26.9, configured Podman and local `postgres:16` image, and available scanner route without loading secrets. If known prerequisites make this bounded outcome unreachable, stop before fixture creation and report/narrow the assignment.

After both explicit approvals and Human dispatch, execute exactly the three commands in `scope.md` with `PYTHONDONTWRITEBYTECODE=1` and Run-scoped writable build/scanner caches where needed. Keep checksum verification enabled. A tooling recovery is permitted only when declared by applicable runbooks and bounded by this active Invocation; preserve failed attempts and stop if recovery does not satisfy the envelope. Preserve dirty workspace and do not edit production source, test code, dependencies, scanner policy, Product/Techplan authority, shared DB, `.env`, or historical evidence.

The full-ID real cleanup claim is successful-path only: verify exact Run label/owner/name/full ID before removal, and verify the owned container and anonymous data volume are absent afterward. Mock negative paths remain separate evidence. If runtime cleanup cannot be positively established, fail closed and do not broaden deletion.

## Execution envelope

- **PREAUTHORIZED only after exact Human approvals and dispatch:** read current bound inputs; safe environment readiness checks; exact three commands; one Run-owned disposable fixture exactly bounded by `scope.md`; bounded caches/captures; this Run's report.
- **ORCHESTRATOR_DECISION:** reconcile actual binding/delta/results and route subsequent obligations.
- **HUMAN_REQUIRED:** exact model approval, fixture authorization, mechanical dispatch, any resource/privilege expansion, shared or persistent DB, secret/provider access, source/dependency/policy change, or material scope/authority change.

Model approval and the bounded fixture authorization are recorded above. Human mechanical dispatch and terminal handoff are reconciled below.

## Terminal reconciliation

- **Terminal report:** `evidence/testing-report-1.md`; SHA-256 `00fe38a3f14c7041d77aba6825a752eb4e8fe1de379a177e09af785893d8819e`; exactly one `## Phase handoff`.
- **Terminal result:** PASS, limited to this Run's Independent Testing claim. All 48 assignment-defining bindings matched before and after execution; scope and manifest identities remain the exact values above. No bound source, test, dependency, policy, Product/Techplan, `.env`, or shared-database input changed.
- **Executed evidence:** five focused cleanup mock tests and all six selected `t2` cases passed. A Run-owned disposable PostgreSQL 16 fixture exercised real PostgreSQL/TCP/SCRAM and contention; successful-path owned container and anonymous-volume cleanup was verified by the runner. HTTP handlers used `httptest`; Provider behavior was simulated. `make verify` exited 0. `go mod verify` evidence was reused only for byte-identical module files.
- **Finding dispositions:** Gosec reported 3 raw findings and Gitleaks history reported 4; all had accepted dispositions and effective gates passed. Gitleaks worktree reported 0 findings; govulncheck reported no vulnerabilities. These results do not mean all scanners had zero raw findings.
- **Evidence limits:** no configured application-runtime credential/role or deployed HTTP listener, real Google/provider exchange, browser, production HTTPS, or abnormal real cleanup path was exercised. Runner cleanup evidence covers its observed successful path; failure branches remain mocked. This is independent Testing evidence, not full T2 acceptance or T2/T3 readiness.
- **Route:** no further Review is applicable to this exact Run because it introduced no source delta; repeating this bounded Testing claim adds no material independent value. Reconcile the owning conditions for `B-T2-002` and `B-T2-003` before deciding whether another narrowly scoped Run is needed. Keep both blockers and Human holds open; do not infer acceptance/readiness or dispatch T3.
