You are implementing against an approved techplan, one iteration of
the build/patch loop at a time. Don't revisit architectural decisions
already made in the techplan — if something genuinely doesn't hold as
written, stop and report it instead of silently working around it.

Guidance folder for this phase: /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/3-build —
guidelines.md and checklist.md referenced below resolve relative to
this.

Response style: keep your own narration minimal during this
iteration — do the work efficiently, don't narrate each step. The
output format below is the one place to be complete
(/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/README.md § Response Style By Phase).

Build target: /home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/2-techplan/tasks/<task-file>.md if
decomposition ran, otherwise /home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/2-techplan/techplan.md
directly. One task per iteration when tasks exist; the whole techplan
in one go when they don't — don't ask for a separate scope on top of
that, the right-sized unit of work was already decided at
decomposition's Step 0 gate (or by its absence).

Test scope for this loop — fixed, not negotiable per task:
{file:/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/3-build/guidelines.md#default-test-scope-always-regardless-of-techplan-content}

Full checklist: {file:/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/3-build/checklist.md}

---

Write your output below to /home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/3-build/report.md for the
initial build, or /home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/3-build/patch-report-<n>.md if this
iteration is executing a patch plan from code-review or testing (see
root README.md § Task Working Directory Structure) — increment <n>
per patch, don't overwrite a previous one.

Output format:

## What changed
[files touched, one line each]

## Tests run
[test name/pattern → category (unit/mocked/API-contract) → result]
Confirm explicitly: no `-race`, perf/load, or security-class test was
run in this iteration.

## Contract check
- [ ] This iteration satisfies its build target in full (the task
      file's scope, or techplan § 4 in full if there's no task file)
- [ ] No contract assumption broke — or, if one did, flagged below
      instead of worked around

## Flagged for techplan/testing review (if any)
[Concurrency/perf/security concern noticed but not tested here, or a
contract assumption that didn't hold — one line each]