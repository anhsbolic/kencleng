# Kencleng — Slice 1 Orchestration Control Surface

Protocol:
Harscode Orchestrator Protocol v0.1 — Pilot Candidate

Pilot Branch:
`validation-03-orchestrator-slice-1`

Project Communication Profile:
`docs/project/communication-profile.md`

Parent Outcome:
S1 — Public Campaign Understanding

Overall Status:
ACTIVE

## NOW

### WU-S1-003 — Slice 1 Backend Public Campaign Delivery
- Status: WAITING_HUMAN
- Current gate: BACKEND_VERIFIED promotion / Tier-1 Human gate
- Testing: `TST-BE-002` — PASS_WITH_FLAGGED_FOLLOWUPS
- Campaign-owned verification: sufficient
- Milestone: READY_FOR_HUMAN_PROMOTION, not yet promoted
- External flags: pre-existing repo gosec + runner-limited full race

### WU-S1-004 — Slice 1 Frontend Public Campaign Understanding
- Status: VERIFYING
- Current Run: `CR-CONF-FE-002`
- Role: Reviewer / targeted confirmation
- Patch: `BLD-FE-PATCH-002` — COMPLETED
- Testing: BLOCKED pending confirmation
- Milestone: NOT_YET_VERIFIED

### WU-S1-005 — Slice 1 Topology & Controlled Media Enablement
- Status: VERIFYING
- Current Run: `TST-TOP-002`
- Role: Verifier / joint runtime re-entry
- TST-TOP-001: PASS_WITH_FLAGGED_FOLLOWUPS
- Remaining: R4 200/404, Campaign-specific R5, R7 retraction fresh-fetch
- Milestone: NOT_YET_COMPLETE

Human Attention:
WU-S1-003 is ready for explicit Human promotion to BACKEND_VERIFIED. FE and TOP continue independently.

No new decision required at this point.

No new decision required. Unrelated pre-existing repository security-tool findings must not be absorbed into WU-S1-003 without explicit owner/authority routing.

None unless Testing/Patch uncovers a material contract or protected-write issue.

None unless Code Review raises a material decision or protected-write boundary.

None during Build unless an Implementer surfaces a material contract contradiction.

Human approval is now required for TP-BE-001, TP-FE-001, and TP-TOP-001 before any Build Run is dispatched.

Confirm Stage 1 independently in each Exploration session. No additional orchestration decision is required before starting them.

## COMPLETED

### WU-S1-002 — Slice 1 Public Contract & Delivery Reconciliation

Type:
RECONCILIATION

Status:
DONE

Milestone:
`CONTRACT_READY` — Human-approved 2026-09-22

Final Verification:
`TST-001` — PASS_WITH_FLAGGED_FOLLOW_UPS

Runtime milestones intentionally unclaimed:
- `BACKEND_VERIFIED`
- `FRONTEND_MOCK_VERIFIED`
- `INTEGRATED_VERIFIED`

Deferred non-blocking follow-up:
- `CR-001-C01` stale threat-model reference labels


### WU-S1-001 — Slice 1 Authority & Current-State Exploration

Type:
ENABLER

Status:
DONE

Completed Run:
`EXP-001`

Evidence:
- `WU-S1-001/runs/EXP-001/evidence/gap-analysis.md`
- `WU-S1-001/runs/EXP-001/evidence/solutioning.md`

## NEXT

### WU-S1-006 — Slice 1 Real Integration & Final Verification
- Type: VERIFICATION
- Status: NOT_STARTED
- Scheduling: PARKED
- Horizon: NEXT
- Readiness: DEFINED
- HARD dependencies:
  - WU-S1-003 → `BACKEND_VERIFIED`
  - WU-S1-004 → `FRONTEND_MOCK_VERIFIED`
  - WU-S1-005 → topology/private-media evidence complete
- Target: real same-origin integration evidence for `INTEGRATED_VERIFIED`, followed by Human Slice-finalization input.

## LATER

Belum committed.

## Human Attention

Tidak ada saat ini.

## Active Blockers

None.

## Stalled Work

Tidak ada.

## Pilot Observations

### OBS-ORCH-001 — Stage 1 routing specificity

