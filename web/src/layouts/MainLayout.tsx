import FireOverlay from '@/components/effects/FireOverlay';
import Header from '@/components/layout/Header';
import Sidebar from '@/components/layout/Sidebar';
import QuickSearch from '@/components/ui/search/QuickSearch';
import { fireStore } from '@/stores/fireStore';
import { searchStore } from '@/stores/searchStore';
import { Outlet, getRouteApi } from '@tanstack/solid-router';
import { createEffect, createSignal, onCleanup } from 'solid-js';

const rootRoute = getRouteApi('__root__');

export default function MainLayout() {
  const data = rootRoute.useLoaderData();
  const [sidebarOpen, setSidebarOpen] = createSignal(false);

  const handleKeyDown = (e: KeyboardEvent) => {
    if ((e.ctrlKey || e.metaKey) && e.key === 'k') {
      e.preventDefault();
      searchStore.open();
    }
  };

  createEffect(() => {
    document.addEventListener('keydown', handleKeyDown);
    onCleanup(() => document.removeEventListener('keydown', handleKeyDown));
  });

  return (
    <div class="flex min-h-svh bg-background">
      <FireOverlay isActive={fireStore.burning} />

      <QuickSearch kits={data().kits} />

      <Sidebar
        data={data()}
        isOpen={sidebarOpen()}
        onClose={() => setSidebarOpen(false)}
      />

      <div class="flex-1 lg:ml-64">
        <Header
          onMenuClick={() => setSidebarOpen(true)}
          onSearchClick={searchStore.open}
        />

        <main>
          <Outlet />
        </main>
      </div>
    </div>
  );
}
