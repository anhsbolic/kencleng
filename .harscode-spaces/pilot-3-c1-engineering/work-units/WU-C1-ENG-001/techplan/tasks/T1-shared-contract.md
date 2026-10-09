# T1 — Shared OpenAPI contract

> Phase: Techplan decomposition; Author: `PARTICIPANT-C1-ENG-DECOMPOSER-001`; Created/Updated: 2026-10-09. Model: `gpt-6-luna`; Reasoning: medium; Session: `SESSION-C1-ENG-DECOMPOSER-001` (Run binding). Target revision: `524ef600c7f71246af6b71671d89c3c040fd9d44`; workflow revision: `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

## Purpose and outcome

Own the coordinated shared OpenAPI source/bundle and generated types for the approved C1 boundary. Capture §8 operations, request/response shapes, authority, error meanings, and privacy semantics so backend and frontend can implement against the same contract. This task makes no consumer production changes.

## Authority

- Parent: `../techplan.md`, SHA-256 `bd732748b486f0327f6cba2cdf14cab32f1ece3c5be8ebb24f9659444947fa47`, status Approved.
- Governing parent items: Q1–Q7; R1–R14; D1, D2, D3, D4, D5, D6, D7, D9; RISK-1–RISK-8; §8 interface contract; §11 verification and §13 Active item 5.
- Shared API concern owner coordinates the contract with backend and frontend. Root and project authority remain applicable. Parent spine controls all meaning and remains read-only.

## Scope

Update `api/openapi/index.yaml` and its derived bundle/types using existing API tooling. Include auth start/callback, `/me`, logout, preparation/read/confirm, my-organizations and organization detail; exact authority and same-person/object-scope semantics; consequence receipt/version and OrganizationView; safe errors/status meanings including indeterminate outcome; private/no-store behavior and callback constraints. Preserve `/api` browser prefix handling as described in the parent. Validate/bundle and generate types using the repository's current tooling.

Coordinate ownership and review with the backend and frontend concern owners. Do not change backend or frontend production files in this task. Never create an independent handwritten frontend model in parallel with generated types.

## Dependencies

No hard task dependency. Parent §8 and current API tooling are direct inputs. Contract publication is this task's result; downstream tasks depend only on the observable contract artifacts and meanings specified in the manifest.

## Verification and handoff

Run API validate/bundle and type-generation commands available in the live repo; record exact commands/results and artifact revision. Review contract coverage against parent §8 and ensure no route grants authority by caller-supplied identity/role/override, no false verification meaning, and no leakage of provider subject/token/email. This is authored Build evidence, not final independent Testing or Human semantic acceptance.

## Explicit boundaries

No product meaning, API semantics, security rule, verification obligation, or unresolved item may be added/changed here to resolve a gap. Route a material gap to parent reconciliation. No production implementation, migrations, provider verification, real runtime evidence, approval, or C1 completion claim.
