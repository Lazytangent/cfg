package status

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/fatih/color"

	"github.com/lazytangent/cfg/git"
)

const notStaged = `Changes not staged for commit:`

func Run(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		fmt.Println("Status command")
	}

	output, err := git.Run(false, false, "status")
	if err != nil {
		return
	}

	outputSplit := strings.Split(output, notStaged)
	notStagedSection := outputSplit[1]
	modifiedRe := regexp.MustCompile(`^\s+modified:\s+`)

	modifiedSection := []string{}

	for _, line := range strings.Split(notStagedSection, "\n") {
		if modifiedRe.MatchString(line) {
			split := modifiedRe.Split(line, -1)
			modifiedSection = append(modifiedSection, split[1])
		}
	}

	c := color.New(color.FgHiWhite).Add(color.Bold)
	c.Println("Modified Files:")
	for _, line := range modifiedSection {
		newLine := fmt.Sprintf("\t%s", line)
		color.Red(newLine)
	}
}
