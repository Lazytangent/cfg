package commit

import (
	"fmt"

	"github.com/lazytangent/cfgo/git"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Commit command")
	runArgs := []string{"commit"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	}

	git.Run(true, true, runArgs...)
}
