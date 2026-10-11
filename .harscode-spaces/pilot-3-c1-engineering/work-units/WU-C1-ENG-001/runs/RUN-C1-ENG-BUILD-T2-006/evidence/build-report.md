# Build report — RUN-C1-ENG-BUILD-T2-006

> Phase: Build; Author / Role: `PARTICIPANT-C1-ENG-T2-IMPLEMENTER-006` / Implementer T2/G1; Session binding: `SESSION-C1-ENG-T2-IMPLEMENTER-006` (fresh). Work Unit: `WU-C1-ENG-001`. Created/Updated: 2026-10-11 Asia/Jakarta. Target HEAD: `059d145d5df64fcda427e5a625486a8e95ce1fd8`; workflow HEAD: `91a97af8ba64f98b0a692a83fbff866a210c45c0`. Invocation selects `gpt-6-sol` / `medium` with exact Run-specific Human approval; actual runtime model/effort not independently established by these captures. User dispatched this specific Run in the current interaction. No registry/Invocation rewrite.

## What changed

**BLOCKED — isolated D8 acquisition/checksum gate failed initially; no tracked backend/dependency write, actual-graph verification, G1 correction, behavioral tests, or commit.** `go mod tidy` on the isolated probe exited 1 because sandbox DNS/socket access to `sum.golang.org` was denied. Harness-required escalation repeated tidy successfully. This recovery does not supersede the Run stop condition or establish gate acceptance.

Only this Run's evidence was written. Existing WU/events/control-tower edits and Loop Health Run directory were preserved. No Product, API, Techplan/task/manifest, frontend, Harscode, or orchestration state was changed. B-T2-002, B-T2-003 and historical B-T2-001 remain open; T3 remains unavailable and T5 independent.

**Execution deviation:** after escalated tidy recovery, Implementer incorrectly treated the environmental recovery as allowing continued read-only gate work. Graph/edge capture, cache inventory, seven upstream legacy-ref checks, seven archive/source comparisons, integrity and focused compile ran after the initial failure. On reopening the bound T2-005 terminal report, its explicit stop/recovery distinction was recognized; all further gate/implementation work stopped. These later results are preserved as observed partial evidence, not acceptance or authorization to reopen this Run. No fresh scanner invocation was started. This deviation must be reconciled by Orchestrator; it must not be hidden or used to close a blocker.

## Authority and input identities

Root/backend AGENTS, current T2 task, approved Techplan §§9–10/12, Build prompt/guidelines/checklist/overlay, workflow/orchestration AGENTS and Run contract, scanner runbook and dependency-verification guide were reopened. Live source anchors: auth Store/LockSession/RevokeSession, OIDC and HTTP boundary, server composition, db.Open, go.mod. The T2-005 trial and Loop Health reconciliation were read; T2-004/005 terminal reports were reopened, with the T2-005 stop explanation recognized late as recorded above. F-T2-003 remains the source-derived stale caller-time expiry issue; it was not corrected.

All 25 bound file identities matched; target/workflow HEAD also matched. The initial hash helper erroneously associated multiple hashes on one line with the first path; a corrected path association and the actual hashes confirmed no input drift. This was a helper parsing error, not artifact drift. Unbound Loop Health reconciliation was present, SHA-256 `f1937f5227ce3033be3d2ba775d9f2aa0862fed79135a599f984a252202d8157`; its older pending-model prose describes preparation, while the bound Invocation records later exact approval and current user dispatch.

