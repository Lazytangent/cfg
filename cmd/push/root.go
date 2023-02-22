package push

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/lazytangent/cfgo/git"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Push command")
	runArgs := []string {"push"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	}

	git.Run(runArgs...)
}
