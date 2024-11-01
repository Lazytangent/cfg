package diff

import (
	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

var Cmd = &cobra.Command{
	Use:     "diff",
	Aliases: []string{"d"},
	Short:   "Show git diff",
	Run:     run,
}

func run(cmd *cobra.Command, args []string) {
	flags := cmd.Flags()
	debug, err := flags.GetBool("debug")
	utils.LogPrintlnIfErr(err)

	runArgs := []string{"diff"}
	cached, err := flags.GetBool("cached")
	utils.LogFatalIfErr(err)

	if cached {
		runArgs = append(runArgs, "--cached")
	}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	}

	git.Run(debug, true, true, cmd, runArgs...)
}

func init() {
	Cmd.PersistentFlags().Bool("cached", false, "Show diff of cached files")
}
