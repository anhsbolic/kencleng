# ORG-01 — Independent Security Findings r1

Assignment `RANCANG-ORG-01-001`; Independent Security Specialist `/root/org01_security`; 2026-10-25. Intended consumers: Coordinator, independent Solution Reviewer and direct backend/frontend owners.

**Design assessment concluded: no unresolved blocking security finding in the exact proposals below.** SEC-01 was LOW/nonblocking and is addressed in the revised documentation. This is scoped independent document assessment, not a blanket security PASS, Human acceptance, implementation verification or deployment readiness.

## Independence, authority and evidence

This actor received a fresh bounded assignment and the workspace instructions, without the parent's prior working transcript or producer deliberations. Original Product/approval, project instructions, applicable standards, actual scaffold and returned producer files were read directly. Producer messages supplied revision pointers, not substitute evidence. This actor did not author X1/O1/O2/O3 or correct their content; correction was requested from the owning Architect and checked again in the actual files. No model/config override, subagent, network, installation, credential inspection, service action, implementation or test execution occurred. The sole write is this owned review file.

The shared filesystem and workspace-write sandbox are **not per-role filesystem isolation**. R2 is the explicitly Human-accepted bounded deviation in [state](state.md), not technical containment. R5 cold resume is untested; R6 semantic/protected/final-acceptance gates remain supervised/instructional rather than universally tool-enforced. Separate contexts and independent authorship support this assessment's limited independence claim; they do not prove process isolation or general Harness readiness.

Kencleng HEAD checked: `3faf3b8a959cdb617958568b158c0805e84107ed`; Harscode HEAD: `df21f47ecb85877a571b15f0793ced2f5c33e003`. Rancang contracts, coordination and Security agent/knowledge originals byte-match formed definition `9663a3fa73ac0101e66764d63af8711f291a66ab`. Its embedded candidate wording predates the directly read 2026-10-25 formation approval. Assignment authorization, source approval, formation approval and final bundle acceptance remain distinct.

### Exact original sources

All project sources listed here were read directly and byte-compared successfully to the Kencleng baseline above. IDs are Git blobs; paths are Kencleng-relative unless marked Harscode.

| Original | Checked blob |
| --- | --- |
| `docs/product/outcomes/org-01/product-contract.md` | `da19d8e7cdb9baac8cf4af1c5e6b61763607013e` |
| `docs/product/outcomes/org-01/approval.md` | `65211210c102afd02143ab1921b2486fb600b241` |
| `AGENTS.md`; `backend/AGENTS.md`; `frontend/AGENTS.md` | `14fd68fe8e25306270be6e3e259f181079ea9390`; `c463cf204ae5dbf32067eaf20c16ceaa76dd6bb9`; `8d5836d6450b2556c91995e57a3c408a7da6b14b` |
| `docs/project/kencleng-backend-tech-stack.md`; frontend equivalent | `f451e63af760ad07991b5c072a3530ed9296799f`; `a1cdd73b816cbeab5509a8729c8f20a3bf012b34` |
| `docs/ui-ux/README.md`; `product-design-principles.md`; `patterns.md` | `426e0d499c73e28b8ffa8af55353d107369f5afa`; `9eb9e5893fc7c1090d28024c51ed507c18cea945`; `da439232e2d17547a986c8a17914b6858b012a79` |
| `backend/cmd/server/main.go`; `backend/internal/platform/db/db.go` | `745ff37cba7471ff4d15c6e624da0d64341301f8`; `1abb2743ee6b991d5d20fb7b070c04897fe9d0fe` |
| `frontend/next.config.ts`; `frontend/app/page.tsx` | `f6e153f76d600f522cf4565ae89974010ad2bf6a`; `95ea11781e110824f4413999e25f102d1939df18` |
| `api/README.md`; `api/openapi/index.yaml` and bundled `api/openapi.yaml` | `3a7be794640f1b243c47e043b035065ba7651856`; `13fb4f31e04d64228eda04c7995548c93c2f649a` (both YAML files) |
| Harscode `AGENTS.md`; `engineering/AGENTS.md` | `7e4d63bc6ef7395480e1edf4f51c01436aef5d4f`; `270bd6a5d0c3a8f5c7752000300fc9da64fb695b` |
| Harscode Rancang `contracts.md`; `coordination.md` | `55d62117c650ad90d6e4c18942181bf5c8db72bc`; `c7ecb97ee7712b1a40e59c4b064d5149b76f7a88` |
| Harscode Rancang Security `agent.md`; `knowledge.md` | `c1b66d4ee09f0fdbb07e7142838176a7f4babdc6`; `743c48bfd0c93e3cb9c79923e71998feac0cb568` |
| Harscode Rancang `formation-approval-2026-10-25.md` | `120baaaee6dd404b5bfa4cfe79960f6b36e2bf21` |

