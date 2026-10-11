# Draft Run Invocation — RUN-C1-ENG-TEST-T2-001

## Identity and route

- **Work Unit / proposed Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-TEST-T2-001`.
- **Phase / Role:** Independent Testing — T2 identity/session and dependency runtime evidence / Tester.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-TESTER-001` / `SESSION-C1-ENG-T2-TESTER-001` (fresh and independent from Implementer/Reviewer).
- **Dispatch:** Human-Assisted; not dispatched.
- **Proposed model / effort:** `gpt-6-sol` / `medium`; approval pending for this exact Run.
- **Target HEAD / branch:** `059d145d5df64fcda427e5a625486a8e95ce1fd8` / `pilot/3-c1-engineering`.
- **Workflow HEAD:** `91a97af8ba64f98b0a692a83fbff866a210c45c0`.
- **Harscode root:** `../harscode-workspace`.

## Route basis and prior evidence

Fresh Code Review `RUN-C1-ENG-CODEREVIEW-T2-002` completed all four passes with Approve and zero findings on the exact T2 six-file scope. Report SHA-256 `3487d5c180d9d60d4c3ad580d20db13e1646b9042361b67fadc7f470fe41a517`; the Review did not execute runtime checks. Earlier `RUN-C1-ENG-CODEREVIEW-T2-001` is terminal STALLED and did not review code.

The current Approved Techplan §12 assigns real Google login/stable person mapping, real PostgreSQL first-login upsert/session contention, browser cookie/session behavior, and real PostgreSQL/SCRAM/restricted-credential behavior to Testing. The T2 task says Build-focused protocol/person-upsert/session-lock checks do not substitute for configured real login, real PostgreSQL concurrency, or independent Testing. Build retry-1 report SHA-256 `4a63608a42dd89c9978c62161b0f6924700115fd0ed06eeed9af42aa04d6b326`; its evidence is a claim to spot-check, not final Testing evidence.

## Bounded T2 Testing scope

Test only behavior exposed by the current T2 identity/session backend and the dependency/runtime compatibility risks owned by Testing. Primary targets:

1. **R1/R9:** actual configured Google authorization/login maps the same stable `(issuer, subject)` person across sessions; verify callback/code-replay and provider/key/config failure behavior where real configuration permits. Protocol-negative cases from Build may be spot-checked rather than recreated. Do not persist or print credentials, tokens, authorization codes, or token payloads.
2. **R10:** observable session cookie flags and dev-only localhost exception; production HTTP rejection / Secure-cookie behavior where HTTPS is configured; Origin/CSRF/logout/session expiry, rotation and revocation behavior at the real HTTP boundary. Use a fresh browser context where the actual surface exists. Do not claim a full user journey that depends on T3/T5 UI.
3. **R14:** actual configured PostgreSQL pool connection and SCRAM authentication using authorized restricted runtime credentials; session query and under-lock behavior; real concurrent first-login upsert; session logout/revocation/confirmation contention, including expiry while waiting for the row lock. Use disposable isolated state, distinct runtime/operator/migration identities where applicable, bounded timeouts, deterministic cleanup, and inspect for leaked resources after abort. Never use production data or bootstrap-superuser-only evidence.
4. **Target verification:** backend `make verify` and `go mod verify` as the Techplan final verification route; run only against the bound tree and record each command/result. `make verify` includes lint, unit, race, contract, secret scanning and dependency scanning. If a constituent command is unavailable or network/configuration constrained, capture the exact limitation and do not claim the gate passed.

Check actual configuration/tool availability without printing secret values. Missing Google test-user/client/callback/consent, HTTPS, Podman/PostgreSQL, or restricted credentials means the applicable real test is `NOT RUN/BLOCKED`, not simulated success. Do not ask the Human to paste secrets into chat. Stop before any production connection, non-disposable DB mutation, broad destructive migration, role escalation, or material test semantic change; route to Orchestrator/Human.

This Run does **not** claim integrated T3 organization/guard behavior, T4 operator runbook rehearsal, T5 frontend/browser workflows, or whole-C1 acceptance. Those remain for their owning implementation and later integrated Testing. R2–R8/R11–R13 scenarios tied to T3/T5 are explicitly deferred. No Testing result may close B-T2-002/B-T2-003 or authorize dependency commit/T3 dispatch without Orchestrator reconciliation.

## Test Focus Pointer execution

Use current Techplan §12 Test Focus Pointer. Open only the applicable anchor in `runs/RUN-C1-ENG-EXPLORATION-001/evidence/stage-3-solutioning.md`: `#material-questions-and-constraints` items 1 and 3, with the `### Attribution and authentication timing` subheading, to recover the bounded person-attribution/authentication rationale. The Organization + Owner consistency anchor is T3-context only; inspect that anchor only if an observed T2 behavior crosses that boundary. Known conflict is outside T2 and is not exercised here. Specialized test design uses current Techplan §§9, 10, 12 and existing retry-1 evidence, not inferred historical tests.

