package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

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

const configFile = "~/.config/cfgo/config.toml"

// TODO: Refactor to use strings.Replacer to explicitly replace the starting
// '~/'
func ParseTildeInPath(path string) string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(dir, path[2:])
}

func ReadConfigFile() string {
	configPath := ParseTildeInPath(configFile)
	dat, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	return string(dat)
}
