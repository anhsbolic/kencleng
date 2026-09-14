# 0029 — Codex optional-capability fallback discipline

**Status:** HOLD CANDIDATE — staging only; not a valid Harscode proposal status
**Date:** 2026-09-14
**Protection Tier:** general
**Triggered by:** Preparing to dogfood optional Codex CLI capabilities such as image generation after proposal 0028 established execution profiles as separate from lifecycle policy.
**Target area if activated:** harness-optimization/codex
**Target file(s) if activated:**
- `harness-optimization/codex/skills.md`

> Do not copy this file into Harscode while it remains HOLD. Before activation, re-check proposal numbering and change `Status` to canonical `Proposed`.

## Candidate gap

Codex Skills are adapters/progressive-disclosure mechanisms, not proof that every backing optional capability is available in the current environment or session.

The candidate portable failure shape is:

```text
skill / wrapper is discoverable
→ agent assumes backing capability is available
→ capability is missing, unsuitable, or needs an unauthorized fallback
→ agent silently degrades, fabricates success, or uses the wrong authority path
```

Potentially affected capability classes include image generation, browser/computer-use tooling, design/Figma integration, network-backed specialized tools, and other hosted/session-dependent capabilities.

No recurring Harscode problem has been established yet. This candidate exists to preserve the idea and its evidence threshold while real dogfood runs.

## Candidate change if evidence justifies it

Add a small capability-generic rule to `harness-optimization/codex/skills.md`:

```text
Optional capability discipline

A skill's presence does not prove its backing capability is available.
Before relying on an optional hosted/tool capability:

1. verify that the backing capability is available in the current environment/session;
2. use the canonical capability when available;
3. use a fallback only when it is explicitly supported and authorized by target-repo rules or user intent;
4. if no valid capability path exists, surface the capability gap instead of fabricating or silently degrading the requested result;
5. keep vendor-specific operating details in current tool documentation or the target project, not in portable Harscode policy.
```

Image generation may motivate the first dogfood, but `$imagegen` must not become a Harscode-specific workflow requirement.

## Activation criteria

Activate this candidate only if evidence shows at least one of:

1. **Recurring real-task failure:** two or more real tasks show the same capability-boundary problem.
2. **Cross-capability recurrence:** the same failure shape appears in at least two optional capability classes.
3. **One structural authority/correctness incident:** for example unauthorized credential/network fallback or a false claim that a tool-backed artifact was produced.

Before activation also confirm:

- the issue is not merely a transient upstream Codex bug;
- current Harscode `harness-optimization/codex/skills.md` is insufficient to prevent it;
- the observed problem belongs to harness translation rather than target-project asset/design policy;
- current Codex behavior has been re-verified;
- the proposal number is re-checked because `0029` is not reserved while HOLD.

## Do not activate for

- normal successful `$imagegen` use;
- a one-off vendor regression with a clear upstream fix;
- project-specific asset-policy gaps;
- prompt-writing uncertainty;
- preferences between image models/tools;
- missing credentials already explained by the explicitly chosen tool path;
- a desire to mirror fast-moving Codex product documentation in Harscode.

## Explicit non-goals

This candidate does not:

- require image generation;
- add an image-generation phase;
- define asset direction, approval, or storage conventions;
- add API-key or credential instructions;
- create a frontend-specific Harscode lifecycle;
- guarantee every optional capability has a fallback;
- permit silent design/verification degradation when a tool is unavailable.

## Resolution after dogfood

```text
ACTIVATE
→ reconcile against current Harscode main
→ assign current proposal number
→ change status to Proposed
→ human review through normal Harscode process

CLOSE / DO NOT PROPOSE
→ existing guidance or upstream tooling already handles the cases

SUPERSEDE
→ a broader capability-routing proposal absorbs the same concern
```

Preferred outcome: no new Harscode rule unless real evidence proves the portable gap exists.
