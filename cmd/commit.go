package cmd

import (
	"github.com/lazytangent/cfgo/cmd/commit"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commitCmd)
}

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit files in the staging area",
	Run:   commit.Run,
}
