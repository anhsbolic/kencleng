# Slice 2 — Work Graph

> State saat ini, diturunkan dari authority dan bukti yang dirujuk Parent Outcome. File/layout ini adalah realisasi project-local Pilot #2, bukan skema storage Harscode canonical.

## Work Units yang diketahui

| ID | Work Unit | Type | Status | Milestone yang dihasilkan |
|---|---|---|---|---|
| `WU-S2-001` | Slice 2 Authority & Current-State Exploration | `ENABLER` | `DONE` | Evidence untuk menurunkan work rekonsiliasi Slice 2 |
| `WU-S2-002` | Slice 2 Donation Domain & Contract Reconciliation | `RECONCILIATION` | `ACTIVE` / `QUEUED` untuk fresh Planner reconciliation `TP-S2-002-014`. O1/O11 decisions post-date TP-011; O8 compatibility was also cleared after approval. New spine must pass independent review/report/Human approval before Task 02 execution. After plan approval, reconcile only affected task snapshots from TPD-001 and retain the accepted split, dependency, and manifest unless topology/dependency changes materially; Human split review reopens only for such a change. Task 01 Human review/acceptance remains independent and parallel. | `CONTRACT_READY` setelah prasyarat kebijakan/contract material dan review owner yang diwajibkan clear; runtime evidence tetap downstream jika hanya dapat dihasilkan saat implementasi/Testing |

## Dependency edges

| Dari | Ke | Strength | Kondisi |
|---|---|---|---|
| `WU-S2-002` | `WU-S2-001` | HARD | `WU-S2-001 = DONE`; evidence Exploration menjadi input rekonsiliasi |

## Runnable frontier

`WU-S2-002` tetap menjadi frontier Work Unit. Human menyetujui TP-011 pada 2026-09-30 setelah report TP-012; TP-013 menyelaraskan status. Keputusan O1/O11 dan clearance O8 datang sesudah approval TP-011 sehingga tidak boleh menjadi asumsi Build yang hanya ada di projection/task files. OIR-006 menetapkan D1; TP-011 membawanya ke spine. RV-006 selesai bersih. TPD-001 membentuk Task 01 domain specs dan Task 02 OpenAPI bergantung Task 01; Human menerima split. Build/review loop Task 01 selesai namun specs tetap `draft` menunggu owner/Human acceptance; review itu tetap jalur paralel independen. BLD-003 menyelesaikan assessment Task 02 dan lint tanpa diff authored/generated API. O8 clear berdasarkan bounded Human/API owner confirmation. O11 route B superseded: verified email retained through terminal notice and Delivery must produce bounded/recoverable terminalization, tanpa memilih numeric bound, architecture, timeout-as-failed, atau residual-risk acceptance. O1 direction disetujui project-wide dan dicatat di standard kanonik; precision/range/fraction/storage specifics tetap deferred. Fresh Planner `TP-S2-002-014` disiapkan untuk membawa keputusan O1/O8/O11 ke spine baru dengan provenance sendiri; TP-011 tetap current-effective Approved predecessor sampai Human menyetujui spine baru. Karena material, route berikutnya adalah independent Complex Review, resolution/re-review bila perlu, Planner report setelah konvergen, lalu Human approval. Setelah approval, hanya task snapshots terdampak dari TPD-001 direkonsiliasi; split/dependency/manifest yang sudah diterima tetap berlaku kecuali topology/dependency berubah material, dan Human split review dibuka kembali hanya bila perubahan itu terjadi. Tidak ada `CONTRACT_READY` atau delivery milestone.

## Batas derivasi

FE/BE delivery dan integration Work Unit belum diturunkan. Evidence mendukung rekonsiliasi contract lebih dahulu; turunkan implementasi setelah `CONTRACT_READY`.
