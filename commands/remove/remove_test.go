package remove

import (
	. "github.com/franela/goblin"
	"github.com/franela/vault/vault"
	"github.com/franela/vault/vault/testutils"
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
			g.It("Should not fail if recipient doesn't exist")
			g.It("Should allow removal of multiple recipients")

			g.It("Should remove recipients", func() {
				v := vault.Vaultfile{}
				v.Recipients = []string{"bob@example.com", "alice@example.com"}
				v.Save()

				c, _ := Factory()

				c.Run([]string{"alice@example.com"})

				newVaultfile, _ := vault.LoadVaultfile()
				g.Assert(newVaultfile.Recipients).Equal([]string{"bob@example.com"})
			})
		})
	})
}