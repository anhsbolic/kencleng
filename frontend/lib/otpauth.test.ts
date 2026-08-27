import { describe, expect, it } from "vitest";
import { parseOtpauthSecret } from "./otpauth";

describe("parseOtpauthSecret", () => {
  it("extracts the secret from a well-formed otpauth:// URI (R24)", () => {
    const uri =
      "otpauth://totp/Kencleng:donatur%40example.com?secret=JBSWY3DPEHPK3PXP&issuer=Kencleng";

    expect(parseOtpauthSecret(uri)).toBe("JBSWY3DPEHPK3PXP");
  });

  it("returns null when the secret param is missing", () => {
    const uri = "otpauth://totp/Kencleng:donatur%40example.com?issuer=Kencleng";

    expect(parseOtpauthSecret(uri)).toBeNull();
  });

  it("returns null (never throws) for a malformed/non-URI string", () => {
    expect(parseOtpauthSecret("not-a-uri")).toBeNull();
    expect(parseOtpauthSecret("")).toBeNull();
  });

  it("returns null when the secret param is present but empty", () => {
    const uri = "otpauth://totp/Label?secret=&issuer=Kencleng";

    expect(parseOtpauthSecret(uri)).toBeNull();
  });
});
