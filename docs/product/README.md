# Kencleng — Product Authority

> Status: **Canonical Product Authority — promoted 2026-09-17**
> Purpose: Own Kencleng product/business truth independently of frontend/backend decomposition while preserving canonical Product Design / Brand Authority as a peer upstream input.

This directory is the canonical product/business authority for Kencleng.

The promotion closes the Product Authority reframe recorded in:

`docs/project/kencleng-product-authority-reframe-audit.md`

Historical domain specs, OpenAPI, ERD/data-model documents, and implementation remain valuable evidence, but they no longer override higher-level Product Authority merely because they are more detailed.

## Authority boundary

Product Authority owns durable product/business truth such as:

- product purpose and boundaries;
- product thesis and character;
- actors/personas;
- major capabilities and end-to-end journeys;
- core business concepts and relationships;
- major lifecycle semantics;
- durable permissions/invariants at product level;
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

## Canonical whole-product owner

`product-overview.md` owns durable whole-product truth.

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

## Approved MVP delivery sequencing

`mvp-delivery-slices.md` is the **approved delivery sequencing authority for the first MVP iteration**.

Approved order:

```text
Slice 1 — Public Campaign Understanding
→ Slice 2 — Guest Donation + Truthful Donation State
→ Slice 3 — Campaign Closure + Persistent Public Result
→ Slice 4 — Accountability Follow-up
```

This sequencing is vertical and product-loop-driven, not historical domain-order-driven.

Account is outside the baseline MVP critical path unless real slice exploration proves an enabling dependency.

## Relationship to existing material

After promotion:

- `product-overview.md` → canonical durable product/business truth;
- `mvp-scope.md` → canonical current MVP release scope;
- `mvp-delivery-slices.md` → canonical current MVP sequencing;
- `docs/ui-ux/` → active Product Design / Brand Authority;
- `docs/spec/<domain>/...` → delivery/domain specifications that are authoritative for reconciled slice detail, but must not silently override Product/MVP authority;
- `api/openapi/` → shared API contract authority for reconciled contract shape, derived from active slice needs;
- ERD/data-model docs → technical/data-model reference unless explicitly reconciled into an active delivery plan;
- backend/frontend code, tests, migrations, and Harscode artifacts → implementation/workflow evidence.

A historical artifact may remain correct. The rule is simply that detail is not promoted upward by accident.

When an active slice exposes a contradiction:

```text
Product/MVP truth clear
→ adapt delivery spec/contract/implementation

Product truth genuinely missing
→ resolve Product Authority first

Design meaning genuinely missing
→ resolve Product Design / Brand Authority

Delivery/security detail missing
→ resolve in the owning delivery/architecture/contract authority
```

Do not preserve a stale detailed decision merely because implementation already exists.

## Validation evidence

`probes/` contains validation evidence rather than authority.

- Probe 01 — Public Campaign Detail: **PASS**; demonstrated forward derivation and narrow contract reconciliation.
- Probe 02 — Account Registration + Email Verification: **PAUSED / REFRAMED**; demonstrated that existing code/specs can be treated as evidence and that security/correctness work can be salvaged without allowing historical Account breadth to drive MVP scope.

Probe 02 does not block MVP delivery. Its findings should be revisited when an Account-dependent capability becomes real scope.

## Promotion checkpoint

Promotion criteria were satisfied on **2026-09-17**:

1. the whole-product model received human review;
2. forward validation demonstrated Product + Design can derive coherent delivery needs;
3. backward reconciliation demonstrated detailed specs/code can be treated as evidence rather than blindly preserved/discarded;
4. material product/design contradictions were surfaced rather than hidden downstream;
5. MVP scope and sequencing were explicitly approved so historical domain breadth no longer drives delivery;
6. root, backend, frontend, delivery-spec, and orchestration routing were updated in the same promotion pass to remove long-lived dual authority.

After this promotion, real MVP slices resume CRTV against Harscode `main` through the canonical Exploration kickoff without custom solution-steering prompts.