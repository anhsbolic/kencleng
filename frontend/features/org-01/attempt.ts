export type Marker = {
    accountId: string;
    requestId: string;
};
export type Attempt = Marker & {
    name?: string;
};
const storageKey = 'kencleng.org-01.pending';
export function canonicalName(name: string) { return name.replace(/^\p{White_Space}+|\p{White_Space}+$/gu, ''); }
export function validName(name: string) { const n = canonicalName(name); return [...n].length >= 1 && [...n].length <= 200 && !/\p{Cc}/u.test(n); }
export function requestId(now = Date.now()) {
    const bytes = crypto.getRandomValues(new Uint8Array(16));
    let time = now;
    for (let i = 5; i >= 0; i--) {
        bytes[i] = time % 256;
        time = Math.floor(time / 256);
    }
    bytes[6] = (bytes[6] & 15) | 112;
    bytes[8] = (bytes[8] & 63) | 128;
    const hex = Array.from(bytes, b => b.toString(16).padStart(2, '0')).join('');
    return `${hex.slice(0, 8)}-${hex.slice(8, 12)}-${hex.slice(12, 16)}-${hex.slice(16, 20)}-${hex.slice(20)}`;
}
export function saveMarker(marker: Marker) { try {
    sessionStorage.setItem(storageKey, JSON.stringify({ accountId: marker.accountId, requestId: marker.requestId }));
    return true;
}
catch {
    return false;
} }
export function readMarker(): Marker | null { try {
    const m = JSON.parse(sessionStorage.getItem(storageKey) ?? 'null');
    if (m && Object.keys(m).length === 2 && /^[0-9a-f-]{36}$/i.test(m.accountId) && /^[0-9a-f]{8}-[0-9a-f]{4}-7[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i.test(m.requestId))
        return m;
    return null;
}
catch {
    return null;
} }
export function clearMarker() { try {
    sessionStorage.removeItem(storageKey);
}
catch { /* Storage unavailable: memory state remains authoritative for this tab only. */ } }
export async function refreshSession(csrf: string) {
    const run = async () => {
        const { request } = await import('./client');
        const r = await request<void>('refreshSession', { body: {}, csrf });
        return r.kind === 'ok' ? request<import('./client').Schema<'SessionResponse'>>('getSession') : r;
    };
    if (navigator.locks)
        return navigator.locks.request('kencleng-session-refresh', { ifAvailable: true }, lock => lock ? run() : Promise.resolve({ kind: 'unknown' } as const));
    return run();
}
