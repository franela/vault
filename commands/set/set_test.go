package set

import (
	. "github.com/franela/goblin"
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/franela/vault/vault/testutils"
	"os"
	"path"
	"testing"
)

func TestSet(t *testing.T) {
	g := Goblin(t)

	g.Describe("Set", func() {
		g.Describe("#Run", func() {
			g.BeforeEach(func() {
				vault.SetHomeDir(testutils.GetTemporaryHomeDir())
				ui.DEBUG = true
			})

			g.AfterEach(func() {
				testutils.RemoveTemporaryHomeDir(vault.UnsetHomeDir())
				ui.DEBUG = false
				ui.GetOutput() //Cleans the output
			})

			g.It("Should create an encrypted file given a text", func() {
				testutils.SetTestGPGHome("bob")

				v := &vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{
					vault.NewRecipient("2B13EC3B5769013E2ED29AC9643E01FBCE44E394:bob@example.com"),
				}
				v.Save()

				c, _ := Factory()
				c.Run([]string{"set_test.asc", "This is a test"})

				_, err := os.Stat(path.Join(vault.GetHomeDir(), "set_test.asc"))
				g.Assert(err == nil).IsTrue()

				out, err := gpg.Decrypt(path.Join(vault.GetHomeDir(), "set_test.asc"))
				g.Assert(err == nil).IsTrue()
				g.Assert(out).Equal("This is a test")
			})
			g.It("Should add .asc extension if not specified", func() {
				testutils.SetTestGPGHome("bob")

				v := &vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{
					vault.NewRecipient("2B13EC3B5769013E2ED29AC9643E01FBCE44E394:bob@example.com"),
				}
				v.Save()

				c, _ := Factory()
				c.Run([]string{"set_test", "This is a test"})

				_, err := os.Stat(path.Join(vault.GetHomeDir(), "set_test.asc"))
				g.Assert(err == nil).IsTrue()

				out, err := gpg.Decrypt(path.Join(vault.GetHomeDir(), "set_test.asc"))
				g.Assert(err == nil).IsTrue()
				g.Assert(out).Equal("This is a test")
			})
			g.It("Should create an encrypted file given another file", func() {
				testutils.SetTestGPGHome("bob")

				v := &vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{
					vault.NewRecipient("2B13EC3B5769013E2ED29AC9643E01FBCE44E394:bob@example.com"),
				}
				v.Save()

				c, _ := Factory()

				c.Run([]string{"-f", path.Join(testutils.GetProjectDir(), "testdata", "set_test"), "set_test"})

				_, err := os.Stat(path.Join(vault.GetHomeDir(), "set_test.asc"))
				g.Assert(err == nil).IsTrue()

				out, err := gpg.Decrypt(path.Join(vault.GetHomeDir(), "set_test.asc"))
				g.Assert(err == nil).IsTrue()
				g.Assert(out).Equal("This is a test")
			})

			g.It("Should fail if Vaultfile recipients is empty", func() {
				c, _ := Factory()

				code := c.Run([]string{"this is a test", "set_test"})

				g.Assert(code).Equal(3)
			})

			g.It("Should fail if file to encrypt doesn't exist or cannot be accesed", func() {
				testutils.SetTestGPGHome("bob")

				v := &vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{
					vault.NewRecipient("2B13EC3B5769013E2ED29AC9643E01FBCE44E394:bob@example.com"),
				}
				v.Save()

				c, _ := Factory()

				code := c.Run([]string{"-f", path.Join("some", "crazy", "unexisting", "path"), "set_test"})

				g.Assert(code).Equal(1)
			})

			g.It("Should print usage incorrect number of parameters are sent", func() {
				testutils.SetTestGPGHome("bob")
				v := &vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{
					vault.NewRecipient("2B13EC3B5769013E2ED29AC9643E01FBCE44E394:bob@example.com"),
				}
				v.Save()

				c, _ := Factory()

				code := c.Run([]string{})

				g.Assert(code).Equal(1)
				g.Assert(ui.GetOutput()).Equal(setHelpText)

				code = c.Run([]string{"-f"})

				g.Assert(code).Equal(1)
				g.Assert(ui.GetOutput()).Equal(setHelpText)

				code = c.Run([]string{"/tmp/test"})

				g.Assert(code).Equal(1)
				g.Assert(ui.GetOutput()).Equal(setHelpText)
			})

			g.It("Should fail to encrypt if destination is not under the current path", func() {
				testutils.SetTestGPGHome("bob")
				v := &vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{
					vault.NewRecipient("2B13EC3B5769013E2ED29AC9643E01FBCE44E394:bob@example.com"),
				}
				v.Save()

				c, _ := Factory()
				code := c.Run([]string{"../set_test", "This is a test"})

				g.Assert(code).Equal(1)
			})

			g.It("Should fail to encrypt files if destination is not under the current path", func() {
				testutils.SetTestGPGHome("bob")

				v := &vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{
					vault.NewRecipient("2B13EC3B5769013E2ED29AC9643E01FBCE44E394:bob@example.com"),
				}
				v.Save()

				c, _ := Factory()

				code := c.Run([]string{"-f", path.Join(testutils.GetProjectDir(), "testdata", "set_test"), "/set_test"})

				g.Assert(code).Equal(1)
			})
		})
	})
}
