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

	output, err := git.Run(debug, false, false, "status")
	if debug {
		fmt.Println("Normal output:")
		fmt.Print(output)
		fmt.Println(utils.CreateDelimiter(""))
	}

	if err != nil {
		return
	}

	stagedStatus := staged(output)
	notModifiedStatus := notModified(output)

	if !notModifiedStatus && !stagedStatus {
		c := color.New(color.FgHiWhite).Add(color.Bold)
		c.Println("No files changed")
	}
}

func staged(output string) bool {
	notStagedSplit := strings.Split(output, notStagedMsg)
	outputSplit := strings.Split(notStagedSplit[0], stagedMsg)

	if len(outputSplit) > 1 {
		notStagedSection := outputSplit[1]
		modifiedRe := regexp.MustCompile(`^\s+modified:\s+`)
		newFileRe := regexp.MustCompile(`^\s+new file:\s+`)
		deletedRe := regexp.MustCompile(`^\s+deleted:\s+`)

		modifiedSection := []string{}
		newFileSection := []string{}
		deletedSection := []string{}

		for _, line := range strings.Split(notStagedSection, "\n") {
			if modifiedRe.MatchString(line) {
				split := modifiedRe.Split(line, -1)
				modifiedSection = append(modifiedSection, split[1])
			}

			if newFileRe.MatchString(line) {
				split := newFileRe.Split(line, -1)
				newFileSection = append(newFileSection, split[1])
			}

			if deletedRe.MatchString(line) {
				split := deletedRe.Split(line, -1)
				deletedSection = append(deletedSection, split[1])
			}
		}

		c := color.New(color.FgHiWhite).Add(color.Bold)
		c.Println("Staged:")

		if len(modifiedSection) > 0 {
			c.Println("    Modified:")
			for _, path := range modifiedSection {
				var newPath string
				if strings.HasPrefix(path, "../../") {
					newPath = substituteTilde(path)
				} else {
					newPath = path
				}

				newLine := fmt.Sprintf("\t%s", newPath)
				color.Green(newLine)
			}
		}

		if len(newFileSection) > 0 {
			if len(modifiedSection) > 0 {
				fmt.Println()
			}

			c.Println("    New File(s):")

			for _, path := range newFileSection {
				var newPath string
				if strings.HasPrefix(path, "../../") {
					newPath = substituteTilde(path)
				} else {
					newPath = path
				}

				newLine := fmt.Sprintf("\t%s", newPath)
				color.Green(newLine)
			}
		}

		if len(deletedSection) > 0 {
			if len(modifiedSection) > 0 || len(newFileSection) > 0 {
				fmt.Println()
			}

			c.Println("    Deleted:")

			for _, path := range deletedSection {
				var newPath string
				if strings.HasPrefix(path, "../../") {
					newPath = substituteTilde(path)
				} else {
					newPath = path
				}

				newLine := fmt.Sprintf("\t%s", newPath)
				color.Green(newLine)
			}
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
		deletedRe := regexp.MustCompile(`^\s+deleted:\s+`)

		modifiedSection := []string{}
		deletedSection := []string{}

		for _, line := range strings.Split(notStagedSection, "\n") {
			if modifiedRe.MatchString(line) {
				split := modifiedRe.Split(line, -1)
				modifiedSection = append(modifiedSection, split[1])
			}

			if deletedRe.MatchString(line) {
				split := deletedRe.Split(line, -1)
				deletedSection = append(deletedSection, split[1])
			}
		}

		c := color.New(color.FgHiWhite).Add(color.Bold)
		c.Println("Unstaged:")
		for _, path := range modifiedSection {
			var newPath string
			if strings.HasPrefix(path, "../../") {
				newPath = substituteTilde(path)
			} else {
				newPath = path
			}

			newLine := fmt.Sprintf("\t%s", newPath)
			color.Red(newLine)
		}

		if len(deletedSection) > 0 {
			if len(modifiedSection) > 0 {
				fmt.Println()
			}

			c.Println("    Deleted:")

			for _, path := range deletedSection {
				var newPath string
				if strings.HasPrefix(path, "../../") {
					newPath = substituteTilde(path)
				} else {
					newPath = path
				}

				newLine := fmt.Sprintf("\t%s", newPath)
				color.Red(newLine)
			}
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
