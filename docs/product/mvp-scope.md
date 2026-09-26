# Kencleng — MVP Scope

> Status: **Approved MVP Release Scope — 2026-09-17; Slice 2 decisions amended by Human approval — 2026-09-26**
> Scope: Time-bounded product commitment for the first Kencleng MVP iteration.
> Upstream: `docs/product/product-overview.md` + canonical `docs/ui-ux/` Product Design / Brand Authority.
>
> This document does **not** redefine what Kencleng is. `product-overview.md` owns whole-product truth. This document defines the subset deliberately delivered first to validate Kencleng's differentiated value without prematurely implementing the whole platform.

## 1. MVP thesis

Kencleng MVP optimizes for **depth over breadth**.

The MVP is not the smallest collection of screens or backend domains. It is the **smallest complete, truthful product loop that demonstrates Kencleng's differentiated value safely**.

The approved loop is:

```text
understand
→ build confidence from evidence
→ donate
→ understand donation state
→ fundraising closes
→ return to the same campaign
→ understand result / accountability / what remains unknown
```

The MVP succeeds when a skeptical-but-open visitor can experience that loop without Kencleng manufacturing trust, hiding uncertainty, or requiring platform breadth the loop does not need.

## 2. MVP inclusion rule

A capability belongs in MVP only when at least one of these is true:

### Product-critical

Without it, the selected Kencleng trust loop cannot be understood or completed.

### Security / correctness-critical

Without it, the delivered loop would be materially unsafe, misleading, privacy-breaking, authorization-breaking, or financially incorrect.

### Enabling-critical

It is not user-facing core value by itself, but the selected loop cannot operate end-to-end without it.

Everything else defaults to **DEFER**, even when it is valuable, technically interesting, already specified, or partially implemented.

Forcing question:

> **What part of the MVP product hypothesis fails if this capability is not built now?**

If the answer is weak, the capability is probably not MVP-critical.

## 3. Product hypothesis being tested

Kencleng can earn user confidence through structured evidence and preserve that trust after donation by making campaign facts, funding state, lifecycle, provenance, and post-campaign accountability understandable before and after the contribution.

The MVP does **not** need to prove that every future actor can self-serve every workflow.

It does need to prove that the public/donor-facing trust loop is coherent, truthful, technically real, and safe.

## 4. Primary MVP actor

### Public Visitor / Guest Donor

This is the primary MVP actor because Kencleng's core value must be understandable before account creation.

The actor must be able to:

- understand a public campaign;
- understand who is responsible for it;
- distinguish platform/system facts from organizer-provided content and unknown/pending information;
- understand funding state without confusing funding with impact;
- donate without being forced to create an account;
- understand the donation's real sandbox processing/result state;
- return after fundraising closes and find the same public campaign identity;
- understand final funding/result context and available accountability/follow-up information.

Account creation is **not a prerequisite** for experiencing the MVP's core value.

## 5. Minimum complete MVP loop

### Stage A — Public understanding

A visitor can reach a truthful Public Campaign Detail experience containing only deliberately public-safe information and enough context to understand:

- campaign identity/purpose;
- responsible Organization/steward context;
- truthful campaign media state;
- current public lifecycle meaning;
- funding facts;
- organizer-provided story/context and its provenance;
- meaningful unknown/pending states;
- a truthful next action.

Public Campaign data must follow Probe 01's safety direction: explicit public-safe projection, not internal-record inheritance.

### Stage B — Guest donation

A visitor can submit a guest donation to an eligible Campaign.

The donation capability must preserve:

- truthful eligibility;
- IDR amount input as whole Rupiah, with minimum Rp5.000 and Rp1 increments (Rp5.001 is valid);
- exact decimal representation for stored and calculated monetary values, including derived values; tax rules and derived-value rounding remain out of scope until separately defined;
- duplicate-submission protection: an ambiguous retry reuses the same idempotency key and refers to the same donation; the same key with a different payload is rejected; the client prevents double-click submission and does not rotate the key while the result is ambiguous; a new key creates a new donation only after an intentional donor action;
- clear sandbox semantics: the backend simulator owns the `pending` → `success`/`failed` result; a failure is produced only by a clearly labeled demo scenario, never by donor selection or a browser request;
- no false implication of real external payment settlement.

The donation screen displays familiar Indonesian methods: QRIS, GoPay, ShopeePay, and bank transfer. **QRIS is the only method that runs the sandbox simulation.** The other methods are visibly unavailable and non-interactive. This Human-approved display choice takes priority over the earlier MVP wording that excluded payment-method breadth; it does not add functional payment-method breadth or real provider integration. No displayed method may move real money or provide usable real-payment instructions.

