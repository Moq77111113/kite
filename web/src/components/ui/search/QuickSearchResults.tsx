import type { TemplateSummary } from "@/types/template";
import { For, Show } from "solid-js";
import QuickSearchEmptyState from "./QuickSearchEmptyState";
import QuickSearchItem from "./QuickSearchItem";

interface QuickSearchResultsProps {
  templates: TemplateSummary[];
  selectedIndex: number;
  onSelectTemplate: (template: TemplateSummary) => void;
}

export default function QuickSearchResults(props: QuickSearchResultsProps) {
  return (
    <div class="max-h-[420px] overflow-y-auto">
      <Show
        when={props.templates.length > 0}
        fallback={<QuickSearchEmptyState />}
      >
        <div class="p-2">
          <For each={props.templates}>
            {(template, index) => (
              <QuickSearchItem
                template={template}
                isSelected={index() === props.selectedIndex}
                onSelect={props.onSelectTemplate}
              />
            )}
          </For>
        </div>
      </Show>
    </div>
  );
}
