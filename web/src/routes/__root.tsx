import { createRootRoute } from "@tanstack/solid-router";
import { fetchKits } from "@/api/kits";
import MainLayout from "../layouts/MainLayout";

export const Route = createRootRoute({
  loader: () => fetchKits(),
  component: MainLayout,
  errorComponent: (e) => <div>Failed to load kits {e.error.message}</div>,
});
