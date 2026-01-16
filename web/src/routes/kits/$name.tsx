import { fetchKit } from "@/api/kits";
import { SectionNav } from "@/components/features/SectionNav";
import { VariablesSection } from "@/components/features/VariablesSection";
import { CodeBlock } from "@/components/ui/CodeBlock";
import { CopyButton } from "@/components/ui/CopyButton";
import { MarkdownRenderer } from "@/components/ui/Markdown";
import { createFileRoute } from "@tanstack/solid-router";

import { For, Show, createMemo } from "solid-js";

export const Route = createFileRoute("/kits/$name")({
  loader: ({ params }) => fetchKit(params.name),
  component: KitDetail,
});

function KitDetail() {
  const kit = Route.useLoaderData();
  const hasVars = () => !!kit().variables?.length;
  const sections = createMemo(() => {
    const secs = [{ id: "install", label: "Installation" }];
    if (kit().readme) {
      secs.push({ id: "readme", label: "README" });
    }
    secs.push({ id: "files", label: "Files" });
    return secs;
  });

  return (
    <div class="w-full h-[calc(100vh-4rem)] overflow-hidden">
      <div class="flex flex-col lg:flex-row gap-6 lg:gap-8 w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 h-full">
        <article class="flex-1 min-w-0 w-full px-2 lg:w-auto overflow-y-auto py-8 sm:py-12">
          <header class="mb-10">
            <div class="flex items-center gap-3 mb-3">
              <h1 class="text-3xl sm:text-4xl font-bold text-foreground">
                {kit().name}
              </h1>
              <span class="px-3 py-1.5 text-xs font-semibold rounded-lg bg-muted/50 text-muted-foreground">
                v{kit().version}
              </span>
            </div>
            <p class="text-base text-muted-foreground mb-4 leading-relaxed">
              {kit().description}
            </p>

            <div class="flex flex-wrap gap-2 mb-4">
              <For each={kit().tags}>
                {(tag) => (
                  <span class="px-2.5 py-1 text-xs font-medium rounded-md bg-accent/50 text-accent-foreground capitalize">
                    {tag}
                  </span>
                )}
              </For>
            </div>

            <div class="flex items-center gap-4 text-sm text-muted-foreground mb-6">
              <span>by {kit().author}</span>
            </div>

            <div
              id="install"
              class="rounded-xl border border-border p-5 flex items-center justify-between shadow-sm bg-code-inline "
            >
              <code class="text-sm font-mono text-card-foreground font-medium break-all">
                kite add {kit().id}
              </code>
              <CopyButton text={`kite add ${kit().id}`} />
            </div>
          </header>

          <div class="lg:hidden mb-10 space-y-6">
            <div>
              <SectionNav sections={sections()} />
            </div>
            <Show when={hasVars()}>
              <VariablesSection kitId={kit().id} variables={kit().variables} />
            </Show>
          </div>

          <Show when={kit().readme}>
            <section id="readme" class="mb-10">
              <h2 class="text-2xl font-bold text-foreground mb-5">README</h2>
              <div class="rounded-xl border border-border bg-card p-6 shadow-sm">
                <MarkdownRenderer markdown={kit().readme} />
              </div>
            </section>
          </Show>

          <section id="files" class="min-w-0 overflow-hidden">
            <h2 class="text-2xl font-bold text-foreground mb-5">Files</h2>
            <div class="space-y-4">
              <For each={kit().files}>
                {(file) => (
                  <CodeBlock content={file.content} filename={file.path} />
                )}
              </For>
            </div>
          </section>
        </article>

        <aside class="hidden lg:block lg:w-72 lg:shrink-0 overflow-y-auto py-8 sm:py-12">
          <div class="space-y-6">
            <div>
              <SectionNav sections={sections()} />
            </div>
            <Show when={hasVars()}>
              <VariablesSection kitId={kit().id} variables={kit().variables} />
            </Show>
          </div>
        </aside>
      </div>
    </div>
  );
}
