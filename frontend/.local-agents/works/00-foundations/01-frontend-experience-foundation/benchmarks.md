# Validation 02 — Frontend Experience Foundation Benchmark

> Status: Active measurement — workflow and human rendered acceptance complete; delivery/closeout pending
> Run ID: `validation-02`
> Measurement owner: Anhar Solehudin
> Prepared with: ChatGPT — GPT-5.6 Sol
> Created: 2026-09-16
> Updated: 2026-09-16
> Task authority: `docs/spec/0-foundations/features/01-frontend-experience-foundation.md`
> Task root: `frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation`
> Raw operator log: `benchmark-logs.md`
> Validation branch: `validation-02-frontend-experience-foundation-cleanstart`
> Kencleng target baseline: `main@0273c4f2b7f3140efe53a3736547fe73fdd2aefe`
> Harscode workflow baseline: `workflow-v2@4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Purpose

This file is **operator benchmark synthesis**, not task authority and not a Harscode phase artifact.

`benchmark-logs.md` preserves the observed per-phase `/status` snapshots and operator notes. This file summarizes those observations into benchmark units and quality evidence without changing phase-agent reasoning.

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

## Measurement protocol

### Benchmark unit

One **Codex thread/session** is one benchmark unit, even when the same healthy thread spans more than one workflow phase.

Observed in this run:

- Exploration → Techplan correctly continued in one healthy thread;
- Build used a fresh thread;
- Code Review used a fresh independent thread;
- Testing used a fresh independent thread.

### Primary measurement source

Use `/status` as the primary observable snapshot.

```text
new/resumed thread
→ /status before work
→ perform workflow work normally
→ /status before leaving the thread
```

`/resume` is navigation, not measurement.

Raw snapshots and operator notes belong in `benchmark-logs.md`. This file stores only the synthesized benchmark view.

Do not reconstruct unavailable values or silently repair suspicious measurements. Mark them unknown/ambiguous instead. Operator-recovered corrections should be recorded explicitly as corrections rather than silently replacing the historical capture.

### Rescue prompt definition

A **rescue prompt** is human coaching added mainly to make the agent succeed because canonical workflow/project guidance was insufficient or missed.

Normal human gates, explicit approvals, requested decisions, permission approvals, and factual corrections are not automatically rescue prompts.

## Session summary

| Session | Phase(s) | Model / reasoning | Posture | Observed context / limits | Human interaction | Rescue | Clarifications | Outcome |
|---|---|---|---|---|---|---:|---:|---|
| 1 | Exploration + Techplan | GPT-5.6 Terra · Medium | Fresh → continued | Exploration checkpoint: 68.8K used / 77% left; Techplan final: 89.1K used / 69% left. 5h 100% → 92%; weekly 83% → 82%. | Stage 2 + Stage 3 gates logged; Techplan human approval occurred outside the logged prompt count. | 0 | 0 | Exploration converged; Techplan synthesized; independent review Skip; decomposition Skip; Build fresh-preferred. |
| 2 | Build | GPT-5.6 Sol · Medium | Fresh | Final: 69.8K used / 77% left. 5h 92% → 86%; weekly 82% → 81%. | No task-direction prompt; environment/tool permission approvals only. | 0 | 0 | Static `/` calibration surface implemented; focused verification/build and rendered desktop/mobile inspection passed; human acceptance deferred to the explicit human gate. |
| 3 | Code Review | GPT-5.6 Terra · Medium | Fresh independent | Final: 77K used / 74% left. 5h final 82%; weekly final 81%. **Valid fresh-thread BEFORE snapshot was not captured in the raw log due copy/paste error.** | None logged. | 0 | 0 | Four-pass review approved with no findings; targeted verification passed; no patch plan. |
| 4 | Testing | GPT-5.6 Terra · Medium | Fresh independent | Final: 49.3K used / 85% left. 5h 82% → 80%; weekly 81% → 80%. | No meaningful task-direction prompt logged; raw log contains one blank human-prompt list item. | 0 | 0 | Pass with flagged follow-up: human rendered desktop/mobile acceptance only; no code defect or patch plan. |

## Secondary token totals recorded by the operator

These totals were recorded separately from `/status`; treat them as secondary evidence rather than the benchmark contract.

| Session | Recorded total | Input | Cached input | Output | Reasoning | Data quality |
|---|---:|---:|---:|---:|---:|---|
| 1 — Exploration + Techplan | 48,272 | 42,843 | 470,784 | 5,429 | 1,544 | Recorded |
| 2 — Build | 136,757 | 114,116 | 1,412,096 | 22,641 | 6,228 | Recorded |
| 3 — Code Review | 75,894 | 62,130 | 1,266,560 | 13,764 | 1,460 | **Operator-corrected after identifying the original copy/paste duplication.** |
| 4 — Testing | 78,172 | 71,292 | 447,232 | 6,880 | 2,537 | Recorded |
| **Run total** | **339,095** | **290,381** | **3,596,672** | **48,714** | **11,769** | Secondary aggregate; cached input remains separate and must not be treated as equivalent to non-cached input. |

## Quality evidence

| Signal | Evidence |
|---|---|
| Requirement/authority misses | None surfaced by Code Review or Testing. Testing's fresh Techplan consistency read reported no contradiction/gap. |
| Human redirections caused by agent misunderstanding | None recorded. |
| Rescue prompts | `0` in Exploration/Techplan, Build, Code Review, and Testing. |
| Clarifications | `0` across all recorded sessions. |
| Exploration → Techplan routing | Continued in the same thread because authorities/anchors/decisions were durable and the session remained healthy; Techplan completed without rescue. |
| Independent Techplan review recommendation | `Skip`. Later independent Code Review and Testing found no evidence that the bounded plan omitted a material contract decision. |
| Decomposition recommendation | `Skip`. Build remained a cohesive single implementation slice; no later phase exposed a need for child-task decomposition. |
| Verification rationale/ownership clarity | Build ran focused implementation verification and rendered inspection; Code Review used targeted checks; Testing independently reran final repo verification; human visual acceptance remained explicitly Human-owned. |
| Code Review findings | Approved with no findings across Safety, Quality, stack-specific best practices, and Consistency; no patch plan. |
| Testing findings | Pass with flagged follow-up; no implementation defect; human rendered acceptance was the only remaining gate. |
| Verification repetition / avoidable reruns | `npm run verify`/`build` were run in Build and independently in Testing. This matches authored-code confidence vs final independent evidence; no broad Playwright suite was ceremonially repeated. Code Review also reran narrowly scoped test/lint checks; this remains a candidate efficiency observation for future validations rather than a correctness issue. |
| Browser automation posture | Build used one-off Chromium inspection for desktop/mobile overflow, fragment navigation, and keyboard focus. No committed Playwright scenario was added; Testing did not repeat broad browser automation. |
| Environment friction | Production build needed network access for approved Google Fonts; initial sandbox restrictions were reported rather than misclassified as product failure. |
| Final human rendered acceptance | **PASS for this foundation/calibration scope.** Operator reviewed desktop and mobile renders and reported satisfaction. The result was judged as a coherent foundation rather than a full-content landing page, consistent with the task boundary. |

## Data-quality notes

These are benchmark-observation issues, not product/workflow defects:

1. Code Review's raw **Before** snapshot is the prior Build thread (`gpt-5.6-sol`, Build session id), while the Review **After** snapshot is a different fresh Terra thread. Treat Review's initial `/status` usage/context as unknown.
2. The original footer token totals for Session 3 accidentally duplicated Session 1. The operator later recovered and supplied the corrected Session 3 totals recorded above; preserve the raw log as historical capture and use the corrected values in synthesized analysis.
3. Testing's `Human prompts` section contains a blank numbered entry. Treat meaningful human task-direction prompts as none unless later evidence shows otherwise.
4. Raw `/status` output may expose account-identifying metadata that is irrelevant to benchmarking; omit/redact such fields in future durable captures.

## Current lifecycle state

Workflow correctness gates through human acceptance are complete:

```text
Exploration ✅
→ Techplan synthesis ✅
→ Independent Techplan review skipped by gate ✅
→ Decomposition skipped by gate ✅
→ Build ✅
→ Code Review ✅
→ Testing ✅
→ Human rendered acceptance ✅
→ Pull Request ⏳
→ Validation closeout ⏳
```

Testing's verdict was **Pass with flagged follow-up**, with no code patch required. The flagged follow-up was the Human-owned rendered acceptance, which has now passed for the intended foundation/calibration scope.

## Benchmark hygiene

Do not alter workflow behavior merely to improve benchmark numbers.

In particular, do not:

- keep a stale thread alive only to avoid starting a new session;
- force a fresh thread merely to lower context usage;
- choose a weaker model only for lower measured usage;
- skip justified verification;
- feed benchmark observations into phase-agent reasoning;
- treat lower usage as success when correctness or outcome quality degrades.

## Baseline note

The target code/document baseline for this run is exactly:

```text
Kencleng main
0273c4f2b7f3140efe53a3736547fe73fdd2aefe

Harscode workflow-v2
4199c6db1b26ef1920ba670f222aff0c6d0f9e59
```

Operator instrumentation/history exists only on the validation branch; it does not change the original product/task baseline.