package cmd

import (
	"github.com/lazytangent/cfg/cmd/push"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push any new commits to the remote repository",
	Run:   push.Run,
}
