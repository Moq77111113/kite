package registry

type Registry struct {
	Name		string			`json:"name" yaml:"name"`
	Homepage	string			`json:"homepage" yaml:"homepage"`
	Items		[]RegistryItem	`json:"items" yaml:"items"`
}


type RegistryItem struct {
	Name			string				`json:"name" yaml:"name"`
	Type			string				`json:"type" yaml:"type"`
	Title			string				`json:"title,omitempty" yaml:"title,omitempty"`
	Description		string				`json:"description,omitempty" yaml:"description,omitempty"`
	Author			string				`json:"author,omitempty" yaml:"author,omitempty"`
	Dependencies	[]string			`json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Files			[]string			`json:"files,omitempty" yaml:"files,omitempty"`
	Meta			map[string]any		`json:"meta,omitempty" yaml:"meta,omitempty"`
	Tags			[]string			`json:"tags,omitempty" yaml:"tags,omitempty"`
}


