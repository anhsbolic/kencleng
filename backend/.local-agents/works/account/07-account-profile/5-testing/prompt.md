You are running the testing phase for a task that has already been
through implementation and code review. Follow this process, in order:

Guidance folder for this phase: /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/5-testing —
guidelines.md, checklist.md, examples.md referenced below resolve
relative to this.

Response style: keep the sweep/coverage work itself efficient — don't
narrate every step. The final report below must be as thorough as a
real testing report demonstrates is possible without an over-narrated
process (see /home/anhar-solehudin/kencleng-workspace/harscode-workspace/best-practices/go/examples/testing-concurrency.md
for a worked example) — that level of detail is the bar for this
report, not an exception (/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/README.md §
Response Style By Phase).

Step 0 — Sweep, don't redo:
{file:/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/5-testing/guidelines.md#process} (item 0 specifically).
Read the implementation report below. For every rule/scenario it claims
is covered by a named test, spot-check it (run the existing test, don't
rewrite it). Then read its "what is not tested, and why" section (or
equivalent) — that is your priority list, not a fresh read of techplan
§ 4 from zero.

Also as part of Step 0: read techplan § 12's Test Focus Pointer table.
For every row still marked relevant, cross-check the raw exploration
docs below for the concrete Sniffing Checklist Risk-lens finding behind
it, then build a Test Execution Plan (scope, tooling, threshold) for
that area — this is a distinct deliverable from the four-category
coverage in Step 1, and covers race/concurrency/perf/security-class
tests that Step 1 does not. If the pointer table is empty or missing
but you notice a genuinely concurrency/perf/security-sensitive area in
the exploration docs or code, flag it back as a possible techplan gap —
don't silently add the test yourself without noting the gap.

Step 1 — Coverage per techplan § 4:
{file:/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/5-testing/guidelines.md#process} (items 1-3). Test every
rule in the techplan's Rules & Validation section using the real
interface. For rules already spot-checked in Step 0 as genuinely
covered, don't re-derive a new test — note it as confirmed. Spend actual
effort only on: rules with no named test, rules whose named test you
could not confirm still passes, and the four categories below.

Cover all four categories, not just the happy path:
- Happy path
- Negative cases (missing fields, invalid input, dependency failures)
- Edge cases (empty/null/boundary values)
- Backward compatibility (old clients, existing data)

Verify error responses precisely — category, actionable message,
correct propagation through the app's error-handling layer, not just
"an error happened."

Full checklist: {file:/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/5-testing/checklist.md}
Known recurring bug patterns worth specifically hunting for:
{file:/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/5-testing/examples.md}

Techplan:
/home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/2-techplan/techplan.md

Raw exploration docs:
/home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/1-exploration/logs

Latest implementation report (build or most recent patch/rebuild):
/home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/3-build/report.md

Target repo build/lint/test commands:
/home/anhar-solehudin/kencleng-workspace/kencleng/backend/README.md