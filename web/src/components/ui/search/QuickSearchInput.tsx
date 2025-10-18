import { SearchIcon } from "@/components/ui/icons";
import { onMount } from "solid-js";

interface QuickSearchInputProps {
  query: string;
  onQueryChange: (query: string) => void;
  inputRef: (el: HTMLInputElement) => void;
}

export default function QuickSearchInput(props: QuickSearchInputProps) {
  let localInputRef: HTMLInputElement | undefined;

  const handleRef = (el: HTMLInputElement) => {
    localInputRef = el;
    props.inputRef(el);
  };

  onMount(() => {
    // Focus the input when component mounts
    localInputRef?.focus();
  });

  return (
    <div class="flex items-center gap-3 px-4 py-3 border-b border-border/50">
      <SearchIcon class="text-muted-foreground" />
      <input
        ref={handleRef}
        type="text"
        placeholder="Search kits..."
        value={props.query}
        onInput={(e) => props.onQueryChange(e.currentTarget.value)}
        class="flex-1 bg-transparent text-card-foreground placeholder:text-muted-foreground outline-none"
        autofocus
      />
    </div>
  );
}
