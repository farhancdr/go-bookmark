package cmd

import (
	"fmt"

	"github.com/farhancdr/go-bookmark/internal/bookmarks"
	"github.com/spf13/cobra"
)

func NewClearCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "Delete all bookmarks",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Print("Are you sure you want to delete ALL bookmarks? (y/N): ")
			if !confirm() {
				fmt.Println("Operation cancelled.")
				return nil
			}

			err := bookmarks.ClearBookmarks()
			if err != nil {
				return fmt.Errorf("failed to clear bookmarks: %v", err)
			}

			fmt.Println("All bookmarks deleted.")
			return nil
		},
	}
	return cmd
}
