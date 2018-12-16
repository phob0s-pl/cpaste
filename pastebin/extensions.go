package pastebin

import (
	"path/filepath"
)

// GetFileFormat returns format of the paste for give file
func GetFileFormat(path string) string { // nolint: gocyclo
	extension := filepath.Ext(path)
	switch extension {
	case ".go":
		return "go"
	case ".sh":
		return "bash"
	case ".c":
		return "c"
	case ".cpp":
		return "cpp"
	case ".html":
		return "html5"
	case ".ini":
		return "ini"
	case ".java":
		return "java"
	case ".md":
		return "markdown"
	case ".pl":
		return "perl"
	case ".py":
		return "python"
	case ".xml":
		return "xml"
	case ".js":
		return "javascript"
	default:
		return "text"
	}
}
