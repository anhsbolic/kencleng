# Harscode handoff — 0028 frontend workflow execution boundaries

Prepared against Harscode main `5f3c9bcf0d9a14d3ae72314864eab982ffa4947e`.

This handoff captures portable learning from a real Codex frontend dogfood without writing directly to `anhsbolic/harscode-workspace`.

## Why this is proposal-only

The dogfood produced useful workflow evidence, but it did **not** complete the full Build → Review → Testing → PR lifecycle. In addition, Harscode's independent Techplan review prompt is still explicitly `DRAFT` and says it should be exercised against 2+ real Complex-tier techplans before formalization.

For that reason this package stages only a proposed Harscode changelog/proposal entry. It deliberately does not pre-author the final replacements for `workflow/`, `best-practices/`, or `harness-optimization/`.

That keeps the follow-up honest:

```text
real dogfood finding
→ portable proposal
→ human/Harscode review
→ exact Harscode edits only if accepted
```

rather than turning one experiment into immediate global policy.

## Portable findings

### 1. Canonical phase prompts were bypassed

Harscode already has explicit phase-entry prompts such as:

- `workflow/1-exploration-kickoff-prompt.md`;
- `workflow/2-1-techplan-synthesis-prompt.md`;
- `workflow/3-build-prompt.md`;
- `workflow/4-code-review-prompt.md`;
- `workflow/5-testing-prompt.md`.

During dogfood, a custom Codex prompt was authored on top of the guidance instead of starting from the canonical phase prompt. The result worked, but duplicated workflow logic and made it easier to over-specify project/tool mechanics.

Portable lesson: the canonical phase prompt should be the default invocation surface. Harness/project overlays should fill variables and add narrow context, not re-author the phase.

### 2. Frontend specialization does not justify a second lifecycle

The frontend needed different best-practice context, rendered/visual capabilities, design authority, and model/client choices. None of that required different lifecycle semantics.

Portable lesson: keep one Harscode lifecycle. Specialize through:

```text
stack best-practices
+ target-repo authority
+ execution capability/profile
```

rather than creating frontend/backend copies of every phase prompt.

### 3. Execution model/client/capability selection is a separate concern

The dogfood switched between Codex models and exposed that model/client selection was being decided ad hoc inside workflow execution.

Portable lesson: phase policy and execution routing are different axes.

```text
workflow
→ what the phase must accomplish

execution profile
→ model / reasoning effort / client / tools suitable for the work
```

Harscode phase prompts should remain model-neutral. Harness translations may explain where project/user execution profiles belong, but generic Harscode should not hard-code a target project's current Codex model table.

### 4. Techplan review needs an explicit convergence boundary

Independent review found meaningful issues in the dogfood plan, including ambiguous error classification, missing project-state evidence, and browser-environment determinism. That validates independent review as useful signal.

However, after the material findings were resolved, the loop continued into increasingly mechanical details such as exact future command selectors and whitespace in command literals. The planning artifact became the optimization target rather than a sufficient contract for Build.

Portable lesson: a Techplan should be approval-ready when no unresolved issue would force Build to invent a material product/domain decision, violate an authority/security boundary, choose a materially wrong architecture, or lack a meaningful verification strategy.

Mechanical/local implementation detail should normally defer to Build/Patch.

Proposed convergence shape for the still-draft independent review mechanism:

```text
synthesis
→ one independent review when the Complex gate actually applies
→ one resolution pass
→ human gate
```

Re-review is warranted only when the resolution materially changes scope, architecture, business/security semantics, or verification strategy. It should not be automatic after every local/mechanical correction.

This proposal does **not** promote the existing draft review prompt to a mandatory protected gate; its own 2+ real-task evidence threshold still applies.

### 5. Browser automation should not become a lifecycle ritual

A target repo had Playwright available and planning started treating an exact Playwright invocation as part of feature-phase correctness before the browser spec even existed.

Portable lesson: named browser-automation frameworks are tooling, not lifecycle phases or default gates.

Harscode should remain tool-neutral:

