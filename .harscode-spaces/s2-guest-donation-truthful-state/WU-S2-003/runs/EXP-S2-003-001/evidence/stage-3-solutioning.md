# Exploration Evidence — Stage 3 Solutioning

## Provenance

- Phase/Stage: Exploration / Stage 3 — Solutioning
- Work Unit: WU-S2-003
- Run: EXP-S2-003-001
- Role / Author: Explorer / KC-EXPLORER (P-S2-003-EXP-001-1)
- Created: 2026-10-01
- Model / reasoning: Invocation requests gpt-6-luna / high; actual runtime model was not exposed
- Session: not exposed; continued from durable Stage-2 evidence as authorized in the Invocation
- Target revision: 7fd8b473b239b20bda3990ab29c51440d321a796 (Invocation baseline; re-read routed live sources)
- Workflow revision: Harscode pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8
- Human gate: Stage 3 authorized by Anhar on 2026-10-01 after Stage 2

This artifact compares delivery paths under current approved authority. It does not change Product/API/spec authority, choose a Tier-0 mechanism, authorize implementation, or accept security/privacy residual risk.

## Governing directions carried forward

The current-effective approved source is .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md (Status: Approved). Its requirements and open items narrow several Stage-2 findings:

- **D1 is settled policy.** Submission eligibility is ordered atomically against Campaign close; accepted-while-eligible Donations remain settleable in full after close; close-first rejects new submission; successful settlement and funding update commit together exactly once; later settlement cannot reopen Campaign or alter the winning close reason. No locking/isolation mechanism is selected, and broader Slice 3 closure/result behavior remains out of scope. See TP-015 §§2–3, 8, and 12–13; Product mvp-delivery-slices.md §5.
- **O3 product behavior is settled, technical fulfillment is not.** Email is optional/opt-in, separately verified, and may receive at most one status-only simulation-labeled notice at terminal success/failed, never pending. Verified address remains eligible through fulfillment under bounded/recoverable terminalization. Controls, architecture, retention/deletion-race evidence, and residual-risk acceptance remain with Security/PII and Donation Delivery. See TP-015 §§3, 8, and 13.
- **O4/O5 contract direction is settled.** Fragment URL with frontend handoff and URL cleanup, one-way HMAC verifier, hard 24-hour lifetime, status-only response, uniform public 404, identical body/header/cache behavior, and Cache-Control: private, no-store. Carrier detail where still unselected, credential/exposure/key/comparison/lifecycle/abuse controls, empirical parity, and Security/PII residual-risk acceptance remain open with their owners. See TP-015 §§3, 8, and 13.
- **O1 shared representation is settled; numeric bounds are not.** Use major-unit decimal strings with explicit currency, exact decimal calculation/persistence, and no float64; Slice 2 input remains whole-Rupiah IDR. Do not infer Donation storage precision/range from Campaign NUMERIC(19,2). See TP-015 §§3, 8, and 13.

Accordingly, Stage-2 F-02 remains an implementation capability/scope-route blocker for D1 and BACKEND_VERIFIED, not an unresolved Product decision. F-03/F-04 identify missing runtime capabilities and owner evidence, not unmade O4/O5 policy or permission to treat a fake sender as delivery.

## Direction considered

| Decision area | Viable paths | Recommended direction | Consequence / boundary |
|---|---|---|---|
| D1 Campaign/Donation integration | **A.** Build only the Campaign lifecycle pieces required to expose an authoritative eligibility/close/funding coordination boundary to Donation, preserving Campaign ownership and settled ordering. **B.** Defer that boundary to a separate Campaign lifecycle Work Unit and make Donation delivery depend on it. | **A, bounded to the D1 boundary.** The approved Work Unit requires D1 and TP-015 explicitly says coordinate only eligibility, threshold crossing, accepted pending Donations, and rejection after closure. Backend architecture is a domain-driven monolith with separate campaign and donation packages and hexagonal ports. Techplan should name the narrow inter-domain seam and decide its owner/files before Build. | Keep full closure/result lifecycle, scheduler, force-close, and public result behavior outside WU-S2-003. If this seam requires writing the root-fenced Donation ledger/transaction-locking path, do not route around the fence: split/re-scope for Human-paired work. D1 verification must cover both submit/close orderings and post-close settlement. |
| D1 when current Campaign code has only a public read model | **A.** Treat status='published' read as sufficient and add Donation independently. **B.** Reconcile the missing Campaign capability before claiming D1. | **B.** A visibility/read predicate cannot order concurrent submission against closure. It leaves an accepted race and cannot satisfy the accepted ordering invariant or preserve the winning close reason/exact-once funding claim. | This is a delivery dependency, not a change to D1 semantics. The Orchestrator must confirm whether the narrow boundary fits WU-S2-003 or needs an explicitly coordinated Campaign sub-scope/Work Unit before Build. Until resolved, D1 implementation and BACKEND_VERIFIED remain blocked; other planning can proceed. |
| Guest terminal notification | **A.** Treat existing FakeSender/development outbox as fulfillment. **B.** Deliver the approved optional guest notice through an explicit guest-capable notification path, with verified ownership, one terminal status-only message, and bounded/recoverable lifecycle. | **B.** Product requires the notice behavior when a guest opts in and verifies; a fake sender that only logs/returns nil does not deliver it. Reuse the established notification boundary where feasible, but keep transport/provider, retry/retention, and exact mechanism open for the assigned Product/Delivery/Security owners. | No real provider/vendor is selected by this Exploration. If Slice 2 must provide production delivery, the current stack's lack of configured SMTP/provider is a concrete dependency to resolve in Techplan/Orchestration. Do not report the email behavior as fulfilled by a simulated local inbox. |
| Public status privacy and abuse | **A.** Implement only the authored fragment/header and uniform-404 response. **B.** Implement the settled contract and carry all open Security/API controls and evidence as named obligations. | **B.** The contract is a necessary baseline, not proof against credential leakage, guessing, cache exposure, or timing distinctions. Preserve the selected fragment/handoff/cleanup, one-way verifier, hard expiry, status-only projection, uniform 404 and private,no-store; route credential strength/key purpose/comparison, logging/referrer/browser/cache/abuse controls to Security/PII/API and later Testing. | Do not invent token entropy, a dedicated-key scheme, rate-limit policy, proxy trust rule, or accept residual risk in Exploration. Never edit root-fenced crypto/auth code here. 200 cache handling should be explicit in the API/implementation review because current authored API explicitly states the directive for 404. |
| Donation monetary persistence | **A.** Copy Campaign's NUMERIC(19,2) schema precedent. **B.** Carry approved exact-decimal representation and derive storage bounds only from supported requirements/evidence. | **B.** TP-015 explicitly marks Campaign precision as precedent only and keeps numeric range, universal scale, and migration detail open. | Techplan may design exact-decimal mapping, but must surface any required range/scale as a decision/evidence need before contract or persistence locks it. Keep whole-IDR validation and decimal wire values consistent with approved authority. |

