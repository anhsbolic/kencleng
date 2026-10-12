# Scanner treatment — authorized Build scope

Preparation input, not execution evidence. Human option 1 decision: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/coordination/2026-10-12-scanner-candidate-decision.md`.

Outcome: implement smallest reviewable candidate-specific false-positive treatment for four exact Gitleaks history fingerprints and three exact G124 candidates already independently reviewed. Do not change auth/browser/DB behavior or rewrite history. Scanner gate must continue rejecting other/new findings and tooling failures.

## Allowed writes

Backend Makefile scan invocation/integration, minimal backend scanner policy/helper file(s) and focused behavioral tests for treatment, backend README explanation of actual gate. Build selects concrete filenames and mechanics based on actual tool capability. Preserve existing scanner coverage: Gosec applicable backend source; Gitleaks backend history plus staged/unstaged tracked changes and untracked nonignored files. Ignored local env remains excluded; force-tracked secrets/new findings must remain detectable. Root Compose only read: existing configured tool image/runtime, no unauthorized root infrastructure changes. No source auth, frontend, API, migrations, modules/pins or higher Product/Techplan/task writes.

## Enforcement bounds

- History dispositions match only approved commit/path/rule/line fingerprints. Validate candidate identity/source mapping to prevent a broader commit/path/rule allowance. Four dispositions are fixture/comment false positives, not permission to accept real credentials or rewrite Git history.
- G124 treatment limited to current reviewed http.go SHA-256 `bd6a7706c4e5b9ad9af69a33a537b9af0e28fa74a852672692a96f3d034b5009`, rule and candidate locations 49/158/277. Preserve constructor/composition assumptions used by Review. Guard materially relied-upon relevant source identity when treatment depends on it; changed source/rule/location/applicability => fail/re-triage, not silent suppression.
- No blanket nosec, whole-rule/file/test-folder exclusion. Use actual structured finding semantics; do not approve by count or regex against arbitrary human output alone.
- Scanner exit/tool execution errors, invalid/incomplete report or unexpected findings must remain visible failures. Successful filtering of known candidates is a scoped accepted-disposition result, not raw scanner zero-findings; report both raw and effective outcomes.
- Discover actual tool versions/capabilities. Gosec earlier capture is dev; record real build metadata and stop on material semantic incompatibility rather than guessing unsupported configuration.
- If tool-native treatment cannot enforce bounds, a minimal safe helper is within scope. Avoid a general scanner framework. If enforcing bounds requires a wider exemption, different security/scan policy, new dependency or auth refactor, stop for owner Decision before expansion.

## Meaningful focused verification

Focused positive/negative checks should demonstrate: exact approved candidate handling; unrelated/new finding rejection; source/fingerprint drift rejection; scanner/report error propagation; existing Git-visible scope retained. These are behavioral gate tests, not tests duplicating implementation details. Use synthetic values only, never real credentials. Run new tests to prove executable oracle.

Targeted Gosec/Gitleaks runs needed to demonstrate actual selected scanner integration are explicitly justified Build security-class checks for this treatment; declare them in canonical Verification scope confirmation. Use current runtime commands where suitable, adapt only within allowed scope; redacted safe output, bounded timeout/disposable resources. Do not replay full make verify/race/vulnerability/provider/DB/browser acceptance suites. Final backend Testing follows after returned implementation reconciliation; holding real Google/E2E does not waive backend gates.
