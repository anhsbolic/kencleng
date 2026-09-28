# Stage 2 — O7 Product Design Surface Gap Analysis

> Phase/Stage: Explorer Open-Item Resolution / Stage 2 (gap analysis)  
> Work Unit / Run: `WU-S2-002` / `OIR-S2-002-004`  
> Author / Participant: `P-S2-002-OIR-004-1` / `KC-EXPLORER`  
> Created: 2026-09-28  
> Model / reasoning: invocation-selected `gpt-6-luna` / `high`; active runtime model is not independently exposed  
> Session: not exposed  
> Target revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535`  
> Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`

This is Stage 2 evidence, not a copy proposal or decision. Product meaning comes from current Product/MVP authority and the Human directions cited in the invocation. Historical Donation specs, bundled/generated API declarations, and code do not replace that authority.

## Design authority applied across the review

- `docs/ui-ux/README.md` routes stable experience judgment to `product-design-principles.md`, recurring interactions to `patterns.md`, and concrete provenance/status treatment to `design-guidelines.md`.
- Principles §§2, 4–6, 8–9 require no unsupported implication, visible uncertainty/consequence, comprehensible next action, and consequence explanation before commitment. §§14–15 distinguish established from materially open design and require owner review for material ambiguity.
- `patterns.md` §§4, 7, 12–14 requires clear form/submission behavior, status → meaning/consequence → relevant action, safe generic failure for sensitive unauthenticated lookup, appropriate retry, and terminal state explanation.
- `design-guidelines.md` §13 calls for visible plain-language pending/unavailable information, actual semantic treatment for system state, and no false independent-verification implication. §26 leaves final production terminology for provenance/truth labels open.
- These sources provide reusable semantics and hierarchy, but no Donation-specific success/failed label set, simulator-source phrase, email disclosure sentence, terminal email copy, or Donation status composition. Current Donation design readiness is therefore not demonstrated by an existing screen.

## Area 1 — `pending` / `success` / `failed` and simulation source labels

**Current state.** Slice 2 Product source says the persisted result belongs to the backend simulator; the approved `pending` copy is “Menunggu hasil simulasi” with no ETA or real-payment instruction. `success` and `failed` are simulator outcomes; demo failure is not donor- or browser-selected. Terminal email must identify the result as simulation, not provider settlement. No Donation/status surface is present in `frontend/app/`; the only current Campaign detail route says the donation flow is unavailable (`frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx`, `CampaignSuccess`, lines 157–161). Old Donation declarations remain in `frontend/lib/api/generated/openapi.ts`, but Product/MVP and Techplan identify historical declarations/specs as evidence only, not current runtime/UI.

**Requirement.** `docs/product/mvp-delivery-slices.md` §5 lines 169–175 and `docs/product/mvp-scope.md` Stage B lines 107–114 govern simulator truth and pending wording. Design requires semantic labels plus context/consequence (`patterns.md` §14), actual system-state treatment and explicit pending language (`design-guidelines.md` §13), and no visual or verbal claim beyond Product truth (`product-design-principles.md` §2).

**Gap.** The Product meaning is established, but live terminal-state labels/source wording and visual/context treatment do not exist. The Design authority leaves final provenance terminology open; it does not establish whether a concise terminal label visibly identifies the result as simulated without being confused with real settlement. No current surface can show whether the state is legible without color or a badge alone.

**Five sniffing lenses.**

- **Risk:** A success treatment or “berhasil” without simulation context could imply provider settlement or actual money movement.
- **Edge cases:** `pending` can persist without a timing promise; terminal state may be success or demo-only failure. A label must remain truthful for both and must not turn the pending state into a timed countdown.
- **Miscontext:** Generated OpenAPI currently declares old Donation operations and `pending`/`success`/`failed`, which can look like an active flow despite no current Donation route/UI and superseded historical semantics.
- **Misleading signals:** A green check, success animation, QR image, or payment confirmation cue could look like a real provider receipt even when accompanied by small sandbox text; none was observed in current UI because the flow is absent.
- **Inconsistency:** Product fixes pending wording, while final provenance terminology is explicitly open in Design guidelines; existing organizer-source labels do not decide Donation simulator-source terminology.

