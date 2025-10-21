import { useNavigate } from "@tanstack/solid-router";
import type { Accessor } from "solid-js";
import { createMemo } from "solid-js";

export type SearchParams = {
  search?: string;
  tags?: string;
};

export function useSearchFilters(searchAccessor: Accessor<SearchParams>) {
  const navigate = useNavigate();
  const search = searchAccessor();

  const searchQuery = () => search.search || "";
  const selectedTags = () => search.tags?.split(",").filter(Boolean) || [];

  const setSearch = (value: string) => {
    navigate({
      to: "/",
      search: (prev) => ({
        ...prev,
        search: value || undefined,
      }),
    });
  };

  const setTags = (tags: string[]) => {
    navigate({
      to: "/",
      search: (prev) => ({
        ...prev,
        tags: tags.length > 0 ? tags.join(",") : undefined,
      }),
    });
  };

  const toggleTag = (tag: string) => {
    const tags = selectedTags();
    if (tags.includes(tag)) {
      setTags(tags.filter((t) => t !== tag));
    } else {
      setTags([...tags, tag]);
    }
  };

  const clearFilters = () => {
    navigate({ to: "/", search: {} });
  };

  const hasFilters = createMemo(() => {
    return searchQuery() !== "" || selectedTags().length > 0;
  });

  return {
    searchQuery,
    selectedTags,
    setSearch,
    setTags,
    toggleTag,
    clearFilters,
    hasFilters,
  };
}

export function validateSearchParams(
  search: Record<string, unknown>
): SearchParams {
  return {
    search: typeof search.search === "string" ? search.search : undefined,
    tags: typeof search.tags === "string" ? search.tags : undefined,
  };
}
