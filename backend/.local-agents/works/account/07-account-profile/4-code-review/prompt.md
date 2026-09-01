You are reviewing a code change. Run four passes, in this exact order,
against the diff below. Do not skip a pass because an earlier pass
found nothing — each pass looks for a different class of problem, and
a clean Safety pass says nothing about Consistency.

Guidance folder for this phase: /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/4-code-review
— guidelines.md, checklist.md, examples.md referenced below resolve
relative to this. best-practices/index.md resolves relative to
/home/anhar-solehudin/kencleng-workspace/harscode-workspace.

Response style: findings must be complete per item (finding, location,
why it matters, suggested fix) but not padded — "No findings" is a
valid and preferred answer over invented non-issues
(/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/README.md § Response Style By Phase).

Guidance for each pass — read before judging, don't rely on general
knowledge alone:
- Pass 1 (Safety): {file:/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/4-code-review/guidelines.md#1-safety-review}
- Pass 2 (Quality): {file:/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/4-code-review/guidelines.md#2-quality-review}
- Pass 3 (Stack-Specific Best Practices): {file:/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/4-code-review/guidelines.md#3-stack-specific-best-practices-review}
  — match against {file:/home/anhar-solehudin/kencleng-workspace/harscode-workspace/best-practices/index.md}, open only the files
  that match this diff's technology, apply their checklists.
- Pass 4 (Consistency): {file:/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/4-code-review/guidelines.md#4-consistency-check}
  — read [TARGET REPO CONVENTION FILE PATH] first, do not assume a
  pattern from a different project applies here.

Full checklist (all four passes): {file:/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/4-code-review/checklist.md}
Known recurring finding patterns worth specifically hunting for:
{file:/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/4-code-review/examples.md}

Diff to review:
/home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/3-build/report.md

Target repo convention file:
/home/anhar-solehudin/kencleng-workspace/kencleng/backend/README.md

---

Write your output below to
/home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/4-code-review/review-findings-<n>.md (see root README.md §
Task Working Directory Structure) — increment <n> per review round. If
any finding requires a code change, also write
/home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/4-code-review/patch-plan-<n>.md listing what needs to
change — the patch itself gets executed and reported in
/home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/3-build/, not here.

Output format — one section per pass, in the same order as run:

## 1. Safety
[Finding | Location | Why it matters | Suggested fix] per item, or
"No findings" if genuinely clean — do not pad with non-issues to look
thorough.

## 2. Quality
(same format)

## 3. Stack-Specific Best Practices
(same format — cite which best-practices file each finding came from,
e.g. "kafka/consumer-and-offset-management.md". If no keyword in
best-practices/index.md matched this diff's technology, say so
explicitly instead of skipping the section silently.)

## 4. Consistency
(same format — cite which existing convention was violated, and where
in the target repo that convention is established.)

## Verdict
One of: Approve / Approve with minor comments / Request changes.
If "Request changes," list which findings are blocking vs. optional.