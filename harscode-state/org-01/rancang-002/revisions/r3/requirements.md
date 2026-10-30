# ORG-01 / rancang-002 / r3 — Engineering Obligation Map (X1)

**Standing: delta/coverage untuk review dokumentasi; Human acceptance PENDING.** Requirements Analyst. Batas [state r3](state.md): deterministic OpenAPI 3.0.3 dan normative conformance handoff untuk 11 operasi efektif r1+r2. X1 ini tidak memilih architecture, wire semantics, final notice, legal basis, atau runtime outcome. Architect memiliki keputusan interface/O2/TECH; Experience memiliki consumer/actor fit/O3/EXP; Security dan Solution Reviewer menilai final bytes secara independen.

## Source binding dan standing

Baseline proyek `4db1612281a311e188722af984efb163783435b1`, formed definition `47297c6f369dad4d4e4c4f90f0e81eb0310be7cf`. Workspace/local routing, project/scoped standards AGENTS, own `agent.md`+`knowledge.md`, state/manifest dan originals dibaca. Header candidate lama pada role bukan pengganti exact approved formation; formation, invocation, containment, acceptance dan runtime readiness tetap terpisah.

**Pemeriksaan independen aktual:** SHA-256 dihitung terhadap seluruh **55/55** actual files pada [source-revisions.json](source-revisions.json); semua cocok, **0 drift** pada checkpoint X1. Tambahan manifest juga diperiksa aktual: **17/17 Harscode definition files dan 3/3 read-only tooling package metadata cocok**, tanpa drift. External technical URLs adalah references, tidak diberi klaim hash atau legal verification oleh X1. Ini pemeriksaan bytes, bukan klaim semua file telah diaudit semantik. `current.md` dipin sebagai navigasi pre-r3 dan boleh diperbarui Koordinator sebagai successor setelah assembly; perubahan navigasi itu bukan perubahan production source.

| Original / standing | Exact revision |
| --- | --- |
| [Product Contract §§1–7](../../../../../docs/product/outcomes/org-01/product-contract.md): binding scoped creation, seluruh 4 SC/8 REQ/9 AC | Blob `da19d8e7cdb9baac8cf4af1c5e6b61763607013e`; SHA-256 `229e87380e28ed868167460198d7682bddf024c93aeb528e32b15f2ae0bdfcb3` |
| [Approval](../../../../../docs/product/outcomes/org-01/approval.md): exact original approved content, bukan Engineering acceptance | Blob `65211210c102afd02143ab1921b2486fb600b241`; SHA-256 `98a4ecaf618dcf64f023df1c78beafaad05961c24f8687f51266e5e85ea737d1`; original content blob `448a8082bcb9ac1c7e41a686f2e3fdd2e3988ac4`, audit 26 normative rows |
| [PD-PRIV-01-r4](../../../../../docs/product/outcomes/org-01/privacy-decision.md): approved principles/simple closure/technical retention/candidate-basis direction pada batasnya | SHA-256 `76dcc93116fd142433d1e3a86bbcf8726d0011f903b51cbea48697f352fa4b5f` |
| [Notice v0.1](../../../../../docs/product/outcomes/org-01/privacy-notice-draft.md): DRAFT, bukan final/publication/legal validation | SHA-256 `b1ebfcbadea720f48892a6add959bc41bc303f3afe35725f26b0c99df13373d8` |
| [T01–T09/front-end index](../../../../../docs/engineering/standards/README.md) dan owning originals; terutama [T06-A–J](../../../../../docs/engineering/standards/api-contract.md) | Exact file hashes/blobs pada manifest r3; T06 SHA-256 `e9a8cb6517fdfdbaf0da0fcfc3a71f2253c8e0c0dd315698bb3e1ba055f76e9f` |
| [X1 r1](../../requirements.md), [X1 r2](../r2/requirements.md): original scoped translation evidence; carry-forward ditentukan di bawah | r1 `055eb247a004a724b130c79e6a07ce662e20f520bb9f44df11802037af8ff951`; r2 `5bf0175545fa7c6617759aeb020ed97a14c8fe722b4385745ddcab4645bf8da7` |
| [Solution r1](../../solution.md) + [r2 exact carry-forward/replacements](../r2/solution.md): effective prior producer design evidence, bukan accepted bundle otomatis | r1 `058c4f00912ea41f882982109aa20b153a6ee983237655ccb3fc9bc04ee7945c`; r2 `c284a9811f098127ea12873e08543e1a956f26a708f214fc134ef0275dc51acc` |

