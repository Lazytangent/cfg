package restore

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

var Cmd = &cobra.Command{
	Use:     "restore",
	Aliases: []string{"res"},
	Short:   "Restores files to previous state",
	Run:     run,
}

func run(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	utils.LogFatalIfErr(err)

	if debug {
		fmt.Println(utils.CreateDelimiter("Restore Command"))
	}

	runArgs := []string{"restore"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	}

	output, err := git.Run(debug, false, false, runArgs...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(output)
}
