export default function QuickSearchFooter() {
  return (
    <div class="px-4 py-3 border-t border-border/50 bg-muted/30 flex items-center justify-between text-xs text-muted-foreground">
      <div class="flex items-center gap-3">
        <div class="flex items-center gap-1.5">
          <kbd class="px-1.5 py-0.5 bg-background border border-border rounded text-[10px] font-medium">
            ↑↓
          </kbd>
          <span class="hidden sm:inline">navigate</span>
        </div>
        <div class="flex items-center gap-1.5">
          <kbd class="px-1.5 py-0.5 bg-background border border-border rounded text-[10px] font-medium">
            ↵
          </kbd>
          <span class="hidden sm:inline">select</span>
        </div>
        <div class="flex items-center gap-1.5">
          <kbd class="px-1.5 py-0.5 bg-background border border-border rounded text-[10px] font-medium">
            esc
          </kbd>
          <span class="hidden sm:inline">close</span>
        </div>
      </div>
    </div>
  );
}
