package status

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
)

func Run(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		fmt.Println("Status command")
	}

	output, err := git.Run(false, false, "status")
	if err != nil {
		return
	}

	fmt.Print(output)
}
