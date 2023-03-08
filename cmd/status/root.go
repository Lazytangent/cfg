package status

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/fatih/color"

	"github.com/lazytangent/cfg/git"
)

const notStaged = `Changes not staged for commit:`
const staged = `Changes to be committed:`

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

	handleNotModified(output)
	handleStaged(output)
}

func handleStaged(output string) {
	outputSplit := strings.Split(output, staged)

	if len(outputSplit) > 1 {
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
		c.Println("Staged:")
		for _, path := range modifiedSection {
			newPath := substituteTilde(path)
			newLine := fmt.Sprintf("\t%s", newPath)
			color.Green(newLine)
		}
	}
}

func handleNotModified(output string) {
	outputSplit := strings.Split(output, notStaged)

	if len(outputSplit) > 1 {
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
		for _, path := range modifiedSection {
			newPath := substituteTilde(path)
			newLine := fmt.Sprintf("\t%s", newPath)
			color.Red(newLine)
		}
	}
}

func substituteTilde(path string) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	cwdTrimmed := strings.Split(cwd, homeDir)
	cwdSplitDirsCount := strings.Count(cwdTrimmed[1], "/")

	newPath := path

	for i := 0; i < cwdSplitDirsCount; i++ {
		newPath = strings.TrimPrefix(newPath, "../")
	}

	newPath = fmt.Sprintf("~/%s", newPath)

	return newPath
}
