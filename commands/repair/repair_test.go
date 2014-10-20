package repair

import (
	. "github.com/franela/goblin"
	"github.com/franela/vault/commands/get"
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/vault"
	"github.com/franela/vault/vault/testutils"
	"path"
	"testing"
)

func TestRepair(t *testing.T) {
	g := Goblin(t)

	g.Describe("Repair", func() {
		g.Describe("#Run", func() {
			g.BeforeEach(func() {
				vault.SetHomeDir(testutils.GetTemporaryHomeDir())
			})

			g.AfterEach(func() {
				testutils.RemoveTemporaryHomeDir(vault.UnsetHomeDir())
			})

			g.It("Should re-encrypt all the files with the recipients in the Vaultfile", func() {
				testutils.SetTestGPGHome("bob")

				v := &vault.Vaultfile{}
				v.Recipients = []string{"bob@example.com"}
				v.Save()

				encryptedFilePath := path.Join(vault.GetHomeDir(), "repair_test")
				gpg.Encrypt(encryptedFilePath, "This is a test", v.Recipients)

				encryptedFilePath2 := path.Join(vault.GetHomeDir(), "foo", "repair_test2")
				gpg.Encrypt(encryptedFilePath2, "This is a test", v.Recipients)

				v.Recipients = append(v.Recipients, "alice@example.com")
				v.Save()
				c, _ := Factory()
				c.Run([]string{})

				getcmd, _ := get.Factory()

				testutils.SetTestGPGHome("bob")
				code := getcmd.Run([]string{"repair_test"})
				g.Assert(code).Equal(0)

				testutils.SetTestGPGHome("alice")
				code = getcmd.Run([]string{"repair_test"})
				g.Assert(code).Equal(0)

				testutils.SetTestGPGHome("bob")
				code = getcmd.Run([]string{path.Join("foo", "repair_test2")})
				g.Assert(code).Equal(0)

				testutils.SetTestGPGHome("alice")
				code = getcmd.Run([]string{path.Join("foo", "repair_test2")})
				g.Assert(code).Equal(0)
			})
		})
	})
}
