package vault

import (
	"os"
	"path"
	"testing"

	. "github.com/franela/goblin"
)

func TestVault(t *testing.T) {
	g := Goblin(t)

	g.Describe("Vault", func() {
		g.Describe("#GetHomeDir", func() {
			g.It("Should return current working dir by default", func() {
				wd, _ := os.Getwd()
				// This is a hack for the test to run because golang returns the path of the test and not the project
				wd = path.Dir(wd)

				d, err := GetHomeDir()

				g.Assert(err == nil).IsTrue()
				g.Assert(d).Equal(wd)
			})

			g.It("Should return VAULTDIR environment variable if it is defined", func() {
				os.Setenv("VAULTDIR", "/tmp")

				d, err := GetHomeDir()

				g.Assert(err == nil).IsTrue()
				g.Assert(d).Equal("/tmp")
			})
		})
	})
}
