package push

import (
	"fmt"

	"github.com/lazytangent/cfg/git"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Push command")
	runArgs := []string{"push"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	}

	git.Run(true, true, runArgs...)
}
