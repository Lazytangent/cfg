package config

import (
	"encoding/json"
	"log"

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
