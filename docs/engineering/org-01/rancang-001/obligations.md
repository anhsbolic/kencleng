# X1 — Engineering Obligation Map v1

Assignment: `RANCANG-ORG-01-001`. Producer: Requirements Analyst `/root/org01_requirements`, 2026-10-25. Intended consumers: Solution Architect, Experience Specialist, independent assessors, then direct backend/frontend implementation owners through the assembled handoff.

**Standing:** source-derived Engineering translation, produced for review; not new Product/Design authority, selected architecture, independent assurance, Human-accepted handoff or implementation readiness. Verification scenarios below are obligations, **NOT RUN**. Product policy is settled within ORG-01; no Product discovery redo is needed.

## Source and authorization boundary

Kencleng HEAD checked: `3faf3b8a959cdb617958568b158c0805e84107ed`, branch assignment `pilot/station5-c1-space-only`. Harscode HEAD checked: `df21f47ecb85877a571b15f0793ced2f5c33e003`. Approved Rancang definition is exact commit `9663a3fa73ac0101e66764d63af8711f291a66ab`; its embedded candidate labels are superseded by the 2026-10-25 formation approval, not by assumption. The four definition/role blobs below were checked against that exact commit and match the working originals read.

The approved self-contained [Product Contract](../../../product/outcomes/org-01/product-contract.md) and [approval](../../../product/outcomes/org-01/approval.md) own bounded Product meaning. Historical C1 material, examples and existing code establish no additional obligations. Product approval does not authorize implementation. [Assignment state](state.md) records the separate Human authorization, supervised Default mode, scoped documentation destination and accepted R2 deviation. This actor writes only this file; R5/R6 limits remain as recorded there.

Checked source revisions below are **Git blob IDs**, including the then-current uncommitted state snapshot. Paths are relative to their named repository; sources were read directly. No scoped `AGENTS.md` was found under Kencleng `docs/`.

| Ref | Repository / original path | Checked blob |
| --- | --- | --- |
| P | Kencleng `docs/product/outcomes/org-01/product-contract.md` | `da19d8e7cdb9baac8cf4af1c5e6b61763607013e` |
| PA | Kencleng `docs/product/outcomes/org-01/approval.md` | `65211210c102afd02143ab1921b2486fb600b241` |
| K | Kencleng `AGENTS.md` | `14fd68fe8e25306270be6e3e259f181079ea9390` |
| DR | Kencleng `docs/ui-ux/README.md` | `426e0d499c73e28b8ffa8af55353d107369f5afa` |
| DP | Kencleng `docs/ui-ux/product-design-principles.md` | `9eb9e5893fc7c1090d28024c51ed507c18cea945` |
| UX | Kencleng `docs/ui-ux/patterns.md` | `da439232e2d17547a986c8a17914b6858b012a79` |
| BE | Kencleng `docs/project/kencleng-backend-tech-stack.md` | `f451e63af760ad07991b5c072a3530ed9296799f` |
| FE | Kencleng `docs/project/kencleng-frontend-tech-stack.md` | `a1cdd73b816cbeab5509a8729c8f20a3bf012b34` |
| S | Kencleng `docs/engineering/org-01/rancang-001/state.md` (snapshot) | `8960c6cb419538040a77ccdbb23c9dd5598ec93b` |
| HR | Harscode `AGENTS.md` | `7e4d63bc6ef7395480e1edf4f51c01436aef5d4f` |
| HE | Harscode `engineering/AGENTS.md` | `270bd6a5d0c3a8f5c7752000300fc9da64fb695b` |
| RC | Harscode `engineering/spaces/feature-solution-design/contracts.md` | `55d62117c650ad90d6e4c18942181bf5c8db72bc` |
| CO | Harscode `engineering/spaces/feature-solution-design/coordination.md` | `c7ecb97ee7712b1a40e59c4b064d5149b76f7a88` |
| RA | Harscode `engineering/spaces/feature-solution-design/agents/requirements-analyst/agent.md` | `8b60a141be4653609fcae06837b015b1748678ad` |
| RK | Harscode `engineering/spaces/feature-solution-design/agents/requirements-analyst/knowledge.md` | `66e878c753ffaacaa6304481964f422186c1eb67` |
| FA | Harscode `engineering/spaces/feature-solution-design/formation-approval-2026-10-25.md` | `120baaaee6dd404b5bfa4cfe79960f6b36e2bf21` |

## Complete Product coverage and checkable derivation

