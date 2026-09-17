# Kencleng — MVP Scope

> Status: **Candidate MVP Scope — human review required**
> Created: 2026-09-17
> Scope: Time-bounded product commitment for the first Kencleng MVP iteration.
> Upstream: `docs/product/product-overview.md` + canonical `docs/ui-ux/` Product Design / Brand Authority.
>
> This document does **not** redefine what Kencleng is. `product-overview.md` owns whole-product truth. This document only defines the smallest product loop that should be delivered first to validate Kencleng's differentiated value without prematurely implementing the whole platform.

## 1. MVP thesis

Kencleng MVP should optimize for **depth over breadth**.

The MVP is not the smallest collection of screens or backend domains. It is the **smallest complete, truthful product loop that demonstrates Kencleng's differentiated value safely**.

For Kencleng, that differentiated loop is:

```text
understand
→ build confidence from evidence
→ donate
→ understand donation state
→ fundraising closes
→ return to the same campaign
→ understand result / accountability / what remains unknown
```

The MVP succeeds when a skeptical-but-open visitor can experience that loop without Kencleng manufacturing trust, hiding uncertainty, or requiring platform breadth that the loop does not need.

## 2. MVP decision rule

A capability belongs in MVP only when at least one of these is true:

### Product-critical

Without it, the selected Kencleng trust loop cannot be understood or completed.

### Security / correctness-critical

Without it, the delivered loop would be materially unsafe, misleading, privacy-breaking, authorization-breaking, or financially incorrect.

### Enabling-critical

It is not user-facing core value by itself, but the selected loop cannot operate end-to-end without it.

Everything else should default to **defer**, even when it is valuable, technically interesting, already specified, or partially implemented.

A useful forcing question is:

> **What part of the MVP product hypothesis fails if this capability is not built now?**

If the answer is weak, the capability is probably not MVP-critical.

## 3. Primary MVP hypothesis

Kencleng can earn user confidence through structured evidence and preserve that trust after donation by making campaign facts, funding state, lifecycle, provenance, and post-campaign accountability understandable before and after the contribution.

The MVP does **not** need to prove that every future actor can self-serve every workflow.

It does need to prove that the public/donor-facing trust loop is coherent, truthful, and technically real.

## 4. Primary MVP actor

### Public Visitor / Guest Donor

This is the primary MVP actor because the core product value must be understandable before account creation.

The actor must be able to:

- understand a public campaign;
- understand who is responsible for it;
- distinguish platform/system facts from organizer-provided content and unknown/pending information;
- understand funding state without confusing funding with impact;
- donate without being forced to create an account;
- understand the donation's real sandbox processing/result state;
- return after fundraising closes and find the same public campaign identity;
- understand final funding/result context and available accountability/follow-up information.

Account creation is therefore **not a prerequisite for experiencing the MVP's core value**.

## 5. Minimum complete MVP loop

### Stage A — Public understanding

A visitor can reach a truthful Public Campaign Detail experience that communicates, at minimum:

- campaign identity and purpose;
- responsible Organization/steward context at an intentionally public-safe level;
- truthful campaign media state;
- current public lifecycle meaning;
- funding target/current funding facts;
- organizer-provided story/context with provenance that does not blur into platform fact;
- meaningful unknown/pending states;
- a calm, truthful next action.

The public response must follow the safety direction already found in Probe 01: explicit public-safe projection, not internal-record inheritance.

### Stage B — Donation

A visitor can submit a guest donation to an eligible campaign.

The MVP donation capability must preserve:

- truthful eligibility;
- money/amount integrity;
- idempotent submission where duplicate client requests could create duplicate contributions;
- clear sandbox semantics;
- no false implication that simulated processing is real external payment settlement.

The MVP does **not** need multiple payment methods merely to resemble a production fundraising platform. One deliberately supported sandbox payment path may be sufficient if it exercises the actual product loop.

### Stage C — Donation state

The guest donor can understand whether the donation is pending, successful, or failed through an appropriate safe mechanism.

The state must be real within the sandbox model, not a frontend-only success illusion.

Personal account history is not required to prove this part of the MVP if a safe guest-tracking mechanism is sufficient.

### Stage D — Campaign closure

The campaign can truthfully transition out of fundraising.

After closure:

- donation is no longer available;
- the same public campaign identity remains reachable if the campaign had entered the public lifecycle;
- final funding/result facts remain understandable;
- the surface shifts emphasis from conversion to result/accountability.

### Stage E — Accountability continuation

The public/donor can understand at least one meaningful post-campaign accountability state.

MVP accountability must demonstrate the product distinction between:

```text
funding result
≠ organizer-reported follow-up
≠ platform-known/verified fact
≠ independently verified real-world impact
```

The MVP does not need a complete future disbursement/report-verification platform to prove this principle.

It may begin with a narrow accountability model that can truthfully show:

- organizer-reported update/result when available;
- provenance/source;
- chronology/timestamp;
- pending/not-yet-reported state;
- any platform-known lifecycle fact that is genuinely known.

A manual or operator-assisted process may create this accountability content during the first MVP iteration if the public experience labels its source truthfully and does not pretend the process is automated or independently verified when it is not.

## 6. Operational breadth intentionally minimized

The first MVP does not need every future operational actor to have a complete self-service product surface.

Initial campaign/organization/accountability setup may use narrower operator-assisted or seeded workflows when all of the following remain true:

1. the public-facing state is backed by real persisted application data, not a hard-coded frontend story;
2. provenance and verification claims remain truthful;
3. no hidden manual action is represented to users as an automated or verified platform capability;
4. the shortcut does not weaken money integrity, authorization boundaries, privacy, or other security-floor requirements;
5. the shortcut is clearly an MVP delivery choice rather than an accidental long-term architecture commitment.

