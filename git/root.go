package git

import (
	"log"
	"os"
	"os/exec"
)

func Run(args ...string) {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}
