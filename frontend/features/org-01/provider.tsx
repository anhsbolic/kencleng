'use client';
import { createContext, useContext, useEffect, useState, type ReactNode } from 'react';
import { useRouter } from 'next/navigation';
const MockContext = createContext(false);
export const useMock = () => useContext(MockContext);
export function Provider({ children }: {
    children: ReactNode;
}) {
    const [ready, setReady] = useState(!(process.env.NODE_ENV === 'development' && process.env.NEXT_PUBLIC_MSW_ENABLED === 'true')), [mock, setMock] = useState(false), [error, setError] = useState(false);
    useEffect(() => {
        if (process.env.NODE_ENV === 'development' && process.env.NEXT_PUBLIC_MSW_ENABLED === 'true') {
            import('./mocks/browser').then(async ({ worker }) => { await worker.start({ quiet: true, onUnhandledRequest: 'bypass', serviceWorker: { url: '/mockServiceWorker.js' } }); setMock(true); setReady(true); }).catch(() => setError(true));
        }
    }, []);
    return <MockContext.Provider value={mock}>{mock && <aside className="mock-banner">MOCK_DEVELOPMENT · Simulasi lokal, hanya data sintetis. Muat ulang mengatur ulang sesi simulasi. Privasi belum siap untuk production.</aside>}{error ? <main><p role="alert">Simulasi tidak dapat dimulai. Muat ulang setelah worker tersedia. Tidak ada peralihan otomatis ke API nyata.</p></main> : ready ? children : <main><p role="status">Menyiapkan lingkungan pengembangan…</p></main>}</MockContext.Provider>;
}
export function useGoogleNavigation() {
    const mock = useMock(), router = useRouter();
    return async (url: string, proof: boolean) => {
        const target = new URL(url);
        if (target.protocol !== 'https:' || target.hostname !== 'accounts.google.com' || target.pathname !== '/o/oauth2/v2/auth' || target.port || target.username || target.password)
            throw Error('Invalid authorization destination');
        if (process.env.NODE_ENV === 'development' && process.env.NEXT_PUBLIC_MSW_ENABLED === 'true' && mock) {
            const { completeMockAuth } = await import('./mocks/browser');
            completeMockAuth(proof);
            router.push(proof ? '/account-deletion/review' : '/auth/complete');
        }
        else
            window.location.assign(url);
    };
}
export function MockControls() { const mock = useMock(); if (process.env.NODE_ENV !== 'development' || process.env.NEXT_PUBLIC_MSW_ENABLED !== 'true')
    return null; return mock ? <details className="mock-controls"><summary>Skenario simulasi pengembangan</summary><label>Respons berikutnya <select aria-label="Skenario simulasi" defaultValue="normal" onChange={async (e) => { const { setScenario } = await import('./mocks/browser'); setScenario(e.target.value as Parameters<typeof setScenario>[0]); }}>{['normal', 'unknown-create', 'denied', 'unavailable', 'pending-delete', 'unknown-delete', 'scope-changed', 'manual-review'].map(s => <option key={s}>{s}</option>)}</select></label><p>Ini mengubah respons mock, bukan fakta backend.</p></details> : null; }
