package git

import (
	"log"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/lazytangent/cfg/config"
)

// TEST: Does this work the same at the root of a worktree and inside the
// worktree?
func Worktree() (*git.Worktree, error) {
	cfg := config.Parse(config.ReadConfigFile())
	gitDir := config.ParseTildeInPath(cfg.GitDir)
	workTree := config.ParseTildeInPath(cfg.WorkTree)

	dotGitFs := osfs.New(gitDir)
	dotGitStorer := filesystem.NewStorage(dotGitFs, cache.NewObjectLRUDefault())
	tree := osfs.New(workTree)

	repo, err := git.Open(dotGitStorer, tree)
	if err != nil {
		log.Fatalf("git.PlainOpen: %s", err.Error())
	}

	return repo.Worktree()
}
