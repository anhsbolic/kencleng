# Handoff Proposal — Harscode 0027 Codex Harness Translation

> Staging location only. This file is not Harscode policy and is not Kencleng project authority.
>
> Prepared against `anhsbolic/harscode-workspace` main at `ee8fba26e9d378220fd2c6facc9e80306458a2c2`.

## Why this handoff exists

Kencleng has now exercised the project-side prerequisites we wanted before translating Harscode into Codex-native mechanics:

- authority/instruction cleanup is merged;
- shared component governance was dogfooded through the Button Secondary migration;
- Playwright browser verification is wired and proven separately from the fast Vitest baseline.

The next risk is not missing frontend guidance. It is **duplicating existing guidance into a second Codex-specific policy layer**.

The proposed Harscode change therefore keeps the same boundary already established by `harness-optimization/`:

> Translation only. Never policy.

## Challenge-first review

Several tempting approaches were rejected deliberately.

### Rejected: a giant `CODEX.md`

Codex already consumes hierarchical `AGENTS.md` instructions, with more-specific scoped instructions layered later. A second always-loaded project instruction file would duplicate target-repo authority and create another drift surface.

### Rejected: copying Harscode workflow prose into target-repo `AGENTS.md`

Harscode owns portable lifecycle/process guidance; the target repo owns project-specific hard rules and routing. Root/nested `AGENTS.md` should point to the relevant Harscode source, not embed it wholesale.

### Rejected: mandatory Codex subagents for every phase

Codex has fast-moving multi-agent support, but Harscode already has a harness-agnostic four-session boundary. Native subagents may optimize isolation where verified, but correctness should not depend on an unstable harness feature.

### Rejected: a frontend-specific Codex folder immediately

Nothing in the initial Codex translation currently differs meaningfully by frontend/backend track. Creating `codex/frontend/` now would violate Harscode's own rule that track splits exist only when the harness translation genuinely differs.

### Rejected: putting project-specific browser commands into Harscode

Kencleng's `npm run test:browser` proved the need for executable rendered verification, but that command belongs to Kencleng. Harscode should translate the requirement as “use the target repo's native rendered/browser verification capability,” not copy Kencleng's tooling instance.

## Proposed Harscode shape

```text
harness-optimization/
└── codex/
    ├── README.md
    ├── instruction-loading.md
    ├── skills.md
    ├── session-boundaries.md
    ├── permissions-and-sandbox.md
    └── token-optimization.md
```

No Harscode root policy files need to change. The existing generic `<harness-name>/` convention already covers this sibling.

## Core decisions

### 1. Hierarchical `AGENTS.md` is the primary project instruction surface

Codex natively composes instruction files from the user/Codex-home layer and the project tree. Target repositories should keep `AGENTS.md` compact: hard rules, commands, write fences, and routing to canonical project/Harscode sources.

`AGENTS.override.md` is treated as an explicit override mechanism, not the normal place for portable policy.

### 2. Skills are progressive-disclosure adapters

Codex skills are a good fit for Harscode guidance that should be discoverable without loading full content into every turn. The translation should keep skills thin and mechanically derived from existing Harscode sources rather than inventing new rules.

The initial recommendation is documentation/pattern guidance only, not checked-in project-specific skills inside Harscode. Actual `.codex/skills/` instances belong in the target repo or user's Codex setup.

### 3. Harscode session boundaries remain the portable isolation baseline

Map the existing four-session guidance directly:

```text
Exploration + Techplan
Build + Patch
Code Review
Testing
```

Pull-request generation follows verified output afterward. Native Codex subagents may be used when they demonstrably improve isolation, but they are an optional harness optimization rather than a new lifecycle requirement.

### 4. Sandbox/approval controls support existing authority; they do not define it

Codex sandbox, network, and approval settings can reinforce read-only review/testing and narrow write scopes. They must not silently weaken project/Harscode human-review requirements or encode a new risk taxonomy.

### 5. Token optimization gets one Codex-native persistent instruction surface

Translate the existing harness-agnostic rule using one persistent Codex instruction location, preferably a dedicated Harscode Codex profile's `developer_instructions` or a single user/project `AGENTS.md` location as appropriate. Do not install the same rule in multiple always-loaded surfaces.

## External capability verification

The proposal was checked against current OpenAI/Codex documentation on 2026-09-13:

- Codex composes `AGENTS.md` / `AGENTS.override.md` hierarchically and supports project document fallback filenames through `config.toml`.
- Codex skills use progressive disclosure: selection metadata first, `SKILL.md` body when used, supporting resources on demand.
- Codex exposes sandbox/approval controls and current configuration includes multi-agent/subagent settings, but these mechanics are fast-moving and must stay under the harness re-verification disclaimer.

Primary references used while preparing the handoff:

- https://openai.com/index/unrolling-the-codex-agent-loop/
- https://openai.com/index/running-codex-safely/
- https://github.com/openai/codex/blob/main/codex-rs/core/config.schema.json
- https://github.com/openai/codex/blob/main/codex-rs/skills/src/assets/samples/skill-creator/SKILL.md

## Acceptance criteria

The Harscode change is ready to accept when:

- proposal `0027` is still the next available number;
- no file introduces project-specific Kencleng paths, commands, or domain policy;
- every Codex file can point to an existing Harscode rule/source for the behavior it translates;
- no new lifecycle phase, risk tier, approval threshold, or testing obligation is invented;
- the Codex README carries Harscode's effective-as-of / last-verified / re-verify disclaimer;
- `AGENTS.md`, skills, session boundaries, and sandbox controls have non-overlapping source-of-truth roles;
- the files remain useful even if Codex native subagent mechanics change, because the portable fallback is Harscode's existing session boundary.

## Deferred intentionally

- project-specific `.codex/skills/` generation;
- a Codex frontend/backend track split;
- MCP/plugin recommendations;
- model/reasoning-effort routing thresholds;
- CI configuration;
- automatic synchronization scripts between Harscode and generated project skills.

Those should be added only after real Codex dogfood produces evidence that a generic translation is missing.