- require the evidence appropriate to the behavior/risk;
- allow target repos to require human rendered/product acceptance where appropriate;
- let projects/humans decide when repeatable browser automation is worth adding;
- do not imply that reaching Build, Testing, or PR automatically requires a Playwright/Cypress/etc. run.

This does not ban browser automation and does not reduce objective verification. It separates **acceptance/evidence requirements** from the **choice to automate a browser behavior**.

## Proposed Harscode changes

If accepted, make the smallest edits that establish these boundaries.

### `workflow/README.md`

Add a short cross-stage invocation rule:

- root `*-prompt.md` files are the canonical phase entrypoints;
- use them directly by filling variables;
- target-project/harness overlays may add only context/capability details not already owned by the prompt;
- do not maintain a second hand-authored copy of phase instructions per project or stack.

Add a general convergence principle:

> A phase artifact is sufficient when the next phase can proceed without inventing a material decision owned by an earlier authority. Do not keep polishing a completed phase for local/mechanical details that the next phase can safely resolve.

### `workflow/2-2-techplan-review-prompt.md` (still Draft)

Keep the existing Complex-tier gate and 2+ real-task formalization threshold.

Add a convergence instruction to findings handling:

- classify blockers by whether they change material contract/architecture/security/verification semantics;
- after one resolution pass, re-review only if the fix materially changes one of those concerns;
- formatting, wording, command spelling, local implementation shape, and other unambiguous mechanical corrections are non-blocking/defer-to-Build unless they make the plan unusable;
- do not turn independent review into recursive plan polishing.

### `harness-optimization/codex/README.md`

Clarify that:

- Codex must start from the canonical Harscode phase prompt instead of replacing it with a re-authored Codex version;
- project/user execution profiles may choose current Codex model, reasoning effort, Desktop/CLI client, browser/image/design capabilities, and similar execution mechanics;
- those execution profiles are translation/configuration, not Harscode lifecycle policy;
- do not add a `codex/frontend/` lifecycle fork merely to store model choices or visual-tool preferences.

### `best-practices/react/testing-automation-boundary.md`

Clarify the distinction between verification and automation:

- objective rendered/runtime behavior needs appropriate evidence;
- that evidence may be manual/human or tool-assisted according to project authority and the property being checked;
- committing/running browser automation is justified by repeatability/value/risk, not simply because a feature is frontend or because the workflow reached Testing;
- a named framework such as Playwright is a target-repo tooling decision, not a portable Harscode requirement.

## Explicit non-goals

This proposal does **not**:

- create separate frontend versions of every Harscode phase prompt;
- remove Harscode Testing;
- weaken executable evidence requirements for objective behavior;
- make all frontend verification manual;
- ban Playwright or any browser automation framework;
- add Kencleng-specific commands, paths, risk tiers, or model names to Harscode;
- replace or repair Harscode's generic `best-practices/model-routing.md` work, which should be handled separately;
- promote the draft independent Techplan review mechanism to a hard gate after only one real Complex-tier dogfood.

## Why this belongs in Harscode

The findings are about source ownership and lifecycle/tooling boundaries, not about one React app:

- canonical prompt vs re-authored prompt applies to any harness/project;
- lifecycle vs stack specialization applies to frontend/backend/infra alike;
- execution profile vs workflow policy applies to any model/client/tool family;
- phase convergence applies to any planning artifact;
- verification vs browser-automation choice applies to any UI-capable project.

The target project should still own its exact model table, client preference, human acceptance policy, browser commands, and test-suite layout.

## Evidence posture

Current evidence is intentionally limited:

- one real Codex frontend dogfood exercised Exploration + Techplan + independent review;
- the independent review produced genuine high-value findings;
- the same review loop also demonstrated over-convergence into mechanical detail;
- Build/Code Review/Testing/PR were intentionally not completed as part of this experiment.

Therefore the proposal is strong enough to record and review, but not strong enough to justify broad new ceremony. The intended direction is **simplification and clearer ownership**, not additional gates.
