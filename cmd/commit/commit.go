package commit

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

func Run(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	utils.LogFatalIfErr(err)

	if debug {
		fmt.Println("Commit command")
	}
	runArgs := []string{"commit"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	}

	git.Run(debug, true, true, runArgs...)
}
