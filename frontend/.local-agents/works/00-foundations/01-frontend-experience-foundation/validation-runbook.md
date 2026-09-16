# Validation 02 — Frontend Experience Foundation Operator Runbook

> Status: Active operator guide
> Run ID: `validation-02`
> Operator: Anhar Solehudin
> Prepared with: ChatGPT — GPT-5.6 Sol
> Created: 2026-09-16
> Task root: `frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation`
> Validation branch: `validation-02-frontend-experience-foundation-cleanstart`
> Product/project baseline: `main@0273c4f2b7f3140efe53a3736547fe73fdd2aefe`
> Harscode baseline: `workflow-v2@4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Purpose

This is an **operator runbook for Validation 02**, not project authority and not a Harscode phase artifact.

It tells the human operator how to run the validation consistently from Exploration through delivery while preserving natural Harscode behavior.

Phase agents do **not** need to read this file. The agent-facing execution source is always the current canonical Harscode phase prompt plus applicable Kencleng authority/artifacts.

`benchmarks.md` is the paired operator telemetry file.

## Core execution rule

Do not use a ChatGPT-authored wrapper or mega-prompt around Harscode.

For every phase:

```text
open the canonical Harscode phase prompt
→ fill the inputs it explicitly requires
→ paste/run that prompt manually in Codex
→ follow its gates and phase handoff
```

The canonical prompt owns the lifecycle mechanics. Kencleng authorities own project/product truth. The execution profile owns model/client selection.

Do not add extra coaching about the expected solution merely because this is a validation run.

## Benchmark rule

Benchmark **actual Codex threads**, not phase names.

For every new thread:

```text
/status
→ run work normally
→ /status before finally leaving that thread
```

Copy the observed snapshots into `benchmarks.md`.

If a healthy thread spans multiple phases, it remains one benchmark session.

If the workflow sends a patch back to an existing Build thread, use `/resume` to navigate back when appropriate. `/resume` is not a measurement command; `/status` is the measurement snapshot.

An optional `/status` checkpoint may be captured before a resumed patch, but do not create measurement ceremony that changes natural workflow behavior.

## Fixed task inputs

These values remain stable unless the run itself exposes a real reason to change them:

```text
TASK = kencleng/docs/spec/0-foundations/features/01-frontend-experience-foundation.md

TASK_PATH = kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation

CODEBASE_CONTEXT = Kencleng — Go backend + Next.js frontend

Ticket = none

Area = frontend — public experience foundation

Kencleng target baseline = 0273c4f2b7f3140efe53a3736547fe73fdd2aefe

Harscode workflow baseline = 4199c6db1b26ef1920ba670f222aff0c6d0f9e59
```

Set `HARSCODE_WORKSPACE_ROOT` to the real local path from the environment/session in which Codex is running. Do not copy an incorrect relative path merely to match this runbook.

## 0. Pre-flight

Before starting the first Codex thread:

```bash
git fetch origin
git checkout validation-02-frontend-experience-foundation-cleanstart
git pull --ff-only
git status --short
git rev-parse HEAD
```

The working tree should be clean.

The branch HEAD should include the current operator instrumentation commit(s), while its product baseline ancestry remains:

```text
main@0273c4f2b7f3140efe53a3736547fe73fdd2aefe
```

Verify the Harscode workspace separately:

```bash
git checkout workflow-v2
git pull --ff-only
git rev-parse HEAD
```

Expected workflow revision:

```text
4199c6db1b26ef1920ba670f222aff0c6d0f9e59
```

Do not start the phase if the wrong project branch/workflow revision is checked out.

---

## 1. Exploration

### Session

Start a **fresh Codex thread**.

Default execution profile:

```text
Model: GPT-5.6 Terra
Reasoning: Medium
```

Run:

```text
/status
```

Record the raw snapshot in `benchmarks.md`.

### Canonical prompt

Open manually:

```text
harscode-workspace/workflow/1-exploration-kickoff-prompt.md
```

Fill its required inputs using the fixed task inputs above, then paste/run the prompt directly.

Do not prepend another kickoff prompt.

### Stage 1 gate

The agent should stop after **Stage 1 — Plan Announcement**.

Check only whether:

- task understanding is materially correct;
- exploration areas/boundaries are reasonable;
- the agent has not started deep implementation inspection or solutioning.

If correct, respond naturally:

```text
lanjut Stage 2
```

If incorrect, correct the **fact/boundary only**. Do not give the solution unless the workflow genuinely needs a human product decision.

### Stage 2 gate

Let the agent complete gap analysis across the announced areas.

Check whether the gap analysis is grounded and whether obvious material areas were missed.

If sound:

```text
lanjut Stage 3
```

A factual correction is not automatically a rescue prompt. Record it accurately in the benchmark notes.

### Stage 3 completion

Let the agent complete solutioning and write durable Exploration evidence.

At the end, read its `## Phase handoff`, especially:

