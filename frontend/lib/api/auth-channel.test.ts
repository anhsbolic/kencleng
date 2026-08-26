import { afterEach, beforeEach, describe, expect, it } from "vitest";
import { FakeBroadcastChannel, installFakeBroadcastChannel } from "@/mocks/fake-broadcast-channel";
import { postAuthChannelMessage, subscribeAuthChannel, type AuthChannelMessage } from "./auth-channel";

// Must match auth-channel.ts's own private CHANNEL_NAME constant.
const CHANNEL_NAME = "kencleng-auth";

let restore: () => void;

beforeEach(async () => {
  restore = await installFakeBroadcastChannel();
});

afterEach(() => {
  restore();
});

describe("postAuthChannelMessage / subscribeAuthChannel", () => {
  it("no-ops (does not throw) when BroadcastChannel is unavailable", async () => {
    restore();
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    delete (globalThis as any).BroadcastChannel;
    const { __resetAuthChannelForTests } = await import("./auth-channel");
    __resetAuthChannelForTests();

    expect(() => postAuthChannelMessage({ type: "logged-out" })).not.toThrow();
    const unsubscribe = subscribeAuthChannel(() => {});
    expect(() => unsubscribe()).not.toThrow();
  });

  it("delivers a posted message to a listener on a separate channel instance (simulating another tab)", () => {
    const received: AuthChannelMessage[] = [];
    const otherTabChannel = new FakeBroadcastChannel(CHANNEL_NAME);
    otherTabChannel.addEventListener("message", (event) =>
      received.push(event.data as AuthChannelMessage)
    );

    postAuthChannelMessage({ type: "refreshed", accessToken: "abc" });

    expect(received).toEqual([{ type: "refreshed", accessToken: "abc" }]);
  });

  it("subscribeAuthChannel's handler receives a message sent from another instance", () => {
    const received: AuthChannelMessage[] = [];
    subscribeAuthChannel((msg) => received.push(msg));

    const otherTabChannel = new FakeBroadcastChannel(CHANNEL_NAME);
    otherTabChannel.postMessage({ type: "logged-out" });

    expect(received).toEqual([{ type: "logged-out" }]);
  });

  it("stops delivering after the returned unsubscribe function is called", () => {
    const received: AuthChannelMessage[] = [];
    const unsubscribe = subscribeAuthChannel((msg) => received.push(msg));
    unsubscribe();

    const otherTabChannel = new FakeBroadcastChannel(CHANNEL_NAME);
    otherTabChannel.postMessage({ type: "logged-out" });

    expect(received).toEqual([]);
  });
});
