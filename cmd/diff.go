package cmd

import (
	"fmt"
	"sgit/internal/repo"

	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff [file]",
	Short: "Add a file to version control",
	Long:  `Add a specified file to the version control system.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		err := repo.DiffFile(file)
		if err != nil {
			fmt.Println("Error diff file:", err)
		}
	},
}

func init() {

}
