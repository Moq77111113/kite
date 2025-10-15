export interface TemplateSummary {
  name: string
  version: string
  description: string
  tags: string[]
  author: string
}

export interface TemplateFile {
  path: string
  content: string
}

export interface TemplateDetail extends TemplateSummary {
  files: TemplateFile[]
  readme: string
}

export interface TemplatesResponse {
  templates: TemplateSummary[]
}
