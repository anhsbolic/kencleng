# Run Invocation — RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001

## Identity and posture

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001`.
- **Phase / route:** Independent targeted Code Review confirmation — Review F1 credential-environment isolation patch.
- **Role / specialization:** Reviewer / backend auth harness.
- **Participant / Session:** `PARTICIPANT-C1-ENG-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-REVIEWER-001` / `SESSION-C1-ENG-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-REVIEWER-001` (fresh identities; not launched).
- **Dispatch:** Human-Assisted; `READY_FOR_HUMAN_DISPATCH`. Human mechanical dispatch is pending; no Review Run is active.
- **Selected model / effort:** `gpt-6-luna` / `medium`; current Human-owned registry supports the bounded repository review and marks it `approval_required: false`.
- **Model routing rationale:** A narrow static confirmation of environment-copy behavior and a mocked regression is within Luna's repository-work/reasoning capabilities. Escalate only if live evidence expands the scope materially.
- **Target revision / branch:** `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `pilot/3-c1-engineering`; exact current target content is bound in `input-bindings.json`.
- **Workflow revision:** `651ce57f044bff23fba99082322cee6c066d25a6`.
- **Harscode root:** `../harscode-workspace`.
- **Work Unit path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001`.
- **Run path:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001`.
- **Communication language:** Bahasa Indonesia.
- **Scope:** `scope.md`; exact target and relevant authority are bound in `input-bindings.json`. Mutable WU/Events/Control Tower lifecycle projections are not assignment bindings.

## Claim and verification boundary

The prior Review returned Request changes for F1 because fixture DB credentials reached `go version`. The patch introduces a copied metadata environment with fixture keys removed and a mocked regression; the Go test child keeps the original configured environment. Independently assess whether this closes F1 on the exact two-file delta. This Review is static and does not establish runtime or Testing claims.

## Required guidance and envelope

Read the canonical current Code Review prompt/overlay/guidelines/checklist; prior F1 Review report and patch plan; terminal Build/Patch report; Approved Techplan/T2 task; root/backend `AGENTS.md`; and relevant routed best-practice guidance.

- **PREAUTHORIZED after Human mechanical dispatch:** inspect exact bound diff and evidence; perform all four Review passes; write this Run's Review report.
- **ORCHESTRATOR_DECISION:** reconcile the F1 verdict and recheck Testing applicability/target.
- **HUMAN_REQUIRED:** mechanical dispatch; any runtime/fixture/DB/provider action, source edit, scope expansion, authority or dependency change, acceptance/commit, or residual-risk decision.

Do not run tests, gates, scanners, fixtures, database operations, Provider flows, browsers, or runtime checks. Do not edit source.

## Human-Assisted kickoff

```text
Jalankan RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-CREDENTIAL-PATCH-001 sesuai Invocation dan Scope. Review statis tepat dua file pada patch F1 dengan empat pass berurutan. Pastikan environment untuk `go version` tidak berisi fixture credential/DSN atau Run identity, dan environment Go test child tetap utuh. Jangan jalankan tes, gate, scanner, fixture/DB, atau runtime. Tulis review-findings-1.md dengan tepat satu Phase handoff.
```

## Terminal closure

This Review occurrence is prepared and not launched. Human mechanical dispatch is pending. Its verdict applies only to F1 and this exact patch. Testing remains held until this Review is terminal and the Testing target/bindings are rechecked.
