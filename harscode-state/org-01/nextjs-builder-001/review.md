# ORG-01 — Independent Frontend Review r1

**Disposition: no unresolved material source finding identified at the exact snapshot below. Runtime/rendered verification remains BLOCKED; this is not complete feature acceptance.** Mode `MOCK_DEVELOPMENT`. Reviewer `/root/reviewer` inspected effective originals and implementation independently, authored no production code, and owns only this report. Native sessions share workspace/tools/permissions; independence means separate assessment context and author separation, not native permission/process isolation.

## Sources and exact revision

- Baseline project HEAD `88a0c4bc1f60424a0012b1a8205cd1c69ad1b0f4`, branch `pilot/station5-c1-space-only`; working changes remain uncommitted.
- Original effective `../rancang-002/handoff.md`, `requirements.md`, `interface-decisions.md`, `experience.md`, Product ORG-01/PD-PRIV and current project/frontend AGENTS, frontend/T04/T06/T07, Design patterns/guidelines/assets read directly. Accepted API bytes independently matched SHA256 `a2d21ee3ca4ea93c7e691a53255ab825b784da00f45baeeb88115e76564df9aa`. API-adoption decision read: approval does not permit current `api/` writes.
- Reviewed changed frontend snapshot: **28 files**, deterministic SHA256 **`e7f56e56f75b190dc980e9e37725059d1bc17d7ea021be170e806cc2250c59e1`**. Algorithm: sorted paths, concatenate each file's `sha256 + two spaces + repo-relative path + LF`, hash UTF-8 bytes. Full per-file identity below. Later change invalidates affected review edges.

## Findings and correction disposition

| ID / severity / accountable owner | Initial consequence and required correction | Final targeted recheck |
| --- | --- | --- |
| R-01 / MEDIUM / Engineer | Generator used UTF-16 `.max(200)`, rejecting valid200 astral-codepoint names and valid paired success. | Corrected generator to codepoint length; generated response accepts200/rejects201. Independently executed Node validator probe and inspected meaningful boundary tests. |
| R-02 / MEDIUM / Engineer | Offset-string mock cursor contradicted accepted base64url `[createdAt,id]` and live keyset ordering. | Corrected base64url tuple, descending timestamp/ID and position comparison independent of surviving cursor row. Source rechecked. |
| R-03 / LOW / Engineer | Mock repeat no-authority logout returned401 instead of204. | Mock now returns204 when authority is absent; source rechecked. This does not verify real cookie/session behavior. |
| R-04 / MEDIUM / Engineer | Old asynchronous list/detail could repopulate private state after account switch/logout. | Principal/epoch guards, response-owner checks, no read with null principal, busy-disabled private actions now present. Producer test covers deferred detail after logout and marker/account mismatch; additional interleavings assessed at source, not browser. |
| R-05 / LOW / Engineer | Mock422 used unsupported `errors.message`, then unsupported detail; generator omitted errors minItems1. | Exact allowed detail values, generated Problem union and array min1 now used. Boundary test rejects empty errors; source rechecked. |

Additional source notes resolved: list-unavailable fixture now emits exact operation-allowed `ProblemUnavailable`; mock startup failure now produces explicit fail-closed error; no real-provider fallback. Craft's staged reauthentication warning, review/edit focus targets, detail announcement, semantic full DRAFT notice and removal of product-facing work ID were inspected in final source. These source fixes do not prove actual rendering/accessibility.

## Eleven-operation and state review

