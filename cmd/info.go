package cmd

import (
	"fmt"
	"os"

	"github.com/farhancdr/go-bookmark/internal/bookmarks"
	"github.com/spf13/cobra"
)

func NewInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info <aliasName>",
		Short: "Display detailed information about a bookmark",
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

			existsStr := "Yes"
			if _, err := os.Stat(bookmark.Path); os.IsNotExist(err) {
				existsStr = "No"
			}

			fmt.Printf("Alias: %s\nPath: %s\nCreated: %s\nUpdated: %s\nExists: %s\n",
				bookmark.Alias, bookmark.Path,
				bookmark.CreatedAt.Format("2006-01-02 15:04"),
				bookmark.UpdatedAt.Format("2006-01-02 15:04"),
				existsStr)
			return nil
		},
	}
	return cmd
}
