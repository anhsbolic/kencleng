import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Kencleng",
  description: "Kencleng research scaffold; Product discovery belum menetapkan fitur aktif.",
};

export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="id">
      <body>{children}</body>
    </html>
  );
}
