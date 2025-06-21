package cmd

import (
	"fmt"
	"os"

	"github.com/farhancdr/go-bookmark/internal/bookmarks"
	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update <aliasName> [newPath]",
		Short: "Update a bookmark's directory path",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			alias := args[0]
			newPath := ""
			if len(args) == 2 {
				newPath = args[1]
			} else {
				var err error
				newPath, err = os.Getwd()
				if err != nil {
					return fmt.Errorf("failed to get current directory: %v", err)
				}
			}

			store, err := bookmarks.LoadBookmarks()
			if err != nil {
				return fmt.Errorf("failed to load bookmarks: %v", err)
			}

			if _, exists := store.FindByAlias(alias); !exists {
				return fmt.Errorf("alias '%s' not found", alias)
			}

			err = store.UpdateBookmark(alias, newPath)
			if err != nil {
				return fmt.Errorf("failed to update bookmark: %v", err)
			}

			fmt.Printf("Bookmark '%s' updated to '%s'.\n", alias, newPath)
			return nil
		},
	}
	return cmd
}