Workspace `AGENTS.md`, assignment state and handoff were also read. State/handoff are mutable coordination records, not substitute solution authority. Initial inspected snapshots: state blob `b2069d1838b4634b8ed7930a9cf3b4ce4eed337a`; handoff blob `95d462e147b81c8cc98bfa81fd49a03a33ae8009`. No scoped AGENTS was found under project docs.

### Exact reviewed producer outputs

Full original X1 and r1 technical/experience outputs were inspected; r2 originals were reopened, with re-assessment focused on the admission-control/429 and affected experience/verification edges. Content hashes were independently computed against real files.

| Output | Git blob | SHA-256 |
| --- | --- | --- |
| [X1 v1](obligations.md), Requirements Analyst | `e1529de3001ae7d261dff30f8b047f5fcd6bb11d` | `93359cf29788af494ba6c98ca485640aaabd6e956a2837157ed2dfff3b5a04e1` |
| Initial technical r1, Architect | `9c873a536b2d7a4ce71185a006d28c7ffa26abfe` | `a6818aaa7a31205fc1d9b9901ac8b999cfe30486fa6ab5af621829cf8b4f6ef6` |
| Initial experience r1, Experience Specialist | `7d91b24dd878f524af829a0ac1cf1f3ea8dffd09` | `432703976620957f3d000a25a21ecee31c345e8ff2eb66f3ff267aeafa7cfa68` |
| Final [technical r2](technical.md), Architect | `03bcfc3d0a6af278cc40e9fa8e38abab9efcd8b7` | `10ef25ca8bfd3eb46f2bf328baa2ee7c21edcdff293747d41c6e65c056c23341` |
| Final [experience r2](experience.md), Experience Specialist | `b84e5a4e05a43e43fbc7f3ff4135be986e87fcae` | `5ff6ef4f05f4792d2aedef7b810d2d49095a0472ad8e48bd337bbc44d566eef0` |

## Threat-boundary assessment

The concrete risks are false actor attribution/Owner grant, cross-account disclosure, forged requests, replay/partial success, claim laundering and resource exhaustion. Browser input and organization names are untrusted; Go owns authentication and resource authorization; the configured OIDC issuer supplies account identity; PostgreSQL owns durable state. Operational DB/provider compromise is outside this document assessment. No money movement, public organization lookup, broad role matrix or legal verification is introduced.

| Concern / binding source | Assessment of the actual proposal; downstream obligation |
| --- | --- |
| Attribution — Product REQ-02/AC-02; root/backend authorization rules | Exact validated issuer/subject maps to a stable actor; no email linking or caller-selected Owner. Account control is bounded attribution, not proof of legal identity/representation. O3 displays the acting account before explicit effect. T-01/02 and E-01 must verify failures, substitution and account changes. |
| Authentication/session/CSRF — technical O2-A | Code/PKCE/state/nonce/browser binding, token validation, opaque hashed server session, cookie flags, rotation, local-only HTTP exception, expiry/revocation ordering, Origin plus session-bound CSRF and closed cross-origin token reads address the identified browser trust boundary. No silent anonymous fallback. Provider revocation is not instant local logout; the proposed eight-hour local session limit is explicit. Actual library/configuration enforcement remains unverified under T-01..03/08. |
| Owner/resource access — Product REQ-02/05; root/backend rules | Backend derives initial Owner from session and restricts list/detail to current owner equality. IDs/keys are not bearer authority; inaccessible/missing IDs share 404. No additional Owner privilege is inferred. Actor-scoped caches and suspended recovery on account change constrain client leakage; test direct access and delayed responses under T-02/03/07 and E-02/04. |
| Atomicity — Product REQ-04/AC-04/07 | Non-null Owner plus one transaction for organization/provenance/request record, session revalidation under lock and success only after positive commit confirmation preserve the combined result. Read Committed with actor/key serialization and database uniqueness is coherent for the bounded slice. This remains a proposed guarantee, requiring real PG fault/concurrency evidence T-04/05. |
| Replay/unknown — Product REQ-07/08; Design uncertainty rules | Actor/key namespace, exact normalized-payload comparison, retained records, conflict on different payload and original-key retries prevent the proposed retry path from silently creating twice. New key/new intent is not automatic. Empty read, denial or 429 cannot settle an earlier unknown. O3 accurately explains replay may create; current reads are separate from formation responses. T-04/05/07 and E-02/03/04 are required. |
| Minimization/provenance — Product REQ-06/AC-06/09; logging rules | Only name/necessary actor and attribution metadata are proposed; no email/avatar, legal documents, persistent provider tokens or public owner details. Name-source and internal-role meaning are explicit in stored/read/UI shapes. Strict bounded parsing, text rendering, parameterized SQL and safe errors/telemetry are required by T-06/08. Local pending intent is untrusted and must not cross accounts. The retained request metadata supports replay safety; future deletion/transfer/account-linking work must revisit affected semantics. |
| Abuse/logging — root/backend rules; SEC-01 | r2 bounds login allocation/outstanding work and formation/replay admission, preserves durable keys and uses safe aggregate signals without raw tokens/PII. Creation actor/time provide durable attribution without requiring a new audit service. Threshold effectiveness and safe telemetry require T-08 execution and deployment configuration. |

