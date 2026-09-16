import { render, screen } from "@testing-library/react";
import Home from "./page";

describe("home page", () => {
  it("offers a visible path from the editorial introduction to the clarity section", () => {
    const { container } = render(<Home />);

    expect(screen.getByRole("banner")).toBeInTheDocument();
    expect(screen.getByRole("navigation", { name: "Navigasi utama" })).toBeInTheDocument();
    expect(screen.getByRole("main")).toBeInTheDocument();
    expect(screen.getByRole("heading", { level: 1 })).toHaveTextContent(
      "Harapan tumbuh dari hal yang jelas.",
    );

    const link = screen.getByRole("link", { name: /Lihat cara kami menjelaskan/i });
    const target = container.querySelector(link.getAttribute("href")!);

    expect(link).toBeVisible();
    expect(target).toHaveAttribute("id", "kejelasan");
    expect(target).toContainElement(
      screen.getByRole("heading", { name: "Ruang untuk melihat lebih dekat." }),
    );
  });
});
