package cmd

import (
	"fmt"
	"sgit/internal/repo"

	"github.com/spf13/cobra"
)

var (
	logLength int
)

var logCmd = &cobra.Command{
	Use:   "log [file]",
	Short: "Add a file to version control",
	Long:  `Add a specified file to the version control system.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		fmt.Println("Log file:", file)
		fmt.Println("Log Length:", logLength)
		err := repo.FileLog(file, logLength)
		if err != nil {
			fmt.Println("Error log file:", err)
		}
	},
}

func init() {
	logCmd.Flags().IntVarP(&logLength, "length", "l", 2, "Number of logs to display")
}
