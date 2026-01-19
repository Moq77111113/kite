package template

import (
	"fmt"
	"strings"
)

type ParseError struct {
	Position int
	Message  string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("template parse error at position %d: %s", e.Position, e.Message)
}

type MissingVariablesError struct {
	Variables []string
}

func (e *MissingVariablesError) Error() string {
	return fmt.Sprintf("missing required variables: %s", strings.Join(e.Variables, ", "))
}
