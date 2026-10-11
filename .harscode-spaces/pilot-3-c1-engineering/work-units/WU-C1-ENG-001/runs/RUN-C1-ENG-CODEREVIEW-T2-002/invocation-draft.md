# Draft Run Invocation — RUN-C1-ENG-CODEREVIEW-T2-002

## Purpose and identity

- **Work Unit / proposed Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-002`.
- **Phase / Role:** Independent Code Review — T2 dependency + G1 session-lock correction / Reviewer.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-CODEREVIEW-002` / `SESSION-C1-ENG-T2-CODEREVIEW-002` (fresh; no continuation of T2-001 context).
- **Dispatch posture:** Human-Assisted; not dispatched.
- **Proposed model / effort:** `gpt-6-sol` / `medium`.
- **Approval:** Pending for this exact new Run. Approval of T2-001 does not transfer.
- **Target revision / branch:** `059d145d5df64fcda427e5a625486a8e95ce1fd8` / `pilot/3-c1-engineering`.
- **Workflow revision:** `91a97af8ba64f98b0a692a83fbff866a210c45c0`.

## Reason for new Run

`RUN-C1-ENG-CODEREVIEW-T2-001` is terminal **STALLED at preflight**; its report `evidence/review-findings-1.md` SHA-256 is `0cd4eaf0a98cadae5b907f15a0ec5f7fa2de43cb19b2add74e180aef52fe1d13` and contains exactly one Phase handoff. Six implementation file hashes matched its Invocation. Three coordination authority hashes (WU, Events, Control Tower) changed after the Invocation was prepared, when lifecycle state was updated to record approval/readiness. The Reviewer correctly stopped. No four-pass review or code verdict was issued.

Current Harscode `workflow/orchestrated-run-overlay.md` §Chronology and re-entry requires a new Run ID, Participant identity, fresh Session/context, and Run path when returning after an execution occurrence ended. This is a new independent review of the same implementation snapshot, not a resumption of T2-001.

## Exact implementation scope

Review only the following six files against target revision `059d145d5df64fcda427e5a625486a8e95ce1fd8`. The three Go test files are untracked in the baseline and must be inspected as current full-file additions.

| File | SHA-256 |
|---|---|
| `backend/go.mod` | `90b7c1c6e870e32828fa5932a6c4e9a28a65d0671a74bdb807ec29dbc5d27cae` |
| `backend/go.sum` | `6923543aab63e09da2b894699afb70b2ff1072a2eb014e063c8542ffed7deaa0` |
| `backend/internal/platform/auth/store.go` | `eb70175ebe756698a3a83b627e0fd6606928d8ac4c355f9dd8a05dab705552b3` |
| `backend/internal/platform/auth/store_test.go` | `bac0d93f0aa6af1591b10dc9a1df81fba0bf77ea2bdb5a165c129e2611795d60` |
| `backend/internal/platform/auth/oidc_test.go` | `2ffd30f781b3aeaa09f6beb0fddd00bea90abcab99b221e126867acfc8331ab4` |
| `backend/internal/platform/auth/store_integration_test.go` | `1b38058a177e0df5f87cf6c732478369e8f798265f2ad91c0762a1ea1879b1e0` |

These hashes were rechecked after T2-001 stalled and match. Recheck all six hashes, branch, target HEAD, and full changed-file scope at dispatch. If any drift, stop and re-ground before review. Do not include orchestration/docs/instruction changes as implementation diff; read target instructions as authority.

## Stable authority and evidence inputs

