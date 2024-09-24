package cmd

import (
	"fmt"
	"sgit/internal/repo"

	"github.com/spf13/cobra"
)

var (
	addOnly bool
	message string
)
var addCmd = &cobra.Command{
	Use:   "add [file]",
	Short: "Add a file to version control",
	Long:  `Add a specified file to the version control system.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		err := repo.AddFile(file, message, !addOnly)
		if err != nil {
			fmt.Println("Error adding file:", err)
		} else {
			fmt.Println("File added:", file)
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&message, "message", "m", "", "Commit message")
	addCmd.Flags().BoolVarP(&addOnly, "addOnly", "u", false, "Commit the file after adding")
}
