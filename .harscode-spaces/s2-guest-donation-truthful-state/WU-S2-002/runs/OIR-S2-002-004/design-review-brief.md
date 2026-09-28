# O7 Product Design Review Brief — WU-S2-002

> Phase/Stage: Explorer Open-Item Resolution / Stage 3 (complete)  
> Work Unit / Run: `WU-S2-002` / `OIR-S2-002-004`  
> Role / specialization: Explorer / O7 Product Design terminology and source-label review facilitation  
> Author / Participant: `P-S2-002-OIR-004-1` / `KC-EXPLORER`  
> Created: 2026-09-28  
> Model / reasoning: invocation-selected `gpt-6-luna` / `high`; active runtime model is not independently exposed  
> Session: not exposed  
> Target revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535`  
> Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`

## Review result

O7 Design review is complete for the assigned terminology and source-label decision surface. Existing Product and Design direction is sufficient for the pending label, method availability, generic status-link failure, and deliberate failed-state recovery. Anhar Solehudin made two explicit, bounded Design decisions in this Participant Session on 2026-09-28: put simulation provenance before terminal state in status/notice labels, and use the complete approved-rule disclosure beside optional email opt-in. Authority provenance is the current Slice 2 Product Design row in `.harscode-spaces/authority-map.md`; the current-Slice-2 scope is bounded to this Run.

This brief does not claim rendered visual acceptance, implementation, Security/PII residual-risk acceptance, API parity, or `CONTRACT_READY`. No repository source outside this Run was changed.

## Per-surface findings and outcomes

### 1. Pending state

- **Finding / current state:** Product fixes the exact copy “Menunggu hasil simulasi”; the current frontend has no Donation status surface.
- **Outcome:** `GUIDANCE_SUFFICIENT` for meaning and wording. Use a visible, calm system-state treatment; do not turn it into warning/error, countdown, or payment instruction. The exact Product wording remains intact.
- **Alternatives and consequence:** A generic “Sedang diproses” is shorter but loses the explicit simulation source; an ETA/countdown would add an unapproved promise. Neither is needed given settled Product copy.
- **Design recommendation:** Keep “Menunggu hasil simulasi” as the user-facing state label. When explanatory context is needed, clarify that this is a sandbox result and not a real external payment settlement. Do not rely on color or a badge alone.
- **Source anchors:** `docs/product/mvp-delivery-slices.md` §5 lines 169–175; `docs/product/mvp-scope.md` Stage B; `docs/ui-ux/patterns.md` §§7, 14; `docs/ui-ux/design-guidelines.md` §13.
- **Deferred dependency:** Simulator timing/recovery remains O2 delivery evidence; no timing promise is approved. This review does not resolve it.
- **Smallest next route:** Carry exact pending copy and no-ETA/no-payment-instruction boundary into contract/spec and later rendered UI review.

### 2. Terminal `success` / `failed` status and source label

- **Finding / current state:** Product makes the backend simulator the source of terminal results and requires terminal email to identify a simulated result, not provider settlement. Design leaves final production provenance terminology open. No current Donation status or email surface exists.
- **Outcome:** `HUMAN_DESIGN_DECISION_RECORDED`.
- **Compared options:**
  1. **Source-first — selected:** “Hasil simulasi donasi: berhasil/gagal.” The user sees the source before the result; this better prevents a payment-settlement reading.
  2. **State-first:** “Status donasi: berhasil/gagal (simulasi).” This is compact and foregrounds state, but leaves simulation provenance in a secondary parenthetical that can be missed.
- **Observed Decision:** On 2026-09-28, Anhar selected the source-first label family for the status page and terminal notice. Provenance: direct Human response in this Participant Session; Anhar is the named current Slice 2 Product Design owner per the authority map. Scope: O7 terminology/source labeling for this Slice 2 Run. This records a Design wording decision; it does not change Product state semantics.
- **Design recommendation:** Use the selected source-first family consistently for terminal `success` and `failed` states and the subject/heading of the one terminal status notice. Keep supporting copy explicit that the result is simulated and is not a real provider settlement. Do not use provider receipt language or imply independent verification.
- **Source anchors:** `docs/product/mvp-delivery-slices.md` §5 lines 169–175; `docs/product/mvp-scope.md` Stage B/C; `docs/ui-ux/product-design-principles.md` §§2, 5; `docs/ui-ux/patterns.md` §§7, 13–14; `docs/ui-ux/design-guidelines.md` §§13, 26.
- **Deferred dependency:** Actual notice delivery timing, retry, retention, and proof that only a verified opted-in address receives it remain under O2/O3 and Security/PII/API work. No risk acceptance is implied.
- **Smallest next route:** Carry the chosen label family and truthful-source restriction into reconciled Donation spec/API and later notification/status UI; reopen only if new Product evidence conflicts.

### 3. QRIS active and other payment methods unavailable

