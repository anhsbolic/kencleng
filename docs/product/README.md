# Kencleng — Product Authority

> Status: **Candidate — not yet canonical**
> Promotion gate: human review plus successful forward/backward slice probes
> Purpose: Own Kencleng product/business truth independently of frontend/backend decomposition while preserving the canonical Product Design / Brand Authority as a peer upstream input to delivery.

This directory is being introduced as part of the Product Authority reframe recorded in:

`docs/project/kencleng-product-authority-reframe-audit.md`

Until this candidate is explicitly promoted, the current repository authorities remain operational. The candidate must therefore not be used to silently override existing specs/contracts during the migration.

## Intended authority boundary

When promoted, Product Authority will own:

- product purpose and boundaries;
- product thesis and durable product character;
- actors/personas in business terms;
- major capabilities;
- end-to-end user/business journeys;
- core business concepts and relationships;
- major lifecycle semantics;
- durable business permissions and invariants;
- trust/accountability semantics at product level;
- material unresolved product questions.

It will **not** own implementation decomposition such as:

- route/component structure;
- frontend folder/component architecture;
- endpoint/method names;
- exact request/response fields;
- database/table/column design;
- package/module boundaries;
- framework-specific mechanisms;
- low-level transaction/locking/crypto implementation;
- exact typography, colors, spacing, composition, iconography, illustration system, or other concrete visual-system rules.

Those concerns are derived later through design, delivery, architecture, and contract authorities.

## Product Authority and Product Design / Brand Authority are peer upstream inputs

The September 2026 design work demonstrated that design exploration can surface durable product insight, not merely visual choices.

Canonical design sources such as:

- `docs/ui-ux/product-design-principles.md`;
- `docs/ui-ux/brand-product-ui-brief.md`;
- `docs/ui-ux/page-map.md`;
- `docs/ui-ux/patterns.md`;
- `docs/ui-ux/design-guidelines.md`;
- `docs/ui-ux/asset-governance.md`;

remain active authorities for their design/experience concerns.

Some design discoveries are also durable product character — for example Evidence-Led Optimism, confidence before conversion, dignity over pity, trust through ordered evidence, and accountability continuing after donation. Those product-level implications belong in Product Authority, while their concrete experience and visual expression remains owned by `docs/ui-ux/`.

The target relationship is therefore not a one-way `product → design` chain. Both product/business truth and canonical product-design/brand authority constrain a delivery slice:

```text
              KENCLENG PRODUCT
                    │
        ┌───────────┴───────────┐
        ▼                       ▼
Product / Business       Product Design / Brand
Authority                 Authority
        │                       │
        └───────────┬───────────┘
                    ▼
              Delivery Slice
               ↙          ↘
        FE Delivery     BE Delivery
           Spec             Spec
               ↘          ↙
          Contract Reconciliation
                    ↓
             Shared Contract
                    ↓
          FE + BE implementation
                    ↓
            Integration evidence
                    ↓
        refine upstream authority
        when real evidence warrants
```

Design must not invent business truth. Product Authority must also not flatten away durable product insight merely because that insight was discovered during design exploration.

## Relationship to existing material

Existing artifacts are not discarded.

During the migration they are treated according to their concern:

- `docs/project/kencleng-business-process-overview.md`, `kencleng-actors-entities.md`, and product portions of phase docs → **product-source evidence**;
- `docs/ui-ux/` → **active Product Design / Brand Authority** plus a source of product discoveries when explicitly promoted into Product Authority;
- `docs/spec/<domain>/...` → **domain/delivery reference pending slice-by-slice reconciliation**;
- `api/openapi/` → **shared-contract reference pending slice-by-slice reconciliation**;
- `kencleng-erd.md` → **technical/data-model reference**;
- current backend/frontend code, tests, migrations, and Harscode artifacts → **implementation/workflow evidence**.

Reference does not mean incorrect. It means that the artifact does not automatically define higher-level product truth merely because it is more detailed.

## Candidate document

`product-overview.md` is the single candidate Product Authority owner for this reframe. It contains the whole-product thesis and character, actors, capability map, end-to-end journey model, trust/accountability model, progressive-delivery posture, and revalidation register.

Do not split those concerns into parallel product-authority files merely for neatness. Add another product document only when a genuinely distinct durable concern becomes large enough to need an independent owner.

## Promotion discipline

The candidate Product Authority should be promoted only after:

1. human review confirms that the whole-product model and product thesis represent the Kencleng we actually intend to build;
2. a forward slice demonstrates that Product + Design authorities can derive coherent FE/BE delivery needs and a shared contract without task-specific steering;
3. a backward slice demonstrates that existing detailed specs/code can be reconciled as evidence rather than blindly preserved or discarded;
4. contradictions between product and design assumptions are resolved explicitly;
5. root/scoped routing is updated deliberately so there is no long-lived dual authority.