| Operation(s) | Source assessment / evidence boundary |
| --- | --- |
| `createGoogleAttempt` | Explicit normal versus same-account deletion-purpose request, pre-Google summary/full draft notice, guarded Google HTTPS destination. |
| `completeGoogleAttempt` | Generated backend operation preserved; fixed frontend normal/proof/failure continuations exist and perform no automatic create/delete. Mock uses explicit synthetic auth helper; callback/OIDC/cookie protocol NOT VERIFIED. |
| `getSession` | Trusted API account/CSRF in memory; account mismatch clears marker/private state. Session200 not treated as automatic mutation replay permission. |
| `refreshSession` | Explicit user action, bounded refresh then principal recheck, browser lock where available. No auto create/delete resend or blind refresh retry. |
| `logoutSession` | Confirmed versus uncertain revocation copy, private-memory/marker clearing, asynchronous stale-read guards. |
| `createOrganization` | Explicit canonical-name/current-account/Owner review; complete pair plus account/request/name consistency needed for success; same key/account/payload retry. Semantic validation stays field-associated. |
| `getOrganizationCreationResult` | Account checked first, exact paired receipt resolves;404/expiry/lost payload remains truthful uncertainty, no automatic new key. |
| `listOrganizations` | Current scoped discovery, cursor pagination, distinct IDs for same names, owner consistency and stale-response guard. |
| `getOrganization` | Current pair read, owner consistency and stale-response guard; cached history not promoted as authorization. |
| `createAccountDeletionPreview` | Explicit fresh full-scope preparation, empty scope supported, IDs/names and consequences before confirmation; proof/manual/scope errors do not imply deletion. |
| `deleteAccount` | Explicit closure true, fresh account check, expiry check; prepared/pending/unknown/active-erased/copies states separate; no automatic confirmation replay, old-account status lookup or login. |

All18 exact full UUID problem literals matched the accepted API set independently, with no missing/extra literals. Generated per-operation schemas constrain actual HTTP status and required body status. No title/status-only/UUID-fragment decision parsing observed. Unknown/partial/malformed success remains unknown. Only same-tab `{accountId,requestId}` persists; name/CSRF/auth/provider secrets are not persisted. Real backend authorization remains external authority.

## Executed evidence assessed

- **Reviewer-executed:** source/hash/literal comparison; actual generated-validator probe for200/201 astral names and status mismatch; seven actual client parser fixtures via transpiled `/tmp` module and stubbed fetch. Complete paired201 accepted; partial201,202,unknown UUID503,status mismatch401,wrong media201 stayed unknown; exact result-unknown503 typed. Exactly one fetch per fixture. No production edits or backend call.
- **Engineer-executed, test sources and coverage summary inspected:**23 tests in5 files passed; affected statement coverage361/423=85.34%, branch265/322=82.29%. Denominator includes attempt/client/workspace/deletion/provider and privacy page; generated outputs, test doubles/MSW and route pass-through excluded with declared rationale. Provider remains included despite lower individual coverage. Tests assert consequences/replay/uncertainty/privacy, not only component mounting or line count. This is a scoped module floor, not every-file80% or browser proof.
- **Engineer-reported commands:** exact-source generation, typecheck, lint passed (vendor worker unused-directive warning); webpack production build passed. Build is compile/static-generation evidence, not runtime or deployment. Default Turbopack build was blocked by worker port permission; offline system fonts used for successful webpack build.
- **Runtime BLOCKED:** dev server including exact approved command fails `listen EPERM 127.0.0.1:3000`; no HTTP application runtime observed. Craft's Chromium failed sandbox host; Coordinator's CUA inventory had no apps/browsers and unavailable in-app browser. Reviewer did not run prohibited escalation or claim renderer access.

## Remaining limitations and next accountable checks

Engineer owns rerunning local mock HTTP and actual browser flows in an environment permitting loopback/browser: start→notice→explicit login/review/create, same names/distinct IDs, unknown→receipt/same-key recovery, current list/detail/pagination, proof→full deletion scope→explicit confirm/pending/unknown; desktop and320px, keyboard/focus/assistive technology, worker startup failure and production mock exclusion. **EXP08/render/responsive/keyboard/real-browser MSW are NOT VERIFIED.** Source focus/ARIA/CSS and jsdom visibility cannot clear them.

Mocks intentionally do not establish actual Google proof, cookie/session/CSRF security, refresh replay, UUIDv7 time admission/receipt retention, durable journal/fence, database ownership/atomicity, erasure, retention or restore guarantees. Backend/API conformance and real integration remain outside this assignment. Production source gates mocks on development plus explicit flag, but live production-browser exclusion was not observed.

Arial/Georgia are disclosed provisional offline fallbacks for Instrument Sans/Newsreader; no visual equivalence or new canonical typography is claimed. Install reported29 dependency advisories; separate audit failed registry DNS `EAI_AGAIN`, so current attribution/validity/remediation remains unresolved, not a verified-clean dependency assessment. Broad fixes were not authorized or applied.

