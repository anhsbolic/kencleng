# Kencleng — Dokumentasi & Konvensi Bahasa (T09)

> **Status:** Project-wide Engineering Standard T09, Human-settled dan persisted pada research branch `pilot/station5-c1-space-only`. **Bukan** promosi ke Kencleng `main`, perubahan Product meaning, atau persetujuan penyuntingan massal.
> **Scope:** Default penulisan human-facing documentation dan handoff baru atau yang sedang diubah secara material, authority routing, terminology, penyuntingan, dan pemeliharaan dokumen lintas-concern.
> **Owner:** Kencleng Engineering Standards untuk panduan lintas-concern. Standar penulisan Product dan pilihan istilah Product tetap dimiliki Product Kencleng, bukan dokumen ini.

## Dasar dan precedence

- Kencleng memilih **Bahasa Indonesia sebagai default narasi untuk human-facing documentation baru** dan bagian dokumen aktif yang sedang ditulis ulang. Gunakan technical English jika lebih natural, merupakan istilah resmi, atau terjemahan mengurangi ketepatan. Ini bukan kewajiban menerjemahkan existing documents.
- More-specific **approved project-owned writing conventions** berlaku dalam scope-nya sendiri: [standar penulisan Product](../../product/README.md#standar-penulisan-dokumen-product-kencleng) dan [pilihan istilah Product](../../product/istilah.md) tetap authoritative bagi dokumen Product aktif. [Product Design authority](../../design/README.md) tetap mengatur concern desain. T09 tidak menggantikan Product terms atau mengesahkan perubahan Product/business semantics.
- Reusable inspiration: [Harscode — Project Language and Writing Profile (reference example)](https://github.com/anhsbolic/harscode-workspace/blob/04fd378cc1a4cab86c39368303f595b8b5a46ed9/engineering/references/project-language-writing-example.md). Sumber tersebut **optional, non-binding, research-only**, **bukan operational Harscode `main` policy** atau Kencleng approval otomatis. Prinsip yang dipilih disetujui Human dan **di-own oleh Kencleng melalui T09**; Kencleng tidak bergantung pada perubahan future branch Harscode.
- Kebijakan format, bahasa, nama field, identifier, dan output untuk interface machine-readable mengikuti owning technical contract (mis. [T06 API](api-contract.md)), bukan default narasi Bahasa Indonesia.

## T09-A — Bahasa human-facing yang jelas

- Pilih Bahasa Indonesia yang umum, langsung, dan tidak bertele-tele untuk penjelasan di documents/handoffs yang baru. Pertahankan technical English ketika memperjelas maksud: contoh `Modular Monolith`, `requestId`, `OpenAPI`, `GIVEN/WHEN/THEN`, `transaction boundary`, `HTTP 422`.
- Nama resmi, project roles, branch, paths, CLI commands, code, JSON/SQL fields, API endpoints, status/error identifiers, serta exact artifact IDs **tidak diterjemahkan**. Jangan mengubah schema atau public interface hanya demi keseragaman bahasa.
- Penulisan dalam English pada dokumen existing bukan pelanggaran yang perlu dibersihkan massal. More-specific audience/consumer requirements dapat membenarkan English sebagai bahasa utama untuk artifact tertentu; jelaskan bila choice itu materially memengaruhi consumption.
- Tidak ada mandatory bilingual delivery, parallel translation, atau project-wide language conversion.

## T09-B — Terminology, clarity, dan checkable wording

- Untuk domain Product, **reuse existing Product terminology authority**, termasuk istilah khusus `Campaign`, `Donatur`, `Journey`, `Owner`, `Staff`, dan istilah yang disepakati dalam [glossary Product](../../product/istilah.md). Jangan membuat glossary kedua untuk T09 atau mengubah definisi domain dari pilihan editorial.
- Tuliskan kalimat singkat, actionable, dan langsung pada decision/rule/requirement. Hindari repeated narrative, filler, atau formulasi yang memerlukan pembaca menebak practical consequence.
- Jika istilah seperti *valid*, *cukup*, *sah*, *relevan*, *approved*, atau *berhasil* menentukan acceptance, authority, atau readiness, sertakan kondisi/criteria yang **dapat diperiksa** dan konsekuensi bila tidak terpenuhi; jangan menciptakan arbitrary threshold.
- Bila ambiguitas istilah akan mengubah business rules, surface material gap kepada owner Product/related concern; jangan menetapkan keputusan business melalui dokumentasi Engineering.

## T09-C — Authority, status, dan semantic preservation

- Setiap artifact harus jelas fungsi, concern owner, dan standing-nya pada tingkat yang diperlukan: **proposal/hypothesis**, **Human-approved decision**, **implemented**, **executed/verified evidence**, atau **historical**. Jangan menaikkan suggestion, experiment, implementation precedent, atau historical artifact menjadi canonical rule hanya karena wording berubah.
- Pertahankan semantics saat menulis ulang atau menerjemahkan: obligation modality (`must/shall/should/may` ↔ `wajib/sebaiknya/boleh`), exclusions, exceptions, scope, permissions, approval gates, failure cases, input/output meaning, dan acceptance criteria. **Translation/editorial cleanup is not semantic authorization**; perubahan material harus disetujui owning Human/authority sesuai aturan yang berlaku.
- Bedakan `NOT RUN`, `NOT VERIFIED`, `FAILED`, `PASSED`, approved design, dan real execution. Jangan menerjemahkan status supaya terdengar lebih positif atau mengklaim test/runtime behavior yang tidak dibuktikan.
- Informasi penting untuk fresh-session reconstruction harus ditemukan melalui current owning artifacts dan exact links/revisions, tidak bergantung pada chat recall. Self-contained berarti cukup untuk concern yang dimiliki, **bukan** menduplikasi isi seluruh referenced documents.
- Existing `main` dan research branch memiliki standing berbeda; T09 approval di branch ini bukan promosi operational main.

## T09-D — Documentation ownership dan readable structure

- Gunakan **satu owning source per semantic concern**; dokumen lain cukup memberikan scoped link/navigation, bukan copy normatif kedua. Product menetapkan business meaning; Engineering menentukan approved technical constraints; feature-level solution menetapkan scoped implementation decisions; executable state/tests menyatakan actual evidence.
- Gunakan headings yang mudah ditemukan; sertakan requirements, constraints, stop/failure paths, explicit unresolved decisions dan references bila benar-benar membantu reader/Agent menjalankan task. Jangan membuat dokumen panjang atau template wajib yang tidak menambah semantic correctness.
- Contoh, referensi, atau history harus dilabeli sesuai statusnya; bukan implied binding implementation. Jangan diam-diam mengubah approved product criteria melalui technical examples.
- Handoffs untuk Agent harus cukup akurat untuk menemukan current effective authority dengan fresh context. Tidak perlu membuat separate language-profile file hanya karena Harscode mempunyai reference example.

## T09-E — Proportional maintenance & translation control

- Update owning docs, index/Agent routing, dan consuming references **ketika meaning, interface, location, ownership, atau effective status benar-benar berubah**. Scope edit sesempit mungkin; verifikasi material diff dan links yang terdampak.
- Tidak perlu bulk-translate existing Engineering/Product Design documents, merevisi historical Rancang/handoffs agar seolah sesuai kebijakan baru, atau mengubah immutable approval/provenance. Historical records dan snapshot tetap evidence sesuai original revision.
- Translation bilingual dibuat **hanya jika ada real consumer need**; project owner menentukan versi authoritative. Versi lain harus jelas berstatus derivative translation, tidak menjadi parallel editable authority. Setelah material original changes, reconcile translations sebelum mengklaim equivalence.
- Tidak ada mandatory markdown linter, documentation CI, per-function documentation, blanket README creation, glossary proliferation, atau automated formatting migration hanya untuk memenuhi T09. Tambahkan tooling ketika ada recurring demonstrated defect yang layak dibayar dengan complexity.
- Editorial update yang tidak mengubah meaning boleh ditangani secara ringan; material Product/Engineering semantics tetap membutuhkan scoped approval dan review sesuai owning standard.

## T09-F — Code comment style (Go dan frontend)

- **Default comments singkat, biasanya satu atau dua baris.** Jelaskan *apa yang dilakukan* dan *hasil/tujuan yang dapat diamati* bila tidak jelas dari kode. Jangan menarasikan setiap statement, mengulang function name tanpa informasi baru, atau menambahkan comment ceremonial. Jika kode sudah self-explanatory, comment tidak diperlukan.
- Utamakan **prosedur dan expected result**, bukan uraian keputusan historis atau rationale panjang. Simpan penjelasan desain/trade-off yang memang perlu dipertahankan pada owning technical/design artifact yang sudah ada; **jangan membuat Techplan baru hanya untuk menampung rationale sebuah komentar**.
- **Pengecualian dengan alasan nyata:** penjelasan lebih panjang boleh ketika diperlukan untuk menjaga correctness, public/external contracts, subtle security/privacy requirements, atau non-obvious business invariants yang rawan dilanggar saat code berubah. Jika invariant kritis tidak jelas tanpa membaca dokumentasi lain, pertahankan short essential safeguard di dekat kode; boleh tautkan owning doc. **Batas 1–2 baris adalah default, bukan hard cap yang menghapus informasi keselamatan.**
- Pertahankan language dan technical vocabulary yang presisi sesuai konvensi source file; T09 Bahasa Indonesia untuk human-facing documents **tidak memaksa penerjemahan identifiers maupun seluruh source comments**. Pada Go, ikuti GoDoc conventions untuk exported declarations saat dokumentasi API diperlukan; jangan mengorbankan required contract detail demi panjang komentar.
- **Pengecualian T07:** tiga comment markers `// GIVEN`, `// WHEN`, dan `// THEN` dalam subtests tetap wajib **sebagai tiga baris terpisah** menurut [T07-I](testing-verification.md#t07-i--unit-coverage-80-and-behavior-quality); aturan 1–2 baris ini tidak berlaku untuk markers tersebut.
- Jangan membersihkan existing comments secara massal. Ketika logic berubah, perbarui/hapus komentar terdekat yang stale agar tidak mendeskripsikan behavior yang salah.

Contoh singkat (ilustratif, bukan implementasi Kencleng):

```go
// indexRowsByField groups rows by field for duplicate checks.

// Preserve the caller-provided value.
```

## Consumption rule

Jika Agent hendak menulis dokumentasi human-facing baru, cari *applicable project writing convention* secara proporsional melalui [Engineering Standards index](README.md), lalu ikuti more-specific rules untuk concern yang sedang diubah. **Jangan membaca semua dokumen lintas-folder hanya karena T09 ada.**

**T09 hanya menetapkan Kencleng writing convention pada research branch.** Tidak menambahkan Product features, implementasi runtime, API endpoints, atau kewajiban menerjemahkan existing artifacts.
