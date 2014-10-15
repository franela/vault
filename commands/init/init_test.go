package init

import (
	"testing"

	. "github.com/franela/goblin"
	"github.com/franela/vault/vault"
	"github.com/franela/vault/vault/testutils"
)

func TestInit(t *testing.T) {
	g := Goblin(t)

	g.Describe("Init", func() {
		g.Describe("#Run", func() {
			g.BeforeEach(func() {
				vault.SetHomeDir(testutils.GetTemporaryHomeDir())
			})

			g.AfterEach(func() {
				testutils.RemoveTemporaryHomeDir(vault.UnsetHomeDir())
			})

			g.It("Should create a Vaultfile with recipients", func() {
				desiredRecipients := []string{"a@a.com", "b@b.com"}
				c, _ := Factory()
				exitCode := c.Run(desiredRecipients)

				v, err := vault.LoadVaultfile()

				g.Assert(err == nil).IsTrue()
				g.Assert(v.Recipients).Equal(desiredRecipients)
				g.Assert(exitCode).Equal(0)
			})
		})
	})
}
