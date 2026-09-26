# WU-S2-001 — Slice 2 Authority & Current-State Exploration

## Definition

- Type: `ENABLER`
- Parent Outcome: `S2-GUEST-DONATION-TRUTHFUL-STATE`
- Coordination owner role: Orchestration Operator
- Primary execution role: Explorer
- Specialization: None established
- Communication language: Bahasa Indonesia
- Communication profile: `docs/project/communication-profile.md`

### Outcome

Menghasilkan evidence durable yang cukup untuk memahami authority Slice 2, current implementation/contract evidence yang relevan, material gap dan decision, serta delivery/reconciliation shape yang memang didukung evidence—tanpa mengunci solusi atau membuat authority baru.

### Scope

- Authority Product/MVP dan design yang dipicu oleh Slice 2.
- Relevansi dan status current dari domain invariants/threat concerns, feature specs, OpenAPI, implementation, dan tests terkait donation, memakai semuanya sebagai evidence sesuai aturan product-first.
- Material dependency, boundary, gap, dan unresolved authority question yang menentukan langkah delivery berikutnya.
- Evidence yang diperlukan untuk Orchestrator menurunkan Work Units dan dependency berikutnya.

### Out of scope

- Implementasi Slice 2.
- Techplan atau pemilihan solusi implementasi.
- Menetapkan endpoint/schema/architecture yang belum direkonsiliasi.
- Mengarang Product, security, interface, design, architecture, atau verification authority.
- Menetapkan FE/BE/dependency topology sebelum cukup evidence.

### Completion condition

Exploration menghasilkan evidence durable yang cukup untuk mengidentifikasi authority/gap/material decisions yang relevan dan memungkinkan Orchestrator menurunkan langkah delivery/reconciliation berikutnya tanpa mengarang authority. Material authority gap/conflict tetap dirutekan ke Human/authority owner.

## Current State

- Execution status: `ACTIVE`
- Scheduling state: `RUNNING`
- Horizon: `NOW`
- Current Run: `EXP-S2-001-001` — Stage 2 active after Human confirmation
- Current milestone: None
- Human gate: Human confirmation diperlukan setelah Stage 2 sebelum Stage 3, sesuai canonical Exploration kickoff
- Authority sync: Not applicable at bootstrap; belum ada authority change yang diajukan
- Active blocker: None identified
- Updated: 2026-09-25

## Current-effective artifacts

Invocation Orchestrator-prepared: `runs/EXP-S2-001-001/invocation.md`. Run launch evidence: `runs/EXP-S2-001-001/launch-record.md`. Stage 2 evidence belum selesai/current-effective.

## Routing note

Run `EXP-S2-001-001`, jika di-dispatch, harus memakai `workflow/1-exploration-kickoff-prompt.md` dan `workflow/orchestrated-run-overlay.md` canonical Harscode. Gunakan prompt kickoff tipis yang dirutekan melalui `orchestration/exploration-kickoff-prompt.md`; jangan menambahkan solution-steering conclusions. Runtime configuration lokal hanya konteks runtime; model registry Human-owned dan read-only bagi Orchestrator.
