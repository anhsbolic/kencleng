# Gitleaks — backend-only recheck

Date: 2026-10-11 Asia/Jakarta. Tool: Gitleaks v8.30.1 in the configured Podman Compose service.
Authorization: Human requested a repeat restricted to backend/. This is a separate tooling check, not a reopening or acceptance update of terminal RUN-C1-ENG-TEST-T2-001.

## Scope and commands

- History: `make -C backend gitleaks-history` invokes `git /repo --log-opts="--all --full-history -- backend/"`; only the backend path across available Git refs. This is not limited to current T2 or HEAD.
- Current files: `make -C backend gitleaks-worktree` invokes `dir /repo/backend`; includes current tracked/untracked and Git-ignored backend files, including the local `.env`.
- Read-only repository mount; secret redaction enabled; no suppressions or allowlists added; bounded 180-second scanner timeout and container auto-removal requested.

## Results

- History: scan completed, 36 commits / approximately 2.34 MB, **4 candidates**. All four occur in historical `*_test.go` paths under backend/. Those files do not exist in the current tree. Test-file location does not alone prove fixture/non-applicability. Sanitized rule/path/line/commit details: `history-findings.json`.
- Current directory: scan completed, approximately 50.34 KB, **1 candidate**, `backend/.env:12`, generic-api-key. Git ignored: True; Git tracked: False. This is local configuration, not evidence that the value is committed or exposed. No other match was emitted for current backend files by this scan.
- Both scanner exits are 1 (findings), corresponding Make exits 2. **Not a clean secret-scanning gate.** No runtime identity/provider/DB/C1 acceptance conclusion follows.

## Evidence and limits

Initial command records and full redacted stdout/stderr: `commands.json`, `gitleaks-history.*`, `gitleaks-worktree.*`. A justified second execution collected sanitized finding locations: `detail-commands.json`, `history-findings.json`, `worktree-findings.json` and detail command logs. Secret/Match/Author/Email fields are not retained in the finding summaries; temporary full redacted reports were removed after sanitization. No secret values were printed, changed or removed, and no Git history was rewritten.

Only `backend/Makefile` was edited to make the backend security gate call these scoped checks, with `make -k` so both execute even if the first finds candidates. Root Compose/Makefile defaults were not edited by this recheck and their repository-wide target is not used by the backend gate. No other verification suite was rerun.