- **Finding / current state:** Product explicitly lists QRIS, GoPay, ShopeePay, and bank transfer; only QRIS runs the sandbox simulation. The others are visibly unavailable and non-interactive. Current frontend has no method selector; Campaign detail states donation is unavailable.
- **Outcome:** `GUIDANCE_SUFFICIENT` for the Product behavior; no new Design-owner choice is needed for O7.
- **Alternatives and consequence:** Enabling or visually emphasizing all four as actionable choices would imply unsupported rails. Hiding the unavailable methods would omit the Human-approved display list. Showing all four with QRIS active and the rest visibly unavailable preserves the approved scope.
- **Design recommendation:** Preserve the complete approved list; give QRIS the sole active action and mark other methods plainly as unavailable using the existing neutral/pending grammar. Do not show usable real-payment instructions, active-looking controls, or provider-settlement cues for inactive methods.
- **Source anchors:** `docs/product/mvp-delivery-slices.md` §5 lines 168–179, 208–215; `docs/ui-ux/product-design-principles.md` §§2, 8; `docs/ui-ux/design-guidelines.md` §13; `frontend/mocks/fixtures/public-campaign.ts` `baseCampaign.donation_action`.
- **Deferred dependency:** No rendered selector exists; keyboard interaction, responsive visibility, and final visual treatment remain for implementation/rendered review.
- **Smallest next route:** Implement only after active Slice 2 contract/spec is ready; verify inactive methods remain non-interactive and clearly unavailable in rendered review.

### 4. Optional email opt-in and 24-hour verification disclosure

- **Finding / current state:** Product makes status email optional and opt-in. The 2026-09-28 Event records the owner’s direction to disclose near the opt-in that an address unverified within 24 hours from email capture is deleted and receives no status email, while Donation processing continues independently. No guest form exists.
- **Outcome:** `HUMAN_DESIGN_DECISION_RECORDED` for the exact local label/helper wording, with disclosure placement and clock origin already set by owner direction.
- **Compared options:**
  1. **Complete disclosure — selected:** Label: “Kirim pemberitahuan status donasi melalui email (opsional)”. Helper: “Verifikasi email dalam 24 jam sejak alamat dicatat. Jika tidak diverifikasi, alamat dihapus dan pemberitahuan tidak dikirim. Donasi tetap berjalan.” This gives the purpose, optionality, time origin, deletion/no-notice consequence, and independence from Donation processing next to the choice.
  2. **First-person opt-in:** Label: “Saya ingin menerima pemberitahuan status donasi lewat email”. Helper: “Alamat email akan dihapus bila tidak diverifikasi dalam 24 jam sejak dicatat; pemberitahuan tidak dikirim. Donasi tetap berjalan.” This is a direct checkbox voice and slightly shorter, while leaving the notification purpose less explicit in the helper itself.
- **Observed Decision:** On 2026-09-28, Anhar selected the complete disclosure wording. Provenance: direct Human responses in this Participant Session; Anhar is the named current Slice 2 Product Design owner. Scope: near-opt-in copy for current Slice 2 only. The 24-hour clock remains from when the address is recorded, not when the verification email is sent.
- **Design recommendation:** Place the selected label/helper adjacent to an unchecked, genuinely optional opt-in control. Do not move the disclosure only to Terms & Conditions, make email required, imply Donation processing stops, or phrase verification as a guarantee that a terminal notification will arrive.
- **Source anchors:** `docs/product/mvp-delivery-slices.md` §5 lines 172–177; `docs/product/mvp-scope.md` Stage B; `.harscode-spaces/s2-guest-donation-truthful-state/events.md` “Guest email verification disclosure confirmed for current Slice 2”; `docs/ui-ux/product-design-principles.md` §§4, 9; `docs/ui-ux/patterns.md` §4.
- **Deferred dependency:** The separate maximum-retention conflict for verified email while the Donation remains `pending` is unresolved. The selected disclosure does not set that cap or change terminal-notification policy.
- **Smallest next route:** Carry the exact selected copy into the reconciled form spec. Keep O2/O3 retention as a separate Product/Security/PII delivery Decision before any contract claims the full notification lifecycle is closed.

### 5. Terminal-only email notice

