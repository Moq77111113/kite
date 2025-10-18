import { useQuickSearch } from "@/hooks/useQuickSearch";
import { searchStore } from "@/stores/searchStore";
import type { KitSummary } from "@/types/kit";
import { Show } from "solid-js";
import QuickSearchFooter from "./search/QuickSearchFooter";
import QuickSearchInput from "./search/QuickSearchInput";
import QuickSearchOverlay from "./search/QuickSearchOverlay";
import QuickSearchResults from "./search/QuickSearchResults";

interface QuickSearchProps {
  kits: KitSummary[];
}

export default function QuickSearch(props: QuickSearchProps) {
  const {
    query,
    setQuery,
    selectedIndex,
    filteredKits,
    selectKit,
    inputRef,
  } = useQuickSearch(props.kits);

  return (
    <Show when={searchStore.isOpen}>
      <QuickSearchOverlay onClose={searchStore.close}>
        <QuickSearchInput
          query={query()}
          onQueryChange={setQuery}
          inputRef={inputRef}
        />
        <QuickSearchResults
          kits={filteredKits()}
          selectedIndex={selectedIndex()}
          onSelectKit={selectKit}
        />
        <QuickSearchFooter />
      </QuickSearchOverlay>
    </Show>
  );
}
