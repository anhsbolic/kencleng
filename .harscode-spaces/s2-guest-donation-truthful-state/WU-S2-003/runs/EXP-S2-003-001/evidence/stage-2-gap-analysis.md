# Stage 2 — Gap Analysis

> Work Unit: `WU-S2-003`  
> Run: `EXP-S2-003-001`  
> Phase/Stage: Exploration / Stage 2  
> Author: `KC-EXPLORER` (`P-S2-003-EXP-001-1`)  
> Created: 2026-10-01  
> Model / Reasoning: Invocation requests `gpt-6-luna` / `high`; actual runtime model was not exposed  
> Target revision: `7fd8b473b239b20bda3990ab29c51440d321a796`  
> Workflow revision: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`  
> Session: not exposed

This evidence records live repository state observed during this Run. Code anchors are coordinates for later phases, which must reopen the live sources.

## Area 1 — Slice 2 requirements and accepted API contract

### Current state

- Product §5 defines guest submission, exact decimal money with whole-IDR minimum/increments, QRIS-only simulation, backend-owned `pending`/`success`/`failed`, idempotent ambiguous retry, optional guest fields, terminal-only verified status email, and a temporary status-only credential. The threshold is not a hard cap; accepted pending donations can settle in full after close.
- Donation Task 01/02, invariants INV-donation-01–11, and Donation threat model encode those directions and explicitly leave O1 parameters and O2–O5 controls/evidence open.
- Approved TP-S2-002-015 carries the accepted decisions and later evidence ownership. Its §12 assigns independent verification for exact money, idempotency, simulator authority, status privacy, exact-once funding, concurrency, D1 ordering, and post-close settlement. §13 keeps O1 numeric parameters and O2–O5 controls/evidence open.
- Authored `api/openapi/donation.yaml` expresses guest `POST /campaigns/{campaignId}/donations`, status `GET /donations/{donationId}/status`, major-unit decimal strings plus currency, `qris` only, status-only guest response, a fragment-carried credential handed to an optional header, and uniform `404` with `Cache-Control: private, no-store`. `common.yaml` supplies the shared Problem and idempotency-header components.
- The authored API retains historical public donor-list and Account operations. Product §5 and Task 01/02 defer the donor list and Account path; WU-S2-003 does not include them.

### Requirement and gap

Requirement: implement against the accepted Slice 2 Product/MVP, Donation specifications, Approved Techplan, and authored split API; do not treat generated or historical surfaces as new authority. The contract expresses requirements but makes no claim of live runtime behavior. In later areas, compare those requirements to backend routes, domain logic, persistence, and existing Campaign integration.

### Sniffing

- **Risk:** Wrong interpretation can widen the backend into Account, donor-list, real-payment, or Slice 3 behavior, or mistake contract prose for runtime proof.
- **Edge cases:** The requirement includes Rp5.001 as valid, threshold overshoot, accepted pending settlement after close, same-key/same-payload retry, same-key/different-payload rejection, and four status-lookup failure cases with identical public behavior.
- **Miscontext:** `api/openapi/donation.yaml` has historical operations beyond the active guest submit/status scope; their presence does not put them in this Work Unit.
- **Misleading signals:** API schema presence and the accepted `CONTRACT_READY` milestone are contract evidence, not backend implementation or O2–O5 empirical evidence.
- **Inconsistency:** `docs/spec/5-donation/tasks.md` has an agreed document header and active Task 01/02 acceptance, but its status tracker rows still say `draft / owner review required`. The 2026-10-01 Event records acceptance of all five Donation specs; the Parent Outcome and Work Graph record `CONTRACT_READY`. Treat the rows as stale tracking text unless a later authoritative status contradicts this evidence; do not infer that acceptance includes residual-risk or runtime evidence.

### Code/source anchors

- `docs/product/mvp-delivery-slices.md` §5 — current Slice 2 scope and correctness floor.
- `docs/spec/5-donation/tasks.md` Task 01, Task 02, Verification ownership, and Status tracker — task behavior, O2–O5 evidence, and the stale row discrepancy.
- `docs/spec/5-donation/invariants.md` INV-donation-01–11 — delivery invariants and explicit deferrals.
- `docs/spec/5-donation/threat-model.md` — trust boundaries and active threats.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` §§2–5, 12–13 — Approved delivery scope, execution requirements, verification and open items.
- `api/openapi/donation.yaml` `/campaigns/{campaignId}/donations`, `/donations/{donationId}/status`, `SubmitDonationRequest`, `Donation`, and `DonationStatusResponse` — authored contract coordinates.
- `api/openapi/common.yaml` `IdempotencyKeyHeader`, `Problem`, and Problem responses — shared contract coordinates.
- `.harscode-spaces/s2-guest-donation-truthful-state/events.md`, 2026-10-01 entries for Task 01 acceptance, RV-014, and Task 02 API acceptance — acceptance/contract-readiness evidence.

