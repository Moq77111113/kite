import { Link } from "@tanstack/solid-router";
import { Show, createMemo, createSignal } from "solid-js";
import type { TemplatesResponse } from "@/types/template";
import CategoryList from "@/components/features/CategoryList";
import SearchBar from "@/components/ui/SearchBar";
import TemplateList from "@/components/features/TemplateList";
import ViewModeToggle from "@/components/features/ViewModeToggle";

interface SidebarProps {
  templates: TemplatesResponse;
  isOpen: boolean;
  onClose: () => void;
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
    <>
      <Show when={props.isOpen}>
        <div
          class="fixed inset-0 bg-black/50 z-40 lg:hidden"
          onClick={props.onClose}
        />
      </Show>

      <aside
        class="fixed left-0 top-0 h-svh w-64 border-r border-sidebar-border bg-sidebar z-50 transition-transform duration-200 lg:translate-x-0"
        classList={{
          "-translate-x-full": !props.isOpen,
          "translate-x-0": props.isOpen,
        }}
      >
        <div class="flex h-full flex-col">
          <header class="flex h-16 items-center justify-between border-b border-sidebar-border px-6">
            <Link
              to="/"
              class="text-xl font-bold text-sidebar-foreground hover:opacity-80"
            >
              Kite
            </Link>

            <ViewModeToggle mode={viewMode()} onModeChange={setViewMode} />
          </header>

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
    </>
  );
}
