# WU-S1-001 — Slice 1 Authority & Current-State Exploration

Type:
ENABLER

Parent Outcome:
S1 — Public Campaign Understanding

Status:
NOT_STARTED

Scheduling:
QUEUED

Horizon:
NOW

Coordination Owner Role:
Orchestration Operator

Primary Execution Role:
Explorer

Specialization:
None

Communication Language:
Bahasa Indonesia

Communication Profile Path:
`docs/project/communication-profile.md`

## Outcome

Menghasilkan durable evidence yang cukup untuk memahami authority, current implementation, existing contracts, dan material gaps yang relevan terhadap Kencleng MVP Slice 1 tanpa prematurely designing implementation.

## Scope

- Product Authority yang relevan terhadap Slice 1.
- Product Design / Brand Authority yang relevan terhadap Public Campaign Understanding.
- Existing delivery/domain specifications yang relevan terhadap slice.
- Existing shared API contract evidence.
- Relevant backend implementation evidence.
- Relevant frontend implementation evidence.
- Applicable reusable project evidence / Project Learning ketika discoverable dan sufficiently fresh.

## Out of Scope

- Mengimplementasikan Slice 1.
- Menulis detailed Techplan.
- Mengarang unresolved Product, Design, Security, atau shared-contract semantics.
- Menentukan backend/frontend decomposition lebih awal.
- Menentukan shared contract changes lebih awal.
- Membuat downstream Work Unit tanpa evidence.
- Memperlakukan pilot-preparation hypothesis sebagai delivery authority.

## Authority Entry Points

Primary routing sources:

- `docs/product/README.md`
- `docs/product/product-overview.md`
- `docs/product/mvp-scope.md`
- `docs/product/mvp-delivery-slices.md`
- `docs/ui-ux/README.md`

Explorer harus mengikuti target-repo routing dari sources tersebut dan applicable `AGENTS.md` files, bukan membaca semua adjacent documents secara default.

## Current Run

`EXP-001`

Run State:
PLANNED

Runtime Harness:
`codex-cli`

Selected Model:
`gpt-5.6-sol`

Reasoning Effort:
`high`

Model Approval:
APPROVED_BY_HUMAN for `EXP-001` only

Run Path:

`.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001`

## Human Escalation

Applicable Human Authority:
Anhar

Orchestration Operator:
Anhar selama v0 pilot

Material authority gap atau conflict harus surfaced untuk human decision, bukan diselesaikan oleh Explorer.

## Completion

Work Unit ini complete ketika Exploration telah menghasilkan durable evidence yang cukup untuk menentukan real downstream delivery / reconciliation structure, termasuk material authority gaps, tanpa mengarang unresolved authority.

## Communication

Human-facing prose menggunakan Bahasa Indonesia.

Canonical Harscode terms/enums serta code/API/schema identifiers tetap dalam English.

## Pilot Notes

Work Unit ini sengaja lebih luas daripada concern backend atau frontend.

Tujuannya adalah menemukan real Slice 1 delivery shape sebelum downstream Work Units committed.
