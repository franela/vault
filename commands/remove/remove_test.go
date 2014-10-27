package remove

import (
	. "github.com/franela/goblin"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/franela/vault/vault/testutils"
	"github.com/mitchellh/cli"
	"testing"
)

func TestRemove(t *testing.T) {
	g := Goblin(t)

	g.Describe("Remove", func() {
		g.BeforeEach(func() {
			vault.SetHomeDir(testutils.GetTemporaryHomeDir())
			ui.DEBUG = true
		})

		g.AfterEach(func() {
			testutils.RemoveTemporaryHomeDir(vault.UnsetHomeDir())
			ui.DEBUG = false
		})

		g.Describe("#Run", func() {
			g.It("Should not fail if recipient doesn't exist", func() {
				v := vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{
					vault.VaultRecipient{Fingerprint: "2B13EC3B5769013E2ED29AC9643E01FBCE44E394", Name: "bob@example.com"},
				}
				v.Save()

				c, _ := Factory()

				code := c.Run([]string{"39A595E45C6C23693074BDA2A74BFF324DC55DBE"})

				g.Assert(code).Equal(0)
			})

			g.It("Should allow removal of multiple recipients", func() {
				v := vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{
					vault.VaultRecipient{Fingerprint: "2B13EC3B5769013E2ED29AC9643E01FBCE44E394", Name: "bob@example.com"},
					vault.VaultRecipient{Fingerprint: "39A595E45C6C23693074BDA2A74BFF324DC55DBE", Name: "alice@example.com"},
				}
				v.Save()

				c, _ := Factory()

				code := c.Run([]string{"2B13EC3B5769013E2ED29AC9643E01FBCE44E394", "39A595E45C6C23693074BDA2A74BFF324DC55DBE"})

				g.Assert(code).Equal(0)
				newVaultfile, _ := vault.LoadVaultfile()
				g.Assert(newVaultfile.Recipients).Equal([]vault.VaultRecipient{})
			})

			g.It("Should remove recipients", func() {
				v := vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{
					vault.VaultRecipient{Fingerprint: "2B13EC3B5769013E2ED29AC9643E01FBCE44E394", Name: "bob@example.com"},
					vault.VaultRecipient{Fingerprint: "39A595E45C6C23693074BDA2A74BFF324DC55DBE", Name: "alice@example.com"},
				}
				v.Save()

				c, _ := Factory()

				repairCommand := cli.MockCommand{}
				removeCmd, _ := c.(removeCommand)
				removeCmd.Repair = &repairCommand

				code := removeCmd.Run([]string{"39A595E45C6C23693074BDA2A74BFF324DC55DBE"})

				newVaultfile, _ := vault.LoadVaultfile()
				g.Assert(newVaultfile.Recipients).Equal([]vault.VaultRecipient{vault.VaultRecipient{Fingerprint: "2B13EC3B5769013E2ED29AC9643E01FBCE44E394", Name: "bob@example.com"}})
				g.Assert(repairCommand.RunCalled).IsTrue()
				g.Assert(code).Equal(0)
			})

			g.It("Should print usage incorrect number of parameters are sent", func() {
				c, _ := Factory()

				removeCommand, _ := c.(removeCommand)
				code := removeCommand.Run([]string{})

				g.Assert(code).Equal(1)
				g.Assert(ui.GetOutput()).Equal(removeHelpText)
			})
		})
	})
}
