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
				v.Recipients = []string{"bob@example.com"}
				v.Save()

				c, _ := Factory()
				code := c.Run([]string{"bob@example.com"})
				g.Assert(code).Equal(0)

				loadedVault, _ := vault.LoadVaultfile()

				g.Assert(v.Recipients).Equal(loadedVault.Recipients)
			})
			g.It("Should allow to add multiple recipients", func() {
				v := vault.Vaultfile{}
				v.Recipients = []string{"bob@example.com"}
				v.Save()

				c, _ := Factory()

				c.Run([]string{"alice@example.com", "third@example.com"})

				newVaultfile, _ := vault.LoadVaultfile()
				g.Assert(newVaultfile.Recipients).Equal([]string{"bob@example.com", "alice@example.com", "third@example.com"})

			})

			g.It("Should add new recipients", func() {
				v := vault.Vaultfile{}
				v.Recipients = []string{"bob@example.com"}
				v.Save()

				c, _ := Factory()

				repairCommand := cli.MockCommand{}
				addCmd, _ := c.(addCommand)
				addCmd.Repair = &repairCommand

				addCmd.Run([]string{"alice@example.com"})

				newVaultfile, _ := vault.LoadVaultfile()
				g.Assert(newVaultfile.Recipients).Equal([]string{"bob@example.com", "alice@example.com"})
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
