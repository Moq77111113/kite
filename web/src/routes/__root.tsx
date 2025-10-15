import { createRootRoute } from "@tanstack/solid-router";
import { fetchTemplates } from "../api/templates";
import MainLayout from "../layouts/MainLayout";

export const Route = createRootRoute({
  loader: () => fetchTemplates(),
  component: MainLayout,
  errorComponent: (e) => <div>Failed to load templates {e.error.message}</div>,
});
