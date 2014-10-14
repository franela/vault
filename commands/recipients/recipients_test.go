package recipients

import (
	. "github.com/franela/goblin"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"os"
	"strings"
	"testing"
)

func TestRecipients(t *testing.T) {
	g := Goblin(t)

	g.Describe("Recipients", func() {
		g.Describe("#Run", func() {
			wd, _ := vault.GetHomeDir()
			ui.DEBUG = true

			g.It("Should output all the recipients in the Vaultfile", func() {
				defer func() {
					os.Remove(wd + "/Vaultfile")
				}()

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
