# Run Invocation — RUN-C1-ENG-CODEREVIEW-T2-PATCH-001

## Identity / status

WORK_UNIT_ID: WU-C1-ENG-001. RUN_ID: RUN-C1-ENG-CODEREVIEW-T2-PATCH-001.
WORK_UNIT_PATH: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`. RUN_PATH: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-PATCH-001`. ARTIFACT_TARGET: none; findings/diagnostic evidence only.
ROLE / Profile: Reviewer / generic independent Reviewer, backend-security specialization; canonical four-pass Code Review plus targeted candidate-attribution diagnostic.
PARTICIPANT: PARTICIPANT-C1-ENG-T2-PATCH-REVIEWER-001.
SESSION: SESSION-C1-ENG-T2-PATCH-REVIEWER-001. SESSION_TRANSITION: FRESH for independence; no hidden prior executor context.
SELECTED_MODEL / REASONING_EFFORT: gpt-6-sol / medium.
MODEL_APPROVAL: APPROVED by Anhar on 2026-10-12 (“approve bro”) for exact RUN-C1-ENG-CODEREVIEW-T2-PATCH-001 / gpt-6-sol / medium. Prior Run approvals do not transfer. READY_FOR_HUMAN_DISPATCH; Reviewer not launched. Approval does not establish runtime identity or execution.
MODEL_ROUTING_RATIONALE: independent security applicability reasoning and disputed history/scanner coordinates require strong cross-cutting source/diagnostic reasoning. Medium lowest supported sufficient Sol effort; model strength does not supply missing policy/authority.
TARGET_REVISION / branch: 9aabf9c580685271982976dfa9e1cb2b18df0288 / pilot/3-c1-engineering, live patched worktree bound below.
WORKFLOW_REVISION: f6856df25a13b004e5fb242eb71f3ecd80728f70.
HARSCODE_WORKSPACE_ROOT: ../harscode-workspace. Working directory: /home/anhar-solehudin/kencleng-workspace/kencleng.
COMMUNICATION_LANGUAGE: Bahasa Indonesia; no separate communication profile.

## Assignment / scope

Exact `input-bindings.json` SHA-256 `8a33c2612c7d2a72a2f862dc1adb08d80f3de5ccc0c33ae87ebc953496ee17b5` binds 14 assignment-defining targets/plan/task/current gate/failure inputs. Preflight before work; assignment drift stop/report. WU/Events/Control Tower, `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/coordination/2026-10-12-patch002-terminal-reconciliation.md`, lifecycle approval/dispatch and Human registry are mutable context references, not byte-locked; semantic contradiction still fails closed. Ordinary applicable guidance current-effective, observed workflow HEAD is provenance.

Run all four canonical passes (Safety, Quality, routed Stack-Specific, Consistency) against **current two-file patch scope**: oidc.go two lowercased error literals compared with git HEAD baseline (same pre-patch hash a9a89117ba682e3963f5a02ec526d566763e1f991994cd5526fb01c60016a812), and new cookie_test.go absent from that baseline. Inspect actual diff/source, not merely report. Context http.go and gate/policy are read-only targets for questions below; no broad re-review of prior six-file Build scope.

Read current root/backend AGENTS, canonical workflow/4-code-review-prompt.md, guidelines/checklist, orchestrated-run-overlay.md, orchestration/run-contract binding/evidence/terminality, Approved Techplan/T2, current Makefile/README and C1 auth environment guide. Route best practices through AGENTS + Go index and applicable REST cookies/secrets guidance. No unrelated Exploration or full historical source ingestion.

Questions:
1. Independently assess PATCH-002 cookie oracle/behavior limits and G124 locations 49/158/277. Distinguish validated production/dev construction from arbitrary mutable config and unexecuted logout handler. Do not infer clean scanner/final runtime acceptance from unit assertions. State exact applicability and whether targeted source refactor vs narrow scanner policy disposition would be needed; policy options are proposals only.
2. Independently inspect three historic fixture assessments against exact commit/path references without logging literals/secret/Match fields. Test filename alone insufficient; disposition requires safe usage/rationale.
3. Reconcile unmapped history candidate ef428be3742eddfc59b86fc205bb03624e669e82 / google_oauth_test.go:487. Preserved metadata and counts are insufficient. Recover exact attribution from available safe original evidence first; if unavailable, perform only the justified bounded diagnostic below. Distinguish scanner diff coordinate from blob coordinate; never guess offset. Retain unresolved state if reproduction does not map.
4. Produce decision-ready minimal accepted-disposition/policy/remediation options for Human/security owner with exact candidate scope, rationale and limitations; **no allowances, risk acceptance or source/policy write**. Code Review verdict covers patch correctness only; candidate gate treatment separately unresolved/proposed. No final Testing or whole-C1 acceptance.

