package cmd

import (
	"context"
	"fmt"
	"os"
	"sgit/internal/repo"
	"sgit/pkg/conf"

	"github.com/spf13/cobra"
)

var (
	autoYes  bool
	repoPath string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new repository",
	Long:  `Initialize a new repository at the given path.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		userInput := make(chan byte, 1)
		if autoYes {
			go func() {
				for i := 0; i < 2; i++ {
					userInput <- 'y'
				}
			}()
		} else {
			go func() {
				var input string
				for i := 0; i < 2; i++ {
					fmt.Scanln(&input)
					userInput <- input[0]
				}
			}()
		}

		if err := repo.InitRepo(ctx, repoPath, userInput); err != nil {
			fmt.Printf("Error initializing repo: %v\n", err)
			os.Exit(0)
		}

		fmt.Println("Repository initialized successfully.")
	},
}

func init() {
	defaultPath := conf.Conf.Repo.Path
	initCmd.Flags().BoolVarP(&autoYes, "yes", "y", false, "Automatically answer 'yes' to prompts")
	initCmd.Flags().StringVarP(&repoPath, "path", "p", defaultPath, "Path to initialize the repository")
}
