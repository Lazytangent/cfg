package config

import (
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
