package add

import (
	"errors"
	"fmt"

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
		filesList := getFilesList(debug)

		if len(filesList) == 0 {
			if debug {
				fmt.Println("len(filesList) == 0", filesList)
			}

			fmt.Println("No files in filesList, exiting")
			return
		} else if len(filesList) == 1 {
			if debug {
				fmt.Println("len(filesList) == 1", filesList)
			}

			fmt.Printf("Only one file in filesList, adding it (%s) to the staging area...\n", filesList[0])
			runArgs = append(runArgs, filesList[0])
		} else {
			if debug {
				fmt.Println("len(filesList) > 1", filesList)
			}

			items, err := getItemsFromTeaProgram(filesList, debug)
			utils.LogFatalIfErr(err)

			runArgs = append(runArgs, items...)
		}
	}

	git.Run(debug, true, true, runArgs...)
}

func getItemsFromTeaProgram(filesList []string, debug bool) ([]string, error) {
	p, err := tea.NewProgram(newModel(filesList, debug)).Run()
	if err != nil {
		return nil, err
	}

	m, ok := p.(model)
	if !ok {
		return nil, errors.New("[add] could not cast tea.Model to add.model")
	}
	data := m.form.Get("items")

	items, ok := data.([]string)
	if !ok {
		return nil, errors.New("[add] could not cast any to []string")
	}

	return items, nil
}
