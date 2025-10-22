import { For, Show, createMemo, createSignal } from "solid-js";

type TagWithCount = {
  name: string;
  count: number;
};

type TagFilterProps = {
  tags: TagWithCount[];
  selectedTags: string[];
  onToggle: (tag: string) => void;
  onClear?: () => void;
  showClear?: boolean;
  maxVisible?: number;
};

export default function TagFilter(props: TagFilterProps) {
  const [expanded, setExpanded] = createSignal(false);
  const maxVisible = () => props.maxVisible ?? 15;

  const visibleTags = createMemo(() => {
    const selected = new Set(props.selectedTags);
    const sortedTags = [...props.tags].sort((a, b) => b.count - a.count);

    if (expanded() || sortedTags.length <= maxVisible()) {
      return sortedTags;
    }

    const topTags = sortedTags.slice(0, maxVisible());
    const topTagNames = new Set(topTags.map((t) => t.name));

    const additionalSelected = sortedTags.filter(
      (t) => selected.has(t.name) && !topTagNames.has(t.name)
    );

    return [...topTags, ...additionalSelected];
  });

  const hasMore = () => props.tags.length > maxVisible();

  return (
    <div class="space-y-3">
      <div class="flex items-center gap-2">
        <h3 class="text-sm font-medium text-muted-foreground">
          Filter by tags
        </h3>
        {props.showClear && (
          <button
            onClick={props.onClear}
            class="text-xs text-muted-foreground hover:text-foreground transition-colors"
          >
            Clear all
          </button>
        )}
      </div>

      <div class="flex flex-wrap gap-2">
        <For each={visibleTags()}>
          {(tag) => {
            const isSelected = () => props.selectedTags.includes(tag.name);
            return (
              <button
                onClick={() => props.onToggle(tag.name)}
                class={`px-3 py-1 rounded-full text-sm font-medium transition-colors ${
                  isSelected()
                    ? "bg-primary text-primary-foreground"
                    : "bg-secondary text-secondary-foreground hover:bg-secondary/80"
                }`}
              >
                {tag.name}
              </button>
            );
          }}
        </For>

        <Show when={hasMore()}>
          <button
            onClick={() => setExpanded(!expanded())}
            class="px-3 py-1 rounded-full text-sm font-medium bg-muted text-muted-foreground hover:bg-muted/80 transition-colors"
          >
            {expanded()
              ? "Show less"
              : `+${props.tags.length - maxVisible()} more`}
          </button>
        </Show>
      </div>
    </div>
  );
}
