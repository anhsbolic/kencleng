# Stage 2 Gap Evidence — O4 Status Credential

> Phase/Stage: Exploration / Stage 2  
> Author: `P-S2-002-OIR-002-1` (`KC-EXPLORER`)  
> Created: 2026-09-28  
> Work Unit / Run: `WU-S2-002` / `OIR-S2-002-002`  
> Model / reasoning: invocation dispatch metadata says `gpt-6-luna` / `high`; active runtime model is not independently exposed  
> Session: not exposed  
> Target revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535` (clean working tree at inspection)  
> Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`

## Current state

- Product fixes the guest status URL lifetime at 24 hours, requires a difficult-to-guess token, limits access to donation status, and says the donor need not revisit after expiry. The same generic public behavior/copy applies to invalid, missing, and expired links. Sources: `docs/product/mvp-scope.md` lines 116 and 126; `docs/product/mvp-delivery-slices.md` §5 lines 174–177.
- Techplan §13 O4 leaves credential carrier/storage, browser history/referrer/log/cache mitigations, token lifecycle/comparison, abuse controls, and residual-risk decision open. The Human Product decision did not accept these risks.
- Historical API `GET /donations/{donationId}/status` places `token` in the query string. Its description says it never expires. Historical `docs/spec/5-donation/features/02-donation-status-check.md` also says no expiry and expects `401` for missing token, wrong token, or absent donation. The draft threat model calls the non-expiring bearer token an accepted residual risk. Those historical choices conflict with current Product and do not carry forward as decisions.
- The historical API status operation returns the broad `Donation` schema, which includes `campaign_id`, `amount`, `payment_method`, `is_anonymous`, `guest_name`, `status`, and `created_at`; `status_token` is described as present only on immediate submission, but the same schema is referenced by the status operation. This leaves the exact status-only projection unclear and creates a schema-level misleading signal.
- The backend currently has no Donation route/domain package or status credential implementation. `backend/cmd/server/main.go` registers health/docs, Campaign, and Account routes. Therefore randomness, storage, comparison, expiry enforcement, URL hygiene, response projection, and abuse behavior have no runtime evidence in this revision.
- No current Donation cache/referrer/log policy is defined. Campaign public-media/detail code sets cache policy for its own resources; this does not establish policy for a private bearer status resource. Root `AGENTS.md` treats sensitive tokens as credentials and prohibits logging them. Crypto/key implementation is Tier-0 fenced.

## Requirement and gap

Product semantics are settled, while the credential and protection design is not. A future contract/runtime must demonstrate the hard 24-hour validity and a status-only projection, plus controls for how the bearer capability is conveyed, stored/checked, kept out of logs/history/referrers/shared caches, and protected against guessing or abuse. The evidence does not establish a suitable carrier, lifecycle/storage representation, or accepted residual risk. Exact projection may need Product clarification if “status only” cannot be translated without changing its meaning.

## Sniffing lenses

- **Risk:** Anyone obtaining the bearer credential can read the permitted status resource. URL query credentials can be copied into browser history, proxy/application logs, referrer data, screenshots, and cache keys. Extending expiry or returning excess fields can expose more data for longer. Residual risk acceptance is explicitly still open.
- **Edge cases:** access at the 24-hour boundary; clock skew; reused, revoked, malformed, or expired credential; URL without token; donation absent; credential exposed then later expired; browser/back-button and referrer behavior; status response cached by browser/proxy; token collision/rotation if later supported.
- **Miscontext:** Historical “never expires” and accepted-risk statements describe the old draft, not current Product. A token-only status route is not equivalent to authenticated identity, and Account token controls do not establish this guest capability.
- **Misleading signals:** `Donation.status_token` says it appears only on immediate submit, but the broad same `Donation` schema is attached to the status response. Presence of a token field and old `INV-donation-05` uniqueness language does not prove secure generation, expiry, carrier handling, or response minimization.
- **Inconsistency:** Product hard 24-hour validity conflicts with the old OpenAPI/spec/threat model's non-expiring token. Product/MVP authority is governing; stale detailed artifacts require reconciliation. Historical old accepted-risk text is not a current owner decision.

## Code and authority anchors

- `docs/product/mvp-scope.md` lines 116 and 126 — binding 24-hour limit, difficult-to-guess credential, status-only access, and generic expired-link behavior.
- `docs/product/mvp-delivery-slices.md` §5 lines 174–177 — active Slice 2 requirement and Security/PII ownership of credential exposure mitigations.
- `api/openapi/donation.yaml` `/donations/{donationId}/status` lines 80–106 — historical query-token shape, non-expiring description, and 401/404 conflict.
- `api/openapi/donation.yaml` `components.schemas.Donation` lines 225–262 — broad response projection and submit-only token property.
- `docs/spec/5-donation/features/02-donation-status-check.md` §§Summary/Behavior/Validation — historical no-expiry and failure behavior.
- `docs/spec/5-donation/threat-model.md` “Token-based status check” and “Knowingly accepted residual risk” — old accepted non-expiring-token posture; not current acceptance.
- `backend/cmd/server/main.go` router lines 154–188 — no Donation route wired at inspected revision.

Stage 2 makes no carrier/control comparison and accepts no risk. API/contract and Security/PII owner discussion belongs after the Stage 3 gate; any unresolved Product meaning around the exact status projection must be routed to Product Authority.
