# Kencleng — Product Authority

> Status: **Candidate — not yet canonical**
> Promotion gate: human review plus successful forward/backward slice probes
> Purpose: Own Kencleng product/business truth independently of frontend/backend decomposition.

This directory is being introduced as part of the Product Authority reframe recorded in:

`docs/project/kencleng-product-authority-reframe-audit.md`

Until this candidate is explicitly promoted, the current repository authorities remain operational. The candidate must therefore not be used to silently override existing specs/contracts during the migration.

## Intended authority boundary

When promoted, Product Authority will own:

- product purpose and boundaries;
- actors/personas in business terms;
- major capabilities;
- end-to-end user/business journeys;
- core business concepts and relationships;
- major lifecycle semantics;
- durable business permissions and invariants;
- trust/accountability semantics;
- material unresolved product questions.

It will **not** own implementation decomposition such as:

- route/component structure;
- frontend folder/component architecture;
- endpoint/method names;
- exact request/response fields;
- database/table/column design;
- package/module boundaries;
- framework-specific mechanisms;
- low-level transaction/locking/crypto implementation.

Those concerns are derived later through delivery/design/architecture/contract authorities.

## Target authority hierarchy

```text
Product / Business Authority
        ↓
Experience / Design Authority
        ↓
Delivery Slice
     ↙       ↘
FE Delivery  BE Delivery
   Spec         Spec
     ↘       ↙
 Contract Reconciliation
        ↓
 Shared Contract
        ↓
 FE + BE implementation
        ↓
 Integration evidence
```

## Relationship to existing material

Existing artifacts are not discarded.

During the migration they are treated according to their concern:

- `docs/project/kencleng-business-process-overview.md`, `kencleng-actors-entities.md`, and product portions of phase docs → **product-source evidence**;
- `docs/ui-ux/` → **active experience/design authority**;
- `docs/spec/<domain>/...` → **domain/delivery reference pending slice-by-slice reconciliation**;
- `api/openapi/` → **shared-contract reference pending slice-by-slice reconciliation**;
- `kencleng-erd.md` → **technical/data-model reference**;
- current backend/frontend code, tests, migrations, and Harscode artifacts → **implementation/workflow evidence**.

Reference does not mean incorrect. It means that the artifact does not automatically define higher-level product truth merely because it is more detailed.

## Candidate documents

- `product-overview.md` — whole-product purpose, actors, capability map, end-to-end journey model, trust/accountability model, and revalidation register.

Additional product documents should be added only when a distinct durable product concern genuinely needs its own owner. Do not rebuild the old all-domain spec tree at a higher directory level.