## Area 2 — Backend Donation submission, amount, idempotency, and simulator

### Current state

- Targeted inventory/search of `backend/internal`, `backend/migrations`, and `backend/cmd` found no Donation domain package, Donation migration, submission/status handler, idempotency persistence, simulator, or settlement service.
- `backend/cmd/server/main.go` registers only health/docs, Campaign public detail/media, Auth, and Account routes. There are no Donation routes.
- The only Campaign schema is migration `000011_create_public_campaigns`; its header explicitly says it does not introduce donation workflows. Its amount columns serve the Slice 1 public Campaign projection.
- Existing Campaign code uses `shopspring/decimal` to map public funding values. This is evidence of an existing decimal dependency/pattern, not authority for Donation-specific storage scale or currency parameters.

### Requirement and gap

Requirement: the Donation API must persist a valid guest request as `pending`, enforce exact decimal/whole-IDR rules, keep request idempotency separate from settlement replay, and let only backend-controlled simulator behavior determine terminal status. Same-key/same-payload retries return the original Donation; same-key/different-payload requests reject. QRIS is the only accepted method. None of this backend-owned behavior exists in the inspected runtime.

### Sniffing

- **Risk:** No guest Donation can currently be submitted or revisited through the backend; the Work Unit's backend capability is absent, not partially implemented.
- **Edge cases:** No runtime path exists to assess Rp5.000/Rp5.001 boundaries, invalid fractions/currencies, ambiguous retries, idempotency-key payload mismatch, donor-selected outcomes, or concurrent/replayed terminal transitions.
- **Miscontext:** Accepted OpenAPI operations describe the target contract; they do not correspond to registered handlers or persisted behavior today.
- **Misleading signals:** The Campaign response contains `donation_action=unavailable/donation_flow_not_available` and explicitly documents it as the Slice 1 posture. Generated API types or a future UI cannot turn that into backend capability.
- **Inconsistency:** The authored Donation OpenAPI exposes submit/status paths while current backend router and migration inventory contain no such operations/schema. This is the concrete implementation gap for WU-S2-003, not evidence that the API authority is unaccepted.

### Code anchors

- `backend/cmd/server/main.go`, `run` router registration near `mux.HandleFunc` — current registered route set; no Donation operation.
- `backend/internal/domain/campaign/entity.go`, `DonationAction` — Slice 1 unavailable action model.
- `backend/internal/domain/campaign/service.go`, `toPublicDetail` — currently returns `donation_flow_not_available`; `mapFunding` uses decimal for the public read model.
- `backend/migrations/000011_create_public_campaigns.up.sql` — current Campaign/public projection schema and explicit non-donation scope.
- `api/openapi/donation.yaml`, submission/status operations and schemas — contract coordinates only.

## Area 3 — Campaign eligibility, close ordering, and funding integration

### Current state

- Campaign's live backend is a Slice 1 public read model. `campaign.Repository` exposes public detail/media reads and a separate seed write; it has no Donation-facing eligibility/close/funding mutation port.
- `FindPublicDetail` selects Campaign rows where `status = 'published'`; `toPublicDetail` reports `public_state: fundraising` and `donation_action` unavailable. It does not implement Slice 2's atomic submit-vs-close ordering.
- Migration `000011_create_public_campaigns` stores `target_amount` and `collected_amount` as `NUMERIC(19,2)` and permits status `closed`, but defines no `closed_reason` or close trigger workflow. Targeted backend source search found no close/force-close/deadline-scheduler implementation.
- Campaign INV-campaign-13 assigns lifecycle and winning `closed_reason` to Campaign, and accepted-submission/settlement contribution to Donation. It requires close-first rejection, accepted-pending full settlement after close, exact-once funding, stable close reason, and possible threshold overshoot. D1 leaves mechanism unspecified.

### Requirement and gap

Requirement: Donation acceptance must atomically win or lose eligibility against Campaign close. Success and the full funding reflection must commit together once; accepted pending Donations can settle after close without reopening Campaign or changing its winning reason. Current backend has neither close transition/close reason nor a mutation boundary that can participate in this ordering. The existing public read predicate cannot serve as evidence of atomic eligibility.

