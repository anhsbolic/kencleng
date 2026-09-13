# Manifest — Harscode 0027 Codex Harness Translation

Prepared against Harscode main `ee8fba26e9d378220fd2c6facc9e80306458a2c2`.

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
2. Copy `proposals/0027-codex-harness-translation.md` first with status `Proposed`.
3. Review the proposal as the human authority for the protected `harness-optimization/` tree.
4. If accepted, change the proposal status to `Accepted` and copy the six `harness-optimization/codex/` files.
5. Review the resulting Harscode diff for project-specific leakage or duplicated policy.
6. Commit/push Harscode using your normal process.
7. Once Harscode is merged/settled, remove this `harscode-handoff/0027-codex-harness-translation/` staging directory from Kencleng.

## Expected validation

A repository search after copying should show:

- exactly one proposal numbered `0027`;
- a new `harness-optimization/codex/` sibling next to `claude-code/`;
- no `Kencleng`, `npm run test:browser`, or target-project-specific path inside the Harscode Codex files;
- no new `CODEX.md` requirement;
- no edits to existing workflow/best-practices policy files.
