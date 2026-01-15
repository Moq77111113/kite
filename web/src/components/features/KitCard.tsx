import TagsCarousel from "@/components/ui/TagsCarousel";

import { getAvatarNumber, timeSince } from "@/lib/utils";
import type { KitSummary } from "@/types/kit";
import { Link } from "@tanstack/solid-router";
import { Show } from "solid-js";


interface KitCardProps {
  kit: KitSummary;
}

export default function KitCard(props: KitCardProps) {
  const primaryTag = () => props.kit.tags[0];
  const avatarNum = () => getAvatarNumber(primaryTag() ?? props.kit.name);

  return (
    <Link
      to="/kits/$name"
      params={{ name: props.kit.id }}
      class="flex flex-col justify-between rounded-xl border border-border bg-card p-5 transition-all hover:border-accent hover:shadow-lg group"
    >
      <div class="flex items-start justify-between gap-3 mb-3">
        <TagsCarousel tags={props.kit.tags} maxVisible={3} />

        <div
          class={`size-14 rounded-2xl flex items-center justify-center text-white shadow-md group-hover:scale-105 transition-transform bg-[var(--avatar)] flex-shrink-0`}
          style={{
            "--avatar": `var(--avatar-${avatarNum()})`,
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
        <Show when={props.kit.author}>
          <span class="opacity-50">•</span>
          <span>by {props.kit.author}</span>
        </Show>
        <Show when={props.kit.lastUpdated}>
          <span class="opacity-50">•</span>
          <span class="opacity-75">
            {timeSince(new Date(props.kit.lastUpdated ?? ""))}
          </span>
        </Show>
      </div>
    </Link>
  );
}
