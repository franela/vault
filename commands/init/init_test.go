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
				desiredRecipients := []string{"3B9CEC3B5069113E2ED39AC9843E01FBCE44AAAA:a@a.com", "BBBBEC3B5069113E2ED39AC9843E01FBCE44BBBB:b@b.com"}
				c, _ := Factory()
				exitCode := c.Run(desiredRecipients)

				v, err := vault.LoadVaultfile()

				g.Assert(err == nil).IsTrue()
				g.Assert(v.Recipients).Equal([]vault.VaultRecipient{vault.NewRecipient("3B9CEC3B5069113E2ED39AC9843E01FBCE44AAAA:a@a.com"), vault.NewRecipient("BBBBEC3B5069113E2ED39AC9843E01FBCE44BBBB:b@b.com")})
				g.Assert(exitCode).Equal(0)
			})
		})
	})
}
