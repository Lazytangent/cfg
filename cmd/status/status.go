package status

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/git"
	"github.com/lazytangent/cfg/utils"
)

var Cmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"s", "st"},
	Short:   "Print the current git status of the dotfiles repo",
	Run:     run,
}

func init() {
	Cmd.PersistentFlags().BoolP("verbose", "v", false, "Show full output from underlying git command")
}

func run(cmd *cobra.Command, args []string) {
	flags := cmd.Flags()
	debug, err := flags.GetBool("debug")
	utils.LogFatalIfErr(err)
	verbose, err := flags.GetBool("verbose")
	utils.LogFatalIfErr(err)
	if debug {
		fmt.Println(utils.CreateDelimiter("Status Command"))
	}

	if verbose {
		git.Run(debug, true, true, cmd, "status")
		return
	}

	output, err := git.Run(debug, false, false, cmd, "status")
	if debug {
		fmt.Println("Normal output:")
		fmt.Print(output)
		fmt.Println(utils.CreateDelimiter(""))
	}

	if err != nil {
		return
	}

	ahead(output)
	stagedStatus := staged(output)
	notModifiedStatus := notModified(output)

	if !notModifiedStatus && !stagedStatus {
		c := color.New(color.FgHiWhite).Add(color.Bold)
		c.Println("No files changed")
	}
}

func ahead(output string) {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Your branch is ahead of") {
			var remote string
			var numOfCommits int
			c := color.New(color.FgHiWhite).Add(color.Bold)

			_, err := fmt.Sscanf(line, "Your branch is ahead of %s by %d commits.", &remote, &numOfCommits)
			if err != nil {
				_, err := fmt.Sscanf(line, "Your branch is ahead of %s by 1 commit.", &remote)
				if err != nil {
					fmt.Println(line)
					log.Println(err)
					return
				}

				c.Printf("Ahead of %s by 1 commit.", remote)
				fmt.Println()
				fmt.Println()
				return
			}

			c.Printf("Ahead of %s by %d commits.", remote, numOfCommits)
			fmt.Println()
			fmt.Println()
		}
	}
}

const notStagedMsg = `Changes not staged for commit:`
const stagedMsg = `Changes to be committed:`

func staged(output string) bool {
	notStagedSplit := strings.Split(output, notStagedMsg)
	outputSplit := strings.Split(notStagedSplit[0], stagedMsg)

	if len(outputSplit) > 1 {
		notStagedSection := outputSplit[1]
		modifiedRe := regexp.MustCompile(`^\s+modified:\s+`)
		newFileRe := regexp.MustCompile(`^\s+new file:\s+`)
		deletedRe := regexp.MustCompile(`^\s+deleted:\s+`)
		renamedRe := regexp.MustCompile(`^\s+renamed:\s+`)

		modifiedSection := []string{}
		newFileSection := []string{}
		deletedSection := []string{}
		renamedSection := []string{}

		// TODO: Refactor this to remove all the repetitiveness
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

			if renamedRe.MatchString(line) {
				split := renamedRe.Split(line, -1)
				renamedSection = append(renamedSection, split[1])
			}
		}

		totalLineCount := len(modifiedSection) + len(newFileSection) + len(deletedSection) + len(renamedSection)

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

		if renamedSectionLength := len(renamedSection); renamedSectionLength > 0 {
			// TODO: Update other similar conditionals based on this one
			if totalLineCount-renamedSectionLength > 0 {
				fmt.Println()
			}

			c.Println("    Renamed:")
			// TODO: Refactor these blocks into a function
			for _, path := range renamedSection {
				var newPath string
				if strings.HasPrefix(path, "../../") {
					newPath = substituteTilde(path)
				} else {
					newPath = path
				}

				split := strings.Split(newPath, " -> ")
				secondPath := split[1]
				if strings.HasPrefix(path, "../../") {
					secondPath = substituteTilde(secondPath)
				}
				newPath = fmt.Sprintf("%s -> %s", split[0], secondPath)

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

	newPath := filepath.Join(cwd, path)
	newPath = filepath.Clean(newPath)

	homeDir, err := os.UserHomeDir()
	utils.LogFatalIfErr(err)

	newPath = strings.ReplaceAll(newPath, homeDir, "~")

	return newPath
}
