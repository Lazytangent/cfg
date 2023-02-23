package cmd

import (
	"github.com/lazytangent/cfgo/cmd/config"
)

func init() {
	rootCmd.AddCommand(config.Cmd)
}
