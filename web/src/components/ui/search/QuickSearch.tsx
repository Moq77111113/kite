import { useQuickSearch } from "@/hooks/useQuickSearch";
import { searchStore } from "@/stores/searchStore";
import type { TemplateSummary } from "@/types/template";
import { Show } from "solid-js";
import QuickSearchFooter from "./QuickSearchFooter";
import QuickSearchInput from "./QuickSearchInput";
import QuickSearchOverlay from "./QuickSearchOverlay";
import QuickSearchResults from "./QuickSearchResults";

interface QuickSearchProps {
  templates: TemplateSummary[];
}

export default function QuickSearch(props: QuickSearchProps) {
  const {
    query,
    setQuery,
    selectedIndex,
    filteredTemplates,
    selectTemplate,
    inputRef,
  } = useQuickSearch(props.templates);

  return (
    <Show when={searchStore.isOpen}>
      <QuickSearchOverlay onClose={searchStore.close}>
        <QuickSearchInput
          query={query()}
          onQueryChange={setQuery}
          inputRef={inputRef}
        />
        <QuickSearchResults
          templates={filteredTemplates()}
          selectedIndex={selectedIndex()}
          onSelectTemplate={selectTemplate}
        />
        <QuickSearchFooter />
      </QuickSearchOverlay>
    </Show>
  );
}
