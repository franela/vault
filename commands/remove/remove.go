package remove

import (
	"fmt"
	"github.com/franela/vault/commands/repair"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
)

const removeHelpText = `
Usage: vault remove recipients...

    Remove specified recipients from the vaultfile. If
    specified recipient doesn't exist, vault will ignore it
`

func Factory() (cli.Command, error) {
	repairCmd, err := repair.Factory()

	if err != nil {
		panic(err)
	}

	return removeCommand{Repair: repairCmd}, nil
}

type removeCommand struct {
	Repair cli.Command
}

func (removeCommand) Help() string {
	return removeHelpText
}

func (self removeCommand) Run(args []string) int {

	if len(args) == 0 {
		ui.Printf(removeHelpText)
		return 1
	}
	if vaultFile, err := vault.LoadVaultfile(); err != nil {
		ui.Printf("Error opening Vaultfile: %s", err)
		return 1
	} else {
		if len(vaultFile.Recipients) == 0 {
			// There is nothing to do, just return
			return 0
		}

		for _, recipient := range args {
			pos := -1
			for i, r := range vaultFile.Recipients {
				if r == recipient {
					pos = i
				}
			}
			if pos != -1 {
				if len(vaultFile.Recipients) == pos+1 {
					// Removing last element
					vaultFile.Recipients = vaultFile.Recipients[:pos]
				} else {
					fmt.Println(vaultFile.Recipients)
					vaultFile.Recipients = append(vaultFile.Recipients[:pos], vaultFile.Recipients[pos+1:]...)
				}
			}
		}
		if err := vaultFile.Save(); err != nil {
			ui.Printf("Error saving Vaultfile: %s", err)
			return 1
		}

		return self.Repair.Run([]string{})
	}
}

func (removeCommand) Synopsis() string {
	return "Removes one or many recipients from your vault"
}
