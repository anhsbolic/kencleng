# ORG-01 / rancang-002 — Security Findings

**Standing: independent bounded design assessment selesai pada checkpoint di bawah; tidak ditemukan open material safeguard violation yang menghalangi review proposal dokumentasi.** Human acceptance PENDING. Real-person persistence/lifecycle edge **HELD** oleh PD-PRIV-01; bukan waived. Runtime/security assurance **UNVERIFIED**; seluruh application, auth, DB, browser, integration dan security tests **NOT RUN**. Penulis: Security Analyst, 29 Oktober 2026.

## Scope, independence dan exact inspection

Security mendapat fresh context, membaca originals sendiri, tidak menulis X1/O1/O2/O3/O4 producer dan hanya menulis file temuan ini. Temuan dikirim kepada Coordinator dan owning Architect; perubahan producer diperiksa ulang dari bytes aktual. Ini evidence context/source independence, **bukan filesystem isolation**. R2 broad writable-runtime deviation diterima Human untuk assignment; narrow own-file scope tetap mengikat. Tidak ada subagent, service, installation, code/API/migration/standard edit atau commit dari Security.

Baseline sumber Kencleng `d0d78fab26a976f9f48c6fc97a45c0adc5cec282`; exact Git blobs + SHA-256 pada [source-revisions.json](source-revisions.json). Workspace/local config, project/root dan applicable stack/scoped AGENTS, state/authorization, Engineering/Rancang routing, Security role agent.md/knowledge.md dan formation approval 27 Oktober dibaca. Role definition diperiksa pada pinned `47297c6f369dad4d4e4c4f90f0e81eb0310be7cf`; candidate header mendahului exact approval, bukan authority baru.

Original yang dibaca langsung dan materially dipakai:

- Product Contract ORG-01 `docs/product/outcomes/org-01/product-contract.md`, blob `da19d8e7cdb9baac8cf4af1c5e6b61763607013e`: SC01–04, REQ01–08, AC01–09, scope/negative outcomes; approval asli blob `65211210c102afd02143ab1921b2486fb600b241`. Tidak menggali historical C1/rancang-001 untuk menciptakan aturan.
- Standards index dan owning T01 backend, T02 identity-authentication, T03 authorization, T04 privacy-data-protection, T05 database-data-model, T06 api-contract, T07 testing-verification, T08 operational-reliability, T09 documentation-conventions; khususnya T02 seluruh session floor, T03 general/feature scaffold, T04 B/F/G/H/I, T05 C/E/I, T06 E/G/I/J. Original Design product-design-principles dibaca untuk consequence, truth classes, unknown dan scope decisions. Bukan pemeriksaan visual fidelity atau seluruh Design asset.
- Actual neutral evidence: `backend/cmd/server/main.go` hanya connectivity/health/lifecycle; `frontend/app/page.tsx` neutral scaffold; `api/README.md` dan `api/openapi/index.yaml` masih tooling dengan `paths: {}`. Tidak menyimpulkan auth/schema/organization delivered.

Executed hash check terhadap 31 sources manifest: 30 cocok worktree; satu authorized navigation edit `docs/engineering/README.md` dibaca lewat diff, hanya tautan assignment/status. Baseline pin tetap sah untuk source authority; tidak ada security/Product/data/API original drift yang teramati.

| Output actual yang dibaca | SHA-256 inspected checkpoint |
| --- | --- |
| [requirements.md](requirements.md), X1 | `055eb247a004a724b130c79e6a07ce662e20f520bb9f44df11802037af8ff951` |
| [solution.md](solution.md), O1/O2/TECH | `058c4f00912ea41f882982109aa20b153a6ee983237655ccb3fc9bc04ee7945c` |
| [experience.md](experience.md), O3/EXP | `475683ebe5e94b83e0930ccdbc5c4f92a4c1368da17229c802fe2b6e2fad1604` |

Reassessment final: setelah substantive O2 checkpoint `e018953000b1dba2fcb99677d5e83bc7d8c4a0e33c40fa59b092bc39fc4049ca`, actual final reciprocal alignment record dan TECH02 dibaca ulang; verification sekarang eksplisit mencakup CSRF reload/tabs/key rollover dengan same mutation key. O3 final dibaca utuh dan ketiga output hashes dicek langsung. Tidak ditemukan material changed mechanism yang membuka concern baru.

