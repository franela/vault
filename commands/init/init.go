package init

import (
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
)

const initHelpText = `
Usage: vault init RECIPIENT [RECIPIENT ...]

Initializes the vault with a space separated list of recipients.
Each recipient identifies someone who can read and write items
in the vault.

`

func Factory() (cli.Command, error) {
	return initCommand{}, nil
}

type initCommand struct {
}

func (initCommand) Help() string {
	return initHelpText
}

func (initCommand) Run(args []string) int {
	v := vault.Vaultfile{}
	for _, r := range args {
		v.Recipients = append(v.Recipients, vault.NewRecipient(r))
	}
	err := v.Save()

	if err != nil {
		ui.Printf("%s", err)
		return 1
	}

	return 0
}

func (initCommand) Synopsis() string {
	return "Initializes a new vault and sets the recipients (people who are allowed to get and set items in the vault)"
}