Stage 1 berhasil memahami authority hierarchy, Run identity, communication profile, dan hard stop tanpa Human menulis custom kickoff prompt. Early specificity tidak berubah menjadi closed scope pada Stage 2.

### OBS-ORCH-002 — Minimal Human checkpoints

Stage 1 dan Stage 2 dapat dilanjutkan dengan Human response yang sangat singkat tanpa mengulang workflow semantics.

### OBS-ORCH-003 — Run / Session separation

`EXP-001` selesai dalam dua Session: Stage 1+2 pada Session pertama dan Stage 3 pada fresh Session. Durable evidence cukup untuk re-ground tanpa membuat Run baru.

### OBS-ORCH-004 — Exploration-derived decomposition

`WU-S1-002` diinstansiasi dari Exploration evidence. Backend/frontend/topology tetap provisional sampai reconciliation menutup contract ambiguity.

### OBS-ORCH-005 — Independent planning review gate

`TP-001` menghasilkan Draft Techplan dengan 14 Rules. Independent review tetap warranted karena plan crosses contracts, carries a breaking contract change, dan menyentuh security-sensitive public/media boundaries. Review dipisahkan ke fresh reviewer Session/Run agar independence nyata.

### OBS-ORCH-006 — Review finding routed without over-escalation

Independent review menemukan satu blocking security/interface gap yang sempit. Orchestrator merutekannya ke fresh Planner resolution Run dengan lower-cost sufficient model (`gpt-5.6-terra / medium`) karena problem sudah terlokalisasi dan tidak memerlukan fresh cross-cutting architecture synthesis.

### OBS-ORCH-007 — Narrow resolution avoided redundant re-review

`TPR-RES-001` menutup blocking finding dengan membuat exact allowlist executable melalui closed-object semantics. Karena field set, runtime behavior, dan material interface/security meaning tidak berubah, flow kembali langsung ke Human Techplan gate tanpa mandatory re-review.

### OBS-ORCH-010 — Orchestrator role bleed in Human report generation

Manual Orchestrator directly authored two `report-techplan.md` artifacts after clean independent review. Although the reports were derived from valid Techplans, this blurred coordination ownership with planning-artifact production and also bypassed the Kencleng Bahasa Indonesia communication profile. The reports were invalidated and removed; replacement generation is routed through explicit Planner Runs using the canonical report template and project communication profile. Candidate learning: orchestration should dispatch derived planning/report artifacts rather than silently authoring them when a participant/workflow boundary exists.

### OBS-ORCH-011 — Human report generation timing authority conflict

Harscode `workflow/2-techplan/report-template.md` says `report-techplan.md` is generated when the current-effective Techplan is ready to enter the Human approval gate after review/resolution convergence. `workflow/2-techplan/guardrails.md §8` says the report is generated only after Approval. The pilot encountered this conflict directly while preparing WU-S1-003/004/005 Human gates. Human approval subsequently occurred, so no Build contract is ambiguous, but the workflow sources should be reconciled after the pilot rather than relying on operator interpretation.

### OBS-ORCH-012 — Local runtime capability was under-specified

BLD-TOP-001 reported runtime Compose validation as unavailable because the `docker` CLI was absent. The Human clarified that the local development environment uses Podman and the repository already exposes `make up-podman` / `make down-podman`. The pilot local config now declares Podman as the container runtime and explicitly states that Docker CLI absence is not equivalent to container-runtime unavailability. Testcontainers compatibility remains probe-at-runtime rather than assumed.

## State Integrity

File ini adalah current-state projection dari orchestration records. Jika projection ini bertentangan dengan Work Unit / Run records, reconcile underlying orchestration state lalu refresh Control Surface.

## Pilot Discipline

- Jangan membuat downstream Work Unit hanya karena pernah muncul di pilot-preparation hypothesis.
- Jangan memperlakukan historical specs, contracts, migrations, tests, atau implementation sebagai authority yang lebih tinggi daripada current Product / Design authority.
- Jangan mengubah Orchestrator Protocol v0.1 selama pilot kecuali execution menemukan hard contradiction, authority-integrity problem, safety problem, atau pilot-blocking representational gap.
- Gunakan Bahasa Indonesia untuk human-facing prose, tetapi pertahankan canonical Harscode terms/enums serta code/API/schema identifiers dalam English.
