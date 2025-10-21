package annotations

type Annotation string

const (
	SkipContainer Annotation = "kite.dev/skip-container"
)

func (a Annotation) String() string {
	return string(a)
}
