package cmd

import (
	"github.com/lazytangent/cfg/cmd/restore"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:     "restore",
	Aliases: []string{"res"},
	Short:   "Restores files to previous state",
	Run:     restore.Run,
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
