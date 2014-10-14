package set

import (
	. "github.com/franela/goblin"
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/vault"
	"os"
	"testing"
)

func TestSet(t *testing.T) {
	g := Goblin(t)

	g.Describe("Set", func() {
		wd, _ := vault.GetHomeDir()

		g.Describe("#Run", func() {

			g.AfterEach(func() {
				os.Remove(wd + "/tmp/set_test")
				os.Remove(wd + "/Vaultfile")
			})

			g.It("Should create an encrypted file given a text", func() {
				v := &vault.Vaultfile{}
				v.Recipients = []string{"bob@example.com"}
				v.Save()
				os.Setenv("GNUPGHOME", wd+"/testdata/bob")
				c, _ := Factory()
				c.Run([]string{"This is a test", wd + "/tmp/set_test"})
				_, err := os.Stat(wd + "/tmp/set_test")
				g.Assert(err == nil).IsTrue()
				out, err := gpg.Decrypt(wd + "/tmp/set_test")
				g.Assert(err == nil).IsTrue()
				g.Assert(out).Equal("This is a test")
			})
			g.It("Should create an encrypted file given another file", func() {
				v := &vault.Vaultfile{}
				v.Recipients = []string{"bob@example.com"}
				v.Save()
				os.Setenv("GNUPGHOME", wd+"/testdata/bob")
				c, _ := Factory()

				c.Run([]string{"-f", wd + "/testdata/set_test", wd + "/tmp/set_test"})

				_, err := os.Stat(wd + "/tmp/set_test")
				g.Assert(err == nil).IsTrue()
				out, err := gpg.Decrypt(wd + "/tmp/set_test")
				g.Assert(err == nil).IsTrue()
				g.Assert(out).Equal("This is a test")
			})

			g.It("Should encrypt only for recipients in the Vaultfile")
			g.It("Should fail if Vaultfile recipients is empty")
			g.It("Should fail if file to encrypt doesn't exist or cannot be accesed")
			g.It("Should fail if encrypted file cannot be saved")
			g.It("Should fail if encrypted file path starts with '..' or '/'")
		})
	})
}
