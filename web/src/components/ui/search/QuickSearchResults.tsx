import type { KitSummary } from '@/types/kit';
import { For, Show } from 'solid-js';
import QuickSearchEmptyState from './QuickSearchEmptyState';
import QuickSearchItem from './QuickSearchItem';

interface QuickSearchResultsProps {
  kits: KitSummary[];
  selectedIndex: number;
  onSelectKit: (kit: KitSummary) => void;
}

export default function QuickSearchResults(props: QuickSearchResultsProps) {
  return (
    <div class="max-h-[420px] overflow-y-auto">
      <Show when={props.kits.length > 0} fallback={<QuickSearchEmptyState />}>
        <div class="p-2">
          <For each={props.kits}>
            {(kit, index) => (
              <QuickSearchItem
                kit={kit}
                isSelected={index() === props.selectedIndex}
                onSelect={props.onSelectKit}
              />
            )}
          </For>
        </div>
      </Show>
    </div>
  );
}
