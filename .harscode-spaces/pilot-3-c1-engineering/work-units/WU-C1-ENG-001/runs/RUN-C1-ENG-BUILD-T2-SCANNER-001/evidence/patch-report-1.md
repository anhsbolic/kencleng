# Scanner treatment — patch-report-1

Phase: Build/Patch. Author: Codex / Implementer, backend-tooling/security. Created/Updated: 2026-10-12.
WORK_UNIT_ID: `WU-C1-ENG-001`. RUN_ID: `RUN-C1-ENG-BUILD-T2-SCANNER-001`.
Participant/Session assignment: `PARTICIPANT-C1-ENG-T2-SCANNER-IMPLEMENTER-001` / `SESSION-C1-ENG-T2-SCANNER-IMPLEMENTER-001`; bukan runtime thread ID terverifikasi.
Selected model/reasoning: `gpt-6-sol / medium`, sesuai exact dispatch approval pada Invocation. Harness mengidentifikasi Codex berbasis GPT-6; exact runtime model/effort tidak terekspos independen dan tidak diklaim dari approval.
Observed target HEAD: `dae55ac44beb0e967b94c2f791d8a0b8d8e72372`, branch `pilot/3-c1-engineering`. Invocation target: `9aabf9c580685271982976dfa9e1cb2b18df0288`. Observed/current-effective workflow HEAD: `651ce57f044bff23fba99082322cee6c066d25a6`.

## Preflight / assignment bindings

Manifest SHA-256 `a8265c7b59b151825ec1c8057b54a20541d1cec68ce1f1bcf267fdc62d0b0773` cocok; **11/11 inputs cocok sebelum edits**, termasuk scope, approved option 1 decision, Approved parent/T2, independent Review/diagnostic, http.go/main.go, Makefile/README dan Compose. Pemeriksaan setelah implementation menunjukkan hanya Makefile/README berubah secara authorized; sembilan input lain tetap cocok. Manifest/Invocation/scope tidak ditulis oleh Run ini.

HEAD delta diperiksa: predecessor→HEAD membawa patch oidc validation literals/cookie tests yang telah independently reviewed, evidence/coordination dan dokumen holds/proportionality. Source http.go/main.go dan gate baseline cocok dengan assignment; live constructor/composition dibuka kembali. Tidak ada material contradiction. Pernyataan progress awal bahwa delta HEAD hanya dokumen/koordinasi dikoreksi setelah pemeriksaan file list lengkap; bukan dasar acceptance.

WU/Events/Control Tower dan reconciliation preparation dibaca sebagai mutable context. Human kickoff ini menjadi launch instruction; metadata prepared/not-launched tidak dijadikan execution evidence. Seluruh pre-existing worktree changes dipertahankan; tidak ada lifecycle projection, Product/Techplan/task, auth, module/pin, API, frontend atau Compose write.

Guidance: root/backend AGENTS; Harscode root/workflow/orchestration/best-practices routers, canonical Build prompt/guidelines/checklist, overlay/run-contract; Go secrets/logging/dependency guidance; project scanner/network/dependency guides, backend architecture dan scoped holds/proportionality. Prior Review `7ea1a91074bdb5e064c167ca26003f3e99b5d79ce201c40bda2aa32bb1e1fe74` dan diagnostic `6d541485fef6c5f3f4d037b610fe689e75901b97b20ee225f08e23f46297d112` tetap applicable untuk rationale/attribution; tidak diulang sebagai audit baru.

## What changed

