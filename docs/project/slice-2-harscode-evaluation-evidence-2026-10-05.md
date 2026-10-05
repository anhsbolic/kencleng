# Lampiran bukti — Evaluasi Slice 2 / Harscode

> Snapshot: 2026-10-05, setelah metadata HOLD diparkir dan sebelum event/report evaluasi ditambahkan. HEAD teramati: `7e731f92eb052e615bd64e84902f5c6fbe61cdbd` plus working tree. Inventaris ini adalah hasil pembacaan file, bukan log jam eksekusi atau hasil Testing baru.

## Metode dan batas hitungan

- Populasi: semua file reguler rekursif di `.harscode-spaces/s2-guest-donation-truthful-state/`; termasuk snapshots, task files, dan browser evidence historis. Report ini serta report evaluasi utama di `docs/project/` tidak termasuk hitungan.
- Run = direktori langsung di `WU-S2-*/runs/`, dibaca berdasarkan ID, bukan urutan filesystem. Jumlah direktori tidak menyatakan jumlah Run yang completed atau benar-benar dispatched.
- Evidence output = file di Run selain `invocation.md`. Keberadaan output tidak sendiri menetapkan verdict, terminal status, atau completion condition.
- Kata dihitung dengan pemisahan whitespace (`str.split`); ini bukan jumlah token model, unique knowledge, waktu baca, atau effort authoring. Byte mencakup semua format.
- Duplicate dihitung hanya bila seluruh byte dua file Markdown identik; perubahan metadata membuat salinan yang hampir sama tidak terhitung duplicate.
- Verifikasi yang dijalankan saat evaluasi: baca dokumen, inventaris/count, git status/log/revision, dan SHA-256. Tidak menjalankan tests, validators, code generators, SQL/migration, database, runtime, browser, atau security actions.

## Statistik snapshot

```json
{
  "run_directories": 126,
  "phase_counts": {
    "EXP": 7,
    "BLD": 20,
    "OIR": 6,
    "RV": 37,
    "TP": 49,
    "TPD": 3,
    "TST": 4
  },
  "space_files": 401,
  "space_markdown": 396,
  "space_bytes": 4403866,
  "markdown_whitespace_words": 507312,
  "extra_exact_duplicate_md": 3,
  "days_inclusive": 11,
  "elapsed_date_days": 10,
  "human_min_hours_self_report": 44,
  "categories": {
    "State / task / dokumen lain": 30,
    "Evidence Exploration / owner": 24,
    "Invocation": 126,
    "Launch record": 62,
    "Build / patch / laporan lain": 20,
    "Handoff": 39,
    "Review findings": 33,
    "Patch plan": 5,
    "Techplan / candidate": 28,
    "Report Techplan": 12,
    "Testing report": 4,
    "Salinan baseline source": 13
  },
  "category_words": {
    "State / task / dokumen lain": 84939,
    "Evidence Exploration / owner": 36406,
    "Invocation": 101409,
    "Launch record": 22292,
    "Build / patch / laporan lain": 16242,
    "Handoff": 14044,
    "Review findings": 26373,
    "Patch plan": 1427,
    "Techplan / candidate": 160374,
    "Report Techplan": 18380,
    "Testing report": 5752,
    "Salinan baseline source": 19674
  },
  "work_units": [
    {
      "wu": "WU-S2-001",
      "runs": 1,
      "phases": {
        "EXP": 1
      },
      "files": 5,
      "md": 5
    },
    {
      "wu": "WU-S2-002",
      "runs": 46,
      "phases": {
        "BLD": 7,
        "OIR": 6,
        "RV": 14,
        "TP": 17,
        "TPD": 2
      },
      "files": 153,
      "md": 153
    },
    {
      "wu": "WU-S2-003",
      "runs": 28,
      "phases": {
        "BLD": 2,
        "EXP": 1,
        "RV": 10,
        "TP": 13,
        "TPD": 1,
        "TST": 1
      },
      "files": 78,
      "md": 78
    },
    {
      "wu": "WU-S2-004",
      "runs": 11,
      "phases": {
        "BLD": 2,
        "EXP": 1,
        "RV": 3,
        "TP": 4,
        "TST": 1
      },
      "files": 35,
      "md": 32
    },
    {
      "wu": "WU-S2-005",
      "runs": 10,
      "phases": {
        "BLD": 2,
        "EXP": 1,
        "RV": 2,
        "TP": 4,
        "TST": 1
      },
      "files": 30,
      "md": 30
    },
    {
      "wu": "WU-S2-006",
      "runs": 22,
      "phases": {
        "BLD": 6,
        "EXP": 1,
        "RV": 6,
        "TP": 9
      },
      "files": 71,
      "md": 69
    },
    {
      "wu": "WU-S2-007",
      "runs": 7,
      "phases": {
        "BLD": 1,
        "EXP": 1,
        "RV": 2,
        "TP": 2,
        "TST": 1
      },
      "files": 21,
      "md": 21
    },
    {
      "wu": "WU-S2-008",
      "runs": 1,
      "phases": {
        "EXP": 1
      },
      "files": 2,
      "md": 2
    }
  ]
}
```

