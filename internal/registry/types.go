package registry

type Registry struct {
	Name		string			`json:"name" yaml:"name"`
	Homepage	string			`json:"homepage" yaml:"homepage"`
	Modules		[]IndexEntry	`json:"items" yaml:"items"`
}


type IndexEntry struct {
	Name			string				`json:"name" yaml:"name"`
	Type			string				`json:"type" yaml:"type"`
	Description		string				`json:"description,omitempty" yaml:"description,omitempty"`
	Author			string				`json:"author,omitempty" yaml:"author,omitempty"`
	RegistryDeps	[]string			`json:"registryDependencies,omitempty" yaml:"registryDependencies,omitempty"`
	Dependencies	[]string			`json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Tags			[]string			`json:"tags,omitempty" yaml:"tags,omitempty"`
	Meta			map[string]any		`json:"meta,omitempty" yaml:"meta,omitempty"`

}


type ModuleFile struct {
	Path    string `json:"path" yaml:"path"`              
	Content string `json:"content" yaml:"content"`         
	Type    string `json:"type,omitempty" yaml:"type,omitempty"` 
	Target  string `json:"target,omitempty" yaml:"target,omitempty"` 
}

type Module struct {
	*IndexEntry
	Files		[]ModuleFile			`json:"files" yaml:"files"`
}

