import KitCard from '@/components/features/KitCard';
import SearchBar from '@/components/ui/SearchBar';
import TagFilter from '@/components/ui/TagFilter';
import { useSearchFilters, validateSearchParams } from '@/hooks/useSearchFilters';
import { createFileRoute, getRouteApi } from '@tanstack/solid-router';
import { For, createMemo } from 'solid-js';

const rootRoute = getRouteApi('__root__');

export const Route = createFileRoute('/')({
  component: Home,
  validateSearch: validateSearchParams,
});

function Home() {
  const data = rootRoute.useLoaderData();
  const search = Route.useSearch();
  const filters = useSearchFilters(search);

  const allTags = createMemo(() => {
    const tagCounts = new Map<string, number>();
    data().kits.forEach((kit) =>
      kit.tags.forEach((tag) => tagCounts.set(tag, (tagCounts.get(tag) || 0) + 1))
    );
    return Array.from(tagCounts.entries()).map(([name, count]) => ({ name, count }));
  });

  const filteredKits = createMemo(() => {
    const query = filters.searchQuery().toLowerCase();
    const tags = filters.selectedTags();

    return data().kits.filter((kit) => {
      if (tags.length > 0 && !tags.some((tag) => kit.tags.includes(tag))) {
        return false;
      }

      if (!query) return true;

      return (
        kit.name.toLowerCase().includes(query) ||
        kit.description?.toLowerCase().includes(query) ||
        kit.tags.some((tag) => tag.toLowerCase().includes(query))
      );
    });
  });

  return (
    <div class="max-w-6xl px-4 sm:px-8 py-8 sm:py-12">
      <div class="mb-10">
        <h1 class="text-4xl font-bold text-foreground mb-3">Browse Kits</h1>
        <p class="text-base text-muted-foreground mb-6">
          Fork your infrastructure, don't worship it. Browse and copy kits into
          your project as editable code.
        </p>

        <div class="max-w-md mb-6">
          <SearchBar
            value={filters.searchQuery()}
            onInput={filters.setSearch}
            placeholder="Search kits..."
          />
        </div>

        <TagFilter
          tags={allTags()}
          selectedTags={filters.selectedTags()}
          onToggle={filters.toggleTag}
          onClear={filters.clearFilters}
          showClear={filters.hasFilters()}
        />
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
        <For each={filteredKits()}>{(kit) => <KitCard kit={kit} />}</For>
      </div>
    </div>
  );
}
