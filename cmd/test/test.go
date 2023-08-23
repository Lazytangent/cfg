package test

import (
	"log"

	"github.com/lazytangent/cfg/git"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "test",
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	worktree, err := git.Worktree()
	if err != nil {
		log.Fatalf("git.Worktree: %s", err.Error())
	}

	subs, err := worktree.Submodules()
	if err != nil {
		log.Fatalf("worktree.Submodules: %s", err.Error())
	}

	for _, submodule := range subs {
		status, err := submodule.Status()
		if err != nil {
			log.Fatalf("submodule.Status %s", err.Error())
		}
		log.Println(status.String())
	}
}
