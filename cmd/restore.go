package cmd

import (
	"github.com/spf13/cobra"
	"github.com/lazytangent/cfg/cmd/restore"
)

var restoreCmd = &cobra.Command{
	Use: "restore",
	Aliases: []string{"res"},
	Short: "Restores files to previous state",
	Run: restore.Run,
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
