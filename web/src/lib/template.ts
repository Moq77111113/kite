const VAR_PATTERN = /\[\[\s*([a-zA-Z_][a-zA-Z0-9_.\-]*)\s*\]\]/g

const ESCAPE_PLACEHOLDER = '\x00ESCAPED_OPEN\x00'

export interface Token {
  type: 'text' | 'variable'
  value: string
}

export function parse(content: string): Token[] {
  const escaped = content.replaceAll('\\[[', ESCAPE_PLACEHOLDER)

  const tokens: Token[] = []
  let lastIndex = 0

  VAR_PATTERN.lastIndex = 0

  let match: RegExpExecArray | null
  while ((match = VAR_PATTERN.exec(escaped)) !== null) {
    if (match.index > lastIndex) {
      tokens.push({
        type: 'text',
        value: unescape(escaped.slice(lastIndex, match.index)),
      })
    }

    tokens.push({
      type: 'variable',
      value: match[1],
    })

    lastIndex = match.index + match[0].length
  }


  if (lastIndex < escaped.length) {
    tokens.push({
      type: 'text',
      value: unescape(escaped.slice(lastIndex)),
    })
  }


  if (tokens.length === 0) {
    tokens.push({ type: 'text', value: content })
  }

  return tokens
}

function unescape(s: string): string {
  return s.replaceAll(ESCAPE_PLACEHOLDER, '[[')
}

export function extractVariables(content: string): string[] {
  const tokens = parse(content)
  const seen = new Set<string>()
  const vars: string[] = []

  for (const token of tokens) {
    if (token.type === 'variable' && !seen.has(token.value)) {
      vars.push(token.value)
      seen.add(token.value)
    }
  }

  return vars
}

export function interpolate(
  content: string,
  values: Record<string, string>,
  defaults: Record<string, string> = {}
): string {
  const tokens = parse(content)

  const lookup: Record<string, string> = { ...defaults, ...values }

  return tokens
    .map((token) => {
      if (token.type === 'text') {
        return token.value
      }
      const resolved = lookup[token.value]
      if (resolved !== undefined && resolved !== '') {
        return resolved
      }
      return `[[${token.value}]]`
    })
    .join('')
}

export interface InterpolatedSegment {
  type: 'text' | 'replaced' | 'unresolved'
  value: string
  variableName?: string
}


export function interpolateWithSegments(
  content: string,
  values: Record<string, string>,
  defaults: Record<string, string> = {}
): InterpolatedSegment[] {
  const tokens = parse(content)
  const lookup: Record<string, string> = { ...defaults, ...values }

  return tokens.map((token): InterpolatedSegment => {
    if (token.type === 'text') {
      return { type: 'text', value: token.value }
    }

    const resolved = lookup[token.value]
    if (resolved !== undefined && resolved !== '') {
      return {
        type: 'replaced',
        value: resolved,
        variableName: token.value,
      }
    }


    return {
      type: 'unresolved',
      value: `[[${token.value}]]`,
      variableName: token.value,
    }
  })
}
