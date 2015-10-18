package edit

import (
	"path"

	"github.com/franela/vault/gpg"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mantika/govipe"
	"github.com/mitchellh/cli"
)

const editHelpText = `
Usage: vault edit vaultfile


`

func Factory() (cli.Command, error) {
	return editCommand{}, nil
}

type editCommand struct {
}

func (editCommand) Help() string {
	return editHelpText
}

func (self editCommand) Run(args []string) int {

	vaultFile, err := vault.LoadVaultfile()

	if err != nil {
		ui.Printf("%s", err)
		return 1
	}

	if len(vaultFile.Recipients) == 0 {
		ui.Printf("Cannot edit in vault if Vaultfile has no recipients. Use `vault add` to add one or more recipients.\n")
		return 3
	}

	if len(args) != 1 {
		ui.Printf(editHelpText)
		return 1
	}

	file := args[0]

	if text, err := gpg.Decrypt(path.Join(vault.GetHomeDir(), file)); err != nil {
		ui.Printf("Error decrypting file %s %s\n", file, err)
		return 1
	} else {
		if output, err := govipe.Edit([]byte(text)); err != nil {
			ui.Printf("Error editing file %s %s\n", file, err)
			return 1
		} else {
			err := gpg.Encrypt(path.Join(vault.GetHomeDir(), file), string(output), vaultFile.Recipients)
			if err != nil {
				ui.Printf("%s", err)
				return 1
			}
		}
	}
	return 0

}

func (editCommand) Synopsis() string {
	return "Edits a vault file using default system editor."
}
