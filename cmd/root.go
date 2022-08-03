package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dignore",
	Short: "dignore list paths included / excluded by your dockerignore files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cobra ran")
	},
}

func Execute() {
	if err := rootCmd.Help(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
