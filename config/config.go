package config

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/BurntSushi/toml"

	"github.com/lazytangent/cfg/utils"
)

type Config struct {
	GitDir   string `toml:"git_dir"`
	WorkTree string `toml:"work_tree"`
	LocalRepoDir *string `toml:"local_repo_dir"`
}

func Parse(data string) Config {
	var conf Config
	_, err := toml.Decode(data, &conf)
	utils.LogFatalIfErr(err)

	return conf
}

func (c Config) Print() string {
	data, err := json.MarshalIndent(c, "", "  ")
	utils.LogFatalIfErr(err)

	return string(data)
}

const configFile = "~/.config/cfg/config.toml"

func ParseTildeInPath(path string) string {
	dir, err := os.UserHomeDir()
	utils.LogFatalIfErr(err)

	return strings.ReplaceAll(path, "~", dir)
}

const defaultConfigFile = `
git_dir = "~/.cfg/"
work_tree = "~/"
`

func ReadConfigFile() string {
	configPath := ParseTildeInPath(configFile)
	dat, err := os.ReadFile(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return defaultConfigFile
	}

	utils.LogFatalIfErr(err)
	return string(dat)
}
