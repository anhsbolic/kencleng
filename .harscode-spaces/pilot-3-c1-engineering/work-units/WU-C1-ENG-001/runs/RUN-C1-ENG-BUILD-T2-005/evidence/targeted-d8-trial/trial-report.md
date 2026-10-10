# Targeted D8 verification trial

Date: 2026-10-11. Human authorization: “oke, kita fokus ke cara ini. coba dulu”. This is a separate method trial, not a new dispatched engineering Run, not a reopening or rewrite of terminal RUN-C1-ENG-BUILD-T2-005. No backend implementation/dependency writes or commit. Captures reside in `/tmp/c1-d8-targeted-f3l9j43t`.

Outcome: cache-first acquisition worked. All selected module files/source directories were available; no `go mod download -json all` or targeted download was needed in this trial. Prior T2-005 escalated download had populated missing cache inputs; this is not proof that the cache was complete before that recovery. Exact D8 probe graph, integrity, focused compile and fresh full source scan completed successfully. Full production acceptance/G1 completion is not claimed.

The probe copies current backend cmd/internal sources and go.mod/go.sum only, excludes .env and secrets, edits solely the approved pgx v5.9.2/text v0.41.0 pins, and lets tidy resolve transitives. Selected Go is 1.26.9 with GOTOOLCHAIN=local and explicit Go1.26.9 PATH prefix. Scanner v1.7.0 is built with Go1.26.9. Writable isolated build cache is `/tmp/c1-d8-targeted-f3l9j43t/build-cache`. Exact environment/cwd: evidence/context.json. Existing module cache is `/home/anhar-solehudin/.gvm/pkgsets/go1.24.8/global/pkg/mod`; cache inventory records .info/.mod/.zip/.ziphash SHA256, stored h1, source directories and upstream Origin metadata when present. No cache deletion, permission change or dependency substitution.

## Commands and results

All cwd values and command/output identities are in evidence/commands.jsonl. source-scan used harness require_escalated with Human authorization for vuln.go.dev access; other commands ran sandboxed. Compile completed before scan. No failed command/retry occurred in this trial.

| Command | Exit | Captures relative to evidence/ |
|---|---:|---|
| `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go version` | 0 | `go-version.stdout` / `go-version.stderr` |
| `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go version -m /home/anhar-solehudin/.gvm/pkgsets/go1.26.9/global/bin/govulncheck` | 0 | `scanner-version.stdout` / `scanner-version.stderr` |
| `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go list -m -json all` | 0 | `baseline.stdout` / `baseline.stderr` |
| `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go mod edit -require=github.com/jackc/pgx/v5@v5.9.2 -require=golang.org/x/text@v0.41.0` | 0 | `pins.stdout` / `pins.stderr` |
| `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go mod tidy` | 0 | `tidy.stdout` / `tidy.stderr` |
| `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go list -m -json all` | 0 | `graph.stdout` / `graph.stderr` |
| `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go mod graph` | 0 | `edges.stdout` / `edges.stderr` |
| `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go mod verify` | 0 | `integrity.stdout` / `integrity.stderr` |
| `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go test -run '^$' ./internal/platform/auth ./internal/platform/db ./cmd/server` | 0 | `focused-compile.stdout` / `focused-compile.stderr` |
| `/home/anhar-solehudin/.gvm/pkgsets/go1.26.9/global/bin/govulncheck -json ./...` | 0 | `source-scan.stdout` / `source-scan.stderr` |

Integrity output: `all modules verified`. Focused compile uses `go test -run '^$'`: the auth/db/server packages compiled, but reported no test files; no behavioral unit test ran.

Full `govulncheck -json ./...` inspected the copied current source under approved pins. Structured parsing consumed all 204 JSON objects: config 1, SBOM 1, progress 2, OSV metadata 200, finding **0**. Empty stderr and exit 0. Go1.26.9, source mode, symbol level, database https://vuln.go.dev, database last modified 2026-10-08T22:31:09Z. SBOM matches pgx v5.9.2/text v0.41.0/auth pins/Go1.26.9. Full stream retained at evidence/source-scan.stdout, parsed summary evidence/scan-summary.json. OSV objects are advisory metadata, not findings; no finding traces exist because finding count is zero.

