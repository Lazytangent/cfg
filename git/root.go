package git

import (
	"os/exec"
	"log"
	"fmt"
	"strings"
)

func Run(args ...string) {
	gitArgs := strings.Join(args, " ")
	out, err := exec.Command("git", gitArgs).Output()
	if err != nil {
		log.Println(err)
	}

	fmt.Print(string(out))
}
