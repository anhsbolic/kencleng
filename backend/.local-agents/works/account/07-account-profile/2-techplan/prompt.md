Read the guidance folder at /home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/2-techplan in
this order: README.md, template.md, rules.md, guardrails.md,
guidelines.md, examples.md, retro.md. All bare file references below
(rules.md, guardrails.md, template.md) resolve relative to this same
folder. This defines how you should classify content and what the
output must look like.

Response style: full detail, execution-grade, no compression — this
techplan becomes the contract other agents and humans execute against
later (/home/anhar-solehudin/kencleng-workspace/harscode-workspace/workflow/README.md § Response Style By Phase).

Then read every file in /home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/1-exploration/logs/ — treat all of
them as raw material. Don't assume a fixed number or fixed names;
classify each piece of content by the function it serves (rules.md
§ 1), not by which file it came from.

Before filling in the Interface Contract section, read this repo's own
convention file (AGENTS.md / README / CONTRIBUTING — whatever exists
here) to know what's mandatory to cover — don't assume the conventions
from the guidance folder apply here.

If two or more raw docs cover the same ground with conflicting or
overlapping detail, follow rules.md § 2 (Dedup & Reconciliation) —
prefer the most specific and most recent version, and call out
anything that's a genuine conflict rather than picking silently.

If any sub-component has its own independent operational lifecycle
(one-time script, cron, separate rollback/cleanup) — evaluate per
rules.md § 3 whether it belongs as a section here or as a separate
linked document.

Follow every guardrail in guardrails.md — in particular: don't invent
technical facts (mark anything uncertain as `TBD — verify`), don't
overwrite an existing Approved/Implemented techplan's contract sections
silently, and STOP and ask me if you find a breaking change or data
risk that isn't already explicit in the raw docs.

Stack-specific risk lens: before finalizing sniffing (risk, edge cases, miscontext,
misleading signals, inconsistency), read /home/anhar-solehudin/kencleng-workspace/harscode-workspace/best-practices/index.md.
Match the trigger keywords in the index table against the area(s)/technology(ies)
touched by this ticket (Go, PostgreSQL, GraphQL, REST API, Kafka, Pub/Sub, Redis).
Open ONLY the matching file(s) — do not scan the entire best-practices/ folder.
Apply the checklist from each matching file as part of the risk lens.

When populating the Testing Checklist section (currently §12 — verify
the actual number against this template.md, don't assume), also
populate its Test Focus Pointer table (see rules.md's Test Focus
Pointer rule) by cross-referencing the raw exploration docs' Sniffing
Checklist Risk findings for this task. Only carry forward areas
genuinely about shared state, concurrency, an expensive primitive
under load, or a security-sensitive boundary — and don't silently drop
one that survived synthesis (see guardrails.md's rule on flagged risk
areas); mark it N/A with a one-line reason instead.

Write the result to /home/anhar-solehudin/kencleng-workspace/kencleng/backend/.local-agents/works/account/07-account-profile/2-techplan/techplan.md, following
template.md's structure exactly. At the end, list out any open items
or unresolved questions you carried forward instead of silently
deciding — I'll review those manually before this goes anywhere
further.