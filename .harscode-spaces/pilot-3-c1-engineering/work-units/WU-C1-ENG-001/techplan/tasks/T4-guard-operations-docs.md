# T4 — Guard operating documentation

> Phase: Techplan decomposition; Author: `PARTICIPANT-C1-ENG-DECOMPOSER-001`; Created/Updated: 2026-10-09. Model: `gpt-6-luna`; Reasoning: medium; Session: `SESSION-C1-ENG-DECOMPOSER-001` (Run binding). Target revision: `524ef600c7f71246af6b71671d89c3c040fd9d44`; workflow revision: `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

## Purpose and outcome

Create the scoped guard runbook and backend operating/configuration examples required to operate T3's restricted guard safely, coordinated with the backend concern owner. Explain provisioning, controlled state changes, evidence recording, monitoring, and rollback within the approved semantics.

## Authority

- Parent: `../techplan.md`, SHA-256 `bd732748b486f0327f6cba2cdf14cab32f1ece3c5be8ebb24f9659444947fa47`, status Approved.
- Governing items: R7, R12, R14; D4, D10; RISK-5, RISK-7; §§8–9 guard contract; §10 DB/Compose/Podman anchors; §11 verification; §12 G3; §13 Active item 5.
- G3 authorization is recorded in WU `events.md` and parent Resolved item 11. Docs must not extend it.

## Hard dependency: T3

**Condition:** Documented commands, roles and examples must match the guard schema/control mechanism and privileges actually implemented by T3, including default-closed initialization, constrained runtime SELECT/UPDATE permission, operator transaction/audit, and rollback behavior. Durable evidence: T3 schema/migration and control implementation/config revision, plus its recorded actual privilege check. Drafting stable semantic sections may proceed earlier; executable instructions must not cross this boundary before those artifacts exist.

## Scoped implementation

Create `docs/project/c1-establishment-guard-runbook.md` and scoped backend setup/config examples. Describe parameterized operator transaction, revision check, atomic append-only audit, accepted input channel, recording outstanding holds, initial opening evidence, reopening only when all holds no longer apply per external owning evidence, safe monitoring, fail-closed repair, and rollback that preserves established Organizations/receipts. Coordinate shared operational docs explicitly with backend/operator concern owners. The guard is an operational availability control, not legal adjudication or independent verification.

## Verification and handoff

Review every executable command/role/permission against T3's durable implementation; check that examples cannot alter audit/state through runtime credentials or imply that opening verifies the outside world. Record document review references. Operator/provider/environmental evidence remains downstream and is not supplied by documentation.

## Explicit boundaries

No new API, admin UI, application override, legal/conflict-resolution judgment, automatic reopen, new conflict source, or claim that a guard has been initialized/opened. No backend implementation or frontend production writes. Any mismatch with parent or T3 implementation routes to owning concern/parent reconciliation.
