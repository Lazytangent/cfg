package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use: "push",
	Short: "Push any new commits to the remote repository",
	Run: push,
}

func push(cmd *cobra.Command, args []string) {
	fmt.Println("Push command")
}
