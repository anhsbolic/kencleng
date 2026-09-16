# Validation 02 — Frontend Experience Foundation Benchmark

> Status: Active measurement
> Run ID: `validation-02`
> Measurement owner: Anhar Solehudin
> Prepared with: ChatGPT — GPT-5.6 Sol
> Created: 2026-09-16
> Updated: 2026-09-16
> Task authority: `docs/spec/0-foundations/features/01-frontend-experience-foundation.md`
> Task root: `frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation`
> Validation branch: `validation-02-frontend-experience-foundation-cleanstart`
> Kencleng target baseline: `main@0273c4f2b7f3140efe53a3736547fe73fdd2aefe`
> Harscode workflow baseline: `workflow-v2@4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Purpose

This file is **operator telemetry**, not task authority and not a Harscode phase artifact.

Phase agents should not use benchmark data to decide implementation, verification, or workflow behavior. The benchmark observes the workflow; it must not change how the workflow is executed merely to produce nicer numbers.

Validation 02 tests the refined workflow on a different real task shape: the first frontend experience foundation built from Kencleng's clean frontend scaffold and current Sunlit Editorial design authority.

## What this validation is trying to learn

Compare Validation 02 with Validation 01 directionally, not token-for-token. The task shape and implementation complexity differ.

Evaluate both:

```text
efficiency
→ session/context usage
→ repeated authority loading
→ avoidable verification repetition
→ human prompts / rescue prompts

correctness
→ authority adherence
→ outcome quality
→ human redirection caused by misunderstanding
→ review findings
→ testing findings
→ final human acceptance
```

The optimization target is **minimum sufficient context with same-or-better correctness/outcome quality**, not minimum token usage in isolation.

## Planned execution profile

| Work | Default model | Reasoning | Session posture |
|---|---|---|---|
| Exploration | GPT-5.6 Terra | Medium | Fresh |
| Techplan synthesis | GPT-5.6 Terra | Medium | Follow Harscode handoff recommendation |
| Independent Techplan review | GPT-5.6 Sol | Medium | Only when recommended/required and its gate applies |
| Material product/UI Build | GPT-5.6 Sol | Medium | Fresh preferred after Approved Techplan |
| Code Review | GPT-5.6 Terra | Medium | Fresh |
| Testing | GPT-5.6 Terra | Medium | Fresh |
| PR/report/admin | GPT-5.6 Luna or Terra | Low | Flexible |

Escalation remains evidence-based. Record the actual model/reasoning used when it differs.

## Measurement protocol

### Benchmark unit

One **Codex thread/session** is one benchmark unit, even when the same healthy thread spans more than one workflow phase.

Examples:

- Exploration → Techplan in one continued thread = one benchmark session;
- Exploration fresh thread + Techplan fresh thread = two benchmark sessions;
- Build resumed later for a narrow Review/Testing patch remains the same Build thread; note the resume/patch activity in that session rather than pretending it is an unrelated new thread.

### Required measurement

For every new Codex thread:

```text
open fresh/resumed thread
→ /status before the first work prompt in that active session
→ perform workflow work normally
→ /status before finally leaving that thread / when its useful work for this validation is complete
```

Copy the `/status` output as observed. Do not reconstruct unavailable metrics from logs and do not invent values.

For a resumed thread, an extra `/status` checkpoint before the resumed patch is useful but optional. The required benchmark remains the thread's observed starting and final state.

### What `/status` is for

Use `/status` as the canonical operator snapshot. Record whatever the installed Codex client exposes, for example model/reasoning, thread/session identity, cumulative usage/context information, or rate-limit state.

Do **not** require a field merely because this document names it. Codex output may change over time.

`/resume` is navigation, not measurement. Use it only when the workflow handoff calls for returning to an existing thread.

### Human interaction counters

Record:

- normal canonical prompt invocation;
- normal human gates (`lanjut Stage 2`, `lanjut Stage 3`, Techplan approve/revise, etc.);
- factual/boundary corrections;
- clarifications requested by the agent;
- rescue prompts.

A **rescue prompt** is human coaching added mainly to make the agent succeed because canonical workflow/project guidance was insufficient or missed.

Normal human gates, requested decisions, and factual corrections are **not automatically rescue prompts**.

## Session summary

Add a row when a new Codex thread starts. Do not add a new row merely because a healthy thread crosses a phase boundary.

| Session | Phase(s) | Model / reasoning | Posture | `/status` before | `/status` after | Human prompts | Rescue prompts | Clarifications | Outcome / notes |
|---|---|---|---|---|---|---:|---:|---:|---|
| 1 | Exploration | GPT-5.6 Terra · Medium | Fresh | pending | pending | pending | 0 | pending | Not started |

## Raw `/status` snapshots

Preserve raw snapshots here so later analysis does not depend on remembered values.

### Session 1 — Exploration

**Before**

```text
pending
```

**After**

```text
pending
```

**Intermediate checkpoints (optional)**

```text
none
```

**Operator notes**

```text
Human prompts: pending
Rescue prompts: 0
Clarifications: pending
Outcome: pending
```

Add another session section only when a genuinely new Codex thread starts.

## Quality evidence

Track lifecycle quality separately from usage snapshots.

| Signal | Evidence |
|---|---|
| Requirement/authority misses | — |
| Human redirections caused by agent misunderstanding | — |
| Rescue prompts | — |
| Techplan review recommendation quality | — |
| Decomposition recommendation quality | — |
| Verification rationale/ownership clarity | — |
| Code Review findings | — |
| Testing findings | — |
| Verification repetition / avoidable reruns | — |
| Final human rendered acceptance | — |
| Reusable workflow/project learning | — |

## Benchmark hygiene

Do not alter workflow behavior merely to improve benchmark numbers.

In particular, do not:

- keep a stale thread alive only to avoid starting a new session;
- force a fresh thread merely to lower context usage;
- choose a weaker model only for lower measured usage;
- skip justified verification;
- feed benchmark observations into phase-agent reasoning;
- treat lower usage as success when correctness or outcome quality degrades.

When a session transition, model escalation, browser check, or extra verification is justified by the actual task/workflow, execute it normally and record the reason.

## Baseline note

The target code/document baseline for this run is exactly:

```text
Kencleng main
0273c4f2b7f3140efe53a3736547fe73fdd2aefe

Harscode workflow-v2
4199c6db1b26ef1920ba670f222aff0c6d0f9e59
```

The benchmark/runbook commits on the validation branch are operator instrumentation only; they do not change the product/task baseline.