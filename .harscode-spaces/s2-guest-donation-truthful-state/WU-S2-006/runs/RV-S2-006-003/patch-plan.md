# Patch Plan — RV-S2-006-003

> Phase: Code Review patch plan  
> Author: P-S2-006-RV-003-1 (KC-REVIEWER)  
> Created: 2026-10-02  
> Updated: 2026-10-02  
> Model: Invocation-configured `gpt-6-luna`  
> Reasoning: high  
> Session: fresh; Session ID not exposed  
> Target revision: `7fd8b473b239b20bda3990ab29c51440d321a796` plus exact six-file Run delta  
> Workflow revision: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`  
> Work Unit / Run: `WU-S2-006` / `RV-S2-006-003`

## Tujuan

Selaraskan aturan penutupan capacity dengan Product dan Techplan, lalu perbaiki checklist penutupan yang masih memakai jumlah alasan lama. Patch sumber harus tetap berada pada enam file spec dalam scope WU-S2-006; tidak ada perubahan API, generated artifact, backend/frontend, migration, atau test dalam patch ini.

## Perubahan yang diminta

1. **Temuan blocking C-01 — tutup saat tak ada Donasi valid yang muat.** Di Campaign INV-campaign-13 dan feature `09-closure.md`, definisikan capacity close terjadi ketika sisa kapasitas setelah settled Funding dan accepted-pending reservations tidak cukup untuk minimum Donation valid Rp5.000, termasuk sisa nol. Tegaskan bahwa admission tetap menolak Donation penuh yang tidak muat dan bahwa capacity close tidak mengubah aturan threshold `max_amount`. Perbarui Donation INV-donation-02/08 dan feature `01-submit-donation-settlement.md` agar memakai trigger yang sama, tidak hanya “admission exhausts capacity”. Pertahankan aturan bahwa pelepasan reservation karena Donation gagal tidak membuka Campaign kembali.
2. **Tambahkan bukti batas.** Perbarui verification/checklist terkait untuk membedakan sisa Rp4.999 (tidak ada nominal valid yang muat, Campaign ditutup) dari Rp5.000 (nominal minimum masih dapat muat, Campaign belum capacity-closed), selain kapasitas tepat habis. Pastikan penolakan nominal yang terlalu besar tidak membocorkan sisa kapasitas ketika nominal lebih kecil masih dapat diterima.
3. **Temuan non-blocking Q-01 — checklist `closed_by`.** Di `docs/spec/4-campaign/features/09-closure.md`, ganti “other two reasons” dengan “all non-admin close reasons” atau enumerasi `max_amount_reached`, `deadline_reached`, dan `funding_capacity_reached`; verifikasi nilai `closed_by` untuk capacity reason juga.

## Batas patch dan handoff

- Perubahan hanya di sumber spec berikut: `docs/spec/4-campaign/invariants.md`, `docs/spec/4-campaign/features/09-closure.md`, `docs/spec/5-donation/invariants.md`, dan `docs/spec/5-donation/features/01-submit-donation-settlement.md`.
- Jangan mengubah Product Authority untuk menyamakan dengan teks spec; Product dan Techplan sudah menyatakan aturan “no additional valid amount can fit”.
- Review ini tidak menentukan mekanisme transaksi/locking, API transport response kapasitas, maupun detail Slice-3. Pertahankan kepemilikan downstream tersebut.
- Sesudah patch, lakukan review/source-owner acceptance sesuai gate WU-S2-006; API/counterpart reconciliation baru mengikuti sumber yang diterima. Runtime concurrency dan settlement evidence tetap dimiliki Testing/delivery.
- Tidak ada patch produksi atau verification runtime yang dilakukan pada Run Review ini.