## Execution envelope

PREAUTHORIZED: read source/history with redaction before output; independent reasoning; bounded targeted reproduction only when necessary for concrete Review uncertainty. No production, Makefile/README, config, .env, git history, credential or lifecycle writes.

For attribution only: read-only repository mounted configured Podman/Compose Gitleaks v8.30.1 route, one bounded history diagnostic matching current backend-history log scope (all/full-history/backend), timeout <=180s. It is a **targeted security diagnostic**, not final gate certification. Use redacted JSON captured privately to a disposable temporary report; sanitize before any tool display/durable write to rule/path/commit/diff-or-blob line/fingerprint and safe rationale, omitting Secret/Match/Author/Email and quoted credential literals. Never mount/copy .env for this diagnostic. Preserve meaningful stdout/metadata/failure facts and exact command/tool/scope/exit provenance; no full raw secret report retained. Stop/cleanup disposable report/container only; do not delete original evidence. No worktree scan/provider/DB/browser run or broad security matrix. Fresh counts may differ with refs; explain without discarding historical candidates.

A narrowly targeted gosec auth-package reproduction is permitted only if resolving a concrete scanner uncertainty needs it; identify actual version, preserve necessary candidate details and report specialized execution rationale. Do not replay all lint/unit/race/security/make verify for confidence. Runtime Test Focus and real-provider prerequisites remain Testing-owned.

Bounded active recovery: preserve sandbox/cache/network failure; one same-command harness-approved scoped retry may continue while active if successful, with no authority/substantive stop. Failure/terminal handoff closes occurrence; subsequent retry needs fresh Run. No checksum bypass/pin changes/global cache clearing.

ORCHESTRATOR_DECISION: reconcile verdict/proposals/ownership, route specific policy decision or fresh Build remediations and scoped Testing readiness.
HUMAN_REQUIRED: exact model approval/launch; any allowance/suppression/security gate/risk policy decision, material auth change, privileged/credential/history/provisioning action. Approval for this Review is not approval of any eventual proposal.

## Outputs / sufficient evidence

Write RUN_PATH/evidence/review-findings-1.md, canonical four sections, executed verification, Verdict and exactly one Phase handoff. Separate confirmed/simulated/reported/not-run/deferred, code verdict, candidate attribution and owner decisions. Include safe exact source/revision/finding pointers and selected disputed/failure-defining capture needed for independent inspection. No blanket stdout/stderr/cache/source copies. Reference original evidence without duplicate copying. patch-plan-1.md only if source/policy remediation needs Build; frame policy as proposed pending Human approval.

Return report path + SHA-256, exact candidate mapping/dispositions or unresolved facts, minimal Human decision options, actual commands/provenance and limitations. Reviewer cannot close B-T2-002/B-T2-003, accept dependencies/commit, mark T2/T3 ready or dispatch Testing. Before future final/real-runtime Testing, Orchestrator resolves/narrows known material prerequisites; Human-performed Google is Testing-owned/Human-attested, not scanner clearance.

## Human-Assisted kickoff after approval

Fresh Reviewer/backend-security Codex Session using gpt-6-sol / medium in `/home/anhar-solehudin/kencleng-workspace/kencleng`:
```
Execute RUN-C1-ENG-CODEREVIEW-T2-PATCH-001 from .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-PATCH-001/invocation.md. Verify assignment bindings, read current canonical Review authority and perform the four passes plus bounded D-PATCH002-01/02 diagnostic questions. Do not modify source/policy or replay final Testing. Return review-findings-1.md SHA-256 and exactly one Phase handoff with concrete disposition/decision options.
```
