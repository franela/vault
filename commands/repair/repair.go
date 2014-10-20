package repair

import (
	"fmt"
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
	"os"
	"path"
	"path/filepath"
)

const repairHelpText = `
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
		ui.Printf("%s", err)
		return 1
	}

	if err := filepath.Walk(vault.GetHomeDir(), func(filepath string, info os.FileInfo, err error) error {
		if path.Base(filepath) == "Vaultfile" {
			return nil
		}
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fmt.Println(filepath)
		if err := gpg.ReEncryptFile(filepath, filepath, vaultFile.Recipients); err != nil {
			return err
		}
		return nil
	}); err != nil {
		ui.Printf("%s", err)
		return 1
	}

	return 0
}

func (repairCommand) Synopsis() string {
	return ""
}
