# Kencleng — Codex Frontend Execution Profile

> Status: Current project-specific execution guidance
>
> Last reconciled: 2026-09-16
>
> Purpose: Route Codex client/model/capability choices for Kencleng frontend work without redefining the Harscode lifecycle or generic model-routing policy.

## 1. Authority boundary

This file is an **execution profile**, not a workflow.

Use these layers separately:

```text
Harscode canonical phase prompt
→ how the current lifecycle phase runs

Kencleng AGENTS/spec/design/architecture authorities
→ what is true for this project

this execution profile
→ which Codex client/model/capabilities are a good fit for the work
```

Do not replace a Harscode phase prompt with a project-authored mega-prompt. Start from the current canonical Harscode prompt for the phase, fill its variables, and add only narrow Kencleng/task context that the prompt does not already own.

Frontend and backend share the same Harscode lifecycle. Do not create a parallel set of frontend-specific Exploration/Techplan/Build/Review/Testing prompts merely because the implementation stack differs. Frontend specialization comes from the relevant React/frontend best-practices, Kencleng frontend authorities, and the execution capabilities selected here.

This file is intentionally separate from Harscode generic model routing. When Harscode's generic routing is revised, reconcile this profile rather than creating two competing generic policies.

## 2. Client and capability routing

Choose the environment according to what the work needs, not according to phase name alone.

| Work | Preferred Codex environment | Why |
| --- | --- | --- |
| Exploration / repository discovery | CLI or Desktop | Primarily read/search/reasoning work. |
| Techplan synthesis | CLI or Desktop | Primarily contract/architecture reasoning; no visual client requirement by default. |
| Routine frontend Build | CLI or Desktop | Code-centric implementation with established UI intent. |
| Material UI / product-design Build | **Desktop preferred** | Better fit for rendered iteration, screenshots/design references, browser feedback, and asset/image work when needed. |
| Code Review | Fresh CLI or Desktop session | Independence matters more than visual tooling by default. |
| Testing | CLI or Desktop according to the verification target | Use the cheapest environment that can objectively verify the required behavior. |
| PR/report/admin work | CLI or Desktop | Low-complexity artifact work; avoid expensive capability by default. |

A tool limitation must not silently become a design limitation. If material frontend work requires rendered inspection, screenshot/design context, image generation/editing, or other visual capability that the current client cannot provide adequately, switch to a capable environment rather than degrading the design.

## 3. Codex model routing

Use the **smallest sufficient capability**, then escalate from evidence.

Current project defaults:

| Work | Default | Escalate when |
| --- | --- | --- |
| Exploration | GPT-5.6 Terra · Medium | GPT-5.6 Sol · Medium when architecture/design authority is genuinely ambiguous or cross-cutting. |
| Techplan synthesis | GPT-5.6 Terra · Medium | GPT-5.6 Sol · Medium for consequential architecture/state/component decisions. |
| Independent Techplan review, when actually warranted | GPT-5.6 Sol · Medium | GPT-6 Astra · Low/Medium only for unusually difficult or novel review problems. |
| Routine frontend Build | GPT-5.6 Terra · Medium | Sol when the implementation becomes materially architectural/visual/cross-cutting. |
| Material product/UI Build | **GPT-5.6 Sol · Medium** | Astra Low/Medium for genuinely hard multi-tool, multimodal, or unfamiliar work. |
| Mechanical/local patch | GPT-5.6 Terra · Low/Medium | Escalate only if the patch exposes a materially harder problem. |
| Code Review | GPT-5.6 Terra · Medium in a fresh session | Sol Medium for risky/cross-cutting diffs or difficult root-cause analysis. |
| Testing/debugging | GPT-5.6 Terra · Medium | Sol Medium for difficult behavioral/rendered failures; Astra Low/Medium for exceptional end-to-end debugging. |
| PR/report/admin | GPT-5.6 Luna or Terra · Low | Rarely needs escalation. |

Model availability and product controls are fast-moving. Re-verify this table when Codex changes model families, reasoning controls, or materially different capability/cost trade-offs.

## 4. Reasoning effort

Reasoning effort is not a severity badge.

```text
Low
→ mechanical, repetitive, clerical, or tightly bounded changes

Medium
→ default for normal engineering work

High
→ difficult but bounded reasoning where Medium has shown insufficient depth

Astra Low/Medium
→ exceptional multi-tool / multimodal / unfamiliar / repeated-failure cases
```

