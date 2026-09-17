# Probe 02 — Account Registration + Email Verification

> Status: **Probe evidence — human product decision required**
> Created: 2026-09-17
> Purpose: Test whether the candidate Product Authority can reconcile an already-implemented backend slice without allowing existing specs/code to silently define product truth.
>
> This document is validation evidence, not canonical Product Authority, Account delivery authority, or API contract.

## 1. Probe question

Starting from current product/design intent rather than the existing Account implementation, what should **registration + email verification** mean to Kencleng, and how much of the existing Account work should survive?

This probe asks:

1. what outcome registration creates;
2. what email verification proves;
3. which capabilities actually require that proof;
4. whether the existing API/data model expresses those meanings cleanly;
5. which existing implementation/security work should be kept, adapted, replaced, or deferred;
6. whether the backend needs any soft reset.

## 2. Upstream product derivation

Relevant candidate Product Authority (`docs/product/product-overview.md`) says:

- a person may understand Kencleng publicly and create/authenticate an account when useful;
- guest donation remains supported, so account creation is not universally required to contribute;
- registered users gain personal/account capabilities;
- identity/account access is a product capability, not itself a trust badge;
- Kencleng must keep verification concepts narrow and truthful;
- exact provider, password, token, MFA, persistence, and endpoint mechanics remain downstream delivery detail until revalidated.

Canonical page-map intent describes Authentication as a form surface for login/register/recovery, while public browsing and guest donation remain available independently when product rules allow.

### Derived product outcome

Registration should establish a Kencleng account and a usable authentication path without implying that the person, their organization, or any future campaign is broadly "verified".

Email verification should prove one narrow fact:

> **Kencleng has evidence that the user controls the specific email address that was verified.**

It does not prove:

- legal identity;
- organization legitimacy;
- campaign legitimacy;
- donor trustworthiness;
- ownership/control of some *other* email address;
- real-world outcome.

## 3. Recommended capability model

Email verification should be **capability-scoped**, not a global account-disabled flag.

Conceptually:

```text
account exists
→ user may authenticate
→ ordinary account/public capabilities may be available

verified control of email X
→ capabilities that rely on control of email X may become available
```

This means an unverified account does not automatically need to be unable to log in.

Examples where verified email control is materially relevant:

- claiming a prior guest donation whose guest email matches the verified account email;
- establishing/using an email-password identity after proving control of that email;
- sensitive organization-participation flows when the product explicitly requires a verified contact address;
- notification/contact semantics where delivery to a verified address is consequential.

Examples where verified email should **not** be inferred as universally required without a separate product rule:

- browsing public campaigns;
- guest donation;
- donation while authenticated merely because the donor has an account;
- generic account login.

## 4. Existing decisions that align well

### 4.1 Login before email verification — KEEP

The existing Login spec explicitly allows an `email_password` user whose identity is not yet verified to log in, with verification restrictions applied at the capability that actually needs them.

This is directionally aligned with the product derivation above and should survive.

### 4.2 Generic registration response / anti-enumeration — KEEP

`POST /auth/register` currently returns the same generic `202` response across:

- new email-password registration;
- existing unverified email-password identity;
- existing verified email-password identity;
- Google-only conflict.

The backend also attempts to normalize branch work to reduce timing-based account enumeration.

The exact implementation can be reviewed independently, but the security intent is strong and should not be discarded by a product reframe.

### 4.3 Single-use/revocable verification tokens — KEEP

The implementation uses hashed, time-bound tokens with guarded single-use redemption and revocation of superseded verification tokens.

This is security/correctness evidence worth preserving even if the surrounding product model changes.

### 4.4 Transactional/concurrency guards — KEEP

Existing work includes unique-index protection for concurrent duplicate registration and atomic guarded token redemption.

These are implementation-quality assets, not assumptions that need to be reset merely because authority hierarchy changes.

### 4.5 PII protection pattern — KEEP as security architecture evidence

Email/identifier persistence currently separates encrypted values from lookup hashes and avoids logging raw PII/tokens in the main registration/verification path.

The exact crypto implementation remains security authority, but there is no product-driven reason to discard the pattern.

## 5. Cross-domain drift found by the probe

The Login feature spec says an unverified user is restricted from "donating as registered" and becoming a representative.

The Donation submission spec, however, explicitly accepts authenticated donations without requiring verified email; guest donation is also a strong Product Authority candidate.

Therefore the Account statement about registered donation is stale/cross-domain drift, not a durable Account-owned rule.

Reconciliation direction:

> Email verification requirements must be owned by the capability that genuinely depends on control of an email address, rather than asserted globally from the Account domain.

For current evidence, ordinary authenticated donation should **not** be treated as requiring verified email unless the Donation slice later establishes a concrete product reason.

Guest-donation claim, by contrast, correctly depends on verified control of the matching email because email ownership is the authorization fact for the claim.

## 6. Material identity-model ambiguity

### Current model

The current backend has:

```text
User
- primary_email

AuthIdentity
- provider_type
- identifier
- verified_at
```

The API exposes:

```text
User.email_verified: boolean
```

with the current definition:

> true if **any** `email_password` AuthIdentity has `verified_at` set.

### Why this becomes ambiguous

Account Linking currently permits a Google-only user to add an `email_password` identity using an explicit new email that does **not** have to equal `User.primary_email`.

Therefore the following state is possible under the existing design:

```text
User.primary_email = old-google@example.com

AuthIdentity google
identifier = old-google@example.com

AuthIdentity email_password
identifier = new-work@example.com
verified_at = verified

User.email_verified = true
```

