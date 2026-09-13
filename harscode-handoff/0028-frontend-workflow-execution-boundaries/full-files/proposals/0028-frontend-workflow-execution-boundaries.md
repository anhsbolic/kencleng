# 0028 — Canonical phase prompts, execution profiles, Techplan convergence, and browser-automation boundary

**Status:** Proposed
**Date:** 2026-09-13
**Protection Tier:** general
**Triggered by:** A real Codex frontend dogfood that completed Exploration + Techplan + independent review. The run surfaced both useful independent-review findings and avoidable process cost from re-authoring canonical prompts, ad-hoc model/client routing, recursive plan polishing, and treating a target project's browser automation as if it were a workflow gate.
**Target area:** workflow + harness-optimization + best-practices
**Target file(s), if accepted:**
- `workflow/README.md`
- `workflow/2-2-techplan-review-prompt.md` (still Draft)
- `harness-optimization/codex/README.md`
- `best-practices/react/testing-automation-boundary.md`

## Gap found

The dogfood exposed five related ownership problems.

### 1. Canonical phase prompts can be bypassed by custom harness/project prompts

Harscode already has paste-ready phase entrypoints such as `1-exploration-kickoff-prompt.md`, `2-1-techplan-synthesis-prompt.md`, `3-build-prompt.md`, `4-code-review-prompt.md`, and `5-testing-prompt.md`.

A custom prompt was nevertheless authored on top of the guidance during dogfood. It worked, but duplicated phase logic and made it easier to carry project/tool-specific mechanics into the wrong layer.

The workspace currently explains prompt shape and path variables, but does not state strongly enough that the root `*-prompt.md` files are the default invocation surface and should be adapted narrowly rather than re-authored per project/harness.

### 2. Stack specialization and lifecycle specialization are easy to conflate

Frontend work genuinely needs different best-practice context, visual/rendered capability, design authority, and sometimes different model/client selection. None of those observations imply a different lifecycle.

Without an explicit boundary, the natural response is to create frontend/backend copies of phase prompts, which would duplicate workflow policy and drift over time.

### 3. Model/client/tool choice has no explicit layer separate from workflow policy

The dogfood moved between Codex models and clients based on perceived task difficulty. That decision is real, but it is not the same thing as Harscode lifecycle policy.

A reusable separation is:

```text
workflow phase
→ what must be accomplished

execution profile
→ model / reasoning effort / client / capabilities suitable for the work
```

Harscode's generic prompts should remain model-neutral. Harness translations and target projects may define execution profiles without turning model names or UI/tool preferences into lifecycle rules.

### 4. Independent Techplan review lacks an explicit convergence boundary

The still-draft Techplan review prompt produced meaningful signal in the dogfood: it found ambiguous error classification, missing project-state evidence, and browser-environment determinism that the synthesis pass missed.

After those material findings were resolved, however, the loop continued into increasingly mechanical details such as exact future command selectors and command-literal whitespace. The plan became the object being perfected rather than a sufficient contract for Build.

The current draft prompt defines what to review but not when a review-fix cycle has converged.

### 5. Browser automation can be mistaken for a phase obligation

A target repository had Playwright available. During planning, exact Playwright execution mechanics began to be treated as part of Techplan approval before the feature browser spec existed.

That is a category error: a named browser-automation framework is tooling. Harscode should require evidence appropriate to the behavior/risk while leaving the choice to automate a browser behavior to target-project policy/human judgment.

## Proposed change

### A. Make canonical phase prompts the default invocation surface

Add a short cross-stage rule to `workflow/README.md`:

- the root `*-prompt.md` file for a phase is the canonical manual invocation surface;
- fill its variables and adapt only where the prompt explicitly allows adaptation;
- project/harness overlays may add narrow context/capability details not already owned by the prompt;
- do not maintain a hand-authored per-project or per-stack copy of the phase instructions unless a proposal establishes a real lifecycle difference.

This does not prohibit wrappers/skills/slash commands. Those should route to or mechanically wrap the canonical prompt, not restate it as independently maintained policy.

### B. Keep one lifecycle; specialize through guidance and execution capability

Document the preferred layering:

```text
Harscode phase prompt
+ matching stack best-practices
+ target-repo authority
+ harness/project execution profile
```

