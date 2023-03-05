package cmd

import (
	"github.com/lazytangent/cfg/cmd/config"
)

func init() {
	rootCmd.AddCommand(config.Cmd)
}
