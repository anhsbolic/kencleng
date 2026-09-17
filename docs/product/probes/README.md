# Kencleng Product Authority Probes

> Status: Validation evidence — **not product authority**
> Last updated: 2026-09-17

These probes validated the Product Authority model against real delivery slices during the authority reframe.

A probe may identify product, design, contract, or implementation gaps. Its conclusions do not become canonical merely because they are written here. Durable conclusions are promoted deliberately into the authority that owns them.

## Validation status

### Product Authority human review

**PASS — 2026-09-17**

`docs/product/product-overview.md` was reviewed by the human product authority and is now promoted as canonical Product Authority.

### MVP scope human review

**PASS — 2026-09-17**

`docs/product/mvp-scope.md` is approved as the current time-bounded release scope for the first MVP iteration.

Approved loop:

```text
truthful campaign understanding
→ guest donation
→ truthful donation state
→ campaign closure
→ persistent result/accountability follow-up
```

Historical domain breadth must not override this release scope merely because detailed specs/code already exist.

### MVP delivery sequencing human review

**PASS — 2026-09-17**

`docs/product/mvp-delivery-slices.md` is approved as the current MVP delivery order:

```text
Slice 1 Public Campaign Understanding
→ Slice 2 Guest Donation + Truthful Donation State
→ Slice 3 Campaign Closure + Persistent Public Result
→ Slice 4 Accountability Follow-up
```

Account is outside the baseline MVP critical path until a real scoped capability proves otherwise.

## Probe 01 — Public Campaign Detail

**PASS — forward derivation + human product decision + narrow contract reconciliation completed**

Artifacts:

- `01-public-campaign-detail.md`
- `01-public-campaign-detail-contract-reconciliation.md`

Human product decision approved 2026-09-17:

> A Campaign that has entered the public fundraising lifecycle remains publicly reachable after fundraising closes, using the same public campaign identity/URL. Donation action is removed and result/accountability context takes priority.

Safety rule:

> Public lifecycle continuity preserves the public Campaign identity, not the internal Campaign record shape.

Key evidence:

- public Campaign Detail must use an explicit public-safe projection rather than inheriting the internal Campaign model;
- only Campaigns that genuinely entered the public lifecycle gain closed-state continuity;
- public Organization context and Campaign media also require deliberate public-safe projections;
- media access must be genuinely revocable after withdrawal/unpublish where product lifecycle requires it;
- donation eligibility remains backend-authoritative;
- Donation Flow remains a separate next-action dependency;
- public donor/social-proof capability is not automatically required.

The operational OpenAPI was intentionally not rewritten during the reframe. The narrow intended contract direction is recorded for Slice 1 reconciliation after promotion.

## Probe 02 — Account Registration + Email Verification

**PAUSED / REFRAMED — backward-reconciliation evidence retained; no longer an MVP gating task**

Artifact:

- `02-account-registration-email-verification.md`

Useful findings retained from the probe:

- generic registration responses / anti-enumeration intent are strong salvage candidates;
- concurrency guards, token single-use/revocation, transactional registration, PII protection, and applicable security tests are valuable implementation evidence;
- allowing login before email verification can be product-coherent when verification is capability-scoped;
- the old Account/Login spec contains stale cross-domain assumptions about verified-email requirements;
- the current `User.email_verified` / `primary_email` / auth-identity relationship becomes ambiguous when provider identifiers diverge;
- whole-backend reset and Account clean-room rewrite are not justified by current evidence.

The probe also exposed the more upstream problem that Account planning had become technical-first and feature-rich before the product needed that breadth.

The previously proposed canonical-email decision was therefore **not promoted**.

Under the approved MVP scope, Account remains outside the baseline critical path unless a real slice proves an enabling dependency.

When an Account-dependent capability becomes active scope, reuse this probe as backward-reconciliation evidence and derive the smallest coherent identity/authentication model from that concrete need.

## Validation closeout

The pre-promotion validation objective is complete enough for Product Authority promotion:

- forward derivation worked on Public Campaign Detail;
- backward reconciliation successfully distinguished reusable implementation evidence from stale/over-broad product assumptions;
- MVP scope corrected the tendency to let detailed historical domains drive product delivery;
- MVP sequencing was approved around the public trust loop;
- repository routing has been promoted so Product/MVP authority now sits above delivery specs/contracts for the concerns it owns.

The next validation mode is **real CRTV on actual MVP development**, beginning with Slice 1 through Harscode `main` and the canonical Exploration kickoff. No custom solution-steering prompt should be added.