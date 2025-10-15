import { Link } from "@tanstack/solid-router";
import { For } from "solid-js";

interface CategoryListProps {
  categories: Array<[string, number]>;
}

export default function CategoryList(props: CategoryListProps) {
  return (
    <div class="space-y-1">
      <For each={props.categories}>
        {([category, count]) => (
          <Link
            to="/categories/$name"
            params={{ name: category }}
            class="block rounded-lg px-3 py-2 text-sm font-medium text-sidebar-foreground transition-colors hover:bg-sidebar-accent hover:text-sidebar-accent-foreground group"
            activeProps={{
              class: "bg-sidebar-accent text-sidebar-accent-foreground",
            }}
          >
            <div class="flex items-center justify-between">
              <span class="capitalize">{category}</span>
              <span class="text-[10px] px-1.5 py-0.5 rounded bg-sidebar-accent text-sidebar-foreground group-hover:bg-sidebar-accent/80">
                {count}
              </span>
            </div>
          </Link>
        )}
      </For>
    </div>
  );
}