Hashes mengikat penilaian ini; perubahan material membuka ulang dependent security edges. Report ini tidak menilai handoff/review final yang belum menjadi objek assessment.

## Threat boundaries dan safeguard assessment

Assets material adalah trusted internal account/provider mapping, session/refresh authority, initial Owner assignment, current own-resource visibility, nama organisasi yang mungkin PII, serta durable pair/request receipt. Browser, submitted IDs/role/name, provider profile, callback parameters, cached UI dan transport outcome tidak trusted authority. Google authentication, Kencleng session, backend current authorization, PostgreSQL atomic effect dan browser disclosure adalah boundary berbeda.

| Boundary / plausible failure | Proposal aktual dan penilaian design | Downstream evidence wajib, semuanya NOT RUN |
| --- | --- | --- |
| Identity spoofing/account takeover | O2 login §§1–4 memvalidasi state/browser binding/nonce/PKCE/OIDC, issuer+subject mapping unique, UUID internal, tanpa email linking/profile. Kencleng JWT berbeda dari token Google; signing/CSRF keys purpose-separated. Prasyarat identity cukup konkret untuk ORG-01, bukan sekadar asumsi existing login. | TECH01/02: negative provider/JWT, replay/concurrent first-login, raw-token absence, real integration secara terpisah |
| Unauthorized Owner grant atau stale session | SOL03 dan permission inventory: create oleh effective internal account, grant dari trusted server principal; forged fields ditolak. Session row lock menyerialkan create versus logout/replay revocation; no global Owner/admin. Future transfer/revoke bukan delivered. | TECH02/03/05: effective-order race, missing/expired/revoked actor, cross-account denied, internal callers tidak bypass |
| Cross-org exposure/current authority drift | Own list/detail scoped di DB sebelum pagination; receipt account-bound + current visibility check. JWT/receipt tidak memulihkan permission; current read terpisah dari historical event. O3 membersihkan private cache/account switch dan melarang stale success sebagai current. | TECH05/06 + EXP03/04/06: raw network, unauthorized reads, current relationship change, same-name distinct IDs |
| Partial/duplicate/ambiguous effects | Satu transaction org+NOT NULL Owner+receipt, composite attribution FK; request lock/PK, immutable-name comparison, unknown COMMIT bukan success. UUIDv7 acceptance window + post-lock expiry + no early purge/tombstone/failclosed rule menjaga retry setelah receipt cleanup. 404/410 bukan cancellation; fresh-key retry tidak otomatis. | TECH03/04: real PostgreSQL failure injection/interleavings, expired-key/purge, original in-flight saat retry expiry; EXP05 truthful unknown |
| Misleading privilege/status dan unnecessary PII | Name-only, no external/legal/Reviewer claim, no public endpoint/profile/documents. Own response minimal/no-store, name tidak masuk URL/log; O3 consequence sebelum effect dan success terbatas. sessionStorage marker hanya pseudonymous account/request, bukan token/nama; storage failure tidak jadi recovery proof. | TECH06/07 + EXP01/02/05/08: DTO/storage/log/cache/aria inspection, escaping, actual UI + payload, no forbidden side effects |

Tidak dibutuhkan security framework, legal verification gate, generic recovery platform, blanket encryption semua fields, atau full role matrix untuk menilai bounded solution ini. Protected plaintext dipilih dengan alasan scope/query; ini tetap memerlukan actual access/TLS/backup controls dan production reassessment. Tidak ada Engineering Standard amendment yang diusulkan untuk melonggarkan baseline.

## Findings dan disposition

### SEC-01 — MEDIUM — CSRF token reload lifecycle sebelumnya tidak dapat direkonstruksi

**Evidence awal:** draft O2 menyimpan hanya random-token `csrf_verifier` tetapi menjanjikan `GET /auth/session` dapat mengeluarkan raw CSRF token setelah reload. One-way verifier tidak bisa merekonstruksi token. T02 CSRF/session safeguards serta T06-E memerlukan mechanism yang konsisten, bukan defense claim saja.

**Consequence:** implementer dapat menghilangkan CSRF check demi reload/refresh berfungsi, atau membuat per-tab rotation yang saling invalidasi dan memicu unsafe mutation retries. Ini design ambiguity pada session prerequisite, bukan exploit yang telah dibuktikan.

