# Patch plan — configured DB harness Review F1

Phase: Code Review. Author/role: Codex / independent Reviewer. Created/Updated: 2026-10-13 Asia/Jakarta. Model / reasoning: `gpt-6-luna` / `medium` per Invocation. Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-PATCH-001`. Participant / Session: `PARTICIPANT-C1-ENG-T2-CONFIGURED-DB-HARNESS-PATCH-REVIEWER-001` / `SESSION-C1-ENG-T2-CONFIGURED-DB-HARNESS-PATCH-REVIEWER-001` (assignment identities). Target / workflow revisions: `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `651ce57f044bff23fba99082322cee6c066d25a6`.

## F1 — Do not pass fixture credentials to `go version`

- **Location:** `backend/scripts/auth_harness.py`, `main()`, `command(["go", "version"], env=env)`.
- **Required change:** Obtain Go version before adding fixture credentials to `env`, or construct a minimal/sanitized environment for the metadata command. Preserve fixture `DATABASE_URL`, provisioner DSN, and Run binding only for the intended Go test subprocess. Do not print either DSN.
- **Acceptance evidence:** Static inspection confirms no fixture DSN is passed to `go version`; the Go test child still receives the configured runtime URL and required Run binding. Any execution evidence belongs to a separately authorized Build/Testing scope.
- **Scope:** Keep the change within the authorized harness delta; no fixture or runtime action is authorized by this Review.
