package template

func IsBinaryContent(content string) bool {
	checkSize := min(8000, len(content))
	for i := range checkSize {
		if content[i] == 0 {
			return true
		}
	}
	return false
}
