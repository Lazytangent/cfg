package status

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Status command")
	output, err := git.Run(false, false, "status")
	if err != nil {
		return
	}

	fmt.Println(output)
}
