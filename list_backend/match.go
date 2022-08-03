package list_backend

import (
	"path/filepath"
	"strings"
)

func matchLine(path string, line string) bool {
	if strings.Contains(line, "**/") {
		return matchDoubleStar(path, line)
	} else if strings.HasPrefix(line, "*") {
		return matchFileExtension(path, line)
	} else {
		return vanillaMatch(path, line)
	}
}

func vanillaMatch(path string, line string) bool {
	matched, err := filepath.Match(line, path)
	if err != nil {
		return false
	}
	return matched
}

func matchFileExtension(path string, line string) bool {
	return filepath.Ext(path) == line[1:]
}

func matchDoubleStar(path string, line string) bool {
	doubleStarIdx := strings.Index(line, "**/")
	beforeDoubleStar := line[:doubleStarIdx]
	afterDoubleStar := line[doubleStarIdx+3:]
	if !strings.HasPrefix(path, beforeDoubleStar) {
		return false
	}
	return matchLine(path, afterDoubleStar)
}
