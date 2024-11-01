package git

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/lazytangent/cfg/config"
	"github.com/lazytangent/cfg/utils"
	"github.com/spf13/cobra"
)

func Run(debug, connectStdin, connectStdout bool, cobraCmd *cobra.Command, args ...string) (string, error) {
	args = addDefaultArgs(cobraCmd, args)

	if debug {
		fmt.Println("Total args passed to git:")
		fmt.Println(args)
	}

	cmd := exec.Command("git", args...)
	if connectStdin {
		cmd.Stdin = os.Stdin
	}
	if connectStdout {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return "", nil
	} else {
		data, err := cmd.Output()
		utils.LogPrintlnIfErr(err)

		return string(data), err
	}
}

func addDefaultArgs(cmd *cobra.Command, args []string) []string {
	cfg, err := config.GetConfig(cmd)
	utils.LogFatalIfErr(err)

	gitDir := config.ParseTildeInPath(cfg.GitDir)
	workTree := config.ParseTildeInPath(cfg.WorkTree)

	gitDirFlag := fmt.Sprintf("--git-dir=%s", gitDir)
	workTreeFlag := fmt.Sprintf("--work-tree=%s", workTree)

	newArgs := []string{gitDirFlag, workTreeFlag}
	newArgs = append(newArgs, args...)

	return newArgs
}
