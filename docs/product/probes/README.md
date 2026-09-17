# Kencleng Product Authority Probes

> Status: Validation evidence — **not product authority**
> Last updated: 2026-09-17

These probes validate the candidate Product Authority model against real delivery slices.

A probe may identify product, design, contract, or implementation gaps. Its conclusions do not become canonical merely because they are written here. Durable conclusions must be promoted deliberately into the authority that owns them.

## Validation status

### Candidate Product Authority human review

**PASS — 2026-09-17**

`docs/product/product-overview.md` has been reviewed by the human product authority and is directionally accepted as the Kencleng product model for validation purposes.

It remains non-canonical until the forward/backward probes complete and repository routing is deliberately promoted.

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

**READY TO START**

This backward-reconciliation probe starts from current Product Authority and derives the simplest correct registration/email-verification outcome before comparing it against the existing Account specs, OpenAPI, backend implementation, migrations, and tests.

Its goal is to classify existing Account work as `KEEP`, `ADAPT`, `REPLACE`, or `DEFER` and provide evidence for the backend soft-reset decision.
