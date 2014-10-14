package get

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestGet(t *testing.T) {
	g := Goblin(t)

	g.Describe("Get", func() {
		g.Describe("#Run", func() {
			g.It("Should output decrypted text of a given file")
			g.It("Should create a file with decrypted text of a given file")
		})
	})
}
