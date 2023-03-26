package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/cmd/config"
	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

var rootCmd = &cobra.Command{
	Use:              "cfg [COMMAND] -- [GIT_ARGS]",
	Short:            "A convenience wrapper for handling dotfiles with a bare git repository.",
	Run:              run,
	PersistentPreRun: preRun,
	TraverseChildren: true,
	Version:          "0.1.2",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	utils.LogFatalIfErr(err)

	if len(args) == 0 {
		if debug {
			fmt.Printf("Args: %v\n", args)
		}
		cmd.Help()
		os.Exit(0)
	}

	gitArgs := utils.GetGitArgs(cmd, args)
	git.Run(debug, true, true, gitArgs...)
}

func preRun(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	utils.LogFatalIfErr(err)

	if debug {
		delim := utils.CreateDelimiter("Root Cmd")
		fmt.Println(delim)
		idx := cmd.ArgsLenAtDash()
		fmt.Printf("Index of Git Args: %d\n", idx)

		if idx >= 0 {
			fmt.Println(args[idx:])
		}
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Set to print extra lines for debugging")

	rootCmd.AddCommand(addCmd, commitCmd, config.Cmd, pushCmd, restoreCmd, statusCmd)
}
