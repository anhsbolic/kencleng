# Validation 02 — Frontend Experience Foundation Closeout

> Status: Complete
> Run ID: `validation-02`
> Closed: 2026-09-16
> Task: `docs/spec/0-foundations/features/01-frontend-experience-foundation.md`
> Task root: `frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation`
> Kencleng start baseline: `main@0273c4f2b7f3140efe53a3736547fe73fdd2aefe`
> Kencleng delivered result: `main@71093b687cd7135495bc6ed62d520a621f96f586`
> Delivery PR: `#24 — Establish frontend experience foundation`
> Harscode validation baseline: `workflow-v2@4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Purpose

This is the operator closeout for the second real CRTV run of Harscode workflow-v2. It records what the run demonstrated, what it did not demonstrate, and what should remain under longitudinal observation.

It is validation/process evidence, not Kencleng product authority and not a replacement for Harscode workflow authority.

## Outcome

The task delivered the intended frontend experience foundation without expanding into a full landing page or inventing unresolved Campaign/product semantics.

Observed lifecycle:

```text
Exploration + Techplan in one healthy continued thread
→ fresh Build
→ fresh independent Code Review
→ fresh independent Testing
→ human rendered acceptance
→ PR delivery
```

Final quality outcome:

```text
Rescue prompts: 0
Clarifications: 0
Human redirections caused by agent misunderstanding: 0
Code Review findings: 0
Testing code defects: 0
Patch loops: 0
Human rendered acceptance: PASS
Delivery: PR #24 merged
```

Independent Techplan review was skipped by the plan's gate and later phases exposed no evidence that this omitted a material contract check. Decomposition was also skipped; Build remained a cohesive execution slice and no later phase exposed a need for child-task boundaries.

## Efficiency evidence

Secondary operator-recorded totals:

| Session | Phase(s) | Total | Input | Cached input | Output | Reasoning |
|---|---|---:|---:|---:|---:|---:|
| 1 | Exploration + Techplan | 48,272 | 42,843 | 470,784 | 5,429 | 1,544 |
| 2 | Build | 136,757 | 114,116 | 1,412,096 | 22,641 | 6,228 |
| 3 | Code Review | 75,894 | 62,130 | 1,266,560 | 13,764 | 1,460 |
| 4 | Testing | 78,172 | 71,292 | 447,232 | 6,880 | 2,537 |
| **Run total** |  | **339,095** | **290,381** | **3,596,672** | **48,714** | **11,769** |

Cached input remains separate evidence of replay/reuse and is not treated as equivalent to non-cached input cost.

The strongest efficiency signal is structural rather than numeric: Exploration → Techplan reused a healthy session successfully, optional Techplan review/decomposition were skipped without later quality loss, no patch loop occurred, and broad browser automation was not turned into ritual.

## What this run supports

Validation 02 provides positive evidence for:

- progressive context loading instead of recursively loading all nearby guidance;
- continuation-fitness routing for Exploration → Techplan;
- fresh Build after an Approved Techplan for a material UI task;
- fresh independent Code Review and Testing;
- durable artifact handoff instead of chat-memory dependency;
- early Techplan recommendations that can skip unnecessary independent review/decomposition;
- Build vs Testing verification ownership separation;
- browser automation as risk-driven capability rather than lifecycle ritual;
- explicit Human-owned rendered acceptance for material UI;
- preserving task/product truth while avoiding speculative shared abstractions.

## What this run does not yet validate

This task did not exercise several important workflow paths:

- Complex Techplan independent review;
- Techplan decomposition into multiple execution tasks;
- Code Review → Build/Patch re-entry;
- Testing → Build/Patch re-entry;
- API/data-contract implementation;
- auth/security-sensitive implementation;
- cross-service, migration, destructive, or major state-ownership changes.

Those paths should be evaluated through future real Kencleng work rather than synthetic validation tasks unless a real need appears.

## Efficiency observations to watch, not fix yet

Two areas are worth longitudinal observation:

1. Code Review consumed substantial context for a small/static diff and reran narrow lint/test checks. This was not a correctness problem, but repeated similar evidence may justify reducing Review startup/loading or verification cost.
2. Testing remained comparatively context-heavy for a bounded task. The current workflow intentionally keeps a fresh whole-Techplan consistency read during early workflow-v2 validation, so one run is not enough evidence to remove that safety net.

Do not refine workflow-v2 from these observations alone. Promote them to workflow changes only when repeated real-task evidence shows a stable inefficiency or quality problem.

## Benchmark-data learning

The run exposed one operator measurement issue: the raw Code Review BEFORE `/status` snapshot and initial copied token totals were incorrect. Corrected Session 3 totals were later recovered and the synthesized benchmark records the correction explicitly.

For future runs:

```text
new/resumed Codex thread
→ /status before work
→ normal workflow execution
→ /status before leaving the thread
```

Keep raw capture in `benchmark-logs.md`, synthesize conclusions separately, and omit account-identifying fields that do not contribute to measurement.

## Workflow-v2 disposition

Based on Validation 01 plus this second different real task, workflow-v2 has enough positive evidence to move from **Validated Candidate** to **Operational Default**.

That disposition means:

- the workflow is suitable to become Harscode `main`;
- it should be used normally in continued Kencleng development;
- the merge is not a claim that every lifecycle path is broadly validated or permanently stable;
- benchmarking should continue longitudinally across real task shapes;
- future refinements should be evidence-driven and preferably based on recurring patterns rather than one-off token optimization.

A useful maturity interpretation is:

```text
Level 0 — Experimental
Level 1 — Validated Candidate
Level 2 — Operational Default      ← current evidence supports this level
Level 3 — Broadly Validated
Level 4 — Mature / Stable
```

Future Kencleng tasks should record the actual Harscode `main` revision used, so outcome and efficiency evidence remain attributable as the workflow evolves.

## Closeout decision

Validation 02 is complete.

The next workflow-level action is to reconcile current Harscode `main` into `workflow-v2`, resolve any integration-only conflicts without speculative workflow redesign, then propose the validated workflow-v2 line for merge to `main` with continued longitudinal benchmarking as an explicit post-merge expectation.
