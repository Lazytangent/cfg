package git

import (
	"log"
	"os"
	"os/exec"
)

func Run(connectStdin, connectStdout bool, args ...string) (string, error) {
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
