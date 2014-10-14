package get

import (
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/ui"
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
	file := args[0]

	if text, err := gpg.Decrypt(file); err != nil {
		ui.Printf("Error decrypting file %s %s", file, err)
		return 1
	} else {
		ui.Printf("%s", text)
		return 0
	}
}

func (setCommand) Synopsis() string {
	return ""
}
