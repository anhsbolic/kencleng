# Review findings — WU-S2-002

> Phase             : Independent Techplan Review
> Work Unit         : WU-S2-002
> Run               : RV-S2-002-004
> Author            : Codex Reviewer
> Participant       : Codex Reviewer
> Session           : Fresh independent Reviewer Session; session ID unavailable
> Created           : 2026-09-27
> Model             : `gpt-6-luna`
> Reasoning         : `high`
> Target revision   : `10457b17059e2da3d97a4f76e4c3fae127227d73`
> Workflow revision : `b122a75d494250d04eb93e71f4c391e82c847842`

**Gate:** Complex — The plan crosses Campaign/Donation and API boundaries and carries payment/money, PII, guest-credential, and settlement correctness risks. The payment/PII criteria independently warrant review; the plan has 10 Rules & Validation entries, below the 15-entry criterion.

**Sections resolved:** Background = §1; Scope = §2; Requirements = §3; Rules & Validation = §4; Decision Log = §5; Backward Compatibility = §6; Edge Cases & Risks = §7; Interface Contract = §8; Architecture / Plan = §9; Implementation Details = §10; Files Changed / Files NOT Changed = §11; Testing Checklist + Test Focus Pointer = §12; Open Items = §13.

### Blocking

- none.

### Non-blocking

- **RISK-7 and RISK-10 cross-reference the wrong rule for atomic funding correctness** — §7, rows RISK-7 and RISK-10 (`techplan.md:114,117`). RISK-7 says threshold ordering must hold “while R4 holds,” and RISK-10 cites “R4” for terminal status/funding divergence, forgery, duplicate increments, and partial commits. §4 defines R4 as simulator-owned state/recovery; §4 R3 defines atomic success/funding coupling, exact-once reflection, replay, and concurrent-update correctness. Correct the references to R3 (and, for RISK-7, retain R7 for threshold/eligibility). This is a mechanical, non-blocking correction because the applicable rule is unambiguous and the checklist/contract elsewhere preserve the required evidence.

### Clean

- **Rule fidelity and checklist coverage:** R1–R10 are each represented in §12. R3 preserves atomic Donation `success`/funding coupling as one committed business outcome, exact-once funding including replay and concurrent settlements, and no lost increment; Testing owns downstream failure, replay, and concurrency evidence. R8 separately covers request submission idempotency and deliberate fresh-key behavior. R4 and R5 carry simulator and guest email/PII evidence; R6/R7/R9/R10 have owner and/or Testing coverage appropriate to their contract, security, compatibility, threshold, API, and design concerns.
- **Decision fidelity:** D1–D14 and the scope/rules preserve the Exploration direction and the approved OIR decisions. Product decisions for O1–O6 and O9 are carried into the plan with their consequences; remaining representation, simulator timing, Security/PII controls, API parity, Design review, and conditional consumer audit remain open. The threshold rule limits Campaign coordination to eligibility/crossing/accepted pending/rejection after closure and does not add Slice 3 lifecycle. Same-key submission retry remains distinct from settlement replay.
- **Open Items lifecycle:** §13 separates active delivery/owner verification from resolved product decisions. Resolved O1–O6/O9 retain their actual outcomes and consequences; O7 remains Design review, O8 remains conditional on a proposed historical operation removal/replacement, and no prior unresolved concern is presented as accepted Security/PII risk.
- **Test Focus Pointer:** The five surviving sensitive areas cover settlement/concurrency, submission retries, guest status credential, guest-email PII, and Campaign eligibility/threshold. Their Stage 2/Stage 3 artifact and heading anchors match the durable Exploration headings; each is marked relevant with a reason. Specialized Testing ownership is limited to sensitive evidence beyond ordinary contract/spec checks.
- **Technical facts and guardrails:** At `TARGET_REVISION`, `backend/cmd/server/main.go` registers public Campaign routes but no Donation route; `backend/internal/domain/campaign/service.go` `toPublicDetail` sets donation availability to `unavailable`; the public transport projects that value. `api/README.md` confirms split domain files are authored, `index.yaml` changes with path changes, and bundle/generated frontend types are derived and must not be hand-edited. These support the plan's implementation and API-workflow claims. Tier-0 write fences are retained in scope and implementation notes.
- **Diagram:** No diagram is present; diagram validation is not applicable.
- Review used artifact and source inspection; no tests or API validator were run.

## Phase handoff

- Completed: independent Complex gate and review of the current amended Techplan against Exploration, Product/MVP, OIR, canonical Techplan rules, and target-repo facts.
- Artifacts: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-004/review-findings.md`
- Human decision: none required before the Planner's resolution pass.
- Open / deferred: the two §7 risk cross-references above; Techplan Open Items remain as recorded in §13.
- Recommended next step: one Planner resolution pass to correct the references, then regenerate the human report as required and proceed to the Human Techplan approval gate. This mechanical correction does not require re-review unless that pass changes material meaning.
- Session transition: return to a Planner Session for the resolution pass; use a fresh Build Session after Human approval, if authorized.
- Context pointers: amended Techplan §4 R3/R4, §7 RISK-7/RISK-10, §12; OIR `OIR-S2-002-001/resolution-brief.md` O6 and R3 references; Product authority `docs/product/mvp-delivery-slices.md` §5.
