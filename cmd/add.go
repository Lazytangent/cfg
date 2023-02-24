package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lazytangent/cfgo/git"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add files to the staging area",
	Run:   add,
}

func add(cmd *cobra.Command, args []string) {
	fmt.Println("Add command")

	runArgs := []string{"add"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	}

	git.Run(runArgs...)
}
