import { cleanup, render, renderHook, screen } from '@testing-library/react';
import { afterEach, it, expect, vi } from 'vitest';
import { Provider, MockControls, useGoogleNavigation } from './provider';
vi.mock('next/navigation', () => ({ useRouter: () => ({ push: vi.fn() }) }));
afterEach(cleanup);
it('keeps development mocks absent outside development and children render', () => { render(<Provider><p>App ready</p><MockControls /></Provider>); expect(screen.getByText('App ready')).toBeVisible(); expect(screen.queryByText(/Skenario simulasi/)).toBeNull(); expect(screen.queryByText(/MOCK_DEVELOPMENT/)).toBeNull(); });
it('rejects unsafe auth destinations before any navigation', async () => { const { result } = renderHook(() => useGoogleNavigation()); for (const url of ['http://accounts.google.com/o/oauth2/v2/auth', 'https://evil.test/o/oauth2/v2/auth', 'https://accounts.google.com/evil', 'https://user@accounts.google.com/o/oauth2/v2/auth', 'https://accounts.google.com:8443/o/oauth2/v2/auth'])
    await expect(result.current(url, false)).rejects.toThrow('Invalid authorization destination'); });
