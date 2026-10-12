# Proteksi engineering yang sebanding dengan outcome

- **Tanggal:** 2026-10-12
- **Status:** catatan praktik/rekomendasi project yang diminta Human.
- **Scope:** menilai biaya dan manfaat tambahan kontrol, verifikasi, evidence dan orchestration.

Dokumen ini membantu memilih proses yang cukup untuk outcome yang dituju. Tidak mengganti root/backend AGENTS, approved scoped contract, canonical Harscode workflow, explicit Human gates atau required independent verification. Menulis panduan ini bukan keputusan menghapus suatu kontrol atau menerima residual risk.

## Apa yang dimaksud over-protective

Over-protective adalah menambah audit, test, approval, constraint, refactor atau dokumentasi yang manfaat tambahannya tidak sebanding dengan waktu, kompleksitas, risiko perubahan dan perhatian Human yang dibutuhkan.

Yang dinilai adalah **manfaat tambahan**. Kontrol pertama yang menemukan defect mungkin bernilai tinggi; mengulang kontrol yang sama dengan input sama dan tanpa pertanyaan baru bisa bernilai rendah. Biaya tinggi saja tidak membuat kontrol berlebihan: tindakan berisiko besar dapat membutuhkan proteksi mahal.

## Pertanyaan sebelum menambah kontrol

Gunakan pertanyaan singkat berikut dalam reasoning/report yang sudah ada; tidak perlu membuat checklist artifact baru untuk setiap perubahan.

1. Outcome atau keputusan apa yang ingin ditetapkan?
2. Risiko konkret apa yang dikurangi, dan pada boundary mana?
3. Uncertainty apa yang belum terjawab oleh evidence yang tersedia?
4. Hasil baru apa yang dapat mengubah keputusan atau tindakan berikutnya?
5. Apakah scope/capability/prerequisites membuat kontrol ini executable?
6. Adakah cara lebih sempit yang tetap membuktikan outcome dengan fidelity yang diperlukan?
7. Siapa pemilik bukti/keputusan, dan kapan proteksi itu cukup sehingga pekerjaan dapat berlanjut?

Jika tidak dapat menjelaskan pertanyaan material yang dijawab atau perubahan keputusan yang mungkin dihasilkan, kontrol tambahan perlu ditinjau. Ini indikator untuk reasoning, bukan automatic rule melewati mandatory gate.

## Pola proses yang sering tidak sebanding

| Pola | Outcome yang seharusnya dilayani | Pendekatan yang lebih proporsional |
|---|---|---|
| Full suite diulang pada tiap edit/phase | Confidence terhadap affected behavior dan final compatibility | Focused checks dalam Build; targeted reproduction dalam Review; final suite pada owner/point yang membutuhkan final claim. |
| Review/audit berulang dengan input sama | Menutup uncertainty atau memperoleh independence yang diwajibkan | Reuse valid evidence, identifikasi material delta; tambah Review hanya saat question/risk/authority membutuhkannya. |
| Refactor besar untuk menghilangkan scanner warning tanpa defect terbukti | Perilaku aman dan gate yang bermakna | Triage candidate/use-path lebih dulu; disposition sempit yang disahkan atau minimal remediation, bukan mengoptimalkan angka findings dengan mengorbankan safety. |
| Exact-bind semua mutable context | Menjaga assignment dari perubahan makna | Exact-bind contract/scope/review target; baca lifecycle projections untuk semantic consequence, bukan hash freshness ritual. |
| Menyalin semua logs/caches/downloads | Claim dapat direkonstruksi dan disputed facts dapat diinspeksi | Simpan semantic results, identities/provenance dan selected necessary captures; reference canonical copies. Jangan hapus historical evidence atas nama minimization. |
| Meminta approval sama yang sudah berlaku | Menghormati Human authority | Periksa scope/validity approval yang ada. Approval exact model pada Run baru atau protected action baru tetap genuine gate, bukan pertanyaan berulang yang otomatis boleh dihapus. |
| Meminta universal environment readiness | Verification scope runnable | Periksa material prerequisites saja; resolve/narrow known gaps. Human-only interaction adalah execution mode, bukan environment defect. |
| Menambah Run untuk deterministic projection update | Durable coordination state benar | Bounded direct reconciliation bila authority menentukan satu hasil, pre/postconditions bisa diperiksa dan tidak memerlukan semantic decision/independent evidence. |
| Terlalu banyak metadata/process artifacts | Reconstruction dan koordinasi jelas | Isi material consequence pada owning record; jangan membuat identitas/proses paralel yang menduplikasi truth. |
| Tests hanya menyalin implementation atau compile/no-tests PASS | Membuktikan invariant yang dapat gagal | Gunakan oracle observable meaningful; jangan hitung command success sebagai coverage yang tidak ada. |