## Daftar Run dan keberadaan output

Keterangan khusus untuk Invocation-only: RV-S2-003-002 superseded sebelum dispatch; TP-S2-006-009 tidak di-dispatch dan status reconciliation kemudian dilakukan langsung; EXP-S2-008-001 prepared, undispatched, PARKED oleh Human HOLD. Run lain di tabel memiliki evidence output; tabel tidak mengaudit atau menyatakan seluruhnya completed.

### WU-S2-001

| Run | Prefix fase | Output selain Invocation |
|---|---|---|
| [EXP-S2-001-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/invocation.md) | EXP | 3 file: launch-record.md, stage-2-gap-analysis.md, stage-3-solutioning.md |

### WU-S2-002

| Run | Prefix fase | Output selain Invocation |
|---|---|---|
| [BLD-S2-002-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-001/invocation.md) | BLD | 2 file: launch-record.md, report.md |
| [BLD-S2-002-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-002/invocation.md) | BLD | 2 file: launch-record.md, patch-report-1.md |
| [BLD-S2-002-003](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-003/invocation.md) | BLD | 2 file: launch-record.md, report.md |
| [BLD-S2-002-004](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-004/invocation.md) | BLD | 2 file: launch-record.md, report.md |
| [BLD-S2-002-005](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-005/invocation.md) | BLD | 2 file: launch-record.md, patch-report-1.md |
| [BLD-S2-002-006](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-006/invocation.md) | BLD | 2 file: launch-record.md, report.md |
| [BLD-S2-002-007](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-007/invocation.md) | BLD | 2 file: launch-record.md, patch-report-1.md |
| [OIR-S2-002-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-001/invocation.md) | OIR | 1 file: resolution-brief.md |
| [OIR-S2-002-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-002/invocation.md) | OIR | 6 file: handoff.md, resolution-brief.md, stage-2-input-provenance.md, stage-2-o3-guest-email.md, stage-2-o4-status-credential.md, stage-2-o5-public-failure-parity.md |
| [OIR-S2-002-003](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-003/invocation.md) | OIR | 5 file: handoff.md, resolution-brief.md, stage-2-input-provenance.md, stage-2-o2-simulator.md, stage-2-o3-pending-retention.md |
| [OIR-S2-002-004](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-004/invocation.md) | OIR | 4 file: design-review-brief.md, handoff.md, stage-2-design-surface-gaps.md, stage-2-input-provenance.md |
| [OIR-S2-002-005](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-005/invocation.md) | OIR | 4 file: amount-contract-brief.md, handoff.md, stage-2-gap-analysis.md, stage-3-solutioning.md |
| [OIR-S2-002-006](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-006/invocation.md) | OIR | 3 file: campaign-donation-ordering-brief.md, handoff.md, stage-2-gap-analysis.md |
| [RV-S2-002-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-001/invocation.md) | RV | 2 file: launch-record.md, review-findings.md |
| [RV-S2-002-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-002/invocation.md) | RV | 2 file: launch-record.md, review-findings.md |
| [RV-S2-002-003](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-003/invocation.md) | RV | 2 file: launch-record.md, review-findings.md |
| [RV-S2-002-004](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-004/invocation.md) | RV | 2 file: launch-record.md, review-findings.md |
| [RV-S2-002-005](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-005/invocation.md) | RV | 2 file: launch-record.md, review-findings.md |
| [RV-S2-002-006](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-006/invocation.md) | RV | 2 file: launch-record.md, review-findings.md |
| [RV-S2-002-007](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-007/invocation.md) | RV | 3 file: launch-record.md, patch-plan.md, review-findings.md |
| [RV-S2-002-008](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-008/invocation.md) | RV | 2 file: launch-record.md, review-confirmation.md |
| [RV-S2-002-009](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-009/invocation.md) | RV | 2 file: launch-record.md, review-findings.md |
| [RV-S2-002-010](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-010/invocation.md) | RV | 2 file: launch-record.md, review-findings.md |
| [RV-S2-002-011](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-011/invocation.md) | RV | 3 file: launch-record.md, patch-plan-1.md, review-findings-1.md |
| [RV-S2-002-012](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-012/invocation.md) | RV | 2 file: launch-record.md, review-confirmation.md |
| [RV-S2-002-013](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-013/invocation.md) | RV | 2 file: patch-plan-1.md, review-findings-1.md |
| [RV-S2-002-014](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-014/invocation.md) | RV | 2 file: launch-record.md, review-confirmation.md |
| [TP-S2-002-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-001/invocation.md) | TP | 2 file: launch-record.md, techplan.md |
| [TP-S2-002-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-002/invocation.md) | TP | 2 file: launch-record.md, techplan.md |
| [TP-S2-002-003](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-003/invocation.md) | TP | 2 file: launch-record.md, techplan.md |
| [TP-S2-002-004](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-004/invocation.md) | TP | 2 file: launch-record.md, report-techplan.md |
| [TP-S2-002-005](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-005/invocation.md) | TP | 1 file: launch-record.md |
| [TP-S2-002-006](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-006/invocation.md) | TP | 2 file: launch-record.md, techplan.md |
| [TP-S2-002-007](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/invocation.md) | TP | 3 file: launch-record.md, report-techplan.md, techplan.md |
| [TP-S2-002-008](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-008/invocation.md) | TP | 1 file: launch-record.md |
| [TP-S2-002-009](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-009/invocation.md) | TP | 2 file: handoff.md, o2-delivery-proposal.md |
| [TP-S2-002-010](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-010/invocation.md) | TP | 2 file: launch-record.md, techplan.md |
| [TP-S2-002-011](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/invocation.md) | TP | 2 file: launch-record.md, techplan.md |
| [TP-S2-002-012](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-012/invocation.md) | TP | 2 file: launch-record.md, report-techplan.md |
| [TP-S2-002-013](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-013/invocation.md) | TP | 1 file: launch-record.md |
| [TP-S2-002-014](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-014/invocation.md) | TP | 2 file: launch-record.md, techplan.md |
| [TP-S2-002-015](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/invocation.md) | TP | 2 file: launch-record.md, techplan.md |
| [TP-S2-002-016](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-016/invocation.md) | TP | 2 file: launch-record.md, report-techplan.md |
| [TP-S2-002-017](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-017/invocation.md) | TP | 1 file: launch-record.md |
| [TPD-S2-002-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001/invocation.md) | TPD | 4 file: 01-donation-domain-spec-reconciliation.md, 02-donation-openapi-reconciliation.md, launch-record.md, manifest.md |
| [TPD-S2-002-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002/invocation.md) | TPD | 4 file: 01-donation-domain-spec-reconciliation.md, 02-donation-openapi-reconciliation.md, launch-record.md, manifest.md |

