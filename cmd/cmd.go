package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/lazytangent/cfg/cmd/add"
	"github.com/lazytangent/cfg/cmd/commit"
	"github.com/lazytangent/cfg/cmd/commit/format"
	"github.com/lazytangent/cfg/cmd/config"
	"github.com/lazytangent/cfg/cmd/diff"
	"github.com/lazytangent/cfg/cmd/push"
	"github.com/lazytangent/cfg/cmd/restore"
	"github.com/lazytangent/cfg/cmd/status"
	"github.com/lazytangent/cfg/constants"
	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

var rootCmd = &cobra.Command{
	Use:              "cfg [COMMAND] -- [GIT_ARGS]",
	Short:            "A convenience wrapper for handling dotfiles with a bare git repository.",
	Run:              run,
	PersistentPreRun: preRun,
	TraverseChildren: true,
	Version:          "0.1.6",
}

var cfgFile string
var v *viper.Viper

func Execute() {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, "viper", v)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
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
	git.Run(debug, true, true, cmd, gitArgs...)
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
	v = viper.New()

	cobra.OnInitialize(initConfig(v))

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Set to print extra lines for debugging")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.config/cfg/config.toml)")

	rootCmd.AddCommand(add.Cmd, diff.Cmd, commit.Cmd, commit.CmCmd, format.Cmd, config.Cmd, push.Cmd, restore.Cmd, status.Cmd)
}

func initConfig(v *viper.Viper) func() {
	return func() {
		var configPath string

		if cfgFile != "" {
			v.SetConfigFile(cfgFile)
		} else {
			home, err := os.UserHomeDir()
			utils.LogFatalIfErr(err)

			configPath = filepath.Join(home, ".config", "cfg")

			v.AddConfigPath(configPath)
			v.SetConfigName("config")
			v.SetConfigType("toml")
		}

		if err := v.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Println("Config file not found. Creating...")

				err := os.MkdirAll(configPath, 0755)
				utils.LogFatalIfErr(err)

				err = os.WriteFile(filepath.Join(configPath, "config.toml"), []byte(constants.DefaultConfig), 0644)
				utils.LogFatalIfErr(err)
				return
			}

			log.Fatal(err)
		}
	}
}
