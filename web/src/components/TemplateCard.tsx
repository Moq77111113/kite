import { Link } from '@tanstack/solid-router'
import type { TemplateSummary } from '../types/template'

interface TemplateCardProps {
  template: TemplateSummary
}

export default function TemplateCard(props: TemplateCardProps) {
  const primaryTag = () => props.template.tags[0] || 'general'

  // Generate a color based on the tag
  const tagColor = () => {
    const colors = [
      'from-blue-500 to-blue-600',
      'from-purple-500 to-purple-600',
      'from-pink-500 to-pink-600',
      'from-red-500 to-red-600',
      'from-orange-500 to-orange-600',
      'from-yellow-500 to-yellow-600',
      'from-green-500 to-green-600',
      'from-teal-500 to-teal-600',
      'from-cyan-500 to-cyan-600',
      'from-indigo-500 to-indigo-600',
    ]
    const index = primaryTag().split('').reduce((acc, char) => acc + char.charCodeAt(0), 0)
    return colors[index % colors.length]
  }

  return (
    <Link
      to="/templates/$name"
      params={{ name: props.template.name }}
      class="block rounded-xl border border-border bg-card p-5 transition-all hover:border-accent hover:shadow-lg group"
    >
      <div class="flex items-start justify-between mb-3">
        <span class="px-2.5 py-1 text-xs rounded-md bg-muted/50 text-muted-foreground capitalize font-medium">
          {primaryTag()}
        </span>
        <div class={`w-14 h-14 rounded-2xl bg-gradient-to-br ${tagColor()} flex items-center justify-center text-white shadow-md group-hover:scale-105 transition-transform`}>
          <span class="text-2xl font-bold">
            {props.template.name.charAt(0).toUpperCase()}
          </span>
        </div>
      </div>

      <h3 class="text-lg font-bold text-card-foreground mb-2 group-hover:text-primary transition-colors">
        {props.template.name}
      </h3>

      <p class="text-sm text-muted-foreground mb-4 line-clamp-2 leading-relaxed">
        {props.template.description}
      </p>

      <div class="flex items-center gap-3 text-xs text-muted-foreground pt-2 border-t border-border/50">
        <span class="font-medium">v{props.template.version}</span>
        <span class="opacity-50">•</span>
        <span>by {props.template.author}</span>
      </div>
    </Link>
  )
}
