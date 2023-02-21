package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/lazytangent/cfgo/cmd/status"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use: "status",
	Short: "Print the current git status of the dotfiles repo",
	Run: status.Run,
}