**Correction diperiksa:** actual solution checkpoint di atas mengganti storage menjadi random session nonce 256-bit; backend derives HMAC-SHA256 dengan dedicated key/domain separator/session UUID/nonce, constant-time compare, token stable lintas reload/tabs/refresh. Revocation mematikan session; rollover refetch session sebelum explicit retry dengan mutation key sama. Inventory/schema dan O3 diperbarui. Security tidak menulis correction.

**Disposition: ADDRESSED AT DESIGN LEVEL.** Wajib verifikasi TECH02 mencakup reload/refresh/concurrent tabs, CSRF-key rollover, wrong session/origin/token dan revoked session; tidak cukup screenshot atau mock. Key custody/rollover implementation tetap protected authorization + security review tersendiri. Runtime closure **NOT VERIFIED**.

### SEC-02 — MEDIUM — OPEN EXTERNAL OWNER DECISION: PD-PRIV-01

**Original evidence:** T04-B/G/H melarang Engineering mengarang legal basis, retention/deletion/erasure exceptions; Product ORG-01 memberi tujuan durable attribution/inspection tetapi tidak menetapkan lifecycle tersebut. Data minimal tetap person-linked: provider issuer+subject, account/Owner association, organization name/creation time, request receipt dan session/verifier history/copies.

**Output evidence:** O2 Actual-field inventory + PD-PRIV-01 secara tepat menyebut fields, owning Human Product/privacy concern, alternatif bounded attribution lifecycle versus disposable fake-data sandbox, dan affected hold. O3 tidak mengklaim consent/retention/readiness. Technical token/retry TTL bukan approval business retention.

**Consequence jika dilompati:** real-person collection dapat berlangsung tanpa effective lifecycle/basis/notice; deletion dapat merusak attribution atau membuka duplicate retry, sementara backups/caches mempertahankan data dengan klaim erasure palsu.

**Required decision/verification:** Human owner menetapkan basis/notice, inspectable retention, deletion/erasure exceptions, session/receipt security retention, backup/copy expiry pada owning source dengan exact revision. Reconcile T04 inventory, FK/receipt-window cleanup dan restore behavior; reopen TECH06/09 + Security/Reviewer setelah decision. Tidak memilih indefinite retention atau automatic consent sebagai default.

**Disposition: OPEN, AFFECTED EDGE HELD.** Ini dependency eksternal yang sudah diidentifikasi secara cukup konkret untuk bounded design completeness; **tidak memblokir review/penerimaan proposal dokumentasi dengan batas ini tetap eksplisit**. Real-person collection/persistence release, lifecycle/deletion/cleanup implementation yang bergantung keputusan tersebut, dan privacy-readiness claim tetap ditahan. Fake-data tests/design reasoning dapat maju sesuai separate authorization. Tidak ada policy waiver atau authority efektif yang dibuat oleh proposal/rekomendasi ini.

## Verdict dan limits

Pada exact checkpoint, tidak ditemukan open HIGH/CRITICAL atau material unresolved safeguard defect untuk proposal dokumentasi scoped. SEC01 telah dikoreksi pada design; SEC02 tetap open owner edge dan harus tetap terlihat di handoff/acceptance. No finding bukan certification bahwa implementasi aman.

Identity/session prerequisite depth, attribution/Owner grant, current authz, data minimization, transaction/retry/unknown dan truthful exposure telah dinilai terhadap originals dan actual producer outputs. Generic deeper assessment tidak diwajibkan untuk menyelesaikan dokumen ini. Saat protected implementation dipilih, targeted security verification T02–T07 dan key/session/privilege implementation review wajib berdasarkan actual code/topology; public/real-user usage juga membuka T08 reassessment. Tidak mengklaim arbitrary numeric TTL, RS256/library/key custody atau deployment aman hanya karena tercantum sebagai proposal.

Seluruh runtime tests **NOT RUN**. Canonical API adoption, separate protected implementation authorization, Human bundle acceptance, downstream consumer fit/discovery dan production readiness tidak diberikan oleh assessment ini. Scope di luar ORG-01 (legal verification, full role matrix, transfer/recovery, Reviewer/Campaign/donation) tidak dinilai. Perubahan pada effective sources, data/lifecycle authority, session/CSRF, request identity/cleanup, authorization atau observed result semantics membuka ulang review pada affected edge.