## Selected graph, origin and Go minimums

| `cloud.google.com/go/compute/metadata` | v0.3.0 | 1.19 | `b952f41fc832fbe5ed74f47ae0e6d6f2e93448e0` |
| `github.com/coreos/go-oidc/v3` | v3.16.0 | 1.24.0 | `e9584733f8bb6c4683d1e98b4fb22eee121f7dff` |
| `github.com/davecgh/go-spew` | v1.1.1 | no go directive (legacy) | Origin field absent in legacy cache metadata |
| `github.com/go-jose/go-jose/v4` | v4.1.4 | 1.24.0 | `0e59876635f3dbf46d7b5e97b52bb75a3f96e7d9` |
| `github.com/jackc/pgpassfile` | v1.0.0 | 1.12 | Origin field absent in legacy cache metadata |
| `github.com/jackc/pgservicefile` | v0.0.0-20240606120523-5a60cdf6a761 | 1.14 | `5a60cdf6a76120dc3d5152b95f3b5fd8aa7cc9eb` |
| `github.com/jackc/pgx/v5` | v5.9.2 | 1.25.0 | `0aeabbcf11d859229c1f0b20e710d3596c76bf27` |
| `github.com/jackc/puddle/v2` | v2.2.2 | 1.19 | `bd09d14bd4018b6d65a9d7770e2f3ddf8b00af1c` |
| `github.com/joho/godotenv` | v1.5.1 | 1.12 | `3fc4292b58a67b78e1dbb6e47b4879a6cc602ec4` |
| `github.com/kr/pretty` | v0.3.0 | 1.12 | Origin field absent in legacy cache metadata |
| `github.com/pmezard/go-difflib` | v1.0.0 | no go directive (legacy) | Origin field absent in legacy cache metadata |
| `github.com/stretchr/objx` | v0.1.0 | no go directive (legacy) | Origin field absent in legacy cache metadata |
| `github.com/stretchr/testify` | v1.11.1 | 1.17 | `2a57335dc9cd6833daa820bc94d9b40c26a7917d` |
| `golang.org/x/mod` | v0.38.0 | 1.25.0 | `792ac169a90372d88fb14e712cb793061ba0c104` |
| `golang.org/x/oauth2` | v0.34.0 | 1.24.0 | `acc38155b7f6f36aefcb58faff6f36d314dd915c` |
| `golang.org/x/sync` | v0.22.0 | 1.25.0 | `1eb64d4bc0cde6da1bb8ebc7f178bb577508e5d0` |
| `golang.org/x/text` | v0.41.0 | 1.25.0 | `acdba6655fd45cdb5ab73c9d6a8981333bd65a39` |
| `golang.org/x/tools` | v0.48.0 | 1.25.0 | `05f9cb5d358503005bd6f82b17916d226ca7b210` |
| `gopkg.in/check.v1` | v1.0.0-20201130134442-10cb98267c6c | 1.11 | Origin field absent in legacy cache metadata |
| `gopkg.in/yaml.v3` | v3.0.1 | no go directive (legacy) | Origin field absent in legacy cache metadata |

Complete module edges: evidence/edges.stdout. All 20 selected external modules have .info/.mod/.zip/.ziphash and extracted source in cache. Maximum declared dependency Go minimum is 1.25.0, compatible with selected Go1.26.9; legacy modules without a Go directive are explicitly recorded rather than assigned an invented minimum. Origin repo/ref/hash metadata is present for all approved exact D8 pins and changed transitives. Several unchanged legacy modules lack the Origin field; their cache/source and integrity are available, but independent upstream tag-to-commit validation for those modules was not performed. This limits any assertion of a complete full-graph provenance gate.

