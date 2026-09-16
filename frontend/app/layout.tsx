import type { Metadata } from "next";
import { Instrument_Sans, Newsreader } from "next/font/google";
import "./globals.css";

const instrumentSans = Instrument_Sans({
  subsets: ["latin"],
  display: "swap",
  variable: "--font-instrument-sans",
});

const newsreader = Newsreader({
  subsets: ["latin"],
  display: "swap",
  variable: "--font-newsreader",
});

export const metadata: Metadata = {
  title: "Kencleng — Harapan tumbuh dari hal yang jelas",
  description:
    "Kencleng adalah ruang penggalangan dana yang menempatkan cerita dan informasi berdampingan.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="id">
      <body className={`${instrumentSans.variable} ${newsreader.variable}`}>
        {children}
      </body>
    </html>
  );
}