While the backend simulator is processing a donation, the status copy is “Menunggu hasil simulasi,” without a time estimate or instruction to make a real payment. If the result is `failed`, the donor may explicitly begin a new donation; `pending` must not trigger an automatic resubmission. The sandbox must not promise a real-world verification time or an SLA.

Guest name is optional and is not public by default. Email is optional and may be supplied only by donors who opt in to donation-status notifications. This is guest notification-email verification, separate from Account email verification. Verify ownership before sending a donation status or access link. Send at most one status-only email when the donation reaches `success` or `failed`, not while it is `pending`, and clearly identify the result as simulated rather than provider settlement. If the donation reaches a terminal state before email verification, hold the notification within a Security/PII-approved verification window; if that window expires, delete the unverified address without sending the status. Security/PII defines the verification and post-terminal delivery-retry windows and the corresponding retention controls. Campaign-wide update emails are not part of this capability.

Provide a temporary guest status URL with a hard 24-hour validity period, a hard-to-guess token, and access limited to the donation status. The donor does not need to revisit it after expiry. An invalid, missing, or expired status link has one generic public behavior and copy, for example: “Link status tidak tersedia atau mungkin kedaluwarsa.” Optional email/Account benefit information may appear after status access but must not block the guest flow or promise benefits that do not exist.

### Stage C — Donation state

The guest donor can safely understand whether the donation is pending, successful, or failed.

The state must be real within the sandbox model, not a frontend-only success illusion.

Account history is not required if a safe guest-tracking mechanism can complete this need.

The temporary guest status URL described in Stage B is available for 24 hours only and reveals donation status only. Without email, the donor does not need continued status access after the link expires. If the donor opts in to email status notifications, the verified address receives one terminal status notification as defined in Stage B; this remains a sandbox result, not evidence of provider settlement.

### Stage D — Campaign closure

The Campaign can truthfully transition out of fundraising.

`max_amount` is a closure threshold, not a hard cap on a donation amount or total funding. A donation submitted while the Campaign is eligible may be accepted in full even if that donation takes funding above the threshold. Donations already accepted while eligible and still pending when the Campaign closes remain eligible to settle in full; total funding may therefore exceed the threshold further. New donation submissions after the Campaign is closed are rejected.

After closure:

- donation is unavailable;
- the same public Campaign identity remains reachable when the Campaign genuinely entered the public lifecycle;
- final funding/result facts remain understandable;
- the surface shifts emphasis from donation to result/accountability.

### Stage E — Accountability continuation

The visitor/donor can understand at least one meaningful post-campaign accountability state.

The MVP must demonstrate that:

```text
funding result
≠ organizer-reported follow-up
≠ platform-known/verified fact
≠ independently verified real-world impact
```

A narrow model is sufficient if it can truthfully show organizer-reported information when available, provenance/source, chronology/timestamp, pending/not-yet-reported state, and genuinely known platform lifecycle facts.

## 6. Operational breadth is intentionally minimized

The first MVP does not require every future operational actor to have a complete self-service surface.

Initial Organization, Campaign, lifecycle, or accountability setup may be seeded or operator-assisted when all of the following hold:

1. user-facing state is backed by real persisted application data, not hard-coded frontend storytelling;
2. provenance and verification claims remain truthful;
3. manual actions are not represented as automated/reviewed/verified platform capability when they are not;
4. the shortcut does not weaken money integrity, authorization, privacy, or other security-floor requirements;
5. the shortcut is explicitly an MVP delivery choice, not accidental long-term architecture.

This permits validation of the trust loop before building full Owner / Staff / Curator / Admin workflow breadth.

## 7. MVP security floor

Security is a floor, not a reason to implement every security feature in the roadmap.

The MVP must preserve security appropriate to capabilities it actually exposes. At minimum:

- server-side authorization for every non-public action;
- explicit public-safe API projections;
- sensitive tokens/secrets are not logged or unnecessarily exposed;
- secure password/credential handling **if** password authentication enters MVP scope;
- secure session/token behavior **if** authenticated Account capability enters MVP scope;
- donation creation/settlement cannot be forged through an exposed transition;
- duplicate/concurrent money transitions cannot corrupt funding totals;
- sensitive personal data follows active protection/logging rules;
- public media access can actually be revoked when lifecycle requires withdrawal;
- sandbox simulation is never presented as real external settlement or real-world evidence.

