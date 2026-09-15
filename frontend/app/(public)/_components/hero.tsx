import Link from "next/link";
import { Badge } from "@/components/ui/badge";

/** Static, route-owned hero for the public home page. */
export function Hero() {
  return (
    <div className="border-b border-primary-100 bg-primary-50">
      <div className="mx-auto flex max-w-[1360px] flex-col items-start gap-4 px-4 py-10 md:gap-5 md:px-6 md:py-20">
        <Badge tone="success">Organisasi terverifikasi</Badge>
        <h1 className="max-w-2xl text-h1 font-extrabold text-neutral-900 md:text-display">
          Berbagi itu mudah, dampaknya nyata
        </h1>
        <p className="max-w-xl text-body text-neutral-700 md:text-body-lg">
          Kencleng menghubungkan Anda dengan kampanye dari organisasi yang
          sudah diverifikasi — dari musala di kampung sampai beasiswa anak
          pesisir.
        </p>
        <div className="flex w-full flex-col gap-2.5 pt-1 md:w-auto md:flex-row">
          <Link
            href="#kampanye"
            className="inline-flex h-13 items-center justify-center rounded-md bg-primary-600 px-6 text-lg font-semibold text-white shadow-sm transition-colors hover:bg-primary-700 md:h-11 md:text-base"
          >
            Mulai berdonasi
          </Link>
          <Link
            href="/register"
            className="inline-flex h-13 items-center justify-center rounded-md border border-neutral-200 px-6 text-lg font-semibold text-neutral-700 transition-colors hover:bg-neutral-100 md:h-11 md:text-base"
          >
            Galang dana untuk organisasi Anda
          </Link>
        </div>
      </div>
    </div>
  );
}
