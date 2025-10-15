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
    <div class="min-h-screen">
      <div class="flex gap-8 max-w-7xl mx-auto px-8 py-12">
        <div class="flex-1 max-w-4xl">
          <div class="mb-10">
            <div class="flex items-center gap-3 mb-3">
              <h1 class="text-4xl font-bold text-foreground">
                {template().name}
              </h1>
              <span class="px-3 py-1.5 text-xs font-semibold rounded-lg bg-muted/50 text-muted-foreground">
                v{template().version}
              </span>
            </div>
            <p class="text-base text-muted-foreground mb-4 leading-relaxed">{template().description}</p>

            <div class="flex flex-wrap gap-2 mb-4">
              <For each={template().tags}>
                {(tag) => (
                  <span class="px-2.5 py-1 text-xs font-medium rounded-md bg-accent/50 text-accent-foreground capitalize">
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
              class="rounded-xl border border-border bg-card p-5 flex items-center justify-between shadow-sm"
            >
              <code class="text-sm font-mono text-card-foreground font-medium">
                kite add {template().name}
              </code>
              <CopyButton text={`kite add ${template().name}`} />
            </div>
          </div>

          <Show when={template().readme}>
            <div id="readme" class="mb-10">
              <h2 class="text-2xl font-bold text-foreground mb-5">README</h2>
              <div class="rounded-xl border border-border bg-card p-6 shadow-sm">
                <pre class="text-sm text-card-foreground whitespace-pre-wrap leading-relaxed">
                  {template().readme}
                </pre>
              </div>
            </div>
          </Show>

          <div id="files">
            <h2 class="text-2xl font-bold text-foreground mb-5">Files</h2>
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
