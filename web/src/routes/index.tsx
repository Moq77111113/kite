import { createFileRoute, getRouteApi } from "@tanstack/solid-router";
import { For, createSignal, createMemo } from "solid-js";
import TemplateCard from "../components/TemplateCard";
import SearchBar from "../components/SearchBar";

const rootRoute = getRouteApi('__root__')

export const Route = createFileRoute("/")({
  component: Home,
});

function Home() {
  const data = rootRoute.useLoaderData()
  const [searchQuery, setSearchQuery] = createSignal("");

  const filteredTemplates = createMemo(() => {
    const query = searchQuery().toLowerCase();
    if (!query) return data().templates;

    return data().templates.filter((template) =>
      template.name.toLowerCase().includes(query) ||
      template.description?.toLowerCase().includes(query) ||
      template.tags.some((tag) => tag.toLowerCase().includes(query))
    );
  });

  return (
    <div class="min-h-screen">
      <div class="max-w-7xl mx-auto px-8 py-12">
        <div class="mb-10">
          <h1 class="text-4xl font-bold text-foreground mb-3">Templates</h1>
          <p class="text-base text-muted-foreground mb-6">
            Fork your infrastructure, don't worship it. Browse and copy templates into your project as editable code.
          </p>
          <div class="max-w-md">
            <SearchBar
              value={searchQuery()}
              onInput={setSearchQuery}
              placeholder="Search templates..."
            />
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
          <For each={filteredTemplates()}>
            {(template) => <TemplateCard template={template} />}
          </For>
        </div>
      </div>
    </div>
  );
}
