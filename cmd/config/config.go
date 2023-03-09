package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	config_ "github.com/lazytangent/cfg/config"
	"github.com/lazytangent/cfg/utils"
)

var Cmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"cfg", "f"},
	Short:   "An alias for `git config` specific to the dotfiles local repository",
	Run:     Run,
}

var (
	list bool
)

func init() {
	Cmd.PersistentFlags().BoolVarP(&list, "list", "l", false, "List values from config")
}

func Run(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	utils.LogPrintlnIfErr(err)

	if debug {
		fmt.Println(utils.CreateDelimiter("Config Command"))
		return
	}

	if list {
		data := config_.ReadConfigFile()
		cfg := config_.Parse(data)
		fmt.Printf("%v\n", cfg.Print())
		return
	}

	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}
}
