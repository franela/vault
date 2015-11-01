package init

import (
	"flag"

	"github.com/franela/vault/gpg"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
)

const initHelpText = `
Usage: vault init [options] RECIPIENT [RECIPIENT ...]

Initializes the vault with a space separated list of recipients.
Each recipient identifies someone who can read and write items
in the vault.

Options: 
    
    --omit-self	Omits keyring owner in Vaultfile when initializing a new vault

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

	cmdFlags := flag.NewFlagSet("init", flag.ContinueOnError)
	var omitSelf bool

	cmdFlags.BoolVar(&omitSelf, "omit-self", false, "Don't add keyring owner to vaultfile")

	if err := cmdFlags.Parse(args); err != nil {
		ui.Printf(initHelpText)
		return 1
	}

	v := vault.Vaultfile{}

	for _, r := range args {
		if recipient, err := vault.NewRecipient(r); err == nil {
			v.Recipients = append(v.Recipients, *recipient)
		}
	}

	if !omitSelf {
		if ownerRecipient, err := gpg.GetKeyringOwnerRecipient(); err != nil {
			ui.Printf("WARN: %s", err)
		} else {
			v.Recipients = append(v.Recipients, *ownerRecipient)
		}
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
