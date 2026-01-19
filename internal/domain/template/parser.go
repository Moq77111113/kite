package template

import (
	"regexp"
	"strings"
)

var varPattern = regexp.MustCompile(`\[\[\s*([a-zA-Z_][a-zA-Z0-9_.\-]*)\s*\]\]`)

func Parse(content string) ([]Token, error) {
	escaped := strings.ReplaceAll(content, "\\[[", "\x00ESCAPED_OPEN\x00")
	matches := varPattern.FindAllStringSubmatchIndex(escaped, -1)

	if len(matches) == 0 {
		return []Token{{Type: TokenText, Value: unescape(escaped)}}, nil
	}

	var tokens []Token
	pos := 0

	for _, match := range matches {
		start, end := match[0], match[1]
		varStart, varEnd := match[2], match[3]

		if start > pos {
			tokens = append(tokens, Token{Type: TokenText, Value: unescape(escaped[pos:start])})
		}
		tokens = append(tokens, Token{Type: TokenVariable, Value: escaped[varStart:varEnd]})
		pos = end
	}

	if pos < len(escaped) {
		tokens = append(tokens, Token{Type: TokenText, Value: unescape(escaped[pos:])})
	}

	return tokens, nil
}

func unescape(s string) string {
	return strings.ReplaceAll(s, "\x00ESCAPED_OPEN\x00", "[[")
}
