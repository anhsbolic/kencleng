# Kencleng Product Authority Probes

> Status: Validation evidence — **not product authority**
> Last updated: 2026-09-17

These probes validate the candidate Product Authority model against real delivery slices.

A probe may identify product, design, contract, or implementation gaps. Its conclusions do not become canonical merely because they are written here. Durable conclusions must be promoted deliberately into the authority that owns them.

## Validation status

### Candidate Product Authority human review

**PASS — 2026-09-17**

`docs/product/product-overview.md` has been reviewed by the human product authority and is directionally accepted as the Kencleng product model for validation purposes.

It remains non-canonical until the forward/backward validation is complete enough and repository routing is deliberately promoted.

### Probe 01 — Public Campaign Detail

**PASS — forward derivation + human product decision + narrow contract reconciliation completed**

Artifacts:

- `01-public-campaign-detail.md`
- `01-public-campaign-detail-contract-reconciliation.md`

Human product decision approved 2026-09-17:

> A Campaign that has entered the public fundraising lifecycle remains publicly reachable after fundraising closes, using the same public campaign identity/URL. Donation action is removed and result/accountability context takes priority.

The approval is constrained by an explicit public-safety rule:

> Public lifecycle continuity preserves the public Campaign identity, not the internal Campaign record shape.

Key reconciliation findings:

- the existing composite campaign-detail concept is reusable but the public response must become an explicit public-safe projection rather than inheriting the full internal `Campaign` schema;
- `closed` may remain public only for Campaigns that previously entered the public lifecycle; other workflow states remain non-public;
- unauthenticated/public reads must not expose internal lifecycle existence/state for non-public Campaigns;
- public Organization context must use a narrow deliberate projection rather than blindly reusing internal status/fields;
- funding data should expose only facts needed for public understanding rather than all calculable/social-proof data;
- donation eligibility must remain backend-authoritative;
- campaign media needs its own public-safe projection;
- the historical public-bucket/direct-URL media design cannot provide true revocation after unpublish/retraction, so media-access architecture must support actual withdrawal of public access;
- Donation Flow remains a separate next-action dependency;
- the public donor list remains deferred for this slice.

The operational OpenAPI has intentionally **not** been edited yet because Product Authority is still candidate/non-canonical during the reframe. The narrow intended contract delta is now recorded and can be applied after authority promotion without rewriting the whole Campaign API.

Probe 01 therefore provides positive evidence that the new hierarchy can derive a real surface, expose contradictions in old detailed contracts, obtain a material product decision, and converge on a narrow contract direction without a full-product respec.

### Probe 02 — Account Registration + Email Verification

**PAUSED / REFRAMED — useful backward-reconciliation evidence, but Account scope must now be reconsidered through the MVP definition first**

Artifact:

- `02-account-registration-email-verification.md`

The probe remains valuable evidence. It found that:

- the core registration/email-verification security work is valuable and mostly reusable;
- allowing login before email verification is directionally correct when verification is enforced only at capabilities that genuinely depend on proof of email control;
- the existing generic `202` registration behavior, anti-enumeration intent, transactional/concurrency guards, single-use token mechanics, PII protection, and applicable tests are strong `KEEP` candidates;
- the Account Login spec contains stale cross-domain wording that implies authenticated donation requires verified email, while the Donation submission spec does not impose that requirement;
- guest-donation claim correctly depends on verified control of the matching email because email ownership is the authorization fact for that capability;
- the current `User.email_verified` / `primary_email` / auth-identity relationship is semantically ambiguous once provider identifiers can diverge;
- whole-backend reset and Account clean-room rewrite are not supported by current evidence;
- if Account continues with the previous breadth, a scoped semantic reset would be required around canonical account email, authentication identities, and verification meaning.

However, the probe also exposed a more upstream problem: previous Account planning was too technical-first and may include feature breadth that the first Kencleng MVP does not need at all.

Therefore the previous human decision gate about canonical email is **not being promoted yet**. The project intentionally moved one level up first to define the MVP product loop.

Current upstream artifact:

- `../mvp-scope.md` — **Candidate MVP Scope — human review required**

The next Account question is no longer:

> How should we repair the full existing Account model?

It is:

> What Account capability, if any, is actually required to complete the approved MVP trust loop safely and coherently?

After MVP scope is approved, Probe 02 should be revisited as a narrower backward-reconciliation exercise. Existing security/correctness work remains salvage evidence even if much of the historical Account feature breadth is deferred.
