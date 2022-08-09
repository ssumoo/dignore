package cmd

import (
	"os"

	"github/ssumoo/dignore/list_backend"

	"github.com/spf13/cobra"
)

func init() {
	cwd, _ := os.Getwd()
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().BoolVarP(&include, "include", "i", false, "set to show only included files")
	listCmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "set to show only included + excluded files")
	listCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "show paths only, don't show reason why/why not a path is in/ex-cluded")
	listCmd.PersistentFlags().StringVarP(&rootPath, "path", "p", cwd, "root path to list from")
	listCmd.PersistentFlags().StringVarP(&dockerignorePath, "dockerignore", "d", ".dockerignore", "path to dockerignore")
}

var include bool
var all bool
var quiet bool
var rootPath string
var dockerignorePath string

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list file names",
	Aliases: []string{"ls"},
	Long:    "list file names long version",
	Run: func(cmd *cobra.Command, args []string) {

		printFilter := list_backend.PrintExclude
		if all {
			printFilter = list_backend.PrintAll
		} else if include {
			printFilter = list_backend.PrintInclude
		}
		list_backend.List(
			dockerignorePath,
			rootPath,
			!quiet,
			printFilter,
		)
	},
}
