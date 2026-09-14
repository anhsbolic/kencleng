# Harscode handoff — candidate 0029 Codex optional-capability fallback discipline

Prepared against Harscode main `03e25b2a95d1e66f6d60cdcbabf7b1b203b0dad7`.

> **Status: HOLD — staging only. Do not apply to Harscode yet.**
>
> `0029` is a provisional candidate number only. Re-check Harscode's proposal sequence before activation.

## Why this exists

Kencleng's frontend workflow now treats model/client/tool choice as an execution-profile concern and explicitly recognizes image/design capability as one possible Codex execution capability. The next practical question is how Codex CLI should behave when a requested optional capability — initially image generation — is represented by a skill or documented mechanism but the backing tool is unavailable, unavailable in the active session, unsuitable for the requested output, or requires an explicitly authorized fallback.

The current Harscode Codex skill guidance already establishes the right broad boundary:

```text
Codex Skill
→ adapter / progressive-disclosure mechanism
→ not a new policy source
```

That means `$imagegen` itself does **not** justify a Harscode proposal. Image-generation specifics belong to current Codex capability documentation and the target project's own asset/design authority.

The potentially portable gap is narrower:

> **A discoverable skill or wrapper does not necessarily prove that its backing optional capability is actually available and usable in the current execution environment.**

If real dogfood proves that this distinction causes repeated failure or unsafe fallback behavior, Harscode may benefit from one small Codex-harness rule for capability detection and safe degradation.

## Candidate portable problem

An agent may encounter an optional capability such as:

```text
image generation
browser / computer-use capability
design/Figma integration
network-backed specialized tools
other hosted or session-dependent capabilities
```

and incorrectly infer:

```text
skill exists
→ capability must be available
→ proceed as if the requested result can be produced
```

Possible failure modes include:

- claiming a tool-backed result that was never actually produced;
- silently substituting a materially weaker artifact because the intended capability is unavailable;
- invoking an API/credential fallback without target-project or human authority;
- generating a placeholder and presenting it as equivalent to the intended asset;
- writing output outside the target workspace and assuming integration succeeded;
- escalating model/reasoning effort when the real problem is missing execution capability;
- duplicating fast-moving vendor-specific tool instructions inside Harscode.

These are plausible failure modes, not yet established recurring Harscode problems. That is why this candidate is HOLD.

## Candidate change if activated

If the activation criteria below are met, propose the smallest possible addition to:

```text
harness-optimization/codex/skills.md
```

Possible portable rule:

```text
Optional capability discipline

A skill's presence does not prove its backing capability is available.
Before relying on an optional hosted/tool capability:

1. verify that the required backing capability is available in the current environment/session;
2. use the canonical capability when available;
3. use a fallback only when that fallback is explicitly supported and authorized by the target project's rules/user intent;
4. if no valid capability path exists, surface the capability gap instead of fabricating or silently degrading the requested result;
5. keep vendor/tool-specific operating details in the current harness/tool documentation or target project, not in portable Harscode policy.
```

The rule should remain capability-generic. Image generation may be the first motivating case, but Harscode should not gain an `$imagegen`-specific workflow or asset policy.

## What would make this worth following up

Move this candidate from **HOLD** toward a real Harscode `Proposed` entry only when at least one of these evidence shapes exists:

### A. Recurring real-task evidence

Two or more real tasks show the same capability-boundary failure, for example:

- Codex discovers/uses a skill but cannot access its backing tool;
- an agent repeatedly chooses an unauthorized or misleading fallback;
- generated output repeatedly fails to land in the expected workspace while the agent treats the operation as successful;
- missing capability repeatedly causes silent design/verification degradation.

### B. Cross-capability evidence

The same failure shape appears across at least two optional capability classes, for example image generation plus browser/design integration. This is especially strong evidence that the gap is generic harness discipline rather than an imagegen-specific issue.

### C. One genuinely structural incident

A single incident may be enough if it demonstrates a clear authority/correctness problem that existing Harscode guidance cannot safely resolve — for example an agent uses credentials/network/API fallback without authorization or falsely reports a capability-backed artifact as produced.

## What is NOT enough to activate it

Keep this candidate HOLD when the evidence is only:

- `$imagegen` works normally in the current Codex CLI setup;
- one temporary Codex regression with a clear upstream fix/workaround;
- a target-project asset-policy gap that belongs in Kencleng or another repo;
- uncertainty about how to prompt image generation;
- preference for one image model/tool over another;
- desire to document current Codex product mechanics inside Harscode;
- a missing API key where the user explicitly chose an API-based path and existing tool documentation already explains the requirement.

Before activation, re-verify current Codex capability/skill behavior. Fast-moving product mechanics must not be fossilized into Harscode from stale observations.

## Activation checklist

Before changing this candidate to a formal Harscode proposal:

- [ ] Record concrete dogfood task(s) and exact observed failure shape.
- [ ] Separate target-project policy issues from Codex/harness capability issues.
- [ ] Confirm the issue is not merely a transient upstream Codex bug already fixed.
- [ ] Confirm existing `harness-optimization/codex/skills.md` does not already give enough guidance to avoid the failure.
- [ ] Re-check whether the failure generalizes beyond image generation or recurs often enough to justify a portable rule.
- [ ] Re-check Harscode proposal numbering; `0029` is not reserved while this candidate is HOLD.
- [ ] Convert the Harscode-facing candidate status to canonical `Proposed` before copying it into `harscode-workspace/proposals/`.
- [ ] Keep the target edit minimal, preferably `harness-optimization/codex/skills.md` only unless evidence proves another file is necessary.

## Explicit non-goals

This candidate does **not** propose to:

- require `$imagegen`;
- document how to use GPT Image or any current vendor-specific image API;
- add an image-generation phase to Harscode;
- define project asset direction, approval, or storage conventions;
- make Codex frontend-specific lifecycle guidance;
- encode API keys, credentials, or network permissions into Harscode;
- require every optional capability to have a fallback;
- treat a tool limitation as permission to silently lower product/design requirements.

## Kencleng dogfood posture

Until activation, Kencleng should simply exercise its existing authority model:

```text
asset/design need
→ Kencleng design + asset authority decides what is required
→ execution profile chooses a capable environment/tool
→ try the supported capability
→ if unavailable, use only an authorized fallback
→ otherwise surface the capability gap / produce the existing asset handoff
→ human acceptance where Kencleng requires it
```

The experiment should observe actual Codex CLI behavior rather than adding new Kencleng or Harscode ceremony in advance.

## Hold resolution

After sufficient dogfood, resolve this candidate in one of three ways:

```text
ACTIVATE
→ evidence shows a portable recurring/structural gap
→ reconcile against current Harscode main
→ assign the then-current proposal number
→ convert status to Proposed
→ copy into Harscode for normal human review

CLOSE / DO NOT PROPOSE
→ existing Codex/Harscode/project guidance already handles the cases
→ retain Kencleng Git history as experiment record

SUPERSEDE
→ a broader capability-routing proposal absorbs the same concern
```

The preferred outcome is still **no new Harscode rule unless dogfood proves one is needed**.
