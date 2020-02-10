package devicons

import (
	"path/filepath"
	"strings"
)

func Icon(path string) rune {
	filename := strings.ToLower(filepath.Base(path))
	if icon, found := exactMatches[filename]; found {
		return icon
	}

	ext := strings.TrimPrefix(filepath.Ext(filename), ".")
	if icon, found := extensions[ext]; found {
		return icon
	}
	return ''
}
