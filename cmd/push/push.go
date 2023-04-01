package push

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

var Cmd = &cobra.Command{
	Use:     "push",
	Aliases: []string{"p", "pu"},
	Short:   "Push any new commits to the remote repository",
	Run:     run,
}

func run(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	utils.LogFatalIfErr(err)

	if debug {
		fmt.Println("Push command")
	}
	runArgs := []string{"push"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	}

	git.Run(debug, true, true, runArgs...)
}
