# Run Invocation — RUN-C1-ENG-TEST-T2-BACKEND-GATE-002

## Identity and route

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-TEST-T2-BACKEND-GATE-002`.
- **Phase / Role:** Independent Testing / Tester — current-tree backend command gate.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-BACKEND-GATE-TESTER-002` / `SESSION-C1-ENG-T2-BACKEND-GATE-TESTER-002` (fresh, independent from prior Build, Review, and Testing contexts).
- **Dispatch:** Human-Assisted; `READY_FOR_HUMAN_DISPATCH`. Prepared, not dispatched.
- **Selected model / effort:** `gpt-6-sol` / `medium`.
- **MODEL_APPROVAL:** APPROVED by Anhar on 2026-10-12 (“approve bro”) for exact `RUN-C1-ENG-TEST-T2-BACKEND-GATE-002` / `gpt-6-sol` / `medium`. Prior Run approvals were not transferred. Approval does not attest runtime identity.
- **Routing rationale:** This bounded final gate spans Go unit/race/contract-tag commands, Gosec/Gitleaks source/disposition interpretation, Go vulnerability scanning and module integrity. `gpt-6-sol` provides the repository/security evidence reasoning used by the previously approved backend-gate route; `medium` is its lowest supported effort and is sufficient for the same bounded claim. Approval does not attest runtime identity.
- **Target revision / branch:** `dae55ac44beb0e967b94c2f791d8a0b8d8e72372` / `pilot/3-c1-engineering`, plus assignment-bound current working-tree inputs listed in `input-bindings.json`.
- **Workflow revision:** `651ce57f044bff23fba99082322cee6c066d25a6`.
- **Harscode root:** `../harscode-workspace`.
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TEST-T2-BACKEND-GATE-002`.
- **Artifact target:** None; this Run writes only Run-owned Testing evidence.
- **Working directory:** `/home/anhar-solehudin/kencleng-workspace/kencleng`; commands run from `backend/`.
- **Communication language:** Bahasa Indonesia.
- **Scope:** `scope.md` (SHA-256 `cb80c79ee76e0f0fef45138ba2e1dd18f7679dc9b15b65780ba4ea98253d76ff`). Assignment-defining inputs: `input-bindings.json` (SHA-256 `dd75870fca2944beae5abb2a1f73f9cf93bcb3626115f6441df59b56cf149eb3`, 43/43 live bindings matched during preparation).

## Claim and phase applicability

Establish only whether the bound current backend tree completes the repository-authoritative `make verify` chain and `go mod verify`, with actual outcomes and properly interpreted scanner dispositions. The preceding `RUN-C1-ENG-TEST-T2-BACKEND-GATE-001` report is terminal PASS on its own bound tree; this new Run exists because current backend `README.md`, logout `http.go`, and logout harness files have different exact identities. No prior result is inherited as a PASS for these changed inputs.

The logout Build regression and independent Review remain sufficient for the narrowly closed object-validation Finding. This gate Run does not claim to exercise the Run-bound `authharness` fixture or the Review's optional `Set-Cookie` oracle. HTTP/DB runtime, SCRAM/restricted privileges, concurrency, real Google, integrated UI, B-T2-002/B-T2-003, T2/T3 readiness, and full C1 acceptance remain outside this claim and open/unestablished as recorded by their owners/holds.

## Required workflow and commands

Before execution, fresh-read the complete current Approved Techplan and relevant T2 task; read the exact bound logout Build and Review reports, prior backend-gate report, current Testing prompt/guidelines/checklist, orchestrated overlay/run contract, root/backend AGENTS, scanner/dependency/network runbooks, and the actual `backend/Makefile` / `backend/README.md`. Open only the recorded Test Focus anchor `stage-3-solutioning.md#material-questions-and-constraints` items 1 and 3 and `Attribution and authentication timing`; other T3/UI anchors are out of this command claim.

Preflight Go 1.26.9, writable per-Run caches, current actual scanner/tool availability and the configured Podman Compose Gitleaks route. Do not inspect secret values. Execute exactly:

1. `make verify` from `backend/`, preserving the actual ordered target/constituent results and obeying its failure/stop behavior.
2. `go mod verify` from `backend/`, with checksum verification retained.

Use safe read-only environment overrides as required by the project runbooks (`GOTOOLCHAIN=go1.26.9`, `GOFLAGS=-mod=readonly`, scoped writable caches, and the approved empty Compose env-file route). Do not load `.env` or application/provider/DB environment variables. A sandbox/tooling constraint may receive one bounded same-command harness-approved recovery only while this Run is active; preserve the first attempt and stop if recovery fails. Do not run scanner subcommands separately after a substantive `make verify` failure to imply the final chain passed.

Follow `../harscode-workspace/workflow/5-testing-prompt.md`, `workflow/5-testing/guidelines.md`, `workflow/5-testing/checklist.md`, and `workflow/orchestrated-run-overlay.md`. Produce `evidence/testing-report-1.md` with sections 0, 0a, 1–4, Verdict, and exactly one `## Phase handoff`. Save only result-bearing captures required to support material failures/disputes and the gate result. No code, tests, policy, dependency, authority, DB, provider, or browser edits/actions.

## Execution envelope

- **PREAUTHORIZED:** read bound inputs and current source; safe tool readiness inspection; the two declared commands from `backend/`; scoped writable caches/result captures; Run-owned Testing report/evidence.
- **ORCHESTRATOR_DECISION:** reconcile report and determine later claim-scoped route.
- **HUMAN_REQUIRED:** exact model approval, Human mechanical launch, protected credentials/non-disposable DB or privilege boundary, source/policy/dependency changes, or material scope/authority changes.

No Testing action has occurred in this prepared Run. Human mechanical dispatch remains pending; approval does not prove a runtime model identity or execution result.
