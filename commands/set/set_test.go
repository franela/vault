package set

import (
	. "github.com/franela/goblin"
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/vault"
	"os"
	"testing"
)

var (
	cwd = os.Getenv("VAULTDIR")
)

func TestSet(t *testing.T) {
	g := Goblin(t)

	g.Describe("Set", func() {
		g.Describe("#Run", func() {

			g.AfterEach(func() {
				os.Remove(cwd + "/tmp/set_test")
				os.Remove("Vaultfile")
			})

			g.It("Should create an encrypted file given a text", func() {
				v := &vault.Vaultfile{}
				v.Recipients = []string{"bob@example.com"}
				v.Save()
				os.Setenv("GNUPGHOME", cwd+"/testdata/bob")
				c, _ := Factory()
				c.Run([]string{"This is a test", cwd + "/tmp/set_test"})
				_, err := os.Stat(cwd + "/tmp/set_test")
				g.Assert(err == nil).IsTrue()
				out, err := gpg.Decrypt(cwd + "/tmp/set_test")
				g.Assert(err == nil).IsTrue()
				g.Assert(out).Equal("This is a test")
			})
			g.It("Should create an encrypted file given another file", func() {
				v := &vault.Vaultfile{}
				v.Recipients = []string{"bob@example.com"}
				v.Save()
				os.Setenv("GNUPGHOME", cwd+"/testdata/bob")
				c, _ := Factory()

				c.Run([]string{"-f", cwd + "/testdata/set_test", cwd + "/tmp/set_test"})

				_, err := os.Stat(cwd + "/tmp/set_test")
				g.Assert(err == nil).IsTrue()
				out, err := gpg.Decrypt(cwd + "/tmp/set_test")
				g.Assert(err == nil).IsTrue()
				g.Assert(out).Equal("This is a test")
			})
		})
	})
}
