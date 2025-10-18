import { getAvatarNumber } from '@/lib/utils';
import type { KitSummary } from '@/types/kit';
import { Link } from '@tanstack/solid-router';

interface KitCardProps {
  kit: KitSummary;
}

export default function KitCard(props: KitCardProps) {
  const primaryTag = () => props.kit.tags[0] || 'general';
  const avatarNum = () => getAvatarNumber(primaryTag());

  return (
    <Link
      to="/kits/$name"
      params={{ name: props.kit.name }}
      class="block rounded-xl border border-border bg-card p-5 transition-all hover:border-accent hover:shadow-lg group"
    >
      <div class="flex items-start justify-between mb-3">
        <span class="px-2.5 py-1 text-xs rounded-md bg-muted/50 text-muted-foreground capitalize font-medium">
          {primaryTag()}
        </span>
        <div
          class={`size-14 rounded-2xl flex items-center justify-center text-white shadow-md group-hover:scale-105 transition-transform bg-[var(--avatar)]`}
          style={{
            '--avatar': `var(--avatar-${avatarNum()})`,
          }}
        >
          <span class="text-2xl font-bold">
            {props.kit.name.charAt(0).toUpperCase()}
          </span>
        </div>
      </div>

      <h3 class="text-lg font-bold text-card-foreground mb-2 group-hover:text-primary transition-colors">
        {props.kit.name}
      </h3>

      <p class="text-sm text-muted-foreground mb-4 line-clamp-2 leading-relaxed">
        {props.kit.description}
      </p>

      <div class="flex items-center gap-3 text-xs text-muted-foreground pt-2 border-t border-border/50">
        <span class="font-medium">v{props.kit.version}</span>
        <span class="opacity-50">•</span>
        <span>by {props.kit.author}</span>
      </div>
    </Link>
  );
}
