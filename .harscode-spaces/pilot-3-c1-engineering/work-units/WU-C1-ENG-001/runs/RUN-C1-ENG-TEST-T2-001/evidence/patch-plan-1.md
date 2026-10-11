# Patch plan — T2 final lint gate

- Phase / Author: Independent Testing / Codex, Tester.
- Created: 2026-10-11 Asia/Jakarta.
- Work Unit / Run: `WU-C1-ENG-001` / `RUN-C1-ENG-TEST-T2-001`.
- Participant / Session bindings: `PARTICIPANT-C1-ENG-T2-TESTER-001` / `SESSION-C1-ENG-T2-TESTER-001`.
- Model / effort: Invocation selected and approved `gpt-6-sol / medium`; actual runtime model/effort not independently exposed.
- Target / workflow revision: `059d145d5df64fcda427e5a625486a8e95ce1fd8` / `91a97af8ba64f98b0a692a83fbff866a210c45c0`.

## F-TEST-T2-001 — Required lint fails

`make verify` recovery exits 2 at Staticcheck: `internal/platform/auth/oidc.go:38:15` and `:42:15`, ST1005 (error strings should not be capitalized). Exact capture: `make-verify-recovery.stdout` / `.stderr`, command record in `commands.jsonl`. Source SHA-256: `a9a89117ba682e3963f5a02ec526d566763e1f991994cd5526fb01c60016a812`.

This file is part of the current T2 surface but outside the prior six-file retry Review scope. No regression introduced by the six-file retry is established. No authentication failure or privilege bypass is inferred from the style finding. It blocks the required final verification route.

## Bounded Build action

Change only the initial capitalization of the two constructor validation error strings, preserving validation branches, safe error propagation, fixed provider trust and all runtime semantics. Build owns production writes; Testing has made none. Do not suppress ST1005, alter the lint gate, or broaden the patch into auth/config redesign.

## Reverification and remaining prerequisites

Run current backend `make verify` with writable Go and Staticcheck caches and the required tools available; retain every constituent outcome. Recheck constructor/protocol coverage and `go mod verify`. The missing `gitleaks` executable is an environment issue, not a production code patch.

Google configuration/registered callback/consent/test-user interaction, authorized disposable PostgreSQL with restricted runtime credentials and distinct provisioning identity, and configured HTTPS remain necessary for their respective real tests. A style patch or unit pass cannot close those gaps. Orchestrator must reconcile this report and route a fresh Build/Patch occurrence, followed by independent Testing as appropriate; do not reopen this terminal report or infer T2 acceptance/T3 readiness.