Do not create `frontend-*` / `backend-*` copies of every workflow prompt by default.

A track-specific harness folder remains justified only when the **harness translation itself** genuinely differs, consistent with the existing `harness-optimization/` rule.

### C. Recognize execution profiles without moving model routing into phase prompts

Clarify in `harness-optimization/codex/README.md` that a target project or user may maintain a Codex execution profile covering fast-moving mechanics such as:

- current model choice;
- reasoning effort;
- Desktop vs CLI client;
- browser/rendered capability;
- image/design capability;
- other harness-specific tool availability.

Such a profile must not redefine Harscode lifecycle, project truth, approval authority, or testing obligations.

Do not add specific Codex model names to generic Harscode as part of this proposal. Generic `best-practices/model-routing.md` remains a separate concern and is expected to be revised independently.

### D. Add an explicit Techplan convergence/stopping rule

Keep `workflow/2-2-techplan-review-prompt.md` in Draft status and keep its own requirement to exercise the mechanism on 2+ real Complex-tier plans before formalization.

Within that draft mechanism, add a convergence rule:

A Techplan is approval-ready when no unresolved issue would require Build to:

- invent a material product/domain decision;
- violate an authority/security boundary;
- choose a materially different architecture/state/component ownership model;
- proceed without a meaningful verification strategy.

Default review shape:

```text
synthesis
→ one independent review when the Complex gate actually applies
→ one resolution pass
→ human gate
```

Re-review only when the resolution materially changes scope, architecture, business/security semantics, or verification strategy.

Do not automatically re-review for unambiguous mechanical/local details such as wording cleanup, formatting, exact command spelling, or implementation shape that Build/Patch can safely resolve. Those become blocking only when the defect makes the contract unusable or changes meaning.

This is a convergence rule, not permission to ignore material contradictions.

### E. Separate rendered verification from browser automation

Update `best-practices/react/testing-automation-boundary.md` to make the distinction explicit:

- objective runtime/rendered outcomes need appropriate evidence;
- the evidence mechanism depends on the property and target-project authority;
- human/manual rendered acceptance may be the correct evidence for subjective product/UX judgment;
- browser automation is useful when repeatability, regression value, or risk justifies it;
- reaching a lifecycle phase does not itself justify adding/running a named browser framework;
- Playwright/Cypress/etc. remain target-repo tooling choices, not Harscode requirements.

This retains the existing principle that mechanically detectable checks should be automated, while preventing "frontend == browser automation every phase" from becoming an unintended interpretation.

## Explicit non-goals

This proposal does **not**:

- create a new frontend lifecycle;
- create separate frontend versions of every phase prompt;
- remove Harscode Testing;
- weaken objective verification requirements;
- require all rendered verification to be manual;
- ban Playwright or another browser framework;
- add project-specific browser commands or paths to Harscode;
- add target-project model names to Harscode;
- replace or repair generic `best-practices/model-routing.md`;
- promote the draft independent Techplan review prompt to a hard gate after only one real Complex-tier dogfood.

## Rationale

These changes preserve the workspace's existing single-source-of-truth discipline.

A phase prompt should own phase mechanics once. Stack-specific correctness should come from `best-practices/`. Project truth should come from the target repo. Harness mechanics should come from `harness-optimization/` or project/user execution profiles. A browser framework should remain tooling rather than a hidden lifecycle phase.

The Techplan convergence rule applies the same principle temporally: once the artifact is sufficient for the next phase without inventing a material decision, continuing to polish local/mechanical detail adds cost without proportional correctness signal.

The intended outcome is **less ceremony with clearer ownership**, not a new layer of mandatory artifacts.

## Evidence posture

Evidence so far:

- one real Codex frontend dogfood completed Exploration + Techplan + independent review;
- independent review found genuine material issues missed by synthesis;
- the same run demonstrated recursive over-convergence after the material issues were resolved;
- Build, Code Review, Testing, and PR were intentionally not completed in that experiment.

This is enough to propose the ownership/convergence corrections above, but not enough to formalize the still-draft independent review mechanism beyond its existing evidence threshold.

---

*If accepted, apply the smallest edits to the target files above, keep the Techplan review prompt Draft until its existing evidence threshold is met, and avoid adding any target-project model table or browser command to Harscode.*