## Material alternatives rejected

- **Donation-only published check as D1:** rejected because it does not serialize submission against close and cannot satisfy the accepted ordering invariant.
- **Full Campaign lifecycle inside WU-S2-003:** rejected because it expands into Slice 3 scheduler, force-close, and public-result behavior; the approved boundary needs only D1 eligibility/threshold coordination.
- **Selecting a database lock/isolation strategy in Exploration:** rejected because D1 deliberately leaves mechanism unspecified, the protected transaction/locking path is fenced, and this phase must not preempt Techplan or Human-paired Tier-0 work.
- **Treating FakeSender or development outbox as production notice delivery:** rejected because current behavior is logging or a local simulated inbox, not fulfillment of the opt-in terminal email promise.
- **Using Campaign NUMERIC(19,2) for Donation by precedent alone:** rejected by the project monetary standard and TP-015 O1 direction.
- **Claiming O4/O5 acceptance from the chosen contract alone:** rejected because implementation controls, empirical evidence, and Security/PII residual-risk acceptance remain separate obligations.

## Remaining owner and evidence routes

1. **Orchestrator / Techplan:** resolve the concrete D1 scope route using current Campaign code and the root Tier-0 fence. Name the smallest Campaign-owned seam and ownership/files, or create an explicit dependency route. Do not claim D1 complete until both orderings and accepted-pending-after-close are implementable and independently verifiable.
2. **Donation Delivery with Security/PII and API owners:** define guest email verification, protection, bounded/recoverable terminalization, deletion/race handling, retry behavior, and actual sender capability. Keep approved product behavior intact. Any residual security/privacy acceptance remains a separate Human gate.
3. **Security/PII, API, and Testing:** define/evidence status-token strength and lifecycle, key/comparison handling, browser/referrer/log/cache protections, abuse controls, status 200 cache policy, response/body/header/timing parity, and runtime behavior. Do not translate missing evidence into accepted risk.
4. **Donation Delivery / shared monetary-standard owner:** carry exact decimal/no-float behavior and whole-IDR input; resolve concrete range/precision/storage requirements only when contract/storage needs are evidenced, without importing Campaign's scale.
5. **Donation Delivery / Testing:** keep simulator terminal outcomes backend-owned and donor-independent; select timing/mechanism in the delivery plan without promising a user-facing estimate, then prove pending-to-terminal behavior and replay/concurrency invariants independently.

These routes do not reopen settled Product choices. No additional Human authority decision is needed to finish Exploration; unresolved implementation scope should be raised through the Orchestrator before Build.

## Phase handoff

- **Completed:** Stage 3 compared bounded D1 integration, notification fulfillment, status-security implementation, and monetary persistence paths; carried forward approved Product/contract decisions and rejected unsafe or out-of-scope alternatives.
- **Artifacts:** .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/EXP-S2-003-001/evidence/stage-2-gap-analysis.md; .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/EXP-S2-003-001/evidence/stage-3-solutioning.md.
- **Human decision:** none required to complete this Exploration. D1 bounded-scope routing belongs to the Orchestrator/Techplan; Tier-0 writes and Security/PII residual-risk acceptance remain explicit Human gates if they arise.
- **Open / deferred:** concrete D1 seam/ownership; sender and guest-email lifecycle controls; credential and abuse controls/evidence; status-200 cache behavior; amount range/storage precision; simulator timing/mechanism. Approved authority directions remain as stated above.
- **Recommended next step:** Techplan synthesis in a new Techplan Run/Participant, reconstructing from these artifacts and current authority; resolve the D1 scope route before any Build authorization.
- **Session transition:** start a fresh Techplan Participant Session. This was an orchestrated Explorer execution occurrence; Harscode requires Techplan synthesis to be a new Run/Participant with fresh context.
- **Context pointers:** docs/product/mvp-delivery-slices.md §5; approved WU-S2-002/runs/TP-S2-002-015/techplan.md §§2–3, 8, 12–13; docs/project/kencleng-backend-tech-stack.md architecture/domain-boundary sections; backend/AGENTS.md; root AGENTS.md §3; Stage-2 anchors in stage-2-gap-analysis.md Areas 2–6; backend/internal/domain/campaign/{repository.go,repository_db.go,service.go} and migration 000011_create_public_campaigns.

No implementation, spec/API edit, or test was performed in Stage 3.
