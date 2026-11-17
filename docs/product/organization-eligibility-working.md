# Kencleng — Organization Eligibility: Working Product Definition

> **Standing:** WORKING / NOT APPROVED AS BINDING PRODUCT DEFINITION  
> **Checkpoint:** 16 November 2026  
> **Owning concern:** Kencleng Product  
> **Human decision-maker:** Anhar  
> **Scope:** Organization Eligibility discovery; **not** an approved release slice, Product Contract, UX flow, Engineering specification, or permission to implement.  
> **Provenance:** Consolidation of Human-agreed working directions from the current Product discussion. Approval of the checkpoint *approach* does not approve these exact document bytes or promote their content into canonical Product commitments.

## 1. Product outcome and boundaries

**Working outcome:** Organisasi nirlaba Indonesia yang memenuhi persyaratan dapat memperoleh dan mempertahankan kelayakan sebagai penyelenggara Campaign di Kencleng berdasarkan bukti legitimasi, kewenangan perwakilan yang sah, dan accountability berkelanjutan.

Nilai yang dituju:
- **Organisasi:** mengetahui dasar kelayakan, hak, pembatasan, dan kewajiban untuk dapat menyelenggarakan penggalangan dana melalui Kencleng.
- **Kencleng:** memiliki dasar yang dapat dipertanggungjawabkan untuk mengakui, meninjau ulang, membatasi, atau mencabut kelayakan.
- **Pengunjung/Donatur:** memperoleh informasi dengan makna yang jujur mengenai pihak yang bertanggung jawab atas Campaign.

**Initial eligible-entity scope (working):** Yayasan dan Perkumpulan berbadan hukum Indonesia dengan entitas hukum sendiri. Komunitas informal, unit di bawah entitas hukum lain sebagai organisasi tersendiri, dan bentuk organisasi lain belum termasuk cakupan awal yang sedang dieksplorasi. Ini bukan penetapan kewajiban legal atau scope rilis.

**Pemisahan domain:** Organization Eligibility, kurasi/penayangan Campaign, kelayakan Campaign menerima donasi, dan verifikasi penggunaan dana adalah keputusan berbeda. Organisasi yang eligible **tidak otomatis** berhak mempublikasikan setiap Campaign, menerima donasi untuk setiap Campaign, atau mengklaim outcome terverifikasi. Proses simulasi bukan bukti pemeriksaan atau hasil nyata.

**Bukan scope dokumen ini:** keputusan MVP/rilis, UX Journey atau screen/flow, mekanisme pemeriksaan teknis, API/database/permissions implementation, settlement/pencairan/refund, standar legal yang telah divalidasi, dan bukti runtime atau dampak nyata.

## 2. Authority, sources, and evidence standing

