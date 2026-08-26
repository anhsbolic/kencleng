// Shared test double for `BroadcastChannel` — jsdom 30.0.1 (this
// project's pinned version) doesn't implement the real API at all
// (confirmed directly against the actual pinned version — see
// techplan account/03-login-session-management's §14 Resolved #6), so
// there is no real API for `lib/api/auth-channel.ts`'s tests, or any
// other test simulating cross-tab behavior, to exercise. Models the one
// semantic that's easy to get subtly wrong: a `postMessage` call is
// never delivered back to the exact instance that sent it (real
// `BroadcastChannel` behavior) — only to *other* instances on the same
// channel name, simulating other tabs.
export class FakeBroadcastChannel {
  static instances: FakeBroadcastChannel[] = [];
  private listeners: Array<(event: MessageEvent) => void> = [];

  constructor(public name: string) {
    FakeBroadcastChannel.instances.push(this);
  }

  postMessage(data: unknown) {
    for (const instance of FakeBroadcastChannel.instances) {
      if (instance === this || instance.name !== this.name) continue;
      const event = { data } as MessageEvent;
      instance.listeners.forEach((listener) => listener(event));
    }
  }

  addEventListener(type: string, listener: (event: MessageEvent) => void) {
    if (type === "message") this.listeners.push(listener);
  }

  removeEventListener(type: string, listener: (event: MessageEvent) => void) {
    if (type === "message") {
      this.listeners = this.listeners.filter((existing) => existing !== listener);
    }
  }
}

/**
 * Installs `FakeBroadcastChannel` as `globalThis.BroadcastChannel` and
 * resets `lib/api/auth-channel.ts`'s cached singleton so it picks up
 * the fake on its next call — call from `beforeEach`. Returns a
 * restore function — call from `afterEach`.
 */
export async function installFakeBroadcastChannel(): Promise<() => void> {
  const original = globalThis.BroadcastChannel;
  FakeBroadcastChannel.instances = [];
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  (globalThis as any).BroadcastChannel = FakeBroadcastChannel;

  const { __resetAuthChannelForTests } = await import("@/lib/api/auth-channel");
  __resetAuthChannelForTests();

  return () => {
    if (original) {
      globalThis.BroadcastChannel = original;
    } else {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      delete (globalThis as any).BroadcastChannel;
    }
    __resetAuthChannelForTests();
  };
}
