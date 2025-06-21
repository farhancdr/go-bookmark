package cmd

import (
	"fmt"
	"os"

	"github.com/farhancdr/go-bookmark/internal/bookmarks"
	"github.com/spf13/cobra"
)

func NewGotoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goto <aliasName>",
		Short: "Navigate to a bookmarked directory",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			alias := args[0]
			store, err := bookmarks.LoadBookmarks()
			if err != nil {
				return fmt.Errorf("failed to load bookmarks: %v", err)
			}

			bookmark, exists := store.FindByAlias(alias)
			if !exists {
				return fmt.Errorf("alias '%s' not found", alias)
			}

			if _, err := os.Stat(bookmark.Path); os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "Warning: directory '%s' no longer exists. Consider removing bookmark '%s'.\n", bookmark.Path, alias)
				return fmt.Errorf("directory does not exist")
			}

			fmt.Println(bookmark.Path)
			return nil
		},
	}
	return cmd
}