```text
6202438fdccd9c47c24d9c1949fa737deee8974c5c8bbcb1b2925d1a5afc5fe9  .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001.md
b4888945c906b26408e1035257a6bf6fef670cffaa87256bad2b0510872d3ea8  .harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/events.md
6d2e41e177e4adaf74f53a06d81dd568def75019c8dc4e4792d938a869896d79  techplan/techplan.md
dfe17b7a3dc0dc89127b040a67f547506a14a4ab591ed0a6343f1a8dd8bcce37  techplan/tasks/T2-backend-identity-session.md
2950408dc48c12d5a8bbed96a88095a20ed7e1e02fd4190293691a90b111b30d  techplan/tasks/manifest.md
18562231cc39a6c1f85ce7c7c20cc8e161d00a203a2b612f2f60117e02b2e611  techplan/report-techplan.md
eff0a3b9c2f0392bbbcaf0f2aa33cf0b93d872afe3eb8b8573f6cef2a81ed051  runs/RUN-C1-ENG-BUILD-T2-004/evidence/build-report.md
07c0330674fb1b43f1d15596e77c2af4a60e3a3260158221ce4c5b590175b78b  runs/RUN-C1-ENG-BUILD-T2-004/evidence/scanner-diagnosis/diagnosis.md
c0bb8330fd984afffe96f8eaecb88612511f42ed34b8483bfab53e8f96f4ff37  runs/RUN-C1-ENG-BUILD-T2-005/evidence/build-report.md
21abe0f80365980275443127399ca3d1ab3e128c3dfa15c6b7806f5d38c4dcef  runs/RUN-C1-ENG-BUILD-T2-005/evidence/orchestrator-reconciliation.md
ff6d6cf3e53651699c5d7b3eb25cf53cf902947fad641901076d5d5b359cacd1  runs/RUN-C1-ENG-BUILD-T2-005/evidence/targeted-d8-trial/trial-report.md
51e1d9c004ce05af37a4bb0c9ce66fa49db55e6424ec0318097c99889ce9051e  runs/RUN-C1-ENG-LOOPHEALTH-002/evidence/loop-health-audit.md
e18ca8bc35601dad839f23b8728b816f5ad1cb25cdd9775ab55c661623474b6e  AGENTS.md
297274385cbbe91ccb498b30edabe540acd8c8108a0a3c4ada4f7b336818d8b3  backend/AGENTS.md
a463e157e8519adfd1ea63c487a4e7546b55ca80cf2186221a6db1f4aacc12ae  docs/project/go-scanner-runbook.md
8e8a183a839c27d887212b1e51bbd96d2e796febed8d2685534d98fe1b3ad4e5  ../harscode-workspace/workflow/3-build-prompt.md
1a85acfa5ebc80837737038198154039d0039bb2c4d6d4addcb62ce34dac9546  ../harscode-workspace/workflow/3-build/guidelines.md
8cb373825edd6e64e2df0e523623bf3b3d6e14fc03580027959e939f9a5ff98a  ../harscode-workspace/workflow/3-build/checklist.md
f806673aa5e7d1d82dae9f433ec54c3562e15ce778b709a141c57d3d22dfaf87  ../harscode-workspace/workflow/orchestrated-run-overlay.md
4610c20ce2670913870718f9886e22c3c00bdf9da5290406a4a4450a818069aa  ../harscode-workspace/orchestration/run-contract.md
98cdc285f7dcd81284ea6d8a1ea23f938a69b8938bb4bc863a3f59aa534113cc  /home/anhar-solehudin/kencleng-workspace/harscode-workspace/orchestration/AGENTS.md
ffbd0521361eac90e82dbee98ba59f01f3e62641afd79bf8e71c4fc8a0e633ce  .harscode-spaces/.local-config.yaml
2c089366f9b54e86da2993dadbeffe3cd7d3786a7b5f0535dd9163521e4db5d3  backend/go.mod
979a350418d3d0e4906ea8af162cb69f8cfcfd38507eaa7d8657f6c7c5d7960f  backend/go.sum
dfeed8714bc1881b9c94e8b9ac8f29ed2226f7396b1f2c789b0bc7dcfd0032b9  backend/internal/platform/auth/store.go
```

`preflight.json` contains copied-source identities and bound-hash verification. Initial and final `git status --short` showed the same pre-existing tracked edits. Final `git diff --exit-code -- backend` exit 0. Live identities remain:

```text
2c089366f9b54e86da2993dadbeffe3cd7d3786a7b5f0535dd9163521e4db5d3  backend/go.mod
979a350418d3d0e4906ea8af162cb69f8cfcfd38507eaa7d8657f6c7c5d7960f  backend/go.sum
dfeed8714bc1881b9c94e8b9ac8f29ed2226f7396b1f2c789b0bc7dcfd0032b9  backend/internal/platform/auth/store.go
```