### Sniffing

- **Risk:** A Donation-only check against `status='published'` cannot establish D1 ordering if a close transition races submission. Partial integration could accept after close, lose an accepted pending donation, duplicate funding, or rewrite Campaign's winning close reason.
- **Edge cases:** Submit and close racing both ways; deadline/threshold/Admin close contenders; settlement after close; concurrent successful Donations; threshold overshoot; nil funding pair versus partial funding pair.
- **Miscontext:** Public visibility (`status='published'`) and donation eligibility are related but not interchangeable. The current public detail query does not expose a transactional Donation eligibility guarantee.
- **Misleading signals:** `NUMERIC(19,2)` and Campaign's `shopspring/decimal` read mapper are Slice 1 precedent, not a resolved Donation storage scale or authority to reuse all Campaign money representation details. The project monetary standard and O1 deferrals govern.
- **Inconsistency:** Current backend Campaign lifecycle remains absent while the accepted Campaign invariant assigns close/`closed_reason` ownership and requires runtime ordering. This is a delivery capability gap; D1 semantics themselves are settled. The Work Unit's boundary excludes broader Slice 3 closure/result behavior, so the exact bounded integration scope needs explicit Techplan/Orchestrator treatment.

### Code/spec anchors

- `docs/spec/4-campaign/invariants.md#inv-campaign-13` — lifecycle ownership, D1 ordering, stable close reason, and Testing obligation.
- `docs/spec/5-donation/invariants.md#inv-donation-02` and `#inv-donation-08` — Donation-side ordering and exact-once contribution.
- `backend/internal/domain/campaign/repository.go`, `Repository` / `SeedRepository` — available Campaign seams.
- `backend/internal/domain/campaign/repository_db.go`, `FindPublicDetail` — current `status='published'` read predicate.
- `backend/internal/domain/campaign/service.go`, `toPublicDetail` / `mapFunding` — Slice 1 projection and current unavailable Donation action.
- `backend/migrations/000011_create_public_campaigns.up.sql`, `campaigns` — current funding columns/status constraint; no close reason.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` §12 R3/R7 and Test Focus Pointer — independent atomicity/concurrency/D1 verification requirements.

## Area 4 — Guest status credential, private response, and abuse boundary

### Current state

- There is no Donation credential issuance or status handler to inspect. The accepted API carries a bearer credential in a URL fragment for frontend handoff, then an optional `X-Donation-Status-Credential` header; the backend must return status only.
- INV-donation-05 and the authored API settle one-way HMAC verification, hard 24-hour expiry, and generic `404` for absent Donation/missing/wrong/expired credential with identical response behavior and `Cache-Control: private, no-store`. O4/O5 still leave credential generation/strength, key and comparison controls, expiry enforcement, browser/referrer/log/cache protections, abuse controls, empirical parity/timing, and residual-risk acceptance unproven.
- Existing Campaign HTTP handlers set `Cache-Control: private, no-store` on success and failure. Their status body is not the Donation status contract.
- Existing `platform/crypto.HMAC` is keyed with configured `HMAC_KEY`; code uses it for account/email identifier lookup hashes. Backend architecture describes `HMAC_KEY` for corresponding PII `*_hash` columns, separate from `ENCRYPTION_KEY`. There is no Donation credential-specific key config or verifier. `crypto/` and auth core are root-fenced read-only paths.
- `RateLimit` exists with idle-key eviction, but router wiring applies it to `/auth/` and Account routes, not current public Campaign routes; there is no Donation route or Donation abuse control. The middleware keys on `RemoteAddr` and documents the reverse-proxy caveat.
- Shared `MapServiceError` returns a generic public 500 but logs unhandled `err` verbatim. Best-practice guidance warns that wrapped errors can carry PII/tokens; no Donation error chain exists yet to audit.
- Donation API declares `Cache-Control` on status `404`, not on status `200`. INV-donation-05 leaves status cache/exposure controls open; do not infer successful-response cache behavior from the error contract. Current Campaign handler's no-store behavior is a relevant existing anchor.

### Requirement and gap

Requirement: status lookup must be temporary, difficult to guess, private, and status-only; credential checks and expiry must avoid enabling enumeration; all 404 cases must match exactly as authored; sensitive values must not reach ordinary logs or unapproved caches. No backend implementation exists, and some exact controls remain explicitly owner-gated O4/O5 questions.

### Sniffing

- **Risk:** A stolen/guessed credential exposes private status. Request logs, wrapped errors, response caches, browser/referrer behavior, or unbounded guessing could leak the credential or existence of a Donation.
- **Edge cases:** Missing header; malformed/unknown Donation ID; wrong credential; expiry boundary; expired credential; equal status/body/headers/cache behavior and timing across all 404 cases; repeated guesses; key rotation/configuration and comparison behavior.
- **Miscontext:** `HMAC_KEY` supports PII lookup hashes, not automatically a safe status-credential key. The Harscode key-management guidance says keys have one purpose; status-verifier key use/separation needs the appropriate owner decision and cannot be fixed by editing the fenced crypto package in this Run.
- **Misleading signals:** A one-way HMAC in the OpenAPI description does not prove credential strength, constant-time comparison, expiry enforcement, cache/referrer/log protection, or parity. Generic 404 prose also does not prove equal runtime timing.
- **Inconsistency:** Authenticated route rate limiting is scoped to `/auth/` and Account, while a future guest Donation status route is public. No owner-selected Donation abuse control exists. Successful status caching is not expressed as an API response header, while its protections remain an O4 obligation; route to Security/API ownership rather than silently treating the 404 header as covering 200.

### Best-practice routing and code/spec anchors

- Harscode `best-practices/index.md` Security Concern Map and `best-practices/go/index.md` route this surface to `go/secrets-and-sensitive-logging.md`, `go/secrets-and-key-management.md`, `go/rate-limiting.md`, `go/authorization-and-idor.md`, `postgresql/encryption-at-rest.md`, and `restapi/anti-enumeration.md`.
- `docs/spec/5-donation/invariants.md#inv-donation-05` / `docs/spec/5-donation/threat-model.md` — credential direction and remaining O4/O5 controls.
- `api/openapi/donation.yaml`, status operation — credential header, status-only response, uniform 404 and declared cache header.
- `backend/internal/platform/crypto/crypto.go`, `HMAC`; `backend/internal/platform/crypto/keys.go`, `Keys` / `New` — existing PII HMAC machinery, read-only due Tier-0 fence.
- `docs/project/kencleng-backend-tech-stack.md`, Encryption Key Management — current key-purpose allocation.
- `backend/internal/transport/http/middleware.go`, `RateLimit`; `backend/cmd/server/main.go`, route registration — current rate-limit scope and router.
- `backend/internal/transport/http/campaign_public.go`, cache-control and Problem writers — existing no-store precedent.
- `backend/internal/transport/http/errors.go`, `MapServiceError` — generic public errors and raw internal error logging boundary.
- Harscode `best-practices/restapi/anti-enumeration.md` — constant-time secret comparisons and response/timing parity guidance.

