/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "dev"

func NewVersionCmd() *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of go-bookmark",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("go-bookmark v%s\n", version)
		},
	}
	return versionCmd
}
