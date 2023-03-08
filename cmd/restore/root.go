package restore

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

func Run(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		log.Fatal(err)
	}

	if debug {
		fmt.Println(utils.CreateDelimiter("Restore Command"))
	}

	runArgs := []string{"restore"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	}

	output, err := git.Run(false, false, runArgs...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(output)
}
