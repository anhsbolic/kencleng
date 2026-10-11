# Run Invocation — RUN-C1-ENG-CODEREVIEW-T2-002

## Identity and dispatch

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-002`.
- **Phase / Role:** Independent Code Review — T2 dependency + G1 session-lock correction / Reviewer.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-CODEREVIEW-002` / `SESSION-C1-ENG-T2-CODEREVIEW-002` (fresh and independent).
- **Dispatch posture:** Human-Assisted; `READY_FOR_HUMAN_DISPATCH`. Prepared, not dispatched.
- **Language:** Bahasa Indonesia; preserve canonical technical identifiers.
- **Target revision / branch:** `059d145d5df64fcda427e5a625486a8e95ce1fd8` / `pilot/3-c1-engineering`.
- **Workflow revision:** `91a97af8ba64f98b0a692a83fbff866a210c45c0` (clean at preparation).
- **WORK_UNIT_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-002`.
- **ARTIFACT_TARGET:** None; Review findings are Run-owned evidence only.

## Re-entry and preflight

Previous occurrence `RUN-C1-ENG-CODEREVIEW-T2-001` is terminal **STALLED at preflight**, not resumed. Its report `evidence/review-findings-1.md` SHA-256 is `0cd4eaf0a98cadae5b907f15a0ec5f7fa2de43cb19b2add74e180aef52fe1d13` and contains exactly one Phase handoff. Six implementation-file hashes matched; its WU/Events/Control Tower hashes differed after lifecycle approval updates. The Reviewer stopped before all four passes; no code verdict was issued.

Current Harscode `workflow/orchestrated-run-overlay.md` §Chronology and re-entry requires a new Run ID, Participant, fresh Session/context, and Run path after the prior occurrence ended. This Invocation is a new independent review of the same exact implementation scope.

At dispatch, recheck all bound hashes, target HEAD/branch, and the entire changed-file scope. Stop and re-ground if any scoped implementation file, bound authority, or relevant execution input has drifted. The bound coordination authority hashes below are the post-approval state; do not change those three lifecycle artifacts during this Run.

## Exact implementation scope

Review exactly the six files below against target revision `059d145d5df64fcda427e5a625486a8e95ce1fd8`. The three Go test files are untracked at baseline and must be inspected as complete additions.

| File | SHA-256 |
|---|---|
| `backend/go.mod` | `90b7c1c6e870e32828fa5932a6c4e9a28a65d0671a74bdb807ec29dbc5d27cae` |
| `backend/go.sum` | `6923543aab63e09da2b894699afb70b2ff1072a2eb014e063c8542ffed7deaa0` |
| `backend/internal/platform/auth/store.go` | `eb70175ebe756698a3a83b627e0fd6606928d8ac4c355f9dd8a05dab705552b3` |
| `backend/internal/platform/auth/store_test.go` | `bac0d93f0aa6af1591b10dc9a1df81fba0bf77ea2bdb5a165c129e2611795d60` |
| `backend/internal/platform/auth/oidc_test.go` | `2ffd30f781b3aeaa09f6beb0fddd00bea90abcab99b221e126867acfc8331ab4` |
| `backend/internal/platform/auth/store_integration_test.go` | `1b38058a177e0df5f87cf6c732478369e8f798265f2ad91c0762a1ea1879b1e0` |

Do not include existing AGENTS, scanner/runbook, Control Tower, or orchestration edits as implementation diff. Read target instructions as review authority.

## Current durable authority bindings

| Input | SHA-256 |
|---|---|
| WU `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md` | `e3590643a5b5c5c785944c67bd1c2a69fb54d10b58e41d67790761c2a7f0fff5` |
| Events `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/events.md` | `1df0344fd2b7cbafc813c4c87068637cbea84fe701fdcae63ceb1dac03a0ad67` |
| Bounded projection `docs/product/pilot-3-control-tower.md` | `4733d32a65ba887820777965e24a3d5b867a5e8cf2d573b620917b4b35e6d7d9` |
| Approved Techplan `techplan/techplan.md` | `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79` |
| Current task `techplan/tasks/T2-backend-identity-session.md` | `dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37` |
| Accepted manifest `techplan/tasks/manifest.md` | `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d` |
| Initial terminal Build report | `4fc08a8cb5533ebf702d364c86be11ce06b0937837ed8941cac828d5e3436191` |
| Human-authorized retry-1 Build report | `4a63608a42dd89c9978c62161b0f6924700115fd0ed06eeed9af42aa04d6b326` |
| Retry authorization | `3c0b06c256b773124fcb9af01207eadc570d90084a1577c551d5a783b1aabb98` |
| Retry isolated gate | `0bbe936da2225601f2e9253d6dc08d432f6044c92da3009109974e99de4da012` |
| Retry actual gate | `c3a7f83133ae9e016c1451a9408bab66d3c8fff9492cde608e8177de9e19c063` |
| Retry changed-file hashes | `5daa2a9b25d13b706bdc8aa5f78ab04894e7cb8523fd031c6661c910328a404d` |
| Retry capture index | `f5913ccb140311022a487704fb57f9e3025993bde086eced2e9af45658c4cad7` |
| Retry final source hashes | `9db9b5239ef9be3361d705e2d518d05b895e270579928e2b2c0896ce6d5a1ecc` |
| Retry final diff capture | `b26f0a28185d57acff16c668240fd7b20ea3aacaaf23b53d2d21a0be36e84ea4` |
| Prior stalled Review report | `0cd4eaf0a98cadae5b907f15a0ec5f7fa2de43cb19b2add74e180aef52fe1d13` |
| Root `AGENTS.md` | `d6bb74b1ffa8cbd1a115006191a40574b5f16508d22477d38e09744681367462` |
| `backend/AGENTS.md` | `817d01ab80598637644c64be9bfd7ec18a9dc5e3f0d22485d505f0760f72769f` |
| Human-owned `.harscode-spaces/.local-config.yaml` | `ffbd0521361eac90e82dbee98ba59f01f3e62641afd79bf8e71c4fc8a0e633ce` |

## Review guidance and method

Follow `../harscode-workspace/workflow/4-code-review-prompt.md` (SHA-256 `d152c1419ea8b0e67a08d1ec2cec7c0235cc15b0fccbd032722aea1ed2332f8a`), `workflow/4-code-review/guidelines.md` (`513f145a5bc1980589a4967f6a0894a3f59f089ade4fa199ba252b2604a735c0`), `workflow/4-code-review/checklist.md` (`5621854c75f8834ff9c32a8a67a326ab03388e11cceac95f0f2c6bdb0e03e079`), and `workflow/orchestrated-run-overlay.md` (`f806673aa5e7d1d82dae9f433ec54c3562e15ce778b709a141c57d3d22dfaf87`). Use `best-practices/AGENTS.md`, `best-practices/index.md`, and `best-practices/go/index.md` as routing authorities; hashes `9dc2b38e5175f4da7e21cbd7c78e4e8aaea07d2325510f24cd9527625980fc6e`, `876a14da409b3a486a790daee542fbafd02ecae82283b892dd189023cfbaaa34`, and `0b51468a6cf4bdc1037422376ef2aeb6884898e0d4a38c9a79b594f4f10bebf6`. Route only relevant entries and open matching Go/PostgreSQL practices.

Run four independent passes in order: Safety, Quality, Stack-Specific Best Practices, Consistency. Examine dependency baseline/evidence; public `LockSession` compatibility and time sampling after the locked row is returned; cancellation, error, and resource behavior; auth/session invariants; test determinism and limits of disposable-PostgreSQL evidence. Treat Build evidence as claims to inspect, not Review conclusions. Do not claim configured Google login, browser/HTTPS-cookie behavior, deployed PostgreSQL/SCRAM/restricted credentials, concurrent upsert/session contention, or independent Testing acceptance from this Build.

Review is primarily read-only reasoning. Use targeted checks only for a concrete Review uncertainty. Do not replay broad Build/Testing suites. Do not edit production, Techplan/task/manifest, Product authority, or blockers. Write only `evidence/review-findings-1.md` and `patch-plan-1.md` only if code changes are needed. Findings must include all four pass sections, verification, verdict, and exactly one structured `## Phase handoff`. Recommend the next route based on findings. Do not close B-T2-002/B-T2-003 or claim T2/T3 readiness.

## Model approval and execution envelope

- **Selected model / effort:** `gpt-6-sol` / `medium`.
- **Rationale:** independent review covers authentication/session security, PostgreSQL lock semantics, concurrency risk, and dependency/supply-chain compatibility; medium is the lowest supported effort for this bounded cross-cutting review.
- **MODEL_APPROVAL:** APPROVED by Anhar on 2026-10-11 (“approve bro”) for this exact Run and model/effort. The registry marks this model approval-required. Earlier Run approvals do not transfer.
- **Runtime identity:** not established by preparation; do not infer it.
- **PREAUTHORIZED:** inspect bound current diff/authority and routed guidance; perform focused read-only Review; write findings/handoff under this Run only.
- **ORCHESTRATOR_DECISION:** reconcile verdict and route next owner.
- **HUMAN_REQUIRED:** Human mechanical launch; later Testing/owning acceptance; any new material semantic/security decision.
- **Status:** READY_FOR_HUMAN_DISPATCH; prepared, not dispatched. No Reviewer action or tests are claimed.
