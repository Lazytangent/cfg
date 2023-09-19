package test

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	git "github.com/libgit2/git2go/v34"

	"github.com/lazytangent/cfg/config"
)

var Cmd = &cobra.Command{
	Use: "test",
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	cfg := config.Parse(config.ReadConfigFile())
	repo, err := git.OpenRepository(config.ParseTildeInPath(cfg.GitDir))
	if err != nil {
		log.Fatalf("git.OpenRepository: %s", err)
	}

	statusOpt := &git.StatusOptions{}
	statusList, err := repo.StatusList(statusOpt)
	if err != nil {
		log.Fatalf("repo.StatusList: %s", err)
	}

	entryCount, err := statusList.EntryCount()
	if err != nil {
		log.Fatalf("statusList.EntryCount: %s", err)
	}

	fmt.Printf("  %d entries in status list\n", entryCount)
}
