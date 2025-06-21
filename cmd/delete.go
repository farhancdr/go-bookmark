package cmd

import (
	"fmt"

	"github.com/farhancdr/go-bookmark/internal/bookmarks"
	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <aliasName>",
		Short: "Delete a bookmark",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			alias := args[0]
			store, err := bookmarks.LoadBookmarks()
			if err != nil {
				return fmt.Errorf("failed to load bookmarks: %v", err)
			}

			if _, exists := store.FindByAlias(alias); !exists {
				return fmt.Errorf("alias '%s' not found", alias)
			}

			fmt.Printf("Are you sure you want to delete bookmark '%s'? (y/N): ", alias)
			if !confirm() {
				fmt.Println("Operation cancelled.")
				return nil
			}

			err = store.DeleteBookmark(alias)
			if err != nil {
				return fmt.Errorf("failed to delete bookmark: %v", err)
			}

			fmt.Printf("Bookmark '%s' deleted.\n", alias)
			return nil
		},
	}
	return cmd
}
