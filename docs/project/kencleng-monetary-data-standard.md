# Kencleng — Project-wide Monetary Data Standard

> Status: Direction approved by the current named owner of this authority area.
> Effective: 2026-09-30.
> Owner: Anhar Solehudin; project-wide and reassignable.

This document owns the shared technical representation direction for monetary
values across Kencleng features, APIs, domains, and persistence. It does not
own feature business policy such as per-feature amount limits, tax, pricing,
or foreign-exchange policy unless separately attributed.

## Approved representation direction

- API and wire monetary values use a major-unit decimal string together with
  an explicit currency code.
- Monetary calculation and persistence preserve exact decimal values
  end-to-end. Do not use binary floating-point (`float`/`float64`) for
  monetary values.
- The active currency set follows current Product Authority. This direction
  does not add a currency.
- For current Slice 2, input remains whole Rupiah (IDR) under the approved
  Product rule.
- Campaign's existing `NUMERIC(19,2)` column is implementation precedent for
  that domain, not a project-wide storage standard.

## Deliberately unresolved

Do not establish a universal database precision or scale now. Concrete
precision, range, per-currency fraction rules, and persistence scale must be
decided when the supported-currency set and actual computation requirements
are known, while preserving every valid exact value. This direction does not
decide additional currencies, numeric range, per-currency fractional
precision, database scale, or migration details.

## Application

Product Authority continues to own the active currency set and feature-level
amount rules. The current Slice-2 whole-Rupiah input rule remains intact.
Active API contracts define the concrete field schemas after reconciliation
against this standard. Domain and storage implementations must preserve exact
decimal values; a concrete database type/scale is selected with the supported
currency and computation needs in view rather than copied from Campaign by
default.

This project-wide direction was explicitly approved by the named authority
owner on 2026-09-30. Decision provenance and the Slice-2 reconciliation are
recorded in `.harscode-spaces/s2-guest-donation-truthful-state/events.md`.