## Pilihan tindakan berdasarkan situasi

| Situasi | Arah yang disarankan |
|---|---|
| Evidence masih valid dan menjawab exact claim | Referensikan; tidak perlu duplicate execution hanya karena context/Run berubah. Tetap penuhi fresh checks yang diwajibkan owning contract. |
| Delta lokal dengan oracle yang jelas | Focused verification dahulu; broaden jika actual failure/risk/delta menuntutnya. |
| Attribution/applicability diperselisihkan | Targeted diagnostic dan selected raw capture yang cukup; jangan mendaur ulang broad matrix. |
| Material defect pada auth/authorization/locking/security ditemukan | Stop improvisation, route owner dan apply authorized remediation; biaya proses sebanding dengan consequence. |
| Prerequisite outcome diketahui belum tersedia | Resolve atau narrow claimed outcome; defer eksplisit jika authority mengizinkan. Tidak memberi label PASS dari simulasi. |
| Keputusan owner sudah evidence-ready | Fasilitasi exact decision; jangan menggantinya dengan audit/Run baru untuk menghindari pertanyaan. |
| Causal blockage berulang setelah reconciliation | Gunakan current Loop Health trigger/stop guidance; jangan menambah iteration tanpa changed premise. |

## Batas perampingan

Perampingan tidak mengizinkan:

- melewati explicit Human approval atau protected money/auth/key/privileged-authority boundary;
- menghapus expected failure handling, secret hygiene, SQL parameterization atau money precision;
- mengganti real/concurrent/independent proof dengan fixture/unit/Human-attested evidence jika owning contract memerlukan bukti lain;
- menyuppress seluruh scanner rule/file atau menerima risk tanpa pemilik;
- menghapus findings/limitations atau merewrite historical/terminal evidence;
- melabel control sebagai “over-protective” hanya karena control menemukan kegagalan;
- menyamakan progress/completion scope dengan acceptance seluruh commitment.

Konteks source/toolchain/config/profile berbeda bisa membuat evidence sebelumnya tidak lagi berlaku. Evidence reuse memerlukan relevansi terhadap claim saat ini; checksum/source cache atau report lama bukan jaminan freshness semua claim.

## Cara menilai outcome tanpa bureaucracy baru

Evaluasi material delta dalam report/handoff yang sudah dimiliki pekerjaan:

- uncertainty atau failure apa yang ditutup;
- keputusan owner atau next route apa yang kini bisa ditentukan;
- behavioral/security evidence apa yang bertambah;
- limitation apa yang masih ada;
- langkah tambahan apa yang tidak diperlukan lagi.

Bila ingin menilai biaya, bedakan nilai terukur dari estimate: waktu command, recovery attempts, jumlah duplicated executions atau blockers yang benar-benar tercatat dapat digunakan. Jangan mengarang angka penghematan atau skor risk. Jumlah file/Run/log bukan proxy manfaat tanpa melihat substance.

## Contoh

**Proporsional:** dua validation error literals diperbaiki, focused lint dijalankan. Tidak perlu membuat tes yang hanya membandingkan dua strings agar jumlah tests bertambah.

**Proporsional:** cookie unit oracle memeriksa emitted headers; real DB lock contention tetap diuji terpisah karena oracle dan risiko berbeda.

**Berlebihan bila diulang tanpa alasan baru:** scanner attribution sudah direproduksi pada exact fingerprint/source; audit lain hanya membacakan assessment yang sama, sementara keputusan owner tidak pernah diminta.

**Tetap justified:** diagnostic independen menemukan bahwa scanner mematch komentar, bukan credential assignment. Biaya diagnostic menghasilkan material changed understanding; bukan sekadar clean/no-finding ritual.

**Perlu owner:** narrow scanner disposition sudah memiliki rationale. Reviewer dapat merekomendasikan; keputusan policy dan implementation scope tetap pada pemilik. Tidak ada automatic allowance dari panduan ini.

## Penerapan dan authority

Orchestrator dapat menggunakan reasoning ini untuk memilih scope/route yang sudah berada dalam authority berlaku. Bila recommended simplification mengubah requirement, mandatory gate, risk acceptance, scope atau protected authorization, keputusan harus datang dari owning authority dahulu. General recommendation ini tidak otomatis menjadi Harscode-wide policy dan tidak mengubah AGENTS atau protected workflow guidance.

Basis: current Harscode Build/Review proportional verification, run-contract binding/evidence-sufficiency/deterministic reconciliation, dan orchestration readiness/decision facilitation. General framing manfaat tambahan adalah sintesis project, bukan kutipan mandatory workflow rule. Harscode guidance tetap current-effective melalui router repository.

Lihat [evaluasi konkret C1/T2](c1-t2-over-protective-process-assessment.md) untuk fakta/risiko/rekomendasi spesifik pada commitment aktif.
