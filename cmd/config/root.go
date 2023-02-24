package config

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
	config_ "github.com/lazytangent/cfgo/config"
)

var Cmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration stuff",
	Run:   Run,
}

var (
	list bool
)

func init() {
	Cmd.PersistentFlags().BoolVarP(&list, "list", "l", false, "List values from config")
}

const configFile = "~/.config/cfgo/config.toml"

func Run(cmd *cobra.Command, args []string) {
	if list {
		listConfig()
	}
}

func listConfig() {
	configPath := getConfigFile()
	dat, err := os.ReadFile(configPath);
	if err != nil {
		log.Fatal(err)
	}
	configData := string(dat)

	cfg := config_.Parse(configData)
	fmt.Printf("%v\n", cfg.Print())
}

func getConfigFile() string {
	usr, _ := user.Current()
	dir := usr.HomeDir

	return filepath.Join(dir, configFile[2:])
}
