import { QRCodeSVG } from "qrcode.react";

export interface QrCodeProps {
  value: string;
}

/**
 * Thin wrapper around `qrcode.react`'s `QRCodeSVG` (techplan account/06-
 * mfa-totp, D3) — SVG output needs no jsdom canvas polyfill, unlike a
 * canvas-based QR library. `size` is a placeholder starting point (`TBD
 * — verify` final value against `design-guidelines.md` spacing once
 * built), not yet confirmed.
 */
export function QrCode({ value }: QrCodeProps) {
  return <QRCodeSVG value={value} size={200} role="img" aria-label="Kode QR untuk aktivasi MFA" />;
}
