package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/moby/buildkit/frontend/dockerfile/dockerignore"
	"github.com/moby/patternmatcher"
	"github.com/spf13/cobra"
)

func init() {
	cwd, _ := os.Getwd()
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().BoolVarP(&printExcluded, "excluded", "e", false, "set to show only excluded files instead of included files")
	listCmd.PersistentFlags().StringVarP(&rootPath, "path", "p", cwd, "root path to list from")
	listCmd.PersistentFlags().StringVarP(&dockerignorePath, "dockerignore", "d", ".dockerignore", "path to dockerignore")
}

var printExcluded bool
var rootPath string
var dockerignorePath string

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list file names",
	Aliases: []string{"ls"},
	Long:    "list file names included with the given .dockerignore file",
	Run: func(cmd *cobra.Command, args []string) {
		printDockerIgnoredFiles(
			dockerignorePath,
			rootPath,
			printExcluded,
		)
	},
}

func printDockerIgnoredFiles(
	dockerignorePath string,
	rootPath string,
	printExcluded bool,
) {
	absDockerignorePath, err := filepath.Abs(dockerignorePath)
	if err != nil {
		log.Fatalf("can't resolve the given dockerignore path to absolute path: %s, (%s)", dockerignorePath, err)
		return
	}
	f, err := os.Open(absDockerignorePath)
	var excludeLines []string
	if err == nil {
		excludeLinesFromDockerignore, excludesErr := dockerignore.ReadAll(f)
		excludeLines = excludeLinesFromDockerignore
		if excludesErr != nil {
			log.Fatalf("error while reading .dockerignore at %s (%s)", absDockerignorePath, excludesErr)
			return
		}
	} else if errors.Is(err, os.ErrNotExist) {
		log.Fatalf("WARNING: .dockerignore file provided at %s doesn't exist, will treat this case as not excluding anything\n", absDockerignorePath)
		excludeLines = make([]string, 0)
	} else {
		log.Fatalf("can't open .dockerignore file provided at %s, (%s)", absDockerignorePath, err)
		return
	}
	defer f.Close()
	pm, matcherErr := patternmatcher.New(excludeLines)
	if matcherErr != nil {
		log.Fatalf("dockerignore file provided at %s does not lead to any valid pattern (%s)", absDockerignorePath, matcherErr)
		return
	}

	walkErr := filepath.Walk(
		rootPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatalf("error checking path (%s) against dockerignore (%s)", path, err)
				return err
			}
			if info.IsDir() {
				return nil
			}
			relPath, err := filepath.Rel(rootPath, path)
			if err != nil {
				log.Fatalf("error computing relative path of %s wrt root path %s", path, rootPath)
				return err
			}
			pathIsExcluded, err := pm.MatchesOrParentMatches(relPath)
			if err != nil {
				log.Fatalf("pattern matcher errored against path %s (%s)", path, err)
				return err
			}
			if pathIsExcluded == printExcluded {
				fmt.Printf("%s\n", relPath)
			}
			return nil
		})
	if walkErr != nil {
		log.Fatalf("error walking directory: (%s) (%s)", rootPath, walkErr)
		return
	}
}