## Environment, commands and outputs

Probe `/tmp/c1-t2-006-jeu2klrm/probe` copies only backend cmd/internal sources and go.mod/go.sum. No .env, credentials or configured database/provider values were read or copied. Writable build cache `/tmp/c1-t2-006-jeu2klrm/build-cache`; existing module cache `/home/anhar-solehudin/.gvm/pkgsets/go1.24.8/global/pkg/mod`. Explicit Go1.26.9 bin prefix in PATH, `GOTOOLCHAIN=local`, `GOPROXY=off` for cache-first module acquisition; checksum database was retained, not disabled. Off proxy does not disable checksum database network access. No cache deletion/permission changes, scanner upgrade, database substitution, replacement, exclusion or alternate dependency pins.

`go version`: go1.26.9 linux/amd64. Scanner executable `/home/anhar-solehudin/.gvm/pkgsets/go1.26.9/global/bin/govulncheck`: v1.7.0 built with Go1.26.9, metadata only; scanner was not run. The Go environment capture records GOMOD pointing to probe. `commands.jsonl` records argv/cwd/env/exit/time/escalation. Corresponding stdout/stderr files retain outputs; `capture-hashes.json` identifies every retained capture and wrapper. Upstream archive source commands record their destination and URL; archives are retained with hashes in `legacy-source-match.json`. Module inventory identifies cached .info/.mod/.zip/.ziphash SHA-256, h1 metadata where exposed, origin and source paths. Durable captures, not /tmp alone, own this evidence.

