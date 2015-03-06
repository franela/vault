package vault

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"testing"

	. "github.com/franela/goblin"
	"github.com/franela/vault/vault/testutils"
)

func TestVaultfile(t *testing.T) {
	g := Goblin(t)
	g.Describe("Vaultfile", func() {
		g.BeforeEach(func() {
			SetHomeDir(testutils.GetTemporaryHomeDir())
		})

		g.AfterEach(func() {
			testutils.RemoveTemporaryHomeDir(UnsetHomeDir())
		})

		g.Describe("#Save", func() {
			g.It("Should work", func() {
				v := &Vaultfile{}
				v.Recipients = []VaultRecipient{
					VaultRecipient{Fingerprint: "3B9CEC3B5069113E2ED39AC9843E01FBCE44AAAA", Name: "a@a.com"},
				}
				v.Save()

				content, err := ioutil.ReadFile(path.Join(GetHomeDir(), "Vaultfile"))
				g.Assert(err == nil).IsTrue()

				v2 := &Vaultfile{}
				er := json.Unmarshal(content, v2)
				g.Assert(er == nil).IsTrue()

				g.Assert(v).Equal(v2)
			})
		})
	})

	g.Describe("NewRecipient", func() {
		g.It("Should run", func() {
			recipient, _ := NewRecipient("3B9CEC3B5069113E2ED39AC9843E01FBCE44AAAA:a@a.com")
			g.Assert(recipient).Equal(&VaultRecipient{Fingerprint: "3B9CEC3B5069113E2ED39AC9843E01FBCE44AAAA", Name: "a@a.com"})
		})

		g.It("Should require both parameters to be present", func() {
			_, err := NewRecipient("a@a.com")
			g.Assert(err != nil).IsTrue()
			_, err = NewRecipient(":a@a.com")
			g.Assert(err != nil).IsTrue()
			_, err = NewRecipient("3B9CEC3B5069113E2ED39AC9843E01FBCE44AAAA:")
			g.Assert(err != nil).IsTrue()
			_, err = NewRecipient("")
			g.Assert(err != nil).IsTrue()
		})

		g.It("Should validate that first argument is a fingerprint", func() {
			_, err := NewRecipient("a@a.com:3B9CEC3B5069113E2ED39AC9843E01FBCE44AAAA")
			g.Assert(err != nil).IsTrue()
		})

		g.It("Should validate that first has the correct amount of characters", func() {
			_, err := NewRecipient("a@a.com:3B9CEC3B50691")
			g.Assert(err != nil).IsTrue()
		})

	})

	g.Describe("LoadVaultfile", func() {
		g.BeforeEach(func() {
			SetHomeDir(testutils.GetTemporaryHomeDir())
		})

		g.AfterEach(func() {
			testutils.RemoveTemporaryHomeDir(UnsetHomeDir())
		})

		g.It("Should load existing Vaultfile", func() {
			v := &Vaultfile{}
			v.Recipients = []VaultRecipient{
				VaultRecipient{Fingerprint: "3B9CEC3B5069113E2ED39AC9843E01FBCE44AAAA", Name: "a@a.com"},
			}
			v.Save()

			v2, err := LoadVaultfile()

			g.Assert(err == nil).IsTrue()
			g.Assert(v).Equal(v2)
		})

		g.It("Should return a new Vaultfile if it doesn't exist", func() {
			v, err := LoadVaultfile()

			g.Assert(err == nil).IsTrue()
			g.Assert(v).Equal(&Vaultfile{})
		})

		g.It("Should return an error when trying to parse the Vaultfile", func() {
			ioutil.WriteFile(path.Join(GetHomeDir(), "Vaultfile"), []byte("Not a JSON"), 0644)
			_, err := LoadVaultfile()
			g.Assert(err == nil).IsFalse()
		})

		g.It("Should traverse directories and look for a Vaultfile ", func() {
			v := &Vaultfile{}
			v.Recipients = []VaultRecipient{
				VaultRecipient{Fingerprint: "3B9CEC3B5069113E2ED39AC9843E01FBCE44AAAA", Name: "a@a.com"},
			}
			v.Save()

			os.Mkdir(path.Join(GetHomeDir(), "test"), 0777)
			SetHomeDir(path.Join(GetHomeDir(), "test"))

			v2, err := LoadVaultfile()

			g.Assert(err == nil).IsTrue()
			g.Assert(v).Equal(v2)
		})
	})
}
