# Configured DB harness patch — independent Code Review

Phase: Code Review. Author/role: Codex / independent Reviewer. Created/Updated: 2026-10-13 Asia/Jakarta. Model / reasoning: `gpt-6-luna` / `medium` per Invocation; runtime model metadata is not independently exposed. Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-CODEREVIEW-T2-CONFIGURED-DB-HARNESS-PATCH-001`. Participant / Session: `PARTICIPANT-C1-ENG-T2-CONFIGURED-DB-HARNESS-PATCH-REVIEWER-001` / `SESSION-C1-ENG-T2-CONFIGURED-DB-HARNESS-PATCH-REVIEWER-001` (assignment identities). Target / workflow revisions: `79ec63d57674d1e4976ff80b2dfee8c50bf17934` / `651ce57f044bff23fba99082322cee6c066d25a6`.

Review target was the exact three-file delta bound by this Run. Current SHA-256 values match the reported post-delta identities: `harness_test.go` `72c63063aa6709216d110d3f6a41cf81d981b510c9e583305b56d43399567ba8`; `auth_harness.py` `fa5ebc1442b18cce82d5b5b1270855baa661ad4fbadb82456ffbf3b488d8948a`; `README.md` `aa48bbeca8f8d186cc806d53769806da6684a1f3fac9cc80ea16de3b66fb0ac4`. Other dirty workspace content was excluded.

## 1. Safety

**F1 — Blocking — fixture credentials reach an unrelated tool subprocess.** `backend/scripts/auth_harness.py`, `main()`, at `command(["go", "version"], env=env)`: after populating `env` with both fixture DSNs and `C1_HARNESS_OWNED_RUN`, the runner passes the credential-bearing environment to `go version`, as well as the intended Go test child. This breaks the stated boundary that the generated restricted URL is child-test-only and exposes both DB credentials to a process that does not need them. Invoke `go version` with a sanitized environment (or obtain toolchain metadata before adding fixture credentials); keep fixture URLs only in the test child environment. The later Go test invocation correctly needs the configured runtime URL. The rest of the isolation, role separation, ownership-checked cleanup, and six-case selection is sound on static inspection.

## 2. Quality

No findings. The Python runner keeps process invocation as argument arrays, validates the required Run ID and suite before resource creation, and fails closed on ambiguous ownership or cleanup state. The Go helper closes pools through test cleanup and propagates the bounded context to pool setup. The README describes the configured-pool path and keeps simulated callback, `httptest`, fixture PostgreSQL, deployed configuration, and independent acceptance claims distinct.

## 3. Stack-Specific Best Practices

No findings. The routed Go integration-testing guidance calls for isolated reproducible state, explicit cleanup, and deterministic resource release; the runner and test helper provide those paths (`best-practices/go/integration-testing-setup.md`). The routed Go input-validation guidance calls for validating boundary input before passing it to external processes; the runner validates the Run ID/suite and uses shell-free argument arrays (`best-practices/go/input-validation-and-injection.md`). The applicable Go test integration trigger matched; no additional stack-specific concern requiring a separate best-practice file was identified.

## 4. Consistency

No findings. The patch follows root/backend instructions to preserve intentional failure handling and avoid leaking credentials, and follows `docs/project/c1-auth-verification-environment.md` by keeping the disposable integration fixture separate from application runtime configuration and real provider/browser evidence. README statements remain within the bounded configured-pool harness claim. Six selected cases and the expiry-under-lock oracle match the bound T2 task and Build report.

## Verification executed during Review

No tests, gates, scanners, fixtures, database operations, provider flows, or browser/runtime checks were run, as required by the Invocation. Static inspection only: read the bound inputs and current diff, confirmed the three current file hashes against the Build report, and inspected relevant source and guidance.

## Verdict

Request changes — resolve blocking F1 before this exact delta is approved. This review does not establish R14, B-T2-003, dependency acceptance, T2/T3 readiness, or C1 acceptance.

## Phase handoff

- **Outcome:** COMPLETED — four-pass static review; Request changes.
- **Result refs:** `evidence/review-findings-1.md`, `evidence/patch-plan-1.md`; reviewed files `backend/internal/platform/auth/harness_test.go`, `backend/scripts/auth_harness.py`, and `backend/README.md` at the hashes recorded above.
- **Findings:** Blocking F1; see report and patch plan.
- **Decision requests:** None.
- **Blockers:** F1 blocks approval of the configured-pool harness delta.
- **Open / unverified:** Independent Testing and all deferred acceptance/evidence claims remain with their owning routes; no runtime evidence was produced by this Review.
- **Recommended continuation:** Build/Patch should address F1 from the patch plan in a fresh Run; then route any required targeted confirmation and independent Testing.
- **Context refs:** This Run's `invocation.md`, `scope.md`, bound Approved Techplan/T2 task, and Build/Patch `evidence/patch-report-1.md`.
