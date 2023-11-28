package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of GistSnip",
	Long:  `All software has versions. This is GistSnip's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GistSnip v0.9")
	},
}
