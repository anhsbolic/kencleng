import pathlib,json,hashlib,difflib,sys
sys.path.insert(0,str(pathlib.Path(__file__).parent));from run import cap,env
cache=pathlib.Path(env['GOMODCACHE']);pgx=cache/'github.com/jackc/pgx/v5@v5.9.2';text=cache/'golang.org/x/text@v0.41.0'
files=[(pgx/'pgxpool/pool.go',[(205,222),(760,816)]),(pgx/'conn.go',[(183,220),(485,522),(1245,1280)]),(pgx/'pgconn/auth_scram.go',[(43,85),(181,208)]),(pgx/'pgconn/config.go',[(90,104),(455,500)]),(text/'secure/precis/profiles.go',[(10,72)]),(pgx/'CHANGELOG.md',[(1,75)])]
parts=[];ident=[]
for p,ranges in files:
 lines=p.read_text().splitlines();h=hashlib.sha256(p.read_bytes()).hexdigest();ident.append({'path':str(p),'sha256':h});parts.append('\nSOURCE '+str(p)+' SHA256 '+h+'\n')
 for a,b in ranges:parts.extend(f'{i}: {lines[i-1]}' for i in range(a,min(b,len(lines))+1))
(cap/'source-excerpts.txt').write_text('\n'.join(parts));(cap/'source-identities.json').write_text(json.dumps(ident,indent=2))
for name,old,new,rel in [('pgx','github.com/jackc/pgx/v5@v5.8.0','github.com/jackc/pgx/v5@v5.9.2',['conn.go','pgconn/auth_scram.go','pgconn/config.go','internal/sanitize/sanitize.go','pgproto3/bind.go','pgproto3/function_call.go','CHANGELOG.md']),('text','golang.org/x/text@v0.29.0','golang.org/x/text@v0.41.0',['secure/precis/profiles.go','secure/precis/profile.go','unicode/norm/iter.go','go.mod'])]:
 parts=[]
 for r in rel:
  a=cache/old/r;b=cache/new/r
  assert a.exists() and b.exists(),r
  parts.extend(difflib.unified_diff(a.read_text().splitlines(True),b.read_text().splitlines(True),fromfile=str(a),tofile=str(b)))
 (cap/(name+'-source.diff')).write_text(''.join(parts))
(cap/'compatibility-review.md').write_text('''# Source compatibility review

Current backend source was reopened (store.go, oidc.go, http.go, db.go, main.go). pgxpool.New parses the supplied connection string; Pool.QueryRow delegates through an acquired connection; BeginTx returns pgx.Tx. Existing pool/transaction/session API use compiled with the exact approved pins. SQL uses positional parameters; current auth queries contain no dollar-quoted literal. No per-query mode override was observed in app source.

pgx defaults to CacheStatement. Connection-string default_query_exec_mode can select simple_protocol; per-query options can also override mode. db.Open forwards DATABASE_URL to pgxpool.New and Ping. Configured environment/URL/password was not read, so actual mode/auth exposure is unverified. The v5.9.2 sanitizer diff includes dollar-quote handling; protocol parser diffs add malformed-message bounds checks.

SCRAM newScramClient calls precis.OpaqueString.String(password), falling back to the original on normalization failure. OpaqueString uses NFC/freeform; Nickname uses NFKC/repeat. This distinguishes source profiles from shared receiver scanner traces. No malformed/profile fixture or exploit reproduction is needed to accept removed affected ranges; configured SCRAM and real DB remain independent Testing.

pgx changes add SCRAM-PLUS/channel binding (default prefer), protocol negotiation and DB OAuth support. These changes were anticipated by the approved plan. No app config/trust/auth toggle was added. Source review + compile establish bounded source/API compatibility, not real connection/DB correctness.

OIDC verifier is configured for exact client and RS256 with skip flags at false defaults. Application checks exact Google issuer (rejects the library alias), subject, audience, azp, nonce and expiry after library signature verification. OAuth2 supplies S256 helpers. No JWE decrypt is introduced. Relevant Go/auth pins stay unchanged.

Full relevant source excerpts, file hashes and baseline→candidate diffs are retained beside this review. No newly required material dependency/security choice was identified in these inspected source paths. Runtime/provider/DB/HTTP2 exposure remains unverified.
''');print('Source excerpts, identities and relevant pgx/text diffs captured.')
