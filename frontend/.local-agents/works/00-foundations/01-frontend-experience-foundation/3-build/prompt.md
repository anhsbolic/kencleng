You are executing the Approved contract through Build/Patch authority.

Read:
- /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/3-build/guidelines.md
- /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/3-build/checklist.md
- /home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/2-techplan/techplan.md as the authoritative spine
- the current task file only when decomposition exists
- the specific patch plan only when this is a Review/Testing re-entry
- target-repo instructions/commands applicable to the files and verification
  you will touch

Do not load raw Exploration logs by default. If the Techplan explicitly points
to unresolved evidence, open that exact source only.

Before editing, reopen current live code/spec at the Techplan/task code anchors.
Earlier Exploration/Techplan descriptions are coordinates/evidence, not a frozen
copy of implementation reality.

Do not re-explore settled product/domain decisions. If live code contradicts a
material Techplan assumption (behavior, authority/security, architecture/
ownership, interface/data contract, risk/verification), stop and report it
instead of silently redesigning.

BUILD TARGET
When task files exist, execute the current task with the parent spine; do not
read unrelated sibling tasks unless a declared hard dependency requires one.
When decomposition did not run, execute the Approved Techplan as the build
target. Do not invent a second ad-hoc scope boundary just because implementation
started.

PATCH RE-ENTRY
Production fixes remain Build work even when the finding came from Review or
Testing. Re-ground on parent Techplan + current task (if any) + the specific
patch plan + relevant live code/diff. Do not import the whole reviewer/tester
conversation as hidden authority.

Remember which phase requested the patch. After the narrow patch, return to
that requesting phase unless the patch materially invalidates the Techplan or
broadens scope enough to justify a different route. Do not automatically create
a full new review loop merely because a patch occurred.

VERIFICATION
Use the Techplan Testing Checklist's primary-owner/rationale fields plus
workflow/3-build/guidelines.md.

Run focused Build-loop verification needed to make the edit credible. If you
add/change an automated test, run it enough to prove the authored test is
executable and exercises the intended state. Do not replay broad Testing-owned
suites merely for extra confidence unless target-repo authority requires them
at this point, the current change materially affects that risk, or the broader
suite is the only credible evidence.

For a narrow patch, rerun affected verification first; broaden only when the
patch changes scope/risk or wider regression evidence is necessary.

Do not pull heavyweight race/perf/security-class verification into this tight
loop merely for extra confidence; those belong to independent Testing when
triggered.

Process narration is terse; do the work. Write:
- initial build: /home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/3-build/report.md
- patch round: /home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/3-build/patch-report-<n>.md

Increment <n> for each patch round and never overwrite earlier patch reports.

Every durable Build/Patch report must start with compact provenance using only
known/exposed values:
- Phase: Build or Build/Patch
- Author
- Created/Updated
- Model / Reasoning / Session when exposed and safe to persist
- Target revision / Workflow revision when known

Do not invent missing metadata or persist account/credential identifiers. Git
history is the default version history.

Report format:

## What changed
[file/symbol or area → concise behavior/contract-relevant change]

## Tests run
[test/command or pattern → verification category → result + why this run belonged in Build when non-obvious]

## Verification scope confirmation
Confirm explicitly: no race/concurrency, performance/load, or security-class
test was executed in this Build iteration. If any was run, list it here and
flag the scope deviation instead of silently treating it as ordinary Build
verification.

Also note any broad Testing-owned suite executed in Build and why it was
necessary for this iteration.

## Contract check
- [ ] Current build target satisfied in full
- [ ] Live-code re-grounding did not invalidate a material contract assumption

## Deferred / not tested here
[verification deliberately left for independent Testing/Human, with reason; "none" if none]

## Flagged for Techplan / Testing
[material assumption break or specialized concern; "none" if none]

## Phase handoff
- Completed: <build target/patch completed or what remains>
- Artifacts: <report path>
- Human decision: <decision needed now or "none">
- Open / deferred: <material blocker/deferred item or "none">
- Recommended next step: Code Review after initial build; return to the requesting Review/Testing phase after a patch
- Session transition: <plain-language action + reason; e.g. start fresh Code Review/Testing for independence, continue this Build session for another focused iteration, or return to/restart Build/Patch as appropriate>
- Context pointers: parent Techplan + current task/patch plan + changed files/tests only