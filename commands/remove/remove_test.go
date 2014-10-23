package remove

import (
	. "github.com/franela/goblin"
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
		})

		g.AfterEach(func() {
			testutils.RemoveTemporaryHomeDir(vault.UnsetHomeDir())
		})

		g.Describe("#Run", func() {
			g.It("Should not fail if recipient doesn't exist", func() {
				v := vault.Vaultfile{}
				v.Recipients = []string{"bob@example.com"}
				v.Save()

				c, _ := Factory()

				code := c.Run([]string{"alice@example.com"})

				g.Assert(code).Equal(0)
			})

			g.It("Should allow removal of multiple recipients", func() {
				v := vault.Vaultfile{}
				v.Recipients = []string{"bob@example.com", "alice@example.com"}
				v.Save()

				c, _ := Factory()

				code := c.Run([]string{"alice@example.com", "bob@example.com"})

				g.Assert(code).Equal(0)
				newVaultfile, _ := vault.LoadVaultfile()
				g.Assert(newVaultfile.Recipients).Equal([]string{})
			})

			g.It("Should remove recipients", func() {
				v := vault.Vaultfile{}
				v.Recipients = []string{"bob@example.com", "alice@example.com"}
				v.Save()

				c, _ := Factory()

				repairCommand := cli.MockCommand{}
				removeCmd, _ := c.(removeCommand)
				removeCmd.Repair = &repairCommand

				code := removeCmd.Run([]string{"alice@example.com"})

				newVaultfile, _ := vault.LoadVaultfile()
				g.Assert(newVaultfile.Recipients).Equal([]string{"bob@example.com"})
				g.Assert(repairCommand.RunCalled).IsTrue()
				g.Assert(code).Equal(0)
			})
		})
	})
}
