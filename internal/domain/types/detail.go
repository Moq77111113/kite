package types

type KitDetailResponse struct {
	Name        string     `json:"name"`
	Version     string     `json:"version"`
	Author      string     `json:"author"`
	Description string     `json:"description"`
	Files       []KitFile  `json:"files"`
	Variables   []Variable `json:"variables"`
	Readme      string     `json:"readme"`
	Tags        []string   `json:"tags"`
}