**Binding foundation:** [Kencleng Product Intent](https://github.com/anhsbolic/kencleng/blob/7778d66d91ba81fff08a6ae1ee0aac9ef71905f8/docs/product/product-intent.md), approved 22 Oktober 2026; blob `cc98e46dd1d8d9e32b8718eb3b95ad9a6ffc2312`, dengan scope whole-product purpose/principles. [Product index](https://github.com/anhsbolic/kencleng/blob/7778d66d91ba81fff08a6ae1ee0aac9ef71905f8/docs/product/README.md) mengatur routing Product dan menyatakan Anhar final decision-maker untuk consequential commitments.

**Research guidance (not Kencleng authority):** [Harscode Product Definition](https://github.com/anhsbolic/harscode-workspace/blob/302bedae43563a66d2ef914191c18139440c0c30/product/product-definition.md), [Authority & Approval](https://github.com/anhsbolic/harscode-workspace/blob/302bedae43563a66d2ef914191c18139440c0c30/product/authority-and-approval.md), dan [Consumer Handoff](https://github.com/anhsbolic/harscode-workspace/blob/302bedae43563a66d2ef914191c18139440c0c30/product/consumer-handoff.md); research candidate dengan exact-content review pending.

**Separately owned constraints:** [Kencleng Design](../design/README.md) dan [Engineering Standards](../engineering/standards/README.md). Product memutuskan tujuan, hak dan consequential semantics; UX menentukan experience journey; Design menentukan penyajian; Engineering menentukan mekanisme dan verifikasi teknis. Ketentuan hukum, privasi, dan keamanan yang relevan tetap perlu diperiksa pada owning authority.

**Meaning of standing in this document:**
- **AGREED WORKING:** arah substantif telah disepakati dalam diskusi; belum menjadi approved exact-revision Product source.
- **CANDIDATE:** analisis atau rekomendasi yang belum disetujui sebagai detail substantif.
- **OPEN:** keputusan atau evidence material masih diperlukan.
- **FOUNDATION APPROVED:** hanya makna dalam scope Product Intent dan sumber project lain yang benar-benar approved.

Belum ada riset pengguna independen, legal opinion khusus, pengujian consumer nyata, atau verifikasi runtime yang dilaporkan sebagai telah dilakukan untuk behavior dalam checkpoint ini.

## 3. Consolidated agreed working meaning

### 3.1 Organisasi, pendaftaran, dan legitimasi

- Kelayakan organisasi memerlukan kombinasi bukti legalitas/legitimasi eksternal dan penilaian sesuai persyaratan Kencleng.
- Satu entitas hukum memiliki satu identitas organisasi yang diakui di Kencleng. Beberapa pemohon boleh memulai permohonan terhadap entitas yang sama; pendaftar pertama tidak otomatis memperoleh kontrol.
- Pemohon boleh memulai permohonan sebelum seluruh bukti tersedia. Pengajuan peninjauan formal memerlukan informasi dan bukti minimum yang cukup untuk dinilai; **minimum final belum disepakati**.
- Menyerahkan dokumen, menyelesaikan kelengkapan, dan berhasil memvalidasi bukti adalah kejadian berbeda. Bukti yang tidak dapat dipastikan bukan otomatis palsu atau pelanggaran.
- Untuk menyatakan organisasi eligible, Kencleng harus memiliki dasar bukti yang dinilai cukup, dapat ditelusuri, dan keputusan yang berwenang; detail ambang bukti belum diselesaikan.

### 3.2 Representative Authority

- Pemohon yang mengajukan pendaftaran tidak otomatis menjadi `Owner`.
- `Owner` pertama boleh merupakan perwakilan dengan mandat resmi, tanpa harus menjadi pimpinan hukum organisasi. Bukti mandat untuk mengajukan permohonan **tidak otomatis** cukup sebagai mandat mengendalikan organisasi.
- Penambahan/pergantian `Owner` membutuhkan dasar kewenangan organisasi yang sah; tindakan `Owner` aktif sendiri tidak otomatis cukup tanpa mandat yang mencakupnya.
- `Owner` dan `Staff` memiliki kewenangan berbeda; batas permission, jumlah Owner, pengelolaan Staff, dan penyelesaian sengketa mandat **belum final**.
- Perubahan atau pencabutan mandat seseorang tidak mengubah identitas legal organisasi dan tidak menghapus riwayat atau kewajiban organisasi.

### 3.3 Accountability and ongoing eligibility

- Organisasi harus bersedia menerima kewajiban accountability Kencleng sebagai bagian dari eligibility. Rincian minimum obligations belum diputuskan; daftar seperti truthfulness, valid mandate, disclosure, Campaign reporting, dan cooperation masih candidate untuk diuji.
- Organisasi wajib proaktif mengungkap perubahan material yang memengaruhi dasar kelayakan atau kewenangan, termasuk ketika belum memiliki Campaign aktif.
- Secara independen dari laporan organisasi, Kencleng wajib proaktif melakukan pengawasan dan penilaian ulang yang proporsional terhadap risiko. Ini tidak menyatakan pemeriksaan seluruh organisasi secara real-time atau verifikasi terus-menerus.
- Peninjauan ulang tidak otomatis berarti eligibility hilang. Tanggal/bukti terakhir diperiksa dan keputusan kelayakan yang masih berlaku tidak boleh diklaim sebagai hal yang sama.
- Jika organisasi berulang kali tidak kooperatif terhadap pemeriksaan material setelah pemberitahuan dan kesempatan respons yang wajar, Kencleng dapat mempertimbangkan penangguhan secara proporsional tanpa menyebutnya terbukti melakukan penipuan.

### 3.4 Suspension, revocation, and recovery

- Eligibility tidak permanen: dapat ditinjau ulang, ditangguhkan, atau dicabut melalui keputusan yang sah.
- Penangguhan sementara sebelum pelanggaran terbukti final memerlukan **dasar kredibel, risiko material, dan intervensi yang proporsional**; tindakan harus berwenang, beralasan, dapat ditinjau ulang, dan tidak dipresentasikan sebagai bukti pelanggaran final.
- Ketika Organization Eligibility **ditangguhkan**, seluruh Campaign di bawahnya tidak boleh **menerima donasi baru**. Riwayat Campaign/donasi dan kewajiban accountability tidak otomatis dihapus.
- Penolakan permohonan berbeda dari pencabutan kelayakan yang sebelumnya berlaku. Penolakan/pencabutan tidak otomatis permanen: penilaian kembali boleh diminta ketika ada bukti atau perubahan material yang baru serta tidak ada pembatasan sah yang menghalangi. Hak tidak otomatis pulih hanya karena permohonan diajukan.
- Perlakuan terhadap donasi *in-flight*, dana yang telah diterima, pencairan/refund, dan hak operasional yang tetap diperlukan selama suspension **belum ditetapkan**.

## 4. Meaningful dimensions — not one mandatory status

| Dimensi | Pertanyaan yang dijawab | Batas penting |
| --- | --- | --- |
| Application/assessment | Apakah informasi dan bukti cukup untuk diajukan/dinilai? | `Incomplete`, `pending`, dan `rejected` adalah *candidate meanings*, bukan state machine final. |
| Organization Eligibility | Apakah keputusan kelayakan saat ini masih berlaku atau dibatasi? | `Eligible`, `suspended`, `revoked` adalah *candidate meanings*, bukan satu status teknis yang diwajibkan. |
| Representative Authority | Siapa mempunyai mandat untuk tindakan apa, dan apakah mandat masih berlaku? | Hak orang tidak identik dengan eligibility organisasi. |
| Evidence/oversight | Apa yang sudah diperiksa, sumber/waktu/hasil, dan apa yang masih belum pasti? | Reassessment dapat berlangsung sementara keputusan eligibility masih berlaku. |

Jangan menyederhanakan seluruh dimensi ini menjadi satu label `verified` yang menutupi perbedaan fakta, review, hak, dan ketidakpastian.

## 5. Candidate minimum information and evidence

**Untuk memulai permohonan (candidate):** nama organisasi yang diklaim; identitas/kontak pemohon untuk melanjutkan permohonan; hubungan pemohon dengan organisasi yang dinyatakan sendiri (*self-declared*). Tidak satu pun merupakan bukti mandat atau kelayakan.

**Untuk mengajukan peninjauan formal (candidate):** identitas legal organisasi (nama/bentuk/identitas legal/representasi alamat relevan), bukti legalitas eksternal, dasar karakter nirlaba, identitas dan hubungan perwakilan, bukti kewenangan sesuai tindakan yang diminta, kontak organisasi, serta penerimaan kewajiban accountability oleh pihak yang berwenang.

**Untuk keputusan eligible (working principle; detail OPEN):** bukti yang memadai mengenai legalitas, karakter nirlaba, kewenangan perwakilan, dan kesesuaian kebijakan Kencleng; provenance dan hasil pemeriksaan dapat ditelusuri; peninjau berwenang dan tidak memiliki conflict of interest.

Jenis dokumen yang diwajibkan untuk Yayasan/Perkumpulan, pemeriksaan sumber resmi, kebutuhan data personal, retention/disclosure, kewajiban perizinan penggalangan dana, dan ambang kecukupan bukti **belum diputuskan atau divalidasi secara legal**. Jangan mengubah contoh KTP, NPWP, rekening, surat kuasa, atau portal pemeriksaan menjadi kewajiban universal hanya dari checkpoint ini.

## 6. Representative scenario review

| Scenario | Yang sudah cukup jelas untuk dieksplorasi | Open consequential edges |
| --- | --- | --- |
| **S1 — Pendaftaran/peninjauan awal** | Boleh mulai belum lengkap; submission dan verifikasi berbeda; belum memperoleh hak hanya dengan mendaftar. | Evidence minimum, decision threshold, rejection/insufficient/contradictory results, authorized decision owner. |
| **S2 — Multiple applicants/Owner change** | Satu identitas organisasi; first applicant bukan otomatis Owner; Owner perlu mandat yang sah. | Konflik mandat, perlindungan data antar-pemohon, continuity saat tidak ada Owner yang sah. |
| **S3 — Ongoing oversight** | Proaktif dua pihak; reassessment tidak otomatis suspension; non-cooperation material dapat memicu penangguhan proporsional. | Parameter oversight, evidence validity/freshness, kesempatan respons yang wajar, siapa memutuskan intervensi. |
| **S4 — Suspension with active Campaign** | Donasi baru dihentikan; Campaign/riwayat/accountability tidak otomatis hilang. | In-flight donations, money treatment, public visibility, permitted accountability actions. |
| **S5 — Rejection/revocation/recovery** | Penilaian kembali dimungkinkan berdasarkan perubahan/bukti material; tidak memulihkan hak secara otomatis. | Kondisi pemulihan, dampak pada Campaign existing, perbedaan prosedural antar-keputusan. |

Review ini adalah analisis koherensi *working decisions*, bukan acceptance independen dari UX/Engineering atau verifikasi software.

## 7. Open decisions and needed evidence

| Open question | Mengapa material | Owner/needed input |
| --- | --- | --- |
| **Eligibility Decision Basis:** bukti apa yang benar-benar sufficient, insufficient, atau contradictory untuk Yayasan dan Perkumpulan? | Mencegah false verification maupun keputusan arbitrer. | Product/Anhar; domain/legal evidence dan review ops. |
| Siapa boleh mengambil keputusan approval, rejection, suspension, revocation; bagaimana conflict of interest/appeal ditangani? | Legitimasi keputusan dan akuntabilitas. | Product/Anhar; operational/legal constraints. |
| Apa minimum kewajiban organisasi yang mengikat, dan siapa sah menyetujuinya? | Mengubah agreement menjadi consequential commitment tanpa sekadar checkbox. | Product/Anhar; legal/privacy evidence. |
| Apa mandat cukup untuk Owner/Staff, perubahan mandat, dan sengketa dua pemohon? | Mencegah pengambilalihan organisasi secara tidak sah. | Product/Anhar; domain authority/security evidence. |
| Apa bukti/interval yang membuat reassessment wajib; kapan evidence outdated menjadi material? | Oversight harus feasible dan klaimnya jujur. | Product/Anhar; operating capability/risk evidence. |
| Apa akibat suspension untuk Campaign yang telah terbit, donasi in-flight, uang, pelaporan, dan pihak terdampak? | Dampak pada donor rights dan risiko uang tinggi. | Product/Anhar dengan legal, finance, security, Engineering constraints saat diperlukan. |
| Apa scope delivery dan measurable acceptance untuk unit yang dipilih? | Menentukan apakah consumer dapat membuat keputusan binding. | Product/Anhar setelah bounded outcome dipilih untuk delivery. |

**Klasifikasi:** Semua di atas adalah OPEN PRODUCT DECISION, kecuali mekanisme teknis pelaksanaan yang merupakan OPEN ENGINEERING CONSTRAINT dan layout/interaction yang merupakan OPEN EXPERIENCE/DESIGN DECISION. Fakta runtime belum diverifikasi.

## 8. Consumer-specific sufficiency

| Receiving responsibility | Current assessment |
| --- | --- |
| Product discovery & non-binding UX exploration | **CONDITIONAL:** boleh mengeksplorasi hipotesis dan skenario, dengan standing/agreed vs proposed/open eksplisit. |
| Binding UX/UI decisions yang menyatakan kelayakan, status, disclosure, atau konsekuensi | **BLOCKED ON AFFECTED EDGES:** memerlukan accepted Product meaning yang material; Design tetap memiliki authority sendiri. |
| Engineering solution/development | **NOT SUFFICIENT:** belum ada exact-revision scoped approved Product behavior, rights, acceptance, dan keputusan risiko lintas domain. |
| Independent consumer acceptance, legal verification, and runtime verification | **NOT CHECKED / NOT VERIFIED.** |

## 9. Recommended next Product inquiry

**Eligibility Decision Basis — Yayasan** sebagai skenario uji pertama:

> Yayasan dengan identitas legal, bukti mandat, dan pernyataan penerimaan kewajiban telah mengajukan permohonan. Apa yang benar-benar dapat disimpulkan Kencleng ketika (a) evidence cukup, (b) evidence belum cukup, atau (c) evidence saling bertentangan?

Tentukan only consequential Product meaning yang diperlukan: sumber/bukti, perbedaan klaim dan fakta diperiksa, authorized decision, observable result, negative/unknown cases. Lanjutkan ke Perkumpulan dengan mengevaluasi perbedaan material, bukan mengasumsikan semua dokumennya sama.

**No automatic promotion:** Checkpoint ini menjaga durable state discovery. Setiap Product commitment binding memerlukan peninjauan exact content, persetujuan Anhar pada scope yang jelas, dan source/approval revision yang dapat direkonstruksi. Tidak ada izin implementasi atau repo mutation yang diturunkan dari dokumen ini.
