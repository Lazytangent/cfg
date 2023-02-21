package status

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/lazytangent/cfgo/git"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Status command")
	git.Run()
}
