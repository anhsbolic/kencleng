# Runtime harness full-ID patch routing — 2026-10-13

## Review reconciliation

Terminal report `RUN-C1-ENG-CODEREVIEW-T2-RUNTIME-HARNESS-CLEANUP-PATCH-001/evidence/review-findings-1.md`, SHA-256 `e1030123ede0df6f9c9620a2f5d18d706cb58ef0e3080f2eb4cabf0322173ab0`, matches the Human-provided identity, has exactly one `## Phase handoff`, and all 25 assignment bindings match. The four-pass read-only Review returned Request changes for blocking F1: cleanup checks that inspected `Id` starts with the discovered Podman ID but does not require a valid full ID before deletion. No source edit, test, gate, Podman, or runtime action occurred.

The finding is within the same bounded disposable-resource cleanup concern as the preceding F1 Review, but it narrows a newly implemented branch: previous Build/Patch materially added exact-name reconciliation and four mock cases; the latest Review identified that full inspected-ID validity was still missing. Preserve both terminal Reviews, the prior Build report, and the prior patch report unchanged.

## Lightweight Loop Health trigger check

Before proposing another ordinary Run, the Orchestrator checked current experimental Loop Health criteria for `WU-C1-ENG-001 / T2 / disposable harness resource cleanup`. One completed Review → patch Build → Review cycle exists; the prior patch added material implementation and four passing Builder-reported mock cases. No HARD H1/H2/H3 trigger is established. No second WARNING trigger is established: no two material planning reconciliations, no completed two-alternation Review/Build bounce, no two routing recovery incidents, no two consecutive low-information re-entry Runs, no owner bounce, and no unexplained Finding-ID churn. The guidance's explicit single Review/patch/Review non-trigger applies; recommendation threshold is not met. This is detector reasoning only: no Loop Health Audit is created, no independent audit verdict is claimed, and the Work Unit is not declared STALLED.

The next premise is materially changed and concrete: validate a full 64-hex inspected ID and discovered-ID prefix before removal; add short/malformed-ID negative assertions. This is a bounded implementation defect under the existing accepted cleanup plan, not a new authority or Product decision.

## Fresh Build/Patch route

Prepared `RUN-C1-ENG-BUILD-T2-RUNTIME-HARNESS-FULL-ID-PATCH-001` with a new Implementer/Session, exact two-file scope and 27 assignment-defining bindings. Selected `gpt-6-luna / medium` from the current Human-owned registry (`approval_required: false`); prior model approvals do not transfer. Scope allows edits only to `backend/scripts/auth_harness.py` and `backend/scripts/test_auth_harness.py`, plus this Run's report. It authorizes only `python3 scripts/test_auth_harness.py` from `backend/`; no real Podman/container/database/fixture/provider/browser/runtime, broad gate, scanner, or network action.

The Builder must require inspect ID exactly 64 hexadecimal characters, validate the discovered 12–64-hex value as its prefix, validate any `podman run` expected ID as full, and fail closed with no deletion on malformed/short/inconsistent identities. Preserve the four existing command-mock scenarios and add short/malformed inspect-ID tests asserting no `podman rm`. Report only command-boundary evidence and exactly one Phase handoff.

Human mechanical dispatch is pending. After the terminal patch report, Orchestrator will reconcile exact bindings/delta and decide whether a fresh targeted Review materially closes the remaining safety claim. Testing applicability will be assessed separately from its claim/evidence needs; no phase is automatic by chronology. T2 runtime/DB acceptance remains unverified; `B-T2-002` / `B-T2-003` and both verification holds remain. No T2/T3 acceptance/readiness or C1 completion is claimed.
