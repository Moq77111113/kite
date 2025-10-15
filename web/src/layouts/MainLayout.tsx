import { Outlet, getRouteApi } from "@tanstack/solid-router";
import { createEffect, createSignal, onCleanup } from "solid-js";
import Sidebar from "@/components/layout/Sidebar";
import QuickSearch from "@/components/ui/QuickSearch";
import Header from "@/components/layout/Header";
import FireOverlay from "@/components/effects/FireOverlay";
import { fireStore } from "@/stores/fireStore";

const rootRoute = getRouteApi("__root__");

export default function MainLayout() {
  const data = rootRoute.useLoaderData();
  const [sidebarOpen, setSidebarOpen] = createSignal(false);
  const [searchOpen, setSearchOpen] = createSignal(false);

  const handleKeyDown = (e: KeyboardEvent) => {
    if ((e.ctrlKey || e.metaKey) && e.key === "k") {
      e.preventDefault();
      setSearchOpen(true);
    }
  };

  createEffect(() => {
    document.addEventListener("keydown", handleKeyDown);
    onCleanup(() => document.removeEventListener("keydown", handleKeyDown));
  });

  return (
    <div class="flex min-h-svh bg-background">
      <FireOverlay isActive={fireStore.isOnFire} />

      <QuickSearch
        templates={data().templates}
        isOpen={searchOpen()}
        onClose={() => setSearchOpen(false)}
      />

      <Sidebar
        templates={data()}
        isOpen={sidebarOpen()}
        onClose={() => setSidebarOpen(false)}
      />

      <div class="flex-1 lg:ml-64">
        <Header
          onMenuClick={() => setSidebarOpen(true)}
          onSearchClick={() => setSearchOpen(true)}
        />

        <main>
          <Outlet />
        </main>
      </div>
    </div>
  );
}