### WU-S2-003

| Run | Prefix fase | Output selain Invocation |
|---|---|---|
| [BLD-S2-003-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/BLD-S2-003-001/invocation.md) | BLD | 2 file: launch-record.md, report.md |
| [BLD-S2-003-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/BLD-S2-003-002/invocation.md) | BLD | 2 file: launch-record.md, report.md |
| [EXP-S2-003-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/EXP-S2-003-001/invocation.md) | EXP | 2 file: stage-2-gap-analysis.md, stage-3-solutioning.md |
| [RV-S2-003-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-001/invocation.md) | RV | 3 file: handoff.md, launch-record.md, review-findings.md |
| [RV-S2-003-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-002/invocation.md) | RV | 0 file: **Invocation-only** |
| [RV-S2-003-003](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-003/invocation.md) | RV | 1 file: review-findings.md |
| [RV-S2-003-004](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-004/invocation.md) | RV | 1 file: review-findings-1.md |
| [RV-S2-003-005](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-005/invocation.md) | RV | 2 file: launch-record.md, review-findings-1.md |
| [RV-S2-003-006](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-006/invocation.md) | RV | 1 file: review-findings-1.md |
| [RV-S2-003-007](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-007/invocation.md) | RV | 1 file: review-findings-1.md |
| [RV-S2-003-008](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-008/invocation.md) | RV | 2 file: launch-record.md, review-findings-1.md |
| [RV-S2-003-009](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-009/invocation.md) | RV | 2 file: launch-record.md, review-findings-1.md |
| [RV-S2-003-010](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-010/invocation.md) | RV | 2 file: launch-record.md, review-findings-1.md |
| [TP-S2-003-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-001/invocation.md) | TP | 2 file: launch-record.md, techplan.md |
| [TP-S2-003-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-002/invocation.md) | TP | 2 file: launch-record.md, techplan.md |
| [TP-S2-003-003](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-003/invocation.md) | TP | 3 file: handoff.md, launch-record.md, techplan.md |
| [TP-S2-003-004](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-004/invocation.md) | TP | 2 file: handoff.md, techplan.md |
| [TP-S2-003-005](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-005/invocation.md) | TP | 2 file: handoff.md, techplan.md |
| [TP-S2-003-006](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-006/invocation.md) | TP | 2 file: handoff.md, techplan.md |
| [TP-S2-003-007](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-007/invocation.md) | TP | 2 file: launch-record.md, report-techplan.md |
| [TP-S2-003-008](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-008/invocation.md) | TP | 1 file: handoff.md |
| [TP-S2-003-009](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-009/invocation.md) | TP | 1 file: handoff.md |
| [TP-S2-003-010](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-010/invocation.md) | TP | 1 file: handoff.md |
| [TP-S2-003-011](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-011/invocation.md) | TP | 2 file: launch-record.md, report-techplan.md |
| [TP-S2-003-012](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-012/invocation.md) | TP | 2 file: handoff.md, launch-record.md |
| [TP-S2-003-013](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-013/invocation.md) | TP | 2 file: handoff.md, launch-record.md |
| [TPD-S2-003-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TPD-S2-003-001/invocation.md) | TPD | 1 file: launch-record.md |
| [TST-S2-003-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TST-S2-003-001/invocation.md) | TST | 2 file: launch-record.md, testing-report-1.md |

