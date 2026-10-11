# Run Invocation — RUN-C1-ENG-CODEREVIEW-T2-001

## Identity and route

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-001`.
- **Phase / Role:** Independent Code Review — T2 dependency + G1 session-lock correction / Reviewer.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-CODEREVIEW-001` / `SESSION-C1-ENG-T2-CODEREVIEW-001` (fresh, independent from Implementer).
- **Dispatch:** Human-Assisted; exact model approval recorded; `READY_FOR_HUMAN_DISPATCH`. Participant is not dispatched.
- **Language:** Bahasa Indonesia; retain canonical technical identifiers.
- **Harscode root:** `../harscode-workspace`; target repository `/home/anhar-solehudin/kencleng-workspace/kencleng`.
- **Target revision / branch:** `059d145d5df64fcda427e5a625486a8e95ce1fd8` / `pilot/3-c1-engineering`.
- **Workflow revision:** `91a97af8ba64f98b0a692a83fbff866a210c45c0` (Harscode worktree clean at preparation).

## Review scope and baseline

Review the current diff against the six exact Build retry-1 files below. Tracked source/dependency diff is measured against target revision above. Three new Go test files are untracked at preparation and must be inspected as complete file contents against their stated purpose; do not omit them because `git diff` does not show untracked files.

| File | SHA-256 |
|---|---|
| `backend/go.mod` | `90b7c1c6e870e32828fa5932a6c4e9a28a65d0671a74bdb807ec29dbc5d27cae` |
| `backend/go.sum` | `6923543aab63e09da2b894699afb70b2ff1072a2eb014e063c8542ffed7deaa0` |
| `backend/internal/platform/auth/store.go` | `eb70175ebe756698a3a83b627e0fd6606928d8ac4c355f9dd8a05dab705552b3` |
| `backend/internal/platform/auth/store_test.go` | `bac0d93f0aa6af1591b10dc9a1df81fba0bf77ea2bdb5a165c129e2611795d60` |
| `backend/internal/platform/auth/oidc_test.go` | `2ffd30f781b3aeaa09f6beb0fddd00bea90abcab99b221e126867acfc8331ab4` |
| `backend/internal/platform/auth/store_integration_test.go` | `1b38058a177e0df5f87cf6c732478369e8f798265f2ad91c0762a1ea1879b1e0` |

Do not include existing root/backend AGENTS, runbook, scanner, control-tower, or orchestration edits in the implementation diff; read the applicable instructions as review authority. Recheck all six identities, target HEAD/branch, full changed-file scope, and bound authority before review and before writing terminal evidence. Stop and ask Orchestrator to re-ground the Run if a scoped file/hash or authority changes.

## Bound authority and evidence

All paths are project-root-relative unless noted.

- Approved Techplan `techplan/techplan.md`, SHA-256 `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79`.
- Current T2 task `techplan/tasks/T2-backend-identity-session.md`, SHA-256 `dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37`.
- Accepted manifest `techplan/tasks/manifest.md`, SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.
- Initial terminal Build report (retained BLOCKED), SHA-256 `4fc08a8cb5533ebf702d364c86be11ce06b0937837ed8941cac828d5e3436191`.
- Same-session Human-authorized retry-1 report, SHA-256 `4a63608a42dd89c9978c62161b0f6924700115fd0ed06eeed9af42aa04d6b326`.
- Retry authorization `evidence/retry-1/authorization.md`, SHA-256 `3c0b06c256b773124fcb9af01207eadc570d90084a1577c551d5a783b1aabb98`.
- Retry isolated gate `evidence/retry-1/isolated-gate.md`, SHA-256 `0bbe936da2225601f2e9253d6dc08d432f6044c92da3009109974e99de4da012`.
- Retry actual gate `evidence/retry-1/actual-gate.md`, SHA-256 `c3a7f83133ae9e016c1451a9408bab66d3c8fff9492cde608e8177de9e19c063`.
- Retry changed-file hashes, SHA-256 `5daa2a9b25d13b706bdc8aa5f78ab04894e7cb8523fd031c6661c910328a404d`; capture index, SHA-256 `f5913ccb140311022a487704fb57f9e3025993bde086eced2e9af45658c4cad7`; final source hashes, SHA-256 `9db9b5239ef9be3361d705e2d518d05b895e270579928e2b2c0896ce6d5a1ecc`; final diff capture, SHA-256 `b26f0a28185d57acff16c668240fd7b20ea3aacaaf23b53d2d21a0be36e84ea4`.
- Current Work Unit at preparation, SHA-256 `47b8db1cbbc3c09082f8e591f1553400447daf0c54c1d6203a26fd0e89b7ad22`; Events, `7c3a2ed42a7750067c333a9e79a557348c0ce292edb6d78bb6bf24ceb9433606`; bounded Control Tower, `f7abf20f887e30542bc8d6d3ae88d85ed25c2e975a828e97541602b83f336d4c`.
- Target instructions: root `AGENTS.md`, SHA-256 `d6bb74b1ffa8cbd1a115006191a40574b5f16508d22477d38e09744681367462`; `backend/AGENTS.md`, SHA-256 `817d01ab80598637644c64be9bfd7ec18a9dc5e3f0d22485d505f0760f72769f`.

