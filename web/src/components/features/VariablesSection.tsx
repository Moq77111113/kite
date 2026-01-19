import { CopyButton } from '@/components/ui/CopyButton'
import type { Variable } from '@/types/kit'
import { For, Show } from 'solid-js'

interface Props {
  kitId: string
  variables: Variable[]
  values: Record<string, string>
  onValuesChange: (values: Record<string, string>) => void
}

export function VariablesSection(props: Props) {
  const values = () => props.values

  const handleChange = (name: string, value: string) => {
    props.onValuesChange({ ...props.values, [name]: value })
  }

  const command = () => {
    const parts = [`kite add ${props.kitId}`]
    for (const v of props.variables) {
      const val = values()[v.name]
      if (val) parts.push(`--var ${v.name}=${val}`)
    }
    return parts.join(' ')
  }

  return (
    <div class="space-y-4 min-w-0 px-2">
      <div class="text-xs font-semibold text-muted-foreground mb-3 uppercase tracking-wider">
        Variables
      </div>
      <div class="space-y-3">
        <For each={props.variables}>
          {(variable) => (
            <div>
              <label class="block text-xs font-semibold text-foreground mb-1.5">
                {variable.name}
                {variable.required && <span class="text-destructive ml-1">*</span>}
              </label>
              <Show when={variable.description}>
                <p class="text-xs text-muted-foreground mb-2">{variable.description}</p>
              </Show>
              <input
                type="text"
                value={values()[variable.name] || ''}
                onInput={(e) => handleChange(variable.name, e.target.value)}
                placeholder={variable.default || variable.name}
                class="w-full px-3 py-2 text-xs rounded-lg border border-border bg-background text-foreground placeholder:text-muted-foreground focus:outline-none focus:ring-2 focus:ring-primary/50"
              />
            </div>
          )}
        </For>
      </div>

      <div class="pt-4 border-t border-border min-w-0">
        <div class="text-xs font-semibold text-muted-foreground mb-2 uppercase tracking-wider">
          Command
        </div>
        <div class="rounded-lg border border-border p-3 flex items-start justify-between gap-2 bg-muted/30 min-w-0 overflow-hidden">
          <code class="text-xs font-mono text-foreground break-all flex-1 min-w-0 overflow-x-auto">
            {command()}
          </code>
          <CopyButton text={command()} />
        </div>
      </div>
    </div>
  )
}
