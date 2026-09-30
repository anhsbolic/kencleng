# Slice 2 — Work Graph

> State saat ini, diturunkan dari authority dan bukti yang dirujuk Parent Outcome. File/layout ini adalah realisasi project-local Pilot #2, bukan skema storage Harscode canonical.

## Work Units yang diketahui

| ID | Work Unit | Type | Status | Milestone yang dihasilkan |
|---|---|---|---|---|
| `WU-S2-001` | Slice 2 Authority & Current-State Exploration | `ENABLER` | `DONE` | Evidence untuk menurunkan work rekonsiliasi Slice 2 |
| `WU-S2-002` | Slice 2 Donation Domain & Contract Reconciliation | `RECONCILIATION` | `ACTIVE` / `QUEUED` untuk rekonsiliasi Task 02 setelah O1 representation direction disetujui project-wide. O8 clear; O11 direction direkonsiliasi dan menuntut terminal notice dipenuhi dengan bounded/recoverable terminalization policy. Concrete currency parameters tetap deferred; tidak menahan arah amount wire. TP-011 perlu direkonsiliasi melalui jalur planning/review/approval yang diwajibkan sebelum `CONTRACT_READY`. Review spec Task 01 oleh Human tetap independen. Runtime evidence O2–O5 adalah downstream, bukan prasyarat contract. | `CONTRACT_READY` setelah prasyarat kebijakan/contract material dan review owner yang diwajibkan clear; runtime evidence tetap downstream jika hanya dapat dihasilkan saat implementasi/Testing |

## Dependency edges

| Dari | Ke | Strength | Kondisi |
|---|---|---|---|
| `WU-S2-002` | `WU-S2-001` | HARD | `WU-S2-001 = DONE`; evidence Exploration menjadi input rekonsiliasi |

## Runnable frontier

`WU-S2-002` tetap menjadi frontier Work Unit. Human menyetujui Techplan TP-007 pada 2026-09-27 dan amendmen material TP-011 pada 2026-09-30 setelah meninjau report TP-012. Planner Run TP-013 menyelaraskan status; TP-011 kini current-effective Approved. OIR-006 menetapkan D1; TP-010 dan TP-011 membawanya ke spine. Review `RV-S2-002-005` menemukan tiga blocker yang ditangani Planner dalam TP-011: O7 Resolved dengan provenance; gap owner/scope O1 dicatat sebagai scoped `AUTHORITY_SYNC`. Independent review baru `RV-S2-002-006` selesai tanpa findings. D1 tetap unchanged. Planner Run `TPD-S2-002-001` selesai dengan dua task: Task 01 untuk Donation domain specs tanpa dependency task dan Task 02 untuk Donation OpenAPI yang bergantung pada Task 01; Human menerima split pada 2026-09-30. Build `BLD-S2-002-001` menyelesaikan Task 01; `RV-S2-002-007` menemukan F-01, `BLD-S2-002-002` menerapkan koreksi Summary yang sempit, dan `RV-S2-002-008` mengonfirmasi penutupannya secara independen; review loop Task 01 selesai. Domain specs masih `draft` menunggu acceptance owner/Human yang berlaku. Build `BLD-S2-002-003` menyelesaikan assessment Task 02 dan OpenAPI lint tanpa diff authored/generated API; Code Review independen N/A karena tidak ada diff. Pemeriksaan batas fase Harscode menyimpulkan semantik publik O2 dan arah O4/O5 tidak memerlukan runtime proof sebelum authoring contract. Human/API owner mengonfirmasi bahwa operasi historis submit/status belum pernah didistribusikan eksternal, sehingga O8 clear dalam scope ini. O11 route B superseded: verified email tidak boleh dihapus sebelum terminal notice selesai dan Delivery harus menghasilkan terminalization policy bounded/recoverable; numeric bound, architecture, timeout-as-failed, dan residual-risk acceptance tidak dipilih. Anhar menjadi owner project-wide Shared Currency Standard dan menyetujui direction: major-unit decimal string + explicit currency code, exact-decimal calculation/persistence, Product-controlled active currency, serta whole-Rupiah IDR input untuk Slice 2. No universal precision/scale or additional currencies. O1 no longer blocks amount-field authoring. WU is `ACTIVE / QUEUED` for Task 02 reconciliation; TP-011 remains current-effective Approved until the required planning/review/approval route reconciles O1/O11 and `CONTRACT_READY` wording. Review Human Task 01 remains available independently. Evidence empiris O2–O5 is downstream. `CONTRACT_READY` belum tercapai.

## Batas derivasi

FE/BE delivery dan integration Work Unit belum diturunkan. Evidence mendukung rekonsiliasi contract lebih dahulu; turunkan implementasi setelah `CONTRACT_READY`.
