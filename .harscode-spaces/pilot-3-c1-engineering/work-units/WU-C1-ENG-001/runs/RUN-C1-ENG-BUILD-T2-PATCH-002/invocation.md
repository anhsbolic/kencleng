# Run Invocation — RUN-C1-ENG-BUILD-T2-PATCH-002

## Identity and preparation posture

- WORK_UNIT_ID: WU-C1-ENG-001
- RUN_ID: RUN-C1-ENG-BUILD-T2-PATCH-002
- WORK_UNIT_PATH: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`
- RUN_PATH: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-PATCH-002`
- PHASE_ROUTE: fresh Build/Patch from terminal Testing finding/current supplemental diagnostic evidence. PATCH-001 was never dispatched and is superseded; no prior Run is reopened.
- ROLE / Profile: Implementer / generic Implementer, backend specialization (canonical Build/Patch plus routed Go/REST/secrets guidance; no separate Profile file).
- PARTICIPANT: PARTICIPANT-C1-ENG-T2-PATCH-IMPLEMENTER-002
- SESSION: SESSION-C1-ENG-T2-PATCH-IMPLEMENTER-002
- SESSION_TRANSITION: FRESH, independent from all prior executor conversations; CONTINUATION_CHECKPOINT: none.
- ARTIFACT_TARGET: oidc.go two error literals and narrowly affected backend auth tests, within patch-scope.md. Results belong under RUN_PATH.
- SELECTED_MODEL / REASONING_EFFORT: gpt-6-sol / medium.
- MODEL_APPROVAL: PENDING exact Human approval for PATCH-002. Prior Run model approvals do not transfer. Status WAITING_HUMAN; prepared only, no dispatch/execution claimed.
- MODEL_ROUTING_RATIONALE: mechanical lint change plus bounded security applicability reasoning across conditional cookie flags, deletion-cookie behavior and historical repository findings needs strong cross-cutting reasoning; medium is lowest supported Sol effort sufficient for this scope. Model choice cannot solve missing policy/authority/prerequisites.
- TARGET_REVISION / branch: 069a84b5d69b089707a76a653437f6de8d9c09d5 / pilot/3-c1-engineering.
- WORKFLOW_REVISION / branch: f6856df25a13b004e5fb242eb71f3ecd80728f70 / pilot/orchestrator-v0.1-c1.
- HARSCODE_WORKSPACE_ROOT: ../harscode-workspace
- Working directory: /home/anhar-solehudin/kencleng-workspace/kencleng
- COMMUNICATION_LANGUAGE: Bahasa Indonesia; COMMUNICATION_PROFILE_PATH: none.

## Assignment inputs and binding discipline

`input-bindings.json` SHA-256 `10930892facc89fd76429f910e6f64dc0fa2b8a7c53a2fd815d0fbff8d444db7` has 18 exact assignment-defining inputs: approved plan/task/topology, reconciled patch scope, affected source/test/compiler baseline, actual gate/guide/contract semantics and materially relied-upon failure/triage evidence. Check at preflight before mutation; drift changing the assignment => stop for Orchestrator re-grounding. Source modifications authorized by this Run are expected after that baseline check; record actual delta rather than treating own patch as preflight drift.

Mutable context references, **not exact bindings**: WU current state, append-only Events, Control Tower, `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/coordination/2026-10-12-patch-checkpoint-reconciliation.md`, PATCH-001/supersession.md, Human-owned `.harscode-spaces/.local-config.yaml` and model approval/dispatch metadata. Read for current meaning: ordinary approval/scheduling/Event/projection byte drift does not invalidate assignment; material authority/scope/approval contradiction still fails closed. Actual Human approval for the exact Run/model/effort is a dispatch precondition. Target/workflow HEADs are preparation provenance, not blanket whole-repo content locks. Do not mutate registry or lifecycle records as Participant.

Ordinary applicable instructions/workflow guidance remain **current-effective**, with observed workflow HEAD above; not exact-bound merely for freshness. Read root/backend AGENTS; Harscode AGENTS routers, orchestration/run-contract.md (binding/evidence/terminality), canonical workflow/3-build-prompt.md, 3-build/guidelines.md and checklist.md, orchestrated-run-overlay.md; Approved Techplan/T2 and patch-scope.md. Read current backend Makefile/README and C1 auth environment guide for gate/evidence boundaries. Route practices through best-practices/AGENTS.md, Go subindex and REST cookie/CSRF + secrets/logging guidance only as relevant. Material changed guidance that alters assignment/safety must be reconciled, not silently adopted or ignored.

PRIOR_ARTIFACTS: original TEST-T2-001 patch/report; PATCH-001 preserved supplemental testing-report.md/patch-plan.md, gosec.stdout and history-findings.json referenced in bindings. The two old patch plans are bounded by current patch-scope.md reconciliation; their ignored-local-env treatment is historical. Build T2-006 retry-1 / Review T2-002 are prior context only, not PATCH-002 tests/acceptance. Do not recreate dependency gates or read full historical execution dumps absent a concrete trigger.

Existing supplemental evidence remains canonical in PATCH-001 for historical preservation; it was not produced by PATCH-001 and must not be attributed to PATCH-002. Do not copy it into this RUN_PATH. This prepared Run has no execution evidence yet.

## Execution envelope and proportional verification