- Approved Techplan `techplan/techplan.md`, SHA-256 `6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79`.
- Current T2 task `techplan/tasks/T2-backend-identity-session.md`, SHA-256 `dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37`.
- Accepted manifest `techplan/tasks/manifest.md`, SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.
- Initial terminal Build report, SHA-256 `4fc08a8cb5533ebf702d364c86be11ce06b0937837ed8941cac828d5e3436191`.
- Human-authorized retry-1 Build report, SHA-256 `4a63608a42dd89c9978c62161b0f6924700115fd0ed06eeed9af42aa04d6b326`.
- Retry-1 `authorization.md`: `3c0b06c256b773124fcb9af01207eadc570d90084a1577c551d5a783b1aabb98`; `isolated-gate.md`: `0bbe936da2225601f2e9253d6dc08d432f6044c92da3009109974e99de4da012`; `actual-gate.md`: `c3a7f83133ae9e016c1451a9408bab66d3c8fff9492cde608e8177de9e19c063`.
- Retry-1 `changed-file-hashes.json`: `5daa2a9b25d13b706bdc8aa5f78ab04894e7cb8523fd031c6661c910328a404d`; `capture-hashes.json`: `f5913ccb140311022a487704fb57f9e3025993bde086eced2e9af45658c4cad7`; `final-source-hashes.json`: `9db9b5239ef9be3361d705e2d518d05b895e270579928e2b2c0896ce6d5a1ecc`; `final-diff.patch`: `b26f0a28185d57acff16c668240fd7b20ea3aacaaf23b53d2d21a0be36e84ea4`.
- Prior stalled Review report `RUN-C1-ENG-CODEREVIEW-T2-001/evidence/review-findings-1.md`, SHA-256 `0cd4eaf0a98cadae5b907f15a0ec5f7fa2de43cb19b2add74e180aef52fe1d13`; use only to understand why that occurrence stopped, not as Review evidence.
- Target instructions: root `AGENTS.md` SHA-256 `d6bb74b1ffa8cbd1a115006191a40574b5f16508d22477d38e09744681367462`; `backend/AGENTS.md` SHA-256 `817d01ab80598637644c64be9bfd7ec18a9dc5e3f0d22485d505f0760f72769f`.
- Human-owned model registry `.harscode-spaces/.local-config.yaml` SHA-256 `ffbd0521361eac90e82dbee98ba59f01f3e62641afd79bf8e71c4fc8a0e633ce`; `gpt-6-sol` supports `medium` and has `approval_required: true`.

### Coordination bindings to finalize after model approval

Before writing the final `invocation.md`, record approval for this exact Run in the existing durable lifecycle owners (WU/Events/Control Tower), then capture their resulting identities below. Do not subsequently edit those three coordination artifacts while this Invocation is being consumed; any required lifecycle change means stop and re-ground again.

- Current WU SHA-256 after approval: **to be captured in final Invocation**.
- Current Events SHA-256 after approval: **to be captured in final Invocation**.
- Current Control Tower SHA-256 after approval: **to be captured in final Invocation**.

## Four-pass review instructions

Follow current `../harscode-workspace/workflow/4-code-review-prompt.md`, `4-code-review/guidelines.md`, `4-code-review/checklist.md`, and `orchestrated-run-overlay.md`. Re-read and bind the exact current hashes in final Invocation. Use `best-practices/AGENTS.md`, `best-practices/index.md`, and the Go index to route only relevant sources (dependency/supply chain, auth/session/token, context/error, concurrency/integration tests, PostgreSQL transaction/locking); cite each source used for any finding. Re-ground exact current WU/Events/Control Tower hashes in final Invocation after approval.

Run four independent passes in order: Safety, Quality, Stack-Specific Best Practices, Consistency. Inspect exact current diff and source, approved dependency pins/graph evidence, public `LockSession` signature/time semantics after row-lock acquisition, cancellation/error/resource behavior, auth/session invariants, and what the disposable-PostgreSQL tests do and do not establish. Build captures are claims to check, not a Review verdict. Keep configured Google login, browser/HTTPS-cookie behavior, deployed PostgreSQL/SCRAM/restricted credentials, concurrent upsert/session contention, and independent Testing unverified unless separately evidenced by their owners.

Use targeted checks only to answer a concrete Review uncertainty; do not repeat broad Build/Testing suites. No production/Techplan/task/manifest/Product/blocker edits. Write only `evidence/review-findings-1.md` (and `patch-plan-1.md` only if code changes are required), with all four pass sections, verification record, verdict, and exactly one structured `## Phase handoff`. Do not close B-T2-002/B-T2-003 or claim T2/T3 readiness. Reconcile the Review verdict before dispatching Testing.

## Proposed dispatch envelope

- **WORK_UNIT_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-002`.
- **ARTIFACT_TARGET:** none; findings are Run-owned evidence only.
- **PREAUTHORIZED:** bound read-only review and Run-owned findings/handoff only.
- **ORCHESTRATOR_DECISION:** reconcile findings and route next owner.
- **HUMAN_REQUIRED:** approval of `gpt-6-sol / medium` for this exact Run before final Invocation/dispatch; material semantic/security choice if newly exposed.
- **Status:** draft package only; no Reviewer action, test, or review result.