This allows Kencleng to validate the trust loop before building the full Organization Owner / Staff / Curator / Admin workflow breadth.

## 7. MVP security floor

Security is a floor, not a justification for implementing every security feature in the roadmap.

The MVP must preserve security appropriate to the capabilities it actually exposes.

At minimum:

- authorization is enforced server-side for every non-public action;
- public contracts expose explicit public-safe projections;
- sensitive tokens/secrets are not logged or exposed unnecessarily;
- password/credential handling is secure **if password authentication is included in the MVP**;
- session/token behavior is secure **if authenticated account capability is included in the MVP**;
- donation creation/settlement cannot be forged through an exposed transition;
- duplicate/concurrent money transitions cannot corrupt funding totals;
- sensitive personal data is protected at rest/in logs according to the active security architecture;
- public media access can actually be revoked when product lifecycle requires withdrawal;
- sandbox simulation is never presented as real external settlement or real-world evidence.

The MVP does **not** automatically require:

- MFA for every user;
- multiple authentication providers;
- provider linking/unlinking;
- backup recovery codes;
- a generalized role-administration platform;
- every future fraud/abuse defense before the corresponding capability exists.

If a privileged MVP operator surface is introduced, stronger authentication/authorization may be required for that specific role without imposing the same feature breadth on ordinary donors.

## 8. Candidate MVP capability classification

### IN — core MVP

- truthful Public Campaign Detail;
- narrow public Organization/steward projection;
- truthful campaign media or explicit missing-media state;
- guest donation;
- sandbox donation processing/result;
- safe guest donation status/tracking;
- campaign closure;
- persistent closed-campaign public identity;
- final funding/result context;
- narrow post-campaign accountability/follow-up with provenance and pending states;
- minimum operational mechanism needed to create/update the above truthfully;
- security/correctness work directly required by the exposed loop.

### MAYBE — include only when the chosen delivery path proves it is enabling-critical

- simple account registration/login;
- canonical account email + email verification;
- registered donation history;
- guest-donation claim;
- minimal Organization Owner surface;
- minimal Curator review surface;
- notifications beyond the minimum necessary guest tracking/follow-up;
- campaign discovery beyond a minimal way to reach the available public campaign(s).

These should not be promoted to MVP merely because historical specs or code already exist.

### DEFER by default

Unless new evidence makes one MVP-critical:

- Google OAuth;
- multiple identity providers;
- account-provider linking/unlinking;
- generalized account-email migration flows;
- MFA for ordinary users;
- backup codes;
- broad admin user/role management;
- full Organization self-service lifecycle;
- full representative-management workflow;
- full Campaign curation workflow UI;
- Event management/discovery;
- public donor list/social proof;
- recommendation/ranking/popularity/trending;
- advanced campaign discovery/filtering;
- comprehensive notification center;
- generalized dashboards/KPIs;
- full disbursement workflow;
- full fund-usage-report submission/review workflow;
- broad impact scoring or impact claims;
- production payment rails/banking integrations.

`DEFER` means **not needed to prove the first product hypothesis**, not "never build" and not "the existing work is wrong."

## 9. Account consequence

The Account domain must be reframed **after** MVP scope is accepted.

The MVP question is not:

> How much of the existing Account platform can we preserve?

It is:

> What account capability, if any, is required to complete the selected MVP trust loop safely and coherently?

Because guest donation is a core product capability, the first MVP may be able to prove the core loop without requiring Account in the critical path at all.

If Account is included, prefer the smallest coherent product model needed by the selected loop rather than restoring the previous provider-rich design by default.

The security/correctness evidence found in Probe 02 remains valuable even if much of its feature breadth is deferred.

## 10. Definition of MVP done

The MVP is not done merely because individual domain endpoints exist.

It is done when a representative end-to-end scenario can be exercised:

```text
visitor opens a public campaign
→ understands purpose/steward/funding/source context
→ makes a guest donation
→ sees truthful donation state
→ campaign closes according to product rules
→ original public campaign remains reachable
→ donation CTA is gone
→ final funding/result is visible
→ accountability follow-up or an explicit pending state is visible
→ provenance remains understandable throughout
```

and the experience remains correct under the security/correctness floor relevant to those exposed capabilities.

## 11. What this MVP deliberately does not prove

The first MVP does not need to prove:

- complete organizer self-service;
- complete curator/admin operations;
- production payment integration;
- mature identity-provider flexibility;
- every future account-security feature;
- every future campaign/disbursement/reporting workflow;
- product-market fit at scale.

It should prove that **Evidence-Led Optimism can exist as a coherent working product loop rather than only as design language or documentation.**

## 12. Human review gate

Confirm or revise this MVP direction before Account reframing or delivery planning resumes:

> **Kencleng MVP prioritizes one complete public trust loop: truthful campaign understanding → guest donation → truthful donation state → campaign closure → persistent result/accountability follow-up. Operational and account breadth are included only when they are required to make that loop real, safe, and coherent. Security remains a non-negotiable floor, but advanced identity/security feature breadth is deferred until a real product capability requires it.**

If approved, the next work should:

1. promote this MVP scope as the current release-scope authority (without replacing whole-product Product Authority);
2. revisit Probe 02 / Account through the MVP lens and reduce Account to the smallest needed capability;
3. derive the first MVP delivery slices from the loop rather than from domain order;
4. reconcile only the contracts required by those slices;
5. resume CRTV against Harscode using those real slices.
