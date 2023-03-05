package cmd

import (
	"github.com/lazytangent/cfg/cmd/status"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the current git status of the dotfiles repo",
	Run:   status.Run,
}