PDEP-R2-01 privacy/provider/final notice/intake/journal/retention/restore gates remain HELD before real-person/public/production use. This review gives no Human acceptance, release/deployment approval, real integration, blanket security assurance or canonical API adoption.

## Snapshot identities

```text
c1b56de752729e87a5cf13312b81198e5ce60831131db614a741a840253f90ad  frontend/README.md
aa9d6075cb023101f959e5af31f94d572ba1fb2b3d6792a5cadb4137a3f48ceb  frontend/app/account-deletion/review/page.tsx
633ac067134344d3beadb8f45df335ff91dfa96f71766fa25772ffa0fe76ef80  frontend/app/auth/complete/page.tsx
69c42fb35799e55f33fe02438fb62c46d70d35383869964d60b4c276b57dd515  frontend/app/auth/failure/page.tsx
c3185c9eb1d2b06b8b062c37c951e27837326481be4e04d38fe880b595e51713  frontend/app/globals.css
cc692d5438ce794927b25b6f1fb76258ae1c5f12e3f2928fe8d408ff49753f2e  frontend/app/layout.tsx
4134df2daecba2a01a896ec01f3b9fed9843cae574654208e6522e580a1f56b2  frontend/app/organizations/page.tsx
daae16df923e0bf83f016b8d8d00be982f070f27d24fdbd6347ee074c8f886ac  frontend/app/page.tsx
4cc6eaca843e40b641740b34ee30c68f481d6cb83a11ce76d127549554912b1e  frontend/app/privacy/page.tsx
bd29edd5e2efb2364746b527399582b3cf47e729738db80d5678f8b9f8735307  frontend/features/org-01/attempt.test.ts
085bdeab5d6a64e58fe280b1bb0297aae0fd759dc2a1da69f3c302662859fbc5  frontend/features/org-01/attempt.ts
a359669c6bc13920946b3a837bbfe2ba21d2b159e1b6c67c75e130481ec867c2  frontend/features/org-01/client.test.ts
693bab64a562c546467f7a62f12ebf988ab5f01cf4cb10889a07e3470f755f09  frontend/features/org-01/client.ts
57ce5af5621a92b14172ec3949c6b2e830a1319295b03f72f75d689389572585  frontend/features/org-01/deletion.tsx
a041e316e25f9f1a0243fc2fb30254f6f315d5e7a35586fa2a2b7b31841501ac  frontend/features/org-01/generated/api.ts
102978d398fa188aa17f245a3a5f3297df054f0962c836771b28a893bb333f13  frontend/features/org-01/generated/validation.ts
c509e510879be701177d8f5b4f3d17acddcc8decadc243ba013564339e86150d  frontend/features/org-01/mocks/browser.ts
126d5f8fa94b74ae459039e908c303c001d5c73634c433f4fbedc2c47dfe0bd6  frontend/features/org-01/mocks/handlers.ts
883baa37e172f5cd59ef9711cc1057693c145af2856571202241dd699db176b1  frontend/features/org-01/privacy.test.tsx
a190b1e4ad0b5c6862ef6a3573eaa6a902a1c05268792f906cbd17a96fec1d5f  frontend/features/org-01/provider.test.tsx
8f762ec2eb31c7cad9eb76fb00955f5b719e9f9ac8da458d2a1d1604b41b6038  frontend/features/org-01/provider.tsx
2818e02c1aaf659d7766ba4344044dc0a7fe13834d095d416c112a0c2c17aaaf  frontend/features/org-01/workspace.test.tsx
74e82d59d88a3e14d7606679fdec07cb9ff7c72b6796853eec8837258b288e61  frontend/features/org-01/workspace.tsx
bf73efff8648b962e09c6a339b7a066b294994860f478a03d6717a6f6adcb986  frontend/package-lock.json
d59c0821c6ff4f31fb85b74b42fba645c756f84df7ef6416a6640473d58aec56  frontend/package.json
26171599d94d5445297303f3a1f1cedec26835ea3d0057dcfa92ec520cb35cca  frontend/public/mockServiceWorker.js
eaba57df1433b6fa33bc3685ee927a7c5de248d711ed7fa7238bcdbd819cd413  frontend/scripts/generate-org-contract.mjs
fb0f11f336a031b2dfe015aa98b84d5e58df973ff71dd4e80cf5ff16d9d8cdc5  frontend/vitest.config.ts
```
