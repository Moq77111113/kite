import CategoryList from '@/components/features/CategoryList';
import KitList from '@/components/features/KitList';
import ViewModeToggle from '@/components/features/ViewModeToggle';
import { FireIcon } from '@/components/ui/icons';
import SearchBar from '@/components/ui/SearchBar';
import { SITE_CONFIG } from '@/lib/site';
import { fireStore } from '@/stores/fireStore';
import type { KitsResponse } from '@/types/kit';
import { Link } from '@tanstack/solid-router';
import { Show, createMemo, createSignal } from 'solid-js';

interface SidebarProps {
  data: KitsResponse;
  isOpen: boolean;
  onClose: () => void;
}

type ViewMode = 'list' | 'categories';

export default function Sidebar(props: SidebarProps) {
  const [viewMode, setViewMode] = createSignal<ViewMode>('categories');
  const [searchQuery, setSearchQuery] = createSignal('');

  const filteredKits = createMemo(() => {
    const query = searchQuery().toLowerCase();
    if (!query) return props.data.kits || [];

    return (props.data.kits || []).filter((kit) => {
      if (viewMode() === 'categories') {
        return kit.tags.some((tag) => tag.toLowerCase().includes(query));
      }

      return kit.name.toLowerCase().includes(query);
    });
  });

  const categorizedKits = createMemo(() => {
    const categories = new Map<string, number>();

    filteredKits().forEach((kit) => {
      const tags = kit.tags.length > 0 ? kit.tags : ['general'];
      tags.forEach((tag) => {
        const normalizedTag = tag.toLowerCase();
        categories.set(normalizedTag, (categories.get(normalizedTag) || 0) + 1);
      });
    });

    return Array.from(categories.entries()).sort((a, b) =>
      a[0].localeCompare(b[0])
    );
  });

  return (
    <>
      <Show when={props.isOpen}>
        <div
          class="fixed inset-0 bg-black/50 z-40 lg:hidden"
          onClick={props.onClose}
        />
      </Show>

      <aside
        class="fixed left-0 top-0 h-svh w-64 border-r border-sidebar-border bg-sidebar z-50 transition-transform duration-200 lg:translate-x-0"
        classList={{
          '-translate-x-full': !props.isOpen,
          'translate-x-0': props.isOpen,
        }}
      >
        <div class="flex h-full flex-col">
          <header class="flex h-16 items-center justify-between border-b border-sidebar-border px-6">
            <Link
              to="/"
              class="text-xl font-bold text-sidebar-foreground hover:opacity-80"
            >
              Kite
            </Link>

            <ViewModeToggle mode={viewMode()} onModeChange={setViewMode} />
          </header>

          <nav class="flex-1 overflow-y-auto p-4">
            <div class="mb-4">
              <SearchBar
                value={searchQuery()}
                onInput={setSearchQuery}
                placeholder="Search kits..."
              />
            </div>

            <Show when={viewMode() === 'list'}>
              <KitList kits={filteredKits()} />
            </Show>

            <Show when={viewMode() === 'categories'}>
              <CategoryList categories={categorizedKits()} />
            </Show>
          </nav>

          <footer class="border-t border-sidebar-border px-4 py-3">
            <a
              href={SITE_CONFIG.github.issues}
              target="_blank"
              rel="noopener noreferrer"
              class="flex items-center gap-2.5 text-xs text-sidebar-foreground/70 hover:text-orange-500 transition-colors group"
              onMouseEnter={() => fireStore.boom()}
              onMouseLeave={() => fireStore.calm()}
            >
              <FireIcon class="flex-shrink-0 group-hover:animate-pulse" />
              <span class="leading-tight">
                Report issue (it's broken anyway)
              </span>
            </a>
          </footer>
        </div>
      </aside>
    </>
  );
}
