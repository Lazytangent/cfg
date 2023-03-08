package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Aliases: []string{"a"},
	Short: "Add files to the staging area",
	Run:   add,
}

func add(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		log.Fatal(err)
	}

	if debug {
		fmt.Println("Add command")
	}

	runArgs := []string{"add", "-f"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	}

	git.Run(true, true, runArgs...)
}
