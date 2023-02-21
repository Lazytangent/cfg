package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add files to the staging area",
	Run: add,
}

func add(cmd *cobra.Command, args []string) {
	fmt.Println("Add command")
}
