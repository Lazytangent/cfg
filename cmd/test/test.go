package test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lazytangent/cfg/git"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "test",
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	repository, err := git.Repository()
	if err != nil {
		log.Fatalf("git.Repository: %s", err.Error())
	}

	cfg, err := repository.Config()
	if err != nil {
		log.Fatalf("repository.Config: %s", err.Error())
	}

	jsonified, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Fatalf("json.MarshalIndent: %s", err.Error())
	}
	fmt.Println(string(jsonified))
}
