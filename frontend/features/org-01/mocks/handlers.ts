import { http, HttpResponse, delay } from 'msw';
import type { Schema } from '../client';
import * as v from '../generated/validation';
import { canonicalName, validName } from '../attempt';
const accountId = 'a1000000-0000-4000-8000-000000000001';
const csrfToken = 'A'.repeat(43);
let signedIn = false, proof = false;
let organizations: Schema<'OrganizationPair'>[] = [];
const receipts = new Map<string, Schema<'CreationResponse'>>();
let preview: Schema<'DeletionPreviewResponse'> | undefined;
export type Scenario = 'normal' | 'unknown-create' | 'denied' | 'unavailable' | 'pending-delete' | 'unknown-delete' | 'scope-changed' | 'manual-review';
let scenario: Scenario = 'normal';
export function setScenario(s: Scenario) { scenario = s; }
export function completeMockAuth(deletionProof: boolean) { signedIn = true; proof = deletionProof; preview = undefined; }
const json = (body: unknown, status = 200, extra?: Record<string, string>) => HttpResponse.json(body as Record<string, unknown>, { status, headers: { 'Cache-Control': 'no-store', ...extra } });
function problem(schema: {
    parse: (v: unknown) => unknown;
}, body: unknown) { const p = schema.parse(body) as {
    status: number;
}; return HttpResponse.json(p as {
    status: number;
}, { status: p.status, headers: { 'Content-Type': 'application/problem+json', 'Cache-Control': 'no-store' } }); }
const unauth = () => problem(v.ProblemUnauthenticated, { type: 'urn:uuid:0c56be0a-40a6-45c9-bb49-c20b075ba28b', title: 'Authentication required', status: 401 });
const absent = () => problem(v.ProblemResultNotObserved, { type: 'urn:uuid:cbf15221-1fdb-47ca-8b8a-8285b9e031b1', title: 'Result not observed', status: 404 });
const unknown = () => problem(v.ProblemResultUnknown, { type: 'urn:uuid:f67a74ed-f3a2-45bb-a2ee-8b14e9c4c96c', title: 'Result unknown', status: 503 });
function auth(request: Request) { if (!signedIn)
    return unauth(); if (scenario === 'denied' || request.headers.get('X-CSRF-Token') !== csrfToken)
    return problem(v.ProblemForbidden, { type: 'urn:uuid:142cb767-7ce1-403f-88b4-b1fe7c6c2f82', title: 'Request forbidden', status: 403 }); }