No selected module additions/removals relative to copied baseline. Version delta: pgx 5.8.0→5.9.2; text 0.29.0→0.41.0; sync 0.17.0→0.22.0; mod 0.27.0→0.38.0; tools 0.36.0→0.48.0. No x/net selected. No new explicit transitive pin.

## Security outcome and compatibility scope

Five-ID OSV metadata was extracted directly from the fresh scanner stream to evidence/five-advisories.json:

| Advisory | Fresh database fixed version for selected module | Probe disposition |
|---|---|---|
| GO-2026-4771 | pgx/v5 5.9.0 | Selected 5.9.2 outside affected range; zero findings |
| GO-2026-4772 | pgx/v5 5.9.0 | Selected 5.9.2 outside affected range; zero findings |
| GO-2026-5004 | pgx/v5 5.9.2 | Selected 5.9.2 outside affected range; other pgx major paths not selected |
| GO-2026-5970 | x/text 0.39.0 | Selected 0.41.0 outside affected range; zero findings |
| GO-2026-6629 | x/text 0.41.0 | Selected 0.41.0 outside affected range; zero findings |

No surviving/new source finding was reported. This is source scanner evidence at the recorded database state, not a guarantee against undiscovered vulnerabilities or a deployed exploit assessment.

Cached tagged source was inspected. evidence/source-excerpts.txt records paths/lines/hashes; evidence/pgx-source.diff and evidence/text-source.diff preserve relevant baseline→candidate diffs. pgx pool.New parses config; Pool.BeginTx returns pgx.Tx; Pool.QueryRow delegates to acquired connection; current app APIs compile against selected pins. Query mode defaults to CacheStatement, while connstring default_query_exec_mode can select simple_protocol and per-query options can override it. Current db.Open passes DATABASE_URL to pgxpool.New; configured values were not inspected. Auth store uses positional SQL parameters with no dollar-quoted query literal observed. v5.9.2 sanitizer diff adds dollar-quote-aware lexing; protocol parser diffs add malformed-message checks.

SCRAM normalization calls precis.OpaqueString.String(password), with upstream fallback to original password on normalization failure. OpaqueString uses NFC/freeform; Nickname uses NFKC and repeat. Source identifies the paths without establishing configured runtime exposure. pgx changes include SCRAM-PLUS/channel binding (default prefer), OAuth support and PostgreSQL protocol negotiation, as already anticipated by approved Techplan. No app trust configuration was changed or extra DB auth activated by this trial. Real PostgreSQL/SCRAM connection compatibility and lock contention remain independent Testing evidence. Source diff/integrity/compile do not prove those runtime behaviors.

## Preserved boundaries and identities

Backend go.mod/go.sum/LockSession remain unchanged. The stale caller-time expiry defect is not corrected by a dependency upgrade or this trial. No production gate closure, T2 completion, blocker closure, T3 readiness or task dispatch is claimed. T2-005 terminal report is preserved byte-for-byte.

Live production identities:

```text
2c089366f9b54e86da2993dadbeffe3cd7d3786a7b5f0535dd9163521e4db5d3  backend/go.mod
979a350418d3d0e4906ea8af162cb69f8cfcfd38507eaa7d8657f6c7c5d7960f  backend/go.sum
dfeed8714bc1881b9c94e8b9ac8f29ed2226f7396b1f2c789b0bc7dcfd0032b9  backend/internal/platform/auth/store.go
```

Isolated input identities after resolution:

```text
06b8dabb343d09d03f1cfa708c9870e1f064151281c7db549cbd4f08bfadfa84  backend/go.mod
6923543aab63e09da2b894699afb70b2ff1072a2eb014e063c8542ffed7deaa0  backend/go.sum
```

No race/concurrency, load/performance, configured Google/DB/browser tests or broad final security suite ran. The sole security tool run was the explicitly requested source dependency scan. Trial evidence supports the proposed cache-first method and source dependency outcome; additional full provenance/actual-backend verification and focused session/protocol/person checks remain necessary before implementation acceptance.
