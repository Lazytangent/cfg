package git

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/lazytangent/cfgo/config"
)

func Run(connectStdin, connectStdout bool, args ...string) (string, error) {
	args = addDefaultArgs(args)

	cmd := exec.Command("git", args...)
	if connectStdin {
		cmd.Stdin = os.Stdin
	}
	if connectStdout {
		cmd.Stdout = os.Stdout
	}
	data, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}

	return string(data), err
}

func addDefaultArgs(args []string) []string {
	cfg := config.Parse(config.ListConfig())
	gitDir := config.ParseTildeInPath(cfg.GitDir)
	workTree := config.ParseTildeInPath(cfg.WorkTree)

	gitDirFlag := fmt.Sprintf("--git-dir=%s", gitDir)
	workTreeFlag := fmt.Sprintf("--work-tree=%s", workTree)

	newArgs := []string{gitDirFlag, workTreeFlag}
	newArgs = append(newArgs, args...)

	return newArgs
}