Tidak mencari hidden requirements dari upstream research, historical code atau previous findings. Isi r1/r2 tetap immutable. Original Product/approval/PD-r4, r2 X1 dan effective r1+r2 solution diperiksa langsung; X1 tidak menganggap r2 review atau source hash sebagai assurance r3.

## Carry-forward tepat dan coverage Product

X1 r3 = carry-forward berikut + delta r3. Semua baris, non-guarantees, consumer implications dan planned verification **X1-01..08** pada tabel obligation map r1 sebagaimana diteruskan bagian r2 “Carry-forward yang tepat: creation tidak berubah” tetap berlaku. Mapping di bawah menjaga original IDs; teks ini tidak mengganti owning Product criteria.

| X1 | Original REQ / AC / SC dan invariant |
| --- | --- |
| X1-01 | REQ-ORG-01-01 / AC-ORG-01-01 / SC-ORG-01-04: boleh mulai; tidak ada klaim representasi legal terbukti |
| X1-02 | REQ-ORG-01-02 / AC-ORG-01-02 / SC-ORG-01-01: actor dan Owner awal teratribusi saat efek efektif |
| X1-03 | REQ-ORG-01-03 / AC-ORG-01-03 / SC-ORG-01-02: organisasi/akibat/batas Owner dipahami sebelum efektif |
| X1-04 | REQ-ORG-01-04 / AC-ORG-01-04 / SC-ORG-01-01: organisasi dan Owner sah/efektif bersama; partial bukan full success |
| X1-05 | REQ-ORG-01-05 / AC-ORG-01-05 / SC-ORG-01-02: current relationship dapat diperiksa durable selama aktif, terpisah dari pesan sukses |
| X1-06 | REQ-ORG-01-06 / AC-ORG-01-06 dan AC-ORG-01-09 / SC-ORG-01-04: claim provenance, tanpa Reviewer/legal/Campaign/publication/donation promotion |
| X1-07 | REQ-ORG-01-07 / AC-ORG-01-07 / SC-ORG-01-03: failed/partial/unattributed/interrupted/unknown bukan sukses; unknown bukan bukti gagal pasti |
| X1-08 | REQ-ORG-01-08 / AC-ORG-01-08 / SC-ORG-01-03: dua actor/nama sama/ID berbeda dapat sama-sama berhasil, tanpa conflict-block/merge/auto-close |

Keempat SC wajib bersama. Deliberate verified privacy closure tidak membatalkan kejujuran creation terdahulu dan tidak sama dengan auto-close karena nama mirip. Seluruh **4 SC, 8 REQ, 9 AC** tetap tercakup, tanpa eligibility atau Reviewer prerequisite baru.

Bagian X1 r2 **“Delta approved privacy → kewajiban consumer dan verification” seluruh X1-P01..10** diteruskan utuh berikut seluruh owning anchors dan planned evidence, tidak digantikan ringkasan ini:

| ID | Obligation yang tetap mengikat r3 |
| --- | --- |
| X1-P01 | Minimum actual-data inventory; provider extra claims bukan izin simpan/expose; tujuan utama tidak memberi indefinite retention |
| X1-P02 | Verified requester, seluruh dependent organizations dan dampak dipahami sebelum explicit closure; nol org sah; tidak ada active orphan |
| X1-P03 | No auto-transfer; continue/other users/Campaign/financial obligations di luar simple flow ditahan dan ditangani owner, tidak menggantung tanpa batas |
| X1-P04 | Session/authority ended dan erasure sebenarnya; bedakan accepted/access-ended/complete/unknown dan residual copies |
| X1-P05 | Successful receipt 48h sejak sukses, retry24h berbeda; cleanup tidak membuka old-key creation; active own-list/detail tetap dapat dipakai |
| X1-P06 | Session/refresh hanya selama security need, maksimal family30d+cleanup24h, ended lebih awal dibersihkan lebih cepat; retensi bukan validity |
| X1-P07 | Necessary logs14d sejak event, tanpa nama/profile/token/full payload; semua copies perlu diperhitungkan |
| X1-P08 | Backup maksimal umur14d seluruh copies/chains; sebelum restore release reapply erasure/revocation; minimum marker bukan replacement profile |
| X1-P09 | Candidate bases untuk divalidasi, applicability OPEN; Google bukan blanket consent; Engineering tidak membuat legal determination |
| X1-P10 | Pre-Google summary+full notice access, pre-create Owner consequence; controller Yayasan Kencleng Indonesia/contact care@kencleng.dev given; vendor/location/final notice/request readiness OPEN |

