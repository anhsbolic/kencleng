Run an independent four-pass review against the CURRENT diff/scope, in order:
1. Safety
2. Quality
3. Stack-specific best practices
4. Consistency with target-repo authority

Read:
- /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/4-code-review/guidelines.md
- /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/4-code-review/checklist.md
- /home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/2-techplan/techplan.md
- current task file only if this diff implements one decomposed task
- target repo's applicable AGENTS/README/CONTRIBUTING/convention source

Do not skip a pass because an earlier pass is clean; each pass asks a different
question. Do not invent findings merely to make the review look thorough.

Use /home/anhar-solehudin/kencleng-workspace/harscode-workspace/best-practices/AGENTS.md as routing authority for
Pass 3: search/scan only clue/security entries relevant to technologies and
concerns in this diff, then open the matching best-practice files. Do not absorb
the full index or browse the whole tree by default.

Do not load raw Exploration logs. The Techplan is the reviewed execution
contract; if material intent/evidence is missing from it, report Techplan drift
rather than rebuilding product intent from history.

For each finding state:
- location;
- problem;
- why it matters;
- suggested resolution;
- blocking vs non-blocking.

For Stack-Specific findings, cite the matching best-practice file. For
Consistency findings, cite the target-repo convention/precedent being violated.
If no best-practice trigger matches, say so briefly instead of silently skipping
Pass 3.

If Safety reveals a specialized concurrency/perf/security area absent from the
Techplan's Test Focus Pointer, report that separately as Techplan drift.

Do NOT edit production code in this session. When code changes are needed,
write a patch plan for Build.

Write:
- /home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/4-code-review/review-findings-<n>.md
- /home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/4-code-review/patch-plan-<n>.md only when code changes are required

Increment <n> per review round; do not overwrite prior review evidence.

Output sections:

## 1. Safety
[findings, or "No findings"]

## 2. Quality
[findings, or "No findings"]

## 3. Stack-Specific Best Practices
[findings + cited best-practice source, or explicit no-match/no-finding result]

## 4. Consistency
[findings + cited target-repo convention/precedent, or "No findings"]

## Verdict
Approve | Approve with minor comments | Request changes

If Request changes, identify which findings are blocking. Minor/non-blocking
comments must not be promoted into a patch loop merely to make the artifact
look cleaner.

## Phase handoff
- Completed: four-pass review + verdict
- Artifacts: <findings path; patch-plan path if any>
- Open / blocked: <blocking findings or none>
- Recommended next step: Testing if approved; otherwise Build/Patch
- Session recommendation: FRESH for Testing; BUILD authority for patches
- Context pointers: Techplan + specific findings/patch plan + diff anchors only