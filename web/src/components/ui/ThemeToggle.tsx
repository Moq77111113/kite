import { createEffect, createSignal, onMount } from "solid-js";
import { MoonIcon, SunIcon } from "./icons";

export default function ThemeToggle() {
  const [theme, setTheme] = createSignal<"light" | "dark">("light");

  onMount(() => {
    const stored = localStorage.getItem("theme");
    const prefersDark = window.matchMedia(
      "(prefers-color-scheme: dark)"
    ).matches;
    const initialTheme =
      (stored as "light" | "dark") || (prefersDark ? "dark" : "light");
    setTheme(initialTheme);
    document.documentElement.classList.toggle("dark", initialTheme === "dark");
  });

  createEffect(() => {
    const currentTheme = theme();
    localStorage.setItem("theme", currentTheme);
    document.documentElement.classList.toggle("dark", currentTheme === "dark");
  });

  const toggleTheme = () => {
    setTheme((prev) => (prev === "light" ? "dark" : "light"));
  };

  return (
    <button
      onClick={toggleTheme}
      class="p-2 cursor-pointer rounded-md hover:bg-sidebar-accent text-sidebar-foreground transition-colors"
      aria-label="Toggle theme"
    >
      {theme() === "light" ? <MoonIcon /> : <SunIcon />}
    </button>
  );
}