**Anchors.** `docs/product/mvp-delivery-slices.md` §5; `docs/product/mvp-scope.md` Stage B; `docs/ui-ux/patterns.md` §§7, 13–14; `docs/ui-ux/design-guidelines.md` §13 and §26; `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` `CampaignSuccess`; `frontend/lib/api/generated/openapi.ts` Donation operations (historical/generated evidence only).

## Area 2 — QRIS active vs GoPay, ShopeePay, and bank transfer unavailable

**Current state.** Product authorizes displaying these four familiar method names, but only QRIS runs the sandbox simulation. Other methods must be visibly unavailable and non-interactive; no real rails, usable real payment instructions, or implication of external settlement. Current Campaign detail has no method chooser and no donation form. Fixture marks `donation_action` unavailable with `donation_flow_not_available` (`frontend/mocks/fixtures/public-campaign.ts`, lines 63–66).

**Requirement.** `docs/product/mvp-delivery-slices.md` §5 lines 168–179 and §5 “Explicitly out of Slice 2” lines 208–215 fix the method meaning. `product-design-principles.md` §§2 and 8 require truthful visual implications and action hierarchy; `design-guidelines.md` §13 distinguishes pending/unavailable from warning/error.

**Gap.** Product resolves active/inactive behavior, but there is no rendered list, unavailable-state wording, or keyboard/accessibility behavior to inspect. Existing design authority states the truthfulness rule but does not prescribe the local control/label presentation for this method chooser. No evidence establishes whether a future UI would make disabled options visibly unavailable without presenting them as selectable or integrated.

**Five sniffing lenses.**

- **Risk:** An unavailable method styled as an actionable payment option can imply an integration or invite a real transfer unsupported by the sandbox.
- **Edge cases:** QRIS is the only active option; inactive methods must remain non-interactive across pointer, keyboard, and responsive presentations. No actual component exists to verify these states.
- **Miscontext:** Historical generated PaymentMethod enums include multiple rails; these do not override current Product’s QRIS-only simulation.
- **Misleading signals:** Familiar provider marks, QR codes, payment instructions, or enabled-looking controls could signal real payment capability; no such current UI was found.
- **Inconsistency:** Product is precise about active versus unavailable; Design has no Donation-specific component precedent, so the exact unavailable-state expression remains unobserved.

**Anchors.** `docs/product/mvp-delivery-slices.md` §5, esp. lines 168–179 and 208–215; `docs/ui-ux/product-design-principles.md` §§2, 8; `docs/ui-ux/design-guidelines.md` §13; `frontend/mocks/fixtures/public-campaign.ts` `baseCampaign.donation_action`.

## Area 3 — Optional email opt-in and verification disclosure

**Current state.** Email is optional and opt-in only for donation-status email. The 2026-09-28 Event records the Product Design owner’s explicit presentation direction: disclose near email opt-in that if the guest does not verify within 24 hours from email capture, the address is deleted and no status email is sent; Donation processing continues independently. This does not move the clock to verification-email send time or resolve the separate verified-email-while-`pending` retention conflict. No guest donation form or email opt-in exists in current frontend.

**Requirement.** `docs/product/mvp-delivery-slices.md` §5 lines 172–179 and `docs/product/mvp-scope.md` Stage B lines 114–116 govern optionality, verification, deletion, and flow independence. The cited 2026-09-28 Event carries the explicit near-opt-in disclosure decision. `product-design-principles.md` §§4, 9 and `patterns.md` §4 require material consequences to be visible at the relevant decision/form field.

**Gap.** Disclosure content, placement near opt-in, 24-hour clock origin, deletion/no-notice consequence, and processing independence are settled. No live field exists to determine whether that information is near the opt-in in the actual flow or whether the optional channel remains clearly optional. Exact concise production wording is not recorded; the present authority supports composition principles but does not show a rendered fit. The independent pending-retention issue is not a Design gap and remains excluded.

**Five sniffing lenses.**

