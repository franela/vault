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
				initArgs := []string{"--omit-self", "3B9CEC3B5069113E2ED39AC9843E01FBCE44AAAA:a@a.com", "BBBBEC3B5069113E2ED39AC9843E01FBCE44BBBB:b@b.com"}
				c, _ := Factory()
				exitCode := c.Run(initArgs)

				v, err := vault.LoadVaultfile()

				g.Assert(err == nil).IsTrue()
				g.Assert(v.Recipients).Equal([]vault.VaultRecipient{vault.VaultRecipient{Fingerprint: "3B9CEC3B5069113E2ED39AC9843E01FBCE44AAAA", Name: "a@a.com"}, vault.VaultRecipient{Fingerprint: "BBBBEC3B5069113E2ED39AC9843E01FBCE44BBBB", Name: "b@b.com"}})
				g.Assert(exitCode).Equal(0)
			})

			g.It("Should add the default keyring secret key if present ", func() {
				testutils.SetTestGPGHome("bob")
				c, _ := Factory()
				exitCode := c.Run([]string{})

				v, err := vault.LoadVaultfile()

				g.Assert(err == nil).IsTrue()
				g.Assert(v.Recipients).Equal([]vault.VaultRecipient{vault.VaultRecipient{Fingerprint: "2B13EC3B5769013E2ED29AC9643E01FBCE44E394", Name: "Bob Example <bob@example.com>"}})
				g.Assert(exitCode).Equal(0)
			})
			g.It("Should not add the fault secret if --omit-self is supplied", func() {
				c, _ := Factory()
				exitCode := c.Run([]string{"--omit-self"})

				v, err := vault.LoadVaultfile()

				var emptyRecipients []vault.VaultRecipient
				g.Assert(err == nil).IsTrue()
				g.Assert(v.Recipients).Equal(emptyRecipients)
				g.Assert(exitCode).Equal(0)
			})
		})
	})
}
