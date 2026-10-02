# Run Invocation — `TP-S2-003-003`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator setelah RV-S2-003-001 selesai dengan dua material/blocking findings. Human-facing prose: Bahasa Indonesia; preserve canonical terms/enums dan technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `TP-S2-003-003`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-003`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Resolve independent Techplan findings
- `PARTICIPANT_ID`: `P-S2-003-TP-003-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — meaningful Planner resolution re-entry setelah completed independent Review; new Participant/Session, preserve prior Draft/Reviewer history.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree; verify live source saat dispatch.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`; ordinary applicable guidance tetap current-effective.
- `MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: Financial range, transaction/idempotency ordering dan cross-domain authority membutuhkan reasoning/source fidelity; configured repository-work/reasoning capability pada high effort memadai untuk bounded resolution/proposal. Escalation hanya untuk demonstrated capability insufficiency, bukan missing authority/context.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Canonical Techplan synthesis sebagai finding-resolution pass; successor tetap Draft / In Review. Materiality declaration dan applicable independent re-review mengikuti resolution sebelum report/Human whole-Techplan approval.

## Current-effective inputs / PRIOR_ARTIFACTS

- Current Draft predecessor: `WU-S2-003/runs/TP-S2-003-002/techplan.md`, `launch-record.md` dengan phase handoff, dan Invocation. Preserve valid owner planning decisions/batasnya; jangan overwrite predecessor.
- Exact finding source: `WU-S2-003/runs/RV-S2-003-001/review-findings.md`, `handoff.md`, `launch-record.md`. Dua Blocking entries tidak diberi finding ID oleh Reviewer; reference exact title/location, jangan mengarang Reviewer-assigned ID.
- Full durable `WU-S2-003/runs/EXP-S2-003-001/evidence/` sesuai fresh-session canonical guidance.
- Approved `WU-S2-002/runs/TP-S2-002-015/techplan.md`, agreed Donation/Campaign D1 specs/invariants/features, authored `api/openapi/donation.yaml` dan referenced common components, monetary standard dan Authority Map.
- Current WU manifest/events/Work Graph/Control Surface/Outcome/tracker; WU-S2-005 Human-approved TP-S2-005-002/report TP-S2-005-003 dan pending Status-only TP-S2-005-004. Final authored Campaign contract belum accepted, scoped GET producer dependency tetap berlaku.
- Root/backend AGENTS, backend architecture dan live source/migration/consumer anchors yang relevan untuk findings. Read source facts, bukan hanya Review prose.
- Canonical `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`, full Techplan template/rules/guardrails; independent review prompt resolution/materiality rules, overlay/context dan applicable workflow/orchestration authorities. Matching best practices hanya sesuai material concerns.

Paths `WU-S2-*` relatif terhadap `.harscode-spaces/s2-guest-donation-truthful-state/`; lainnya relatif repository root. Overlay memetakan prior plan/Exploration inputs di atas dan output successor ke RUN_PATH.

## Task and completion condition

Resolve kedua material findings melalui successor execution-grade `RUN_PATH/techplan.md`, dengan traceable resolution/remaining gates dan Run-local phase handoff/provenance:

1. **Idempotent retries can be rejected after Campaign closure.** Reconcile admission/retry ordering dengan accepted same-key/same-payload return-original behavior dan atomic eligibility ordering untuk genuinely new intent. Preserve payload-conflict semantics, D1 financial invariants, Campaign-first lock ownership dan explicit concurrency/verification obligations; jangan weakening contract atau menyatakan empirical proof dari planning. Derive proposal dari accepted authority/current source, bukan meminta owner memilih ulang accepted idempotency rule.
2. **Unconstrained Donation values cannot all be reflected by the existing Campaign funding column.** Verify full cross-domain amount representability, cumulative funding, current storage/projection/consumer limits dan migration scope. Produce concrete alternatives, impact/risks, recommendation, owning-authority/source changes yang diperlukan, verification dan exact Human decision. Jangan diam-diam membuat cap, widen authority/source, menganggap no declared maximum sebagai infinite platform capacity, atau memilih policy/range dari implementation precedent. Anhar adalah project-wide monetary-standard owner serta current Slice-2 Campaign/Donation/API owner; feature policy/Product authority tetap dibedakan dari representation/storage authority.

Bila proposal monetary decision sudah decision-ready, fasilitasi exact bounded owner question langsung di Participant Session dan catat explicit answer/provenance. Jika membutuhkan missing external authority/evidence atau tidak decision-ready, catat exact gap/next owner, jangan minta approval desain yang belum konkret. Valid prior decisions tidak dibuka ulang kecuali finding/owner decision menuntut bounded change; bedakan technical correction terhadap chosen design dari perubahan owner semantics. Any durable authority/spec/API promotion dirutekan ke owner/gate yang sesuai, bukan ditulis oleh backend Planner.

Declare apakah resolution mengubah material scope, architecture/ownership, business/security/interface/data semantics atau verification strategy. Jika material, independent re-review mengikuti canonical rule sebelum report, kecuali explicit Human waiver yang benar-benar ada; tidak ada waiver sekarang. Jangan generate report selama resolution/re-review/owner churn. Self-check rule coverage/Test Focus/Open Items lifecycle dan seluruh existing controls/gates tetap dipertahankan.

No production/migration/spec/API/monetary-standard writes, tests/validators/generation/runtime/security checks. Stop setelah successor Draft dan handoff; jangan self-approve, claim findings verified closed, mulai Build atau menerima residual risk. Rekomendasikan applicable next gate dari actual resolution/owner status.

## Execution envelope

- `PREAUTHORIZED`: read finding/plan/evidence/current authority/live source; write hanya successor planning/provenance/handoff pada RUN_PATH; facilitate decision-ready owner question dan record explicit answer.
- `ORCHESTRATOR_DECISION`: reconcile resolution/owner evidence lalu route applicable independent re-review/report/Human approval serta owning-source reconciliation.
- `HUMAN_REQUIRED`: material monetary/product/API decisions, whole-Techplan approval, owning-authority changes, exact Tier-0 source authorization, residual-risk acceptance dan any review waiver.

Protected `campaign/donation_coordinator_db.go`, `donation/ledger.go`, any balance transaction/locking file, crypto/auth fencing tetap berlaku. Draft mechanism/impact proposal diperbolehkan; protected implementation tidak diotorisasi. O3/O4/O5 controls/provider/key/abuse/evidence dan scoped Campaign GET producer dependency tetap sesuai current plan; jangan menghilangkannya dalam resolution.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Planner / KC-PLANNER Session; configured `gpt-6-luna` / `high`. Start canonical Techplan synthesis entrypoint dengan Invocation/overlay ini.

Kickoff: `Jalankan Planner resolution Run TP-S2-003-003 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-003/invocation.md dan canonical ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md dengan orchestrated-run overlay. Tulis successor Draft Techplan dan phase handoff pada RUN_PATH, lalu berhenti sesuai gate fase.`

Laporkan completion, exact owner question yang pending, atau blocker kepada Orchestrator. Tidak ada Participant auto-dispatch atau solusi monetary yang telah dipilih oleh Invocation.
