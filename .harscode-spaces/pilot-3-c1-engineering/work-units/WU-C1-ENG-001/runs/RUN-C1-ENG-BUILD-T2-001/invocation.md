# Run Invocation — RUN-C1-ENG-BUILD-T2-001

## Identity and assignment

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-BUILD-T2-001`.
- **Phase route / Role:** Build — T2 backend identity/session boundary (G1) / Implementer.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-IMPLEMENTER-001` / `SESSION-C1-ENG-T2-IMPLEMENTER-001` (FRESH).
- **Session reason:** New task Run after T1 contract Build, Review, bounded patch, and patch Review completed; reconstruct from durable invocation/evidence.
- **Dispatch posture:** Human-Assisted; Anhar approved `gpt-6-sol` / `medium` for this Run on 2026-10-10 (“approve bro”) and returned the terminal Build report. Exact launch timestamp/runtime model are not independently inferred.
- **Communication:** Bahasa Indonesia; retain canonical technical identifiers/terms.
- **Escalation owner:** Anhar via Orchestrator.

## Current-effective authority and inputs

Paths are project-root-relative unless otherwise stated.

- Approved Techplan: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.md`, SHA-256 `bd732748b486f0327f6cba2cdf14cab32f1ece3c5be8ebb24f9659444947fa47`.
- Current task: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/T2-backend-identity-session.md`, SHA-256 `075170761afb971cd09aba90b1e112c617f12d27d4b8bb4ffbd1a3599163378b`.
- Accepted task manifest: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/tasks/manifest.md`, SHA-256 `2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d`.
- T1 dependency evidence: approved [patched T1 Code Review findings](.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T1-PATCH-001/evidence/review-findings-1.md), SHA-256 `f8a61a3bfd591a5310e291e84ec0aac393b615c731216238abbacee0342782ea`; [T1 Build/Patch report](.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T1-PATCH-001/evidence/patch-report-1.md), SHA-256 `92d8d1f8cfa86eb4b4d1c197bbaa493b09531d504b0a43846a62937c0016e3f7`.
- Exact active contract source: `api/openapi/index.yaml` SHA-256 `ddefffbe588459cb0cba9f68dfedaccba3da8a1313f1956511cf628df7199965`; bundle `api/openapi.yaml` SHA-256 `38bb6c531804b54e9d5ce56d623c397a590722c56dcb7a3c12e646bb8280fa09`; generated types `api/openapi.d.ts` SHA-256 `d21e8aa1eec92a95cc1fed87d976f726f59b48bd334aa8e463844d48e16cb80c`.
- **T1 hard dependency satisfied:** accepted manifest requires coordinated auth/session contract shapes and durable source/bundle revision with validation/bundle result. The exact contract hashes above are review-approved; the bound T1 patch report records successful validate/bundle/types and reproducible outputs. This establishes the task edge only, not runtime correctness.
- **TARGET_REVISION:** `59f2b25940fa4baf243b4f1d2ca2ba04c9b17fd4`, branch `pilot/3-c1-engineering`; inspect current working tree and preserve unrelated changes. At preparation no backend implementation paths are changed.
- **WORKFLOW_REVISION:** `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`; Harscode tree observed clean.
- **WORK_UNIT_PATH / TASK_PATH compatibility root:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-001`.
- **ARTIFACT_TARGET:** Backend T2 implementation only, following live anchors and scope in T2 task/Techplan; Run-owned evidence under `RUN_PATH/evidence/`.
- **CWD / Harness:** `/home/anhar-solehudin/kencleng-workspace/kencleng` / `codex-cli`.

## Task and permission boundary

Implement exactly T2: Google OIDC identity mapping and local opaque session/authentication boundary on the approved T1 contract; provide T3 the authenticated person/session interface and same-person/session-row-lock revalidation behavior required by the task.

