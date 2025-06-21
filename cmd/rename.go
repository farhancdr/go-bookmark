package cmd

import (
	"fmt"

	"github.com/farhancdr/go-bookmark/internal/bookmarks"
	"github.com/spf13/cobra"
)

func NewRenameCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rename <oldAliasName> <newAliasName>",
		Short: "Rename a bookmark alias",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			oldAlias, newAlias := args[0], args[1]
			if !isValidAlias(newAlias) {
				return fmt.Errorf("invalid new alias: must contain only alphanumeric characters, underscores, or hyphens")
			}

			store, err := bookmarks.LoadBookmarks()
			if err != nil {
				return fmt.Errorf("failed to load bookmarks: %v", err)
			}

			if _, exists := store.FindByAlias(oldAlias); !exists {
				return fmt.Errorf("alias '%s' not found", oldAlias)
			}
			if _, exists := store.FindByAlias(newAlias); exists {
				return fmt.Errorf("alias '%s' already exists", newAlias)
			}

			err = store.RenameBookmark(oldAlias, newAlias)
			if err != nil {
				return fmt.Errorf("failed to rename bookmark: %v", err)
			}

			fmt.Printf("Bookmark '%s' renamed to '%s'.\n", oldAlias, newAlias)
			return nil
		},
	}
	return cmd
}
