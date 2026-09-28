# Stage 2 — Input Provenance

> Phase/Stage: Exploration / Stage 2  
> Author: `P-S2-002-OIR-003-1` (`KC-EXPLORER`)  
> Created: 2026-09-28  
> Model / reasoning: invocation dispatch metadata `gpt-6-luna` / `high`; active runtime model is not independently exposed  
> Session: FRESH per invocation; Session ID not exposed  
> Work Unit / Run: `WU-S2-002` / `OIR-S2-002-003`  
> Target revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535` (`HEAD` at inspection; working tree had pre-existing changes outside this Run)  
> Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`

## Assignment-defining inputs

| Input | Expected SHA-256 | Observed SHA-256 | Result |
|---|---|---|---|
| `.harscode-spaces/participant-profiles/profiles.md` | `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32` | `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32` | Match; `KC-EXPLORER` profile pin verified |
| `.harscode-spaces/authority-map.md` | `65a73a5ab1233ff98699f425222d3cc9679e8c34176a9fcd3e1f60c9e20d433c` | `65a73a5ab1233ff98699f425222d3cc9679e8c34176a9fcd3e1f60c9e20d433c` | Match |
| `runs/TP-S2-002-007/techplan.md` | `441577782f7600fd5507d228c4dd6491009c32d6e8b81607b28a38a5ebb1a1a3` | `441577782f7600fd5507d228c4dd6491009c32d6e8b81607b28a38a5ebb1a1a3` | Match |
| `runs/OIR-S2-002-001/resolution-brief.md` | `3d2b22f6fca9bd23024eca41cda79417bcc5cf929513223867d73ddcc6440782` | `3d2b22f6fca9bd23024eca41cda79417bcc5cf929513223867d73ddcc6440782` | Match |
| `runs/OIR-S2-002-002/resolution-brief.md` | `5de8e28170e9b5ebc3b1789010433589c03268d8b4e7e7923ed3ce8363e70e9e` | `5de8e28170e9b5ebc3b1789010433589c03268d8b4e7e7923ed3ce8363e70e9e` | Match |
| `runs/OIR-S2-002-002/handoff.md` | `86d50645b32ec6eca36cb8694b254fc544873235158c48abbfe34981240748c6` | `86d50645b32ec6eca36cb8694b254fc544873235158c48abbfe34981240748c6` | Match |

The orchestrated target revision matches `HEAD`. The working tree was not clean: pre-existing modifications were present in the Work Unit manifest, Control Surface, events, outcome, readiness reconciliation, work graph, and development tracker, plus untracked `OIR-S2-002-002` evidence/brief/handoff and this Run's invocation. This Run did not modify those paths. No orchestration state was changed.

## Governing current-effective sources

- `docs/product/mvp-scope.md` §§Stage B and Stage C — simulator-owned truthful result, no time estimate/SLA, terminal-only status email, Security/PII windows and retention controls.
- `docs/product/mvp-delivery-slices.md` §5 — active Slice 2 version of those requirements, including backend-configured/fixture demo-only failure.
- `runs/TP-S2-002-007/techplan.md` §13 O2/O3 and §4 Q4/R4/R5 plus Risk-3 — unresolved simulator timing/scenario delivery detail and guest-email controls.
- `runs/OIR-S2-002-001/resolution-brief.md` O2 — prior Product state/recovery decision; exact timing/mechanism remains open.
- `runs/OIR-S2-002-002/resolution-brief.md` O3 and `handoff.md` — 24-hour verification and post-terminal retry windows; maximum verified-email retention while state remains `pending` explicitly deferred.
- `.harscode-spaces/authority-map.md` rows 1 and 4 — current Slice 2 Donation delivery/domain and Security/PII owner is Anhar Solehudin. Owner attribution does not itself decide timing or set a retention cap.

