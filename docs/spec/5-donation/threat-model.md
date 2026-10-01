# Threat Model — Donation

> File: `docs/spec/5-donation/threat-model.md`
> Status: draft — Slice 2 reconciliation; Security/PII and applicable Human review required before `agreed`
> Last updated: 2026-10-01
> Active slice: Slice 2 — Guest Donation + Truthful Donation State

## Scope and trust boundaries

Active surfaces are guest submission, backend-controlled sandbox processing, status-only guest revisit, and the Campaign eligibility/funding boundary. Account history, public donor list, claim, guest-email reveal, and real payment rails are deferred. Historical risk acceptances are not current authority.

| Actor / component | Authenticated? | Trust boundary |
|---|---|---|
| Guest donor | No | Sends amount, optional name, opt-in email, and submission request; may revisit status through temporary bearer URL. All request data is untrusted. |
| Donation API | Server-side | Validates request and eligibility, persists Donation state, and returns only approved public data. |
| Donation simulator | Backend-controlled | Owns terminal sandbox result; request/UI cannot select outcome. No client-callable settlement transition. Timing/scenario detail remains O2. |
| Campaign state | Server-side, Campaign-owned | Submission eligibility and close are ordered under D1; successful funding reflection remains atomic and exact-once. |
| Email delivery | External delivery boundary | No status/access content before ownership verification; verified opt-in remains eligible until its required terminal notice is fulfilled, with bounded, recoverable terminalization. Numeric bounds, architecture, timeout meaning, controls, and lifecycle evidence remain O3. |
| Holder of guest status URL | No | Possession grants temporary status-only access through the settled fragment URL/frontend handoff direction, one-way HMAC verifier, and hard 24-hour expiry from issuance; implementation controls and residual-risk decision remain O4/O5. |

## STRIDE by active surface

### Guest submission and retry

| Category | Concrete threat | Required mitigation / current direction | Residual risk / owner gate |
|---|---|---|---|
| Spoofing | A caller submits as another registered identity or forges authenticated context. | If authenticated, use validated session identity; guest fields cannot override it. Explicit backend authorization remains required where applicable. | Authentication implementation is outside this spec change. |
| Tampering | Invalid or fractional amount, unavailable method, altered payload on retry, or client-selected terminal outcome. | Whole IDR, minimum Rp5.000, Rp1 increments; monetary API/wire values are major-unit decimal strings with explicit currency codes; calculation/persistence are exact decimal with no `float`/`float64`. Only QRIS is active sandbox; backend owns result. Same key/same payload returns original, same key/different payload rejects. | O1 retains only supported currencies and concrete range/fraction/precision/scale/storage parameters; O2 simulator mechanics remain open. |
| Repudiation | Guest disputes whether a new donation was intentionally submitted after an ambiguous response. | Same idempotency key remains through ambiguity; client suppresses double-click; deliberate new intent uses a new key. | Retry-record lifetime/serialization is delivery detail for contract reconciliation. |
| Information disclosure | Guest name/email or Donation details leak through response/log/email before ownership verification. | Name not public by default; email opt-in/status-only; verify ownership before status/access email; established encryption/HMAC and safe-logging pattern. Preserve verified-address eligibility until required terminal notice fulfillment. | O3 retention, controls, deletion-race evidence, and no accepted residual risk. |
| Denial of service | Automated public submission burdens the service or creates unwanted records. | Minimum amount is a product validation, not a sufficient abuse control. | Security/API owner must determine applicable abuse controls; no prior sandbox acceptance is carried forward. |
| Elevation of privilege | External caller invokes a success/failure settlement transition. | No client-callable transition; backend simulator exclusively owns terminal state. | Requires later route/security review and runtime evidence. |

### Simulator and money-state transition

| Category | Concrete threat | Required mitigation / current direction | Residual risk / owner gate |
|---|---|---|---|
| Spoofing | Forged settlement causes a false successful donation. | Simulator is backend-controlled; no exposed HTTP settlement route. | Implementation and route audit remain downstream. |
| Tampering | Replay, concurrent transition, or partial commit changes status without matching funding or increments twice. | Successful status and full funding reflection commit atomically and exactly once; pending/failed do not count. | Mechanism is deliberately unspecified; Tier-1 review and specialized runtime Testing remain required. |
| Repudiation | Observable status and funding disagree about whether settlement occurred. | Every committed/observable state includes both success and full funding reflection, or neither. | Runtime atomicity, failure, replay, and concurrency evidence is deferred. |
| Information disclosure | Simulator leaks internal scenario configuration or non-public Donation data. | Expose only approved status; no internal fixture/configuration details in public responses. | O2 scenario detail remains open. |
| Denial of service | Simulator work is delayed or exhausted. | No timing/SLA is promised; pending copy is “Menunggu hasil simulasi.” | Timing and operational controls are O2; no availability promise is set. |
| Elevation of privilege | Donor controls whether the simulator returns success or failure. | Failure only through clearly labeled backend-configured/fixture demo scenario, never donor choice/request. | O2 mechanics remain open. |

### Temporary guest status URL

