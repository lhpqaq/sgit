package repo

import (
	"context"
	"fmt"
	"sgit/pkg/repo"
	"sgit/utils/paths"
)

func InitRepo(ctx context.Context, path string, userInput chan string) error {
	if exist, _ := paths.PathExists(path); !exist {
		fmt.Printf("Directory %s does not exist. Do you want to create it? (y/n): ", path)

		select {
		case response := <-userInput:
			if response == "y" {
				err := paths.EnsureDirExists(path)
				if err != nil {
					fmt.Printf("Failed to create directory: %v\n", err)
				} else {
					fmt.Println("Directory created successfully.")
				}
			} else {
				fmt.Println("Directory creation canceled.")
			}
		case <-ctx.Done():
			fmt.Println("Operation canceled.")
			return nil
		}
	}
	fmt.Printf("Create a repository in %s ? (y/n): ", path)
	select {
	case response := <-userInput:
		if response == "y" {
			err := repo.PlainInit(path)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("canceled.")
		}
	case <-ctx.Done():
		fmt.Println("Operation canceled.")
		return nil
	}
	return nil
}
