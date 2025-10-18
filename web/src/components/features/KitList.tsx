import type { KitSummary } from '@/types/kit';
import { Link } from '@tanstack/solid-router';
import { For, Show } from 'solid-js';

interface KitListProps {
  kits: KitSummary[];
}

export default function KitList(props: KitListProps) {
  return (
    <div class="mb-4">
      <h2 class="px-3 text-xs font-semibold uppercase tracking-wider text-sidebar-foreground opacity-60 mb-2">
        Kits
      </h2>
      <Show when={props.kits.length}>
        <div class="space-y-1">
          <For each={props.kits}>
            {(kit) => (
              <Link
                to="/kits/$name"
                params={{ name: kit.name }}
                class="block rounded-lg px-3 py-2 text-sm font-medium text-sidebar-foreground transition-colors hover:bg-sidebar-accent hover:text-sidebar-accent-foreground"
                activeProps={{
                  class: 'bg-sidebar-accent text-sidebar-accent-foreground',
                }}
              >
                {kit.name}
              </Link>
            )}
          </For>
        </div>
      </Show>
    </div>
  );
}