- **Risk:** Hiding the 24-hour consequence only in Terms & Conditions would fail the explicit owner direction; a nearby disclosure that makes email appear mandatory would contradict guest flow optionality.
- **Edge cases:** Email omitted; email entered without opt-in; opt-in followed by no verification before 24 hours from capture; terminal notification held for an unverified address. The disclosure must not imply that donation processing stops if email verification does not occur.
- **Miscontext:** Treating this as Account email verification would add the wrong account requirement or imply account creation.
- **Misleading signals:** A checked-by-default box or compulsory-looking email field could contradict the opt-in/optional semantics; no current form exists to inspect.
- **Inconsistency:** Product/event fix privacy consequence and clock origin, while no exact short sentence or UI arrangement is yet evidenced. The separate verified-email pending cap remains unresolved in O2/O3 and must not be answered by design copy.

**Anchors.** `docs/product/mvp-delivery-slices.md` §5 lines 172–177; `docs/product/mvp-scope.md` Stage B lines 114–116; `.harscode-spaces/s2-guest-donation-truthful-state/events.md` “Guest email verification disclosure confirmed for current Slice 2”; `docs/ui-ux/product-design-principles.md` §§4, 9; `docs/ui-ux/patterns.md` §4.

## Area 4 — Terminal-only email notice and source wording

**Current state.** Product allows at most one status-only message to a verified opted-in address after terminal `success` or `failed`, never for initial `pending`; it must identify the result as simulation, not provider settlement. Unverified addresses do not receive the status and are deleted after the approved verification window. There is no Donation email template, sender/runtime, or current UI/preview in this frontend tree. Window/retention controls have their own Security/PII authority; verified-email pending retention remains separately open.

**Requirement.** `docs/product/mvp-delivery-slices.md` §5 lines 172–177; `docs/product/mvp-scope.md` Stage B and C lines 114–126; `docs/ui-ux/patterns.md` §§13–14 (terminal consequence, calm status communication); `docs/ui-ux/product-design-principles.md` §§5, 11 (truth class and calmness around identity/security).

**Gap.** Delivery boundaries are explicit; exact email subject/body terminology and the degree of visible simulator provenance are not specified by current Design artifacts. The status/notification source is not an organizer report or independent verification; current guidance does not provide a Donation-specific source label or message precedent. No rendered message exists for assessment.

**Five sniffing lenses.**

- **Risk:** Email that says only “donation successful” can be read as payment-provider settlement; sending while pending or to unverified email breaches current product/privacy meaning.
- **Edge cases:** Failure terminal notice; terminal result before verification; no email opt-in; repeated delivery attempts for the single logical notice. Copy must not promise pending updates or multiple notices.
- **Miscontext:** An organizer/report provenance label would imply the wrong source; the outcome is a backend simulator state.
- **Misleading signals:** Provider logos, a receipt/invoice presentation, or “pembayaran diterima” would imply a real external rail absent from the product.
- **Inconsistency:** Product says simulation, Design requires source/truth distinction but leaves final production provenance terminology open; no approved Donation-specific wording closes the gap.

**Anchors.** `docs/product/mvp-delivery-slices.md` §5 lines 172–177; `docs/ui-ux/patterns.md` §§13–14; `docs/ui-ux/product-design-principles.md` §§5, 11; `docs/ui-ux/design-guidelines.md` §13 and §26. Runtime/email implementation anchors are absent at this revision.

## Area 5 — Generic unavailable behavior for missing/invalid/expired status links

**Current state.** Product requires one generic public behavior/copy for absent, invalid, or expired guest status links; its example is “Link status tidak tersedia atau mungkin kedaluwarsa.” OIR-S2-002-002 records owner direction for a generic `404`, but transport parity, body/headers/cache/timing and residual risk are Security/API matters and are not resolved by shared UI copy. There is no Donation status route/page in current frontend.

**Requirement.** `docs/product/mvp-delivery-slices.md` §5 line 174; `docs/product/mvp-scope.md` Stage B lines 116–126; `docs/ui-ux/patterns.md` §§7 and 12 require visible failure without unsafe distinctions and useful recovery only where appropriate.

