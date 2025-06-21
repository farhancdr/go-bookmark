package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/farhancdr/go-bookmark/internal/bookmarks"
	"github.com/spf13/cobra"
)

func NewSaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save <aliasName>",
		Short: "Save the current directory with an alias",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			alias := args[0]
			if !isValidAlias(alias) {
				return fmt.Errorf("invalid alias: must contain only alphanumeric characters, underscores, or hyphens")
			}

			currentDir, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get current directory: %v", err)
			}

			store, err := bookmarks.LoadBookmarks()
			if err != nil {
				return fmt.Errorf("failed to load bookmarks: %v", err)
			}

			if existing, exists := store.FindByAlias(alias); exists {
				fmt.Printf("Alias '%s' already exists for '%s'. Overwrite? (y/N): ", alias, existing.Path)
				if !confirm() {
					fmt.Println("Operation cancelled.")
					return nil
				}
			}

			err = store.AddBookmark(alias, currentDir)
			if err != nil {
				return fmt.Errorf("failed to save bookmark: %v", err)
			}

			fmt.Printf("Directory '%s' bookmarked as '%s'.\n", currentDir, alias)
			return nil
		},
	}
	return cmd
}

func isValidAlias(alias string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, alias)
	return matched
}

func confirm() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		return false
	}
	return response == "y" || response == "Y"
}
