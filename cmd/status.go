package cmd

import (
	"github.com/lazytangent/cfg/cmd/status"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"s", "st"},
	Short:   "Print the current git status of the dotfiles repo",
	Run:     status.Run,
}
