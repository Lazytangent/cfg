package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

var rootCmd = &cobra.Command{
	Use:              "cfg [COMMAND] -- [GIT_ARGS]",
	Short:            "A convenience wrapper for handling dotfiles with a bare git repository.",
	Run:              run,
	PersistentPreRun: preRun,
	TraverseChildren: true,
	Version:          "0.1.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	gitArgs := utils.GetGitArgs(cmd, args)
	git.Run(true, true, gitArgs...)
}

func preRun(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	utils.LogFatalIfErr(err)

	if debug {
		delim := utils.CreateDelimiter("Root Cmd")
		fmt.Println(delim)
		idx := cmd.ArgsLenAtDash()
		fmt.Printf("Index of Git Args: %d\n", idx)

		fmt.Println(args)
		if idx >= 0 {
			fmt.Println(args[idx:])
		}
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Set to print extra lines for debugging")
}