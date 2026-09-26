# Slice 2 — Work Graph

> State saat ini, diturunkan dari authority dan bukti yang dirujuk Parent Outcome. File/layout ini adalah realisasi project-local Pilot #2, bukan skema storage Harscode canonical.

## Work Units yang diketahui

| ID | Work Unit | Type | Status | Milestone yang dihasilkan |
|---|---|---|---|---|
| `WU-S2-001` | Slice 2 Authority & Current-State Exploration | `ENABLER` | `DONE` | Evidence untuk menurunkan work rekonsiliasi Slice 2 |
| `WU-S2-002` | Slice 2 Donation Domain & Contract Reconciliation | `ACTIVE` / `QUEUED` | `CONTRACT_READY` setelah rekonsiliasi dan acceptance yang berlaku |

## Dependency edges

| Dari | Ke | Strength | Kondisi |
|---|---|---|---|
| `WU-S2-002` | `WU-S2-001` | HARD | `WU-S2-001 = DONE`; evidence Exploration menjadi input rekonsiliasi |

## Runnable frontier

`WU-S2-002` tetap menjadi frontier. Human menyetujui Techplan efektif pada 2026-09-26 setelah `RV-S2-002-003` bersih dan report tersedia; Planner menyelaraskan status ke `Approved`. Brief OIR selesai. Human lalu memperbarui dan menyetujui `mvp-scope.md` serta `mvp-delivery-slices.md`, sehingga arah O1/O3 kini tercakup dalam current Product/MVP authority bersama O2/O4/O5/O6/O9. Techplan `TP-S2-002-003` belum mencerminkan amendmen tersebut; invocation `TP-S2-002-006` siap untuk fresh Planner amendment (`gpt-6-luna` / `high`). Human-Assisted dispatch adalah runnable frontier. Setelah amendment, independent review diperlukan (Techplan tetap Complex/high-stakes), lalu Planner membuat report dan Human melakukan approval setelah review/resolution konvergen. O1/O2/O3/O4/O5 menyisakan detail/owner reviews; O6/O9 perlu diterjemahkan ke spec/API. O8 bersyarat sebelum breaking API change; O9 menahan final submit contract/`CONTRACT_READY`. Work Unit `ACTIVE` / `QUEUED`.

## Batas derivasi

FE/BE delivery dan integration Work Unit belum diturunkan. Evidence mendukung rekonsiliasi contract lebih dahulu; turunkan implementasi setelah `CONTRACT_READY`.
