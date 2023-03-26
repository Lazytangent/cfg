package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add files to the staging area",
	Run:     add,
}

func add(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	utils.LogFatalIfErr(err)

	if debug {
		fmt.Println("Add command")
	}

	runArgs := []string{"add", "-f"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	}

	git.Run(debug, true, true, runArgs...)
}
