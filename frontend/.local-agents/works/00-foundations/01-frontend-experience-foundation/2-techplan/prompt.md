You are synthesizing the execution-grade Techplan for this task.

Read these Techplan authorities in full:
- /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/2-techplan/template.md
- /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/2-techplan/rules.md
- /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/2-techplan/guardrails.md

Use /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/2-techplan/guidelines.md only when
you need deeper process clarification; this canonical prompt already owns
the normal execution sequence. Do not default-load examples.md, retro.md,
techplan-example.md, report-template.md, or diagram-guidelines.md. Open them
only when their documented trigger applies.

Output quality: complete, unambiguous, execution-grade, non-redundant.
A fresh Build agent must be able to execute without inventing a material
product/domain, authority/security, architecture/ownership, interface/data,
risk, or verification decision. Do not pad the Techplan with duplicated
source prose, historical narrative, generic framework knowledge, or exact
mechanical detail Build can safely derive from current code.

EXPLORATION COVERAGE
- Fresh/compacted Techplan session: enumerate and read every durable file in
  /home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/1-exploration/logs/ once.
- Same healthy session that completed Exploration: enumerate the durable files,
  reuse evidence still active/unchanged, and open anything not already covered
  or whose exact wording is material. Do not mechanically reread unchanged
  evidence merely for ceremony.
- Correctness must remain reconstructable from durable artifacts. Do not place
  a decision in the Techplan if it exists only in remembered chat context.

SYNTHESIS
- Classify evidence by function per rules.md §1.
- Reconcile overlaps/conflicts per rules.md §2. Genuine contradictions become
  Open Items; do not silently choose the convenient source.
- Evaluate any independently operable migration/script/cron/runbook concern per
  rules.md §3 rather than forcing it into the feature plan.
- Preserve material rejected alternatives in the Decision Log with enough
  rationale that a later agent does not re-litigate settled choices.

PROJECT / PORTABLE AUTHORITY
Read the target repo's applicable AGENTS/README/spec/convention sources for any
project-specific contract or implementation assumption. Paths, symbols,
signatures, schema/API facts, UI/design requirements, and current behavior must
come from durable sources or direct live-code/spec checks.

For stack/security concerns that materially affect this plan, route through
/home/anhar-solehudin/kencleng-workspace/harscode-workspace/best-practices/AGENTS.md: search/scan only matching
clue/index entries, then open the matching best-practice files. Do not read the
whole index or best-practices tree by default. If Exploration already cites a
matching best-practice, verify the current authoritative file rather than
copying the Exploration paraphrase as policy.

IMPLEMENTATION DETAIL
In Implementation Details, record code anchors as path + symbol/section + why
relevant + intended change/precedent. Prefer anchors over copied code. Recheck
non-obvious current facts against live code/spec before presenting them as
executable instructions.

BEFORE FINALIZING
- every Rules & Validation rule ID has Testing Checklist verification coverage;
- every surviving concurrency/perf/security-sensitive Exploration risk has a
  Test Focus Pointer row with its exact Exploration evidence anchor;
- a scoped-out sensitive risk is N/A with a reason/Decision Log pointer, not
  silently absent;
- unresolved material uncertainty is in Open Items rather than guessed;
- an existing Approved/Implemented Techplan has not been materially changed
  without the guardrail/human gate required for a contract revision.

There is NO embedded Summary in techplan.md. Do not generate the human report
during Draft/In Review; report-techplan.md is generated separately only after
Approval.

Write the result to /home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/2-techplan/techplan.md using template.md.
Treat the written Techplan as a Draft/In-Review checkpoint until the human gate
approves it; do not automatically continue into Build merely because synthesis
completed.

At completion, report:

## Phase handoff
- Completed: Techplan synthesized and self-checked (or state what remains)
- Artifacts: /home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/2-techplan/techplan.md
- Open / blocked: <material unresolved items or none>
- Recommended next step: human Techplan gate; independent review only if its Complex gate applies
- Session recommendation: FRESH for Build after Techplan approval; if review/decomposition remains, CONTINUE only while planning context stays focused
- Context pointers: techplan path + only source anchors needed for unresolved follow-up