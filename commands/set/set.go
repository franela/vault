package set

import (
	"flag"
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
)

const setHelpText = `
`

func Factory() (cli.Command, error) {
	return setCommand{}, nil
}

type setCommand struct {
}

func (setCommand) Help() string {
	return setHelpText
}

func (setCommand) Run(args []string) int {

	vaultFile, err := vault.LoadVaultfile()

	if err != nil {
		ui.Printf("%s", err)
		return 1
	}

	cmdFlags := flag.NewFlagSet("set", flag.ContinueOnError)

	var fileName string

	cmdFlags.StringVar(&fileName, "f", "", "specify the file to encrypt")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	args = cmdFlags.Args()

	if len(fileName) > 0 {
		path := args[0]
		err := gpg.EncryptFile(path, fileName, vaultFile.Recipients)
		if err != nil {
			ui.Printf("%s", err)
			return 1
		}
	} else {
		text := args[0]
		path := args[1]

		err := gpg.Encrypt(path, text, vaultFile.Recipients)

		if err != nil {
			ui.Printf("%s", err)
			return 1
		}
	}

	return 0
}

func (setCommand) Synopsis() string {
	return ""
}
