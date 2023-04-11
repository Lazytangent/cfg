package commit

import (
	"fmt"
	"log"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
	"github.com/spf13/cobra"
)

var CmCmd = &cobra.Command{
	Use: "cm [message]",
	Short: "commit -m alias",
	Run: cmRun,
}

func cmRun(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	utils.LogFatalIfErr(err)

	if debug {
		fmt.Println(utils.CreateDelimiter("Commit -m Command"))
	}

	if len(args) < 1 {
		log.Fatal("Message argument required")
	}

	msg := args[0]

	commitArgs := []string{"commit", "-m", msg}

	git.Run(debug, true, true, commitArgs...)
}
