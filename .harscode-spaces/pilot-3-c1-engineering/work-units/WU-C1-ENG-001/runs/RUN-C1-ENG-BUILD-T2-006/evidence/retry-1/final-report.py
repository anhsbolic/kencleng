import pathlib,json,hashlib,shutil,subprocess,sys
base=pathlib.Path(__file__).parent;root=pathlib.Path.cwd();dst=root/'.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-BUILD-T2-006/evidence/retry-1';act=base/'actual';iso=base/'capture'
sys.path.insert(0,str(base));from inventory import stream
normalize=lambda p:{o['Path']:o.get('Version',o.get('GoVersion')) for o in stream(p.read_text())}
assert normalize(act/'final-graph.stdout')==normalize(iso/'graph.stdout');assert (act/'final-edges.stdout').read_bytes()==(iso/'edges.stdout').read_bytes()
assert json.loads((act/'post-g1-scan-summary.json').read_text())['finding_count']==0
assert not (act/'fixture-cleanup.stdout').read_text().strip()
old=dst.parent/'build-report.md';assert hashlib.sha256(old.read_bytes()).hexdigest()=='4fc08a8cb5533ebf702d364c86be11ce06b0937837ed8941cac828d5e3436191'
for p in act.iterdir():shutil.copy2(p,dst/'actual'/p.name)
for p in base.glob('*.py'):shutil.copy2(p,dst/p.name)
changed=['backend/go.mod','backend/go.sum','backend/internal/platform/auth/store.go','backend/internal/platform/auth/store_test.go','backend/internal/platform/auth/oidc_test.go','backend/internal/platform/auth/store_integration_test.go']; identities=[{'path':p,'sha256':hashlib.sha256((root/p).read_bytes()).hexdigest()} for p in changed];(dst/'changed-file-hashes.json').write_text(json.dumps(identities,indent=2))
for n in ['go.mod','go.sum']:shutil.copy2(root/'backend'/n,dst/'actual'/('final.'+n))
source=[{'path':str(p.relative_to(root)),'sha256':hashlib.sha256(p.read_bytes()).hexdigest()} for p in sorted((root/'backend').rglob('*.go'))];(dst/'final-source-hashes.json').write_text(json.dumps(source,indent=2))
(dst/'final-diff.patch').write_text(subprocess.check_output(['git','diff','--','backend/go.mod','backend/go.sum','backend/internal/platform/auth/store.go'],text=True));(dst/'final-status.txt').write_text(subprocess.check_output(['git','status','--short'],text=True));subprocess.run(['git','diff','--check'],check=True)
# Enrich command captures with execution context and output content identities.
for folder in [dst/'isolated',dst/'actual']:
 commands=[json.loads(l) for l in (folder/'commands.jsonl').read_text().splitlines()]
 for c in commands:
  name=c['name'];escalated=(name in {'missing-sums','source-scan','post-g1-source-scan','focused-postgres','fixture-port','fixture-version','fixture-image','fixture-stop','fixture-cleanup'} or name.startswith(('origin-','ref-','commit-')) or c.get('network'))
  c['harness_context']='require_escalated' if escalated else 'sandbox'
  for streamname in ['stdout','stderr']:
   p=folder/(name+'.'+streamname)
   if p.exists():c[streamname+'_sha256']=hashlib.sha256(p.read_bytes()).hexdigest()
 (folder/'commands-context.jsonl').write_text(''.join(json.dumps(c)+'\n' for c in commands))
