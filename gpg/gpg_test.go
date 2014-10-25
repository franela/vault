package gpg

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func TestGPG(t *testing.T) {
	g := Goblin(t)

	g.Describe("GPG", func() {
		g.Describe("#getGPGHomeDir", func() {
			g.It("Should return GNUPGHOME env if it is set", func() {
				os.Setenv("GNUPGHOME", "/a/test/path")

				args := getGPGHomeDir()

				g.Assert(args).Equal([]string{"--homedir", "/a/test/path"})
			})
			g.It("Should return empty args if GNUPGHOME env is not set", func() {
				os.Setenv("GNUPGHOME", "")

				args := getGPGHomeDir()

				g.Assert(args).Equal([]string{})
			})
		})
	})
}
