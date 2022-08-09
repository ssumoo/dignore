package list_backend

import (
	"fmt"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
)

func matchLine(path string, line string) bool {
	path = filepath.ToSlash(path)
	line = filepath.ToSlash(line)
	vanillaMatched := vanillaMatch(path, line)
	if vanillaMatched {
		return vanillaMatched
	}
	addedStarMatched := vanillaMatch(path, filepath.Clean(fmt.Sprintf("%s/**", line)))
	if addedStarMatched {
		return addedStarMatched
	}
	return false
}

func vanillaMatch(path string, line string) bool {
	matched, err := doublestar.Match(line, path)
	if err != nil {
		return false
	}
	return matched
}
