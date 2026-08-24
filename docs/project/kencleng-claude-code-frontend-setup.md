# Kencleng Frontend — Claude Code Setup

> File: `docs/project/kencleng-claude-code-frontend-setup.md`
> Status: **Agreed** — Tier 1 only (slash commands + skills); Tier 2
> (subagents, protected-file hook) deliberately not adopted, see §4.
> Last updated: 2026-08-24

## 1. What this is

`frontend/` uses Claude Code as its agent harness, following the generic
pattern defined in `harscode-workspace/harness-optimization/claude-code/
frontend/`. This file documents the actual instance running in this
repo — what's installed, what it depends on, and how to operate it. The
harness-optimization files themselves stay project-agnostic by design;
this file is the project-specific counterpart, same split as `AGENTS.md`
(this repo) vs `best-practices/` (the generic workspace).

## 2. Hard dependency on `harscode-workspace`

Every command and skill in `frontend/.claude/` resolves content from
`harscode-workspace` via relative path or `@import` — **none of it is a
standalone copy.** Layout assumed (matches `frontend/opencode.jsonc`):

```
kencleng-workspace/
├── kencleng/              (this repo)
└── harscode-workspace/    (sibling repo, required)
```

If `harscode-workspace` isn't present at `../../harscode-workspace`
relative to `frontend/`, every command and skill breaks — commands
reference guideline files by path, skills `@import` best-practices
files directly. This is intentional (single source of truth — a
`harscode-workspace` proposal getting accepted updates behavior here
with zero edits to this repo) but means **this repo cannot run its
Claude Code setup standalone.** Anyone forking `kencleng` without also
cloning `harscode-workspace` needs to either clone it as a sibling or
adjust every path in `frontend/.claude/`.

## 3. What's installed

```
frontend/
├── AGENTS.md                          # existing file, + token-optimization
│                                       # section appended (§5)
├── .claude/
│   ├── commands/
│   │   ├── explore.md
│   │   ├── techplan.md
│   │   ├── techplan-review.md
│   │   ├── code-review.md
│   │   ├── test.md
│   │   └── pr.md
│   └── skills/
│       ├── accessibility-fundamentals/SKILL.md
│       ├── server-client-component-boundary/SKILL.md
│       ├── form-validation-boundary/SKILL.md
│       ├── data-fetching-conventions/SKILL.md
│       ├── component-test-mocking-discipline/SKILL.md
│       ├── api-client-centralization/SKILL.md
│       ├── loading-empty-error-state-conventions/SKILL.md
│       ├── ai-prototype-to-production-translation/SKILL.md
│       └── testing-automation-boundary/SKILL.md
```

Each command's body references its `workflow/<phase>/` guideline files
in `harscode-workspace` by path — never inlines their content. Each
skill's body is a single `@../../../../../harscode-workspace/
best-practices/react/<file>.md` import, not a copy — a `best-practices/`
proposal getting accepted is picked up automatically next session, no
regeneration needed here.

Verified working (2026-08-24): `/explore` correctly loads
`workflow/1-exploration/` guidance and runs the 3-stage process; the
`accessibility-fundamentals` skill auto-triggers on relevant prompts and
its `@import` resolves real file content (confirmed via specific detail
in the response matching the source file, not generic knowledge).

## 4. What's deliberately NOT installed

`harness-optimization/claude-code/frontend/subagents.md` and
`hooks-protected-files.md` (Tier 2) are **not** instanced here. This is
a decision, not a gap:

- **Protected-file hook**: skipped in favor of relying on Claude Code's
  default per-edit confirmation plus manual review. Accepted trade-off:
  this workspace's own `techplan/retro.md` records that manual
  vigilance has slipped even from the workspace owner directly (proposal
  0003's driving-story pass 3) — noted here so the trade-off is explicit,
  not because manual review is expected to be perfect.
- **Subagents**: skipped along with the hook, since the tool-restriction
  half of subagents serves the same access-control goal manual review
  already covers here. The model-tiering half of subagents (see §5) is
  handled manually instead — no per-phase model is auto-selected.

If this changes later, `harscode-workspace`'s generic pattern is already
written and doesn't need to be redrafted — only instanced.

## 5. Token-optimization behavior

`frontend/AGENTS.md` carries the enforcement block from
`harscode-workspace/harness-optimization/claude-code/token-
optimization.md`: terse by default, full completeness for
human-facing/self-check-guarded content (summaries, risk notes, PR
descriptions). Confirmed working in the `/explore` + skill test above —
concise on a routine question, without losing the checklist's actual
substance.

## 6. Model + effort per command

No subagent auto-pins a model per phase (§4), so this is manual: switch
before running a command via `/model` (arrow keys adjust effort) or
`/effort`. Anthropic's own guidance is to leave effort at default and
not fiddle per-task — the two exceptions below are deliberate departures
from that default, each backed by a specific reason, not general
caution.

| Command | Model | Effort | Why |
|---|---|---|---|
| `/explore` | Sonnet 5 | default | Routine investigation. Escalate only if Claude visibly misses something (skips an area, misapplies a sniffing lens) — that's the concrete-reason bar Anthropic's guidance sets for raising effort |
| `/techplan` | Sonnet 5 for Simple/Medium-tier stories; **Opus 4.8 for Complex tier** (per `best-practices/model-routing.md`'s tier definition — ≥15 rules, cross-service, breaking-change, or auth/payment/PII) | **high, deliberately** | Techplan synthesis is the one phase with a documented, repeated failure history (R18 dropping, severity leaking) that recurred across different models and even a manual pass by the workspace owner — this is exactly the kind of concrete history that justifies raising effort preemptively rather than waiting for a visible miss |
| `/techplan-review` | Sonnet 5 | default | Review pass, less generative than synthesis |
| `/code-review` | Sonnet 5 default; **Opus 4.8 for the Safety pass specifically** on any file matching `best-practices/index.md`'s Security Concern Map | high for the Safety pass only, default for Quality/Consistency | Mirrors this workspace's own dual-model-adversarial principle at Complex tier (`model-routing.md`) — applied here as a model/effort escalation on the highest-stakes pass only, not the whole review |
| `/test` | Sonnet 5 | default | Mechanical, follows `checklist.md` |
| `/pr` | Sonnet 5 | low | Diff-based, mechanical description drafting — routine work is exactly what Anthropic's guidance says to run at lower effort/smaller model |

Practical note: switching model/effort per command adds friction Claude
Code doesn't remove for you (no subagent to do it automatically, per
§4). Default the session to Sonnet 5 at default effort and only bump
for `/techplan` (Complex tier) and `/code-review`'s Safety pass — don't
re-litigate the other rows per task.