## Review method

Follow current `../harscode-workspace/workflow/4-code-review-prompt.md`, `workflow/4-code-review/guidelines.md`, `workflow/4-code-review/checklist.md`, and `workflow/orchestrated-run-overlay.md`. Their preparation hashes are respectively `d152c1419ea8b0e67a08d1ec2cec7c0235cc15b0fccbd032722aea1ed2332f8a`, `513f145a5bc1980589a4967f6a0894a3f59f089ade4fa199ba252b2604a735c0`, `5621854c75f8834ff9c32a8a67a326ab03388e11cceac95f0f2c6bdb0e03e079`, and `f806673aa5e7d1d82dae9f433ec54c3562e15ce778b709a141c57d3d22dfaf87`. Use current best-practices routing from `best-practices/AGENTS.md`, `best-practices/index.md`, and `best-practices/go/index.md`; preparation hashes `9dc2b38e5175f4da7e21cbd7c78e4e8aaea07d2325510f24cd9527625980fc6e`, `876a14da409b3a486a790daee542fbafd02ecae82283b892dd189023cfbaaa34`, `0b51468a6cf4bdc1037422376ef2aeb6884898e0d4a38c9a79b594f4f10bebf6`. Relevant clue files include Go dependency/supply-chain, authorization/IDOR, JWT/token lifecycle, context/error handling, concurrency/integration testing, and PostgreSQL transactions/locking. Reviewer must route from the current index, open only matching guidance, and cite sources for any finding.

Run the four independent passes in order: Safety, Quality, Stack-Specific Best Practices, Consistency. Independently assess the exact approved dependency pins/graph evidence and their interaction with T2 auth code; public `LockSession` signature compatibility and actual time sampling only after the locked row is returned; cancellation/error/resource semantics; auth/session/token invariants; test determinism/coverage and accurate boundaries of disposable PostgreSQL evidence; repository instructions and local patterns. Treat Build report/captures as claims to inspect, not as Review conclusions. Do not claim real provider, browser, deployed PostgreSQL, SCRAM/restricted credentials, concurrent-upsert or real session-contention acceptance from Build evidence.

Review is primarily reasoning/read-only inspection. Run targeted checks only for a concrete review uncertainty, record why and the result, and do not replay broad Build/Testing matrices. Do not edit production, Techplan/task/manifest, blockers, Product authority, or dispatch Testing from this Run. Write only `evidence/review-findings-1.md` (and `patch-plan-1.md` only if code changes are required), with all four pass sections, verification record, verdict, and exactly one structured `## Phase handoff`. Recommend the next route from findings. The Review does not close B-T2-002/B-T2-003 or establish T2/T3 readiness.

## Model and approval

- **Selected model / effort:** `gpt-6-sol` / `medium`.
- **Rationale:** independent review spans auth/session security, row-lock correctness, concurrency concerns, and dependency/supply-chain compatibility. The security and cross-cutting scope warrants Sol; medium is the lowest supported effort that fits this bounded but multi-surface review.
- **MODEL_APPROVAL:** APPROVED by Anhar on 2026-10-11 (“approve bro”) for `RUN-C1-ENG-CODEREVIEW-T2-001` using `gpt-6-sol / medium`. This approval is Run-specific; Build and Loop Health approvals do not transfer. Approval establishes permission to launch this Run, not runtime model identity or Review outcome.
- **Registry basis:** Human-owned `.harscode-spaces/.local-config.yaml`, SHA-256 `ffbd0521361eac90e82dbee98ba59f01f3e62641afd79bf8e71c4fc8a0e633ce`; `gpt-6-sol` lists strong reasoning/coding/repository/architecture/cross-cutting capabilities, supports `medium`, and has `approval_required: true`. Current Harscode `orchestration/local-runtime-config.md` policy §Model selection requires stopping before dispatch for such a model and obtaining explicit approval for this Run.
- **Runtime identity:** not established by preparation; do not infer it.

## Dispatch package

- **WORK_UNIT_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-001`.
- **ARTIFACT_TARGET:** none; findings are Run-owned evidence only.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-CODEREVIEW-001` / `SESSION-C1-ENG-T2-CODEREVIEW-001`.
- **PREAUTHORIZED:** read the bound diff, current authority, and routed best-practice sources; perform focused read-only review; write Review findings/handoff under this Run only.
- **ORCHESTRATOR_DECISION:** reconcile Review verdict and route the next owner.
- **HUMAN_REQUIRED:** exact model approval before Human launch; later Testing evidence and owning acceptance; any newly exposed material semantic/security decision.
- **Status:** READY_FOR_HUMAN_DISPATCH; prepared, not dispatched; no Reviewer action or test is claimed.
