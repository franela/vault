package cmdimport

import (
	. "github.com/franela/goblin"
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/vault"
	"github.com/franela/vault/vault/testutils"
	"strings"
	"testing"
)

func TestImport(t *testing.T) {
	g := Goblin(t)

	g.Describe("Import", func() {
		g.Describe("#Import", func() {
			g.BeforeEach(func() {
				vault.SetHomeDir(testutils.GetTemporaryHomeDir())
			})

			g.AfterEach(func() {
				testutils.RemoveTemporaryHomeDir(vault.UnsetHomeDir())
			})

			g.It("Should work", func() {
				v := vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{
					vault.VaultRecipient{Fingerprint: "2B13EC3B5769013E2ED29AC9643E01FBCE44E394", Name: "bob@example.com"},
					vault.VaultRecipient{Fingerprint: "3B6094CF22AEC3F24274F389F8987FE0142E59FA", Name: "marcos@example.com"},
				}
				v.Save()

				c, _ := Factory()

				mockExecutor := &testutils.MockCMDExecutor{}
				gpg.SetExecutor(mockExecutor)

				code := c.Run([]string{})

				g.Assert(code).Equal(0)
				g.Assert("gpg --batch --yes --recv-keys 2B13EC3B5769013E2ED29AC9643E01FBCE44E394 3B6094CF22AEC3F24274F389F8987FE0142E59FA").Equal(strings.Join(mockExecutor.Arguments, " "))
			})
			g.It("Should allow to specify a keyserver", func() {
				v := vault.Vaultfile{}
				v.Recipients = []vault.VaultRecipient{
					vault.VaultRecipient{Fingerprint: "2B13EC3B5769013E2ED29AC9643E01FBCE44E394", Name: "bob@example.com"},
					vault.VaultRecipient{Fingerprint: "3B6094CF22AEC3F24274F389F8987FE0142E59FA", Name: "marcos@example.com"},
				}
				v.Save()

				c, _ := Factory()

				mockExecutor := &testutils.MockCMDExecutor{}
				gpg.SetExecutor(mockExecutor)

				code := c.Run([]string{"--keyserver", "hkp://test.keyserver.com"})

				g.Assert(code).Equal(0)
				g.Assert("gpg --batch --yes --keyserver hkp://test.keyserver.com --recv-keys 2B13EC3B5769013E2ED29AC9643E01FBCE44E394 3B6094CF22AEC3F24274F389F8987FE0142E59FA").Equal(strings.Join(mockExecutor.Arguments, " "))
			})
		})
	})
}