# Preserve isolated capture manifest from before backend write, and separately index final evidence.
manifest=[{'path':str(p.relative_to(dst)),'sha256':hashlib.sha256(p.read_bytes()).hexdigest()} for p in sorted(dst.rglob('*')) if p.is_file() and p.name not in {'build-report.md','capture-hashes.json'}];(dst/'capture-hashes.json').write_text(json.dumps(manifest,indent=2))
mods=json.loads((iso/'inventory.json').read_text());legacy={m['module']:m['commit'] for m in json.loads((iso/'legacy-source-match.json').read_text())};module_rows='\n'.join('| `'+m['path']+'` | '+m['version']+' | '+m['go']+' | `'+(m['origin']['Hash'] if m['origin'] else legacy[m['path']])+'` |' for m in mods)
hashes='\n'.join(x['sha256']+'  '+x['path'] for x in identities)
summary=json.loads((act/'post-g1-scan-summary.json').read_text());conf=summary['config'];dbdate=conf.get('db_last_modified','not exposed');fixture=(act/'fixture-version.stdout').read_text().strip();image=(act/'fixture-image.stdout').read_text().strip()
report=f'''# Build retry report — RUN-C1-ENG-BUILD-T2-006 / retry-1

> Phase: Build retry; Author / Role: Implementer T2/G1 in the current conversation; WU `WU-C1-ENG-001`; original Participant/Session bindings `PARTICIPANT-C1-ENG-T2-IMPLEMENTER-006` / `SESSION-C1-ENG-T2-IMPLEMENTER-006`, retained as bindings, not a new execution identity. Created/Updated: 2026-10-11 Asia/Jakarta. Target HEAD `059d145d5df64fcda427e5a625486a8e95ce1fd8`; workflow HEAD `91a97af8ba64f98b0a692a83fbff866a210c45c0`. Original Invocation selected/approved `gpt-6-sol` / `medium`; actual runtime model/effort not independently established. Latest Human instruction explicitly authorized retry in this session after the network-recovery discussion. Same-session retry and separate result carrier are explicit Human exceptions to ordinary Run/session/path conventions; no new model approval/Run/Participant identity is invented.

## What changed

**COMPLETED within the authorized retry scope:** isolated D8 including provenance passed before backend writes; actual backend graph acceptance passed before G1; post-lock expiry correction and authored focused checks passed. No commit or downstream dispatch. Work Unit completion, blocker closure and T3 readiness are not promoted by this report; Orchestrator owns reconciliation and independent Review/Testing/Human evidence remains.

- `backend/go.mod` / `go.sum`: exact pgx v5.9.2 + explicit text v0.41.0, resolved sync v0.22.0; selected mod/tools graph versions also change to v0.38.0/v0.48.0. Go1.26.9, OIDC v3.16.0, OAuth2 v0.34.0, jose v4.1.4 retained. After authored tests import jose, tidy promotes its existing exact requirement from indirect to direct; no new version/module or transitive pin.
- `store.go`: public `LockSession` preserves its signature, but expiry uses `time.Now` sampled after successful locked-row Scan. Caller timestamp remains accepted for compatibility and is explicitly ignored for expiry authorization. A private clock seam makes acquisition ordering deterministic in unit checks. Revocation, person attribution, digest-only lookup, error chains and lock/query behavior remain unchanged. No other production auth/DB/HTTP behavior was modified.
- `store_test.go`: simulated delayed-lock expiry, exact expiry boundary, revoked/missing/storage error/valid attribution, public stale-caller regression, malformed state/browser/session rejection before storage.
- `oidc_test.go`: production `NewGoogleProtocol` construction and exchange against simulated discovery/JWKS/token transport with signed fixtures; forged issuer/audience/azp/signature/algorithm, expiry/nonce/key/token/provider failure; S256 PKCE/missing input checks. Valid fixture maps only issuer/subject, ignoring extra email/domain/person claims. No real Google login claim.
- `store_integration_test.go` (`integration` tag): serial person-upsert and pending-login binding/replay/expiry checks using the real migration in disposable PostgreSQL schemas. Test URL comes solely from `C1_AUTH_TEST_DATABASE_URL`; no application .env. Schema identifiers are generated and sanitized, values parameterized; bounded contexts and cleanup.

Only six backend dependency/source/test files above belong to this retry. Previous Human-requested network guidance edits and existing Orchestrator/WU/events/control-tower changes were preserved. No upstream Product, API, Techplan/task/manifest, frontend or orchestration state was edited. Original terminal BLOCKED report SHA-256 `4fc08a8cb5533ebf702d364c86be11ce06b0937837ed8941cac828d5e3436191` remains byte-for-byte unchanged; this is separate Human-authorized retry evidence, not a rewrite or silently reopened result.

## Authority and sequence

`authorization.md` preserves the current Human retry instruction and narrow recovery/scope interpretation. All 25 Invocation-bound identities were rechecked; only root/backend AGENTS and scanner runbook have known Human-authorized network-guidance deltas. Current network/dependency guidance applies. Production baseline inputs matched before retry; target/workflow HEAD matched. Techplan §§9–10/12, T2 scope and live auth/DB/server/module anchors were re-grounded. Original pins, G1 permission and semantic stop gates remain.

1. Fresh probe copied cmd/internal/go.mod/go.sum only, no secrets. Cached inputs supported tidy immediately; no initial sandbox checksum failure recurred. Only six graph module archive sums absent from go.sum required targeted authenticated download metadata. No full download-all was forced.
2. All isolated provenance/graph/source/minimum/delta/integrity/compile/fresh-scan/compatibility evidence was durably captured in `isolated/` and `isolated-gate.md` **before tracked backend dependency writes**.
3. Actual backend pin update, tidy, graph/edges/integrity/compile and fresh scan completed. Actual source and dependency inputs matched the isolated evidence at this stage. `actual-gate.md` was preserved **before G1 production correction**.
4. G1 correction and focused tests executed. A public-API regression was also run on copied original source (expected failure) and copied patched source (pass). This deliberate negative control is not an acceptance check failure; it demonstrates the test detects the known defect.
5. Post-G1 fresh scan and per-ID reconciliation passed. Final tidy only changes existing jose directness; final selected graph and all edges remain identical to the scanned/accepted graph. Final integrity and diff checks passed. Fixture stopped and auto-removed; container lookup output empty.

Analysis helper corrections are recorded without hiding failures: the advisory parser initially could not read a prerelease affected-range boundary, then was corrected; the actual evidence wrapper initially parsed before inventory was copied, then capture order was corrected. Neither changed scanner output, selected dependencies or required disposition. All subsequent structured parsing/assertions completed successfully. These helper errors are distinct from a substantive gate finding or failed source scan.

## Environment and durable commands

Explicit Go `/home/anhar-solehudin/.gvm/gos/go1.26.9/bin/go`, `GOTOOLCHAIN=local`, PATH Go1.26.9 prefix. Go version go1.26.9 linux/amd64. Scanner `/home/anhar-solehudin/.gvm/pkgsets/go1.26.9/global/bin/govulncheck` v1.7.0 built with Go1.26.9. Writable GOCACHE `{base}/build-cache`; GOMODCACHE `/home/anhar-solehudin/.gvm/pkgsets/go1.24.8/global/pkg/mod`; observed final GOPATH `/home/anhar-solehudin/.gvm/pkgsets/go1.26.9/global`. GOPROXY=off for cache-first acquisition; GOSUMDB=sum.golang.org retained. No checksum bypass, global cache deletion, permission change, version substitution, database substitution or finding suppression.

Probe cwd `{base}/probe`; actual cwd `{root}/backend`. Exact argv/cwd/env/time/exit/harness context and output hashes are in `isolated/commands.jsonl` and `actual/commands.jsonl`, with full stdout/stderr. Network-dependent targeted checksum, upstream refs/source/commit API, source scans and disposable fixture commands used harness escalation. Other Go graph/integrity/compile/unit commands ran sandboxed with writable build cache. Runtime application credentials/URLs were not read.

All relied-upon temporary captures/wrappers were copied into this retry directory. `capture-hashes.json` identifies evidence, `changed-file-hashes.json`/`final-source-hashes.json` identify backend results, `final-diff.patch` preserves production/dependency delta. The original isolated capture-hashes file remains its pre-write snapshot; the retry-wide manifest identifies the later command-context enrichment separately.

## Gates and dependency evidence

| Gate | Disposition |
|---|---|
| Live authority/baseline identities | PASS; known approved documentation deltas explicitly adopted |
| Isolated exact D8 graph, full transitive delta, Go minima | PASS; 20 external modules, five version deltas, no additions/removals/x/net |
| Full graph origin/tag/commit/source/checksums | PASS; metadata/cache hashes, live upstream tags, legacy archive matches and pseudo-version timestamps |
| Isolated integrity + compile | PASS; all modules verified; auth/db/server compiled before scan |
| Fresh isolated source scan + all required advisory dispositions | PASS; 0 finding objects, 55 required IDs reconciled |
| pgx API/query-mode/config/SCRAM-OpaqueString source compatibility | PASS in source/compile scope; configured runtime deferred |
| Actual backend graph/source/origin/integrity/compile/fresh scan | PASS; graph/source identities explicitly matched, separate actual source scan |
| Bounded G1 post-lock expiry and focused checks | PASS; unit/protocol plus serial PostgreSQL person/login tests |
| Post-G1 scan, final graph/integrity/diff and fixture cleanup | PASS |
| Independent Review/Testing, configured Google/DB/SCRAM/privileges/contention | DEFERRED / NOT ESTABLISHED |

Complete selected graph and source origins:

| Module | Version | Declared dependency Go minimum | Upstream commit |
|---|---|---|---|
{module_rows}

Maximum declared dependency minimum is Go1.25.0, compatible with Go1.26.9. Modules with omitted graph GoVersion fields were checked in cached version-specific go.mod; true legacy absence is labeled explicitly. Exact cache metadata and archive/source identities are in `isolated/inventory.json`.

All twelve Origin-backed tag refs matched live upstream peeled commits. The pgservicefile pseudo-version commit and check.v1 pseudo-version commit timestamps match upstream committer dates. Seven legacy modules without Origin have upstream refs/commits plus **133 module archive files matching upstream commit archives byte-for-byte**. Archive h1 hashes match ziphash and all present go.sum entries; six omitted archive entries have authenticated exact-version `go mod download -json` Sum/GoModSum matching computed cache hashes, with no structured Error. Missing go.sum archive entries for graph-only modules are not fabricated or silently interpreted as missing verified source. Cache authenticity, upstream origin and integrity remain separate evidence.

Actual graph/source/edges were checked against accepted isolated content identities, allowing reuse of exact source/origin checks without claiming the isolated scan validates actual source. Actual and post-G1 scans are fresh separate commands. Final jose directness changes no selected module/source/edge. No new material dependency/trust choice was needed.

## Source scan and advisory disposition

All three fresh full source scans (isolated, actual before G1, actual after G1) finished exit 0 with empty stderr and parsed complete streams: config 1, SBOM 1, progress 2, OSV metadata 200, **finding 0**. Scanner v1.7.0, Go1.26.9, source/symbol mode, https://vuln.go.dev; final recorded database modification `{dbdate}`. Metadata OSV objects are not source findings. No finding traces exist because no finding objects were emitted. This does not establish exploit reproduction or configured runtime security.

`isolated/advisory-dispositions.json`, `actual/advisory-dispositions.json` and `actual/post-g1-advisory-dispositions.json` contain **each of the prior 50 Go/jose IDs plus the five pgx/text IDs**, fresh affected/fixed ranges, selected/absent modules and disposition. Selected stable versions are outside applicable ranges or the named module is absent. No unreviewed surviving/new finding or risk waiver.

| Required ID | Selected module/version | Removed affected range / disposition |
|---|---|---|
| GO-2026-4771 | pgx/v5 v5.9.2 | fixed v5.9.0; selected version outside range |
| GO-2026-4772 | pgx/v5 v5.9.2 | fixed v5.9.0; selected version outside range |
| GO-2026-5004 | pgx/v5 v5.9.2 | fixed v5.9.2; other named major module paths absent |
| GO-2026-5970 | text v0.41.0 | fixed v0.39.0; selected version outside range |
| GO-2026-6629 | text v0.41.0 | fixed v0.41.0; selected version outside range |

Compatibility/source review retains API/query/config excerpts, file hashes and relevant pgx/text release/source diffs. Pool.New/BeginTx/QueryRow and current positional session queries compile. Default query mode is CacheStatement, but URL/per-query options can choose simple protocol; configured mode was not inspected. Current queries contain no dollar-quoted literal. SCRAM calls OpaqueString.String with original-password fallback on normalization error; OpaqueString uses NFC/freeform, Nickname uses NFKC/repeat. Shared-method traces do not prove the profile-specific trigger. Approved pgx feature changes (channel binding default prefer, protocol negotiation, OAuth support) were reviewed; no application config toggle was added. Real PostgreSQL/SCRAM compatibility and deployed exposure remain independent.

## Tests run

All Go commands use the explicit Go1.26.9 executable/env above; patterns/argv and outputs are retained.

| Command/check | Result and scope |
|---|---|
| `go mod verify` isolated, actual, post-G1 and final | PASS: all modules verified |
| `go test -run '^$' ./internal/platform/auth ./internal/platform/db ./cmd/server` isolated/actual | PASS: focused compile; pre-authoring packages had no tests |
| `govulncheck -json ./...` isolated/actual/post-G1 | PASS: complete stream, 0 findings; explicit dependency gate |
| `go test -count=1 -v ./internal/platform/auth ./internal/platform/db ./cmd/server` | PASS, final authored unit/protocol checks; db/server compile with no test files |
| `TestLockSessionExpiredDuringLockWait` | PASS: fake clock advances across Scan barrier, clock cannot be sampled before simulated acquisition |
| `TestLockSessionRevalidation` | PASS: expiry equality/past, revoked/missing/query error and valid attribution |
| `TestLockSessionIgnoresStaleCallerTime` | PASS on public API; original-source control intentionally FAILS, patched-source control PASSES |
| `TestGoogleProtocolNegatives` | PASS: valid control and 14 negatives through production constructor, simulated discovery/JWKS/token HTTP and signed fixtures |
| `TestGoogleProtocolPKCEAndMissingExchangeInputs`, malformed browser authority | PASS: S256 bindings and incomplete/malformed input rejection |
| `go test -count=1 -tags=integration -run '^(TestPersonUpsertPostgreSQL\\|TestPendingLoginBindingAndReplayPostgreSQL)$' -v ./internal/platform/auth` | PASS: two focused serial PostgreSQL tests, no skips |
| `git diff --check` | PASS |
| Fixture stop and scoped container lookup | PASS; disposable container absent |

Disposable fixture: `{fixture}`. Image `docker.io/library/postgres:16-alpine`, image content ID `{image}`; temporary container `c1-t2-006-retry1-pg`, ephemeral loopback port, fixture-only trust auth, no volume/production data. Schemas created by tests were cleaned, pool contexts bounded and container auto-removed. Person-upsert checks prove stable case-sensitive issuer/subject uniqueness and parameter safety on this real isolated migration/DB. Login checks prove correct state/browser matching, consumed local flow cannot replay, and expiry rejection on this fixture. This is not real Google authorization-code replay evidence, configured SCRAM, restricted runtime-role acceptance or concurrent first-login proof. Fixture uses provisioning credentials; no privilege gate is claimed.

## Verification scope confirmation

No race detector/concurrency or performance/load suite ran. Lock wait was deterministic simulation, not real PostgreSQL contention. Protocol negatives are the explicitly required focused functional checks with fake provider/signing fixtures. The only security tool was the explicitly required dependency source scan; no broad final security sweep or full `make verify`/Testing-owned suite ran. Serial disposable DB checks ran because authored person/login checks needed execution against the actual migration; final configured/runtime/concurrency correctness remains Testing-owned.

## Contract check

- [x] Authorized retry target fulfilled: complete isolated D8, actual backend acceptance, bounded G1 correction and focused evidence.
- [x] Required pre-write and pre-G1 gate sequence preserved in this retry.
- [x] No new material authority/interface/risk contradiction silently repaired.
- [x] Exact pins and post-lock expiry/revocation/person semantics preserved; caller API compatibility retained.
- [x] Existing edits and original terminal report preserved; no commit/downstream dispatch or automatic blocker/readiness promotion.

## Changed-file identities

```text
{hashes}
```

## Deferred / not tested here

Real configured Google login/browser/HTTPS-cookie behavior; deployed DB authentication/SCRAM and restricted credentials; concurrent person upsert, session logout/revoke/confirmation contention, final integration/race/security/browser/Human semantic acceptance. Tests here do not establish those outcomes. T3/T5 routing and durable blocker acceptance remain Orchestrator-owned; original terminal Run history is preserved.

## Flagged for Review / Orchestrator / Testing

No new substantive Finding/Decision/Blocker from retry gates. Review should inspect retained public timestamp compatibility, private post-Scan clock and regression evidence, exact dependency/provenance/fresh scan captures and focused test scope. Orchestrator must reconcile this explicit same-session retry with existing terminal history rather than replacing the original result. B-T2-002/B-T2-003 remain durably open until their owning acceptance/reconciliation, even though this report supplies the requested bounded implementation evidence; independent runtime limitations remain.

## Phase handoff

- **Outcome:** COMPLETED — Human-authorized retry scope; isolated/actual D8 and bounded G1/focused checks passed, no commit.
- **Result refs:** this `retry-1/build-report.md`; `authorization.md`, `isolated-gate.md`, `actual-gate.md`, `capture-hashes.json`, `changed-file-hashes.json`, `final-source-hashes.json`, `final-diff.patch`; actual/isolated command/scan/provenance captures.
- **Findings:** none new; known F-T2-003 corrected with focused evidence. Original BLOCKED occurrence/deviation remains historical evidence unchanged.
- **Decision requests:** none for implementation; Orchestrator reconciles authorized retry/terminal history and owning blocker acceptance.
- **Blockers:** no gate blocked this retry. Existing B-T2-002/B-T2-003 not directly mutated; no T3 readiness claim.
- **Open / unverified:** independent Review/Testing, configured Google/browser/DB/SCRAM/restricted-role and real contention evidence.
- **Recommended continuation:** Orchestrator reconcile retry report, then independent Code Review of dependency/G1 diff and required Testing; do not infer automatic downstream dispatch or commit permission.
- **Context refs:** original Invocation + `authorization.md`; current Approved Techplan §§9–10/12 and T2 task; changed backend files/tests; isolated/actual gate and scan/provenance evidence; original terminal report.
'''
(dst/'build-report.md').write_text(report);print('retry report',dst/'build-report.md');print('report SHA256',hashlib.sha256((dst/'build-report.md').read_bytes()).hexdigest())
