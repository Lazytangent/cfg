package cmd

import (
	"github.com/spf13/cobra"
	"github.com/lazytangent/cfgo/cmd/push"
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use: "push",
	Short: "Push any new commits to the remote repository",
	Run: push.Run,
}
