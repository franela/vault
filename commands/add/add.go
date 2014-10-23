package add

import (
	"github.com/franela/vault/commands/repair"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
	"log"
	"strings"
)

const addHelpText = `
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
	if vaultFile, err := vault.LoadVaultfile(); err != nil {
		ui.Printf("Error opening Vaultfile: %s", err)
		return 1
	} else {
		for _, recipient := range args {
			if !strings.Contains(strings.Join(vaultFile.Recipients, " "), recipient) {
				log.Printf("Adding [%s] to Vaultfile\n", recipient)
				vaultFile.Recipients = append(vaultFile.Recipients, strings.TrimSpace(recipient))
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
