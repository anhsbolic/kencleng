import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { HowItWorks } from "./how-it-works";

describe("HowItWorks", () => {
  it("renders the three fixed authority-bounded messages", () => {
    render(<HowItWorks />);

    expect(screen.getByText("Temukan kampanye yang sedang berjalan")).toBeInTheDocument();
    expect(
      screen.getByText("Jelajahi kampanye dari organisasi yang telah diverifikasi.")
    ).toBeInTheDocument();
    expect(screen.getByText("Pilih metode donasi")).toBeInTheDocument();
    expect(
      screen.getByText(
        "Transfer bank, kartu debit, e-wallet, atau QRIS tersedia untuk dipilih saat berdonasi."
      )
    ).toBeInTheDocument();
    expect(screen.getByText("Pantau status donasi")).toBeInTheDocument();
    expect(
      screen.getByText("Setelah donasi dibuat, cek statusnya melalui tautan yang tersedia.")
    ).toBeInTheDocument();
  });

  it("does not render removed money, fee, or reporting promises", () => {
    render(<HowItWorks />);

    expect(screen.queryByText(/Rp\s*10[.]?000/i)).not.toBeInTheDocument();
    expect(screen.queryByText(/tanpa biaya tersembunyi/i)).not.toBeInTheDocument();
    expect(screen.queryByText(/laporan.*berkala/i)).not.toBeInTheDocument();
  });
});
