package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "bm",
		Short: "bm is a CLI tool for bookmarking and navigating directories",
		Long:  `bm (bookmark) is a cross-platform CLI tool to save, manage, and navigate to bookmarked directories efficiently.`,
	}

	rootCmd.AddCommand(NewSaveCmd())
	rootCmd.AddCommand(NewGotoCmd())
	rootCmd.AddCommand(NewListCmd())
	rootCmd.AddCommand(NewDeleteCmd())
	rootCmd.AddCommand(NewClearCmd())
	rootCmd.AddCommand(NewUpdateCmd())
	rootCmd.AddCommand(NewRenameCmd())
	rootCmd.AddCommand(NewInfoCmd())

	return rootCmd
}

func Execute() error {
	return NewRootCmd().Execute()
}
