package template

import "testing"

func TestParse_SimpleText(t *testing.T) {
	tokens, err := Parse("hello world")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 || tokens[0].Type != TokenText || tokens[0].Value != "hello world" {
		t.Errorf("expected text token 'hello world', got %+v", tokens)
	}
}

func TestParse_SingleVariable(t *testing.T) {
	tokens, err := Parse("[[name]]")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 || tokens[0].Type != TokenVariable || tokens[0].Value != "name" {
		t.Errorf("expected variable token 'name', got %+v", tokens)
	}
}

func TestParse_VariableWithWhitespace(t *testing.T) {
	tokens, err := Parse("[[ name ]]")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 || tokens[0].Value != "name" {
		t.Errorf("expected trimmed variable 'name', got %+v", tokens)
	}
}

func TestParse_MixedContent(t *testing.T) {
	tokens, err := Parse("Hello [[name]], welcome!")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
}

func TestParse_EscapeSequence(t *testing.T) {
	tokens, err := Parse("Use \\[[literal]] for brackets")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 || tokens[0].Value != "Use [[literal]] for brackets" {
		t.Errorf("expected escaped text, got %+v", tokens)
	}
}

func TestParse_DottedVariables(t *testing.T) {
	tokens, err := Parse("[[repo.name]] and [[project.version]]")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d: %+v", len(tokens), tokens)
	}
	if tokens[0].Value != "repo.name" || tokens[2].Value != "project.version" {
		t.Errorf("expected dotted vars, got %+v", tokens)
	}
}

func TestParse_InvalidPatternsKeptAsText(t *testing.T) {
	tests := []string{"[[123abc]]", "[[my var]]", "[[]]", "Hello [[name"}
	for _, test := range tests {
		tokens, err := Parse(test)
		if err != nil {
			t.Errorf("unexpected error for %q: %v", test, err)
		}
		hasVar := false
		for _, tok := range tokens {
			if tok.Type == TokenVariable {
				hasVar = true
			}
		}
		if hasVar {
			t.Errorf("expected no variables for invalid %q, got %+v", test, tokens)
		}
	}
}

func TestParse_GoreleaserTemplate(t *testing.T) {
	content := `project_name: [[ repo.name]]
release:
  name: [[ repo.name]]
  header: |
    ## [[ project.name ]] {{ .Version }}`

	tokens, err := Parse(content)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	varCount := 0
	for _, tok := range tokens {
		if tok.Type == TokenVariable {
			varCount++
		}
	}
	if varCount != 3 {
		t.Errorf("expected 3 variables, got %d", varCount)
	}
}
