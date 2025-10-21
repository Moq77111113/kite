export interface KitSummary {
  name: string
  version: string
  description: string
  tags: string[]
  author: string
  lastUpdated?: string
}

export interface KitFile {
  path: string
  content: string
  lastUpdated?: string
}

export interface KitDetail extends KitSummary {
  files: KitFile[]
  readme: string
}

export interface KitsResponse {
  kits: KitSummary[]
}
