# Run Invocation — RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-PATCH-001

## Identity and posture

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-PATCH-001`.
- **Phase / route:** Independent Code Review — exact configured-pool harness delta.
- **Role / specialization:** Reviewer / backend auth harness.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-CONFIGURED-DB-HARNESS-PATCH-REVIEWER-001` / `SESSION-C1-ENG-T2-CONFIGURED-DB-HARNESS-PATCH-REVIEWER-001` (fresh identities; not launched).
- **Dispatch:** Human-Assisted; `READY_FOR_HUMAN_DISPATCH`. Human mechanical dispatch is pending; no Review Run is active.
- **Selected model / effort:** `gpt-6-luna` / `medium`; current Human-owned runtime registry supports the required repository-review/reasoning work and marks it `approval_required: false`.
- **Model routing rationale:** Low-cost Luna at medium is sufficient for a bounded four-pass static review of a three-file test-harness/documentation delta. Escalate only if the assignment reveals material cross-cutting security or authority reasoning beyond the bound scope.
- **Target revision / branch:** Build target revision `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering`; exact review target is the three post-delta file hashes bound in `input-bindings.json`.
- **Workflow revision:** `651ce57f044bff23fba99082322cee6c066d25a6`.
- **Harscode root:** `../harscode-workspace`.
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-PATCH-001`.
- **Communication language:** Bahasa Indonesia.
- **Scope:** `scope.md`; exact source/report/authority identities are recorded in `input-bindings.json`. Mutable lifecycle projections are not assignment bindings.

## Claim and independent value

The terminal Build/Patch report claims the restricted test harness now passes child fixture `DATABASE_URL` through production `db.Open`, while retaining six `t2` assertions and ownership-checked cleanup. Independent Review is applicable because credential/environment boundaries, ownership validation/deletion, and the configured-pool oracle changed. Review examines the exact three-file delta; it is not runtime acceptance and must not claim Testing-owned R14 or B-T2-003 evidence.

## Required guidance and envelope

Read the current canonical Code Review prompt, overlay, guidelines, and checklist; the bound Approved Techplan/T2 task; root/backend `AGENTS.md`; applicable auth verification environment guidance; current diff; and only the matching Go/Python/process/resource/security best practices routed by the index.

- **PREAUTHORIZED:** read bound source/evidence and perform four-pass static reasoning; write this Run's Review report and a patch plan only if required.
- **ORCHESTRATOR_DECISION:** none.
- **HUMAN_REQUIRED:** mechanical dispatch; any fixture/database/provider/browser/runtime operation; source edits; authority, dependency, scanner-policy, acceptance, or commit changes.

Do not run tests, gates, scanners, fixtures, DB operations, provider flows, or browser/runtime checks in this Review. No production edits. If a question requires runtime evidence, state it as unresolved for Testing.

## Human-Assisted kickoff

```text
Jalankan RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-PATCH-001 sesuai Invocation dan scope yang terikat. Lakukan empat pass Review secara statis pada delta tepat tiga file. Periksa batas environment/credential fixture, `DATABASE_URL` child menuju `db.Open`, ownership dan cleanup, enam kasus t2 termasuk expiry-under-lock, serta batas klaim README. Jangan menjalankan tes, gate, scanner, fixture/DB, Provider, browser/runtime, atau mengubah source. Tulis review-findings-1.md dengan tepat satu Phase handoff.
```

## Terminal closure

This Review occurrence is prepared only. Human mechanical dispatch is pending. A terminal Review handoff closes this occurrence; any follow-up uses a new Run/Participant/Session. Testing preparation and any fixture authorization remain a later independent route after Review reconciliation.
