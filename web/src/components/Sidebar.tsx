import { Link } from '@tanstack/solid-router'
import { Show, For } from 'solid-js'
import type { TemplatesResponse } from '../types/template'

interface SidebarProps {
  templates: TemplatesResponse
}

export default function Sidebar(props: SidebarProps) {

  return (
    <aside class="fixed left-0 top-0 h-screen w-64 border-r border-sidebar-border bg-sidebar">
      <div class="flex h-full flex-col">
        <div class="flex h-16 items-center border-b border-sidebar-border px-6">
          <Link to="/" class="text-xl font-bold text-sidebar-foreground hover:opacity-80">
            Kite
          </Link>
        </div>

        <nav class="flex-1 overflow-y-auto p-4">
          <div class="mb-4">
            <h2 class="px-3 text-xs font-semibold uppercase tracking-wider text-sidebar-foreground opacity-60 mb-2">
              Templates
            </h2>
            <Show when={props.templates.templates}>
              <div class="space-y-1">
                <For each={props.templates.templates}>
                  {(template) => (
                    <Link
                      to="/templates/$name"
                      params={{ name: template.name }}
                      class="block rounded-lg px-3 py-2 text-sm font-medium text-sidebar-foreground transition-colors hover:bg-sidebar-accent hover:text-sidebar-accent-foreground"
                      activeProps={{
                        class: "bg-sidebar-accent text-sidebar-accent-foreground"
                      }}
                    >
                      {template.name}
                    </Link>
                  )}
                </For>
              </div>
            </Show>
          </div>
        </nav>

        <div class="border-t border-sidebar-border p-4">
          <p class="text-xs text-sidebar-foreground opacity-60">
            Infrastructure templates you can fork
          </p>
        </div>
      </div>
    </aside>
  )
}
