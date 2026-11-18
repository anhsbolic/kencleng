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

**Prior working-domain context (not binding):** [Organization Eligibility — Working Product Definition](../organization-eligibility-working.md), pre-selection checkpoint, blob `e2f65984bb39a3d6d0b6b2d31c6240180be22964`. Its broader ongoing-eligibility decisions are context and boundary constraints, not automatically in-scope delivery requirements for PO-01.

**Research guidance, not Kencleng authority:** Harscode [Product Definition](https://github.com/anhsbolic/harscode-workspace/blob/302bedae43563a66d2ef914191c18139440c0c30/product/product-definition.md), [Consumer Handoff](https://github.com/anhsbolic/harscode-workspace/blob/302bedae43563a66d2ef914191c18139440c0c30/product/consumer-handoff.md), and [Authority & Approval](https://github.com/anhsbolic/harscode-workspace/blob/302bedae43563a66d2ef914191c18139440c0c30/product/authority-and-approval.md); selected research guidance, still pending its own human exact-revision review.

**Other owning inputs:** [Design Authority](../../design/README.md), [Engineering Standards](../../engineering/standards/README.md), and applicable legal/privacy/security authorities when relevant. These are separately owned sources; their factual or operational approval is not supplied by this document.

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
> **Scope:** Selected PO-01, representative Yayasan scenario; differences for Perkumpulan still to be tested. The shared [Actor & Authority Model](../actor-authority-model.md) supplies reusable Product meaning. The wider [Organization Eligibility working checkpoint](../organization-eligibility-working.md) remains historical/adjacent discovery and is not a second competing authority.
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

Reviewer may classify a *material claim/evidence finding*, not the entire organization, as:

- **Sufficient:** reliable support for the material claim following appropriate assessment.
- **Insufficient:** evidence not yet adequate; seek relevant clarification or support.
- **Unverifiable:** evidence exists but its material assertion cannot be established from available inspection; do not assert fraud or verified truth.
- **Contradictory:** materially conflicting evidence; hold affected decisions and escalate where risk requires.
- **Disqualifying:** adequate evidence establishes that a material eligibility criterion is not satisfied; can support a reasoned substantive rejection.

Product needs evidence source/provenance, relevant freshness, examiner, what was actually checked, limitations, and relationship to decision. This does not mandate a specific document-upload method or technical verification mechanism. Required legal document sets, source authority, and precise sufficiency thresholds for Yayasan/Perkumpulan remain OPEN for proportional legal/domain review.

### 5.4 Human-led decision authority and traceable review

- Kencleng uses **manual human approval** for initial eligibility. Completing fields, submitting documents, or automated completeness checks cannot automatically grant eligibility.
- Reviewer and Decision-Maker are separate responsibilities, not necessarily separate people: an appropriately authorized person may perform both in normal, non-conflicted cases.
- Disqualifying conflicts of interest prohibit that actor's assessment; material authority disputes, credible unauthorized-takeover risk, or decision-authority concerns require a hold/independent authorized escalation. An escalated reviewer still cannot waive material evidence sufficiency.
- Every consequential decision must allow reconstructing the application/organization, examiner, decision-maker and applicable authority, evidence and findings, reasoning/criteria, timestamps of material events and effective consequences, and subsequent corrections without silently replacing history.
- Privacy-appropriate disclosure is part of legitimate decision meaning. Product states *what must be traceable*; Engineering chooses audit-log, authorization, and storage mechanisms. No blanket two-person approval or automated document-verification requirement is implied.

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
