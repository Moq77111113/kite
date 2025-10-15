import { createFileRoute } from "@tanstack/solid-router";
import { For, Show, createMemo } from "solid-js";
import { fetchTemplate } from "../../api/templates";
import { CodeBlock } from "../../components/CodeBlock";
import { CopyButton } from "../../components/CopyButton";
import { SectionNav } from "../../components/SectionNav";

export const Route = createFileRoute("/templates/$name")({
  loader: ({ params }) => fetchTemplate(params.name),
  component: TemplateDetail,
});

function TemplateDetail() {
  const template = Route.useLoaderData();

  const sections = createMemo(() => {
    const secs = [{ id: "install", label: "Installation" }];
    if (template().readme) {
      secs.push({ id: "readme", label: "README" });
    }
    secs.push({ id: "files", label: "Files" });
    return secs;
  });

  return (
    <div class="px-8 py-8">
      <div class="flex gap-8 max-w-7xl mx-auto">
        <div class="flex-1 max-w-4xl">
          <div class="mb-8">
            <div class="flex items-center gap-3 mb-2">
              <h1 class="text-3xl font-bold text-foreground">
                {template().name}
              </h1>
              <span class="px-2 py-1 text-xs rounded-md bg-secondary text-secondary-foreground">
                v{template().version}
              </span>
            </div>
            <p class="text-muted-foreground mb-4">{template().description}</p>

            <div class="flex flex-wrap gap-2 mb-4">
              <For each={template().tags}>
                {(tag) => (
                  <span class="px-2 py-1 text-xs rounded-md bg-accent text-accent-foreground">
                    {tag}
                  </span>
                )}
              </For>
            </div>

            <div class="flex items-center gap-4 text-sm text-muted-foreground mb-6">
              <span>by {template().author}</span>
            </div>

            <div
              id="install"
              class="rounded-lg border border-border bg-card p-4 flex items-center justify-between"
            >
              <code class="text-sm text-card-foreground">
                kite add {template().name}
              </code>
              <CopyButton text={`kite add ${template().name}`} />
            </div>
          </div>

          <Show when={template().readme}>
            <div id="readme" class="mb-8">
              <h2 class="text-xl font-semibold text-foreground mb-4">README</h2>
              <div class="rounded-lg border border-border bg-card p-6">
                <pre class="text-sm text-card-foreground whitespace-pre-wrap">
                  {template().readme}
                </pre>
              </div>
            </div>
          </Show>

          <div id="files">
            <h2 class="text-xl font-semibold text-foreground mb-4">Files</h2>
            <div class="space-y-4">
              <For each={template().files}>
                {(file) => (
                  <CodeBlock content={file.content} filename={file.path} />
                )}
              </For>
            </div>
          </div>
        </div>

        <div class="hidden lg:block w-64">
          <SectionNav sections={sections()} />
        </div>
      </div>
    </div>
  );
}
