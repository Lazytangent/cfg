package add

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add files to the staging area",
	Run:     add,
}

func add(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	utils.LogFatalIfErr(err)

	if debug {
		fmt.Println("Add command")
	}

	runArgs := []string{"add"}

	if len(args) > 0 {
		runArgs = append(runArgs, args...)
	} else {
		p, err := tea.NewProgram(newModel(debug)).Run()
		utils.LogFatalIfErr(err)

		m, ok := p.(model)
		if !ok {
			log.Fatalln("[add] could not cast tea.Model to add.model")
		}
		data := m.form.Get("items")

		items, ok := data.([]string)
		if !ok {
			log.Fatalln("[add] could not cast any to []string")
		}

		runArgs = append(runArgs, items...)
	}

	git.Run(debug, true, true, runArgs...)
}