export const handlers = [
    http.post('/api/auth/google/attempts', async ({ request }) => { const body = await request.json(); const parsed = v.GoogleAttempt.safeParse(body); if (!parsed.success)
        return HttpResponse.error(); if (parsed.data.purpose) {
        const denied = auth(request);
        if (denied)
            return denied;
    } return json(v.AttemptResponse.parse({ data: { authorizationUrl: 'https://accounts.google.com/o/oauth2/v2/auth?scope=openid&state=synthetic&code_challenge=synthetic' } }), 201); }),
    http.get('/api/auth/session', () => signedIn ? json(v.SessionResponse.parse({ data: { accountId, csrfToken } })) : unauth()),
    http.post('/api/auth/refresh', ({ request }) => auth(request) ?? new HttpResponse(null, { status: 204, headers: { 'Cache-Control': 'no-store' } })),
    http.post('/api/auth/logout', ({ request }) => { const denied = signedIn ? auth(request) : undefined; if (denied)
        return denied; signedIn = false; proof = false; return new HttpResponse(null, { status: 204, headers: { 'Cache-Control': 'no-store' } }); }),
    http.post('/api/organizations', async ({ request }) => { const denied = auth(request); if (denied)
        return denied; await delay(150); const body = v.CreateOrganization.safeParse(await request.json()); if (!body.success)
        return HttpResponse.error(); if (!validName(body.data.name))
        return problem(v.ProblemValidation, { type: 'urn:uuid:1b1d2eb4-2c68-4809-8910-f40bd94e1e42', title: 'Validation failed', status: 422, errors: [{ field: 'name', detail: !canonicalName(body.data.name) ? 'A name is required.' : [...canonicalName(body.data.name)].length > 200 ? 'Name exceeds 200 Unicode code points.' : 'Name contains a control character.' }] }); const key = request.headers.get('Idempotency-Key') ?? ''; const name = canonicalName(body.data.name); const existing = receipts.get(key); if (existing)
        return existing.data.organization.name === name ? json(existing) : problem(v.ProblemIdempotencyMismatch, { type: 'urn:uuid:144ea3b4-83cc-4b0f-961a-892f3fc09b86', title: 'Request identity conflict', status: 409 }); if (scenario === 'unavailable')
        return unknown(); const pair = { organization: { id: crypto.randomUUID(), name, createdAt: new Date().toISOString() }, owner: { accountId, role: 'Owner' as const, scope: 'internal' as const } }; const receipt = v.CreationResponse.parse({ data: { requestId: key, ...pair } }); organizations.unshift(pair); receipts.set(key, receipt); if (scenario === 'unknown-create')
        return unknown(); return json(receipt, 201, { 'Location': '/organizations/' + pair.organization.id }); }),
    http.get('/api/organizations', ({ request }) => { if (!signedIn)
        return unauth(); if (scenario === 'unavailable')
        return problem(v.ProblemUnavailable, { type: 'urn:uuid:bc32fcd9-a6ce-4c42-bfa2-d09431017922', title: 'Service unavailable', status: 503 }); const url = new URL(request.url); const sorted = [...organizations].sort((a, b) => b.organization.createdAt.localeCompare(a.organization.createdAt) || b.organization.id.localeCompare(a.organization.id)); let start = 0; const cursor = url.searchParams.get('cursor'); if (cursor) {
        try {
            const [at, id] = JSON.parse(atob(cursor.replace(/-/g, '+').replace(/_/g, '/')));
            const next = sorted.findIndex(p => p.organization.createdAt < at || (p.organization.createdAt === at && p.organization.id < id));
            start = next < 0 ? sorted.length : next;
        }
        catch {
            return HttpResponse.error();
        }
    } const page = sorted.slice(start, start + 2); const last = page.at(-1)?.organization; return json(v.OrganizationList.parse({ data: page, ...(last && start + 2 < sorted.length ? { meta: { nextCursor: btoa(JSON.stringify([last.createdAt, last.id])).replace(/=/g, '').replace(/\+/g, '-').replace(/\//g, '_') } } : {}) })); }),
    http.get('/api/organizations/:organizationId', ({ params }) => { if (!signedIn)
        return unauth(); const pair = organizations.find(p => p.organization.id === params.organizationId); return pair ? json(v.OrganizationResponse.parse({ data: pair })) : problem(v.ProblemResourceUnavailable, { type: 'urn:uuid:5ed0b852-a063-4c56-85a7-a0d7c1775000', title: 'Resource unavailable', status: 404 }); }),
    http.get('/api/organization-creation-requests/:requestId', ({ params }) => { if (!signedIn)
        return unauth(); const result = receipts.get(String(params.requestId)); return result ? json(result) : absent(); }),
    http.post('/api/account-deletion-previews', ({ request }) => { const denied = auth(request); if (denied)
        return denied; if (!proof)
        return problem(v.ProblemDeletionProofRequired, { type: 'urn:uuid:0ddfd153-87b2-4619-a11e-f48418401c21', title: 'Fresh account proof required', status: 409 }); if (scenario === 'manual-review')
        return problem(v.ProblemManualReviewRequired, { type: 'urn:uuid:5a799183-67f6-49a9-90cd-ae9e06af1cf2', title: 'Manual review required', status: 422 }); preview = v.DeletionPreviewResponse.parse({ data: { previewId: crypto.randomUUID(), expiresAt: new Date(Date.now() + 600000).toISOString(), organizations: organizations.map(p => ({ id: p.organization.id, name: p.organization.name })).sort((a, b) => a.id.localeCompare(b.id)), effect: 'closeContextsAndDeleteAccount', backupCopies: 'pendingExpiry' } }); return json(preview, 201); }),
    http.post('/api/account-deletions', async ({ request }) => { const denied = auth(request); if (denied)
        return denied; const parsed = v.DeleteAccount.safeParse(await request.json()); if (!parsed.success || parsed.data.previewId !== preview?.data.previewId || Date.now() >= Date.parse(preview.data.expiresAt))
        return problem(v.ProblemDeletionPreviewInvalid, { type: 'urn:uuid:25c9699d-d474-496b-8214-65ceb830f197', title: 'Deletion preview invalid', status: 409 }); if (scenario === 'scope-changed')
        return problem(v.ProblemDeletionScopeChanged, { type: 'urn:uuid:bd8b4ade-aaef-4ea9-9d73-2585c7b31ea6', title: 'Deletion scope changed', status: 409 }); if (scenario === 'unknown-delete')
        return unknown(); if (scenario === 'pending-delete') {
        signedIn = false;
        return problem(v.ProblemDeletionPending, { type: 'urn:uuid:6d143e2a-e3ba-4e71-8c67-1c8f7d4bae85', title: 'Deletion pending', status: 503 });
    } signedIn = false; organizations = []; receipts.clear(); preview = undefined; return json(v.DeletionResponse.parse({ data: { accessEnded: true, activeDataErased: true, remainingCopies: 'pendingVerification' } })); })
];
