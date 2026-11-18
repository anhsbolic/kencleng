# PO-01 — Organization Entry & Initial Eligibility Decision

> **Standing:** SELECTED PRODUCT OUTCOME / WORKING DEFINITION OF DONE — NOT APPROVED AS BINDING PRODUCT CONTRACT  
> **Selection:** Anhar, Product working session, 17 November 2026  
> **Product decision-maker:** Anhar  
> **Owning concern:** Kencleng Product (Product Harscode phase)  
> **Purpose:** Durable bounded outcome and completion checklist for Product Definition; not an MVP/release plan, UX deliverable, implementation specification, or Engineering authorization.  
> **Interpretation:** Selection and permission to checkpoint are confirmed. Detailed Product semantics and the DoD checklist below remain candidate criteria until substantively reviewed/accepted at exact revision. Checking a box requires supporting accepted evidence, not just authoring text.

## 1. Label and intended Product outcome

**Local label:** `PO-01` = first selected Kencleng **Product Outcome**. This is a navigation/working identifier, not a mandatory Harscode identifier, feature ID, runtime status, or implementation sequence.

**Name:** Organization Entry & Initial Eligibility Decision.

**Selected bounded outcome:** Perwakilan organisasi nirlaba dalam cakupan Kencleng dapat memulai dan mengajukan permohonan kelayakan, lalu mendapatkan keputusan **awal** yang jelas, sah, berbasis bukti, dapat dipertanggungjawabkan, dan memiliki konsekuensi kewenangan/hak yang tidak ambigu.

**Intended value:** organisasi memahami basis dan hasil keputusan; Kencleng dapat memberi atau menahan pengakuan eligibility tanpa klaim verifikasi palsu; pihak downstream memperoleh aturan yang cukup jelas untuk membedakan organisasi yang mengklaim identitas, pemohon yang memiliki kewenangan, dan organisasi yang telah dinilai eligible.

**Working eligible entity scope:** Yayasan dan Perkumpulan berbadan hukum Indonesia dengan identitas hukum sendiri. Dimulai dengan representative case **Yayasan**, lalu memeriksa perbedaan material pada **Perkumpulan**. Jenis dokumen dan persyaratan hukum terperinci **belum diverifikasi atau disetujui**.

**Scope boundary** — IN:
- Memulai permohonan sebelum semua bukti lengkap, melengkapinya, dan mengajukan review formal.
- Identitas hukum organisasi, karakter nirlaba, dasar kewenangan pemohon, dan satu identitas Kencleng per entitas hukum.
- Minimum informasi yang diperlukan untuk start, submit, dan keputusan; evidence provenance dan sufficiency.
- Mandat yang cukup untuk First Owner; tanpa mengasumsikan pemohon pertama berhak mengendalikan organisasi.
- Wewenang peninjau/pengambil keputusan dan conflict-of-interest safeguards dalam Product meaning.
- Hasil awal: eligible, rejected, belum dapat diputuskan karena bukti tidak cukup/bertentangan; makna hak, kewajiban, kejelasan hasil, dan jalan klarifikasi/koreksi yang material.
- Kewajiban accountability organisasi yang harus diterima secara sah sebelum eligible.

**OUT / adjacent boundary (kecuali dependency material terhadap keputusan awal):**
- Ongoing eligibility monitoring, reassessment berkala, suspension/revocation/recovery lifecycle secara penuh.
- Owner delegation lanjutan, Staff management, serta penyelesaian semua kemungkinan sengketa organisasi.
- Campaign creation/curation/publication, donor conversion, donation eligibility, pembayaran/settlement/disbursement/refund, dan verifikasi penggunaan dana.
- UX journey, layar, UI components/copy final, API, database, security implementation, rencana pengembangan, atau klaim runtime.

**Cross-domain invariant:** Organization Eligibility bukan otomatis Campaign approval, donation eligibility, atau penggunaan dana yang terverifikasi. Kondisi initial decision tidak boleh mengklaim lebih dari fakta yang benar-benar diketahui dan diperiksa.

## 2. Authority and provenance

