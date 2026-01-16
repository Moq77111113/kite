package template

import (
	"strings"

	"github.com/moq77111113/kite/internal/domain/models"
)

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) ExtractVariables(content string) ([]string, error) {
	tokens, err := Parse(content)
	if err != nil {
		return nil, err
	}
	return ExtractVariableNames(tokens), nil
}

func (e *Engine) ExtractFromFiles(files []models.File) []string {
	seen := make(map[string]bool)
	var vars []string

	for _, file := range files {
		if IsBinaryContent(file.Content) {
			continue
		}
		tokens, err := Parse(file.Content)
		if err != nil {
			continue
		}
		for _, name := range ExtractVariableNames(tokens) {
			if !seen[name] {
				vars = append(vars, name)
				seen[name] = true
			}
		}
	}
	return vars
}

func (e *Engine) Interpolate(content string, values map[string]string, metadata []models.Variable) (string, error) {
	tokens, err := Parse(content)
	if err != nil {
		return "", err
	}

	lookup := buildLookup(values, metadata)
	if err := validateRequired(tokens, lookup, metadata); err != nil {
		return "", err
	}
	return render(tokens, lookup), nil
}

func render(tokens []Token, lookup map[string]string) string {
	var result strings.Builder
	for _, t := range tokens {
		switch t.Type {
		case TokenText:
			result.WriteString(t.Value)
		case TokenVariable:
			if val, ok := lookup[t.Value]; ok {
				result.WriteString(val)
			}
		}
	}
	return result.String()
}
