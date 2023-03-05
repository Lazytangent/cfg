package config

import (
	"fmt"

	config_ "github.com/lazytangent/cfgo/config"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "config",
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
		data := config_.ListConfig()
		cfg := config_.Parse(data)
		fmt.Printf("%v\n", cfg.Print())
	}
}
