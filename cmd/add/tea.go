package add

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

type model struct {
	form *huh.Form
}

// Get list of changed files from `git status --porcelain=1`
func getFilesList(debug bool) []string {
	var list []string

	output, err := git.Run(debug, false, false, "status", "--porcelain=1")
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

		list = append(list, split[1])
	}

	return list
}
