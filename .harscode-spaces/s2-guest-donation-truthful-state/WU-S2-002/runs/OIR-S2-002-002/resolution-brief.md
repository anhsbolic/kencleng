# OIR Resolution Brief — O3–O5

> Phase/Stage: Explorer Open-Item Resolution / Stage 3 (complete)  
> Author: `P-S2-002-OIR-002-1` (`KC-EXPLORER`)  
> Created: 2026-09-28  
> Updated: 2026-09-28 — recorded Human decisions and terminal Stage 3 synthesis  
> Work Unit / Run: `WU-S2-002` / `OIR-S2-002-002`  
> Model / reasoning: invocation dispatch metadata says `gpt-6-luna` / `high`; active runtime model is not independently exposed  
> Session: not exposed  
> Target revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535`  
> Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`

Stage 3 synthesis is complete for this Run. O3–O5 have explicit owner direction where the evidence supported a choice, with residual-risk acceptance and runtime proof kept as separate later gates. This brief is not `CONTRACT_READY` evidence.

## O3 — Guest email verification, retry, retention

- **Outcome:** `PARTIALLY_RESOLVED`
- **Problem / current condition:** Product requires ownership verification before sending a guest donation status or access link. Unverified email is held only within a Security/PII-approved verification window and then deleted without sending. Product assigns verification, post-terminal retry windows, and retention controls to Security/PII. There is no Donation email runtime or sender/queue evidence in this revision. See Stage 2 evidence `evidence/stage-2-o3-guest-email.md`.
- **Options and trade-offs:** (1) Remove email verification and rely on donor-provided address plus disclaimer; lower donor friction, but wrong/third-party addresses can receive the bearer status link and this conflicts with active Product. (2) Retain verification; it adds a step but preserves the approved ownership requirement. For timing, (a) choose durations without delivery evidence, which enables contract closure but risks arbitrary retention/reliability, or (b) adopt a bounded proposal and defer the remaining pending-state cap until simulator timing is known.
- **Recommendation:** Retain verification; set a 24-hour verification window from guest email capture and a 24-hour post-terminal retry window. Delete unverified email at verification-window expiry. Delete verified email after successful terminal notification delivery or at retry-window expiry. Retry represents one logical notice and should avoid duplicate sends where sender semantics permit. Do not retain email indefinitely while status remains `pending`; maximum pending-state retention is deferred until the simulator's maximum terminal timing or an explicit Security/PII cap is known.
- **Observed Human decision:** On 2026-09-28, the Human accepted the recommendation to keep verification and use the recommended windows/retention: 24 hours for verification from email capture; 24 hours for delivery retry from terminal state; delete unverified email at verification-window expiry; delete verified email after successful delivery or retry-window expiry. The Human agreed to defer the maximum retention while status remains pending until relevant timing evidence or an explicit Security/PII cap is available. Provenance: direct Human messages in the Participant Session on 2026-09-28; no separate Product change was approved.
- **Authority:** Anhar Solehudin is the named current Slice 2 Security/PII owner and API/contract owner per `.harscode-spaces/authority-map.md` rows 11–12. Human approval of a Product change is not inferred. Product authority remains `docs/product/mvp-scope.md` §Stage B and `docs/product/mvp-delivery-slices.md` §5.
- **Residual risk:** Not accepted. An incorrect address can still be entered until ownership is verified; the verification step mitigates sending status/access to an unintended recipient. Delivery/retry idempotency is not implemented or tested.
- **Next route:** Preserve this policy in later spec/API reconciliation; resolve maximum retention for a verified address while Donation remains pending when O2 simulator timing or a Security/PII cap is available. Later Testing must verify no send before verification, deletion boundaries, and retry/dedup behavior.

## O4 — Status URL credential and residual risk

