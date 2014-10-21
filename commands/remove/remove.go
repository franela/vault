package remove

import (
	"fmt"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
)

const removeHelpText = `
`

func Factory() (cli.Command, error) {
	return removeCommand{}, nil
}

type removeCommand struct {
	Repair cli.Command
}

func (removeCommand) Help() string {
	return removeHelpText
}

func (self removeCommand) Run(args []string) int {
	if vaultFile, err := vault.LoadVaultfile(); err != nil {
		ui.Printf("Error opening Vaultfile: %s", err)
		return 1
	} else {
		if len(vaultFile.Recipients) == 0 {
			// There is nothing to do, just return
			return 0
		}

		recipient := args[0]

		pos := 0
		for i, r := range vaultFile.Recipients {
			if r == recipient {
				pos = i
			}
		}
		if len(vaultFile.Recipients) == pos+1 {
			// Removing last element
			vaultFile.Recipients = vaultFile.Recipients[:pos]
		} else {
			fmt.Println(vaultFile.Recipients)
			vaultFile.Recipients = append(vaultFile.Recipients[:pos], vaultFile.Recipients[pos+1:]...)
		}
		if err := vaultFile.Save(); err != nil {
			ui.Printf("Error saving Vaultfile: %s", err)
			return 1
		}

		return self.Repair.Run([]string{})
	}
}

func (removeCommand) Synopsis() string {
	return ""
}
