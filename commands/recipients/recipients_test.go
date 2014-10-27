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
				v := vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{
					vault.VaultRecipient{Fingerprint: "2B13EC3B5769013E2ED29AC9643E01FBCE44E394", Name: "bob@example.com"},
					vault.VaultRecipient{Fingerprint: "39A595E45C6C23693074BDA2A74BFF324DC55DBE", Name: "alice@example.com"},
				}
				v.Save()

				c, _ := Factory()

				c.Run([]string{})

				listedRecipients := ui.GetOutput()

				g.Assert(strings.Contains(listedRecipients, v.Recipients[0].ToString())).IsTrue()
				g.Assert(strings.Contains(listedRecipients, v.Recipients[1].ToString())).IsTrue()
			})
		})
	})
}
