# ORG-01 — Concept C UX research checkpoint (2026-11-13)

> **Standing:** BOUNDED EXPERIMENTAL CHECKPOINT / POSITIVE HUMAN REVIEW OF SCOPED UX / SUBSEQUENT MATERIAL ACTOR-JOURNEY MISMATCH IDENTIFIED. UX v1.1 investigation paused for now; this is not exact-revision Product/UX/Hi-Fi acceptance, user-tested usability, operational implementation, release approval or Harscode framework promotion.
> **Owner:** Anhar (Human qualitative reviewer); target-project research only, Kencleng `pilot/station5-c1-space-only`. **Scope:** Step 1A–1E actor discovery and Step 2A–2E Concept C exploration with inspectable structural Lo-Fi v1.
> **Prior project baseline inspected:** `e43bc0276a9d027813e07d19099d86a5ee9e8c49`. No changes to Kencleng Product Contract, privacy policy, Rancang, accepted UI handoff/prototype, application code, or `main` in this checkpoint.

## Post-checkpoint Human clarification — actor-journey mismatch (2026-11-13)

**Evidence:** After the positive scoped Lo-Fi review, Anhar clarified that the produced UX was coherent and the translation was acceptable for the **narrow task**, but **did not match the intended broader human experience**. The target actor was envisioned as a person whose **goal is to raise funds via Kencleng**, not someone whose intrinsic goal is to establish an internal organization/Owner relationship. This feedback is **new, material Human product-intent/context evidence**; it does not retroactively negate the previous favorable assessment of the bounded Lo-Fi execution.

**Human-envisioned experience (illustrative direction, not newly approved Product scope):** a public Kencleng landing experience that earns justified trust; a clear way to discover how to fundraise and enter via fundraising CTA or organization menu; straightforward explanation of the process and its requirements; initial organization registration; the necessary identity/login binding; then an authenticated organization context/dashboard showing the actor's actual organization/Owner relationship and truthful available next actions. The actor's fundraising goal is the overarching motivation; ORG-01 may be one enabling outcome. These are **Human expectations / investigation input**, not a mandate to build every screen, add mandatory data, claim verified status, or automatically grant fundraising capability.

**Boundary to investigate:**
- Product Intent already covers fundraising generally. The approved ORG-01 Product Contract intentionally covers a narrower, attributable organization + internal first-Owner outcome; Campaign eligibility, public visibility and donations remain distinct. Therefore **"Product input was wrong" is a hypothesis**, not a demonstrated defect in the approved ORG-01 contract.
- The gap could arise from **parent actor journey/value framing, division of Product outcomes, selection/packaging of input for UX, or UX discovery's failure to challenge a too-narrow task boundary**. Audit the actual formation/decision chronology and current project authority before assigning causality.
- Anhar's illustrative "register, then login, then dashboard" must not bypass ORG-01's binding identity-attribution-before-effect requirement or pre-Google-login privacy explanation. Any sequence can be explored only within those safeguards. Initial organization data beyond scoped minimum and actual dashboard capabilities require Product/Engineering authority where material.
- This new feedback is **not representative user research, implementation evidence, release approval, or rejection of the previously checkpointed Concept C Lo-Fi as a bounded experiment**.

**Research disposition (Human-directed):** UX v1.1 investigation is **paused/concluded for now without framework promotion or final generalized success claim**. Preserve the artifact and positive review as scoped evidence **alongside** this newly discovered end-to-end experience mismatch. Move to **Product-level ORG-01 revisit and formation audit**, starting with Product Intent → parent actor/journey intent → outcome decomposition/ORG-01 Contract → decision and handoff trail. Do not preemptively amend approved Product policy or create a new roadmap. No changes to runtime, Hi-Fi candidate or `main` are authorized by this checkpoint.

**Authority/standing:** This section is the **latest working research interpretation** and supersedes the earlier “compare old/new Lo-Fi and pursue UX refinement next” suggestion below; the original earlier review/notes remain as dated evidence.

## Human feedback — observed

On 2026-11-13 Anhar reviewed the Concept C experiment and said: **“hasil eksperimennya, menghasilkan UX yang sesuai. translasi menjadi UX nya juga oke.”** Anhar then requested that the result be updated and persisted as a checkpoint before the next discussion.

**Decision interpretation:** The Human considered the experimental UX direction and its translation to the presented Lo-Fi satisfactory as a bounded research checkpoint. This feedback is a favorable *Human design review*, not evidence that representative organization actors successfully understood/completed tasks, nor exact-revision final UX/Hi-Fi approval. It does not supersede the existing scoped UI candidate or Product authority.

## Why this was tried

An earlier bounded Relationship-First Lo-Fi drew Human criticism for feeling system-centric, too administrative, and insufficiently natural for an actual actor. A corrective experimental UX v1.1 framework was then strengthened to explicitly investigate actor experience across all material inputs, challenge Product sufficiency, route genuine owner-owned gaps, and evaluate naturalness rather than only document completeness. This checkpoint tests that *correction's usefulness in one case*, not universal effectiveness.