```text
Human decision
Open / deferred
Recommended next step
Session transition
Context pointers
```

Do **not** decide fresh vs continue by habit. Follow the handoff unless there is an observable reason it is wrong.

### Exploration → Techplan transition

If the handoff says the thread remains healthy for continuation:

```text
continue in the same Codex thread
```

Do not close the benchmark session yet.

If the handoff recommends fresh Techplan context:

```text
/status
→ record Exploration thread final snapshot
→ leave thread
→ start fresh Techplan thread
```

---

## 2. Techplan synthesis

### Session/model

Default:

```text
GPT-5.6 Terra · Medium
```

Use the Exploration handoff to choose continued vs fresh.

If fresh, run `/status` before the Techplan prompt and create a new session row in `benchmarks.md`.

### Canonical prompt

Open manually:

```text
harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md
```

Fill only the inputs the canonical prompt requires. Do not paste a second summary of Exploration unless the prompt explicitly needs information that is not durably available.

Expected primary artifact:

```text
TASK_PATH/2-techplan/techplan.md
```

### Human Techplan gate

Synthesis must stop for human review before Build.

Review:

- material product/domain semantics;
- design/architecture ownership decisions;
- interface/data/authority decisions;
- meaningful risk and verification strategy;
- Active Open Items;
- whether Build would still need to invent a material decision.

Also record the agent's recommendations:

```text
Independent Techplan review: Skip | Recommend | Required by project policy
Decomposition: Skip | Consider
```

Do not invoke either optional mechanism merely because it exists.

---

## 3. Optional Independent Techplan Review

Run only when synthesis recommends/require it or the human has a concrete reason to request it.

### Session

Use a **fresh independent thread**.

Default:

```text
GPT-5.6 Sol · Medium
```

Run `/status`, create a benchmark session row, then open:

```text
harscode-workspace/workflow/2-2-techplan-review-prompt.md
```

Fill its required inputs and run it directly.

The prompt owns its own Complex gate. If the gate says review is not warranted, accept that outcome unless project policy requires otherwise.

If findings are blocking, return them to Techplan authority for one resolution pass. Do not let the independent reviewer silently rewrite the plan.

When the independent review thread is finished, run `/status` and record it.

---

## 4. Techplan resolution and Approval

Resolve material review findings/Open Items through Techplan authority.

Reuse the synthesis thread if it remains healthy and current; otherwise start a fresh Techplan resolution thread grounded on durable artifacts.

Do not approve the plan while a material Active Open Item still requires a human/project decision before Build.

When satisfied, explicitly approve the Techplan.

After Approval, produce the human-facing Techplan report using the current Harscode Techplan reporting guidance/template as required by the workflow. The Approved `techplan.md` remains the execution contract.

---

## 5. Optional decomposition

Only after the parent Techplan is **Approved**.

Run this only if synthesis said `Consider` and there is a real execution/review/context boundary worth splitting.

Canonical prompt:

```text
harscode-workspace/workflow/2-3-techplan-decomposition-prompt.md
```

The prompt owns its own gate. If the answer is NO, stop; do not manufacture task files.

If task files are generated, review/accept the split before Build.

Decomposition is not a new source of product truth. The Approved parent Techplan remains the spine.

---

## 6. Build

### Session

Build is **fresh preferred** after Approved Techplan.

For this task, treat the initial implementation as material product/UI engineering unless the actual Approved plan proves otherwise.

Default execution profile:

```text
GPT-5.6 Sol · Medium
```

Use the client/environment that preserves implementation quality. Do not degrade visual/product work merely to keep all execution inside one CLI environment.

For a fresh Build thread:

```text
/status
```

Record a new benchmark session.

### Canonical prompt

Open manually:

```text
harscode-workspace/workflow/3-build-prompt.md
```

Fill the required inputs and run it directly.

Build should execute:

```text
Approved Techplan
+ current task slice when decomposition exists
+ current live code
```

It should not broadly reload raw Exploration history.

Expected initial report:

```text
TASK_PATH/3-build/report.md
```

Let Build perform its normal edit → focused verify → fix loop.

When Build is complete, read the handoff. Before leaving the Build thread for independent Code Review, run `/status` and record the snapshot.

Keep the Build thread resumable because Review or Testing may send a narrow patch back to it.

---

## 7. Code Review

### Session

Start a **fresh independent thread**.

Default:

```text
GPT-5.6 Terra · Medium
```

Run `/status` and create a new benchmark row.

### Canonical prompt

Open manually:

```text
harscode-workspace/workflow/4-code-review-prompt.md
```

Fill its required inputs, especially the **current diff / changed-file scope**, and run it directly.

Review must not edit production code.

### Review outcome

If verdict is:

```text
Approve
or
Approve with minor comments
```

follow its handoff toward Testing.

If verdict is:

```text
Request changes
```

