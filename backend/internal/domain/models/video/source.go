package video

type SourceType string

const (
	SourceUrl  SourceType = "url"
	SourceFile SourceType = "file"
)

func (s SourceType) isValid() bool {
	switch s {
	case SourceUrl, SourceFile:
		return true
	default:
		return false
	}
}
