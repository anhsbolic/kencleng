# Run Invocation — `EXP-S2-001-001`

> Invocation prepared by Orchestration Operator on 2026-09-25. This Run has not been launched. Human-facing prose: Bahasa Indonesia; preserve canonical Harscode terms/enums and code/API/schema identifiers in English.

## Invocation identity

- `WORK_UNIT_ID`: `WU-S2-001`
- `RUN_ID`: `EXP-S2-001-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli` (runtime context dari `.harscode-spaces/.local-config.yaml`)
- `TARGET_REVISION`: `ee0d4b072d9f5cf279952fe309049f687c95e30e`
- `WORKFLOW_REVISION`: `b2d7ca4918b520d960139bc392f87619410b27ed`
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `Trigger`: current runnable frontier `WU-S2-001`, `NOT_STARTED`, `QUEUED`; new Slice 2 Exploration Run.

## Workflow routing and assignment

- `PHASE_ROUTE`: `REQUIRED` — canonical Harscode Exploration, initial phase for this Work Unit. Follow Stage 1 hard stop; Stage 2 and Stage 3 require the Human confirmations specified by the canonical prompt.
- `ROLE`: Explorer
- `SPECIALIZATION`: None established; do not infer one.
- `PARTICIPANT`: Codex Explorer — non-human executor assigned to this Run. Participant identity is distinct from the Codex CLI harness and from its Session.
- `SESSION`: Not created yet; create a fresh Participant Session at actual launch and record its identity then.
- `SESSION_ID`: Belum tersedia; Orchestrator mencatat ID aktual saat launch.
- `SESSION_TRANSITION`: `FRESH` at launch.
- `SESSION_TRANSITION_REASON`: This is a new workflow Run with no existing Run Session to continue.
- `CONTINUATION_CHECKPOINT`: Not applicable; this is not continuation of an active Run.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by the current Human-owned local registry (`approval_required: false`).
- `MODEL_ROUTING_RATIONALE`: Eksplorasi ini memetakan authority dan evidence lintas area untuk satu MVP slice. `gpt-6-luna` memiliki capability `reasoning` dan `repository-work` dengan `low` cost tier, sehingga memenuhi kebutuhan dengan biaya terendah yang tersedia. `medium` adalah effort terendah yang memadai untuk menyusun urutan area Stage 1 dan mendukung eksplorasi lintas area setelah Human confirmation; model `gpt-6-sol` yang berbiaya lebih tinggi dan memerlukan approval tidak diperlukan berdasarkan bukti saat ini.

## Current-effective inputs

- `PRIOR_ARTIFACTS`: `none` — belum ada Exploration artifact atau prior workflow Run untuk Work Unit ini.
- Orchestration context: `.harscode-spaces/s2-guest-donation-truthful-state/outcome.md`, `work-graph.md`, dan `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/manifest.md`.
- Product/MVP authority: `docs/product/README.md`, `docs/product/product-overview.md`, `docs/product/mvp-scope.md`, dan `docs/product/mvp-delivery-slices.md` (bagian Slice 2).
- Project routing/orchestration: root `AGENTS.md`, `docs/kencleng-agentic-workflow.md`, serta routing applicable dari target-repo authorities.
- Product Design routing: `docs/ui-ux/README.md`; ikuti hanya bila concerns Slice 2 memicu design authority.
- Communication profile: `docs/project/communication-profile.md`.
- Canonical Harscode phase: `../harscode-workspace/workflow/1-exploration-kickoff-prompt.md`.
- Orchestrated identity/path contract: `../harscode-workspace/workflow/orchestrated-run-overlay.md` dan `../harscode-workspace/orchestration/run-contract.md`.
- Thin Exploration wrapper: `../harscode-workspace/orchestration/exploration-kickoff-prompt.md`.

## Task and codebase context

`TASK`: Explore `Slice 2 — Guest Donation + Truthful Donation State` berdasarkan authority yang dirujuk di atas. Ikuti target-repository authority routing dan seluruh tahapan canonical Harscode Exploration. Mulai dengan Stage 1 plan announcement dan berhenti untuk Human confirmation sesuai canonical prompt. Setelah confirmation diberikan, lanjutkan tahapan berikutnya sesuai prompt canonical. Jangan mengasumsikan requirement, solution, authority, atau Work Unit/dependency topology yang belum didukung source/evidence.

`CODEBASE_CONTEXT`: Kencleng — Go backend + Next.js frontend.

`Ticket`: None.

`Area`: not sure yet; Stage 1 menentukan routing/area yang relevan.

## Canonical invocation

Jalankan wrapper tipis berikut tanpa mengubah phase semantics:

```text
Run the canonical Exploration contract from:
../harscode-workspace/workflow/1-exploration-kickoff-prompt.md

Apply orchestration identity/path semantics from:
../harscode-workspace/workflow/orchestrated-run-overlay.md

Work Unit: WU-S2-001
Run: EXP-S2-001-001
Run path: .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001
Role: Explorer
Specialization: None established
Participant: Codex Explorer
Session: assign a fresh Session at launch; record its ID
Communication language: Bahasa Indonesia
Communication profile: docs/project/communication-profile.md

Task: Explore Slice 2 — Guest Donation + Truthful Donation State under its approved Product/MVP authority and applicable target-repository routing. Use Stage 1 to state task understanding and identify exploration areas/order. Stop for Human confirmation as the canonical prompt requires. Continue later stages only after the corresponding Human confirmation.

Codebase context: Kencleng — Go backend + Next.js frontend
Ticket: none
Area: not sure yet
PRIOR_ARTIFACTS: none

For this Run, write durable Exploration evidence under RUN_PATH, not an ordinal legacy TASK_PATH phase directory. Follow the canonical Exploration prompt and the orchestrated-run overlay. Use Bahasa Indonesia for human-facing prose and preserve canonical Harscode terms/enums and code/API/schema identifiers. The communication profile is available at docs/project/communication-profile.md. Stage 1 is a hard stop. Do not steer Exploration toward expected gaps, Work Units, or solutions.
```

## Execution envelope

- `PREAUTHORIZED`: Baca routing/authority yang relevan dan evidence repository sesuai canonical Exploration. Participant boleh menulis artifact Exploration milik Run ini hanya di `RUN_PATH` setelah stage yang bersangkutan diizinkan oleh Human checkpoint. Sampaikan Stage 1 plan announcement; tidak perlu membuat durable artifact Stage 1 kecuali workflow meminta.
- `ORCHESTRATOR_DECISION`: Tidak ada wewenang untuk mengubah Work Unit state, Work Graph, Control Surface, atau project tracker. Laporkan rekomendasi/routing evidence kepada Orchestrator; Orchestrator yang merekonsiliasi state.
- `HUMAN_REQUIRED`: Human confirmation sebelum Stage 2, dan confirmation sebelum Stage 3 sebagaimana canonical prompt. Material authority decision, perubahan authority/spec/contract/implementation, protected-path write, atau keputusan risk acceptance memerlukan owner/gate yang berlaku dan berada di luar authorization Run ini.

## Expected Run boundary

- Pada first launch, jalankan Stage 1 saja dan berhenti untuk Human confirmation.
- Invocation preparation tidak mengubah scheduling: state tetap `QUEUED` sampai actual dispatch.
- Belum ada artifact atau klaim hasil Run. Setelah eksekusi, Participant menulis hanya evidence milik Run ini; Orchestrator merekonsiliasi verdict menjadi Work Unit state.
