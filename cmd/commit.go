package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	rootCmd.AddCommand(commitCmd)
}

var commitCmd = &cobra.Command{
	Use: "commit",
	Short: "Commit files in the staging area",
	Run: commit,
}

func commit(cmd *cobra.Command, args []string) {
	fmt.Println("Commit command")
}
