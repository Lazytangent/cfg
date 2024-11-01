package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/lazytangent/cfg/utils"
)

type Config struct {
	GitDir       string  `toml:"git_dir" mapstructure:"git_dir"`
	WorkTree     string  `toml:"work_tree" mapstructure:"work_tree"`
	LocalRepoDir *string `toml:"local_repo_dir" mapstructure:"local_work_dir"`
}

func (c Config) String() string {
	data, err := json.MarshalIndent(c, "", "  ")
	utils.LogFatalIfErr(err)

	return string(data)
}

func ParseTildeInPath(path string) string {
	dir, err := os.UserHomeDir()
	utils.LogFatalIfErr(err)

	return strings.ReplaceAll(path, "~", dir)
}

func GetConfig(cmd *cobra.Command) (*Config, error) {
	vpr, err := utils.GetViper(cmd)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = vpr.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