use the generated specific Review patch plan and return production fixes to **Build/Patch authority**.

Before leaving the Review thread, run `/status` and record it.

---

## 8. Review → Build/Patch loop

If Code Review requests code changes:

1. follow the Review handoff's session-transition recommendation;
2. when the existing Build thread remains healthy, use `/resume` to return to it;
3. otherwise start a fresh Build/Patch thread grounded only on the Approved Techplan + specific patch plan + relevant live code;
4. run `/status` on resume/fresh entry if useful for benchmark checkpointing;
5. run the canonical `3-build-prompt.md` in patch mode using the specific patch plan;
6. write a new `patch-report-<n>.md`; never overwrite earlier Build evidence;
7. return to the phase that requested the patch.

Do not automatically run a full new Code Review merely because any patch happened. Follow Harscode's proportional re-review rule and the requesting phase's handoff.

---

## 9. Testing

### Session

Start a **fresh independent thread**.

Default:

```text
GPT-5.6 Terra · Medium
```

Run `/status` and create a benchmark row.

### Canonical prompt

Open manually:

```text
harscode-workspace/workflow/5-testing-prompt.md
```

Fill the required inputs and run it directly.

Testing should follow **SWEEP, DON'T REDO**:

- verify Build claims rather than rewriting equivalent tests;
- close deferred/flagged gaps first;
- execute Testing-owned verification;
- run required final repo verification;
- exercise meaningful observable interfaces;
- respect Human-owned checks instead of marking them passed automatically.

Expected report:

```text
TASK_PATH/5-testing/testing-report-<n>.md
```

### Testing outcome

If verdict is `Pass`, continue toward human acceptance / PR.

If verdict is `Fail — send back to Build`, use the specific Testing patch plan and return fixes to Build/Patch authority.

Before leaving the Testing thread, run `/status` and record it.

---

## 10. Testing → Build/Patch loop

For a Testing-discovered defect:

```text
Testing patch plan
→ Build/Patch authority
→ affected verification first
→ return to Testing
```

Resume the prior Build thread when the handoff says it remains healthy; otherwise start a fresh Build/Patch thread.

Testing re-entry should verify the affected gap first. Do not replay unrelated expensive checks only because a new round began; rerun broad/final verification when the approved contract/risk requires it.

---

## 11. Human rendered acceptance

Material frontend UI requires human acceptance before delivery.

This is separate from agent Build/Testing evidence.

Exercise the rendered implementation in a real browser at representative scope, including at minimum the desktop/mobile surfaces and important states justified by the implemented slice.

Focus on what automation/reasoning cannot fully accept for you:

- visual hierarchy and clarity;
- interaction comprehension;
- responsive usability;
- consistency with approved product/design intent;
- asset appropriateness;
- obvious state/error behavior in context.

Record the result in `benchmarks.md` Quality Evidence and in the appropriate workflow/delivery artifact when required.

If acceptance finds a real implementation defect, send it through Build/Patch authority and re-verify proportionally.

---

## 12. Pull Request

Harscode currently has **no root `6-pull-request-prompt.md`**.

Use the current PR guidance directly:

```text
harscode-workspace/workflow/6-pull-request/guidelines.md
harscode-workspace/workflow/6-pull-request/template.md
```

Use `examples.md` only when a concrete calibration need exists.

PR truth must come from:

```text
final repository state
+ final durable workflow evidence
```

not from planned changes that never landed.

PR/report/admin execution can use the smallest sufficient model/client according to the Kencleng execution profile.

Whether PR work gets a fresh thread is flexible. If it is a new thread and part of benchmark measurement, take `/status` before/after and add the row.

---

## 13. Validation closeout

After final implementation/testing/human acceptance/PR state is known, finish `benchmarks.md`.

For each actual Codex thread record:

- actual phase(s);
- model/reasoning;
- fresh/continued/resumed posture;
- raw `/status` before/after snapshots;
- normal human prompts;
- factual corrections/redirections;
- clarifications;
- rescue prompts;
- material outcome/notes.

Then complete Quality Evidence:

- requirement/authority misses;
- redirections caused by misunderstanding;
- optional Techplan review recommendation quality;
- decomposition recommendation quality;
- verification ownership/rationale quality;
- Code Review findings;
- Testing findings;
- avoidable repeated work;
- final human rendered acceptance;
- reusable workflow/project learning.

Evaluate the run directionally against Validation 01.

Do not conclude that lower token/context usage is an improvement if it came with worse correctness, more rescue prompts, weaker verification, or worse product outcome.

## Operator decision rule

At every unexpected branch in the workflow, prefer this order:

```text
1. follow the current canonical phase prompt
2. follow its explicit phase handoff
3. follow current Kencleng project authority
4. use this runbook only to keep operator/benchmark procedure consistent
```

If this runbook ever conflicts with current Harscode or Kencleng authority, the current authority wins and the runbook should be updated afterward.