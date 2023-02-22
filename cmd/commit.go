package cmd

import (
	"github.com/spf13/cobra"
	"github.com/lazytangent/cfgo/cmd/commit"
)

func init() {
	rootCmd.AddCommand(commitCmd)
}

var commitCmd = &cobra.Command{
	Use: "commit",
	Short: "Commit files in the staging area",
	Run: commit.Run,
}
