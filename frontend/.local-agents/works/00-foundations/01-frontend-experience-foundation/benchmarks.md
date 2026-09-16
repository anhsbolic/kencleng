# Validation 02 — Frontend Experience Foundation Benchmark

> Status: Active measurement
> Run ID: `validation-02`
> Measurement owner: Anhar Solehudin
> Prepared with: ChatGPT — GPT-5.6 Sol
> Created: 2026-09-16
> Task authority: `docs/spec/0-foundations/features/01-frontend-experience-foundation.md`
> Task root: `frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation`
> Validation branch: `validation-02-frontend-experience-foundation-cleanstart`
> Kencleng target baseline: `main@0273c4f2b7f3140efe53a3736547fe73fdd2aefe`
> Harscode workflow baseline: `workflow-v2@4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Purpose

This file is **operator telemetry**, not task authority and not a Harscode phase artifact.

Phase agents should not use benchmark data to decide implementation, verification, or workflow behavior. Capture measurements outside the task reasoning loop so optimization pressure does not distort correctness.

Validation 02 tests the refined workflow on a different real task shape: the first frontend experience foundation built from Kencleng's clean frontend scaffold and current Sunlit Editorial design authority.

## Comparison posture

Compare Validation 02 with Validation 01 directionally, not token-for-token. Task complexity and implementation shape differ.

Evaluate both:

```text
efficiency
→ context usage / turns / repeated reading / verification repetition / rescue prompts

correctness
→ authority adherence / outcome quality / human redirection / review findings / testing findings
```

Do not add cached and non-cached input as if they were equivalent cost. Cached input is still useful as evidence of context replay/reuse.

## Planned execution profile

| Work | Default model | Reasoning | Session posture |
|---|---|---|---|
| Exploration | GPT-5.6 Terra | Medium | Fresh |
| Techplan synthesis | GPT-5.6 Terra | Medium | Follow Harscode handoff recommendation |
| Independent Techplan review | GPT-5.6 Sol | Medium | Only when recommended/required |
| Material product/UI Build | GPT-5.6 Sol | Medium | Fresh preferred after approved Techplan |
| Code Review | GPT-5.6 Terra | Medium | Fresh |
| Testing | GPT-5.6 Terra | Medium | Fresh |

Escalation remains evidence-based; record actual model/reasoning used when it differs.

## Session measurement rule

One **actual agent session** is one benchmark unit, even when one session spans more than one workflow phase.

For every session:

1. capture available usage/context/quota telemetry **before** sending the first task prompt;
2. capture the same telemetry after the session's useful work is complete;
3. record actual phase(s), model, reasoning, fresh/continued posture, human prompts, clarifications, rescue prompts, and outcome;
4. note repeated authority loading or verification repetition when materially visible;
5. never invent a metric the client does not expose.

A **rescue prompt** is human coaching added mainly to make the agent succeed because canonical workflow/project guidance was insufficient or missed. Normal human gates, requested decisions, and factual corrections are not automatically rescue prompts.

## Sessions

| Session | Phase(s) | Model | Reasoning | Fresh / continued | Input | Cached input | Output | Reasoning tokens | Context used | 5h before → after | Weekly before → after | Human prompts | Rescue prompts | Clarifications | Outcome / notes |
|---|---|---|---|---|---:|---:|---:|---:|---:|---|---|---:|---:|---:|---|
| 1 | Exploration | GPT-5.6 Terra | Medium | Fresh | — | — | — | — | — | — | — | — | 0 | — | Not started |

Add rows only when a new actual session starts.

## Quality evidence

Track lifecycle evidence separately from usage totals.

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

## Baseline note

The target code/document baseline for this run is exactly:

```text
Kencleng main
0273c4f2b7f3140efe53a3736547fe73fdd2aefe

Harscode workflow-v2
4199c6db1b26ef1920ba670f222aff0c6d0f9e59
```

The benchmark-seeding commit on the validation branch is instrumentation only; it does not change the product/task baseline.