### WU-S2-004

| Run | Prefix fase | Output selain Invocation |
|---|---|---|
| [BLD-S2-004-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/BLD-S2-004-001/invocation.md) | BLD | 2 file: launch-record.md, report.md |
| [BLD-S2-004-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/BLD-S2-004-002/invocation.md) | BLD | 2 file: launch-record.md, patch-report-001.md |
| [EXP-S2-004-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/EXP-S2-004-001/invocation.md) | EXP | 2 file: stage-2-gap-analysis.md, stage-3-solutioning.md |
| [RV-S2-004-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/RV-S2-004-001/invocation.md) | RV | 1 file: review-findings.md |
| [RV-S2-004-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/RV-S2-004-002/invocation.md) | RV | 2 file: patch-plan-001.md, review-findings-001.md |
| [RV-S2-004-003](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/RV-S2-004-003/invocation.md) | RV | 1 file: review-findings-001.md |
| [TP-S2-004-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TP-S2-004-001/invocation.md) | TP | 2 file: handoff.md, techplan.md |
| [TP-S2-004-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TP-S2-004-002/invocation.md) | TP | 2 file: handoff.md, techplan.md |
| [TP-S2-004-003](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TP-S2-004-003/invocation.md) | TP | 2 file: handoff.md, techplan.md |
| [TP-S2-004-004](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TP-S2-004-004/invocation.md) | TP | 2 file: launch-record.md, report-techplan.md |
| [TST-S2-004-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TST-S2-004-001/invocation.md) | TST | 5 file: .last-run.json, launch-record.md, playwright.config.ts, status-credential-handoff.spec.ts, testing-report-001.md |

