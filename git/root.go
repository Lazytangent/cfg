package git

import (
	"os/exec"
	"log"
	"fmt"
)

func Run() {
	out, err := exec.Command("git").Output()
	if err != nil {
		log.Println(err)
	}

	fmt.Print(string(out))
}
