package get

import (
	"flag"
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
	"path"
	"path/filepath"
)

const getHelpText = `
Usage: vault get [options]

  Add specified recipients to the Vaultfile. If specified recipients
  already exist, vault will ignore them
`

func Factory() (cli.Command, error) {
	return getCommand{}, nil
}

type getCommand struct {
}

func (getCommand) Help() string {
	return getHelpText
}

func (getCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("get", flag.ContinueOnError)

	var outputFile string

	cmdFlags.StringVar(&outputFile, "o", "", "specify the output file to store decrypted text")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	args = cmdFlags.Args()

	if len(args) != 1 {
		ui.Printf(getHelpText)
		return 1
	}

	file := args[0]
	if filepath.Ext(file) != ".asc" {
		file = file + ".asc"
	}

	if len(outputFile) > 0 {
		if err := gpg.DecryptFile(outputFile, path.Join(vault.GetHomeDir(), file)); err != nil {
			ui.Printf("Error decrypting file %s %s", file, err)
			return 1
		}
	} else {
		if text, err := gpg.Decrypt(path.Join(vault.GetHomeDir(), file)); err != nil {
			ui.Printf("Error decrypting file %s %s", file, err)
			return 1
		} else {
			ui.Printf("%s", text)
		}
	}
	return 0
}

func (getCommand) Synopsis() string {
	return ""
}
