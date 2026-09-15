You are exploring this task in Kencleng — Go backend + Next.js frontend. This phase has three
stages; do not merge solutioning into gap analysis.

Task: /home/anhar-solehudin/kencleng-workspace/kencleng/docs/spec/0-foundations/features/01-frontend-experience-foundation.md
Area: "not sure yet"
Working directory: /home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation

If /home/anhar-solehudin/kencleng-workspace/kencleng/docs/spec/0-foundations/features/01-frontend-experience-foundation.md points to a task-specific requirement/spec document, read that
source before Stage 1. If it points to a large multi-topic source, read the
authoritative section(s) governing this task plus referenced dependencies;
do not infer requirements from a title, filename, or stale summary.

Guidance for Exploration:
- /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/1-exploration/guidelines.md
- /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/1-exploration/sniffing-checklist.md

Do not load Exploration examples unless you need calibration for an ambiguous
output shape.

Response style: Stage 1 is concise because it is only a routing/checkpoint
step. Stages 2 and 3 preserve concrete evidence and material rationale because
Techplan synthesis must be able to reconstruct the work from durable artifacts.

STAGE 1 — PLAN ANNOUNCEMENT (hard stop)
- Read the target repo's applicable AGENTS/README/convention authority only
  far enough to identify the relevant areas and any non-negotiable boundaries.
- State your 1–2 sentence understanding of the task so a requirement misread
  can be corrected early.
- Identify the areas you intend to explore, their order, and why that order.
- Do not inspect implementation deeply and do not propose solutions yet.
- STOP for human confirmation before Stage 2.

STAGE 2 — GAP ANALYSIS (after confirmation)
For each area, finish that area before moving on:
- Current state — concrete live behavior, relevant files/symbols/contracts.
- Requirement — the exact relevant expectation; cite the governing source
  section/anchor when the task came from a document.
- Gap — the specific difference between current state and requirement.
- Sniffing — run the five lenses in sniffing-checklist.md for this area.
- Code anchors — record path + symbol/section + why it matters for later
  verification/Build; copied implementation text is not a new authority.

When Stage 2 surfaces a concrete stack/security concern that matches portable
best-practice guidance, use
/home/anhar-solehudin/kencleng-workspace/harscode-workspace/best-practices/AGENTS.md as the routing rule: search
or scan only the relevant clue/index entries and open only matching authority
files. Do not browse the whole best-practices tree for completeness theater.

Do not develop solutions or compare alternatives in Stage 2. A one-line
observation may be parked without turning it into a decision.

Write durable Stage-2 evidence under /home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/1-exploration/logs/. Report
progress after each area so the human can redirect, but do not require a new
approval between every area unless asked.

STAGE 3 — SOLUTIONING (only after human confirms Stage 2)
Compare viable options/trade-offs and record the chosen direction. Preserve
material rejected alternatives and the rationale/consequence that prevents a
later agent from reopening a settled choice accidentally.

Keep source evidence and decisions durable in
/home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/1-exploration/logs/; do not rely on chat memory alone. The raw-doc
shape should follow the evidence rather than a forced per-file template.

At Stage-3 completion, output:

## Phase handoff
- Completed: <areas explored + solutioning state>
- Artifacts: <durable Exploration paths>
- Open / blocked: <material unresolved items or none>
- Recommended next step: Techplan synthesis
- Session recommendation: CONTINUE | FRESH — based on continuation fitness
- Context pointers: <source/spec paths + code anchors likely needed next>

Recommend FRESH when the session was compacted/reset, accumulated substantial
dead ends/unrelated investigation, had major human redirection, or cannot name
the current durable authorities cleanly. Otherwise CONTINUE is valid; do not
use a fuzzy task-size label as the deciding rule.