import { createFileRoute } from "@tanstack/solid-router";

export const Route = createFileRoute("/")({
  component: Home,
});

function Home() {
  return (
    <div class="flex items-center justify-center min-h-screen px-8">
      <div class="max-w-2xl text-center">
        <h1 class="text-4xl font-bold text-foreground mb-4">Welcome to Kite</h1>
        <p class="text-lg text-muted-foreground mb-8">
          Fork your infrastructure, don't worship it. Browse templates from the
          sidebar and copy them into your project as editable code.
        </p>

        <div class="rounded-lg border border-border bg-card p-6 text-left">
          <h2 class="text-xl font-semibold text-card-foreground mb-3">
            Getting Started
          </h2>
          <ol class="space-y-2 text-sm text-muted-foreground">
            <li class="flex gap-3">
              <span class="font-semibold text-foreground">1.</span>
              <span>Select a template from the sidebar</span>
            </li>
            <li class="flex gap-3">
              <span class="font-semibold text-foreground">2.</span>
              <span>Review the template files and README</span>
            </li>
            <li class="flex gap-3">
              <span class="font-semibold text-foreground">3.</span>
              <span>
                Copy the
                <code class="px-1 py-0.5 rounded bg-muted text-muted-foreground">
                  kite add
                </code>
                command
              </span>
            </li>
            <li class="flex gap-3">
              <span class="font-semibold text-foreground">4.</span>
              <span>Run it in your project directory</span>
            </li>
          </ol>
        </div>
      </div>
    </div>
  );
}
