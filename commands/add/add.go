package add

import (
	"github.com/franela/vault/commands/repair"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
	"strings"
)

const addHelpText = `
Usage: vault add recipients...

  Add specified recipients to the Vaultfile. If specified recipients
  already exist, vault will ignore them
`

func Factory() (cli.Command, error) {
	repairCmd, err := repair.Factory()

	if err != nil {
		panic(err)
	}

	return addCommand{Repair: repairCmd}, nil
}

type addCommand struct {
	Repair cli.Command
}

func (addCommand) Help() string {
	return addHelpText
}

func (self addCommand) Run(args []string) int {

	if len(args) == 0 {
		ui.Printf(addHelpText)
		return 1
	}

	if vaultFile, err := vault.LoadVaultfile(); err != nil {
		ui.Printf("Error opening Vaultfile: %s", err)
		return 1
	} else {
		for _, recipient := range args {
			if !strings.Contains(recipient, strings.Join(vaultFile.Recipients, " ")) {
				vaultFile.Recipients = append(vaultFile.Recipients, recipient)
			}
		}
		if err := vaultFile.Save(); err != nil {
			ui.Printf("Error saving Vaultfile: %s", err)
			return 1
		}
		return self.Repair.Run([]string{})

		return 0
	}
}

func (addCommand) Synopsis() string {
	return ""
}
