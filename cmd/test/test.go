package test

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use: "test",
	Run: run,
}

func run(cmd *cobra.Command, args []string) {

}
