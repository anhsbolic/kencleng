import type { Metadata } from "next";
import { Inter, Plus_Jakarta_Sans } from "next/font/google";
import { AuthBootstrapProvider } from "@/components/providers/auth-bootstrap-provider";
import { MockingProvider } from "@/components/providers/mocking-provider";
import { QueryProvider } from "@/components/providers/query-provider";
import { SessionGuardProvider } from "@/components/providers/session-guard-provider";
import { SwRegister } from "@/components/providers/sw-register";
import "./globals.css";

const plusJakartaSans = Plus_Jakarta_Sans({
  variable: "--font-plus-jakarta-sans",
  subsets: ["latin"],
  weight: ["600", "700", "800"],
});

const inter = Inter({
  variable: "--font-inter",
  subsets: ["latin"],
  weight: ["400", "500", "600"],
});

export const metadata: Metadata = {
  title: "Kencleng",
  description: "Kencleng — transparent, curated crowdfunding.",
  manifest: "/manifest.json",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html
      lang="en"
      className={`${plusJakartaSans.variable} ${inter.variable} h-full antialiased`}
    >
      <body className="min-h-full flex flex-col">
        <MockingProvider>
          <QueryProvider>
            {/* Mounted inside QueryProvider (needs useQueryClient) and
                inside MockingProvider (so MSW is ready first in
                mock-dev mode — MockingProvider already gates all
                children behind `ready`). Techplan account/02-google-
                oauth-login-register, R13. SessionGuardProvider nested
                inside AuthBootstrapProvider — both are root-level
                session-lifecycle providers, grouped together (techplan
                account/03-login-session-management, task-02). */}
            <AuthBootstrapProvider>
              <SessionGuardProvider>{children}</SessionGuardProvider>
            </AuthBootstrapProvider>
          </QueryProvider>
        </MockingProvider>
        <SwRegister />
      </body>
    </html>
  );
}