What exactly does `email_verified=true` mean here?

It proves control of `new-work@example.com`.

It does **not automatically prove** control of `User.primary_email = old-google@example.com` under the local email-verification mechanism.

This matters because other specs describe capabilities such as guest-donation claim in terms of the user's **verified primary email**.

The current single boolean therefore risks collapsing two different facts:

```text
"the account has at least one verified email identity"

vs

"the account's canonical/contact email is verified"
```

Those are not equivalent once multiple provider identities or email changes exist.

## 7. Recommended product direction for the ambiguity

Prefer explicit semantics over a generic account-level `email_verified` boolean.

### Recommended model

Kencleng should distinguish at least conceptually:

```text
Account
  canonical/contact email
  verification state for that canonical/contact email

Authentication identities
  email/password login identifier(s) as product allows
  Google/provider identity
  provider-specific verification/evidence
```

For v1, keep this simpler rather than more flexible:

1. a User has **one canonical account/contact email**;
2. product capabilities that depend on "the user's verified email" mean verified control of this canonical email;
3. authentication providers may prove or use identifiers, but provider identity must not silently redefine what another capability means by "verified account email";
4. if a user intentionally changes the canonical email, treat that as an explicit account-email-change/product flow with verification, rather than an incidental side effect of adding a login method;
5. avoid a public/API boolean whose meaning is "some email identity somewhere is verified" when consumers actually need "canonical email is verified".

This direction does **not** require choosing the final table/column design in Product Authority.

## 8. Consequence for existing Account Linking

The current Google-only → set-password flow allows a different email primarily to support migration from one login email to another.

That user need is legitimate, but the current mechanism mixes two concerns:

```text
add/change authentication method
+
change the account's canonical email
```

Recommended reconciliation:

- keep the user need;
- do not assume the current overloaded `set-password(email, password)` contract is the final expression;
- when the Account slice is next delivered, separate "establish/change password login" from "change canonical account email" unless exploration shows a simpler contract with equally explicit semantics.

The existing implementation around linking should therefore be treated as **ADAPT**, not automatically preserved as product truth.

## 9. Existing implementation classification

### KEEP

Preserve unless later technical/security evidence contradicts them:

- generic registration response / anti-enumeration intent;
- duplicate-registration concurrency protection;
- password hashing and password-secret handling pattern;
- PII encryption/HMAC-at-rest pattern;
- verification-token hashing;
- single-use guarded redemption;
- superseded-token revocation on verification resend;
- transactional registration creation;
- send-email-after-commit posture;
- sanitized logging posture;
- ability to authenticate before email verification;
- tests proving concurrency/security properties that remain semantically applicable.

### ADAPT

- account-level meaning/exposure of `email_verified`;
- relationship between `User.primary_email` and provider/auth-identity identifiers;
- Account Linking flow that allows a different email as part of `set-password`;
- downstream capability checks that currently refer ambiguously to "verified email";
- stale Account-domain statement that authenticated donation requires verified email;
- exact registration/verification API response copy and FE flow once Authentication UI is actually delivered.

### REVALIDATE / DEFER

Do not make these Product Authority merely because they are implemented/spec'd:

- exact 24-hour verification-token TTL;
- exact password minimum and breach-list provider/fail-open policy;
- exact provider set (`email_password`, Google, future phone OTP);
- exact endpoint naming/path split;
- exact notification/nudge email behavior;
- exact token-status HTTP distinctions where they are UX/contract detail rather than security invariants.

### REPLACE

No major registration/email-verification security mechanism currently warrants wholesale replacement based on this probe alone.

A future implementation may replace bounded code as a consequence of the identity/email semantic adaptation, but that is different from declaring the existing Account subsystem disposable.

## 10. Backend soft-reset verdict — provisional

### Whole-backend reset

**Not justified.**

The probe found substantial reusable security/correctness implementation and no evidence that the backend architecture as a whole conflicts with the new Product Authority.

### Account clean-room rewrite

**Not justified at this point.**

A clean rewrite would unnecessarily throw away race-condition handling, anti-enumeration work, token correctness, PII controls, tests, and security mechanics that remain useful.

### Scoped Account semantic reset

**Recommended.**

Interpret "soft reset" as:

```text
existing Account implementation = valuable implementation evidence

but

canonical email / login identity / verified-email semantics
= reopen and reconcile before future Account delivery
```

When Account development resumes:

1. freeze current semantic assumptions as reference;
2. derive the next Account delivery slice from Product Authority;
3. define canonical-account-email semantics explicitly;
4. map surviving implementation to the new model;
5. selectively migrate/modify code and migrations only where the new semantics require it;
6. preserve applicable security tests/mechanics throughout.

## 11. Human product decision gate

Before marking Probe #2 PASS and before promoting repository authority routing, confirm this product direction:

> **A Kencleng account has one canonical account/contact email. "Verified email" for cross-product capabilities means verified control of that canonical email, not merely that some authentication identity has a verified email identifier. Adding/changing an authentication method must not silently change the canonical account email; changing the canonical email is an explicit, verified account action.**

If approved:

- Probe #2 can be marked PASS;
- existing Account registration/verification is classified mostly KEEP with a scoped identity/email semantic ADAPT;
- whole-backend soft reset is rejected;
- a scoped Account semantic reset becomes the recommended migration posture;
- repository authority promotion can proceed before the next real development slice.

If rejected, define the alternative relationship between canonical email, provider identities, and verified-email-dependent capabilities before Account delivery resumes.
