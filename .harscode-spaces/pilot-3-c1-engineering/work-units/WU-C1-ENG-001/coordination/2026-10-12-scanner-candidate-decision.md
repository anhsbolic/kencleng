# Human/security owner decision — candidate-specific scanner disposition

Status: **APPROVED — bounded option 1 dispositions; implementation pending**, 2026-10-12. No gate implementation, scanner clearance or acceptance established. Decision owner: Anhar / Human security authority. Orchestrator facilitates based on terminal independent Review `7ea1a91074bdb5e064c167ca26003f3e99b5d79ce201c40bda2aa32bb1e1fe74`.

## Recommended option 1 — accept only these bounded dispositions

Historical Gitleaks v8.30.1 `generic-api-key` findings, exact fingerprint scope:

| Fingerprint | Reviewed rationale |
|---|---|
| `6f036c6384868644989134386fc378ed1c19efeb:backend/internal/transport/http/account_security_test.go:generic-api-key:169` | Synthetic fixture request/policy input with fake service; exact usage independently reviewed. |
| `6f036c6384868644989134386fc378ed1c19efeb:backend/internal/transport/http/account_security_test.go:generic-api-key:245` | Synthetic fixture request/policy input with fake service; exact usage independently reviewed. |
| `ef428be3742eddfc59b86fc205bb03624e669e82:backend/internal/domain/account/google_oauth_test.go:generic-api-key:487` | Comment-only ES256/JWT explanation; exact match/blob line 487. |
| `d68b40ce7d8d7f3b7334fdf45d5d4a1899e61c0f:backend/internal/domain/account/service_test.go:generic-api-key:471` | Synthetic fixture request/policy input with fake service; exact usage independently reviewed. |

These are four exact repository-history candidates, not a whole rule/path/commit/test-file exemption. Matching fingerprint and finding source must remain reconstructable; unrelated/new findings still fail. No credential rotation/history rewriting is justified by these proposals.

Gosec G124 candidates: exact `backend/internal/platform/auth/http.go` SHA-256 `bd6a7706c4e5b9ad9af69a33a537b9af0e28fa74a852672692a96f3d034b5009`, locations 49 (cookie helper), 158 (login clear), 277 (logout clear). Capture version `dev` (not independently reproducible version lock); current future scanner version/outcome must be recorded and material changed analysis re-triaged. Rationale: validated constructor/composition preserves production Secure/HttpOnly/Lax/host/path and exact HTTP localhost dev exception; deletion preserves helper flags. No arbitrary config-mutation or actual logout/browser/DB correctness accepted by this decision.

Accepting source-derived scanner dispositions means only **bounded false-positive treatment for these stated constructions**, not declaring R10 runtime fully verified. Material source/hash/location/rule or applicability drift invalidates G124 treatment and must fail closed/re-triage. No entire G124 rule, http.go file or test folder exclusion; no unaudited nosec. Runtime/independent verification remains mandatory per existing owners.

If option 1 approved, next fresh Build designs/applies the smallest reviewable exact-candidate gate treatment consistent with these constraints; feasibility is not established by this proposal. Do not auto-accept any implementation that cannot enforce the bounds. Return the concrete change for owning review/reconciliation before final Testing. Any need for broader allowance is a new Human decision. Final scanner gate must actually execute under current semantics after implementation; old exit 1 remains historical FAIL. No dependency/blocker acceptance implied.

## Alternative option 2 — G124 source remediation first

Accept no G124 exception now; route fresh scoped Build proposal to reinforce validated configuration/cookie construction without changing approved production HTTPS/cookie security or exact dev exception. Actual scanner behavior must be demonstrated; source refactor alone is not guaranteed to clear conditional Secure/helper analysis. Four exact historical false-positive proposals can be accepted separately; never alter history just to remove matches. New protected/security semantics outside existing G1 require explicit owning approval.

## Alternative option 3 — defer acceptance

Keep candidate/gate treatment unresolved; specify any further material evidence needed. No final acceptance claim or downstream dispatch. No full-matrix rerun required merely to postpone policy choice.

## Evidence and limits

Independent Review Approve covers patch only. Diagnostic Gitleaks remains exit 1; Gosec not rerun. Four history fingerprint assessments and three conditional cookie assessments are policy proposals. Existing B-T2-002/B-T2-003 remain open. No T2/T3 readiness/acceptance or dependency commit. Final Testing must resolve/narrow scope-specific prerequisites, distinguish real/simulated/reported/Human-attested evidence and not infer production HTTPS or real lock invariants from cookie units/fixture history.

## Human decision — 2026-10-12

Anhar instructed “oke bro. lanjut” in response to the immediately preceding recommendation “disposition kandidat spesifik → Build treatment sempit → Testing backend T2”. Orchestrator records approval of recommended option 1 within its stated bounds: four exact history fingerprints and three G124 constructions at exact source identity, with runtime limitations unchanged. This continuation is not approval of alternatives, broader exceptions, protected auth changes, dependency/commit/acceptance, or a new Run model. Treatment must fail closed when candidate/rationale identity is no longer valid; if implementation cannot enforce the bounds, stop and return an owner Decision. Actual gate verification remains required.
