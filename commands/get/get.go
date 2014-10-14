package get

import (
	"github.com/mitchellh/cli"
)

const setHelpText = `
`

func Factory() (cli.Command, error) {
	return setCommand{}, nil
}

type setCommand struct {
}

func (setCommand) Help() string {
	return setHelpText
}

func (setCommand) Run(args []string) int {
	return 1
}

func (setCommand) Synopsis() string {
	return ""
}
