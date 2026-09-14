# Manifest — HOLD candidate 0029 Codex Optional-Capability Fallback Discipline

Prepared against Harscode main `03e25b2a95d1e66f6d60cdcbabf7b1b203b0dad7`.

> **Status: HOLD — no Harscode copy/apply action is authorized by this package.**
>
> Candidate number `0029` is provisional and must be re-checked before activation.

## Current operation

```text
NO COPY TO HARSCODE YET
```

This staging package exists only to preserve the candidate proposal and its activation criteria while Kencleng dogfoods optional Codex capabilities such as image generation.

## Candidate target if activated

Expected primary target:

```text
harscode-workspace/
└── harness-optimization/codex/skills.md
```

Expected Harscode proposal entry after activation:

```text
proposals/<then-current-number>-codex-optional-capability-fallback-discipline.md
```

Do not assume `0029` remains available. Harscode uses a single sequential proposal counter.

## Activation prerequisites

Before any file from this package is copied into Harscode:

1. satisfy the evidence criteria in `PROPOSAL.md`;
2. re-verify current Codex capability/skill behavior;
3. confirm the observed problem is not merely an upstream transient bug;
4. confirm existing Harscode guidance is insufficient;
5. re-check Harscode's highest proposal number;
6. update the candidate proposal to canonical Harscode `Status: Proposed`;
7. reconcile the target wording against current `harness-optimization/codex/skills.md`.

## Expected final scope if activated

Prefer the smallest diff:

```text
CREATE  proposals/<number>-codex-optional-capability-fallback-discipline.md
MODIFY  harness-optimization/codex/skills.md
```

Do not add imagegen-specific workflow, Kencleng asset rules, credentials, model names, or project-specific commands to Harscode.

## Close conditions

Delete this staging package from active Kencleng main after one of these outcomes is settled:

- **Activated and applied:** Harscode owns the proposal/history afterward.
- **Closed / not proposed:** dogfood shows existing guidance is sufficient or issue is tool-specific/transient.
- **Superseded:** another proposal absorbs the capability-fallback concern.

Kencleng Git history remains sufficient as the experiment record.
