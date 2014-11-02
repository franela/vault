package validator

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
	"github.com/franela/vault/ui"
	"github.com/franela/vault/vault"
	"github.com/franela/vault/vault/testutils"
)

func TestGPG(t *testing.T) {
	g := Goblin(t)

	g.Describe("Validator", func() {
		g.Describe("#Validate", func() {
			var tempdir string
			g.BeforeEach(func() {
				vault.SetHomeDir(testutils.GetTemporaryHomeDir())
				ui.DEBUG = true
			})

			g.AfterEach(func() {
				testutils.RemoveTemporaryHomeDir(tempdir)
				ui.DEBUG = true
				os.Setenv("GNUPGHOME", "")
			})
			g.It("Should print a message when no private keys are found", func() {
				os.Setenv("GNUPGHOME", vault.GetHomeDir())

				validator := CommandValidator{}
				validator.Validate()

				g.Assert(ui.GetOutput()).Equal(noKeyMessage)
			})
			g.It("Should  not print a message when no private keys are found", func() {
				testutils.SetTestGPGHome("bob")
				validator := CommandValidator{}
				validator.Validate()

				g.Assert(ui.GetOutput()).Equal("")

			})
		})
	})
}
