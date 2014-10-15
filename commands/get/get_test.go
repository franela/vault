package get

import (
	. "github.com/franela/goblin"
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/franela/vault/vault/testutils"
	"io/ioutil"
	"path"
	"testing"
)

func TestGet(t *testing.T) {
	g := Goblin(t)

	g.Describe("Get", func() {
		testutils.SetTestGPGHome("bob")

		g.BeforeEach(func() {
			vault.SetHomeDir(testutils.GetTemporaryHomeDir())
			ui.DEBUG = true
		})

		g.AfterEach(func() {
			testutils.RemoveTemporaryHomeDir(vault.UnsetHomeDir())
			ui.DEBUG = false
		})

		g.Describe("#Run", func() {
			g.It("Should output decrypted text of a given file", func() {
				v := vault.Vaultfile{}
				v.Recipients = []string{"bob@example.com"}
				v.Save()

				gpg.Encrypt(path.Join(vault.GetHomeDir(), "get_test"), "This is a test", v.Recipients)

				c, _ := Factory()
				c.Run([]string{"get_test"})

				g.Assert(ui.GetOutput()).Equal("This is a test")
			})

			g.It("Should create a file with decrypted text of a given file", func() {
				v := vault.Vaultfile{}
				v.Recipients = []string{"bob@example.com"}
				v.Save()

				gpg.Encrypt(path.Join(vault.GetHomeDir(), "get_test"), "This is a test", v.Recipients)

				c, _ := Factory()
				c.Run([]string{"-o", path.Join(vault.GetHomeDir(), "get_test_output"), "get_test"})

				output, _ := ioutil.ReadFile(path.Join(vault.GetHomeDir(), "get_test_output"))
				g.Assert(string(output)).Equal("This is a test")
			})

			g.It("Should show help when called without arguments")
		})
	})
}