| Capture name | Exact argv | Exit | Context |
|---|---|---:|---|
| `go-version` | `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go version` | 0 | none (sandbox) |
| `go-env` | `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go env GOTOOLCHAIN GOROOT GOPATH GOMOD GOCACHE GOMODCACHE` | 0 | none (sandbox) |
| `scanner-version` | `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go version -m /home/anhar-solehudin/.gvm/pkgsets/go1.26.9/global/bin/govulncheck` | 0 | none (sandbox) |
| `baseline` | `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go list -m -json all` | 0 | none (sandbox) |
| `pins` | `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go mod edit -require=github.com/jackc/pgx/v5@v5.9.2 -require=golang.org/x/text@v0.41.0` | 0 | none (sandbox) |
| `tidy` | `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go mod tidy` | 1 | none (sandbox) |
| `tidy-escalated` | `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go mod tidy` | 0 | require_escalated |
| `graph` | `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go list -m -json all` | 0 | none (sandbox) |
| `edges` | `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go mod graph` | 0 | none (sandbox) |
| `origin-github.com_davecgh_go-spew` | `git ls-remote https://github.com/davecgh/go-spew refs/tags/v1.1.1 refs/tags/v1.1.1^{}` | 0 | require_escalated |
| `origin-github.com_jackc_pgpassfile` | `git ls-remote https://github.com/jackc/pgpassfile refs/tags/v1.0.0 refs/tags/v1.0.0^{}` | 0 | require_escalated |
| `origin-github.com_kr_pretty` | `git ls-remote https://github.com/kr/pretty refs/tags/v0.3.0 refs/tags/v0.3.0^{}` | 0 | require_escalated |
| `origin-github.com_pmezard_go-difflib` | `git ls-remote https://github.com/pmezard/go-difflib refs/tags/v1.0.0 refs/tags/v1.0.0^{}` | 0 | require_escalated |
| `origin-github.com_stretchr_objx` | `git ls-remote https://github.com/stretchr/objx refs/tags/v0.1.0 refs/tags/v0.1.0^{}` | 0 | require_escalated |
| `origin-gopkg.in_check.v1` | `git ls-remote https://github.com/go-check/check refs/heads/* refs/tags/*` | 0 | require_escalated |
| `origin-gopkg.in_yaml.v3` | `git ls-remote https://github.com/go-yaml/yaml refs/tags/v3.0.1 refs/tags/v3.0.1^{}` | 0 | require_escalated |
| `upstream-source` | `curl --fail --silent --show-error --max-time 30 --output /tmp/c1-t2-006-jeu2klrm/capture/upstream-github.com_davecgh_go-spew.tar.gz https://codeload.github.com/davecgh/go-spew/tar.gz/8991bc29aa16c548c550c7ff78260e27b9ab7c73` | 0 | require_escalated |
| `upstream-source` | `curl --fail --silent --show-error --max-time 30 --output /tmp/c1-t2-006-jeu2klrm/capture/upstream-github.com_jackc_pgpassfile.tar.gz https://codeload.github.com/jackc/pgpassfile/tar.gz/99d8e8e28945ffceaf75b0299fcb2bb656b8a683` | 0 | require_escalated |
| `upstream-source` | `curl --fail --silent --show-error --max-time 30 --output /tmp/c1-t2-006-jeu2klrm/capture/upstream-github.com_kr_pretty.tar.gz https://codeload.github.com/kr/pretty/tar.gz/a883a8422cd235c67c6c4fdcb7bbb022143e10b1` | 0 | require_escalated |
| `upstream-source` | `curl --fail --silent --show-error --max-time 30 --output /tmp/c1-t2-006-jeu2klrm/capture/upstream-github.com_pmezard_go-difflib.tar.gz https://codeload.github.com/pmezard/go-difflib/tar.gz/792786c7400a136282c1664665ae0a8db921c6c2` | 0 | require_escalated |
| `upstream-source` | `curl --fail --silent --show-error --max-time 30 --output /tmp/c1-t2-006-jeu2klrm/capture/upstream-github.com_stretchr_objx.tar.gz https://codeload.github.com/stretchr/objx/tar.gz/facf9a85c22f48d2f52f2380e4efce1768749a89` | 0 | require_escalated |
| `upstream-source` | `curl --fail --silent --show-error --max-time 30 --output /tmp/c1-t2-006-jeu2klrm/capture/upstream-gopkg.in_check.v1.tar.gz https://codeload.github.com/go-check/check/tar.gz/10cb98267c6cb43ea9cd6793f29ff4089c306974` | 0 | require_escalated |
| `upstream-source` | `curl --fail --silent --show-error --max-time 30 --output /tmp/c1-t2-006-jeu2klrm/capture/upstream-gopkg.in_yaml.v3.tar.gz https://codeload.github.com/go-yaml/yaml/tar.gz/f6f7691b1fdeb513f56608cd2c32c51f8194bf51` | 0 | require_escalated |
| `integrity` | `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go mod verify` | 0 | none (sandbox) |
| `focused-compile` | `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go test -run ^$ ./internal/platform/auth ./internal/platform/db ./cmd/server` | 0 | none (sandbox) |

Initial failure, preserved in `tidy.stderr`:

```text
github.com/jackc/pgx/v5@v5.9.2: verifying module: github.com/jackc/pgx/v5@v5.9.2: initializing sumdb.Client: checking tree#67994042: Get "https://sum.golang.org/tile/8/1/x001/037.p/129": dial tcp: lookup sum.golang.org on 127.0.0.53:53: dial udp 127.0.0.53:53: socket: operation not permitted
```

Escalated tidy exit 0 is environment recovery only. The failure is not a vulnerability or source incompatibility finding. Network access to vuln.go.dev and fresh scanner completion were not tested.

## Partial isolated evidence and gate dispositions

| Gate | Disposition |
|---|---|
| Bound authority/live baseline identities | MATCH; no detected input drift |
| Isolated cache-first resolution/checksum acquisition | Initial FAILED, escalated recovery exit 0; terminal stop applies |
| Complete selected graph / exact pins / delta / Go minima | Later partial evidence captured after stop; no acceptance |
| Origin/tagged source/checksums | Later cache metadata and legacy refs/source matches captured; full gate acceptance not established |
| Integrity | Later `go mod verify` exit 0, `all modules verified`; not gate reopening |
| Focused compile | Later exit 0; auth/db/server `[no test files]`, no behavioral test |
| Fresh full govulncheck / finding JSON / traces / per-ID disposition | NOT RUN; incomplete |
| pgx API/query-mode/config/SCRAM-OpaqueString source review and relevant diffs | Full source-review gate NOT COMPLETED |
| Actual backend graph and actual scan/compatibility acceptance | NOT RUN; prohibited by incomplete isolated gate |
| G1 LockSession correction, expired-during-lock-wait/protocol-negative/person-upsert checks | NOT IMPLEMENTED / NOT RUN |

