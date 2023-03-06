package config

import (
	"fmt"

	config_ "github.com/lazytangent/cfg/config"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "config",
	Aliases: []string{"cfg", "f"},
	Short: "Configuration stuff",
	Run:   Run,
}

var (
	list bool
)

func init() {
	Cmd.PersistentFlags().BoolVarP(&list, "list", "l", false, "List values from config")
}

func Run(cmd *cobra.Command, args []string) {
	if list {
		data := config_.ReadConfigFile()
		cfg := config_.Parse(data)
		fmt.Printf("%v\n", cfg.Print())
	}
}
