package recipients

import (
	. "github.com/franela/goblin"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/franela/vault/vault/testutils"
	"strings"
	"testing"
)

func TestRecipients(t *testing.T) {
	g := Goblin(t)

	g.Describe("Recipients", func() {
		g.Describe("#Run", func() {
			g.BeforeEach(func() {
				vault.SetHomeDir(testutils.GetTemporaryHomeDir())
				ui.DEBUG = true
			})

			g.AfterEach(func() {
				testutils.RemoveTemporaryHomeDir(vault.UnsetHomeDir())
				ui.DEBUG = false
			})

			g.It("Should output all the recipients in the Vaultfile", func() {
				vault := vault.Vaultfile{}
				vault.Recipients = []string{"bob@example.com", "alice@example.com"}
				vault.Save()

				c, _ := Factory()

				c.Run([]string{})

				listedRecipients := ui.GetOutput()

				g.Assert(strings.Contains(listedRecipients, vault.Recipients[0])).IsTrue()
				g.Assert(strings.Contains(listedRecipients, vault.Recipients[1])).IsTrue()
			})
		})
	})
}
