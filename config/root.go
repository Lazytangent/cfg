package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	GitDir   string `toml:"git_dir"`
	WorkTree string `toml:"work_tree"`
}

func Parse(data string) Config {
	var conf Config
	if _, err := toml.Decode(data, &conf); err != nil {
		log.Fatal(err)
	}

	return conf
}

func (c Config) Print() string {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

const configFile = "~/.config/cfg/config.toml"

func ParseTildeInPath(path string) string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return strings.Replace(path, "~", dir, -1)
}

const defaultConfigFile = `
git_dir = "~/.cfg/"
work_tree = "~/"
`

func ReadConfigFile() string {
	configPath := ParseTildeInPath(configFile)
	dat, err := os.ReadFile(configPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return defaultConfigFile
		}
		log.Fatal(err)
	}
	return string(dat)
}
