import { create } from "zustand";

/**
 * Client-only UI state for the landing-page login/register modal
 * (`frontend/AGENTS.md` §3 — Zustand for UI state, not server data).
 * `mode: null` means closed. `/login` and `/register` remain real,
 * directly-navigable routes independent of this store (needed for the
 * Google OAuth callback's hardcoded `/login?error={code}` redirect
 * target, and for direct/shared links) — this store only drives the
 * *additional* modal presentation triggered from the Public Shell.
 */
export type AuthModalMode = "login" | "register" | null;

type AuthModalState = {
  mode: AuthModalMode;
  openLogin: () => void;
  openRegister: () => void;
  close: () => void;
};

export const useAuthModalStore = create<AuthModalState>((set) => ({
  mode: null,
  openLogin: () => set({ mode: "login" }),
  openRegister: () => set({ mode: "register" }),
  close: () => set({ mode: null }),
}));
