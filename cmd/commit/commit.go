package commit

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

func init() {
	Cmd.PersistentFlags().StringP("message", "m", "", "Commit message")
}

var Cmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"c"},
	Short:   "Commit files in the staging area",
	Run:     run,
}

func run(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	utils.LogFatalIfErr(err)

	msg, err := cmd.Flags().GetString("message")
	utils.LogFatalIfErr(err)

	if debug {
		fmt.Println("Commit command")
	}
	runArgs := []string{"commit"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	}

	if msg != "" {
		runArgs = append(runArgs, "-m", msg)
	}

	git.Run(debug, true, true, runArgs...)
}