## Area 5 — Guest email verification and terminal status notification

### Current state

- No Donation-specific email field persistence, verification lifecycle, terminal notice, retry worker, or delivery record exists in the inspected backend.
- `platform/notification.Sender` is an Account seam with only verification, nudge, and password-reset methods. `FakeSender` logs only that Account email was queued and does no network delivery. In development `DevSender` writes a simulated inbox/outbox file with recipient and token, with owner-only directory/file modes; it is not the structured log stream.
- Server wiring selects `FakeSender` outside development and `DevSender` only in development. No real SMTP sender is configured. The accepted Donation behavior requires a status-only terminal notice when opted-in email is verified; whether the current fake/development behavior is an acceptable sandbox realization of that user-facing promise is not stated in the inspected current Product/API authority.
- Account verification is explicitly separate from guest email verification. Existing `auth_tokens` rows reference users and have Account-specific purpose/redeem semantics; reusing the Account route would impose an Account prerequisite and conflicts with the accepted guest path.
- O3 keeps implementation controls, verification/retry windows, retention/deletion behavior, terminalization bounds, timeout meaning, deletion-race evidence, and residual-risk acceptance open. Product copy discloses a 24-hour verification period, but this does not resolve the separate post-terminal delivery retry window or terminalization mechanics.

### Requirement and gap

Requirement: guest email is optional and opt-in, encrypted/HMAC-protected, and not public. Verify ownership before any status/access email. At most one status-only, simulation-labeled email may be sent on terminal `success`/`failed`, never on initial `pending`. Unverified email is deleted after its approved window; verified email remains eligible until the notice is fulfilled, with bounded/recoverable terminalization. Current backend contains only Account email behavior and no actual Donation delivery path. The existing production `FakeSender` does not send an email.

