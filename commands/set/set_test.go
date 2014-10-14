package set

import (
	"os"
	"testing"
	. "github.com/franela/goblin"
	"github.com/franela/vault/vault"
	"github.com/franela/vault/gpg"
)


var (
        cwd = os.Getenv("VAULTDIR")
)

func TestSet(t *testing.T) {
	g := Goblin(t)

	g.Describe("Set", func() {
		g.Describe("#Run", func() {
			g.It("Should create an encrypted file given a text", func() {
				defer func() {
					os.Remove("Vaultfile")
				}()



                                v := &vault.Vaultfile{}
                                v.Recipients = []string{"bob@example.com"}
                                v.Save()
                                os.Setenv("GNUPGHOME", cwd+"/test/bob")

                                c, _ := Factory()
                                c.Run([]string {cwd+"/tmp/set_test", "This is a test"})


                                _, err := os.Stat(cwd+"/tmp/set_test")
                                g.Assert(err == nil).IsTrue()



                                out, err := gpg.Decrypt(cwd+"/tmp/set_test")

                                g.Assert(err == nil).IsTrue()
                                g.Assert(out).Equal("This is a test")


			})
		})
	})
}
