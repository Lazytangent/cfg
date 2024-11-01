package utils

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultLen = 80

func CreateDelimiter(msg string) string {
	if msg == "" || msg == "-" {
		return strings.Repeat("-", defaultLen)
	}

	n := len(msg)
	remaining := defaultLen - n
	halfLen := remaining / 2
	lenOfDash := halfLen - 1
	padding := halfLen*2 < remaining

	left := strings.Repeat("-", lenOfDash)
	mid := fmt.Sprintf(" %s ", msg)
	right := strings.Repeat("-", lenOfDash)
	if padding {
		right = fmt.Sprintf("%s-", right)
	}

	return fmt.Sprintf("%s%s%s", left, mid, right)
}

func LogFatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogPrintlnIfErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func GetGitArgs(cmd *cobra.Command, args []string) []string {
	idx := cmd.ArgsLenAtDash()

	if idx == -1 {
		return []string{}
	}

	return args[idx:]
}

func GetViper(cmd *cobra.Command) (*viper.Viper, error) {
	ctx := cmd.Context()

	vpr := ctx.Value("viper")

	v, ok := vpr.(*viper.Viper)
	if !ok {
		return nil, errors.New("could not cast 'viper' from context to *viper.Viper")
	}

	return v, nil
}
