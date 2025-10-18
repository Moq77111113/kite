import KitCard from '@/components/features/KitCard';
import SearchBar from '@/components/ui/SearchBar';
import { createFileRoute, getRouteApi } from '@tanstack/solid-router';
import { For, createMemo, createSignal } from 'solid-js';

const rootRoute = getRouteApi('__root__');

export const Route = createFileRoute('/')({
  component: Home,
});

function Home() {
  const data = rootRoute.useLoaderData();
  const [searchQuery, setSearchQuery] = createSignal('');

  const filteredKits = createMemo(() => {
    const query = searchQuery().toLowerCase();
    if (!query) return data().kits;

    return data().kits.filter(
      (kit) =>
        kit.name.toLowerCase().includes(query) ||
        kit.description?.toLowerCase().includes(query) ||
        kit.tags.some((tag) => tag.toLowerCase().includes(query))
    );
  });

  return (
    <div class="max-w-6xl px-4 sm:px-8 py-8 sm:py-12">
      <div class="mb-10">
        <h1 class="text-4xl font-bold text-foreground mb-3">Browse Kits</h1>
        <p class="text-base text-muted-foreground mb-6">
          Fork your infrastructure, don't worship it. Browse and copy kits into
          your project as editable code.
        </p>
        <div class="max-w-md">
          <SearchBar
            value={searchQuery()}
            onInput={setSearchQuery}
            placeholder="Search kits..."
          />
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
        <For each={filteredKits()}>{(kit) => <KitCard kit={kit} />}</For>
      </div>
    </div>
  );
}
