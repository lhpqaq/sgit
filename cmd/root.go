package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sgit",
	Short: "A simple file version control tool",
	Long:  `sgit is a simple file version control tool built using Go and git concepts.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sgit - single file version control tool")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(diffCmd)
}
