export interface KitSummary {
  id: string
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
  lastSync?: string
}

export interface SyncResponse {
  status: string
  lastSync?: string
}
