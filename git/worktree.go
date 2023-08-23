package git

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/lazytangent/cfg/config"
)

// TEST: Does this work the same at the root of a worktree and inside the
// worktree?
func Worktree() (*git.Worktree, error) {
	cfg := config.Parse(config.ReadConfigFile())
	gitDir := config.ParseTildeInPath(cfg.GitDir)

	repo, err := git.PlainOpen(gitDir)
	if err != nil {
		log.Fatalf("git.PlainOpen: %s", err.Error())
	}

	return repo.Worktree()
}
