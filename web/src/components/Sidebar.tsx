import { Link } from "@tanstack/solid-router";
import { Show, createMemo, createSignal } from "solid-js";
import type { TemplatesResponse } from "../types/template";
import CategoryList from "./CategoryList";
import SearchBar from "./SearchBar";
import TemplateList from "./TemplateList";
import ViewModeToggle from "./ViewModeToggle";

interface SidebarProps {
  templates: TemplatesResponse;
}

type ViewMode = "list" | "categories";

export default function Sidebar(props: SidebarProps) {
  const [viewMode, setViewMode] = createSignal<ViewMode>("categories");
  const [searchQuery, setSearchQuery] = createSignal("");

  const filteredTemplates = createMemo(() => {
    const query = searchQuery().toLowerCase();
    if (!query) return props.templates.templates || [];

    return (props.templates.templates || []).filter((template) => {
      if (viewMode() === "categories") {
        return template.tags.some((tag) => tag.toLowerCase().includes(query));
      }

      return template.name.toLowerCase().includes(query);
    });
  });

  const categorizedTemplates = createMemo(() => {
    const categories = new Map<string, number>();

    filteredTemplates().forEach((template) => {
      const tags = template.tags.length > 0 ? template.tags : ["general"];
      tags.forEach((tag) => {
        const normalizedTag = tag.toLowerCase();
        categories.set(normalizedTag, (categories.get(normalizedTag) || 0) + 1);
      });
    });

    return Array.from(categories.entries()).sort((a, b) =>
      a[0].localeCompare(b[0])
    );
  });

  return (
    <aside class="fixed left-0 top-0 h-screen w-64 border-r border-sidebar-border bg-sidebar">
      <div class="flex h-full flex-col">
        <div class="flex h-16 items-center justify-between border-b border-sidebar-border px-6">
          <Link
            to="/"
            class="text-xl font-bold text-sidebar-foreground hover:opacity-80"
          >
            Kite
          </Link>

          <ViewModeToggle mode={viewMode()} onModeChange={setViewMode} />
        </div>

        <nav class="flex-1 overflow-y-auto p-4">
          <div class="mb-4">
            <SearchBar
              value={searchQuery()}
              onInput={setSearchQuery}
              placeholder="Search templates..."
            />
          </div>

          <Show when={viewMode() === "list"}>
            <TemplateList templates={filteredTemplates()} />
          </Show>

          <Show when={viewMode() === "categories"}>
            <CategoryList categories={categorizedTemplates()} />
          </Show>
        </nav>
      </div>
    </aside>
  );
}
