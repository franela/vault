package add

import (
	. "github.com/franela/goblin"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/franela/vault/vault/testutils"
	"github.com/mitchellh/cli"
	"testing"
)

func TestAdd(t *testing.T) {
	g := Goblin(t)

	g.Describe("Add", func() {
		g.BeforeEach(func() {
			vault.SetHomeDir(testutils.GetTemporaryHomeDir())
			ui.DEBUG = true
		})

		g.AfterEach(func() {
			testutils.RemoveTemporaryHomeDir(vault.UnsetHomeDir())
			ui.DEBUG = false
		})

		g.Describe("#Run", func() {
			g.It("Should not fail if recipient already exist", func() {
				v := vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{vault.VaultRecipient{Fingerprint: "2B13EC3B5769013E2ED29AC9643E01FBCE44E394", Name: "bob@example.com"}}
				v.Save()

				c, _ := Factory()
				code := c.Run([]string{"2B13EC3B5769013E2ED29AC9643E01FBCE44E394:bob@example.com"})
				g.Assert(code).Equal(0)

				loadedVault, _ := vault.LoadVaultfile()

				g.Assert(v.Recipients).Equal(loadedVault.Recipients)
			})
			g.It("Should allow to add multiple recipients", func() {
				v := vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{vault.VaultRecipient{Fingerprint: "2B13EC3B5769013E2ED29AC9643E01FBCE44E394", Name: "bob@example.com"}}
				v.Save()

				c, _ := Factory()

				c.Run([]string{"39A595E45C6C23693074BDA2A74BFF324DC55DBE:alice@example.com", "69A595E45C6C23693075BDA2A74BFF324DC55DBF:third@example.com"})

				newVaultfile, _ := vault.LoadVaultfile()

				expectedRecipients := []vault.VaultRecipient{
					vault.VaultRecipient{Fingerprint: "2B13EC3B5769013E2ED29AC9643E01FBCE44E394", Name: "bob@example.com"},
					vault.VaultRecipient{Fingerprint: "39A595E45C6C23693074BDA2A74BFF324DC55DBE", Name: "alice@example.com"},
					vault.VaultRecipient{Fingerprint: "69A595E45C6C23693075BDA2A74BFF324DC55DBF", Name: "third@example.com"},
				}

				g.Assert(newVaultfile.Recipients).Equal(expectedRecipients)

			})

			g.It("Should add new recipients", func() {
				v := vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{vault.VaultRecipient{Fingerprint: "2B13EC3B5769013E2ED29AC9643E01FBCE44E394", Name: "bob@example.com"}}
				v.Save()

				c, _ := Factory()

				repairCommand := cli.MockCommand{}
				addCmd, _ := c.(addCommand)
				addCmd.Repair = &repairCommand

				addCmd.Run([]string{"39A595E45C6C23693074BDA2A74BFF324DC55DBE:alice@example.com"})

				expectedRecipients := []vault.VaultRecipient{
					vault.VaultRecipient{Fingerprint: "2B13EC3B5769013E2ED29AC9643E01FBCE44E394", Name: "bob@example.com"},
					vault.VaultRecipient{Fingerprint: "39A595E45C6C23693074BDA2A74BFF324DC55DBE", Name: "alice@example.com"},
				}

				newVaultfile, _ := vault.LoadVaultfile()
				g.Assert(newVaultfile.Recipients).Equal(expectedRecipients)
				g.Assert(repairCommand.RunCalled).IsTrue()
			})

			g.It("Should add recipients to empty Vaultfile", func() {
				c, _ := Factory()

				repairCommand := cli.MockCommand{}
				addCmd, _ := c.(addCommand)
				addCmd.Repair = &repairCommand

				addCmd.Run([]string{"39A595E45C6C23693074BDA2A74BFF324DC55DBE:alice@example.com"})

				newVaultfile, _ := vault.LoadVaultfile()

				expectedRecipients := []vault.VaultRecipient{
					vault.VaultRecipient{Fingerprint: "39A595E45C6C23693074BDA2A74BFF324DC55DBE", Name: "alice@example.com"},
				}

				g.Assert(newVaultfile.Recipients).Equal(expectedRecipients)
				g.Assert(repairCommand.RunCalled).IsTrue()
			})

			g.It("Should print usage if no parameters are sent", func() {
				c, _ := Factory()

				addCmd, _ := c.(addCommand)
				code := addCmd.Run([]string{})

				g.Assert(code).Equal(1)
				g.Assert(ui.GetOutput()).Equal(addHelpText)
			})
		})
	})
}
