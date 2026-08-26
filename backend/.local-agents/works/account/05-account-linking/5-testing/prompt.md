You are running the testing phase for a story that has already been
through implementation and code review. Follow this process, in order:

Step 0 — Sweep, don't redo:
{/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/5-testing/guidelines.md#process} (item 0 specifically).
Read the implementation report below. For every rule/scenario it claims
is covered by a named test, spot-check it (run the existing test, don't
rewrite it). Then read its "what is not tested, and why" section (or
equivalent) — that is your priority list, not a fresh read of techplan
§ 4 from zero.

Step 1 — Coverage per techplan § 4:
{/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/5-testing/guidelines.md#process} (items 1-3). Test every
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

Full checklist: {/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/5-testing/checklist.md}
Known recurring bug patterns worth specifically hunting for:
{/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/5-testing/examples.md}

Techplan:
/home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/05-account-linking/2-plan/techplan.md

Latest implementation report (build or most recent patch/rebuild):
/home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/05-account-linking/4-patch/report.md

Real interface entry point(s):
Read /home/anhar-solehudin/kencleng-workspace/kencleng/backend/README.md

Target repo build/lint/test commands:
read /home/anhar-solehudin/kencleng-workspace/kencleng/backend/README.md

---

Output format:

## 0. Sweep Summary
- Confirmed (spot-checked, still passing): [rule/scenario → test name]
- Closed from report's own gap list: [gap → what was added/verified]
- Not carried over — required fresh testing: [what, and why the report
  didn't cover it]

## 1. Test Coverage
[Rule/Scenario | Category (happy/negative/edge/backward-compat) |
Real-interface test performed | Result] per item. Cite § 4 rule IDs
where applicable. If a § 4 scenario can't be exercised through the real
interface, flag it explicitly — don't skip silently.

## 2. Error Verification
[Error case | Expected category | Actual category | Message
actionable? | Propagation correct?] per case, or "No error paths
exercised beyond §1" if genuinely none apply.

## 3. Final Verification
- [ ] Target repo build/lint/test commands: pass/fail (paste output)
- [ ] Migration/schema version collision check: result
- [ ] Backward compatibility: explicitly verified, not assumed
- [ ] Fresh end-to-end techplan read: gaps/contradictions found, or none

## 4. New Bug Patterns
Only include entries that meet the threshold in
{/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/5-testing/guidelines.md#threshold-for-adding-to-examplesmd}
(a category of mistake, not a one-off). Otherwise state "No new pattern
— see /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/5-testing/examples.md for handling this ticket-specific
bug directly."

## Verdict
One of: Pass / Pass with flagged follow-ups / Fail — send back to
implementation.
If "Fail" or flagged follow-ups, list which findings are blocking vs.
optional, and whether they trace to a genuinely new gap or a
Step-0-confirmed area that regressed.