package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"syscall"
	"testing"

	. "github.com/franela/goblin"
	"github.com/franela/vault/vault"
	"github.com/franela/vault/vault/testutils"
)

func TestMain(t *testing.T) {
	g := Goblin(t)

	g.Describe("Main", func() {
		g.Describe("#isGPGInstalled", func() {
			g.It("Should return false if gpg is not in the PATH", func() {
				prevPath := os.Getenv("PATH")

				defer func() {
					os.Setenv("PATH", prevPath)
				}()
				os.Setenv("PATH", "")

				g.Assert(isGPGInstalled()).IsFalse()
			})
			g.It("Should return true if gpg is in the PATH", func() {
				g.Assert(isGPGInstalled()).IsTrue()
			})
		})

		g.Describe("#Run", func() {
			g.It("Should log to dev/null by default if no verbose flag is set ", func() {
				dir := testutils.GetTemporaryHomeDir()
				filePath := path.Join(dir, "stderr")
				logFile, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0644)
				syscall.Dup2(int(logFile.Fd()), 2)

				initializeCli([]string{})
				log.Println("Test")

				defer func() {
					logFile.Close()
				}()

				logFile, _ = os.Open(filePath)
				fileContent, _ := ioutil.ReadAll(logFile)

				g.Assert(len(fileContent)).Equal(0)
			})

			g.It("Should log to stderr if verbose flag is set ", func() {
				dir := testutils.GetTemporaryHomeDir()
				filePath := path.Join(dir, "stderr")
				logFile, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0644)
				syscall.Dup2(int(logFile.Fd()), 2)

				initializeCli([]string{"-verbose"})
				log.Println("Test")

				logFile.Close()

				defer func() {
					logFile.Close()
				}()

				logFile, _ = os.Open(filePath)
				fileContent, _ := ioutil.ReadAll(logFile)

				g.Assert(len(fileContent)).Equal(25)
			})

			g.It("Should run validator if command returns with an error", func() {
				testutils.SetTestGPGHome("nokey")
				vault.SetHomeDir(testutils.GetTemporaryHomeDir())

				called := false

				mockValidator := &testutils.MockValidator{
					ValidateMock: func() { called = true },
				}
				commandValidator = mockValidator

				runCommand(initializeCli([]string{"set"}))

				g.Assert(called).IsTrue()

			})
		})
	})
}
