package format

import (
	"fmt"
	"os"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "format TOPIC MESSAGE",
	Aliases: []string{"f"},
	Short:   "Commit with a pre-formatted message based on the positional arguments provided.",
	Run:     run,
}

func run(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	utils.LogFatalIfErr(err)

	if debug {
		fmt.Println(utils.CreateDelimiter("Commit Format command"))
	}

	if len(args) < 2 {
		cmd.Help()
		os.Exit(1)
	}

	message := fmt.Sprintf("(%s) %s", args[0], args[1])
	runArgs := []string{"commit", "--message", message}

	git.Run(debug, true, true, cmd, runArgs...)
}