PREAUTHORIZED: inspect bound safe evidence/live affected source; perform patch-scope.md two literal changes; author/run focused constructor/protocol/cookie/config unit oracles only when needed for the affected claim; run focused staticcheck with Go1.26.9 and writable isolated caches. Preserve exact dependency/compiler baseline, readonly module mode, checksum verification. No need for redundant tests mirroring a two-literal edit. Newly authored nontrivial cookie tests must run enough to establish their oracle/executability, clearly labeled unit/fixture evidence.

Candidate triage is targeted use of existing captures and safe source/history inspection. A targeted `gosec ./internal/platform/auth` rerun is allowed only if needed to resolve applicability/current tool behavior; record its justification, actual version, outcome/candidate locations and specialized verification in canonical Verification scope confirmation. This explicit narrow exception does not authorize a broad scanner sweep or final clearance. Do not alter scanner configuration or add suppressions. The repository Gitleaks gate/final make verify, broad race/concurrency/security and real DB/provider/browser acceptance remain fresh independent Testing-owned. Use current commands from Makefile when final Testing is later assigned; make constituents or no-test contract tags do not substitute for asserted coverage.

Bounded environment recovery while active: preserve meaningful initial cache/sandbox/network failure; one harness-approved scoped retry of the same necessary focused command may continue the Build loop if it succeeds and no substantive finding/authority stop exists. Writable caches are permitted; no checksum bypass, dependency pin change or global-cache clearing. Follow current docs/project/go-network-recovery.md, go-dependency-verification.md and go-scanner-runbook.md if triggered. Retry failure => terminal report; **terminal closure monotonic**, no retry child/reopening after Phase handoff. A later Human retry intent requires another fresh Run.

ORCHESTRATOR_DECISION: actual delta/dispositions and additional remediation routing, Code Review applicability, independent Testing readiness/scope, blocker/acceptance reconciliation.
HUMAN_REQUIRED: exact model approval/launch; protected credential access, history repair/rotation, populated DB/migrations/roles, material security/scan-policy/risk choices outside approved G1. G1 already settled in Approved parent Resolved 9; do not ask it again or expand it. Stop on unresolved authority before production changes beyond the envelope.

## Evidence sufficiency / output

Write `RUN_PATH/evidence/patch-report-1.md` following canonical What changed, Tests run, Verification scope confirmation, Contract check, Deferred/not-tested, Flagged and exactly one structured Phase handoff per overlay. Report actual exposed provenance; never invent runtime model/Session identity. Candidate dispositions may live in the report or a separate referenced file when useful.

Preserve the **smallest sufficient** evidence: semantic change/disposition and rationale; materially relied-upon source/commit identities; safe command/tool/environment provenance and actual outcomes; selected failure-defining, disputed, security-sensitive or non-reproducible captures needed to substantiate the claim. Do not require blanket stdout/stderr retention, every temporary/cache/download/helper or duplicate capture. Identical raw captures may share one identity while execution facts remain distinct. Never drop uncertainty/failure facts or project-required evidence; never capture secret/Match/token values. Original referenced evidence remains immutable. A table with meaningful outcomes may suffice for routine focused successful checks; unresolved candidate details need safe exact refs/rationale, not count alone.

## Return to Testing with readiness guard

Requesting phase: Testing. Return actual patch/dispositions to Orchestrator, no automatic full Review loop. Material source/security delta may require separate Review per owning reconciliation. Before final/real-runtime Testing dispatch, resolve declared material prerequisites or explicitly narrow the claimed outcome; do not launch an unreachable acceptance Run. Restricted application DB/schema, disposable fixture identity, callback/consent/test-user/network/browser route, and production HTTPS are distinct prerequisites. Unknown safe availability can be probed when authorized; PATCH-002 does not provision or verify them.

Real Google private account interaction is Human-performed under project guide, still Testing-owned: bounded observable criteria before request, Human-attested result and separately reported agent-side evidence. Human participation is not environment failure. No credentials/tokens requested; no repeat-person/provider code-replay/concurrency/production-cookie sub-results inferred from a broad Human login PASS. Integration fixture, simulated provider/fake locks, real runtime and deferred T3/T4/T5 results remain distinct.

B-T2-002/B-T2-003 remain open; no dependency commit/acceptance, T2 acceptance, T3 readiness/dispatch or whole-C1 completion from this patch. Never replace required real/independent evidence with Human attestation where owning authority does not permit it.

## Human-Assisted dispatch package — only after exact approval

Fresh Codex Session, Implementer/backend Profile, gpt-6-sol / medium, working directory `/home/anhar-solehudin/kencleng-workspace/kencleng`. Participant/Session identifiers above are assignment identities, not exposed runtime proof.

Kickoff:
```
Execute RUN-C1-ENG-BUILD-T2-PATCH-002 from .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-PATCH-002/invocation.md. Check assignment-defining input-bindings.json before edits, read mutable coordination for semantic drift, and follow current Build/Patch authority, patch-scope.md and the execution envelope. Do not continue PATCH-001 or any terminated Run. Return patch-report-1.md with exactly one Phase handoff and SHA-256, actual delta, candidate dispositions/Decisions and verification limits.
```

Do not dispatch before exact Human approval. On block/completion Human returns terminal report path/digest and material result refs/remaining owner requests. No Testing launch automatic from patch completion.
