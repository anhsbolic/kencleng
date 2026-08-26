# Manifest — Code-review follow-up tasks

> Generated   : 2026-08-26
> Snapshot    : generation-time index, NOT a progress tracker (status lives in PRs/`tasks.md`)
> Source      : `.local-agents/works/account/03-login-session-management/4-code-review/report.md`
> Back-ref    : `.local-agents/works/account/03-login-session-management/2-plan/techplan.md` (Approved) · `3-build/report.md`

## Task files

| File | Title | Priority |
|---|---|---|
| `task-01-writeattempt-failopen-vs-failclosed.md` | Reconcile `writeAttempt` fail-open doc vs fail-closed code + add test | **Blocking** |
| `task-02-collapse-duplicated-ttl-constant.md` | Collapse duplicated access-token TTL constant | Optional |
| `task-03-logout-500-doc-exception.md` | Document logout 500-on-infra exception or align contract | Optional |

## Splitting axis

**Risk/blast-radius** — the blocking task (Q1) touches login-path error
semantics (security-relevant: fail-open vs fail-closed on an audit table).
The two optional tasks are low-risk cleanups (constant dedup, doc wording)
that can land independently or be folded into the Tier 0 paired-pass commit.

## Dependency graph

```
task-01 (blocking) ─────► must land before any commit
task-02 (optional) ─────► independent (can fold into Tier 0 paired commit)
task-03 (optional) ─────► independent (doc-only)
```

No hard dependencies between the three. task-01 is the only one that blocks
the commit gate; task-02 and task-03 can be deferred or skipped at Anhar's
discretion without correctness impact.

## Model routing

Per `best-practices/model-routing.md` — this is a follow-up to a Complex-tier
slice, but the tasks themselves are small surgical fixes (not multi-step
design work):

| Task | Build model | Rationale |
|---|---|---|
| 01 fail-open/closed | GLM 5.2 (max) | Security-relevant error-path change; needs careful reasoning about the fail-open trade-off + a new test case |
| 02 TTL dedup | DeepSeek V4 Pro | Mechanical constant consolidation; precision, no judgment |
| 03 logout doc | DeepSeek V4 Pro | Doc-only wording; precision |

Downstream reminder: the Tier 0 paired rewrite pass (techplan Resolved #13)
is a **human** task, not model-routed — it covers `platform/auth/token.go`,
`repository_db.go` rotation methods, and `login.go` reuse/race-loser branch
regardless of which model executes these follow-up tasks.

## ⚠️ Human decision required before task-01 execution

task-01 implements **fail-open** (the recommended option, matching the
existing `writeAttempt` doc and build-report deviation #5). The alternative
**fail-closed** (keep the current code, fix the doc) is documented inside
task-01 as option (b). If Anhar prefers fail-closed, task-01 must be
adjusted before execution — the test shape differs between the two.

## Cross-references

- Review report: `.local-agents/works/account/03-login-session-management/4-code-review/report.md`
- Contract techplan: `.local-agents/works/account/03-login-session-management/2-plan/techplan.md`
- Build report: `.local-agents/works/account/03-login-session-management/3-build/report.md`
- Review checklist: `harscode-workspace/workflow/4-code-review/checklist.md`
