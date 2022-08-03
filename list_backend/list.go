package list_backend

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func List(
	dockerignorePath string,
	rootPath string,
	explain bool,
	printFilter PrintFilter,
) {
	lines, err := readDockerignoreFile(dockerignorePath)
	if err != nil {
		fmt.Println("error matching file")
		return
	}

	files := make([]string, 0)
	walkErr := filepath.Walk(
		rootPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			files = append(files, path)
			return nil
		})
	if walkErr != nil {
		fmt.Println("error listing files under directory")
		return
	}

	for _, path := range files {
		relPath, err := filepath.Rel(rootPath, path)
		if err != nil {
			fmt.Printf("error computing rel path for %s from %s\n", path, rootPath)
			return
		}
		matchRes := checkPathAgainstDockerignore(relPath, lines)
		printMatches(path, matchRes, printFilter, explain)
	}
}

func readDockerignoreFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{""}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			continue
		}
		if line == "." {
			// match Docker's historical reasons note
			continue
		}
		if len(line) == 0 {
			continue
		}
		content = append(content, line)
	}

	if err := scanner.Err(); err != nil {
		return []string{""}, err
	}

	return content, err
}

func checkPathAgainstDockerignore(path string, dockerignoreLines []string) matchResult {
	for i := len(dockerignoreLines) - 1; i >= 0; i-- {
		line := dockerignoreLines[i]
		res := matchResult{
			line,
			checkPathAgainstLine(path, line),
		}
		if res.mode != matchNone {
			return res
		}
	}
	return matchResult{"", matchNone}
}

func checkPathAgainstLine(path string, line string) matchMode {
	path = filepath.Clean(path)
	checkLine := line
	if strings.HasPrefix(line, "!") {
		checkLine = line[1:]
	}
	checkLine = filepath.Clean(checkLine)
	matched := matchLine(path, checkLine)
	if !matched {
		return matchNone
	}
	if strings.HasPrefix(line, "!") {
		return matchInclude
	} else {
		return matchExclude
	}
}