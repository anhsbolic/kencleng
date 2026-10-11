# Run Invocation — RUN-C1-ENG-BUILD-T2-PATCH-001

## Identity / dispatch

- WORK_UNIT_ID: WU-C1-ENG-001
- RUN_ID: RUN-C1-ENG-BUILD-T2-PATCH-001
- WORK_UNIT_PATH: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`
- RUN_PATH: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-PATCH-001`
- ROLE / Profile: Implementer / generic Implementer with backend specialization (canonical Build/Patch guidance + Go/REST/secrets routes; no separate project Profile artifact).
- PHASE_ROUTE: fresh Build/Patch requested by failed T2 Testing; not continuation of TEST-T2-001 or BUILD-T2-006.
- PARTICIPANT: PARTICIPANT-C1-ENG-T2-PATCH-IMPLEMENTER-001
- SESSION: SESSION-C1-ENG-T2-PATCH-IMPLEMENTER-001
- SESSION_TRANSITION: FRESH — reconstruct from durable inputs, no prior executor context.
- ARTIFACT_TARGET: bounded backend production patch surface described below; Run-owned patch report.
- SELECTED_MODEL / REASONING_EFFORT: gpt-6-sol / medium.
- MODEL_APPROVAL: PENDING exact Human approval for this Run. Prior approvals do not transfer. Registry approval_required=true. No Participant launched; WAITING_HUMAN.
- MODEL_ROUTING_RATIONALE: two mechanical lint edits plus security candidate applicability/disposition across cookie construction and secret-scan scope requires strong scoped security reasoning; medium lowest selectable effort for Sol. Model strength does not settle missing policy or risk authority.
- TARGET_REVISION / branch: 059d145d5df64fcda427e5a625486a8e95ce1fd8 / pilot/3-c1-engineering, with pre-existing worktree changes.
- WORKFLOW_REVISION: 91a97af8ba64f98b0a692a83fbff866a210c45c0.
- HARSCODE_WORKSPACE_ROOT: ../harscode-workspace
- Working directory: /home/anhar-solehudin/kencleng-workspace/kencleng
- COMMUNICATION_LANGUAGE: Bahasa Indonesia; no additional profile path.

## Inputs and preflight

Exact `input-bindings.json` SHA-256 `b2975ded1c2a8677762afbf154a76a34593e60dc9208baa60515ad2026cb882b` binds current parent/task/manifest, original and supplemental reports/patch plans, reconciliation, live backend sources, Makefile/README/Compose, active contract, lifecycle authorities, registry and routed guidance. Preflight every identity before execution; stale relevant input => stop/report and re-ground via Orchestrator. Target HEAD is provenance plus bound live worktree, not evidence of accepted implementation. Do not edit WU/Events/Control Tower while this Invocation is consumed.

PRIOR_ARTIFACTS: Approved `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`; T2 task and accepted manifest; original TEST-T2-001 patch plan/report; `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-PATCH-001/evidence/supplemental-retest/patch-plan.md` and `testing-report.md`; retry-1 Build and Review T2-002 reports in bindings. Safe supplemental captures and per-candidate Gitleaks metadata are indexed in `evidence/preserved-evidence-hashes.json`. They are prior diagnostic evidence to inspect, not this Run's execution or accepted dispositions.

Read current root/backend AGENTS, canonical workflow/3-build-prompt.md, 3-build/guidelines.md and checklist.md, orchestrated-run-overlay.md, Approved parent and T2 task. Reopen live affected source. Route practices through best-practices/AGENTS.md and Go index, then REST cookie/CSRF and secrets/logging guidance as relevant. Do not load unrelated raw Exploration/history or sibling tasks by default.

## Bounded patch and triage