Do not default to High simply because a task matters. Missing context, missing permissions, missing files, or missing visual/browser capability must be fixed as capability/context problems; more reasoning does not substitute for them.

## 5. Escalation triggers

Stay on the current model when the task only needs more file reading, one or two normal implementation iterations, ordinary test fixes, or local visual tuning.

Escalate **Terra → Sol** when one or more of these becomes true:

- architecture or state ownership is materially ambiguous;
- existing production patterns conflict;
- a change has meaningful blast radius across shared components or API abstractions;
- product/visual judgment materially determines the user experience;
- the task requires coordinated reasoning across several frontend layers;
- a first well-grounded attempt failed because the reasoning, not the context/tooling, was insufficient.

Escalate **Sol → Astra** only when the problem genuinely changes class, for example:

- unfamiliar multi-step debugging across several systems/tools;
- browser + code + design/visual + assets must be reasoned about as one problem;
- repeated well-grounded attempts fail without a clear root cause;
- strongest independent reasoning is justified by the stakes and complexity.

## 6. Frontend Build classification

Treat these as different execution shapes even though both remain Harscode Build:

### Routine frontend engineering

Examples:

- API plumbing against an established contract;
- local refactor;
- test addition;
- straightforward component implementation from an established pattern;
- generated-client integration;
- mechanical fixes.

Default: **Terra · Medium**, CLI or Desktop.

### Material product/UI engineering

Examples:

- layout/composition judgment;
- responsive translation;
- interaction design interpretation;
- prototype/reference → production translation;
- new visual pattern;
- expressive/semantic asset work;
- complex accessibility behavior;
- significant visual hierarchy decisions.

Default: **Sol · Medium**, Desktop preferred.

This classification is about implementation cognition/capability. It is independent from Kencleng risk tier and design readiness.

## 7. Rendered iteration and human acceptance

During material frontend Build, agent-driven `edit → render → inspect → fix` is normal implementation feedback when the chosen environment supports it.

That agent feedback is **not final product acceptance**.

Before merge/delivery of material frontend UI, a human must manually exercise the rendered result in a real browser and accept the observable experience at a representative desktop/mobile scope appropriate to the change.

Human acceptance should concentrate on what automation and code review cannot fully decide:

- visual hierarchy and clarity;
- interaction comprehension;
- responsive usability;
- product/design intent;
- whether assets feel appropriate;
- obvious state/error behavior in context.

Keep the acceptance scope proportional to the change; this is not a requirement for an exhaustive manual test matrix on every PR.

## 8. Playwright boundary

Playwright is an **independent, on-demand browser automation capability**, not a Harscode phase requirement for Kencleng.

Ordinary Exploration, Techplan, Build, Code Review, Testing, and PR sessions must not add or run broad Playwright coverage merely because the lifecycle reached a particular phase.

For a non-trivial proposed browser check/regression, the planning/verification contract should state:

```text
why browser automation is useful
risk if omitted
which phase owns the authoritative run
```

An agent may recommend Playwright when concrete browser, interaction, responsive, or regression risk justifies repeatable automation. It becomes part of the task when the human explicitly requests it **or approves a task/spec/Techplan that includes that verification contract**. Do not require a second redundant permission prompt after the human has already approved the plan containing it.

Code Review may use a targeted browser reproduction when needed to prove or disprove a suspected finding; this is not permission to rerun broad final verification by default. Build should use focused implementation feedback. Broad/final independent verification belongs to Testing when the workflow uses a Testing phase.

Committed browser tests live under `frontend/tests/browser/` and are scoped by the behavior they protect, for example:

```text
smoke
feature flow
regression for a known bug
broad critical-journey regression
```

The workflow task path does not define Playwright scope. The protected behavior and approved verification contract define the automation breadth.

Human rendered acceptance and Playwright serve different purposes:

```text
human rendered acceptance
→ product/UX acceptance before delivery

Playwright
→ optional repeatable regression automation when risk justifies it
```

Do not use Playwright availability as a reason to turn every visual check into a permanent E2E test.

## 9. Reconciliation rule

If this profile conflicts with:

- Harscode lifecycle policy → Harscode owns the lifecycle;
- Kencleng product/spec/design truth → the Kencleng authority owns the product decision;
- current Codex mechanics/model availability → re-verify the execution profile rather than rewriting project/workflow truth to fit stale tooling assumptions.