- **Outcome:** `PARTIALLY_RESOLVED`
- **Problem / current condition:** Product requires a difficult-to-guess token, a hard 24-hour URL lifetime, and status-only access. Historical API puts token in query and claims no expiry; no Donation runtime exists. URL exposure, storage, comparison, abuse, response projection, and residual risk remain. See `evidence/stage-2-o4-status-credential.md`.
- **Options and trade-offs:** (1) Query URL with log redaction, no-store, and referrer controls is simpler, but retains token exposure in URL surfaces. (2) Fragment URL, frontend sends token to API and clears fragment, reduces exposure on the initial request/log/referrer path but requires a client handoff and careful browser handling.
- **Recommendation:** Use fragment URL, enforce expiry 24 hours from token issuance without extending expiry on access, and return a purpose-built status-only projection. A one-way HMAC verifier is preferred over recoverable plaintext/ciphertext if the server does not need to resend the original token. Exact implementation controls and residual risk require evidence before acceptance.
- **Observed Human decision:** On 2026-09-28, Human selected option 2: fragment URL, with frontend handoff and URL cleanup as described in the option. This carrier decision was not itself residual-risk acceptance or a storage decision. Provenance: direct Human message in the Participant Session on 2026-09-28.
- **Observed Human decision:** On 2026-09-28, Human selected storage option 1: store a one-way HMAC verifier, since the original token need not be recovered for link re-send. Provenance: direct Human reply in the Participant Session on 2026-09-28. Key handling, comparison implementation, and operational controls remain implementation/evidence work; no Tier-0 crypto changes are authorized by this Run.
- **Authority:** Anhar Solehudin is named current Slice 2 Security/PII and API/contract owner. Human residual-risk acceptance remains a separate required gate.
- **Residual risk:** Not accepted. Fragment handling, URL cleanup, token storage/comparison, log redaction, cache behavior, rate/abuse controls, and runtime parity are unverified. The query-token and non-expiry choices in old Donation artifacts are not carried forward.
- **Next route:** API/contract reconciliation records the fragment carrier, frontend handoff/URL cleanup, HMAC verifier, hard 24-hour expiry, and status-only projection. Security/PII evidence must verify key handling, comparison, URL/log/cache protections, and abuse controls before any residual-risk acceptance. Later Testing must exercise 24-hour expiry, URL cleanup/exposure, status-only projection, and failure parity.

## O5 — Public failure parity / anti-enumeration

- **Outcome:** `PARTIALLY_RESOLVED`
- **Problem / current condition:** Product requires generic behavior for missing, invalid, and expired links. Historical feature spec specifies uniform `401`, while authored API lists both `401` and `404`; the shared 401 example describes an Account access token. No Donation runtime exists. See `evidence/stage-2-o5-public-failure-parity.md`.
- **Options and trade-offs:** (1) One uniform `404` for absent Donation, missing/wrong token, and expiry aligns with “status link unavailable” and the existing Campaign anti-enumeration precedent. (2) One uniform `401` treats all failures as invalid guest credentials but must not reuse JWT-specific text. In either case, all failure cases need the same Problem Details body, headers, and cache behavior; timing/abuse evidence remains downstream.
- **Recommendation:** Uniform `404` with the generic public wording; use `Cache-Control: private, no-store` and do not distinguish the failure cases in response body/headers/cache. Assess timing parity, constant-time secret comparison, and rate/abuse policy from the actual runtime/topology; no rate values are selected here.
- **Observed Human decision:** On 2026-09-28, Human selected a generic `404` response. The decision follows the recommendation that absent Donation, missing/wrong token, and expired token share one public failure result with identical body/headers/cache behavior. Provenance: direct Human reply in the Participant Session on 2026-09-28. A Product-style generic message should be used; the exact Problem Details coordinates and `Cache-Control: private, no-store` still need explicit contract reconciliation if not already adopted by the API owner.
- **Authority:** Anhar Solehudin is named current Slice 2 API/contract owner; Security review and explicit Human residual-risk acceptance remain required.
- **Residual risk:** Not accepted. Contract choice alone cannot establish timing parity, proxy-aware abuse protection, or cache behavior in runtime.
- **Next route:** API/contract reconciliation records the 404 response matrix and generic Problem Details body. Later Testing verifies equality across absent Donation, missing/wrong token, and expired token in status/body/headers/cache and relevant timing cases. Choose rate/abuse controls using the actual runtime/proxy topology; no numeric limit is selected here.

## Cross-item dependencies

- O3 verified email delivery may carry the O4 status credential. Verification stays required by current Product; removal would require a separate Product Authority change, not a Security/API reinterpretation.
- O4 determines how the browser presents the credential and how the API receives it; O5 applies the same 404 public failure behavior regardless of whether token lookup fails because of absence, mismatch, or expiry.
- O3 pending-state retention depends on O2's unresolved maximum terminal timing or an explicit Security/PII upper bound.
- None of these decisions accepts residual security/privacy risk or establishes runtime behavior.

## Verification and current state

- Verified by inspection: assignment-defining input checksums match; Stage 2 evidence exists for O3/O4/O5 and input provenance.
- Not performed: tests, runtime checks, API validation, deployment/topology analysis, or independent security review.
- Stage 3 decision facilitation is complete. Terminal handoff records these decisions, remaining concerns, and next-route recommendation; this does not close downstream contract reconciliation, risk acceptance, implementation, or Testing.