The MVP does **not** automatically require MFA, multiple authentication providers, provider linking/unlinking, backup codes, generalized role administration, or defenses for capabilities that do not yet exist.

## 8. MVP capability classification

### IN — baseline MVP

- truthful Public Campaign Detail;
- narrow public Organization/steward projection;
- truthful campaign media or explicit missing-media state;
- guest donation;
- sandbox donation processing/result;
- safe guest donation status/tracking;
- the approved QRIS-only simulated processing path, with other familiar Indonesian methods shown as unavailable;
- optional guest name and verified, opt-in terminal donation-status email as defined in Stage B;
- Campaign closure;
- persistent closed-Campaign public identity;
- final funding/result context;
- narrow post-campaign accountability/follow-up with provenance and pending states;
- minimum operational mechanism required to create/update those states truthfully;
- security/correctness work directly required by exposed capabilities.

### MAYBE — only if delivery evidence makes it enabling-critical

- simple Account registration/login;
- canonical Account email + email verification;
- registered donation history;
- guest-donation claim;
- minimal Organization Owner surface;
- minimal Curator review surface;
- notifications beyond the approved minimum guest donation-status email and guest tracking/follow-up;
- Campaign Discovery beyond a minimal way to reach available public Campaigns.

These do not become MVP requirements merely because historical specs/code exist.

### DEFER by default

Unless new evidence makes one MVP-critical:

- Google OAuth and multiple identity providers;
- provider linking/unlinking;
- generalized account-email migration flows;
- MFA for ordinary users and backup codes;
- broad Admin user/role management;
- full Organization self-service and representative management;
- full Campaign curation workflow UI;
- Event management/discovery;
- public donor list/social proof;
- ranking/popularity/trending and advanced discovery/filtering;
- comprehensive notification center and generalized dashboards/KPIs;
- full disbursement workflow;
- full fund-usage report submission/review workflow;
- broad impact scoring/impact claims;
- production payment rails/banking integrations.

`DEFER` means **not needed to prove the first product hypothesis**, not "never build" and not "existing work is wrong."

## 9. Account consequence

Account is **not automatically on the MVP critical path**.

The relevant question is:

> What Account capability, if any, is required to complete the approved trust loop safely and coherently?

Because Guest Donation is a core capability, the baseline MVP can be delivered without requiring Account first unless later slice exploration finds a real enabling dependency.

Existing Account security/correctness work discovered in Probe 02 remains valuable salvage evidence. Historical feature breadth does not become scope merely because it exists.

When an Account-dependent product capability enters real scope, Account Product Reframing should derive the smallest coherent identity/authentication model from that need.

## 10. Definition of MVP done

The MVP is done when this representative scenario is real end-to-end:

```text
visitor opens a public campaign
→ understands purpose/steward/funding/source context
→ makes a guest donation
→ sees truthful donation state
→ campaign closes according to supported product rules
→ original public campaign remains reachable
→ donation action is gone
→ final funding/result is visible
→ accountability follow-up or explicit pending state is visible
→ provenance remains understandable throughout
```

and the experience satisfies the security/correctness floor applicable to those exposed capabilities.

## 11. What MVP deliberately does not prove

The first MVP does not need to prove:

- complete organizer self-service;
- complete curator/admin operations;
- production payment integration;
- mature identity-provider flexibility;
- every future Account-security feature;
- every future Campaign/disbursement/reporting workflow;
- product-market fit at scale.

It should prove that **Evidence-Led Optimism can exist as a coherent working product loop rather than only as design language or documentation.**

## 12. Approval record and next derivation

Human product review: **APPROVED — 2026-09-17**. Human follow-up approval for the Slice 2 product decisions in this document: **APPROVED — 2026-09-26**.

Approved direction:

> **Kencleng MVP prioritizes one complete public trust loop: truthful campaign understanding → guest donation → truthful donation state → campaign closure → persistent result/accountability follow-up. Operational and Account breadth are included only when required to make that loop real, safe, and coherent. Security remains a non-negotiable floor, while advanced identity/security feature breadth is deferred until a real product capability requires it.**

The current delivery sequencing derived from this scope is recorded in:

- `docs/product/mvp-delivery-slices.md`

Delivery planning must proceed from that product loop rather than historical domain order.

The 2026-09-26 Human-approved Slice 2 decisions take priority over conflicting pre-amendment MVP wording. Product and delivery documents must reflect those decisions; technical, contract, Design, and Security/PII details remain subject to their owning reviews.