## Discovery and selected approach — working meanings

- **Step 1A:** Verified approved ORG-01 Product outcome, actor/identity/Owner boundaries, privacy direction, and relevant Design Standards. The source specifies what outcome must be true; actual actor awareness/motivation is not established by those facts.
- **Step 1B–1D:** Explored actor situations A–F (new visitor, ready visitor, signed-in without organization, signed-in with an organization, uncertain about authority, interrupted/unknown result) and candidate experience needs: relevance/orientation, appropriate entry, proportionate effort, informed commitment, and truthful result/continuity. These are scenario hypotheses, not observed actor research.
- **Step 1E:** Framed the core experience problem as helping a person decide whether organization formation is relevant, proceed at fitting readiness, understand consequences, and inspect truthful results. No new mandatory Product blocker was established for exploratory design.
- **Step 2A–2B:** Compared education-first, direct formation, and contextual progressive formation. **Concept C** was explicitly selected by Anhar as working direction and retained after scenario-based stress test.
- **Step 2C–2D:** Proposed public orientation without mandatory onboarding; privacy explanation before simulated Google Login; login after orientation but before name entry as a *reversible UX choice*; authenticated existing-organization context only when knowable and relevant; minimum form; account/Owner meaning near explicit creation action; separate honest success/failure/partial/unknown states; inspectable relationship following genuine success.
- **Step 2E:** Built and reviewed the attached client-only structural HTML. It provides six reviewer-controlled actor situations, synthetic result switching, example similar-name behavior, desktop/mobile presentation, and a separate inspector surface that is not part of user-facing UI.

## Exact inspectable artifacts

- [Concept C Structural Lo-Fi v1](ux-concept-c-lofi-v1.html) — HTML/CSS/JS, simulated local state and reviewer tools. Original file SHA-256 before Git persistence: `2c54e50b20f885d49303f7893a74b3263fc15ab17aea7c2f731bff41de89c470`.
- [Lo-Fi v1 original review notes](ux-concept-c-review-notes-v1.md) — frozen pre-feedback rationale/trace and *at-that-time* proposed status. Original SHA-256: `2656aa91054673c5786b91eebfe58429146ce54be409262869850bb7cfb2c601`. **The earlier 'NOT APPROVED' label describes the artifact before the Human review; the present checkpoint records the later positive review with its bounded standing.**
- Screen captures and a ZIP review package were created in the originating session, but are **not claimed to be persisted in this Git checkpoint**. The HTML is the source artifact, and screenshots can be regenerated for visual review.

## Binding boundaries and unresolved dependencies

1. **Product truth:** ORG-01 is an actual organization context plus first attributable internal Owner as one outcome, not legal representation, Reviewer approval, Campaign creation, public listing, or donation eligibility. Similar names alone do not auto-block/merge. Authority: `docs/product/outcomes/org-01/product-contract.md` (approved v0.3; prior blob `da19d8e7cdb9baac8cf4af1c5e6b61763607013e`).
2. **Privacy:** Minimum necessary data and pre-Google-login notice direction apply; login/notice in Lo-Fi are simulated, and final notice/legal basis and production data-real gate remain open. Source: `docs/product/outcomes/org-01/privacy-decision.md` (prior blob `9dcbc448cd387633422aecfc9bc8a39dced70c04`).
3. **Unknown/partial result:** Real authoritative read-after-submit/status resolution, duplicate safety and safe recovery were not runtime verified. No blind retry entitlement is approved or simulated as a real implementation.
4. **Existing-organization context:** Inspectability is a Product obligation, but actual relationship lookup, access control, authenticated session behavior and integration were not verified.
5. **Value and quality evidence:** No representative actor interviews, comprehension tests, independent usability/accessibility testing, Google login, backend persistence, or product/runtime verification. Positive Human qualitative review is narrower evidence.

## Scope and next research boundary

**Checkpoint decision:** Preserve Concept C and Lo-Fi v1 as the **current preferred bounded research direction**; do not reopen settled Product rules, replace existing project Hi-Fi, modify accepted Harscode UX/UI definitions, start implementation, or promote to `main` because this experiment was liked.

**Next:** Compare the accepted-for-checkpoint UX candidate with the earlier negative Human critique; identify what materially changed and what remains hypothetical. Then decide whether another targeted actor/consumer validation, UX source refinement, or design handoff is justified. Persist any later binding scoped UX decision separately with explicit approval and exact revision.

## Source references

- Kencleng approved ORG-01 Product: `docs/product/outcomes/org-01/product-contract.md`, `approval.md`, `privacy-decision.md`.
- Kencleng Design Standards: `docs/design/product-design-principles.md`, `patterns.md`, `design-guidelines.md`.
- Harscode experimental framework: `anhsbolic/harscode-workspace` `research/cross-stage-mapping@98f3b1077fc224857902a818e6ce9be05a352fc0`, `design/ux-v1-1/` — research candidate, not operational.
