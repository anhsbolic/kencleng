Read /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/2-3-techplan-decomposition-prompt.md in
full first — "When To Use This," "What This Must Not Do," and "Agent
Workflow" (steps 0-4) define what you're allowed to do here and what
you must not do (no compression, no reinterpretation, contract section
and any Test Focus Pointer table stay with the original techplan, not
redistributed).

Response style: full detail, no compression, when redistributing
content into task files. State the splitting axis you choose and why,
explicitly — don't apply one silently
(/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/README.md § Response Style By Phase).

The techplan under consideration is at /home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/2-techplan/techplan.md
— its contract must already be locked (not Draft). Read it in full
before doing anything else.

Step 0 — Gate question (mandatory, answer explicitly before proceeding):
Is it worth it to decompose this techplan? Answer no and stop here,
reporting a one-line reason, if it's small/linear enough for an
execution agent to run through start to finish without losing focus,
or if splitting it would mostly produce trivial single-task
boundaries. Do not proceed to Step 1 on autopilot just because this
prompt was invoked.

If yes: read the full techplan (Step 1), choose one splitting axis
from "Agent Workflow" § 2 — defaulting to dependency/sequence chain
when it's ambiguous — and state your choice and rationale before
redistributing content (Step 3). Each task file must remain
full-detail for its scope; nothing gets shortened or summarized in the
split, only relocated. Include a back-reference to the originating
techplan in every task file.

Generate the manifest last (Step 4) — task list, splitting axis +
rationale, dependency graph (or an explicit "no hard dependency"
marker), back-reference to the techplan, and which model to route each
task to (see /home/anhar-solehudin/kencleng-workspace/harscode-workspace/best-practices/model-routing.md).

Write every task file plus the manifest to
/home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/2-techplan/tasks/. Report the gate-question answer and,
if you proceeded, the splitting axis chosen — I'll review the split
before any task is executed.