- **Finding / current state:** Product permits at most one status-only email after terminal `success` or `failed`, never initial `pending`, after address verification. The notice must identify simulation rather than provider settlement. No template, sender, or preview exists.
- **Outcome:** `GUIDANCE_SUFFICIENT` for eligibility and source family; exact message-body layout awaits the delivery artifact.
- **Alternatives and consequence:** A subject such as “Donasi berhasil” could imply a real payment result. The selected source-first label “Hasil simulasi donasi: berhasil/gagal” communicates the source at the most prominent label. A general “Status donasi diperbarui” is neutral but does not itself identify simulation.
- **Design recommendation:** Apply the selected source-first label to each terminal notice; keep body status-only and state that the outcome is simulated, not provider settlement. Do not send pending notices or imply repeated notices. Avoid excessive celebration around money.
- **Source anchors:** `docs/product/mvp-delivery-slices.md` §5 lines 172–177; `docs/product/mvp-scope.md` Stage B/C; `docs/ui-ux/patterns.md` §§13–14; `docs/ui-ux/product-design-principles.md` §§5, 11.
- **Deferred dependency:** The maximum verified-email retention while `pending`, delivery retry and retention controls, and runtime proof remain unresolved outside Design. No residual privacy risk is accepted.
- **Smallest next route:** Carry the label and terminal-only/no-provider implication into the reconciled notification contract; later verify send/no-send and deletion boundaries.

### 6. Generic unavailable behavior for status links

- **Finding / current state:** Product requires the same public behavior/copy for missing, invalid, or expired links; its example is “Link status tidak tersedia atau mungkin kedaluwarsa.” A prior owner decision selected generic `404`; transport parity remains open. There is no current Donation status route.
- **Outcome:** `GUIDANCE_SUFFICIENT` for visible UX direction; no new Design decision is needed.
- **Alternatives and consequence:** Distinguishing “sudah kedaluwarsa”, “tautan salah”, and “donasi tidak ditemukan” could disclose lookup state and contradict the generic behavior. One shared message avoids that distinction while satisfying the Product direction.
- **Design recommendation:** Use one generic message, e.g. “Link status tidak tersedia atau mungkin kedaluwarsa.” Do not present case-specific reasons or imply a status is available. Do not treat a shared message as evidence of transport parity.
- **Source anchors:** `docs/product/mvp-delivery-slices.md` §5 line 174; `docs/product/mvp-scope.md` Stage B; `docs/ui-ux/patterns.md` §§7, 12; `OIR-S2-002-002/resolution-brief.md` O4/O5.
- **Deferred dependency:** API/Security still need parity evidence for response status/body/headers/cache/timing, abuse controls, and residual-risk decision. That dependency is not a Design decision.
- **Smallest next route:** Preserve generic UX copy in authored status-link requirements; API/Security resolves transport controls and later Testing verifies parity.

### 7. Failed-state next action

- **Finding / current state:** Product allows the donor to explicitly start a new Donation after a terminal `failed`; `pending` never auto-resubmits. Existing UX guidance orders state → consequence → relevant next action, and says retry belongs only where retry can reasonably succeed. No failed Donation UI exists.
- **Outcome:** `GUIDANCE_SUFFICIENT` for action intent; Design recommendation recorded, no separate Human decision required because Product meaning is explicit and the wording follows it directly.
- **Alternatives and consequence:** “Coba lagi” is shorter but could imply retrying the same Donation or repeating settlement. “Mulai donasi baru” makes the new intent explicit and distinguishes it from the same-key ambiguous request retry handled by O9.
- **Design recommendation:** Show “Mulai donasi baru” after terminal failure, with any supporting sentence clearly indicating this is a new Donation. Activating it may open a new donation flow; it must not submit automatically. Do not expose it as an automatic action while `pending`.
- **Source anchors:** `docs/product/mvp-delivery-slices.md` §5 lines 170–171; `docs/product/mvp-scope.md` Stage B; `docs/ui-ux/patterns.md` §§7, 12–13; `docs/ui-ux/product-design-principles.md` §§8–9.
- **Deferred dependency:** New-key creation and ambiguous retry semantics remain delivery/API concerns under O9; status display does not authorize client settlement or key rotation while the result is ambiguous.
- **Smallest next route:** Carry explicit new-Donation intent into flow requirements and later test that the action opens, but does not silently submit, a new flow.

## Cross-cutting deferred concerns

- O2 simulator terminal timing and recovery remain without a Product ETA/SLA. This Run neither proposes nor implies one.
- The O2/O3 verified-email-while-`pending` retention conflict remains separately open. The selected opt-in disclosure covers only the owner-approved unverified-address consequence.
- O4/O5 token exposure, response parity, cache/timing/abuse controls, and residual-risk acceptance remain Security/API matters.
- No current Donation UI or email template exists; all visual/rendered judgments remain unverified until implementation.

## Verification and boundary

- **Verified by inspection:** invocation/profile pin; current Product, Design, authority, and referenced Run evidence; live frontend route inventory and Campaign detail unavailable state; target revision matches the invocation.
- **Human decisions observed:** two current-Slice-2 Product Design decisions are recorded above with date, scope, and direct-session provenance.
- **Not performed:** tests, runtime/API validation, browser automation, rendered visual acceptance, accessibility execution checks, implementation, Security/PII review, or residual-risk acceptance.
- This artifact is a focused Design review brief. It does not amend Product/MVP, canonical Design docs, spec/OpenAPI, code/tests, Approved Techplan, prior Run artifacts, or orchestration state.
