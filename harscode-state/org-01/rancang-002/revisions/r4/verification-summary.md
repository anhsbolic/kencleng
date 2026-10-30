# ORG-01 r4 — verification summary

Coordinator memiliki catatan executed static evidence ini, 30 Oktober 2026. Ini bukan runtime conformance, penerimaan Human atau izin implementasi. Final output manifest mengikat file ini; tidak ada self-hash sirkular.

## Exact input dan tooling

- Corrected [OpenAPI](openapi-handoff.yaml): SHA256 `a2d21ee3ca4ea93c7e691a53255ab825b784da00f45baeeb88115e76564df9aa`.
- [Normative companion](interface-decisions.md): SHA256 `c742908cd90c8bf454298d7da4a5b189f89dbcad4409129a7d1f47695952088b`.
- Baseline/historical evidence: Git `8a3e38a713491bfb6a6198ff9e1f5598c2749e9a`; exact source/history pins pada [source manifest](source-revisions.json).
- Actual versions: Node24.22.1, Redocly openapi-core1.34.19, openapi-typescript7.13.0, TypeScript5.9.3, Python3.14.4, jsonschema4.19.2. Dependencies terpasang hanya dibaca; tidak ada installation. Canonical `api/` Redocly2.47.0 tidak terpasang/dijalankan.

## Commands dan hasil yang dijalankan sebelum cleanup

Commands dijalankan dari root Kencleng, memakai dua script r3 yang disalin **byte-exact** ke r4 untuk execution sementara. Nama script di bawah adalah historical execution paths, bukan dependency permanent handoff.

```sh
/home/anhar-solehudin/.nvm/versions/node/v24.22.1/bin/node harscode-state/org-01/rancang-002/revisions/r4/check-contract.mjs
PYTHONDONTWRITEBYTECODE=1 python3 harscode-state/org-01/rancang-002/revisions/r4/check-examples.py
```

| Executed check | Hasil pada exact corrected input |
| --- | --- |
| Parse OpenAPI3.0.3, Redocly recommended lint, dereferenced bundle/ref resolution | PASS;0errors/0warnings,79 internal refs,0 external refs, tepat11 operasi |
| openapi-typescript generation + strict noEmit TypeScript compile | PASS;11 operation slices dan77 concrete request/response examples,0diagnostics |
| JSON Schema Draft4 + documented OAS3 nullable conversion + format checks | PASS;218 instance/schema/missing/unknown/null checks,0failures;7 request examples |
| Tambahan response assertions | PASS; required no-store/cookie headers, no-body204/303, setiap problem variant mempunyai required type/title/status dan status=HTTP |
| One-off Python JSON/UUID/equivalence assertions | PASS;18 unique canonical RFC9562 UUIDv4/variant URNs, mapping bijective/exhaustive,94 occurrences, status/title tetap,0 old identifiers dalam OAS |
| Bounded semantic regression | PASS; inverse mapping dan normalisasi **hanya** info.title/info.version mereproduksi exact r3 OAS bytes/SHA256 `741b227ca58fe2bccddf39e0c2b7ebfd54ce7fe07283045df3ba3073a63ea2eb` |

One-off assertions memakai Python3 stdlib `json`, `uuid`, `hashlib`, dan Git read-only; tidak dibuat permanent test framework. Schema checks tidak membuktikan registered-NID validity: [Architect](interface-decisions.md#api-r3-01--verified-disposition-allocation-dan-compatibility), [Security](security-findings.md) dan [Reviewer](solution-review.md) memeriksa RFC8141/RFC9457/IANA/RFC9562 secara terpisah. Historical r3 lint PASS tidak membatalkan defect API-R3-01 yang baru terkonfirmasi.

## Historical machinery dan raw evidence

Full original r3 ada pada [pinned tree](https://github.com/anhsbolic/kencleng/tree/8a3e38a713491bfb6a6198ff9e1f5598c2749e9a/harscode-state/org-01/rancang-002/revisions/r3); semua20 file dipin hash-nya dalam source manifest r4. [Original output manifest](https://github.com/anhsbolic/kencleng/blob/8a3e38a713491bfb6a6198ff9e1f5598c2749e9a/harscode-state/org-01/rancang-002/revisions/r3/output-revisions.json) SHA256 `40577b769160d2d39d6fdb9e5ba1a597ac03e635d80465f4e3cff33460db1b4b` mengikat original evidence pada kontrak r3, bukan corrected r4.

Execution-script provenance: `check-contract.mjs` SHA256 `34568c28427b15e749329eee0da005cc64e6acb6f1e2badb15335347d98232b7`; `check-examples.py` SHA256 `72b8fea8e8b6e28b6da0d2a2fa1a5e44f0387190b044bd028234f6a4630fccd2`. Keduanya dapat dibaca dengan `git show 8a3e38a713491bfb6a6198ff9e1f5598c2749e9a:harscode-state/org-01/rancang-002/revisions/r3/<filename>`. Script/layout dan versi yang sama dengan corrected r4 input merekonstruksi derived checks; tidak memerlukan artifact generated sebagai authority baru.

Transient corrected parsed/dereferenced views, generated declaration/smoke file dan verbose reports hanya digunakan untuk checks dan independent review sebelum dibersihkan. **Tidak diklaim bahwa raw corrected r4 transient bytes sudah diarsipkan Git.** Summary ini menyimpan exact input, commands, versi, hasil dan limitations; r3 raw evidence/machinery tetap immutable dan reconstructable. `check-documentation.py` r3 tidak dipakai sebagai checker r4 karena layout source manifest berubah; current source/hash/link/scope checks dilakukan inline, tanpa framework permanen.

## Final source, navigation dan cleanup checks

Executed **PASS setelah cleanup** melalui inline Python/Git read-only checks:95 source checkpoints reconciled (54 actual project files+1 navigation preimage Git8a3e38a,20 historical r3 files via Git,17 formation/runtime files,3 installed tooling metadata),21 r1/r2 historical payload/navigation checks,40 local link instances dan9 pinned historical blob links diperiksa;2 section anchors cocok,0 dangling link. Sembilan durable r4 files tepat sesuai inventory; current r3 directory tidak ada, semua20 original files recoverable dan hash-match di Git8a3e38a. Exact corrected pair tidak berubah sesudah tooling cleanup. Git changed-path set tepat30:20 retirements r3,9 additions r4 dan1 assignment navigation update; tidak ada path lain. Summary mencatat checkpoint ini; final manifest/hash reconciliation dan independent administrative addendum mengikat final assembly tanpa self-hash cycle.

## Limits

OAS lint/types/examples/UUID/equivalence adalah document evidence, bukan executable API. Python conversion adalah bounded OAS3 schema checking, bukan pengganti seluruh semantics validator runtime. Registry observation2026-10-30 tidak membuktikan deployed consumer compatibility atau global noncollision measured; UUID allocation tetap fixed dan locally unique. No tests menguji actual Google/browser cookie/CSRF/authz/transaction/DB/concurrency/unknown delivery/deletion/journal/retention/physical disposal/restore/provider/legal/inbox/rendered accessibility. Semua application/runtime conformance **NOT RUN**, PDEP-R2-01 HELD, Harness PARTIAL, Human acceptance PENDING. Canonical adoption dan protected implementation memerlukan otorisasi terpisah.
