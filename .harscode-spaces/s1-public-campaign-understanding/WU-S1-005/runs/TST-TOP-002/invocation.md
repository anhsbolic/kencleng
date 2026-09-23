# TST-TOP-002 — Joint Runtime Verification Re-entry

WORK_UNIT_ID:
`WU-S1-005`

RUN_ID:
`TST-TOP-002`

RUN_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TST-TOP-002`

WORK_UNIT_PATH:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005`

ROLE:
`Verifier`

SPECIALIZATION:
Joint Testing re-entry — controlled media proxy/retraction closure

PARTICIPANT:
Codex CLI agent

SESSION:
Fresh independent joint-runtime Testing session

COMMUNICATION_LANGUAGE:
Bahasa Indonesia

SELECTED_MODEL:
`gpt-5.6-terra`

REASONING_EFFORT:
`high`

MODEL_APPROVAL:
`NOT_REQUIRED`

PRIOR_ARTIFACTS:
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TP-TOP-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TST-TOP-001/testing-report-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TP-BE-001/techplan.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/TST-BE-002/testing-report-1.md`
- `.harscode-spaces/s1-public-campaign-understanding/WU-S1-003/runs/BLD-BE-PATCH-001/patch-report-1.md`

## Objective

Close only the remaining WU-S1-005 runtime evidence that depends on the implemented WU-S1-003 backend.

Required:
1. Create/use disposable or explicitly test-only persisted Campaign/media state; do not mutate shared/manual Human data.
2. Establish one eligible public Campaign with private JPEG/PNG media and record its browser-facing same-origin `content_url`.
3. Through root Caddy, verify:
   - eligible media 200 with expected JPEG/PNG bytes and `Cache-Control: private, no-store`;
   - absent/non-public/non-member public 404 behavior/header preservation;
   - eligible dependency failure 503/header preservation when credibly inducible without production-code edits.
4. R5: confirm the seeded Campaign object is not anonymously retrievable from MinIO and public response does not expose direct object/signed URL.
5. R7: after the content_url is already known, withdraw persisted parent public eligibility and, if practical as a separate scenario, media membership using test-only persistence manipulation owned by the backend fixture. Then issue a fresh request through `localhost:8080` to the exact same content_url and confirm no new media bytes are delivered, public non-disclosure semantics remain, and no redirect/object/signed URL appears.
6. Clearly distinguish test fixture/data mutation from product/operator capability; do not add production endpoints or broaden WU-S1-003.
7. Use Podman-aware local runtime guidance from `.harscode-spaces/.local-config.yaml`.

If a required state transition cannot be exercised without a production change, stop and report the exact missing test seam instead of inventing a product API.

Do not edit production code in Testing.
Write:
`.harscode-spaces/s1-public-campaign-understanding/WU-S1-005/runs/TST-TOP-002/testing-report-1.md`

If a production defect or missing necessary test seam is found, write a patch plan and route to the owning Build Work Unit.
