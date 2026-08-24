import { create } from "zustand";

/**
 * In-memory access-token holder. Shape only — no login logic lives
 * here (that's Account Task #3's job); this exists so Tasks #2 and
 * #3 share one store shape instead of each inventing their own, and
 * so `lib/api/client.ts` has a real store to read from.
 *
 * Deliberately not persisted (no `persist` middleware): the access
 * token is short-lived and re-obtainable via `POST /auth/refresh`
 * (the `HttpOnly` refresh-token cookie is the actual durable
 * credential) — keeping it in-memory-only avoids leaving a token
 * readable in `localStorage`.
 */
type AuthState = {
  accessToken: string | null;
  setAccessToken: (token: string | null) => void;
  clearAccessToken: () => void;
};

export const useAuthStore = create<AuthState>((set) => ({
  accessToken: null,
  setAccessToken: (token) => set({ accessToken: token }),
  clearAccessToken: () => set({ accessToken: null }),
}));
