package cmd

import (
	"github.com/lazytangent/cfg/cmd/commit"
	"github.com/spf13/cobra"
)

func init() {
	commitCmd.PersistentFlags().StringP("message", "m", "", "Commit message")

	rootCmd.AddCommand(commitCmd)
}

var commitCmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"c"},
	Short:   "Commit files in the staging area",
	Run:     commit.Run,
}
