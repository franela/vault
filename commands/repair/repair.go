package repair

import (
	"github.com/mitchellh/cli"
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

//	vaultFile, err := vault.LoadVaultfile()
//
//	if err != nil {
//		ui.Printf("%s", err)
//		return 1
//	}
//
//	cmdFlags := flag.NewFlagSet("repair", flag.ContinueOnError)
//
//	var fileName string
//
//	cmdFlags.StringVar(&fileName, "f", "", "specify the file to encrypt")
//
//	if err := cmdFlags.Parse(args); err != nil {
//		return 1
//	}
//
//	args = cmdFlags.Args()
//
//	if len(fileName) > 0 {
//		vaultPath := args[0]
//		err := gpg.EncryptFile(path.Join(vault.GetHomeDir(), vaultPath), fileName, vaultFile.Recipients)
//		if err != nil {
//			ui.Printf("%s", err)
//			return 1
//		}
//	} else {
//		text := args[0]
//		vaultPath := args[1]
//
//		err := gpg.Encrypt(path.Join(vault.GetHomeDir(), vaultPath), text, vaultFile.Recipients)
//
//		if err != nil {
//			ui.Printf("%s", err)
//			return 1
//		}
//	}

	return 0
}

func (repairCommand) Synopsis() string {
	return ""
}
