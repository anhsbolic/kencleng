# Manifest — Harscode 0027 Codex Harness Translation

Prepared against Harscode main `ee8fba26e9d378220fd2c6facc9e80306458a2c2`.

Human review of this handoff completed on 2026-09-13. The staged target proposal is now marked `Accepted`.

## Target operations

All target files are **new** on the audited Harscode baseline.

```text
CREATE  proposals/0027-codex-harness-translation.md

CREATE  harness-optimization/codex/README.md
CREATE  harness-optimization/codex/instruction-loading.md
CREATE  harness-optimization/codex/skills.md
CREATE  harness-optimization/codex/session-boundaries.md
CREATE  harness-optimization/codex/permissions-and-sandbox.md
CREATE  harness-optimization/codex/token-optimization.md
```

No existing `workflow/`, `best-practices/`, or `harness-optimization/` root file needs modification for this proposal.

## Copy source

Each file under `full-files/` mirrors its target Harscode path exactly:

```text
full-files/proposals/0027-codex-harness-translation.md
→ proposals/0027-codex-harness-translation.md

full-files/harness-optimization/codex/*
→ harness-optimization/codex/*
```

## Apply order

1. Re-check Harscode `main` and confirm proposal number `0027` is still available.
2. Copy `full-files/proposals/0027-codex-harness-translation.md` to `proposals/0027-codex-harness-translation.md`.
3. Copy all six files under `full-files/harness-optimization/codex/` to `harness-optimization/codex/`.
4. Review the resulting Harscode diff for project-specific leakage or duplicated policy.
5. Commit/push Harscode using your normal process.
6. Once Harscode is merged/settled, remove this `harscode-handoff/0027-codex-harness-translation/` staging directory from Kencleng.

## Expected validation

A repository search after copying should show:

- exactly one proposal numbered `0027`;
- the proposal status is `Accepted`;
- a new `harness-optimization/codex/` sibling next to `claude-code/`;
- no `Kencleng`, `npm run test:browser`, or target-project-specific path inside the Harscode Codex files;
- no new `CODEX.md` requirement;
- no edits to existing workflow/best-practices policy files.
