import { getTagColor } from "@/lib/utils";
import type { TemplateSummary } from "@/types/template";

interface QuickSearchItemProps {
  template: TemplateSummary;
  isSelected: boolean;
  onSelect: (template: TemplateSummary) => void;
}

export default function QuickSearchItem(props: QuickSearchItemProps) {
  return (
    <button
      class="w-full px-3 py-2.5 flex items-center gap-3 hover:bg-accent rounded-lg transition-colors text-left group"
      classList={{
        "bg-accent": props.isSelected,
      }}
      onClick={() => props.onSelect(props.template)}
    >
      <div
        class={`size-10 rounded-lg bg-gradient-to-br ${getTagColor(
          props.template.tags[0] || "general"
        )} flex items-center justify-center text-white shadow-sm flex-shrink-0 group-hover:scale-105 transition-transform`}
      >
        <span class="text-lg font-bold">
          {props.template.name.charAt(0).toUpperCase()}
        </span>
      </div>
      <div class="flex-1 min-w-0">
        <div class="font-medium text-card-foreground mb-0.5">
          {props.template.name}
        </div>
        <div class="text-sm text-muted-foreground truncate">
          {props.template.description}
        </div>
      </div>
      <div class="flex items-center gap-2 flex-shrink-0">
        <span class="text-xs text-muted-foreground hidden sm:block">
          v{props.template.version}
        </span>
      </div>
    </button>
  );
}
