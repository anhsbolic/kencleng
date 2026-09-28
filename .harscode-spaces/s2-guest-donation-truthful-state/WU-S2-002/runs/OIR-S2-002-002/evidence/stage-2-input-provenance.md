# Stage 2 Input Provenance Check

> Phase/Stage: Exploration / Stage 2  
> Author: `P-S2-002-OIR-002-1` (`KC-EXPLORER`)  
> Created: 2026-09-28  
> Work Unit / Run: `WU-S2-002` / `OIR-S2-002-002`  
> Model / reasoning: invocation dispatch metadata says `gpt-6-luna` / `high`; active runtime model is not independently exposed  
> Session: not exposed  
> Target revision: `02e6bc77aea8630afb3d4ffdff25a20d5af3a535`  
> Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`

The assignment-defining checksums in `invocation.md` match the files read for this Run:

| Input | SHA-256 | Result |
|---|---|---|
| `.harscode-spaces/participant-profiles/profiles.md` | `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32` | Match |
| `.harscode-spaces/authority-map.md` | `65a73a5ab1233ff98699f425222d3cc9679e8c34176a9fcd3e1f60c9e20d433c` | Match |
| `TP-S2-002-007/techplan.md` | `441577782f7600fd5507d228c4dd6491009c32d6e8b81607b28a38a5ebb1a1a3` | Match |
| `OIR-S2-002-001/resolution-brief.md` | `3d2b22f6fca9bd23024eca41cda79417bcc5cf929513223867d73ddcc6440782` | Match |

The current Techplan frontmatter says `Status: Approved`. `.harscode-spaces/s2-guest-donation-truthful-state/events.md` records Human approval on 2026-09-27 and a status-only reconciliation Run changing that field from `Draft / In Review` to `Approved` on 2026-09-27. Techplan §11 still contains a stale historical sentence describing the then-Draft/In-Review state. The invocation, current frontmatter, and event provenance agree that TP-007 is approved; the stale sentence is not treated as a current gate. Its approval does not resolve O3–O5 or accept residual Security/PII risk.

At inspection the repository working tree was clean and `HEAD` was `02e6bc77aea8630afb3d4ffdff25a20d5af3a535`. This Run has written only its own `evidence/` files. No tests, runtime checks, or API validation were run; no implementation or contract behavior is claimed as verified.