## Finding history and disposition

**SEC-01 — LOW, nonblocking: explicit transient admission control absent from initial proposal.** Affected output: technical r1 hash above, O2-A login transaction allocation and O2-B/C persistent formation/replay; root/backend proportional security rules apply. Initial evidence: r1 specified 4 KiB bodies, short transactions and ten-minute login records, but no explicit request-rate/outstanding-work bounds. A caller could repeatedly allocate login records or an authenticated caller could exhaust transactional capacity/persistent writes. This was a design omission, not an observed exploit or an unmet Product success invariant.

Required correction sent separately to Coordinator and Architect: bounded per-source login and per-actor formation/replay admission/concurrency plus bounded work lifetime; safe temporary 429/backoff preserving original key/payload and earlier unknown. Thresholds may remain downstream capacity configuration. No permanent Product quota, organization-name gate or Reviewer condition was requested. The initial finding was also shared with Experience and Solution Reviewer.

**Disposition: ADDRESSED IN DESIGN, implementation unverified.** Architect technical r2 O2-A now explicitly requires those bounds, outstanding login expiry, bounded total work, safe aggregate signals and no idempotency eviction; O2-C defines `429 temporarily_throttled`, this request `effect: none`, `Retry-After`, original-key retention and preservation of prior unknown. T-08 adds checks. Experience r2 maps login/formation throttling without permanent ineligibility or false resolution; E-01/E-03 add checks. The actual changed sections and both final hashes were independently reread/verified. Original severity/evidence are retained here. No open security design finding remains; this disposition does not certify implemented controls.

## Conditions, residual risks and limits

Direct backend/frontend design compatibility is conditional on adopting the exact proposals through authorized canonical API/type work and separately authorized authentication/privileged Owner implementation. The existing API has no product operations, backend only health/database wiring, and frontend only a neutral page/dev rewrite. Their absence is observed implementation scope, not a design finding or evidence that the proposed security controls already exist. No Engineering Standard amendment was proposed or assessed.

Before deployment, the downstream owner must supply and verify a trusted issuer/stable subject, client/redirect registration and secret handling, maintained library behavior, TLS, cookie/origin configuration, same-origin gateway behavior and admission thresholds. The development rewrite does not supply production routing. Local non-Secure cookies cannot become a production default. Provider compromise, stolen valid sessions and account loss remain residual identity risks; this slice does not promise MFA, account recovery, provider-wide logout or legal identity proof.

All T-01..08 and E-01..06 implementation, PG integration/race/fault, browser/render/accessibility and security tests are **NOT RUN**. No penetration test, dependency audit, deployed configuration inspection, provider validation or runtime readiness check was performed. These are downstream evidence/deployment gates, not retroactively fabricated design blockers. Final bundle Human acceptance is pending and separate from this finding disposition. A changed material source/output invalidates only its affected review joins; compare hashes before reusing this assessment.
