# Kencleng — Product

> **Authority:** Current Product artifacts for branch `pilot/station5-c1-space-only` only; Kencleng operational `main` is unchanged. Product owner: Anhar.
> **Tracking:** Use the phase/output types of [Harscode Product](https://github.com/anhsbolic/harscode-workspace/blob/refactor/harscode-foundation/product/README.md); project identifiers and exact approvals are recorded in [Identifier Migration](identifier-migration-2026-10-23.md).
> **Boundary:** Product-semantic approval is not an Engineering work order or Rancang invocation.

## Product phases and current outputs

| Phase | Phase ID | Output type | Current Kencleng artifact | Standing |
| --- | --- | --- | --- | --- |
| Product Foundation | `PROD-FND` | `PROD-INTENT` | [Product Intent](product-intent.md): purpose, actor value, responsibilities, truth boundaries | **APPROVED** (22 Oktober 2026) |
| Product Journey Mapping | `PROD-JRN` | `PROD-JOURNEYS` | [Interconnected Product Journeys](interconnected-journeys/README.md) and [Level 2](interconnected-journeys/level-2-network.md) | **APPROVED map**, with node IDs migrated structurally |
| Outcome Selection | `PROD-SEL` | `PROD-SELECTION` | [Selected Outcome ORG-01](outcome-selection.md) | **SETTLED**, specific to ORG-01 |
| Outcome Definition | `PROD-DEF` | `PROD-CONTRACT` | [Product Contract — ORG-01](outcomes/org-01/product-contract.md) | **APPROVED**, with new tracking IDs |
| Product Review & Approval | `PROD-APR` | `PROD-APPROVAL` | [Approval Record — ORG-01](outcomes/org-01/approval.md) | **CLOSED / pinned revision** |
| Product Handoff | `PROD-HND` | `PROD-HANDOFF` | Refer to approved Contract + Approval Record when authorized | **NOT INVOKED**; no separate packet created |

The output types identify **what kind of information exists**, not six required files, separate runtime stages or a universal delivery sequence. The Product Contract for `ORG-01` is authoritative within its scope. Other Journey nodes are not approved Product Contracts by implication.

## Product value and Design boundary

[Product Intent](product-intent.md) owns Kencleng's approved value direction, actor responsibilities and Product truth principles. Future Value Definition elaboration belongs to Product Foundation when a material gap is demonstrated and approved by Anhar; do not import unapproved business value narratives from historical experiments. Product Journey Mapping connects these values to actor experience. [Product Design Principles](../ui-ux/product-design-principles.md) owns the way approved Product meaning is expressed and understood in UI/UX; it does not independently create new business promises.

## Product identity and approval

Current selected Product outcome is **`ORG-01`**, corresponding to Journey node `J-ORG-01`. Track artifacts as `PROD-SELECTION-ORG-01`, `PROD-CONTRACT-ORG-01`, and `PROD-APPROVAL-ORG-01`. The **26 Journey nodes and 26 approved Contract identifiers** have explicit old→new mapping in the [migration record](identifier-migration-2026-10-23.md).

The exact original Human-approved Product Contract blob is `448a8082bcb9ac1c7e41a686f2e3fdd2e3988ac4`. The approved meaning was preserved during the identifier migration; the current canonical content/revision is pinned by its [approval record](outcomes/org-01/approval.md). Historic codes belong to historical snapshots only.

## Boundaries for consumers

- Product Contract must be self-contained in behavior, authority, success, negative paths and acceptance; IDs are traceability metadata, not instructions to reconstruct upstream research.
- Rancang formed v3 still has source-inspection clauses. The desired consumer boundary is **not yet an approved Rancang definition amendment**.
- `I2/I3/I4`, technical constraints, current implementation/security evidence and Human-direct assignment remain separate. No Engineering or Space invocation is authorized here.
- No change to the Kencleng Product Intent's approved semantics, operational `main`, Product Design authority, or the applicable Harscode `main` baseline.

## Standar penulisan dokumen Product Kencleng v0.1

> **Disetujui Anhar: 22 Oktober 2026.** Berlaku untuk penyusunan dan penyuntingan **dokumen Product Kencleng yang aktif**, dari tujuan produk hingga dokumen kebutuhan untuk tim berikutnya. **Tidak mengubah Harscode**, aturan Product, atau keputusan yang telah disetujui.

1. **Utamakan Bahasa Indonesia yang umum.** Ikuti [pilihan istilah Product Kencleng](istilah.md), termasuk istilah Inggris yang sengaja disepakati (misalnya `Campaign`, `Donatur`, `Journey`, `development`, dan `conflict of interest`). Untuk istilah lain, gunakan bahasa Inggris hanya bila nama resminya harus dipertahankan atau padanan Indonesianya kurang jelas. Jelaskan istilah teknis ketika pertama dipakai.
2. **Tulis langsung dan ringkas.** Satu gagasan utama per kalimat. Hindari pengulangan, pengantar panjang, dan istilah yang tidak diperlukan.
3. **Susun agar mudah dicari.** Pisahkan tujuan, pengguna, aturan, syarat keberhasilan, kondisi gagal, batas pekerjaan, dan hal yang belum diputuskan sesuai kebutuhan dokumen. Jangan memaksakan format panjang.
4. **Jelaskan batas kata yang bisa ditafsirkan berbeda.** Jika memakai kata seperti *sah, cukup, relevan, jelas,* atau *berhasil* untuk suatu keputusan penting, sebutkan kondisi yang dapat diperiksa dan akibat jika tidak terpenuhi. Jangan menciptakan angka atau skor tanpa dasar pengukuran.
5. **Bedakan fakta, usulan, dan keputusan.** Jangan menaikkan kesimpulan riset atau pernyataan yang belum disetujui menjadi aturan yang mengikat.
6. **Jaga makna ketika menyunting.** Terjemahan tidak boleh menambah/mengurangi aturan, mengganti kewenangan, mengubah cakupan, atau menentukan cara teknis. Revisi dokumen yang sudah disetujui perlu dibandingkan dengan sumbernya dan ditinjau sebelum menggantikan sumber yang berlaku.

**Acuan istilah:** [Pilihan Istilah Product Kencleng](istilah.md) adalah satu tempat untuk padanan kata yang sudah disepakati; jangan menebak atau mengganti istilah secara massal.

**Tahap pertama selesai (22 Oktober 2026):** [Product Intent Bahasa Indonesia](product-intent.md) telah disetujui sebagai sumber utama Product pada branch riset ini. [Draf sebelumnya](product-intent-id-draft.md) disimpan untuk riwayat penyuntingan, bukan sumber kedua. Dokumen pilot lama tetap menjadi arsip sejarah.

**Tahap kedua selesai (22 Oktober 2026):** [Pengantar Peta Journey Product — Bahasa Indonesia](interconnected-journeys/README.md) telah disetujui dan dipromosikan sebagai sumber pengantar riset. [Draf sebelumnya](interconnected-journeys/README-id-draft.md) hanya riwayat penyuntingan, bukan sumber kedua. Persetujuan Level 2 v0.3 dan aturan O1 tidak berubah.

**Tahap ketiga selesai (22 Oktober 2026):** Anhar menyetujui dan mempromosikan [Peta Journey Level 2 v0.3 Bahasa Indonesia](interconnected-journeys/level-2-network.md) sebagai sumber riset Level 2 yang berlaku di branch ini. Peta tetap mencakup **26 node**, dua diagram, `EX-OPS`, **17 hubungan penting**, **17 pemeriksaan batas makna**, dan lima kelompok pertanyaan Product **OPEN**. OPEN **bukan** tuntutan untuk segera mengambil keputusan. [Draf Bahasa Indonesia sebelumnya](interconnected-journeys/level-2-network-id-draft.md) hanya riwayat penyuntingan. Persetujuan ini **bukan** persetujuan kontrak O1, delivery atau kesiapan Engineering.

## History and evidence

Older Product experiments, prior journey drafts, previous contract IDs, and discarded historical commitment rules are Git evidence only. The [identifier mapping](identifier-migration-2026-10-23.md) and immutable approved commits preserve their lineage. Historical approvals retain only their original scope and are not automatic instructions to future Product work.
