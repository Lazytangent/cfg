package add

import (
	"fmt"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
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

	runArgs := []string{"add"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	} else {
		list := getFilesList(debug)
		fmt.Printf("%#v\n", list)
		return
	}

	git.Run(debug, true, true, runArgs...)
}
