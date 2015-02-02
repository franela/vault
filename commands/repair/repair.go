package repair

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/franela/vault/gpg"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
)

const repairHelpText = `
    Ensures that all your vault files are encrypted for the recipients in your vault.
`

func Factory() (cli.Command, error) {
	return repairCommand{}, nil
}

type repairCommand struct {
}

func (repairCommand) Help() string {
	return repairHelpText
}

func (repairCommand) Run(args []string) int {
	vaultFile, err := vault.LoadVaultfile()

	if err != nil {
		ui.Printf("%s\n", err)
		return 1
	}

	if err := filepath.Walk(vault.GetHomeDir(), func(filepath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if path.Ext(filepath) != ".asc" {
			return nil
		}
		log.Printf("Re-encrypting %s\n", filepath)
		if err := gpg.ReEncryptFile(filepath, filepath, vaultFile.Recipients); err != nil {
			ui.PrintErrorf("Error trying to re-encrypt file %s\n", filepath)
		}
		return nil
	}); err != nil {
		return 1
	}

	return 0
}

func (repairCommand) Synopsis() string {
	return "Re-encrypts all your vault files for your current recipients."
}
