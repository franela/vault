package vault

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func TestVaultfile(t *testing.T) {
	g := Goblin(t)
	g.Describe("Vaultfile", func() {
		g.AfterEach(func() {
			os.Remove("Vaultfile")
		})

		g.Describe("#Save", func() {
			g.It("Should work", func() {
				v := &Vaultfile{}
				v.Recipients = []string{"a@a.com"}
				v.Save()

				content, err := ioutil.ReadFile("Vaultfile")
				g.Assert(err == nil).IsTrue()

				v2 := &Vaultfile{}
				er := json.Unmarshal(content, v2)
				g.Assert(er == nil).IsTrue()

				g.Assert(v).Equal(v2)
			})
		})
	})

	g.Describe("LoadVaultfile", func() {
		g.AfterEach(func() {
			os.Remove("Vaultfile")
		})

		g.It("Should load existing Vaultfile", func() {
			v := &Vaultfile{}
			v.Recipients = []string{"a@a.com"}
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
	})
}
