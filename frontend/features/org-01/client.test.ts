import { afterEach, it, expect, vi } from 'vitest';
import { request, message, types } from './client';
import * as schemas from './generated/validation';
afterEach(() => vi.unstubAllGlobals());
it('validates exact complete responses and rejects malformed/unknown without retries', async () => { const accountId = crypto.randomUUID(), key = '019a3438-9000-7000-8000-000000000001'; const good = { data: { requestId: key, organization: { id: crypto.randomUUID(), name: 'K', createdAt: '2026-11-02T00:00:00Z' }, owner: { accountId, role: 'Owner', scope: 'internal' } } }; for (const [status, body, media, kind] of [[201, good, 'application/json', 'ok'], [201, { data: { organization: good.data.organization } }, 'application/json', 'unknown'], [202, good, 'application/json', 'unknown'], [503, { type: 'urn:uuid:f67a74ed-f3a2-45bb-a2ee-8b14e9c4c96c', title: 'Result unknown', status: 503 }, 'application/problem+json', 'problem'], [401, { type: types.unauthenticated, title: 'Authentication required', status: 403 }, 'application/problem+json', 'unknown'], [200, good, 'text/plain', 'unknown']] as const) {
    const fetcher = vi.fn().mockResolvedValue(new Response(JSON.stringify(body), { status, headers: { 'Content-Type': media } }));
    vi.stubGlobal('fetch', fetcher);
    expect((await request('createOrganization', { body: { name: 'K' }, key, csrf: 'a' })).kind).toBe(kind);
    expect(fetcher).toHaveBeenCalledTimes(1);
    expect(fetcher.mock.calls[0][0]).toBe('/api/organizations');
    expect(fetcher.mock.calls[0][1].redirect).toBe('error');
} });
it('handles204, bounds body, bad JSON, transport errors and encodes cursor only', async () => { vi.stubGlobal('fetch', vi.fn().mockResolvedValue(new Response(null, { status: 204 }))); expect(await request('logoutSession', { body: {} })).toEqual({ kind: 'ok', data: undefined }); for (const body of ['{', 'x'.repeat(65537)]) {
    vi.stubGlobal('fetch', vi.fn().mockResolvedValue(new Response(body, { headers: { 'Content-Type': 'application/json' } })));
    expect(await request('getSession')).toEqual({ kind: 'unknown' });
} vi.stubGlobal('fetch', vi.fn().mockRejectedValue(Error('network'))); expect(await request('getSession')).toEqual({ kind: 'unknown' }); const f = vi.fn().mockResolvedValue(new Response(JSON.stringify({ data: [] }), { headers: { 'Content-Type': 'application/json' } })); vi.stubGlobal('fetch', f); await request('listOrganizations', { cursor: 'a+/b' }); expect(f.mock.calls[0][0]).toBe('/api/organizations?cursor=a%2B%2Fb'); });
it('all18 exact accepted problems remain typed; unrecognized statuses/types stay unknown', async () => { for (const [key, v] of Object.entries(schemas)) {
    if (!key.startsWith('Problem'))
        continue;
    const shape = (v as typeof schemas.ProblemUnauthenticated).shape;
    const body = { type: shape.type.value, title: shape.title.value, status: shape.status.value, ...(key === 'ProblemValidation' ? { errors: [{ field: 'name', detail: 'A name is required.' }] } : {}) };
    expect((v as typeof schemas.ProblemUnauthenticated).safeParse(body).success).toBe(true);
    const text = message({ kind: 'problem', problem: body as never, wait: 3 });
    expect(text.length).toBeGreaterThan(15);
} expect(message({ kind: 'ok', data: null })).toBe(''); expect(message({ kind: 'unknown' })).toContain('belum dapat dipastikan'); const f = vi.fn().mockResolvedValue(new Response(JSON.stringify({ type: types.rate, title: 'Too many requests', status: 429 }), { status: 429, headers: { 'Content-Type': 'application/problem+json', 'Retry-After': '3' } })); vi.stubGlobal('fetch', f); expect(await request('createOrganization')).toMatchObject({ kind: 'problem', wait: 3 }); });
