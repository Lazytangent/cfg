package commit

import (
	"github.com/lazytangent/cfgo/git"
	"github.com/spf13/cobra"
	"fmt"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Commit command")
	runArgs := []string {"commit"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	}

	git.Run(runArgs...)
}
