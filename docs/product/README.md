# Kencleng — Product Authority

> Status: **Candidate — not yet canonical**
> Promotion gate: human-reviewed whole-product model, validated product hierarchy, approved MVP scope, and deliberate repository-routing promotion.
> Purpose: Own Kencleng product/business truth independently of frontend/backend decomposition while preserving canonical Product Design / Brand Authority as a peer upstream input.

This directory is part of the Product Authority reframe recorded in:

`docs/project/kencleng-product-authority-reframe-audit.md`

Until this candidate is explicitly promoted, current repository authorities remain operational. Candidate product material must not silently override existing specs/contracts during migration.

## Authority boundary

When promoted, Product Authority owns durable product/business truth such as:

- product purpose and boundaries;
- product thesis and character;
- actors/personas;
- major capabilities and end-to-end journeys;
- core business concepts and relationships;
- major lifecycle semantics;
- durable permissions/invariants;
- trust/accountability semantics;
- material unresolved product questions.

It does **not** own implementation decomposition such as routes/components, endpoint names, exact request/response fields, database design, package boundaries, framework mechanisms, low-level security implementation, or concrete visual-system rules.

Those concerns are derived through design, delivery, architecture, contract, and implementation authorities.

## Product Authority and Product Design / Brand Authority

Canonical design sources under `docs/ui-ux/` remain active authorities for product experience and visual expression.

Product/business truth and Product Design / Brand Authority are peer upstream inputs to delivery:

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

Design must not invent business truth. Product Authority must not discard durable product insight merely because it was discovered during design exploration.

## Relationship to existing material

During the migration:

- project/business/phase docs → product-source evidence;
- `docs/ui-ux/` → active Product Design / Brand Authority plus explicit product discoveries;
- `docs/spec/<domain>/...` → delivery/domain reference pending slice-by-slice reconciliation;
- `api/openapi/` → shared-contract reference pending slice-by-slice reconciliation;
- ERD/data-model docs → technical reference;
- backend/frontend code, tests, migrations, and Harscode artifacts → implementation/workflow evidence.

Reference does not mean incorrect. It means the artifact does not automatically define higher-level product truth merely because it is more detailed.

## Whole-product authority candidate

`product-overview.md` is the single candidate owner for durable whole-product truth in this reframe.

It contains the product thesis, character, actors, capability map, connected journey, trust/accountability model, progressive-delivery posture, and revalidation register.

Do not split those concerns into parallel product-authority files without a genuinely distinct ownership need.

## Approved MVP release scope

`mvp-scope.md` is the **approved time-bounded release scope** for the first MVP iteration.

Relationship:

```text
product-overview.md
→ what Kencleng is / durable whole-product truth

mvp-scope.md
→ what subset we deliberately prove first
```

Human approval: **2026-09-17**.

The approved MVP prioritizes one complete public trust loop:

```text
truthful campaign understanding
→ guest donation
→ truthful donation state
→ campaign closure
→ persistent result/accountability follow-up
```

Account and operational breadth enter MVP only when required to make that loop real, safe, and coherent. Security remains a non-negotiable floor without turning every future security feature into an MVP requirement.

## MVP delivery sequencing

`mvp-delivery-slices.md` is the current **candidate sequencing artifact** derived from the approved MVP scope.

It proposes:

```text
Slice 1 — Public Campaign Understanding
→ Slice 2 — Guest Donation + Truthful Donation State
→ Slice 3 — Campaign Closure + Persistent Public Result
→ Slice 4 — Accountability Follow-up
```

This sequencing is vertical and product-loop-driven, not domain-order-driven.

Account is outside the baseline MVP critical path unless real slice exploration proves an enabling dependency.

## Validation evidence

`probes/` contains validation evidence rather than authority.

- Probe 01 — Public Campaign Detail: **PASS**; demonstrated forward derivation and narrow contract reconciliation.
- Probe 02 — Account Registration + Email Verification: **PAUSED / REFRAMED**; useful backward-reconciliation evidence, but historical Account breadth is no longer allowed to drive MVP scope.

Probe 02's security/correctness findings remain salvage evidence for a future Account-dependent slice.

## Promotion discipline

The Product Authority candidate should be promoted only after:

1. human review confirms the whole-product model;
2. forward validation demonstrates Product + Design can derive coherent delivery needs;
3. backward reconciliation demonstrates detailed specs/code can be treated as evidence rather than blindly preserved/discarded;
4. product/design contradictions are explicitly resolved;
5. approved MVP scope prevents historical domain breadth from driving delivery;
6. repository root/scoped routing is deliberately updated so no long-lived dual authority remains.

After promotion, real MVP slices should resume CRTV against Harscode `main` through the canonical Exploration kickoff without custom solution-steering prompts.