package git

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/lazytangent/cfg/config"
	"github.com/lazytangent/cfg/utils"
)

func Run(debug, connectStdin, connectStdout bool, args ...string) (string, error) {
	args = addDefaultArgs(args)

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
		cmd.Run()
		return "", nil
	} else {
		data, err := cmd.Output()
		utils.LogPrintlnIfErr(err)

		return string(data), err
	}
}

func addDefaultArgs(args []string) []string {
	cfg := config.Parse(config.ReadConfigFile())
	gitDir := config.ParseTildeInPath(cfg.GitDir)
	workTree := config.ParseTildeInPath(cfg.WorkTree)

	gitDirFlag := fmt.Sprintf("--git-dir=%s", gitDir)
	workTreeFlag := fmt.Sprintf("--work-tree=%s", workTree)

	newArgs := []string{gitDirFlag, workTreeFlag}
	newArgs = append(newArgs, args...)

	return newArgs
}
