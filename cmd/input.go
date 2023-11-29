package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var inputFilePath string

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Reads file input as stdin",
	Run: func(cmd *cobra.Command, args []string) {
		if inputFilePath == "" {
			fmt.Println("Please provide an input file using --input or -i")
			return
		}
		// Open the input file
		file, err := os.Open(inputFilePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Read and print the contents of the file
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	},
}

func init() {
	inputCmd.Flags().StringVarP(&inputFilePath, "input", "i", "", "Input file path")
}