**Effective whole-product foundation:** [Kencleng Product Intent](https://github.com/anhsbolic/kencleng/blob/7778d66d91ba81fff08a6ae1ee0aac9ef71905f8/docs/product/product-intent.md), original blob `cc98e46dd1d8d9e32b8718eb3b95ad9a6ffc2312`, approved 22 Oktober 2026 within its stated whole-product scope.

**Prior working-domain context (not binding):** [Organization Eligibility — Working Product Definition](../../organization-eligibility-working.md), pre-selection checkpoint, blob `e2f65984bb39a3d6d0b6b2d31c6240180be22964`. Its broader ongoing-eligibility decisions are context and boundary constraints, not automatically in-scope delivery requirements for PO-01.

**Research guidance, not Kencleng authority:** Harscode [Product Definition](https://github.com/anhsbolic/harscode-workspace/blob/302bedae43563a66d2ef914191c18139440c0c30/product/product-definition.md), [Consumer Handoff](https://github.com/anhsbolic/harscode-workspace/blob/302bedae43563a66d2ef914191c18139440c0c30/product/consumer-handoff.md), and [Authority & Approval](https://github.com/anhsbolic/harscode-workspace/blob/302bedae43563a66d2ef914191c18139440c0c30/product/authority-and-approval.md); selected research guidance, still pending its own human exact-revision review.

**Other owning inputs:** [Design Authority](../../../design/README.md), [Engineering Standards](../../../engineering/standards/README.md), and applicable legal/privacy/security authorities when relevant. These are separately owned sources; their factual or operational approval is not supplied by this document.

A future binding Product Contract, if selected, must express the material approved rules directly. The original approved content and a separate approval reference must identify exact immutable bytes/scope. This checkpoint does **not** itself grant that approval.

## 3. Definition of Done — checklist

**Checklist convention:** `[x]` = supported as completed **for this checkpoint's limited meaning**; `[ ]` = pending. Do not mark a consequential requirement complete merely because a proposed rule is written, a prototype looks convincing, or Harscode guidance lists it. Where a criterion is intentionally non-applicable, explain why instead of treating it silently as passed.

### O1 — Value intent and bounded outcome

- [x] **O1.1 — Outcome selected by Human owner:** PO-01 named and bounded as the Product Definition focus; this is not an approved release/MVP scope.
- [x] **O1.2 — Initial IN/OUT boundaries captured:** Scope, adjacent domains, and representative Yayasan/Perkumpulan cases recorded above for review.
- [ ] **O1.3 — Actor value and trade-offs validated proportionally:** Identify intended beneficiaries and affected actors (pemohon, legitimate representative, organisasi, peninjau/pengambil keputusan, affected downstream users), real problem, rationale, competing interests, and evidence confidence without inventing demand.
- [ ] **O1.4 — Product success conditions set:** Observable, falsifiable conditions for a legitimate initial decision and a non-misleading outcome; distinguish decision correctness from conversion/growth or implementation completion.

### O2 — Legitimate Product commitments and consequential meaning

- [ ] **O2.1 — Eligibility basis sufficient:** Define what must be true for initial eligibility for Yayasan and Perkumpulan, distinguishing actual law/domain constraints from Kencleng policy.
- [ ] **O2.2 — Minimum information stages clear:** Define what is necessary to start, submit for formal assessment, and decide; permissible absence, correction, and material privacy/data purposes.
- [ ] **O2.3 — Evidence sufficiency and provenance clear:** What claim each evidence item can support; source, freshness, examiner, verification limitations, and criteria for sufficient, insufficient, conflicting, unverifiable evidence.
- [ ] **O2.4 — Representative and First Owner authority clear:** Mandate needed to apply versus accept obligations versus hold First Owner authority; duplicate applicant/organization and materially conflicting mandate cases.
- [ ] **O2.5 — Legitimate decision authority clear:** Who may review and decide initial eligibility/rejection, under what constraints, and how material conflict of interest is handled; do not invent an implementation permission matrix.
- [ ] **O2.6 — Accountability obligation at eligibility clear:** Minimum organizational commitments, authorized acceptance, and which Campaign-specific responsibilities remain for later decisions.
- [ ] **O2.7 — Outcomes and rights clear:** Consequences of approval, rejection, pending/unknown, missing or conflicting evidence, and corrections/resubmission. Specify what is *not* authorized by initial eligibility.
- [ ] **O2.8 — Material invariants coherent:** One recognized organization identity per legal entity; application is not authorization; assessed fact is not submitted claim; Organization Eligibility is not Campaign or donation eligibility.
- [ ] **O2.9 — Observable acceptance specified:** SC/REQ/AC or equivalent for positive, negative, partial, unknown, authority-conflict and misuse cases. Acceptance claims are not proof that tests or implementations have run.

### O3 — Legitimate authority, constraints, and acceptance

- [ ] **O3.1 — Project/external authority reconciled:** Identify effective exact source revisions, applicable legal/privacy/security constraints, and decisions that require external/domain validation; do not assert unverified Indonesian regulatory details as settled.
- [ ] **O3.2 — Decision ownership documented:** Human Product owner, scope of any delegation, responsibility for review/approval and resolution of consequential ambiguities.
- [ ] **O3.3 — Explicit bounded Product approval:** Anhar approves exact substantive Product content and scope; original accepted source and external approval evidence/reference identify exact immutable revision. Working-chat agreement or this checklist alone is not sufficient.
- [ ] **O3.4 — Exclusions and freedoms explicit:** Clarify what UX, UI, and Engineering may decide independently, and what consequential Product questions must return to Product.

### O4 — Uncertainty, evidence standing, and continuity

- [ ] **O4.1 — Evidence standing inspectable:** Accepted vs working vs hypothesized vs observed vs unknown, including source/date and uncertainty for all material Product claims.
- [ ] **O4.2 — Open questions bounded:** For each material unresolved issue, record impacted rule/actor, owner, needed evidence or decision, and whether it blocks a downstream decision; no hidden assumption.
- [ ] **O4.3 — Whole outcome reconstructable:** A fresh reader can locate the current accepted original Product source, approval/revision, scope, constraints, unresolved limits, and supersession without this conversation.

### Consumer-specific Product information readiness (not doing their jobs)

- [ ] **D.1 — UX/Interaction & Experience:** Consumer can explore/design the relevant actor behavior using legitimate Product goals, rights, allowed/denied/unknown results, and constraints **without inventing Product policy**. UX owns actual journeys/flows/navigation.
- [ ] **D.2 — UI/Content:** Required claims, source attribution, disclosure, decision meaning, and prohibited false-success labels are unambiguous. UI/Design own expression and visual solutions.
- [ ] **D.3 — Engineering Solution & Development:** Accepted invariants, actor authority, domain outcomes, acceptance conditions, constraints, dependencies, and meaningful freedoms are usable without inventing consequential semantics. Technical solution/security work and execution authorization remain separate.
- [ ] **D.4 — Actual sufficiency check:** Inspect at least relevant project-specific downstream consumer needs; record `SUFFICIENT / CONDITIONAL / BLOCKED / UNVERIFIED` by receiving responsibility and evidence. Do not claim independent consumer sign-off unless performed.

### Checkpoint provenance (not Product DoD)

- [x] **C.1 — Earlier working discovery is durable:** Wider Organization Eligibility working checkpoint exists in Kencleng repo and is distinguishable from binding Product authority.
- [x] **C.2 — PO-01 selected for Product Definition:** Explicit session agreement, not a Product behavior acceptance.
- [ ] **C.3 — Actual approved PO-01 Product contract/revision recorded:** Not yet applicable as an accomplished checkpoint; will require separate exact-content approval.

## 4. Completion rule and current standing

**PO-01 Product Definition is DONE only when:** material O1–O4 and relevant consumer-sufficiency criteria above have supporting accepted information; any unfulfilled criterion is demonstrably non-blocking for the intended receiver/decision or explicitly resolved; substantive accepted Product meaning is self-contained in its actual owning source(s) and references exact effective authority and approved immutable revision. No false claim of consumer verification, legal clearance, Engineering readiness, or deployed capability.

**Current outcome standing:** **SELECTED FOR PRODUCT DEFINITION, NOT PRODUCT DONE**. This checklist is the working completion contract for planning/review, **not yet an approved binding Product Contract**.

**Open decisions requiring near-term Product work:** Eligibility Decision Basis; required evidence and sufficiency threshold for Yayasan/Perkumpulan; legitimate review/approval authority; First Owner mandate; organization accountability obligations; result semantics when evidence is insufficient or contradictory; and consequential rights after the initial decision.

**Next inquiry:** Use Yayasan representative case and test `sufficient / insufficient / contradictory` evidence against authority, decision outcome, rights, and falsifiable acceptance. Then test differences material to Perkumpulan. Legal facts must be verified from owning sources before they are treated as binding constraints.

**Out-of-scope retained:** Ongoing monitoring, suspension, revocation, and recovery remain relevant boundary context in the broader working checkpoint, not a promise to fully define or implement their lifecycles within PO-01.

**Non-authorizations:** This document does not approve Harscode reusable policy, design deliverables, Engineering solutions, protected writes, API adoption, implementation, transactions, deployment, or real-world evidence claims.

## 5. Checkpoint #2 — agreed working direction (18 November 2026)

> **Standing:** HUMAN-AUTHORIZED WORKING CHECKPOINT ONLY. Anhar approved persisting the discussion's agreed working decisions on 18 November 2026, **not** these exact bytes as a binding PO-01 Product Contract, legal sign-off, Engineering handoff, or implementation authorization.
> **Scope:** Selected PO-01, representative Yayasan scenario; differences for Perkumpulan still to be tested. The shared [Actor & Authority Model](../../actor-authority-model.md) supplies reusable Product meaning. The wider [Organization Eligibility working checkpoint](../../organization-eligibility-working.md) remains historical/adjacent discovery and is not a second competing authority.
> **Status:** DoD in §3 remains uncompleted unless evidence and exact-revision acceptance independently support completion. Working agreements below are retained to prevent a fresh-chat durability gap.

### 5.1 Actor outcomes, invariants, and initial eligibility gates

- An Applicant with a valid authenticated Kencleng account can initiate and manage a bounded application, even before holding a proven organizational mandate. Application submission, account existence, and applicant priority never grant Owner, Staff, or organizational control.
- One underlying legal entity should have one recognized organization identity. Multiple applicants and materially conflicting authority claims must not cause automatic takeover or unauthorized exposure of others' evidence.
- Initial eligibility requires **all three substantive gates**: (A) organization legal/nonprofit qualification and applicable Kencleng criteria evidenced sufficiently; (B) organizational baseline accountability obligations validly accepted by a sufficiently authorized party; and (C) adequately mandated First Owner with a valid Kencleng account, explicit acceptance, and Kencleng recognition. All require a legitimate Kencleng decision.
- The First Owner can be a non-chair/non-management individual where a truly authorized organizational party has sufficient power to delegate that scope. Evidence of authority to file an application alone is insufficient. No implicit acceptance or right assignment.
- Organization qualification, applicant relationship, representative authority, Owner rights, and organization eligibility remain independently meaningful; a partial positive finding is not final eligibility.
- Eligibility does **not** imply Campaign approval/publication, permission to receive donations, fund settlement, verified use of funds, or real-world impact.

### 5.2 Start, submission, and approval: different thresholds

| Stage | Minimum Product meaning | Not implied |
| --- | --- | --- |
| Start/draft | Valid account; claimed organization name, applicant contact/identity needed to continue, and self-declared relationship | Proven entity identity, organizational mandate, or reviewable evidence |
| Formal submission | Reviewable identification of the claimed Yayasan/legal entity; document or inspectable official/legal reference relevant to status; applicant attribution; initial identification of claimed authorized representative and basis; applicant acknowledgement of submission's limited meaning | Completed verification, accepted organizational accountability obligations, First Owner acceptance, eligibility, or organizational control |
| Initial eligibility approval | Manual authorized Kencleng decision grounded in sufficient checked evidence for legal identity, Kencleng criteria/nonprofit qualification, representative mandate, valid organizational obligations acceptance, and First Owner authorization/acceptance | Any downstream Campaign or monetary privilege |

Evidence minimum is **source-neutral**: an inspectable official reference may be sufficient for submission without mandatory upload; a submitted reference is not automatically verified. One evidence item may support multiple claims. A formal submission asks Kencleng to assess claims; it does not itself bind the legal entity.

**Reconciliation of earlier candidate:** §5 of the wider 16 November Organization Eligibility working note listed accepted obligations among *candidate formal-submission requirements*. For the bounded PO-01 working direction, **authorized organizational acceptance and First Owner acceptance are approval prerequisites, not automatic submission prerequisites**. This later scoped decision supersedes that earlier candidate solely for PO-01; it does not rewrite historical evidence or finalize a platform-wide rule.

### 5.3 Evidence domains and assessment outcomes

| Evidence domain | Material question to substantiate for approval |
| --- | --- |
| EB-01 — Legal identity | Is the specific Yayasan legally identifiable and sufficiently established as a legal entity? |
| EB-02 — Kencleng eligibility | Does the organization meet applicable nonprofit/domain and Kencleng eligibility criteria? |
| EB-03 — Representative authority | Does the person/org binding the entity have valid authority for that act? |
| EB-04 — First Owner authorization | Are appointing authority, mandate scope/current validity, candidate identity, explicit acceptance, and Kencleng recognition sufficiently supported? |
| EB-05 — Organizational commitment | Were the relevant accountability obligations validly accepted for the legal entity, with traceable actor/authority/version/time? |

**Interpretation of EB-01..05:** these are **Reviewer-facing material review criteria**, not five mandated machine-verification engines, document parsers, separately persisted claim objects or independent screens. An authorized human reads the relevant sources, determines whether the approval requirements are met, and **attests to that conclusion**. Findings such as sufficient, insufficient, unverifiable, contradictory or disqualifying are helpful human assessment language, **not a required per-document classification taxonomy or technical status enumeration**. A routine positive review can use a concise attestation against the applicable eligibility requirements and linked evidence; material exceptions, rejection, uncertainty, conflicts and overrides need adequate reason to explain the action. Evidence sufficiency and legal interpretation remain the reviewer's/authorized decision-maker's responsibility.

**Minimum Product record:** Link the consequential action to the application/organization, attributable actor(s) and relevant capacity, time, event/action, affected person or entity, result, and concise material context. For reviewer attestation and decisions, include **identifiable evidence references** sufficient to trace what was relied upon, the applicable criteria/commitment version where material, and the reason where consequential or exceptional. Preserve material event chronology and corrections; a later change must not silently overwrite an earlier decision. This is not a mandate to capture every document read, transcribe legal documents, produce a finding for every sentence, or store every raw input forever.

### 5.4 Human-led decision authority and minimal traceable audit

- Kencleng uses **manual human assessment and approval** for initial eligibility. Evidence submission/completeness and automated auxiliary checks cannot grant eligibility. Systems do not have to understand, interpret, or independently verify legal-document contents.
- Reviewer and Decision-Maker are separate **responsibilities**, not necessarily separate people: the same appropriately authorized and non-conflicted person may perform both in ordinary cases.
- Disqualifying conflict of interest prevents that actor's review; credible material mandate conflict or takeover risk holds the **affected** authority decision and requires appropriately authorized independent handling. Assessment and escalation are accountable human judgments, not automatic document-credibility scores.
- **Minimum business audit-event meaning:** (1) **what** action/event, (2) **when** it occurred, (3) **where** in Kencleng's domain it occurred (application/organization/affected resource), (4) **who** acted and, where relevant, in which authorized capacity, (5) **for whom/what** it was done, (6) **result/effective consequence**, and (7) **necessary context** (for decisions/attestations, reference to relevant evidence and material rationale). These are *questions that must be answerable*, **not a prescribed database schema, one-row-per-document design, or exhaustive logging specification**.
- **Scope of business events:** Capture meaningful formal submission; material evidence addition/replacement/withdrawal; organizational/First Owner acceptance; Reviewer attestation; consequential approval/rejection/closure/withdrawal; effective rights/authority changes within PO-01; material correction/escalation. Do not require recording every field edit, page view, document-reading step, or incidental interaction as a business audit event.
- **Security logging is a separate Engineering responsibility:** Follow current Kencleng authorization/privacy/security standards and risk-appropriate [OWASP Logging guidance](https://cheatsheetseries.owasp.org/cheatsheets/Logging_Cheat_Sheet.html). Business audit's minimalism does **not** waive security-relevant logging (e.g., access-control failures or sensitive privilege operations). IP address, device, HTTP metadata and location are **not universal Product audit fields**; collect technical details only where security purpose, risk and lawful processing justify them. Protect integrity and access to both kinds of record.
- Keep traceability proportionate and purpose-bound; actual retention/deletion terms and lawful basis remain OPEN privacy/legal decisions. Product owns *what must be accountable*; Engineering chooses secure audit, event and logging mechanisms. Neither an OWASP citation nor a Product checkpoint proves implementation or compliance.

### 5.5 Distinct application outcomes and rights

| Resolution | Required Product meaning | Consequence |
| --- | --- | --- |
| **Approved** | All eligibility gates supported; authorized human decision recorded | Initial organization eligibility can become effective; recognized First Owner rights only within valid mandate/defined scope |
| **Rejected** | Material eligibility criterion demonstrably not met; substantive, reasoned authorized decision | No initial eligibility or Owner privileges from this application; rejection does not automatically allege fraud |
| **Closed Inconclusive** | Evidence remains unestablished/unverifiable following reasonable notice and opportunity to clarify; reasoned process closure | No eligibility or Owner grant; does **not** assert that the organization has been proven unlawful or disqualified |
| **Withdrawn** | Valid applicant-initiated withdrawal in an allowed context; relevant consequential actions and authority still require checking | No eligibility from withdrawal; existing history/evidence obligations are not silently erased |

Clarification Required and Authority Conflict are blocking **process/evidence conditions**, not automatic final rejections. Closing an unresolved application is distinct from finding a proven negative: never keep a request open forever by default, nor invent an arbitrary closure timeout. A fresh assessment may be sought on materially changed facts/new evidence subject to legitimate restrictions; refiling does not itself restore/grant rights. Applicant and authorized actors receive truthful, access-bounded reasons and permitted next steps; third-party confidential application evidence is not automatically exposed.

### 5.6 Baseline organizational accountability obligations

Before initial eligibility, a legitimately authorized representative must accept the applicable, identifiable **version** of the organization's obligations; acceptance must be distinct from Applicant terms and the First Owner's personal role acceptance.

| Candidate ID | Minimum organization-level commitment |
| --- | --- |
| OA-01 — Truthful Information | Supply truthful material information and correct material inaccuracies when known |
| OA-02 — Valid Representation | Ensure organizational acts and appointments are within valid authority and mandate scope |
| OA-03 — Material Change Disclosure | Proactively disclose material changes affecting organization eligibility or representative authority, even without an active Campaign |
| OA-04 — Assessment & Oversight Cooperation | Cooperate reasonably with legitimate, proportionate Kencleng assessments and follow-up; no unrestricted request for unrelated personal or organizational records |
| OA-05 — Continuing Accountability | Commit to applicable Campaign-specific reporting, funds-use, and post-fundraising accountability **when relevant**; details are owned by future Campaign/fund accountability decisions |

Acceptance traceability includes legal entity, acting person and authority, accepted terms/version, meaningful act/evidence, time, and Kencleng assessment. Future substantial changes to obligations cannot silently inherit blanket acceptance. Specific Campaign reporting cadence, financial operations, fund verification, sanctions, monitoring/reassessment and suspension mechanics remain adjacent, not PO-01 specifications. A suspected breach is not an automatic revocation or proven fraud claim.

### 5.7 Representative behavior tests to carry into O2.8/O2.9

| Scenario | Expected Product invariant |
| --- | --- |
| Applicant starts draft with no proven mandate | Draft may exist; no organization rights or false verified claim |
| Applicant submits reviewable source-neutral legal reference | Formal review may begin; provenance and verification remain separate |
| Legal identity is sufficient, but First Owner mandate or acceptance is pending | Partial finding remains truthful; no effective initial eligibility or First Owner activation |
| Two materially conflicting First Owner mandates | No first-filer privilege; affected authority decision held; independent authorized handling and bounded disclosure |
| Routine complete case with one authorized unconflicted reviewer/decision-maker | Human decision can be legitimate with accountable evidence, rationale and chronology; two-person approval is not universally required |
| Applicant fails to clarify unverifiable evidence after fair opportunity | Closure can be inconclusive without substantive legal rejection; reason/history preserved |
| Entity definitively fails a material eligibility criterion | Substantive rejection supported with evidence and authorized reason; no fraud allegation by default |
| Eligible organization requests Campaign publication or receives an attempted donation | No implied Campaign/donation/funds approval from organization eligibility alone |

### 5.8 OPEN / not yet ready for binding Product handoff

- **OPEN PRODUCT DECISION:** Specific eligibility criteria and proportionate, legally sufficient sources/document requirements for Yayasan and Perkumpulan (including material differences); distinction between law and Kencleng policy.
- **OPEN PRODUCT DECISION:** Exact scope of initial First Owner privileges, detailed duplicate-claim reconciliation and ownership conflict/appeal routes, withdrawal after material organizational action, and consequential access boundaries.
- **OPEN PRODUCT DECISION:** Material escalation thresholds/authority assignment exceptions, fairness of notification/opportunity for inconclusive closure, and resubmission restrictions.
- **OPEN PRODUCT DECISION:** Detailed organizational obligations language and version-change handling; applicable Campaign-specific duties remain for later Product Outcomes.
- **OPEN PRODUCT / PRIVACY-LEGAL DECISION:** Purpose, minimum sensitive data, lawful basis, disclosure, retention/deletion and verified legal/regulatory requirements. No universal KTP, NPWP, bank account, file upload, or mandatory identity-document collection is assumed.
- **OPEN EXPERIENCE/DESIGN DECISION:** How actors understand, exercise, and recover within these meaningful decisions; no screens, fixed status vocabulary, or flow is prescribed.
- **OPEN ENGINEERING CONSTRAINT / IMPLEMENTATION:** Permission matrix, backend mechanisms, verification tooling, and audit storage. None is approved for delivery by this checkpoint.
- **Next Product inquiry:** O2.8 cross-outcome invariants and negative/misuse case consistency, then O2.9 falsifiable behavioral acceptance and remaining legal/domain validation before marking DoD or seeking binding exact-revision acceptance.

## 6. Checkpoint #3 and minimum legal/privacy evidence — 18 November 2026

- [Working checkpoint #3 — post-#2 decisions](po-01-working-checkpoint-3.md): records explicitly agreed working directions on duplicate/existing organization claims, bounded First Owner rights, organization eligibility policy, claim-based evidence, auditability, privacy minimization and purpose-based retention. **Not binding Product acceptance.**
- [Minimum Legal / Privacy Evidence](po-01-minimum-legal-privacy-evidence.md): source-linked statutory and official status review for Yayasan, Perkumpulan and UU PDP; separates narrow source-checked facts, Kencleng Product implications and remaining legal/privacy validation blockers. **Not legal sign-off.**
- **Standing unchanged:** original Product DoD checkboxes remain uncompleted; exact-revision binding Product Contract, consumer sufficiency and implementation authorization remain outstanding. Read this section with original checkpoint #2 in §5; neither checkpoint silently supersedes original accepted Product Intent.
## 7. Agreed working Actor Value and Product Success Conditions — 18 November 2026

> **Standing:** Anhar explicitly agreed to the Actor Value and SC-01..SC-06 in the Product working session and authorized repository persistence. **Working Product commitments, not an exact-revision binding Product Contract, verified demand, actual successful acceptance tests, completed DoD, or implementation authorization.** This section is the current O1.3/O1.4 candidate basis; original DoD boxes in §3 remain unchecked pending proportional evidence/acceptance assessment.

### 7.1 Actor Value

| Actor | Intended outcome / value |
| --- | --- |
| **Applicant** | Can start an application, understand its outcome and next permitted step, without first having to prove every organizational mandate; application access does not confer organizational control. |
| **Legitimate organization / authorized representative** | Can seek recognition within Kencleng's policy while avoiding control being transferred to an unauthorized applicant; organizational acceptance remains an independently authorized act. |
| **Reviewer** | Can exercise professional judgment on appropriate evidence and provide a concise human attestation, without a document-understanding engine or unnecessary per-document verification workflow. |
| **Kencleng / authorized Decision-Maker** | Can grant, withhold, or decline initial eligibility with legitimate authority, traceable evidence references, meaningful reasoning when material and a reconstructable minimal audit trail. |
| **Visitor / Donor and other downstream actors** | Are not misled about what organization eligibility confirms: it does not certify Campaign approval, permission to receive donations, fund settlement or proven real-world impact. |

**Trade-offs / evidence standing:** protect organizational mandate against first-filer takeover without creating an automatic veto for unsupported claims; preserve decision accountability without excessive Reviewer workload, sensitive-data collection, or indefinite raw-evidence retention. These are reasoned Product outcomes grounded in representative scenarios and the agreed intent, **not independently validated demand, user research or implementation findings**.

### 7.2 Observable Product Success Conditions — accepted working candidates

| ID | Falsifiable Product success condition |
| --- | --- |
| **SC-01 — Legitimate application start and submission** | An Applicant meeting the minimum entry/submission conditions can initiate and request review without being granted Owner, Staff or other organizational rights solely for applying. |
| **SC-02 — Human-led review** | An authorized Reviewer can examine linked evidence and attest to material eligibility criteria; successful review does not depend on automated reading or interpretation of document contents. |
| **SC-03 — Valid approval gate** | Initial Approved cannot take effect unless the necessary organization qualification, authorized accountability acceptance and valid First Owner mandate/acceptance have been judged sufficient and an authorized Kencleng human decision has occurred. |
| **SC-04 — Truthful distinct outcomes** | Approved, Rejected, Closed Inconclusive and Withdrawn convey materially different reasons and consequences; pending/unknown findings are not automatically disqualification or fraud. |
| **SC-05 — Minimal decision accountability** | For consequential PO-01 actions, authorized parties can reconstruct what happened, when, in which application/organization context, by whom and for whom/what, result, and relevant evidence/attestation/decision context; material history is not silently overwritten. |
| **SC-06 — No unauthorized or misleading grants/disclosure** | Application, evidence submission or eligibility status alone cannot confer unauthorized Owner/Campaign/money rights, expose another applicant's confidential evidence, or claim verification/impact beyond the authorized decision. |

These are **Product-level acceptance targets**, not measured system performance, conversion/approval-rate targets, completed tests, engineering test plans or proof of deployment. High approval rate is **not** a proxy for correct, legitimate decisions. Later O2.9 acceptance scenarios may test these conditions without redefining their meaning.

## 8. Agreed minimum privacy baseline — 18 November 2026

> **Standing:** Anhar agreed to this **minimal PO-01 Product privacy direction** in the working session; persisted with the Actor Value / Success Conditions checkpoint. It establishes necessary purpose/access boundaries, **not** legal clearance of a particular lawful basis, retention period, data subject notice, or feature implementation. Read with [minimum legal/privacy source review](po-01-minimum-legal-privacy-evidence.md) and current Engineering [Privacy Standard T04](../../../engineering/standards/privacy-data-protection.md).

- **Purpose-driven minimum information:** (a) Applicant account attribution, necessary contact and declared organizational relationship; (b) legal organization identity and reviewable supporting references; (c) sufficient representative/First Owner attribution, mandate and acceptance; and (d) material decision and audit-event references. No universal KTP/NIK, full ID scans, personal NPWP, bank account data or system understanding of legal documents. Additional sensitive data only where a real material purpose justifies it.
- **Audience-specific disclosure:** Applicant sees their own application as authorized, not every competing Applicant's private file; authorized Reviewer sees evidence needed for assigned review; Decision-Maker receives information necessary for a valid decision; First Owner receives organization information/obligations within effective rights; public disclosure is restricted to specifically authorized, truthful organization/eligibility claims. Business audit trail and private mandate evidence are **not public by default**.
- **Three retention purposes, not one universal lifetime:** (1) draft/application information is held only for legitimate application and related purposes; (2) raw evidence is held only while review, lawful challenge/accountability or other justified purpose requires it; (3) decision/attestation and material business audit records remain traceable while a lawful accountability purpose requires them. Record references may remain after lawful evidence deletion, but must not falsely imply the raw evidence is still retrievable or independently re-verifiable. Respect corrections and lawful deletion/exception handling.
- **Lawful-basis validation still OPEN:** Map actual purposes separately for Applicant application processing, third-party representative/mandate evidence, First Owner attribution/acceptance, and decision accountability. For each actual purpose, validate legal basis, appropriate notice, audience/recipients, retention triggers/periods, data-subject rights, lawful exceptions and any regulated requirements. Do not assume universal consent or arbitrary time periods.
- **Responsibility boundary:** Product defines what is necessary and who may legitimately receive it; human Reviewer/Decision-Maker owns legal/evidence judgment; Engineering owns technical data protection, authorization, protected audit and security logging in accordance with effective standards. Security logs, business records and raw evidence have distinct purposes; IP/device/user-agent are not universal PO-01 business audit requirements.

**Next Product inquiry:** consolidate OA-01..05 acceptance wording and minimum clarification/closure exception behavior, then reassess remaining Product/legal/privacy blockers and prepare an exact-revision Product Contract candidate. Keep DoD and downstream handoff acceptance separate.
## 9. Agreed minimum OA commitments and exception behavior — 18 November 2026

> **Standing:** Anhar explicitly agreed to the minimum OA-01..05 wording and single organizational acceptance model, then agreed to the bounded clarification/withdrawal/closure behavior in the Product working session. **Working Product decisions only:** not external legal sign-off, an approved enforceable contract, Product DoD completion, UX flow, or Engineering authorization. This is the latest working interpretation of §5.5–5.6 where wording differs.

### 9.1 Organizational accountability commitments — plain meaning

| ID | Working organizational commitment | Meaningful limit |
| --- | --- | --- |
| **OA-01 — Truthful Information** | Organisasi memberikan informasi material yang benar, tidak menyesatkan, dan memperbaiki ketidakakuratan material setelah diketahui. | Bukan jaminan tanpa kesalahan atau verifikasi otomatis oleh Kencleng. |
| **OA-02 — Valid Representation** | Organisasi memastikan tindakan, penerimaan komitmen, dan penunjukan perwakilan/First Owner di Kencleng dilakukan pihak yang berwenang sesuai cakupan mandatnya. | Applicant tidak otomatis menjadi perwakilan atau Owner. |
| **OA-03 — Material Change Disclosure** | Organisasi memberitahu Kencleng perubahan yang diketahui dan secara material memengaruhi eligibility atau kewenangan perwakilan. | Tidak mewajibkan pelaporan setiap perubahan administratif kecil. |
| **OA-04 — Proportionate Cooperation** | Organisasi bersedia memberi klarifikasi dan bekerja sama secara wajar dalam pemeriksaan/pengawasan Kencleng yang sah dan relevan. | Bukan akses tanpa batas ke seluruh informasi/dokumen pribadi atau organisasi. |
| **OA-05 — Continuing Accountability** | Apabila organisasi melakukan Campaign atau kegiatan dana melalui Kencleng, organisasi memenuhi komitmen accountability yang dinyatakan berlaku untuk kegiatan tersebut, termasuk setelah fundraising berakhir. | Bukan izin Campaign/donasi/dana dan bukan persetujuan otomatis atas seluruh persyaratan baru/masa depan yang belum dikomunikasikan. Campaign-specific duties remain owned by the applicable future Product Outcome. |

**Single authorized organizational acceptance:** OA-01..05 accepted as **one identifiable, versioned package** by a party sufficiently authorized to bind the organization for this act. No universal five separate acceptance steps, mandatory on-product legal signatory account, mandatory digital-signature method, or required screen/checkbox pattern. External acceptance evidence may be assessed if sufficiently traceable and lawful.

**Keep the acts distinct:** (a) Applicant application terms/submission; (b) authorized organization acceptance of the OA package; (c) proposed First Owner's separate explicit personal role acceptance; (d) Kencleng human judgment and eligibility decision. One human may perform multiple acts only where actually authorized, but none is silently inferred from another.

**Minimum traceability:** identity of the legal entity, authorized accepting person/organ and capacity, *actual* OA package/version, attributable acceptance act and when it occurred, relevant evidence link, Reviewer attestation and eventual authorized decision. This is one consequential business event with connected decision evidence, not machine legal-document understanding or a required database schema.

**Change handling:** Material OA version changes before approval require acceptance of the version applicable to the decision. Later material versions cannot silently inherit previously granted acceptance; detailed post-approval amendment/renewal handling is adjacent. The legal effectiveness of a particular form of acceptance and the effect of an application withdrawn before approval remain case-specific legal/Product questions; retain the historical acceptance act without claiming more than it established.

### 9.2 Minimal exception and recovery behavior

| Condition | Required Product behavior / bounded consequence |
| --- | --- |
| **Clarification Required** | Authorized Reviewer requests information relevant to an unresolved material requirement. The application remains undecided; the Applicant receives an access-appropriate explanation of what is missing and what may be done next. Clarification is a process condition, **not rejection**. |
| **Closed Inconclusive** | After reasonable opportunity to clarify, necessary evidence or authority still cannot be established. An authorized, reasoned closure conveys **no initial eligibility / First Owner rights**, without asserting proven legal ineligibility, misconduct or fraud. |
| **Rejected** | Authorized decision based on sufficiently established material ineligibility or a proven relevant failed requirement. A clearly evidenced refusal of baseline OA by a party authorized to bind the organization may support rejection; Applicant's unsupported refusal or an unverifiable claim does **not** automatically establish the organization's refusal. |
| **Withdrawn before effective approval** | A party legitimately entitled to manage their application may request withdrawal in a valid context. Withdrawal ends that application without granting organizational rights. It does not automatically rescind separate authorized organizational acts/acceptance, erase material decision history, or override rights of the organization or other parties. Record any unresolved consequential effects rather than inventing a universal revocation outcome. |
| **After effective approval** | Applicant withdrawal is **not** a mechanism to revoke an eligible organization or its Owner rights. Post-approval eligibility/authority changes belong to relevant separately authorized lifecycle processes. |
| **New evidence / resubmission** | Material new evidence may justify a legitimate subsequent review; it does not automatically change or overwrite a prior decision, bypass restrictions, or confer rights without a new authorized decision. |

**No arbitrary Product timeout:** Fair clarification opportunity, notification and case closure must be proportionate; do not silently create an indefinite pending state or an automated closure deadline without a legitimate scoped operational decision. A credible material authority conflict holds the affected decision and requires authorized independent handling without exposing another Applicant's private evidence.

**Minimum consequential audit:** for clarification, closure, rejection, withdrawal, material correction or reassessment, preserve what/when/in which application/organization context/who/for whom/result and material reason/evidence reference where needed. Do not require raw-document interpretation, exhaustive procedural logs, or infinite retention. Rights and disclosure follow §5.4 and §8.

### 9.3 Standing and next review

- O2.6 and exception edges of O2.7/O2.9 now have **agreed working behavior**; their DoD checkboxes are **not** automatically complete or binding.
- Still OPEN: legal sufficiency of specific mandate and OA acceptance; privacy lawful bases/notice and purpose-specific retention; fair operational clarification/closure boundaries, detailed post-approval OA version changes and exceptional authority dispute routing; exact-revision Product acceptance and downstream consumer sufficiency.
- **Next:** review the separate [consolidated PO-01 Product Contract candidate](po-01-product-contract-candidate.md) as *one* bounded draft. That candidate does not promote these working agreements into approved Product Authority by being written.

## 10. Consolidated review clarification — 18 November 2026

> **Standing:** Anhar explicitly agreed to the bounded First Owner minimum capabilities, **private-by-default eligibility**, time-contextual evidence traceability and separated approval/readiness levels during consolidated Product review, then authorized revising PO-01 and its candidate. **Agreed WORKING Product meaning, not binding exact-revision approval, legal/privacy clearance, completed DoD, or implementation permission.** This later scoped clarification resolves only the affected ambiguity in §5.8 and other earlier working text; historical checkpoints remain historical.

### 10.1 Minimum First Owner behavior at effective initial approval

The legitimate recognized First Owner, **only after** effective initial Organization Eligibility approval and subject to current mandate/authorization, must be able to:

1. **View** the recognized organization's relevant eligibility decision, context and applicable organizational accountability commitments.
2. **Provide** relevant eligibility clarification or correction on behalf of the organization within the applicable mandate; providing material evidence/corrections does not itself alter a prior approval or bypass authorized review.
3. **Receive** material eligibility-related Kencleng communication as an authorized organization contact. This is a responsibility/ability to receive relevant information, **not** a prescribed notification channel, guaranteed delivery mechanism, or exclusive right to all organization correspondence.

These three capabilities do **not** confer Campaign creation/publication, donation acceptance, fund control, Staff/Owner administration, universal evidence access, or organizational power beyond valid mandate. The detailed operational permission model remains Engineering-owned and other Owner functions remain downstream Product decisions.

### 10.2 Eligibility disclosure — private by default

Initial Organization Eligibility decision/status and its evidence are **not automatically public**. Only legitimately authorized applicants, organizational actors and Kencleng reviewers/decision actors may receive the information permitted for their specific purposes and rights. Necessary internal downstream Product decisions can use truthful, access-bounded eligibility meaning; that is **not** authorization to expose it to visitors or donors.

Any public profile, badge, public-facing status, publication of reasons or evidence requires a **separate explicit Product disclosure decision** covering audience, allowed claim, timing and privacy consequences. No public surface or status text is specified by PO-01. Private-by-default is a Product disclosure rule, not a requirement for any particular UI/technical protection.

### 10.3 Evidence at decision time — proportionate traceability

For consequential Reviewer attestation or eligibility decisions, evidence references must identify **what sources/material were actually relied on at the time of the decision** well enough to reconstruct the then-known basis and its limitations, even if a live source later changes. If raw evidence is lawfully no longer retained, the record must not imply it remains available or independently verifiable. This is **not** a universal requirement for snapshots, content hashes, immutable documents, exhaustive per-claim findings, or indefinite original-file storage; Engineering selects the proportionate mechanism under privacy/legal constraints.

### 10.4 Separate Product acceptance, legal clearance and delivery readiness

- **WORKING behavior settled:** substantive scenario-based Product direction agreed by Anhar and durably recorded; **not binding contract acceptance**.
- **Binding Product semantics (future):** requires a separately approved, **exact immutable revision and explicitly bounded Product scope** by Anhar, including clear conditional assumptions and remaining legal/domain validation; acceptance does **not** declare any real mandate, OA contract or personal-data processing legally sufficient.
- **Legal/privacy and operational clearance (future):** before real use of affected evidence, acceptances and personal data, qualified review must establish applicable lawful basis, notice, retention/deletion limits and case-specific legal sufficiency. UX/Engineering consumer sufficiency, delivery/deployment authority and any runtime security claims remain separate later gates.

**Not resolved by this clarification:** case-specific Yayasan/Perkumpulan representative authority and legal enforceability of OA acceptance; lawful basis/notice, third-party personal data and retention terms; detailed post-approval authority lifecycle, appeals/escalation rules and future public disclosure. Do not mark the original DoD complete merely because this clarification is written.

**Minimum scenario checks for future Product acceptance:** after approval the legitimate First Owner can see the decision/obligations, provide an authorized correction and receive relevant communications; an unrelated Applicant cannot take those rights; a visitor cannot see a private eligibility status without later authorized public-disclosure rules; evidence changed or deleted after a decision cannot silently rewrite what was relied upon; exact-revision Product approval cannot be mistaken for legal/privacy or Engineering clearance.

## 11. Final review clarifications — 18 November 2026

> **Standing:** Anhar explicitly agreed to these **two narrow WORKING Product clarifications** during the final scenario review and authorized their persistence. They do not constitute exact-revision binding Product approval, external legal/privacy sign-off, completion of the DoD, or Engineering authorization. This is the latest interpretation of §10.1 and the legal/privacy readiness boundary; existing checkpoints remain historical.

### 11.1 Material correction reported after Approved

A recognized First Owner with valid organizational mandate may submit a material correction or clarification after effective initial Organization Eligibility approval. Kencleng must **receive and trace** that consequential submission and make it available for **appropriate authorized Kencleng handling**, including where it could call an earlier eligibility finding into question. The person submitting the correction **does not** independently change the organization's effective eligibility and the submission does **not automatically** suspend, revoke, reapprove or alter the prior decision. Any subsequent eligibility reassessment and effective rights consequences belong to a separately authorized ongoing eligibility/authority lifecycle, **outside the bounded PO-01 initial decision**. Product does not prescribe technical routing, queueing, notifications, review deadlines or suspension mechanisms here. No indefinite pending claim or silent overwrite of the original decision is implied.

### 11.2 Exact-revision Product semantics approval versus legal/operational clearance

- **Working Product behavior:** human-agreed scenarios and constraints recorded durably; not binding merely because they were discussed or committed.
- **Exact-revision Product semantics approval (potentially conditional):** Anhar may consider approving specifically bounded Product meaning at an immutable source revision **only if** the approval clearly names its scope, unresolved legal/privacy assumptions, affected clauses and non-authorized uses. Such an approval **cannot certify** that an actual legal mandate, OA acceptance, processing purpose or retention approach is lawful or ready for real use. No automatic acceptance occurs now.
- **Legally dependent commitments:** statements implying verified legal sufficiency/enforceability, lawful personal-data processing or operational readiness remain **BLOCKED** until qualified validation supports them. An unconditional Product acceptance or claims of legal compliance cannot be inferred from this candidate or source-linked legal research.
- **Delivery readiness:** consumer-specific sufficiency, Engineering implementation permission, legal/privacy clearance for production handling and deployment verification remain separate future decisions, irrespective of bounded approval of Product semantics.

**Next review input:** [PO-01 Contract Candidate](po-01-product-contract-candidate.md) and [minimum legal/privacy source review](po-01-minimum-legal-privacy-evidence.md). Original unchecked DoD items remain unchecked. No legal thresholds, retention durations or technical workflows are invented by these clarifications.
