# ORG-01 — Concept C Structural Lo-Fi v1: Review Notes

**Status:** UX EXPERIMENTAL / PROPOSED / NOT APPROVED. Date 2026-11-13. No real participants, authentic login, server requests, actual persistence, actual organization lookup, or real-world onboarding/analytics were tested.

## Open

Open `org-01-concept-c-structural-lofi-v1.html` in a modern browser. Use the **Kontrol reviewer** panel outside the simulated application to switch six actor situations, simulated submission outcomes, and the similar-name case. The product-facing surface is the white preview; the inspector is not part of the product UI. After a simulated success, the newly formed organization is discoverable in **Organisasi Anda** during that demo session; this is local in-memory behavior only, not persistent storage evidence. Narrow the browser for mobile; it is responsive.

## Current authority read

- **Kencleng:** `anhsbolic/kencleng`, branch `pilot/station5-c1-space-only`, HEAD `e43bc0276a9d027813e07d19099d86a5ee9e8c49`. Research-scoped approved Product, not Kencleng `main`.
  - `docs/product/outcomes/org-01/product-contract.md` blob `da19d8e7cdb9baac8cf4af1c5e6b61763607013e` (Product ORG-01 approved v0.3)
  - `docs/product/outcomes/org-01/privacy-decision.md` blob `9dcbc448cd387633422aecfc9bc8a39dced70c04` (approved approach, final notice and data-real release gate open)
  - `docs/design/patterns.md` blob `5c45ba62fc83dec6b802359317d642a408188e57` (canonical design pattern system)
  - `docs/design/product-design-principles.md` blob `92b74d9dd7c80da055dd632f18e3059fe2013eaa`
- **Harscode experimental UX:** `anhsbolic/harscode-workspace`, branch `research/cross-stage-mapping`, `design/ux-v1-1/definition.md` blob `70f7c76f42f894978d66af688db633cdfb8f144c`. Research candidate only, not operational `main`.

## Concept C thesis & trade-off

**Contextual Progressive Formation:** Orient without requiring onboarding; recognize actual account state when authorized; require only relevant initial data; disclose consequences close to commitment; confirm the actual combined outcome and allow separate inspection. Context informs, never guesses intention. More transparency does not mean more panels.

**Provisional choice:** Google Login **after initial orientation and before the form**. This is not Product policy. It avoids relying on unapproved draft preservation while keeping orientation public. The UX alternative (name before login) remains reversible if testing suggests significantly less friction and feasibility supports safe continuity.

## Flow / inspection coverage

| Situation | Reviewer selection | Expected inspection |
|---|---|---|
| A — new visitor | `A` | Public explanation, CTA, optional information, privacy before Google simulation, one-field formation, brief consequence, result |
| B — ready visitor | `B` | Can go directly from public CTA to identity boundary, without mandatory education |
| C — logged-in user without organization | `C` | Bypass login and irrelevant organization selector; arrive at form |
| D — logged-in user with organization | `D` | Existing context/Owner inspectable; may choose new organization rather than forced repeat or forced prohibition |
| E — uncertain actor | `E` | Role boundary available before commitment without legal-document requirement |
| F — interrupted/unknown | `F` | Distinct uncertain result; no false success/failure and no blind resubmission |

Reviewer **Simulasikan hasil** offers success, definitive failure, unknown, partial; results are synthetic and chosen externally by the reviewer, never computed by the app. A separate similar-name reviewer input does **not** auto-block formation or claim universal collision detection; it supplies only an example name in this Lo-Fi.

## Trace to source / provisional design

| Structural UX choice | Standing and source |
|---|---|
| Internal organization + first Owner effective together; success only when both real | Binding ORG-01 PD-01/02, SC-01; rendering simulated, not runtime verified |
| Actor should understand first Owner before effect | Binding PD-03/SC-02; briefing near submit is **UX hypothesis** pending comprehension review |
| Relationship separately inspectable after success | Binding REQ-05/AC-05; detail page is **illustrative UX structure**, not existing route/backend proof |
| Similar organization names don't automatically reject/merge | Binding PD-04/REQ-08; no UI collision lookup implied |
| Organization creation does not imply external verification, Reviewer approval, Campaign, public listing or donations | Binding PD-05; short copy is UX candidate |
| Identity available by effective act, not necessarily the beginning | Binding PD-02; login-before-form is **UX provisional choice** |
| Privacy summary before Google Login and link to full notice | Privacy approach approved, but final notice and legal validation unresolved; prototype *draft placeholder* not release-ready |
| Minimum data | Privacy PRIV-01: account, Google ID, org name/ID, Owner relationship, timestamp/necessary technical data. **Only actor-entered name** illustrated; ID/time managed by system hypothesis |
| Avoid duplicate submit, distinguish validation vs request failure, safe recovery | Canonical UX Pattern System; backend implementation not verified |
| Less explanatory burden and clear next action | Product Design Principles: Ordered Transparency, One Confident Next Action, Explain Consequences Before Commitment; naturalness is a hypothesis |
| Existing-organization context | Product says previously valid org need not repeat ORG-01. New/inspect options provisional; authentication/permissions and retrieval not tested |

## Material open dependencies

1. **Unknown / partial outcome**: Product requires honest states. There is no verified API/receipt/status resolution or safe retry design for loss of response; demo intentionally withholds retry and does not pretend `Kembali ke awal` resolves uncertainty. This needs Engineering/Security/UX verification before binding acceptance.
2. **Existing organization lookup and access**: Product requires durable inspectability, but exact account relationship discovery, routing and permission behavior have not been checked against runtime.
3. **Google login and privacy notice**: Login is simulated; final privacy notice, legal basis and production gate remain open. Demo copy is not approved final notice.
4. **Value proposition**: Actor research has not demonstrated that the phrase `Mulai kehadiran organisasi` or internal Owner benefit is understandable/motivating to real organization representatives. Do not invent Campaign or public visibility entitlement to strengthen copy.
5. **Responsiveness/accessibility**: static responsive layout inspected technically where possible; no assistive-tech usability validation or production accessibility certification.

## Critical review questions

1. On entry, would an organization representative understand why this is useful **before** knowing the internal concept of Owner?
2. Does any page request more explanation or work than the person needs to continue?
3. Does the login boundary feel reasonable, or too early for an unconvinced actor?
4. Can a person infer what `Owner` means and does **not** mean without reading an administrative policy?
5. Does the immediate result clearly distinguish creation from verification/review/donation permission?
6. Would a person with an existing organization see their current context without being falsely blocked from a second one?
7. Is the unknown state truthful yet unsatisfying? If yes, record as evidence of missing actual recovery capability, not as a mandate to invent a UI retry.

## Evidence calibration

- **Observed now:** Current source inspection, implementation of a working local simulated HTML state flow, source-to-concept trace. Visual rendering and interaction tests if run are expert artifact checks only.
- **Not observed:** real-user task comprehension, preference, delight, accessibility by assistive technology, Google identity process, persistence/duplicate safety/transactionality, backend status resolution, real notice/legal readiness.
- **No repository writes:** files are local research artifacts for Human critique. Harscode `design/ux/` and Kencleng Project UX remain unchanged.

**Recommended next move:** Human reviews public entry, form, success, and unknown on desktop/mobile; classify the most material naturalness issue first; refine structural Lo-Fi before UI-level visual work or acceptance.