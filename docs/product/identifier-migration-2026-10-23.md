# Kencleng — Product Identifier Migration (23 Oktober 2026)

> **Status:** Approved structural/refactoring direction from Anhar, applied on `pilot/station5-c1-space-only`. This file is a durable old→new ID crosswalk, **not** a new Product behavior or approval for Engineering.
> **Output catalog:** Harscode Product phases `PROD-FND/JRN/SEL/DEF/APR/HND`; output kinds `PROD-INTENT`, `PROD-JOURNEYS`, `PROD-SELECTION`, `PROD-CONTRACT`, `PROD-APPROVAL`, `PROD-HANDOFF`. Human-facing labels remain in plain language.

## Active output IDs

| Project artifact | Tracking ID | Location |
| --- | --- | --- |
| Product Intent & Value Definition | `PROD-INTENT` | [Product Intent](product-intent.md) |
| Interconnected Product Journeys | `PROD-JOURNEYS` | [Journey Map](interconnected-journeys/level-2-network.md) |
| Selected Organization outcome | `PROD-SELECTION-ORG-01` | [Outcome Selection](outcome-selection.md) |
| Approved Organization Product Contract | `PROD-CONTRACT-ORG-01` | [Product Contract](outcomes/org-01/product-contract.md) |
| Exact-revision Human approval | `PROD-APPROVAL-ORG-01` | [Approval Record](outcomes/org-01/approval.md) |
| Future consumer handoff | `PROD-HANDOFF-ORG-01` | **Not created / not authorized**; use approved Product Contract + approval record when authorized |

**Outcome key `ORG-01`** corresponds to Journey node `J-ORG-01`. They are **different kinds of identifier**: a selected bounded Product outcome versus a node on the Journey map.

## Journey node IDs — all 26 mapped

| Former node ID | Current node ID |
| --- | --- |
| `O1` | `J-ORG-01` |
| `O2` | `J-ORG-02` |
| `O3` | `J-ORG-03` |
| `O4` | `J-ORG-04` |
| `O5` | `J-ORG-05` |
| `P1` | `J-CAM-01` |
| `P2` | `J-CAM-02` |
| `P3` | `J-CAM-03` |
| `P4` | `J-CAM-04` |
| `P5` | `J-CAM-05` |
| `P6` | `J-CAM-06` |
| `P7` | `J-CAM-07` |
| `P8` | `J-CAM-08` |
| `V1` | `J-VIS-01` |
| `V2` | `J-VIS-02` |
| `V3` | `J-VIS-03` |
| `V4` | `J-VIS-04` |
| `V5` | `J-VIS-05` |
| `V6` | `J-VIS-06` |
| `V7` | `J-VIS-07` |
| `A1` | `J-ACC-01` |
| `A2` | `J-ACC-02` |
| `A3` | `J-ACC-03` |
| `A4` | `J-ACC-04` |
| `A5` | `J-ACC-05` |
| `A6` | `J-ACC-06` |

**Original approved Journey source:** blob `5783fdb303e34ecef7988d6ec3c7fb8b15326a95` for [Level 2 map at the pre-migration checkpoint](https://github.com/anhsbolic/kencleng/blob/757f715a1aca9f420b48919f03ff36e980a94ef0/docs/product/interconnected-journeys/level-2-network.md). **Current mapped Journey blob:** `650fbaddac526963d56de12e42b36b3a603d969d`. Local Mermaid node aliases such as `ORG_01` are diagram syntax only, not canonical Product tracking identifiers.

## Approved Product Contract IDs — all 26 mapped

| Category | Old range | New range |
| --- | --- | --- |
| Product Decision | `PD-O1-01..05` | `PD-ORG-01-01..05` |
| Success Criteria | `SC-O1-01..04` | `SC-ORG-01-01..04` |
| Requirement | `O1-R01..08` | `REQ-ORG-01-01..08` |
| Acceptance Criteria | `O1-AC01..09` | `AC-ORG-01-01..09` |

Expansion is **one-to-one and order-preserving**. For example `PD-O1-01 → PD-ORG-01-01`, `SC-O1-03 → SC-ORG-01-03`, `O1-R08 → REQ-ORG-01-08`, and `O1-AC08 → AC-ORG-01-08`. Existing shorthand trace references to `R01` etc. become fully qualified IDs within the new contract.

**Exact Human-approved Product content:** original approved candidate blob `448a8082bcb9ac1c7e41a686f2e3fdd2e3988ac4`; original canonical blob `0a28f7431f749e608402711be704a3f065a1398a`. **New canonical Product Contract:** `293bc59dbb738a17e763de361c412b42ca9cd565`. [New approval record](outcomes/org-01/approval.md) reconciles their authority.

The Product Contract was compared against the original Human-reviewed blob. **26/26 normative rows match after inverse ID mapping** (PD semantics, SC meanings/boundaries, REQ behavior/conditions/invariants, AC trace and expected result). No new internal `Owner`, Reviewer or Potential Conflicts behavior was introduced by this migration.

## Historical and authority discipline

Existing dated historical pilot/draft artifacts and the old exact approval remain **immutable history** via Git SHA. Do not rewrite their original codes as if those were used at the time, and do not inherit their product decisions into the current approved outcome. In current Project Product documents, use new codes only; when historical audit is necessary, consult this crosswalk.

Changing a project-local tracking identifier is not an automatic request for a new Product outcome, a new process phase, an Engineering invocation, or a change in operational `main`. Downstream reads **self-contained approved Product semantics**, not upstream code archaeology. Any future semantic revision to the approved Product Contract requires separate Human approval on the exact new revision.
