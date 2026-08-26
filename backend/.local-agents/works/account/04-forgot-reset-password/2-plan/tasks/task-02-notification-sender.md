# Task 02 — Notification platform: SendPasswordResetEmail

> Back-reference (contract): `../techplan.md` — sections 1–8 are the source of truth. Techplan wins over this file on any apparent conflict.
> Splitting axis: dependency/sequence chain (see `manifest.md`). No dependency on Task 01 — different package; can execute in parallel with it. Task 03 depends on BOTH.

## Scope

**In scope:**
- `SendPasswordResetEmail` added to the `notification.Sender` interface
- Implementations on `FakeSender` and `DevSender`
- Compile-ripple fix through every existing test fake implementing `Sender`
- Sender-level unit tests mirroring the existing FakeSender test style

**Out of scope (this task):**
- Who calls the method or when (Task 03 owns post-commit send policy)
- Any change to `SendVerificationEmail`/`SendNudgeEmail` semantics or nudge constants

## Dependencies

None (parallel with Task 01). Must land before Task 03 (service calls the new method).

## Interface addition (exact signature — techplan §8)

```go
// platform/notification/sender.go (Sender interface addition)
SendPasswordResetEmail(ctx context.Context, to, token string) error
```

## Implementation details (redistributed from techplan §10)

**File**: `internal/platform/notification/sender.go`
- Add method to `Sender` interface alongside the existing two (explicit-method style preserved — do NOT generalize into a `SendTokenEmail(kind, ...)` refactor; that option was considered and rejected implicitly by keeping house style).
- `FakeSender`: no-op returning nil, consistent with its current log-safe behavior (sender.go:34–43 pattern).

**File**: `internal/platform/notification/dev_sender.go`
- `DevSender.SendPasswordResetEmail` appends a THIRD outbox line type (recipient + token) to the dev outbox file; preserve mode 0600 and the existing `append` helper (dev_sender.go:38–55 pattern). This is the dev stand-in for an SMTP inbox — tokens stay out of `log.Printf` output (AGENTS.md golden rule: outbox ≠ log stream).

**Compile ripple**: any test fake in the codebase implementing `notification.Sender` gains the method (expected locations: `internal/domain/account/*_test.go` helpers like `integrationSilentSender`). Mechanical addition returning nil; no behavior.

## Binding constraints (techplan §7 risks + golden rules)

- The plain token leaves the process exactly once, through this method's argument — never logged. Log lines may carry the FACT of sending + sanitized category on failure (the category-classification helper lives in domain/account, not here; platform stays business-rule-free per AGENTS.md §1 layout rules).
- PII: recipient address goes to the sender but is never logged by it (FakeSender redacts — existing convention, sender_test.go:14–58).

## Known doc-drift consequence (accepted, techplan §14 Active #3)

Root `AGENTS.md` §5 describes dev-outbox contents as verification tokens; after this task the outbox carries a third line type. Doc patch is human-owned — do NOT edit root AGENTS.md from this task. Flag in the task report instead.

## Testing checklist (this task's items)

- [ ] `TestFakeSender_SendPasswordResetEmail_ReturnsNilNoPIIInLog` — mirror `TestFakeSender_SendNudgeEmail_ReturnsNilNoPIIInLog` (sender_test.go:40): call it, capture stderr/log sink, assert nil error and no recipient/token material in output
- [ ] DevSender path exercised under `APP_ENV=development` manual check OR a unit test asserting a third line lands in a temp outbox file with mode 0600 preserved (follow however `DevSender` is currently tested; TBD — verify current DevSender test coverage at build time and match its pattern)

## Gate

`go build ./...` green across ALL packages (proves the interface ripple is complete — any missed fake fails here, which is the point); `go test ./...` green.