### WU-S2-005

| Run | Prefix fase | Output selain Invocation |
|---|---|---|
| [BLD-S2-005-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/BLD-S2-005-001/invocation.md) | BLD | 2 file: handoff.md, report.md |
| [BLD-S2-005-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/BLD-S2-005-002/invocation.md) | BLD | 2 file: handoff.md, report.md |
| [EXP-S2-005-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001/invocation.md) | EXP | 2 file: stage-2-gap-analysis.md, stage-3-solutioning.md |
| [RV-S2-005-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/RV-S2-005-001/invocation.md) | RV | 2 file: handoff.md, review-findings.md |
| [RV-S2-005-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/RV-S2-005-002/invocation.md) | RV | 1 file: review-findings-1.md |
| [TP-S2-005-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/TP-S2-005-001/invocation.md) | TP | 1 file: techplan.md |
| [TP-S2-005-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/TP-S2-005-002/invocation.md) | TP | 2 file: handoff.md, techplan.md |
| [TP-S2-005-003](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/TP-S2-005-003/invocation.md) | TP | 2 file: handoff.md, report-techplan.md |
| [TP-S2-005-004](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/TP-S2-005-004/invocation.md) | TP | 2 file: handoff.md, launch-record.md |
| [TST-S2-005-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/TST-S2-005-001/invocation.md) | TST | 3 file: handoff.md, launch-record.md, testing-report-1.md |

### WU-S2-006

| Run | Prefix fase | Output selain Invocation |
|---|---|---|
| [BLD-S2-006-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-001/invocation.md) | BLD | 1 file: report.md |
| [BLD-S2-006-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-002/invocation.md) | BLD | 8 file: 01-campaign-creation-draft-crud.md, 01-submit-donation-settlement.md, 02-campaign-detail-listing.md, 09-closure.md, invariants.md, report.md, source-delta.patch |
| [BLD-S2-006-003](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-003/invocation.md) | BLD | 9 file: 01-submit-donation-settlement.md, 09-closure.md, campaign-09-closure.md, campaign-invariants.md, donation-01-submit-donation-settlement.md, donation-invariants.md, invariants.md, patch-report-1.md, source-delta.patch |
| [BLD-S2-006-004](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-004/invocation.md) | BLD | 1 file: report.md |
| [BLD-S2-006-005](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-005/invocation.md) | BLD | 1 file: report.md |
| [BLD-S2-006-006](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-006/invocation.md) | BLD | 1 file: report.md |
| [EXP-S2-006-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/EXP-S2-006-001/invocation.md) | EXP | 3 file: handoff.md, stage-2-gap-analysis.md, stage-3-solutioning.md |
| [RV-S2-006-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-001/invocation.md) | RV | 1 file: review-findings.md |
| [RV-S2-006-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-002/invocation.md) | RV | 1 file: review-findings.md |
| [RV-S2-006-003](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-003/invocation.md) | RV | 2 file: patch-plan.md, review-findings.md |
| [RV-S2-006-004](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-004/invocation.md) | RV | 1 file: review-findings.md |
| [RV-S2-006-005](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-005/invocation.md) | RV | 1 file: review-findings.md |
| [RV-S2-006-006](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-006/invocation.md) | RV | 1 file: review-findings.md |
| [TP-S2-006-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-001/invocation.md) | TP | 2 file: handoff.md, techplan.md |
| [TP-S2-006-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-002/invocation.md) | TP | 2 file: handoff.md, techplan.md |
| [TP-S2-006-003](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-003/invocation.md) | TP | 2 file: handoff.md, techplan.md |
| [TP-S2-006-004](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-004/invocation.md) | TP | 3 file: handoff.md, report-techplan.md, techplan.md |
| [TP-S2-006-005](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-005/invocation.md) | TP | 2 file: handoff.md, report-techplan.md |
| [TP-S2-006-006](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-006/invocation.md) | TP | 1 file: handoff.md |
| [TP-S2-006-007](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-007/invocation.md) | TP | 2 file: handoff.md, techplan.md |
| [TP-S2-006-008](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-008/invocation.md) | TP | 3 file: handoff.md, report-techplan.md, techplan.md |
| [TP-S2-006-009](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-009/invocation.md) | TP | 0 file: **Invocation-only** |