## Required workflow and report

Follow `../harscode-workspace/workflow/5-testing-prompt.md`, `workflow/5-testing/guidelines.md`, `workflow/5-testing/checklist.md`, and `workflow/orchestrated-run-overlay.md`. Fresh-read the entire current Approved Techplan end-to-end before testing. Read the exact latest retry-1 Build report and independent Review report. Route matching practices only through `best-practices/AGENTS.md` and the current clue map; likely Go integration/concurrency, JWT/token lifecycle, authorization, dependency/supply-chain, and PostgreSQL locking guidance.

Inspect actual target-repo test authority `backend/Makefile` and `backend/README.md`; do not infer commands. Testing report must use Harscode sections 0, 0a, 1–4, New Recurring Bug Patterns, Verdict and one structured Phase handoff. Separate confirmed, assumed, simulated, not-run and deferred evidence. Where a rule cannot be exercised at the current T2 observable boundary, state the gap and owner rather than mark it passed. Do not modify production code; any required code changes go to Build via a patch plan.

## Stable input identities

- Approved Techplan `techplan/techplan.md`: `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79`.
- T2 task `techplan/tasks/T2-backend-identity-session.md`: `dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37`.
- Accepted manifest `techplan/tasks/manifest.md`: `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.
- Retry-1 Build report: `4a63608a42dd89c9978c62161b0f6924700115fd0ed06eeed9af42aa04d6b326`.
- Code Review T2-002 report: `3487d5c180d9d60d4c3ad580d20db13e1646b9042361b67fadc7f470fe41a517`.
- Backend `Makefile`: `9ef7028c55e544e829684b4588dee3b89a9651d68d8f6395d5c3b3bb6803cbe2`.
- Backend `README.md`: `430592e4aaed7d4717c2cd95f60c1e5bc0487d6d7721259406ee5a59b8e2ee3b`.
- Root `AGENTS.md`: `d6bb74b1ffa8cbd1a115006191a40574b5f16508d22477d38e09744681367462`.
- `backend/AGENTS.md`: `817d01ab80598637644c64be9bfd7ec18a9dc5e3f0d22485d505f0760f72769f`.
- Current code hashes: go.mod `90b7c1c6e870e32828fa5932a6c4e9a28a65d0671a74bdb807ec29dbc5d27cae`; go.sum `6923543aab63e09da2b894699afb70b2ff1072a2eb014e063c8542ffed7deaa0`; store.go `eb70175ebe756698a3a83b627e0fd6606928d8ac4c355f9dd8a05dab705552b3`; store_test.go `bac0d93f0aa6af1591b10dc9a1df81fba0bf77ea2bdb5a165c129e2611795d60`; oidc_test.go `2ffd30f781b3aeaa09f6beb0fddd00bea90abcab99b221e126867acfc8331ab4`; store_integration_test.go `1b38058a177e0df5f87cf6c732478369e8f798265f2ad91c0762a1ea1879b1e0`.
- Test Focus Exploration anchor `stage-3-solutioning.md`: `e995d13b72438cd9a68a66ef34927ccb9d9000020a0b406de9024cc38383b34b`.

## Coordination hashes to bind after model approval

Approval recording updates lifecycle-owned WU/Events/Control Tower state. Therefore do not copy the current draft-time hashes into final Invocation. After Anhar approves this exact Run/model/effort, record approval in those lifecycle owners, set the Run READY_FOR_HUMAN_DISPATCH, and then capture the resulting three hashes in final `invocation.md`. Do not mutate them again while the Testing Run consumes its Invocation; if needed, stop and re-ground.

- WU current hash after approval: **capture in final Invocation**.
- Events current hash after approval: **capture in final Invocation**.
- Control Tower current hash after approval: **capture in final Invocation**.

## Proposed execution envelope

- **ARTIFACT_TARGET:** None; write `RUN_PATH/evidence/testing-report-1.md`, plus `patch-plan-1.md` only if required.
- **PREAUTHORIZED:** read bound authority/source and current T2 interface; inspect safe configuration presence; execute approved-scope independent tests on disposable/authorized environments; write Testing evidence only.
- **ORCHESTRATOR_DECISION:** reconcile report, blockers, and next route.
- **HUMAN_REQUIRED:** exact model approval before final dispatch; any access to protected credentials, non-disposable data, privilege escalation, or material scope/semantic choice.
- **Status:** draft only, awaiting exact model approval; no test/runtime action is claimed.
