import type { Metadata } from "next";
import "./globals.css";
import { Provider } from "@/features/org-01/provider";
export const metadata: Metadata = {
    title: "Kencleng",
    description: "Konteks organisasi dan tanggung jawab Owner internal Kencleng",
};
export default function RootLayout({ children, }: Readonly<{
    children: React.ReactNode;
}>) {
    return (<html lang="id">
      <body>
        <Provider>{children}</Provider>
      </body>
    </html>);
}
