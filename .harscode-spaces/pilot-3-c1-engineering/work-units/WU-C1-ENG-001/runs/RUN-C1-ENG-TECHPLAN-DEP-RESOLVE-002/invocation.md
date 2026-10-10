# Run Invocation — RUN-C1-ENG-TECHPLAN-DEP-RESOLVE-002

## Identity and routing

- **Work Unit / Run:** `WU-C1-ENG-001` / `RUN-C1-ENG-TECHPLAN-DEP-RESOLVE-002`.
- **Phase route:** Bounded post-Review factual chronology resolution and exact-candidate Human report generation / Planner.
- **Participant / Session:** `PARTICIPANT-C1-ENG-DEPENDENCY-RESOLVER-002` / `SESSION-C1-ENG-DEPENDENCY-RESOLVER-002` (fresh Planner Session).
- **Dispatch posture:** Human-Assisted; prepared, not dispatched.
- **Model / effort:** `gpt-6-luna` / `medium`; registry does not require Run-specific approval.
- **Communication:** Bahasa Indonesia; canonical identifiers retained.

## Assignment and exact boundary

Review `RUN-C1-ENG-TECHREVIEW-DEP-002` completed with 0 blocking / 0 mechanical findings against candidate SHA-256 `7fc6d4539b388fda3cd0c7065609969443d578286a5ae3dfa26a09075f78627c`. During Orchestrator reconciliation, one inherited §1 chronology sentence was identified: it can be read as saying the current candidate already completed independent Review and bounded resolution before Human approval, although the candidate header/intro correctly says this candidate is In Review and not approved. Treat this as an Orchestrator-observed factual/mechanical clarification, not as a Reviewer finding; do not modify the Review findings or imply the Reviewer reported it.

Bounded task:
1. Recheck candidate and predecessor hashes and inspect the exact §1 wording/context plus Review result.
2. Mechanically clarify §1 so it distinguishes the previously approved predecessor and its review history from the current In Review candidate. Preserve all execution meaning, proposed pins, five advisory dispositions, trust semantics, Build gates, scope, task topology, G1–G3, and Open Items. Do not edit any other section unless exact report-gate metadata required by the template is demonstrably stale; if so, stop and report rather than expanding.
3. Record before/after candidate hashes and exact textual delta in `evidence/resolution.md`. State plainly that the delta is factual chronology only, changes no executable meaning, and does not require re-review under the applicable guardrail.
4. After that correction converges, regenerate `techplan/report-techplan.md` in full from the resulting exact candidate using the current report template. It must say current candidate is In Review / awaiting Human approval; identify the predecessor review as history and the current Review verdict (0 blocking/0 mechanical); accurately present proposed pgx/x/text baseline and five unresolved advisories; no new decision.
5. Write exactly one structured `evidence/phase-handoff.md` containing candidate/report identities, source review/evidence identities, correction classification/re-review rationale, open blockers, remaining exact Human gate, and recommended next route.

Do not promote candidate or change current Approved `techplan.md`. Do not edit dependencies, production code, T2 task, manifest, API, Product/pre-engineering authority, Control Tower, Space, WU, Events, or other Run evidence. Do not run Build/tests/runtime/dependency downloads, close `B-T2-001`/`B-T2-002`, accept residual risk, or dispatch Build/T3/T5. If the wording cannot be corrected as purely mechanical or any material ambiguity emerges, stop and route it; do not broaden the candidate.

## Bound inputs

- Current candidate before correction: `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/techplan/techplan.candidate.md`, SHA-256 `7fc6d4539b388fda3cd0c7065609969443d578286a5ae3dfa26a09075f78627c`.
- Current Approved predecessor: `techplan/techplan.md`, SHA-256 `bae2470f9038521b5d2d51fe4bc8de192fe5a9cf91bc0975831a3046fbf7a8ad`.
- Current report path is stale for this successor and must be regenerated: `techplan/report-techplan.md`, prior SHA-256 `c8b59544f80f74bf64ec4d7aedb4babbd45f232a9d31cdc4663e7b3bbfb78779`.
- Current independent Review findings SHA-256 `eeb82a25b3f566fd16a983a9d1ff2d74d006d3f8a99103cd9dab4eab392385e2`; handoff SHA-256 `49de2036b61ebe15f305f6e67c7de5ec82fc3753c14f53f4863e87e085d9d019`.
- Candidate Planner handoff SHA-256 `9cb0cbc3d3db330911e6366aa4f87d775b658a15185ad8897bfe60ff97284d1f`; dependency analysis SHA-256 `3996b4eee9c2df5cc6ff4a2e31e70e2526174cc8077d90595ee853642c39f198`; candidate self-check/allowed delta SHA-256 `1b3a4bb149c3103c4dacdc32b9ad761bb511097f22920783e117b1e60b04a647`.
- Approved report template: `../harscode-workspace/workflow/2-techplan/report-template.md`; synthesis prompt, Techplan template/rules/guardrails, and orchestrated overlay govern.
- Target/workflow revision: `59f2b25940fa4baf243b4f1d2ca2ba04c9b17fd4` / `3d9cbafaf4c2921c4ada4e3e5be8f6c9cb238590`.

## Execution envelope

- **RUN_PATH:** `.harscode-spaces/pilot-3-c1-engineering/work-units/WU-C1-ENG-001/runs/RUN-C1-ENG-TECHPLAN-DEP-RESOLVE-002`.
- **ARTIFACT_TARGET:** candidate chronology sentence and derived Human report only.
- **PREAUTHORIZED:** Read bound files/guidance; make the one factual/mechanical sentence clarification; regenerate exact candidate report; write this Run's evidence.
- **ORCHESTRATOR_DECISION:** Reconcile hashes/handoff and route exact Human approval gate.
- **HUMAN_REQUIRED:** Exact approval of the resulting candidate/report before lifecycle promotion; no implementation authorization changes.
- **Out of scope:** semantic plan change, re-review unless correction proves material, predecessor promotion, implementation/dependency/test/build/runtime, blocker closure, downstream dispatch.

## Canonical guidance

Follow `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`, `workflow/orchestrated-run-overlay.md`, current `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`, and `report-template.md`. Report generation is allowed only after exact candidate correction and review convergence; do not hand-maintain the digest.

## Prepared state

- **Dispatch readiness:** READY_FOR_HUMAN_DISPATCH; Luna/medium needs no model approval.
- **Participant dispatched:** No.
- **Run outcome:** Not started.
- **Next action:** Human mechanically dispatches the fresh Planner Session with this invocation.
