package get

import (
	. "github.com/franela/goblin"
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"io/ioutil"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	g := Goblin(t)

	g.Describe("Get", func() {

		wd, _ := vault.GetHomeDir()
		ui.DEBUG = true
		os.Setenv("GNUPGHOME", wd+"/testdata/bob")

		g.AfterEach(func() {
			os.Remove(wd + "/Vaultfile")
			os.Remove(wd + "/tmp/get_test")
			os.Remove(wd + "/tmp/get_test_output")
		})

		g.Describe("#Run", func() {
			g.It("Should output decrypted text of a given file", func() {
				vault := vault.Vaultfile{}
				vault.Recipients = []string{"bob@example.com"}
				vault.Save()

				gpg.Encrypt(wd+"/tmp/get_test", "This is a test", vault.Recipients)

				c, _ := Factory()
				c.Run([]string{wd + "/tmp/get_test"})

				g.Assert(ui.GetOutput()).Equal("This is a test")
			})

			g.It("Should create a file with decrypted text of a given file", func() {
				vault := vault.Vaultfile{}
				vault.Recipients = []string{"bob@example.com"}
				vault.Save()

				gpg.Encrypt(wd+"/tmp/get_test", "This is a test", vault.Recipients)

				c, _ := Factory()
				c.Run([]string{"-o", wd + "/tmp/get_test_output", wd + "/tmp/get_test"})

				output, _ := ioutil.ReadFile(wd + "/tmp/get_test_output")
				g.Assert(string(output)).Equal("This is a test")
			})
		})
	})
}