### Sniffing

- **Risk:** Status or access content could go to an unverified address; plaintext email or access credential could leak; failed post-commit delivery could lose the required notice; retries could send duplicates; retaining PII without a finite lifecycle violates purpose limitation.
- **Edge cases:** Terminal result before verification; verification racing deletion at 24 hours; terminal notice delivery failure/retry and duplicate sends; verified address retained while notice pending; pending never sends; address and Donation expiration/retention; simulated development outbox retention.
- **Miscontext:** Account email verification cannot stand in for guest email ownership verification: it is user-bound, and Product says it is separate and cannot gate guest Donation.
- **Misleading signals:** Existing `FakeSender` returns nil after logging “queued,” which is not proof of delivery; development outbox is a simulated inbox, not external email. A configured sender interface does not itself fulfill the optional guest-notice outcome.
- **Inconsistency:** Product/Donation acceptance promises an optional terminal status email, while current production sender behavior discards messages and current API/Donation backend has no guest verification or delivery. Whether simulated delivery suffices for this sandbox is an unresolved product/delivery boundary; resolve it with Product/Delivery owners before claiming notification fulfillment.

### Best-practice routing and code/spec anchors

- Harscode `best-practices/index.md` Security Concern Map routes PII/encryption and sensitive logging to `postgresql/encryption-at-rest.md` and `go/secrets-and-sensitive-logging.md`; those guides were checked. The latter warns against logging error values that may wrap recipient/token data.
- `docs/product/mvp-scope.md` Slice 2 guest email direction; `docs/product/mvp-delivery-slices.md` §5; `docs/spec/5-donation/invariants.md` INV-donation-03/04; `docs/spec/5-donation/features/01-submit-donation-settlement.md` notification acceptance — purpose and O3 boundaries.
- `backend/internal/platform/notification/sender.go`, `Sender` / `FakeSender` — current Account-only interface and no-network sender.
- `backend/internal/platform/notification/dev_sender.go`, `DevSender` / `append` — simulated local inbox and file permissions.
- `backend/cmd/server/main.go`, `newEmailSender` — environment-specific delivery wiring.
- `backend/internal/domain/account/service.go`, `Register` / `VerifyEmail` / `sendVerification`; `backend/internal/domain/account/repository_db.go`, token persistence/redeem — Account-only flow, not guest semantics.
- `docs/project/kencleng-backend-tech-stack.md`, Encryption Key Management and Notification — established key/data pattern and lack of configured SMTP sender.

## Area 6 — Risk tier, verification ownership, and available test evidence

### Current state

- Donation Task 01/02 are Tier 1 in `docs/spec/5-donation/tasks.md`; WU-S2-003 says Donation is Tier 1 baseline and asks Exploration/Techplan to assess O2–O5 obligations. Project workflow Tier 1 requires independent review/testing and Human review before merge.
- Project workflow defines Tier 0 for donation-ledger locking, money-rounding/calculation core, and encryption/key-handling core. Root `AGENTS.md` fences the Donation ledger/transaction-locking path and crypto/auth core. WU-S2-003 authorizes no protected-path write; if needed, surface it for Human-paired re-scope.
- Backend `Makefile` exposes lint, unit, full race, contract-tag, and security targets. Existing PostgreSQL integration tests use `//go:build integration`, testcontainers, migrations, and Postgres 16. The Makefile does not provide a separate integration target.
- There are no Donation unit, contract, or integration tests because there is no Donation implementation. The approved Techplan assigns later runtime/testing to exact money, idempotency, simulator authority, status privacy, atomic settlement/funding, replay/concurrency, D1 orderings, accepted-pending-after-close, stable close reason, and threshold overshoot; O2–O5 require downstream owner/security/runtime evidence.
- No tests, contract validation, migrations, or executable checks were run in this Exploration Run.

### Requirement and gap

Requirement: Tier 1 evidence must independently establish the acceptance criteria and threat-model controls. Race detection alone does not prove an invariant; the close, settlement, and funding claims require invariant-asserting concurrency and real persistence evidence. Contract validation/generation from WU-S2-002 only establishes contract source correspondence, not backend runtime correctness/security.

### Sniffing