Selected graph has 20 external modules, no additions/removals against captured baseline and no x/net. Delta: pgx 5.8.0→5.9.2, text 0.29.0→0.41.0, sync 0.17.0→0.22.0, mod 0.27.0→0.38.0, tools 0.36.0→0.48.0. Exact Go/OIDC/OAuth2/jose pins retained. Go minima below were checked in cached version-specific go.mod, including modules with omitted GoVersion fields in graph JSON; maximum declared minimum 1.25.0 fits compiler 1.26.9. Legacy modules without a directive are explicitly labeled. This is metadata/compile evidence only.

| Module | Probe version | Cached go.mod minimum | Partial provenance |
|---|---|---|---|
| `cloud.google.com/go/compute/metadata` | v0.3.0 | 1.19 | Origin repo/ref/commit tersedia |
| `github.com/coreos/go-oidc/v3` | v3.16.0 | 1.24.0 | Origin repo/ref/commit tersedia |
| `github.com/davecgh/go-spew` | v1.1.1 | no go directive (legacy) | refs + source archive cocok; belum full gate acceptance |
| `github.com/go-jose/go-jose/v4` | v4.1.4 | 1.24.0 | Origin repo/ref/commit tersedia |
| `github.com/jackc/pgpassfile` | v1.0.0 | 1.12 | refs + source archive cocok; belum full gate acceptance |
| `github.com/jackc/pgservicefile` | v0.0.0-20240606120523-5a60cdf6a761 | 1.14 | Origin repo/ref/commit tersedia |
| `github.com/jackc/pgx/v5` | v5.9.2 | 1.25.0 | Origin repo/ref/commit tersedia |
| `github.com/jackc/puddle/v2` | v2.2.2 | 1.19 | Origin repo/ref/commit tersedia |
| `github.com/joho/godotenv` | v1.5.1 | 1.12 | Origin repo/ref/commit tersedia |
| `github.com/kr/pretty` | v0.3.0 | 1.12 | refs + source archive cocok; belum full gate acceptance |
| `github.com/pmezard/go-difflib` | v1.0.0 | no go directive (legacy) | refs + source archive cocok; belum full gate acceptance |
| `github.com/stretchr/objx` | v0.1.0 | no go directive (legacy) | refs + source archive cocok; belum full gate acceptance |
| `github.com/stretchr/testify` | v1.11.1 | 1.17 | Origin repo/ref/commit tersedia |
| `golang.org/x/mod` | v0.38.0 | 1.25.0 | Origin repo/ref/commit tersedia |
| `golang.org/x/oauth2` | v0.34.0 | 1.24.0 | Origin repo/ref/commit tersedia |
| `golang.org/x/sync` | v0.22.0 | 1.25.0 | Origin repo/ref/commit tersedia |
| `golang.org/x/text` | v0.41.0 | 1.25.0 | Origin repo/ref/commit tersedia |
| `golang.org/x/tools` | v0.48.0 | 1.25.0 | Origin repo/ref/commit tersedia |
| `gopkg.in/check.v1` | v1.0.0-20201130134442-10cb98267c6c | 1.11 | refs + source archive cocok; belum full gate acceptance |
| `gopkg.in/yaml.v3` | v3.0.1 | no go directive (legacy) | refs + source archive cocok; belum full gate acceptance |

