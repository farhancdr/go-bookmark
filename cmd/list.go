package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/farhancdr/go-bookmark/internal/bookmarks"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all bookmarks",
		RunE: func(cmd *cobra.Command, args []string) error {
			store, err := bookmarks.LoadBookmarks()
			if err != nil {
				return fmt.Errorf("failed to load bookmarks: %v", err)
			}

			if len(store.Bookmarks) == 0 {
				fmt.Println("✨ No bookmarks found. Start by adding some! ✨")
				return nil
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			for _, b := range store.Bookmarks {
				fmt.Fprintf(w, "✅ %s\t%s\t%s\t\n", b.Alias, b.Path, b.UpdatedAt.Format("2006-01-02 15:04"))
			}
			w.Flush()
			fmt.Println("\n💡 * Marked bookmarks point to directories that no longer exist. You might want to remove them!")
			return nil
		},
	}
	return cmd
}