All identifiers in this table have prefix `ORG-01`; abbreviated REQ/AC/SC numbers refer to P §§1, 3.2, 4. Normative behavior is approved Product; consumer implications and verification prompts are derived Engineering obligations. No row selects an endpoint, persistence schema or screen count.

| X1 obligation | Binding trace | Checkable behavior and negative boundary | Required consumer consequence |
| --- | --- | --- | --- |
| X1-01 Start without false authority | REQ-01; AC-01; SC-04 | Actor can begin without claiming completed legal-representation verification. Starting alone proves neither successful formation nor Owner status. External legal verification is not a universal formation prerequisite. | O2/O3 distinguish starting from effective success and distinguish internal role from external authority; do not add new actor eligibility policy. |
| X1-02 Effective actor attribution | REQ-02; AC-02; SC-01 | At effective success, both the action and first Owner relationship are genuinely attributable to the acting person. Missing/invalid attribution cannot yield effective success. Identity need not be complete at the first interaction. | O2 must define trusted attribution and effective-time enforcement; FE assertions alone cannot confer ownership. Check identity loss/substitution and unauthorized requests against the chosen mechanism. |
| X1-03 Understanding before effect | REQ-03; AC-03; SC-02 | Before the consequential effect, actor can understand which organization is formed, that they become first internal Owner, and what that role does not prove externally. | O3 must place consequential meaning before effect and reconcile its action boundary with O2. A post-success explanation alone is insufficient. No mandatory separate confirmation screen is inferred. |
| X1-04 One truthful combined success | REQ-04; AC-04; SC-01 | Success requires valid effective organization and valid attributable first internal Owner together. Organization-only or Owner-only state cannot be ordinary full success. | O2 specifies the joint guarantee, observable result and what consumers may conclude. Verification must exercise failure between material operations and ensure partial/unknown outcomes do not become success. Product does not prescribe a transaction recipe. |
| X1-05 Durable current inspection | REQ-05; AC-05; SC-02 | After true success, a separate inspection from the transient success message shows the actual current organization/Owner relationship and internal meaning. | O2 provides authoritative current-state semantics and access boundaries; O3 distinguishes loading, unavailable/stale information and verified current state. Local optimistic state or toast alone cannot establish durability. |
| X1-06 Claims retain provenance | REQ-06; AC-06; SC-04 | Organization-provided information remains recognizable by its source; formation does not turn claims into independent verification, Reviewer approval, or Campaign curation. | O2/O3 preserve provenance and avoid manufactured status/badge implications. Check stored/read representations and copy, not only the success message. No external evidence subsystem is implied. |
| X1-07 Honest negative and unknown outcomes | REQ-07; AC-07; SC-03 | Failure, partial effect, insufficient attribution, interruption and uncertain result are not ordinary success. Any offered retry/recovery cannot manufacture success from partial or unknown state. | O2 defines outcomes and safe action semantics for actual guarantees; O3 reflects them without optimistic assumptions. Check lost response, permission/identity failure and subsequent inspection. A recovery subsystem is not a Product prerequisite. |
| X1-08 Separate contexts with similar claims | REQ-08; AC-08; SC-03 | A and B each form organization K with distinct internal IDs; when each meets ORG-01 conditions each has effective internal Owner. Similar names/claims or Potential Conflicts alone cannot auto-block, defer, reject, merge, close or cancel them. Reviewer approval remains separate. | O2 must preserve distinct identities and avoid uniqueness/conflict rules that contradict this acceptance. O3 must not imply an automatic winner or approval. No universal conflict detector/adjudicator is required. |
| X1-09 No cross-Journey side-effect claim | REQ-06; AC-09; SC-04 | Completion alone does not perform Campaign preparation, publication, donation eligibility or organization review. | O2 describes only authorized ORG-01 effects; O3 next actions and success language must not imply those separate outcomes already happened. |

SC coverage is conjunctive: SC-01 → X1-02/04; SC-02 → X1-03/05; SC-03 → X1-07/08; SC-04 → X1-01/06/09. Passing one path cannot substitute for the remaining material conditions. Positive spot-check: effective attributed organization/Owner plus pre-effect understanding and subsequent inspection satisfies the combined definition only when its claims remain bounded. Negative/trust spot-check: an apparently successful response with no attributable Owner violates AC-02/04/07; two similarly named contexts are not that failure and must satisfy AC-08 independently.

## Applicable standards and verification implications

