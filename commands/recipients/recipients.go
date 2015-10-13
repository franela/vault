package recipients

import (
	"github.com/fatih/color"
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/mitchellh/cli"
)

const recipientsHelpText = `
Usage

    Lists current vault recipients.
`

var missingRecipients = []string{}
var trustedRecipients = []string{}
var untrustedRecipients = []string{}

func Factory() (cli.Command, error) {
	return recipientsCommand{}, nil
}

type recipientsCommand struct {
}

func (recipientsCommand) Help() string {
	return recipientsHelpText
}

func (recipientsCommand) Run(args []string) int {
	if vaultFile, err := vault.LoadVaultfile(); err != nil {
		ui.Printf("Error opening Vaultfile: %s", err)
		return 1
	} else {
		if gpgFingerprints, err := getGPGFingerprints(); err != nil {
			return 1
		} else {
			for _, recipient := range vaultFile.Recipients {
				if isTrusty, exists := gpgFingerprints[recipient.Fingerprint]; exists {
					if isTrusty {
						trustedRecipients = append(trustedRecipients, recipient.ToString())
					} else {
						untrustedRecipients = append(untrustedRecipients, recipient.ToString())
					}

				} else {
					missingRecipients = append(missingRecipients, recipient.ToString())
				}
			}

		}
	}
	if len(trustedRecipients) > 0 {
		ui.Printf("Trusted recipients:\n")
		for _, recipient := range trustedRecipients {
			ui.Printf("    %s\n", color.GreenString(recipient))
		}
	}

	if len(untrustedRecipients) > 0 {
		ui.Printf("Untrusted recipients:\n")
		ui.Printf("  (use gpg --sign-key <fingerprint> to sign recipient key):\n")
		for _, recipient := range untrustedRecipients {
			ui.Printf("    %s\n", color.CyanString(recipient))
		}
	}

	if len(missingRecipients) > 0 {
		ui.Printf("Missing recipients:\n")
		ui.Printf("  (use vault import to add missing recipients):\n")
		for _, recipient := range missingRecipients {
			ui.Printf("    %s\n", color.RedString(recipient))
		}
	}
	return 0

}

func (recipientsCommand) Synopsis() string {
	return "Lists all your current vault recipients."
}

func getGPGFingerprints() (map[string]bool, error) {
	return gpg.GetKeysWithFingerprints()
}
