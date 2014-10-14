package set

import (
        "log"
	"github.com/mitchellh/cli"
        "github.com/franela/vault/gpg"
        "github.com/franela/vault/vault"
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

    path := args[0]
    text := args[1]

    if vaultFile, err := vault.LoadVaultfile(); err != nil {
        log.Print(err)
        return 1
    } else {
        err := gpg.Encrypt(path, text, vaultFile.Recipients)

        if err != nil {
            log.Print(err)
            return 1
        }

        return 0
    }
}

func (setCommand) Synopsis() string {
	return ""
}