- `backend/scripts/scanner_gate.py`: helper Python standard-library kecil untuk execute→parse→validate→disposition. G124 hanya tiga exact rule/file/line/column di http.go; exact reviewed production-source hashes/set menjaga constructor/composition dan caller baseline. Perubahan/penambahan production Go source memerlukan re-triage, tanpa silent allowance. Build metadata Gosec harus module v2.28.0 dengan checksum yang diamati, tanpa replacement; report schema/version, loading errors, coverage/statistics, suppression, duplicate/missing candidates dan status diperiksa.
- History hanya empat exact commit/path/rule/start/end-line/fingerprint, dengan SHA-256 blob independently reviewed melalui `git show`; tidak ada allowance path/rule/commit-wide. Kandidat hilang/duplikat, fingerprint/source drift, finding lain atau report invalid gagal.
- Gitleaks upstream memakai exit 1 untuk findings maupun partial errors. Helper menetapkan `--exit-code 10`: complete finding reports harus exit 10; execution/partial errors exit 1 tidak dapat diterima walaupun report berisi kandidat approved. Ini diverifikasi terhadap [versioned primary source, findingSummaryAndExit](https://raw.githubusercontent.com/gitleaks/gitleaks/v8.30.1/cmd/root.go) dan [git scan error path](https://raw.githubusercontent.com/gitleaks/gitleaks/v8.30.1/cmd/git.go), lalu behavior aktual diuji. Tidak memakai regex human output sebagai disposition authority.
- `backend/Makefile`: lint memanggil `make gosec`; history/worktree menggunakan helper dengan existing Compose runner/image. `make test-scanner-gate` ditambahkan. All/full-history/backend tetap; staged, unstaged tracked dan semua nonignored untracked backend files tetap attempted meski ada failure. Worktree memiliki nol disposition. Gosec tetap seluruh applicable production source/rules dan tidak mengizinkan nosec comments; Gitleaks native ignore files/comments tidak memperluas exception dan repository-native config ditolak. Ignored local env tetap excluded kecuali force-tracked.
- `backend/scripts/test_scanner_gate.py`: tujuh behavioral tests dengan positive/negative oracles; source guards nyata, synthetic mutation/new caller/history failure, scanner status/report/schema/source/location/fingerprint drift, metadata rejection, no payload output dan seluruh worktree attempts setelah finding.
- `backend/README.md`: menjelaskan actual scope, versi, error/disposition semantics, Python 3 prerequisite dan batas effective PASS.

Output hanya safe status/count/category; tidak menulis atau mencetak Secret/Match, credential, raw scanner JSON atau stderr. Source/history probe bytes dipakai private untuk hashes. Unknown/incomplete/error tidak dinormalisasi menjadi PASS.

### Exact material identities

| File | Before SHA-256 | After SHA-256 |
|---|---|---|
| backend/Makefile | `7a36b24c3cdc0e018d43c47a4934b9c818a5fd8842c72b2e72de7214b258232f` | `d0ddf3678f8f3a8a071159e6ffeb11e9eaf761193de8ec19e1b63f04f4aca76f` |
| backend/README.md | `bf7839d364d9b8cb0f2db9331acb83f8fd4bc4ffe02040cacd23f44814185e0d` | `5cbf1d89b28f93c82d619d0d12fecd2b5a01d67e1ac89e862e90cd2557c60a17` |
| backend/scripts/scanner_gate.py | absent | `1c231ed1dd8461f2b16eeef327508646a3f65d8ca55fafa9ad82854ae8ef8d30` |
| backend/scripts/test_scanner_gate.py | absent | `4097a9b3cd5fa0e9c5cc8187ad2c57cbc4c3dff9c3351dcb1469177999ac14fe` |

http.go `bd6a7706c4e5b9ad9af69a33a537b9af0e28fa74a852672692a96f3d034b5009` dan main.go `d217dfd21898ce89d3eee96683b7512fd7010ec0e55da04c6d211bb5686e2246` tetap; additional current production guards ada pada helper, bukan perubahan source. Historical blob identities tetap yang dicatat Review.

## Tests run

Commands scanner/Make dijalankan dari `backend/` kecuali disebutkan. Environment: Python **3.14.4**; actual selected Go **go1.26.9 linux/amd64** (`GOTOOLCHAIN=go1.26.9 go version`); Gosec executable `/home/anhar-solehudin/.gvm/pkgsets/go1.24.8/global/bin/gosec`, display `dev`, build Go **1.26.7**, module **github.com/securego/gosec/v2 v2.28.0**, sum `h1:ZsSdiDb0AtTpLFVol5z91gbMei9ZiLEPG/pZjZujp7c=` (observed `go version -m`, including structured `-json`). Gitleaks configured image **ghcr.io/gitleaks/gitleaks:v8.30.1**, actual `version` **v8.30.1**; Podman **5.7.0**. Image digest tidak direkam; tidak mengklaim immutable image provenance. Tidak ada dependency acquisition/pin write atau checksum bypass.

| Verification | Actual result / oracle |
|---|---|
| SHA-256 preflight and final checks | 11/11 cocok sebelum edit; final 9 unchanged + 2 authorized delta. |
| `make test-scanner-gate` (then final `make -C backend test-scanner-gate`) | Final **7 tests, OK, exit 0**. Candidate matching, unrelated/new findings, disappearing/duplicate candidates, location/rule/source drift, invalid/duplicate JSON, loading errors, unexpected exit/signal, timeout/executable failure, version drift and payload privacy covered. Six tests initially passed; seventh integration oracle added and seven passed again after exit-code correction. |
| `GOTOOLCHAIN=go1.26.9 GOCACHE=/tmp/c1-scanner-snkrni9o/cache make gosec` | Final **exit 0 effective PASS**; underlying `gosec -fmt=json -nosec ./...` **raw exit 1 / 3 G124 / accepted 3**. Full applicable source report: 6 files, 750 lines, nosec 0, no Golang errors. Two actual integrated executions passed, second on final helper. `GOFLAGS=-mod=readonly` enforced by helper; cache writable in this sandbox. |
| `make gitleaks` via approved harness escalation | Final **exit 0 effective PASS**. History **raw exit 10 / 4 findings / accepted 4**; unstaged/staged **raw exit 0 / 0 findings**; three inventoried untracked files **raw exit 0 / 0 findings**. Third was this Run's temporary generated Python bytecode, subsequently removed; final deliverables are two .py files already scanned. Initial gate also passed before findings-exit disambiguation (history raw exit 1); superseded by final execution with exit 10, not concealed. |
| `python3 /tmp/c1-scanner-coverage.py` via approved escalation | Disposable synthetic Git repo; image `--network=none`, repo read-only mount, no env/credential mount. Clean unstaged/staged: **0/0**. Synthetic unstaged `backend/tracked.txt`, force-staged ignored `backend/.env.forced`, untracked `new file.txt`: each **raw exit 10 / 1 finding**, effective **FAIL as required**, all three attempted. Ignored untracked `.env.ignored` excluded. Invalid target: **raw exit 1 / report bytes 0**, rejected. Fixture repo removed by TemporaryDirectory. Earlier same coverage probe at default findings-exit 1 also rejected all synthetic findings; final probe repeated only because exit semantics changed. |
| `git diff --check -- backend` | Exit 0. |

Preflight recovery: sandbox `podman version --format '{{.Client.Version}}'` failed with `Failed to obtain podman configuration: set sticky bit on: chmod /run/user/1000/libpod: read-only file system`. One same-command harness-approved retry succeeded (5.7.0); no permission mutation/global cache clear. Subsequent container version/help and selected gate executions explicitly used harness-approved escalation. No failed recovery/substantive stop was overridden.

An initial private asynchronous Gosec capability probe produced a valid report with the same three locations, but its raw exit metadata was not retained after yield; **not used as acceptance evidence**. Actual integrated commands above establish raw/effective outcomes. An accidental `make test-scanner-gate` from repository root returned exit 2/no target; corrected with `make -C backend test-scanner-gate`, seven tests passed. Raw scanner reports/payloads were not duplicated into evidence. Disposable probe cache/report/script remain under bounded `/tmp` paths; no persistent infrastructure was created.

## Verification scope confirmation

**Authorized targeted security-class exception:** only Gosec/Gitleaks runs necessary to verify this treatment, their actual coverage/error behavior and synthetic scanner integration. This is explicitly authorized by scope.md/Invocation, not an unreported scope deviation. No broad `make verify`, race/concurrency, performance/load, govulncheck/dependency reacquisition, DB/runtime/HTTP acceptance, browser/provider or Final Testing executed. No real Google login. Build result establishes scanner-treatment correctness/integration confidence only, not raw zero-findings or whole-T2 acceptance.

## Contract check

- [x] Current scanner-treatment Build target satisfied within scope.md: seven approved candidates only, coverage retained, unexpected/source/fingerprint/report/tool failures fail closed.
- [x] Live-code re-grounding did not invalidate a material contract assumption. Current hashes/rationale remain the bound ones; no auth/security-core refactor or broader exception introduced.

## Deferred / not tested here

Real Google/manual and final integrated journey/rendered UI/cache/Human walkthrough remain **deferred** under `docs/project/c1-google-login-verification-hold.md` and `c1-integrated-verification-hold.md`. These decisions provide scheduling authority, not scanner/runtime evidence.

DB/SCRAM/session-lock/HTTP/production HTTPS/final backend/dependency obligations are **not applicable to this Build's scanner-treatment verification claim** and remain required at their owning claim/dependency boundaries. Prior dependency evidence is not reacquired or newly accepted. No waiver/residual-risk acceptance, dependency commit, B-T2-002/B-T2-003 closure, T2→T3 edge satisfaction, T3 readiness/dispatch or C1 completion.

## Flagged for Techplan / Testing

Current claim: candidate-specific filtering correctness plus selected gate integration. Required evidence: exact approved source/rationale, bounded match semantics, fail-closed negatives and retained actual tool coverage. Current sufficient evidence: cited independent Review for attribution/rationale, matching source hashes, new seven-test oracles, actual Gosec/Gitleaks gates and synthetic Git-visible/error integration.

Material remaining gap: **independent scrutiny of novel filtering/error/coverage implementation**, particularly Gosec source-guard completeness and Gitleaks partial-error disambiguation. Recommend scoped independent Code Review of these four files and this evidence; material value is finding a permissive/missed failure path, not redoing settled candidate attribution. No phase dispatched here and no N/A verdict predeclared. Additional independent scanner Testing is not automatically required by chronology: Orchestrator should identify any material enforcement/integration question not answered by these results, or any current authority requiring independent evidence, before defining such a Run. Final backend/runtime Testing obligations remain separate and unsatisfied by this Build.

## Phase handoff

- **Outcome:** COMPLETED — bounded scanner-treatment Build; no whole-T2 acceptance.
- **Result refs:** this `evidence/patch-report-1.md`; four exact implementation identities in “Exact material identities”.
- **Findings:** no unresolved implementation finding observed in the executed scope; raw approved scanner findings remain 3 G124 + 4 history candidates, with effective accepted disposition.
- **Decision requests:** none for existing bounded treatment; no wider exemption requested.
- **Blockers:** none for this Build claim; existing owning B-T2-002/B-T2-003 are not closed or reclassified.
- **Open / unverified:** independent review of novel gate implementation; runtime/final/deferred obligations and provenance limits above.
- **Recommended continuation:** Orchestrator reconcile exact claim/delta/evidence; assess scoped independent Code Review for remaining gap. Do not dispatch Final Testing/real Google or claim readiness from this handoff.
- **Context refs:** this Run Invocation/scope/input-bindings; Human scanner-candidate decision; bound independent Review/diagnostic; Approved parent/T2; four changed backend files only.
