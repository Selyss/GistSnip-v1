package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of GistSnip",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GistSnip v0.9")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
