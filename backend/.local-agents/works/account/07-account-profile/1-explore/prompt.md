read docs inside /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/exploration before doing exploration 

now you are exploring a task for feature 07-account-profile, read /home/anhar-solehudin/kencleng-workspace/kencleng/docs/spec/README.md then read /home/anhar-solehudin/kencleng-workspace/kencleng/docs/spec/domains/account/features/07-account-profile.md.

STAGE 1 — Plan Announcement (do this first, then STOP and wait for me):

Read this repo's own convention file (AGENTS.md / README / whatever
exists) just enough to identify which areas of the codebase are
relevant to this task. Then tell me:
- Which areas you intend to explore
- In what order, and why that order
- Do NOT read the actual implementation files yet, and do NOT propose
  any solution yet. Wait for my go-ahead before Stage 2.

STAGE 2 — Gap Analysis (after I confirm the plan):

For each area, one at a time, fully before moving to the next:
- Current state: what exists today, concretely — actual function/file
  names, actual behavior, not a vague description
- Requirement: what the new requirement expects here
- Gap: the specific difference
- Sniffing: run the five lenses in sniffing-checklist.md (risk, edge
  cases, miscontext, misleading signals, inconsistency) on this area
- Do NOT propose solutions or compare options here. A bare one-line
  observation is fine if something occurs to you, but don't develop it.
- Report after each area so I can redirect if needed, then continue to
  the next area without waiting for explicit approval each time.

STAGE 3 — Solutioning (only after I've reviewed Stage 2's output):

This is where trade-offs, options, and rationale get written. Don't
start this until I explicitly confirm the gap analysis is accurate and
tell you to proceed.

Write whatever form of raw doc best captures what's found at each
stage inside /home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/1-explore/logs — the shape should follow the content, not a preset template.