| Category | Concrete threat | Required mitigation / current direction | Residual risk / owner gate |
|---|---|---|---|
| Spoofing | Attacker guesses or steals the bearer credential. | Require a difficult-to-guess bearer credential; settled direction also uses a fragment-carried status URL with frontend handoff and URL cleanup, one-way HMAC verifier, hard 24-hour expiry from issuance, and status-only access. | Concrete credential generation/strength choice and evidence, storage, key/comparison controls, expiry enforcement, browser/infrastructure exposure protections, abuse controls, and residual-risk acceptance remain O4. No numeric entropy/length target is selected and this spec does not claim empirical verification. |
| Tampering | URL or parameters alter Donation or request a broader response. | Read-only status surface; do not expose mutation or details beyond status. | API shape is not defined in this task. |
| Repudiation | A donor cannot distinguish the legitimate status URL from an invalid/expired one. | User-facing generic copy is “Link status tidak tersedia atau mungkin kedaluwarsa.” | Design does not resolve technical response parity. |
| Information disclosure | Token appears in browser history, referrer, logs, cache, or forwarded URL; response reveals Donation existence. | Settled public behavior: absent Donation and missing/wrong/expired credential share a uniform `404`, identical Problem Details body, headers, and cache behavior, including `Cache-Control: private, no-store`. | Browser/referrer/log/cache protections and residual-risk acceptance remain O4; exact authored API expression is Task 02. Empirical response/timing parity and abuse evidence remain O5. Uniform behavior is a contract direction, not proof of parity. |
| Denial of service | Token guessing or repeated lookup exhausts resources. | No product-level abuse control selected here. | Security/API must decide and evidence rate/abuse handling (O4/O5). |
| Elevation of privilege | Credential grants access to email, account, or mutation capability. | Scope access to Donation status only. | Verify against final contract and implementation. |

### Optional notification email

| Category | Concrete threat | Required mitigation / current direction | Residual risk / owner gate |
|---|---|---|---|
| Spoofing | Unverified address is used to access a Donation or receive its status. | Verify ownership before sending any status or access link. A verified opt-in address remains eligible until its required terminal notice is fulfilled. | Verification window and controls remain O3. |
| Tampering | Donor-controlled content changes the notice into a payment request or implies real settlement. | Status-only notice; terminal result labeled simulation; no real-payment instructions. | Delivery template review is downstream. |
| Repudiation | Notice is sent more than once or before terminal result. | At most one notice for terminal success/failed; never initial pending. Delivery must provide bounded, recoverable terminalization. | Numeric bound, architecture, timeout meaning, retry/deduplication, and retention details remain O3. |
| Information disclosure | PII or Donation status reaches an unverified recipient or persists too long. | Optional opt-in, verify first, encrypt/HMAC, safe logs, delete unverified address after approved window; preserve verified-address eligibility until terminal notice fulfillment. | O3 controls and retention/deletion-race evidence remain unresolved; no residual risk accepted. |
| Denial of service | Repeated opt-ins or verification messages abuse delivery. | No detailed control selected in this spec. | O3 Security/PII and delivery owners decide controls. |
| Elevation of privilege | Email link grants broader access than status-only guest URL. | Any access link follows INV-donation-05 and O4/O5 gates. | Do not finalize contract fields until owner decisions. |

## Campaign eligibility, threshold, and close

| Threat | Required mitigation / current direction | Deferred evidence |
|---|---|---|
| Submission/close race accepts a new Donation after close or rejects one that won eligibility. | Apply D1: atomically order eligibility against close; close-first rejects; accepted Donation remains settleable in full. | Runtime concurrency evidence; no lock/isolation mechanism selected. |
| Accepted pending Donation settles after close and is lost or changes close reason. | Success and full funding commit exactly once; later settlement cannot reopen Campaign or replace winning close reason; funding can exceed `max_amount`. | Runtime integration evidence. |
| Donation spec imports unrelated closure/result lifecycle. | Scope this reference to eligibility, threshold, accepted pending Donation, and rejection after close only. | Broader closure/result remains Slice 3. |

## Knowingly accepted residual risk

None is accepted by this reconciliation. Historical statements accepting non-expiring credentials, missing abuse controls, public-list correlation, or other sandbox risks are not carried forward. Security/PII and API owners must resolve applicable controls and residual-risk decisions through O3–O5; O11/D19 direction is settled, while its delivery mechanics and lifecycle evidence remain under O3.

## Open items and downstream evidence

- **O1**: supported currencies and concrete range/fraction/precision/scale/storage parameters; shared representation direction is settled.
- **O2**: simulator timing and backend-controlled scenario mechanics.
- **O3**: email verification, retention, retry, and control evidence.
- **O4/O5**: status-link and uniform-404 directions are settled. Browser/infrastructure exposure, key/comparison controls, expiry enforcement, abuse controls, empirical response/timing parity, and residual-risk decision remain open.
- **O8 conditional**: consumer/distribution evidence before any historical operation removal/replacement.
- **O11/D19 direction is settled**: verified opted-in address remains eligible until required terminal notice fulfillment; Donation Delivery provides bounded, recoverable terminalization. O3 mechanics, controls, lifecycle/deletion-race evidence, and residual-risk gate remain open.
- Runtime Testing must cover atomic success/funding, replay, concurrency, submit/close ordering, accepted-pending settlement after close, stable winning close reason, threshold overshoot, status privacy, and retry behavior as assigned by the Approved Techplan.

## References

- `docs/spec/5-donation/invariants.md`
- `docs/product/mvp-scope.md` §§4–7
- `docs/product/mvp-delivery-slices.md` §§5–6
- `docs/spec/4-campaign/invariants.md#inv-campaign-13`
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` §§3–13
