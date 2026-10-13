# Testing scope — current auth harness runtime and cleanup

Status: PREPARED; Human model approval, fixture authorization, and mechanical dispatch are pending. No command, fixture, or runtime operation has occurred in this Run.

## Claim and applicability

Independently verify the current full-ID cleanup implementation at its real Podman/PostgreSQL successful-fixture boundary, execute the six-case `t2` auth harness against this Run's own isolated PostgreSQL fixture, and establish the current backend `make verify` result for the post-patch tree. This is bounded Testing evidence, not final T2 acceptance.

Testing is applicable because the latest two-file cleanup patch has only Builder-reported mock tests and static Review evidence. The prior real-fixture harness Build used an earlier cleanup implementation and was not independent Testing. A fresh Run-owned fixture materially verifies that the final harness can identify, inspect, and clean its actual owned PostgreSQL container and anonymous data volume while exercising real PostgreSQL/TCP/SCRAM and contention cases.

## Evidence sweep and limits

- Reuse the original runtime-harness Build report only for historical six-case executability, PostgreSQL/TCP/SCRAM, simulated-provider, and successful cleanup observations on its own fixture and earlier runner version.
- Reuse the full-ID Build report for its five mocked boundary cases and exact two-file delta; the new Testing Run will independently execute the focused mock suite and the selected real-fixture suite.
- Reuse the exact-delta Review for static approval of validation and no-delete ordering. Review did not run tests or runtime.
- `backend/go.mod` and `backend/go.sum` remain byte-identical to `RUN-C1-ENG-TEST-T2-BACKEND-GATE-002` (hashes in `input-bindings.json`). Its `go mod verify` result remains sufficient for those unchanged module inputs; this Run will not claim new dependency provenance.
- Fresh `make verify` is required because backend harness files changed after the prior gate. Preserve actual command/target outcomes and scanner raw findings separately from accepted dispositions.
- Still open/out of this claim: `B-T2-002`, `B-T2-003`, full R1–R14 T2 acceptance, configured application-runtime privilege/SCRAM proof, the post-lock expiry finding, real Google/provider/browser/HTTPS, integrated FE/UI/cache isolation, Human semantic walkthrough, T2/T3 readiness, and C1 completion.
- Google login and integrated verification remain deferred under their owning Human holds. No hold is expanded to backend runtime obligations.

## Authorized commands and stop conditions

After exact approvals and Human dispatch, run only from `backend/`:

1. `PYTHONDONTWRITEBYTECODE=1 python3 scripts/test_auth_harness.py` — current five mocked command-boundary tests, including malformed/short inspect IDs and no-removal assertions.
2. `PYTHONDONTWRITEBYTECODE=1 python3 scripts/auth_harness.py --run-id RUN-C1-ENG-TEST-T2-AUTH-HARNESS-RUNTIME-001 --suite t2` — all six selected auth-harness cases using this Run's disposable fixture.
3. `PYTHONDONTWRITEBYTECODE=1 make verify` — current ordered backend lint/unit/race/contract/security gate. Preserve nonzero raw Gosec/Gitleaks findings and their accepted dispositions; never describe accepted findings as zero findings.

Do not rerun `go mod verify` unless module inputs drift; if they drift, stop and return for scope reconciliation. Follow the declared Go 1.26.9 and current scanner/network runbooks, bounded recovery rules, and stop conditions. Do not continue after a substantive failed `make verify` constituent or disguise a scanner failure as environmental recovery.

Do not start the app, use `DATABASE_URL`, load `.env`, contact Google/provider, use browser/HTTPS, alter dependencies/source/policy/allowances, access shared DB/state, or run unrelated tests/gates. No broad runtime/T2/T3 acceptance claim.

Preserve the current dirty workspace. `backend/scripts/__pycache__/` is currently untracked and is intentionally excluded from semantic input bindings; do not delete or rewrite it. At execution, report whether the repository's worktree scanner includes it and the actual coverage/outcome. Do not add an ignore rule or allowance to make the gate pass.

## Fresh disposable fixture — exact Human approval required

Authorize only for this Run: one unique, ephemeral Run-owned `postgres:16` container through configured Podman, existing local image only, loopback ephemeral port, no host bind mounts, no shared named volume or existing Compose network. Use fixture-only current identity/session schema, distinct disposable provisioner and non-owner restricted runtime roles, and generated ephemeral credentials passed only to the child process. Never read/load `.env`, application `DATABASE_URL`, or provider credentials; never persist DSNs, passwords, tokens, or payloads.

Before deletion, validate this Run's exact resource label/owner, name, and full container ID. On success/failure, remove only the verified owned container and its anonymous data volume with bounded timeouts, then verify absence. Do not use broad-prefix cleanup or affect unrelated resources. If ownership/isolation/restricted grants/cleanup cannot be established, fail closed and report; do not fall back to shared or elevated application state.

## Results and evidence classification

The report must distinguish actual PostgreSQL/TCP/SCRAM/contention, `httptest` HTTP behavior, simulated callback Provider, mocked command-boundary cases, gate/scanner outcomes, and cleanup observations. State exactly which six cases ran. Successful cleanup proves only this Run's observed path; it does not establish every abnormal cleanup branch. No Human-attested or real-provider evidence is expected from this Run.

Write only `evidence/testing-report-1.md` and minimal result-bearing evidence needed for material claims. Follow canonical Testing prompt/guidelines/checklist and orchestrated overlay. Include exact current bindings/delta, readiness, required Test Focus anchor execution, full Techplan contradiction read, PASS/FAIL/flagged results, claim limits, and exactly one `## Phase handoff`. Production fixes return to Build/Patch authority.
