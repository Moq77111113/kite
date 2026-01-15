import { GithubIcon, MenuIcon } from "@/components/ui/icons";
import SyncButton from "@/components/ui/SyncButton";
import ThemeToggle from "@/components/ui/ThemeToggle";
import { SITE_CONFIG } from "@/lib/site";

interface HeaderProps {
  onMenuClick: () => void;
  onSearchClick: () => void;
  lastSync?: string;
  onSyncComplete?: (lastSync?: string) => void;
}

export default function Header(props: HeaderProps) {
  return (
    <>
      <header class="sticky top-0 z-30 flex h-16 items-center justify-between border-b border-border bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 px-4 lg:hidden">
        <div class="flex items-center gap-4">
          <button
            onClick={props.onMenuClick}
            class="p-2 hover:bg-accent rounded-md"
            aria-label="Open menu"
          >
            <MenuIcon />
          </button>
          <span class="text-lg font-bold">{SITE_CONFIG.name}</span>
        </div>

        <div class="flex items-center gap-1">
          <button
            onClick={props.onSearchClick}
            class="p-2 hover:bg-accent rounded-md text-muted-foreground"
            aria-label="Search"
          >
            <kbd class="pointer-events-none inline-flex h-5 select-none items-center gap-1 rounded border bg-muted px-1.5 font-mono text-[10px] font-medium opacity-100">
              ⌘K
            </kbd>
          </button>
          <SyncButton
            lastSync={props.lastSync}
            onSyncComplete={props.onSyncComplete}
          />
          <ThemeToggle />
          <a
            href={SITE_CONFIG.github.url}
            target="_blank"
            rel="noopener noreferrer"
            class="p-2 hover:bg-accent rounded-md"
            aria-label="GitHub"
          >
            <GithubIcon />
          </a>
        </div>
      </header>

      <header class="hidden lg:flex sticky top-0 z-30 h-16 items-center justify-end border-b border-border bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 px-8 gap-2">
        <button
          onClick={props.onSearchClick}
          class="flex items-center gap-2 px-3 py-2 text-sm text-muted-foreground hover:bg-accent rounded-md transition-colors"
        >
          <span>Search</span>
          <kbd class="pointer-events-none inline-flex h-5 select-none items-center gap-1 rounded border bg-muted px-1.5 font-mono text-[10px] font-medium opacity-100">
            <span class="text-xs">⌘</span>K
          </kbd>
        </button>
        <SyncButton
          lastSync={props.lastSync}
          onSyncComplete={props.onSyncComplete}
        />
        <ThemeToggle />
        <a
          href={SITE_CONFIG.github.url}
          target="_blank"
          rel="noopener noreferrer"
          class="p-2 hover:bg-accent rounded-md"
          aria-label="GitHub"
        >
          <GithubIcon />
        </a>
      </header>
    </>
  );
}
