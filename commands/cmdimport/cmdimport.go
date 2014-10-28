package cmdimport

import (
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
)

const importHelpText = `
Usage: vault import

  Imports all recipients from the Vaultfile by fingerprint.
`

type importCommand struct {
}

func (importCommand) Synopsis() string {
	return "Import all recipients from the Vaultfile"
}

func (self importCommand) Run(args []string) int {
	if vaultFile, err := vault.LoadVaultfile(); err != nil {
		ui.Printf("Error opening Vaultfile: %s", err)
		return 1
	} else {
		errReceive := gpg.ReceiveKey(vaultFile.Recipients)
		if errReceive != nil {
			ui.Printf("%s", err)
			return 1
		}
		return 0
	}
}

func Factory() (cli.Command, error) {
	return importCommand{}, nil
}

func (importCommand) Help() string {
	return importHelpText
}
