import { Link } from "@tanstack/solid-router";
import { For, Show } from "solid-js";
import type { TemplateSummary } from "../types/template";

interface TemplateListProps {
  templates: TemplateSummary[];
}

export default function TemplateList(props: TemplateListProps) {
  return (
    <div class="mb-4">
      <h2 class="px-3 text-xs font-semibold uppercase tracking-wider text-sidebar-foreground opacity-60 mb-2">
        Templates
      </h2>
      <Show when={props.templates}>
        <div class="space-y-1">
          <For each={props.templates}>
            {(template) => (
              <Link
                to="/templates/$name"
                params={{ name: template.name }}
                class="block rounded-lg px-3 py-2 text-sm font-medium text-sidebar-foreground transition-colors hover:bg-sidebar-accent hover:text-sidebar-accent-foreground"
                activeProps={{
                  class: "bg-sidebar-accent text-sidebar-accent-foreground",
                }}
              >
                {template.name}
              </Link>
            )}
          </For>
        </div>
      </Show>
    </div>
  );
}