1. Fix only initial capitalization of two constructor validation errors in backend/internal/platform/auth/oidc.go:38/42. Preserve branches, error meaning, trust and propagation. No ST1005 suppression or lint weakening.
2. Triage three G124 candidates at http.go:49/158/277 against live cookie helper, BrowserConfig validation and set/clear paths. Record exact applicability and bounded focused emitted-cookie/config checks where necessary. Preserve production HTTPS/Secure/HttpOnly/Lax and exact development localhost exception. No broad nosec/allowlist. Do not infer exploitation or clearance from source/count alone. Mechanical semantics-preserving correction may be proposed; material cookie/auth/security change stops before implementation and routes to owning authority. G1 already authorizes its exact approved surface, not a new trust policy.
3. Triage four historical test-file Gitleaks candidates and one ignored local .env candidate using sanitized rule/path/line/commit metadata. Do not print/persist secret/Match/token/credential fields or dump .env/history matches. Establish factual fixture-vs-credential status without leaking contents; test filename alone is insufficient. Current ignored local configuration is not proof of committed exposure. No Git history rewrite, credential removal/rotation, changes to live .env, secret-scanning exemptions/policy/allowlist or residual-risk acceptance. If existing authority cannot settle gate treatment, return explicit Decision to Human/security/source owner; do not silently force a clean exit.
4. Report migration tracking limitation only; no local populated DB mutation, migrate-down, migrate-up or role/provisioning change in this patch.

Allowed production edits: two error literals in oidc.go; focused backend auth test(s) needed for cookie applicability evidence. http.go changes require a bounded semantics-preserving proposal reconciled by Orchestrator before expansion beyond initial literals/tests. No dependency pins/lockfiles, frontend, API, Compose, Makefile/README gate policy or upstream Product/Techplan/task/manifest writes. Preserve unrelated worktree changes.

## Verification and execution envelope

PREAUTHORIZED: read bound source/safe captures; bounded lint literal patch and relevant focused tests; writable isolated caches; focused staticcheck and affected Go constructor/protocol/cookie unit checks under Go1.26.9, readonly module mode, checksum verification enabled. Run newly authored tests enough to establish executable oracle. Record exact command/env/tool version/exit/capture hash. No test counts or no-test contract tags relabeled as contract proof.

Security triage here is targeted inspection of existing captured findings and justified focused cookie oracles, not a broad security sweep or final Testing acceptance. If any specialized security-class verification is necessary, explicitly justify/list deviation in Verification scope confirmation; default broad gosec/gitleaks/vulnerability/race/concurrency make verify belongs to fresh independent Testing. Do not replay final suite solely for confidence or claim make verify PASS from separate commands. Report confirmed, assumed, simulated, not-run and deferred separately.

Bounded environment recovery: on sandbox/cache/module/checksum network restriction, preserve initial failure and use one harness-approved scoped retry for the same necessary focused command with writable caches; successful same-command retry permits focused Build gate continuation only while this occurrence is still active and no substantive finding/authority stop exists. Do not disable GOSUMDB, alter pins or clear global caches. Retry failure => terminal BLOCKED with evidence; no post-terminal reopening. Follow bound network/scanner/dependency guides where relevant. Scanner final gates remain Testing-owned.

ORCHESTRATOR_DECISION: returned patch delta/dispositions, additional scoped source write route, Review applicability for actual change, blocker reconciliation and fresh independent Testing preparation.
HUMAN_REQUIRED: exact model approval/launch, material security/policy or risk choice, protected credential access, non-disposable DB/privilege/migration operations, credential rotation/history repair, scan allowance policy. Stop before those actions. Do not ask for secrets in chat.

## Output / return route

Write `RUN_PATH/evidence/patch-report-1.md`, candidate-dispositions with safe refs where needed, command captures/hashes and changed-file before/after hashes. Follow canonical report sections including What changed, Tests run, Verification scope confirmation, Contract check, Deferred/not tested, Flagged for Techplan/Testing and exactly one structured Phase handoff per overlay. Unresolved candidate treatment => report partial/blocked; never mark required final gate passed.

Requesting phase: Testing. After bounded patch, return to Orchestrator for fresh independent Testing on new bindings. No automatic full Code Review loop: material security/source changes may require review as determined from returned delta. No dependency commit, B-T2-002/B-T2-003 closure, T2 acceptance, T3 dispatch/readiness or whole-C1 claim.

## Human-Assisted kickoff (only after exact approval)

Open a fresh Participant Session using gpt-6-sol / medium in `/home/anhar-solehudin/kencleng-workspace/kencleng`. Prompt: `Execute RUN-C1-ENG-BUILD-T2-PATCH-001 from .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-PATCH-001/invocation.md. Read and verify its input-bindings.json before any execution. Follow current canonical Build/Patch authority and the bounded envelope. Return patch-report-1.md with exactly one Phase handoff and final SHA-256; stop on authority/input mismatch.`

On block/completion return report path + digest, actual changed file hashes, unresolved disposition/Decision and verification limits. No prior Session continuation.
