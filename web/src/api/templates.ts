import type { TemplatesResponse, TemplateDetail } from '../types/template'

export async function fetchTemplates(tag?: string): Promise<TemplatesResponse> {
  const url = tag ? `/api/templates?tag=${encodeURIComponent(tag)}` : '/api/templates'
  const response = await fetch(url)

  if (!response.ok) {
    throw new Error(`Failed to fetch templates: ${response.statusText}`)
  }

  return response.json()
}

export async function fetchTemplate(name: string): Promise<TemplateDetail> {
  const response = await fetch(`/api/templates/${name}`)

  if (!response.ok) {
    throw new Error(`Failed to fetch template: ${response.statusText}`)
  }

  return response.json()
}
