package add

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

type model struct {
	form *huh.Form
}

// Get list of changed files from `git status --porcelain=1`
func getFilesList(cmd *cobra.Command, debug bool) []string {
	var list []string

	output, err := git.Run(debug, false, false, cmd, "status", "--porcelain=1")
	utils.LogFatalIfErr(err)

	homeDir, err := os.UserHomeDir()
	utils.LogFatalIfErr(err)

	for _, line := range strings.Split(output, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		split := strings.Split(trimmed, " ")
		if debug {
			fmt.Printf("%#v\n", line)
		}

		list = append(list, filepath.Join(homeDir, split[1]))
	}

	return list
}

func newModel(opts []string, debug bool) model {
	return model{
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Key("items").
					Options(huh.NewOptions(opts...)...),
			),
		),
	}
}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	if m.form.State == huh.StateAborted || m.form.State == huh.StateCompleted {
		return m, tea.Quit
	}

	return m, cmd
}

func (m model) View() string {
	return m.form.View()
}
