package template

type TokenType int

const (
	TokenText TokenType = iota
	TokenVariable
)

type Token struct {
	Type  TokenType
	Value string
}

func ExtractVariableNames(tokens []Token) []string {
	seen := make(map[string]bool)
	var vars []string
	for _, t := range tokens {
		if t.Type == TokenVariable && !seen[t.Value] {
			vars = append(vars, t.Value)
			seen[t.Value] = true
		}
	}
	return vars
}
