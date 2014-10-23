package recipients

import (
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
)

const recipientsHelpText = `
`

func Factory() (cli.Command, error) {
	return recipientsCommand{}, nil
}

type recipientsCommand struct {
}

func (recipientsCommand) Help() string {
	return recipientsHelpText
}

func (recipientsCommand) Run(args []string) int {
	if vaultFile, err := vault.LoadVaultfile(); err != nil {
		ui.Printf("Error opening Vaultfile: %s", err)
		return 1
	} else {
		for _, recipient := range vaultFile.Recipients {
			ui.Printf("%s\n", recipient)
		}
		return 0
	}
}

func (recipientsCommand) Synopsis() string {
	return "Lists all your current vault recipients."
}
