package status

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

const notStagedMsg = `Changes not staged for commit:`
const stagedMsg = `Changes to be committed:`

func Run(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	utils.LogFatalIfErr(err)
	if debug {
		fmt.Println(utils.CreateDelimiter("Status Command"))
	}

	output, err := git.Run(false, false, "status")
	if debug {
		fmt.Println("Normal output:")
		fmt.Print(output)
		fmt.Println(utils.CreateDelimiter(""))
	}

	if err != nil {
		return
	}

	notModifiedStatus := notModified(output)
	stagedStatus := staged(output)

	if !notModifiedStatus && !stagedStatus {
		c := color.New(color.FgHiWhite).Add(color.Bold)
		c.Println("No files changed")
	}
}

func staged(output string) bool {
	outputSplit := strings.Split(output, stagedMsg)

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

		return true
	}

	return false
}

func notModified(output string) bool {
	outputSplit := strings.Split(output, notStagedMsg)

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

		return true
	}

	return false
}

func substituteTilde(path string) string {
	cwd, err := os.Getwd()
	utils.LogFatalIfErr(err)

	homeDir, err := os.UserHomeDir()
	utils.LogFatalIfErr(err)

	cwdTrimmed := strings.Split(cwd, homeDir)
	cwdSplitDirsCount := strings.Count(cwdTrimmed[1], "/")

	newPath := path

	for i := 0; i < cwdSplitDirsCount; i++ {
		newPath = strings.TrimPrefix(newPath, "../")
	}

	newPath = fmt.Sprintf("~/%s", newPath)

	return newPath
}
