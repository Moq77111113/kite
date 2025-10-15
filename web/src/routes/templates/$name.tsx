import { fetchTemplate } from "@/api/templates";
import { SectionNav } from "@/components/features/SectionNav";
import { CodeBlock } from "@/components/ui/CodeBlock";
import { CopyButton } from "@/components/ui/CopyButton";
import { createFileRoute } from "@tanstack/solid-router";
import { For, Show, createMemo } from "solid-js";
import { marked } from "marked";

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

  const renderedReadme = createMemo(() => {
    if (!template().readme) return "";
    return marked.parse(template().readme, { async: false });
  });

  return (
    <div class="flex gap-8 max-w-6xl px-4 sm:px-8 py-8 sm:py-12">
      <article class="flex-1 max-w-4xl">
        <header class="mb-10">
          <div class="flex items-center gap-3 mb-3">
            <h1 class="text-3xl sm:text-4xl font-bold text-foreground">
              {template().name}
            </h1>
            <span class="px-3 py-1.5 text-xs font-semibold rounded-lg bg-muted/50 text-muted-foreground">
              v{template().version}
            </span>
          </div>
          <p class="text-base text-muted-foreground mb-4 leading-relaxed">
            {template().description}
          </p>

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
            <code class="text-sm font-mono text-card-foreground font-medium break-all">
              kite add {template().name}
            </code>
            <CopyButton text={`kite add ${template().name}`} />
          </div>
        </header>

        <Show when={template().readme}>
          <section id="readme" class="mb-10">
            <h2 class="text-2xl font-bold text-foreground mb-5">README</h2>
            <div class="rounded-xl border border-border bg-card p-6 shadow-sm">
              <div class="prose prose-slate dark:prose-invert max-w-none" innerHTML={renderedReadme()} />
            </div>
          </section>
        </Show>

        <section id="files">
          <h2 class="text-2xl font-bold text-foreground mb-5">Files</h2>
          <div class="space-y-4">
            <For each={template().files}>
              {(file) => (
                <CodeBlock content={file.content} filename={file.path} />
              )}
            </For>
          </div>
        </section>
      </article>

      <aside class="hidden lg:block w-64">
        <SectionNav sections={sections()} />
      </aside>
    </div>
  );
}
