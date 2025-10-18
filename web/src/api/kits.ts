import type { KitsResponse, KitDetail } from '../types/kit'

export async function fetchKits(tag?: string): Promise<KitsResponse> {
  const url = tag ? `/api/kits?tag=${encodeURIComponent(tag)}` : '/api/kits'
  const response = await fetch(url)

  if (!response.ok) {
    throw new Error(`Failed to fetch kits: ${response.statusText}`)
  }

  return response.json()
}

export async function fetchKit(name: string): Promise<KitDetail> {
  const response = await fetch(`/api/kits/${name}`)

  if (!response.ok) {
    throw new Error(`Failed to fetch kit: ${response.statusText}`)
  }

  return response.json()
}
