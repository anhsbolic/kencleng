# Source compatibility review

Current backend source was reopened (store.go, oidc.go, http.go, db.go, main.go). pgxpool.New parses the supplied connection string; Pool.QueryRow delegates through an acquired connection; BeginTx returns pgx.Tx. Existing pool/transaction/session API use compiled with the exact approved pins. SQL uses positional parameters; current auth queries contain no dollar-quoted literal. No per-query mode override was observed in app source.

pgx defaults to CacheStatement. Connection-string default_query_exec_mode can select simple_protocol; per-query options can also override mode. db.Open forwards DATABASE_URL to pgxpool.New and Ping. Configured environment/URL/password was not read, so actual mode/auth exposure is unverified. The v5.9.2 sanitizer diff includes dollar-quote handling; protocol parser diffs add malformed-message bounds checks.

SCRAM newScramClient calls precis.OpaqueString.String(password), falling back to the original on normalization failure. OpaqueString uses NFC/freeform; Nickname uses NFKC/repeat. This distinguishes source profiles from shared receiver scanner traces. No malformed/profile fixture or exploit reproduction is needed to accept removed affected ranges; configured SCRAM and real DB remain independent Testing.

pgx changes add SCRAM-PLUS/channel binding (default prefer), protocol negotiation and DB OAuth support. These changes were anticipated by the approved plan. No app config/trust/auth toggle was added. Source review + compile establish bounded source/API compatibility, not real connection/DB correctness.

OIDC verifier is configured for exact client and RS256 with skip flags at false defaults. Application checks exact Google issuer (rejects the library alias), subject, audience, azp, nonce and expiry after library signature verification. OAuth2 supplies S256 helpers. No JWE decrypt is introduced. Relevant Go/auth pins stay unchanged.

Full relevant source excerpts, file hashes and baseline→candidate diffs are retained beside this review. No newly required material dependency/security choice was identified in these inspected source paths. Runtime/provider/DB/HTTP2 exposure remains unverified.