Carry-forward lintas concern: r1 **“Constraints efektif lintas-kewajiban” seluruh 13 baris**, **“Negatif, unknown dan batas disclosure” seluruhnya**, dan delegated freedom/gaps sebagaimana replacement eksplisit r2 **“Gaps yang tetap bermakna dan affected gates”** tetap berlaku. Generic r1 G-B1 tidak hidup kembali: approved PD-r4 menggantinya. Baseline/status/executed claims r1/r2 tidak diwarisi. Gaps r2 kategori (a) input/(b) owner/(c) delegated mechanisms/(d) evidence tetap klasifikasi; PDEP-R2-01 pada handoff r2 tetap release dependency, bukan alasan mengarang solution atau membuka ulang settled creation Product.

## Delta r3 — deterministic technical handoff

| ID | Kewajiban r3 dan accountable consumer |
| --- | --- |
| X1-R3-01 | Architect menyediakan satu exact OpenAPI3.0.3 candidate dalam r3 yang konsisten dengan normative technical decisions dan seluruh11 operasi. Required/optional/null, closed input policy, schema limits/parsing, success/error discrimination dan conditional security harus deterministik bagi dua implementer. No invented operation atau canonical API adoption. |
| X1-R3-02 | Seluruh item T06-J harus ditulis **per operation**; shared rules boleh direferensikan dengan exact applicability. Inapplicable item diberi alasan, bukan blank. Architect memilih ordinary interface details yang masih terbuka dalam authority; Analyst tidak menetapkan problem map, TTL baru, input normalization atau status precedence. Material new trust/Product/standard decision diekskalasi hanya affected edge. |
| X1-R3-03 | Untuk overlap failure, state dan multiple cookie/credential paths, downstream harus dapat menentukan supported request dan outcome tanpa menebak prose: normal versus delete-purpose login, callback success/failure, access versus refresh-bootstrap session, refresh rotation/replay, repeat logout, create/replay/expired receipt, preview/confirm before versus after intent. Architect memiliki formalization/precedence; normative cases mengunci hal yang tidak diekspresikan OAS. |
| X1-R3-04 | Normative conformance cases menautkan operation, exact inputs/preconditions/effective state, expected HTTP/wire/headers/cookies, allowed effect/non-effect, uncertainty dan original X1. Minimal positive, denied/missing/expired/cross-account, malformed/unknown fields/oversize, replay/concurrent/timeout, no false success dan privacy-negative cases. Spec examples, generated types, mocks atau prose checks bukan executed backend/FE conformance. |
| X1-R3-05 | Architect TECH dan Experience EXP harus bertemu pada state/certainty/actor meaning yang sama. FE handoff mencakup typed client/result boundary untuk envelope/problem/transport parse failure; account-switch/cache/key cleanup, safe re-entry, no blind mutation replay dan no auto-login sesudah delete. Consumer-fit assessment harus menyebut input exact revisions serta held gates dan verification yang belum dijalankan. |

## T06-J — checklist coverage seluruh11 operasi

Ini **required coverage checklist**, bukan penilaian bahwa final contract sudah memenuhi. Pada setiap baris, semua kolom checklist J1–J11 di bawah wajib dilengkapi Architect dengan applicable shared anchor atau explicit N/A, serta concrete cases. Method/path berikut hanya inventaris existing effective r1+r2 solution, bukan pilihan endpoint baru oleh X1.

**J1** authority/effect/invariant; **J2** operation/resource/visibility; **J3** security/principal/scope/guest/service; **J4** path/query/header/body/limits/validation; **J5** success/status/envelope/optional/null/no-body; **J6** RFC9457 types/status/extensions/safe detail; **J7** collections/limit/filter/order/cursor consistency; **J8** state/atomicity/concurrency/replay/async truth; **J9** minimum outputs/cache/PII; **J10** compatibility/consumer impact; **J11** spec lint/bundle plus planned contract/FE/security/negative/retry evidence.