| Source | Material obligation for O1–O4 and downstream owners |
| --- | --- |
| K golden rules/security; BE engineering rules | Backend owns enforcement when authorization is introduced; frontend gates are UX. Parameterize SQL; handle expected failures; do not log secrets/raw tokens/unnecessary PII or expose SQL, stack traces, paths or secrets in client failures. Security choices require current risk evidence, not copied historical mechanisms. Proposed identity/Owner control needs independent security scrutiny. |
| DP §§2, 4–6, 9, 11–15 | Do not imply unsupported verification; distinguish platform fact, organization claim and unknown state. Explain consequences before commitment, keep material uncertainty reachable, preserve task/consequence priority responsively. Ordinary presentation choices are delegated; material new interaction/brand choices need the owning Design decision if actually encountered. High-fidelity UI is not an automatic prerequisite. |
| UX §§3–4, 7, 10–14, 16, 18–19 | Server validation is authoritative; distinguish field validation from request failure; prevent accidental duplicate non-idempotent submission and communicate progress. Preserve input unless successful transition intentionally leaves it. Recovery only when meaningful; status/success explains consequences rather than relying on color or badge. Sensitive lookup distinctions respect anti-enumeration needs. |
| BE neutral baseline/API relationship; FE state/API boundaries | Available tools do not prove delivered authentication/domain/API. Server state remains authoritative; client/mock state is not production evidence. Use active generated contract types when a canonical contract exists; O2 is a proposal, not that adoption. Keep the least additional complexity sufficient for accepted behavior and security. |

For O4, every X1 row and material standard needs a named accountable producer and an intended execution/verification owner in the final bundle. Architect supplies technical scenario detail; Experience supplies observable experience detail. These are coordination requests, not unilateral assignments. Coordinator confirms ownership. Evidence must distinguish document inspection, independent review and executed tests. No runtime tests, browser checks or independent assessments were performed by X1's producer.

## Freedoms, exclusions and real gaps

**Delegated freedoms (P §5):** authentication mechanism and timing consistent with effective attribution, data model/persistence, consistency technique, API/component shape, proportionate recovery and testing are Engineering decisions. Form fields, interaction and copy may be derived from accepted Product/Design authority without inventing business rules. Simplicity must preserve mandatory safeguards. X1 makes no final mechanism, endpoint, role-permission matrix or UI architecture selection.

**Exclusions:** full Owner/Staff permission matrix; staff transfer/delegation; external legal verification; Reviewer procedures, standards, corrections or conflict adjudication; universal duplicate detection; auto-merge/closure; Campaign; publication/donation; financial movement/accountability; mandatory recovery subsystem. Existing context does not automatically require repeating ORG-01. Formation success is no proof of legal authority, Reviewer approval, organizational quality or downstream readiness.

| Gap classification | Actual issue / impacted obligation | Disposition |
| --- | --- | --- |
| (a) Invalid/incomplete bounded input | None identified in P/PA for the scoped semantic derivation. | Do not relitigate settled Product decisions. This is an analyst conclusion, not independent input approval. |
| (b) Material Product/Design owner decision | None identified from inspected originals. Additional eligibility rules or new material interaction meaning would exceed this derivation if introduced. | Hold only a genuinely dependent edge if new evidence exposes such a gap; route exact contradiction to Coordinator/Human. Do not preemptively demand a new Product decision. |
| (c) Delegated Engineering questions | Trusted actor attribution and effective Owner enforcement (X1-02); joint success and interrupted/duplicate/unknown handling (04/07); authorized durable current inspection (05); provenance representation (06); distinct IDs without name-based blocking (08). | Architect investigates and decides substantively in O1/O2, with security assessment and O3 reconciliation. Unknown mechanisms are not invalid input or Product blockers. |
| (d) Relevant evidence deficiency | Baseline architecture explicitly supplies no delivered authentication/product domain/API; this actor has not inspected runtime implementation or proved current technical readiness. Intended backend/frontend consumer acceptance and project-bound compatibility have not yet been checked against completed outputs. | Architect checks actual technical evidence; Coordinator checks actual per-consumer output compatibility. Generic reusable consumer Input Contract is NOT APPLICABLE for direct implementation owners; project-bound fit remains required. No standard deficiency needing a standards change is established by this X1. |

Downstream gating remains explicit: this assignment authorizes documentation only. Canonical API adoption, protected implementation, services and deployment need their own authorization. Human acceptance of the exact assembled handoff remains pending. Source/output changes require affected trace/review joins to be rechecked; this file cannot certify the final bundle by itself.