G1 implementation authorization is already recorded in WU `events.md` and parent Resolved item 9. It explicitly covers Google OIDC/PKCE/state/nonce/token verification, local person mapping, opaque session issuance/digest/expiry/rotation/revocation, cookies including the explicit localhost exception, CSRF, and secret handling. Do not ask for that unchanged approval again or exceed it. Do not introduce a new privileged G2/G3 surface; if schema/auth work crosses those unapproved-for-this-task boundaries materially, stop for Orchestrator routing.

No custom cryptography, local JWT/signing key system, refresh token, password/MFA/recovery, email/domain/body-person authority, broad role engine, Organization aggregate/guard, frontend writes, or Product/Techplan/task/manifest edits. Do not copy historical security mechanisms merely because they exist. Follow root and `backend/AGENTS.md`, approved parent, T2 task, current backend architecture and active API contract. Parameterize SQL, preserve error chains, and never log secrets, raw tokens, or unnecessary PII.

Before edits, re-ground live `backend/cmd/server/main.go`, `backend/internal/platform/db/db.go`, `backend/go.mod`, migrations, API contract, current dependencies and their source. If a material contract assumption or protected boundary differs, stop and surface the smallest blocker instead of redesigning.

## Focused Build verification and evidence

Follow current `../harscode-workspace/workflow/3-build-prompt.md`, `3-build/guidelines.md`, `3-build/checklist.md`, and `orchestrated-run-overlay.md`. Author focused protocol-negative, person-upsert, and session-lock/revalidation checks required by T2. Record dependency source/graph review, `go mod verify`, vulnerability scan and available compile/auth checks per parent §10/§12 and actual repo commands. Distinguish fakes/compile checks from real Google login, configured PostgreSQL concurrency, provider/runtime, and independent Testing evidence. Respect actual dependency/network permissions; do not substitute a different dependency or version without routing a material gap.

Write `RUN_PATH/evidence/build-report.md` in canonical Build format with exact commands/results, dependency and changed-file identities, limitations, and exactly one structured `## Phase handoff`. State explicitly that no race/performance/security-class final sweep or broad Testing-owned suite ran unless such a scope deviation is actually necessary and recorded. Return T2 evidence to Orchestrator for T3 dependency reconciliation; do not dispatch another task.

## Execution envelope and model route

- **PREAUTHORIZED:** Read bound/current inputs; write T2 backend auth/config/person/session/migrations and focused tests within T2/G1; use existing approved contract and focused repo tooling; write Run-owned evidence.
- **ORCHESTRATOR_DECISION:** Stop on material authority/contract/architecture drift; after completion assess durable T3 prerequisites and route the next task.
- **HUMAN_REQUIRED:** Run-specific model approval before dispatch; any new material scope/risk/security authority, additional protected G2/G3 surface, destructive/external action, or deployment.
- **Selected model / effort:** `gpt-6-sol` / `medium`.
- **Routing rationale:** G1 implements cross-cutting authentication/session security including OIDC validation, opaque session lifecycle, CSRF, and T3's locked session revalidation interface. Human-declared Sol capabilities include strong reasoning, architecture, and cross-cutting analysis; medium is the lowest supported effort for this bounded but security-sensitive implementation. Luna's registry does not declare those stronger cross-cutting capabilities. This is routing for settled authority, not permission or runtime proof.
- **MODEL_APPROVAL:** APPROVED by Anhar on 2026-10-10 (“approve bro”) for `RUN-C1-ENG-BUILD-T2-001` only. Earlier Run approvals were not inherited.
- **Dispatch / outcome:** Launched according to Anhar; terminal outcome **BLOCKED** before production writes.
- **Result ref:** `RUN_PATH/evidence/build-report.md`, SHA-256 `130b888011a5c5286b818078410d2b022faf60d0033f20d0e89d628098c0ad50`; scanner/dependency evidence is indexed in that report.
- **Next route:** Material dependency/toolchain reconciliation under Approved Techplan §10 is required before a fresh T2 Build/Patch Run. T3's T2 session/person prerequisite is not satisfied; do not resume or redispatch this blocked Run.