| Existing operation | X1 authority joins | Per-operation edge beyond common J1–J11 checklist |
| --- | --- | --- |
| POST /auth/google/attempts | X1-01/02/03, P01/02/09/10 | Normal versus deletion-purpose request/security; before-Google information, no org effect; attempt binding/rate/expiry |
| GET /auth/google/callback | X1-02/07, P01/02/04/06/08 | Protocol redirect/envelope exception, safe failure, bounded one-use proof, same-account deletion purpose, cutoff/fence/no resurrection |
| GET /auth/session | X1-02, P01/04/06 | Exact access/refresh-bootstrap selection and invalid credentials; inspection does not grant refresh-only business mutation |
| POST /auth/refresh | X1-02/07, P04/06 | Single-use rotation, concurrent/replay/revoked/fenced cases, CSRF/Origin and uncertain issuance |
| POST /auth/logout | X1-02/07, P04/06 | Authenticated versus repeated/absent/invalid session; CSRF/Origin behavior and revocation uncertainty must agree |
| POST /organizations | X1-01..08, P01/04/05 | Attributed combined commit; key/account/payload binding, concurrency/admission/expiry; same name never conflict gate |
| GET /organization-creation-requests/{requestId} | X1-04/05/07, P01/04/05 | Own-scope confirmed result; absent/expired observation is not failure, receipt is not current authority |
| GET /organizations/{organizationId} | X1-05/06/08, P01/04 | Current authorized pair; unavailable/unowned/invariant failure, no hidden-resource disclosure |
| GET /organizations | X1-05/06/08, P01/04/05 | J7 applicable: bounded deterministic pagination, authorization before pagination, empty list, cursor validation/consistency |
| POST /account-deletion-previews | X1-03/05/07, P01..04/08/10 | Fresh verified same account; complete dependent set or no actionable preview, zero set, expiry, prepared is not accepted intent |
| POST /account-deletions | X1-04/05/07, P01..08/10 | Explicit exact-set closure, proof/scope race, durable intent/fence/active erase/copies distinction, lost response and no post-delete authority |

J7 untuk operasi selain own-list tetap perlu disposition: preview organizations adalah complete bounded impact set, bukan silently paginated partial confirmation. J10 harus mencatat delta r2→r3/defaults/security/schema/error map dan consumer implications; contract version tidak menciptakan external URL versioning. J11 memisahkan checks dokumentasi yang executed dari runtime cases NOT RUN.

## Non-OAS obligations dan gates

Machine-readable schema/security declarations tidak dapat membuktikan attribution, effective authorization, session/account/fence races, transaction ordering, durable receipt time/cleanup, independent deletion journal durability/coverage, copy/chain expiry atau restore admission safety. Normative technical behavior dan O4 wajib mengikatnya pada X1-02/04/05/07 serta P02–08. OAS juga tidak membuktikan pre-effect comprehension, full-set review, prelogin notice access, accessible interaction, truthful uncertainty atau actual inbox handling; Experience/O4 dan operator evidence memiliki concern itu. Jangan mengubah expiry, weaken safeguards atau menjanjikan all-copies erasure demi membuat schema sederhana.

Synthetic documentation handoff tetap bisa dinilai. **PDEP-R2-01 HELD:** actual basis applicability/necessity/balancing, actual processors/vendors/locations, approved final notice/date, request verification/follow-up readiness, independent journal/backup/restore/physical erasure evidence sebelum real-person/public/privacy-readiness claims. Complex lifecycle/transfer tetap conditional owner gate. Identity/contact telah given; tidak perlu diminta ulang. No email/inbox operation dalam assignment ini.

**Gate consumer:** exact r3 assembly → independent Security dan Solution Review atas originals/final bytes → Human bundle acceptance. Canonical API adoption, protected auth/crypto/authorization code, migrations/implementation dan deployed/runtime conformance adalah izin/evidence terpisah. Broad workspace-write tetap supervised containment deviation, bukan technical r3-only isolation. Source lint/bundle/types tidak memberi cold-resume/native runtime qualification atau FE/BE implementation permission.

**Executed pada X1 ini:** actual source/hash inspection55 files (55 match/0 drift), original Product coverage dan privacy anchors terhadap carry-forward; positive combined outcome serta attribution/partial/unknown/conflict/non-claim boundaries diperiksa pada source. **NOT RUN:** application/unit/coverage/DB/auth/session/security/browser/FE integration/conformance, cleanup/restore/retention/email/vendor/legal checks. Hanya `requirements.md` r3 ditulis. Tidak ada canonical API/code/standards/env/service/install/commit/push. Final hash dilaporkan ke Koordinator; material source drift membuka affected translation dan review kembali.