### WU-S2-007

| Run | Prefix fase | Output selain Invocation |
|---|---|---|
| [BLD-S2-007-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/BLD-S2-007-001/invocation.md) | BLD | 1 file: report.md |
| [EXP-S2-007-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/EXP-S2-007-001/invocation.md) | EXP | 2 file: stage-2-gap-analysis.md, stage-3-solutioning.md |
| [RV-S2-007-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/RV-S2-007-001/invocation.md) | RV | 2 file: launch-record.md, review-findings-1.md |
| [RV-S2-007-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/RV-S2-007-002/invocation.md) | RV | 2 file: handoff.md, review-findings-1.md |
| [TP-S2-007-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/TP-S2-007-001/invocation.md) | TP | 1 file: handoff.md |
| [TP-S2-007-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/TP-S2-007-002/invocation.md) | TP | 1 file: handoff.md |
| [TST-S2-007-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/TST-S2-007-001/invocation.md) | TST | 1 file: testing-report-1.md |

### WU-S2-008

| Run | Prefix fase | Output selain Invocation |
|---|---|---|
| [EXP-S2-008-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-008/runs/EXP-S2-008-001/invocation.md) | EXP | 0 file: **Invocation-only** |

## Anchor sumber utama analisis

- [Timeline Events](../../.harscode-spaces/s2-guest-donation-truthful-state/events.md), [Work Graph](../../.harscode-spaces/s2-guest-donation-truthful-state/work-graph.md), manifests setiap WU dan tracker: dasar state/timeline; Events dibaca bersama exact Run evidence untuk klaim sebab.
- [First launch EXP-S2-001-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/launch-record.md): tanggal mulai 2026-09-25; launch 14:13:34 UTC / 21:13:34 Asia/Jakarta.
- [BLD-S2-003-001](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/BLD-S2-003-001/report.md): bounded backend cap projection, unapplied 000012, lalu STALLED.
- [BLD-S2-003-002](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/BLD-S2-003-002/report.md): prerequisite migration-design Review belum terpenuhi, tanpa code/migration write.
- [RV-S2-003-005](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-005/review-findings-1.md): enam blocking schema-design findings.
- [RV-S2-003-006](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-006/review-findings-1.md): authority mismatch concrete credential controls.
- [RV-S2-003-008](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-008/review-findings-1.md): required token pada retry tidak dapat direproduksi dari verifier.
- [RV-S2-003-009](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-009/review-findings-1.md) dan [RV-S2-003-010](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-010/review-findings-1.md): lifecycle attribution, mechanical pointer, batas OI9.
- [WU-S2-004 manifest](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/manifest.md): frontend mock outcome dan follow-ups, Review/Testing/Human rendered receipt.
- [WU-S2-007 handoff](../../.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/handoff-to-WU-S2-003.md): exact acceptance Funding-unavailable 503.
- [Stage A baseline](../../.harscode-spaces/s2-guest-donation-truthful-state/experiments/current-state-simplification/stage-a-baseline.md): pengulangan state sudah diamati pada 2026-09-29.
