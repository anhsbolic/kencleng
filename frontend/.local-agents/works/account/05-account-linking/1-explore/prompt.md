Read docs inside /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/1-exploration before doing exploration.

You are exploring a task for feature 05-account-linking (frontend surface). Read:
1. /home/anhar-solehudin/kencleng-workspace/kencleng/docs/spec/README.md
2. /home/anhar-solehudin/kencleng-workspace/kencleng/docs/spec/1-account/features/05-account-linking.md
3. /home/anhar-solehudin/kencleng-workspace/kencleng/docs/ui-ux/page-map.md — check ALL rows relevant to this feature, not just the most obvious page; a feature can surface as a full page, a non-page flow (e.g. a link/redirect landing route), or both. Report every surface you find, not just the primary one.
4. /home/anhar-solehudin/kencleng-workspace/kencleng/docs/ui-ux/patterns.md — read whichever pattern(s) apply to the page(s) found in page-map.md above.
5. /home/anhar-solehudin/kencleng-workspace/kencleng/docs/ui-ux/prototype-reference.md — check which tier this task's page(s) fall under per the Tier 1/Tier 2 tables. Don't stop at the table: separately list the actual contents of /home/anhar-solehudin/kencleng-workspace/kencleng/docs/design-reference/ and check whether any file there plausibly corresponds to this task's page(s), even if the filename doesn't exactly match a route named in prototype-reference.md's table. Report what you actually find there rather than assuming the table's tier assignment is complete.
6. /home/anhar-solehudin/kencleng-workspace/kencleng/docs/ui-ux/design-guidelines.md

STAGE 1 — Plan Announcement (do this first, then STOP and wait for me):
Read this repo's own convention file (frontend/AGENTS.md, frontend/.agents/docs/README.md, and whatever else exists) just enough to identify which areas of the frontend codebase are relevant to this task — including which Shell (Public/Auth/Dashboard from Phase 0, or otherwise) this task's page falls under, and what that Shell currently contains. Then tell me:
- Which areas you intend to explore
- In what order, and why that order
- Do NOT read deeper implementation files yet beyond what's needed to answer the above, and do NOT propose any solution yet. Wait for my go-ahead before Stage 2.

STAGE 2 — Gap Analysis (after I confirm the plan):
For each area, one at a time, fully before moving to the next:
- Current state: what exists today, concretely — actual file/component names, actual behavior, not a vague description. Explicitly confirm whether this task's page(s) already exist as routes or not.
- Requirement: what the feature spec + page-map.md + patterns.md expect here
- Gap: the specific difference
- Page-consolidation check (per workflow §14 step 1): does this task's page(s) already exist from an earlier frontend task in this domain? Is there any page-map.md action for this page with no backing endpoint in this domain's tasks.md, or vice versa?
- Sniffing: run the five lenses in sniffing-checklist.md (risk, edge cases, miscontext, misleading signals, inconsistency) on this area
- Do NOT propose solutions or compare options here. A bare one-line observation is fine if something occurs to you, but don't develop it.
- Report after each area so I can redirect if needed, then continue to the next area without waiting for explicit approval each time.

STAGE 3 — Solutioning (only after I've reviewed Stage 2's output):
This is where trade-offs, options, and rationale get written, including how any non-page surface identified in Stage 1/2 gets implemented. Don't start this until I explicitly confirm the gap analysis is accurate and tell you to proceed.

Write whatever form of raw doc best captures what's found at each stage (from stage 1) inside /home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/account/05-account-linking/1-explore/logs — the shape should follow the content, not a preset template.