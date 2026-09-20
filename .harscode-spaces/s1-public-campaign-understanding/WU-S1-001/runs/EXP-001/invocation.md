# EXP-001 — Orchestrated Exploration Invocation

Status:
READY_FOR_REVIEW

Prepared By Role:
Orchestration Operator

Work Unit:
WU-S1-001

Run:
EXP-001

Role:
Explorer

Specialization:
None

Run Path:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001`

Work Unit Path:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-001`

Prior Artifacts:
None

Target Branch:
`validation-03-orchestrator-slice-1`

Workflow Branch:
`pilot/orchestrator-v0.1`

Communication Language:
Bahasa Indonesia

Communication Profile Path:
`docs/project/communication-profile.md`

## Operator Binding Before Dispatch

Bind:

`HARSCODE_WORKSPACE_ROOT`

ke local checkout `harscode-workspace` yang sedang berada di branch:

`pilot/orchestrator-v0.1`

Ini hanya environment/path binding. Jangan rewrite Task dan jangan menambahkan solution-steering instructions.

## Entrypoint

Gunakan:

`{HARSCODE_WORKSPACE_ROOT}/orchestration/exploration-kickoff-prompt.md`

Wrapper tersebut harus menerapkan:

- canonical Exploration authority dari `workflow/1-exploration-kickoff-prompt.md`;
- orchestrated identity/path semantics dari `workflow/orchestrated-run-overlay.md`.

## Codebase Context

`Kencleng — Go backend + Next.js frontend`

## Communication Directive

Human-facing prose:
Bahasa Indonesia

Preserve in English:

- canonical Harscode terms;
- protocol enums/status/type values;
- code/API/schema identifiers;
- file paths, branch names, commit SHAs, dan CLI commands.

Jika exact wording dari authority source materially important, pertahankan wording sumber apa adanya.

## Task

Eksplorasi apa saja yang diperlukan untuk deliver Kencleng MVP Slice 1 — Public Campaign Understanding — berdasarkan current authoritative Product dan Product Design sources serta existing project specifications, shared contracts, backend/frontend implementation, dan current evidence lain yang relevan.

Governing slice source:

`docs/product/mvp-delivery-slices.md`

Jangan mengasumsikan historical specs, contracts, migrations, tests, atau implementation otomatis menjadi current authority. Surface material gaps atau conflicts melalui authority yang memiliki concern tersebut.

Jangan mengasumsikan downstream Work Unit graph, frontend/backend split, atau contract changes lebih awal. Derive hanya hal yang didukung current authority dan evidence.

## Optional Routing Inputs

Ticket:
None

Area:
Not sure yet — Stage 1 harus menentukan relevant areas.

## Dispatch Contract

Jalankan canonical Exploration contract dengan:

- `WORK_UNIT_ID = WU-S1-001`
- `RUN_ID = EXP-001`
- `RUN_PATH = .harscode-spaces/s1-public-campaign-understanding/WU-S1-001/runs/EXP-001`
- `WORK_UNIT_PATH = .harscode-spaces/s1-public-campaign-understanding/WU-S1-001`
- `ROLE = Explorer`
- `SPECIALIZATION = none`
- `PRIOR_ARTIFACTS = none`
- `COMMUNICATION_LANGUAGE = Bahasa Indonesia`
- `COMMUNICATION_PROFILE_PATH = docs/project/communication-profile.md`
- `TASK = Task section di atas`
- `CODEBASE_CONTEXT = Kencleng — Go backend + Next.js frontend`

## Mandatory First Stop

Execute **Stage 1 — Plan Announcement only**.

Jangan lanjut ke Stage 2 sampai human checkpoint mengonfirmasi Stage 1 understanding dan exploration routing.

## Invocation Review Checklist

Sebelum dispatch, verifikasi:

- [ ] Work Unit yang dipilih benar.
- [ ] Exploration adalah next workflow phase yang benar.
- [ ] Explorer adalah Role yang benar.
- [ ] Tidak ada downstream solution / Work Unit hypothesis yang di-inject sebagai fact.
- [ ] Product / Design authority tetap upstream terhadap historical implementation evidence.
- [ ] Run Path unik untuk `EXP-001`.
- [ ] Prior Artifacts benar-benar `none`.
- [ ] Communication Directive mengikuti Kencleng communication profile tanpa menerjemahkan canonical protocol semantics.
- [ ] Satu-satunya operator-specific mutation adalah binding `HARSCODE_WORKSPACE_ROOT`.
