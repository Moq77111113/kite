interface SearchBarProps {
  value: string;
  onInput: (value: string) => void;
  placeholder?: string;
}

export default function SearchBar(props: SearchBarProps) {
  return (
    <div class="relative">
      <svg
        class="absolute left-3 top-1/2 -translate-y-1/2 size-4 text-muted-foreground"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
        />
      </svg>
      <input
        type="text"
        value={props.value}
        on:input={(e) => props.onInput(e.currentTarget.value)}
        placeholder={props.placeholder || "Search..."}
        class="w-full pl-9 pr-3 py-2 text-sm rounded-lg border border-input bg-background text-foreground placeholder:text-muted-foreground focus:outline-none focus:ring-2 focus:ring-ring"
      />
    </div>
  );
}
