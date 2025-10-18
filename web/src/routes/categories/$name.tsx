import { fetchKits } from '@/api/kits';
import KitCard from '@/components/features/KitCard';
import SearchBar from '@/components/ui/SearchBar';
import { createFileRoute } from '@tanstack/solid-router';
import { For, createMemo, createSignal } from 'solid-js';

export const Route = createFileRoute('/categories/$name')({
  component: CategoryView,
  loader: async ({ params }) => {
    const kits = await fetchKits(params.name);
    return kits.kits;
  },
});

function CategoryView() {
  const data = Route.useLoaderData();
  const params = Route.useParams();
  const [searchQuery, setSearchQuery] = createSignal('');

  const filteredKits = createMemo(() => {
    const query = searchQuery().toLowerCase();
    if (!query) return data();

    return data().filter((kit) => kit.name.toLowerCase().includes(query));
  });

  return (
    <div class="max-w-6xl px-4 sm:px-8 py-8 sm:py-12">
      <header class="mb-10">
        <h1 class="text-3xl sm:text-4xl font-bold text-foreground mb-3 capitalize">
          {params().name}
        </h1>
        <p class="text-base text-muted-foreground mb-6">
          {filteredKits().length} kit
          {filteredKits().length !== 1 ? 's' : ''} in this category
        </p>
        <div class="max-w-md">
          <SearchBar
            value={searchQuery()}
            onInput={setSearchQuery}
            placeholder="Search in category..."
          />
        </div>
      </header>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
        <For each={filteredKits()}>{(kit) => <KitCard kit={kit} />}</For>
      </div>
    </div>
  );
}