For seven legacy modules without Origin, live upstream refs resolve to commits and every file present in cached module zip matched the corresponding upstream commit tarball (133 files total). `legacy-provenance.json` retains annotated tags and peeled commits; pseudo-version check.v1 resolves its commit prefix through upstream v1 head. `legacy-source-match.json` records commits, archive identities and zero mismatches. This does not claim full independently accepted provenance: pseudo-version timestamp verification and remaining complete source/origin/checksum reconciliation were not completed as an acceptance gate. Modern modules' Origin fields were inventoried, not independently fetched anew. Source comparison and integrity are distinct from a source vulnerability scan.

No fresh advisory disposition is established. Required five IDs GO-2026-4771/4772/5004/5970/6629 remain governed by Approved Techplan; planned pgx/text fixed ranges and prior trial zero findings are upstream evidence, not a fresh scan in this Run. No current finding stream exists, so no assertion of zero findings, absent affected ranges, no new finding, or risk clearance is made. The earlier Go/jose advisory set likewise remains prior evidence; all required fresh reconciliation is deferred. Actual backend remains pgx v5.8.0/text v0.29.0.

## Tests run and verification scope confirmation

Only `go test -run '^$' ./internal/platform/auth ./internal/platform/db ./cmd/server` ran on the probe after the stop boundary as the recorded deviation. It compiled the three packages, with no test files. No protocol-negative, person-upsert, session-lock, fake-clock regression, real PostgreSQL contention or configured auth test ran. No race/concurrency, performance/load, security-class test or broad Testing-owned suite was executed. Scanner metadata inspection is not a source scan.

## Contract check

- [ ] Current build target satisfied in full — blocked at isolated gate.
- [x] Live source re-grounding did not expose a new material product/interface contradiction; known F-T2-003 retained.
- [x] Backend/dependency tracked write boundary preserved; no commit.
- [ ] Required stop sequence observed in full — post-recovery read-only continuation deviation recorded.

## Deferred / not tested here

Complete isolated provenance/source/scan/compatibility gate; actual backend dependency acceptance; bounded post-lock-time G1 correction and focused behavioral evidence. Real configured Google login, PostgreSQL/SCRAM connection and contention, independent Testing and Human acceptance remain unverified. Compile/source identity does not establish those runtime outcomes.

## Flagged for Orchestrator / Techplan / Testing

Scoped environmental blocker at D8 checksum acquisition; recovery succeeded but strict stop closed this occurrence. Orchestrator/Anhar owns reconciliation of the occurrence and the post-recovery execution deviation. Any future execution envelope should explicitly settle whether bounded environment retry permits continued gate evaluation; no waiver is assumed here. Techplan owner is needed only if a future dependency/API/risk choice arises. No additional Run is dispatched or blocker closed by this report.

## Phase handoff

- **Outcome:** BLOCKED — initial isolated D8 acquisition failure; escalated recovery and later partial evidence do not reopen this Run.
- **Result refs:** this `evidence/build-report.md`; `commands.jsonl`, `preflight.json`, `capture-hashes.json`, `legacy-provenance.json`, `legacy-source-match.json`, `inventory.json`, `final-boundary.txt`.
- **Findings:** checksum database DNS/socket denied in sandbox; after recovery Implementer continued read-only gate checks beyond the strict stop boundary, recorded above. No vulnerability/incompatibility finding is claimed.
- **Decision requests:** Orchestrator/Anhar to reconcile the terminal occurrence and execution deviation, and explicitly define permitted environment recovery/continuation in any later authorized envelope.
- **Blockers:** B-T2-002 remains open: no complete isolated gate or actual-backend acceptance. B-T2-003 remains open: correction/focused evidence not done. T3 unavailable; T5 unaffected/independent.
- **Open / unverified:** fresh scan/per-ID reconciliation, complete provenance/API compatibility, actual graph, G1/focused checks and configured provider/DB/independent Testing.
- **Recommended continuation:** Orchestrator reconciles this terminal report and bounded environment-recovery semantics; do not resume this occurrence automatically or dispatch from this Run.
- **Context refs:** this Invocation; Approved Techplan §§9–10/12; current T2 task; scanner runbook/dependency guide; T2-005 terminal stop distinction and separate cache-first trial.
