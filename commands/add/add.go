package add

import (
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
)

const addHelpText = `
`

func Factory() (cli.Command, error) {
	return addCommand{}, nil
}

type addCommand struct {
}

func (addCommand) Help() string {
	return addHelpText
}

func (addCommand) Run(args []string) int {
	if vaultFile, err := vault.LoadVaultfile(); err != nil {
		ui.Printf("Error opening Vaultfile: %s", err)
		return 1
	} else {
		recipient := args[0]

		vaultFile.Recipients = append(vaultFile.Recipients, recipient)
		if err := vaultFile.Save(); err != nil {
			ui.Printf("Error saving Vaultfile: %s", err)
			return 1
		}

		return 0
	}
}

func (addCommand) Synopsis() string {
	return ""
}