**Gap.** The user-facing semantic direction is established but no page exists to verify that it collapses all three cases or avoids revealing Donation existence. Exact recovery path and interaction are not shown. UI evidence cannot establish the unresolved transport, cache, timing, or residual-risk behavior.

**Five sniffing lenses.**

- **Risk:** Distinguishing “expired” from “not found” in page copy may reveal whether a Donation or valid credential exists; technical parity is still unverified.
- **Edge cases:** Missing token, incorrect token, nonexistent record, and expired token must not produce visibly distinguishable UI outcomes; valid but stale client state is also unobserved.
- **Miscontext:** A shared message alone does not prove identical API response status/body/headers/cache/timing; OIR-002 explicitly leaves those API/Security controls open.
- **Misleading signals:** Old generated API declarations list credential failures but do not prove parity in a current runtime or route.
- **Inconsistency:** Product/UI genericity is set; API/Security response and residual-risk evidence is still open. These are separate concerns, not a reason to infer closure from the copy.

**Anchors.** `docs/product/mvp-delivery-slices.md` §5 line 174; `docs/product/mvp-scope.md` Stage B; `docs/ui-ux/patterns.md` §§7, 12; `OIR-S2-002-002/resolution-brief.md` O4/O5; no current Donation route/page in `frontend/app/`.

## Area 6 — Explicit failed-state next action

**Current state.** Product allows the donor to explicitly start a new donation after `failed`; `pending` never auto-resubmits. The failed result is produced only through a clearly labeled demo scenario. `patterns.md` §7 says status should lead through state and consequence to relevant next action; §12 permits retry only when retry could reasonably succeed. Current frontend has no failed Donation state or recovery control.

**Requirement.** `docs/product/mvp-delivery-slices.md` §5 line 170 and `docs/product/mvp-scope.md` Stage B lines 107–112; `docs/ui-ux/patterns.md` §§7, 12–13; `product-design-principles.md` §§8–9.

**Gap.** Product distinguishes a deliberate new donation after failure from automatic resubmission while pending, but no UI explains that distinction or exposes the failed-state action. Current Design authority supplies the status/consequence/action sequence but no Donation-specific composition precedent. O9 idempotency/new-key rules are delivery behavior; they must not be silently converted into a UI-created new attempt.

**Five sniffing lenses.**

- **Risk:** A generic “Coba lagi” may be understood as retrying the failed transaction or generating an unintended second donation rather than starting a new one deliberately.
- **Edge cases:** Repeat action after terminal failure versus refresh/revisit while pending; a pending state must not show an action that automatically resubmits.
- **Miscontext:** The approved recovery is a new donation after a terminal failure, not retry of simulator/payment settlement; O9 ambiguous-request retry remains same-key/same-donation.
- **Misleading signals:** An automatic retry spinner/button on `failed` or a new donation flow triggered by page refresh would misstate user intent; none is present currently.
- **Inconsistency:** Product action meaning and generic pattern are compatible; exact action label, placement, and any explanatory sentence are absent from current artifacts/UI.

**Anchors.** `docs/product/mvp-delivery-slices.md` §5 line 170 and duplicate-submission lines 171; `docs/product/mvp-scope.md` Stage B; `docs/ui-ux/patterns.md` §§7, 12–13; `docs/ui-ux/product-design-principles.md` §§8–9.

## Cross-area current-state check

`frontend/app/` currently contains the public home and `/campaigns/[campaignId]` detail route only; there is no guest donation form, Donation status route, status lookup UI, or email preview. Campaign detail states “Alur donasi untuk campaign ini belum tersedia” and the mock fixture sets `donation_action.availability` to `unavailable`. Thus the present evidence is product/design-source analysis plus a bounded live-UI absence check, not a rendered design review. Generated Donation API types and historical Donation source materials are not treated as proof of active Slice 2 delivery.

## Stage 2 boundary

No options are compared, no copy is selected, and no Design decision is attributed to the Human in this artifact. Stage 3 needs the separate Human gate. The open verified-email-while-`pending` retention conflict remains excluded; it is not a Design decision and no email behavior that changes that promise can be derived here. No final visual acceptance, implementation, or `CONTRACT_READY` claim is made.
