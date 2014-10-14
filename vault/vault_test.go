package vault

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func TestVault(t *testing.T) {
	g := Goblin(t)

	g.Describe("Vault", func() {
		g.Describe("#GetHomeDir", func() {
			g.It("Should return current working dir by default", func() {
        env := os.Getenv("VAULTDIR")
        defer func() {
          os.Setenv("VAULTDIR", env)
        }()
        os.Setenv("VAULTDIR", "")

				wd, _ := os.Getwd()

				d, err := GetHomeDir()

				g.Assert(err == nil).IsTrue()
				g.Assert(d).Equal(wd)
			})

			g.It("Should return VAULTDIR environment variable if it is defined", func() {
        defer func() {
          os.Setenv("VAULTDIR", "")
        }()
				os.Setenv("VAULTDIR", "/tmp")

				d, err := GetHomeDir()

				g.Assert(err == nil).IsTrue()
				g.Assert(d).Equal("/tmp")
			})
		})
	})
}
