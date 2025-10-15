import { Outlet, getRouteApi } from "@tanstack/solid-router";
import { createSignal } from "solid-js";
import Sidebar from "@/components/layout/Sidebar";

const rootRoute = getRouteApi("__root__");

export default function MainLayout() {
  const data = rootRoute.useLoaderData();
  const [sidebarOpen, setSidebarOpen] = createSignal(false);

  return (
    <div class="flex min-h-svh bg-background">
      <Sidebar
        templates={data()}
        isOpen={sidebarOpen()}
        onClose={() => setSidebarOpen(false)}
      />

      <div class="flex-1 lg:ml-64">
        <header class="sticky top-0 z-30 flex h-16 items-center border-b border-border bg-background px-4 lg:hidden">
          <button
            onClick={() => setSidebarOpen(true)}
            class="p-2 hover:bg-accent rounded-md"
            aria-label="Open menu"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <line x1="3" y1="12" x2="21" y2="12" />
              <line x1="3" y1="6" x2="21" y2="6" />
              <line x1="3" y1="18" x2="21" y2="18" />
            </svg>
          </button>
          <span class="ml-4 text-lg font-bold">Kite</span>
        </header>

        <main>
          <Outlet />
        </main>
      </div>
    </div>
  );
}