- **Risk:** A green unit/race or OpenAPI validation result could be mistaken for exact-once funding, D1 ordering, status privacy, O3 retention, or O5 parity; failures affect monetary truth and guest PII.
- **Edge cases:** Concurrent submit-vs-close in both orders; duplicate and concurrent settlement; database rollback between Donation status and funding; accepted pending settlement after close; verification/deletion races; status response/header/cache/timing parity; abuse behavior.
- **Miscontext:** The current Campaign integration test proves migration/constraint behavior for the Slice 1 schema only; it is not Donation settlement or concurrency evidence.
- **Misleading signals:** `make verify` runs broad unit/race/contract/security checks but does not include the `integration` build tag. The previously reported API validation/bundle/type generation is contract evidence, not execution proof.
- **Inconsistency:** Task verification ownership and Approved Techplan correctly leave runtime/security evidence downstream; no current backend tests contradict them. Tier-1 remains the feature baseline while any need to change a protected Tier-0 surface would require a separate Human gate.

### Code/source anchors

- `docs/spec/5-donation/tasks.md`, Verification ownership — required invariant/runtime evidence and tier.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` §12 and Test Focus Pointer — detailed independent testing evidence.
- `docs/kencleng-agentic-workflow.md` §4 and §15 — Tier definitions and risk-driven evidence ownership.
- `backend/AGENTS.md`, Testing conventions; `backend/Makefile`, `verify` / test targets — stack-specific command routing.
- `backend/internal/domain/campaign/repository_integration_test.go`, `//go:build integration` / `isolatedCampaignPostgres` — current PostgreSQL testcontainers pattern.
- Root `AGENTS.md` §3 — protected Donation ledger, transaction/locking, crypto and auth paths.

## Findings and progression effect

### F-01 — Donation backend capability is absent

**Finding:** No Donation routes, domain package, persistence schema, simulator, guest status handler, or Donation tests exist in the inspected backend. The accepted contract is not runtime evidence.  
**Effect:** Informational for Stage 3; it confirms the delivery gap and gives Techplan concrete implementation areas. It does not prevent Stage 3 solutioning.

### F-02 — D1 needs a Campaign close/funding integration boundary that does not exist

**Finding:** D1 requires atomic ordering against Campaign close, but current backend has only a Campaign read model and no close transition, persisted `closed_reason`, or Donation-facing funding mutation seam. WU-S2-003 requires D1 while excluding broader Slice 3 closure/result behavior.  
**Effect:** **Blocking for D1 implementation and the `BACKEND_VERIFIED` milestone** until the bounded backend scope/owner route is resolved. This does not block Stage 3: Stage 3 can compare bounded integration directions and route any WU scope change to the Orchestrator; changes to root-fenced ledger/transaction-locking remain Human-required. Other Donation behavior can be solutioned in parallel.

### F-03 — O4/O5 controls are not implemented or fully expressed

**Finding:** Credential generation/strength, dedicated key purpose, constant-time comparison, expiry, cache/log/referrer/browser protections, abuse handling, status-200 cache behavior, and empirical response/timing parity are not present or proven. The current status route does not exist.  
**Effect:** Decision-relevant and needs Security/API owner routing before status access can be accepted as secure. Stage 3 can proceed to record decisions/evidence needs; runtime parity and residual-risk acceptance remain later Testing/Human gates.

### F-04 — O3 guest-notification delivery realization is unresolved

**Finding:** Product requires an optional terminal status email, but production `FakeSender` does not deliver and no guest verification, persistence, retry, or terminalization path exists. Current authority does not say whether simulated delivery suffices for this sandbox experience.  
**Effect:** Decision-relevant to the email-enabled behavior and any `BACKEND_VERIFIED` claim that includes it. Stage 3 can frame the Product/Delivery decision and preserve owner gates; no guest-notice fulfillment can be claimed meanwhile.

### F-05 — Remaining Delivery/Testing evidence is unproduced

**Finding:** Exact-money boundaries, request idempotency, simulator authority, D1 ordering, settlement exact-once/concurrency, O3 lifecycle races, status parity/abuse, and credential controls have no backend runtime evidence.  
**Effect:** Deferable from Stage 3 to implementation and independent Testing where the Approved Techplan assigns them. Race detector alone is insufficient for business invariant proof. No test/check was run during this Exploration.

### F-06 — Task acceptance tracker rows are stale

**Finding:** `tasks.md` document status and the accepted Event/Work Graph agree, but the Task 01/02 status rows still say `draft / owner review required`.  
**Effect:** Informational; the Event records acceptance and `CONTRACT_READY` evidence. Reconcile the tracker in its owning status/orchestration surface; do not treat the stale rows as an unresolved Product or spec decision.
