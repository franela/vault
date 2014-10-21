package main

import (
	. "github.com/franela/goblin"
	"github.com/franela/vault/vault/testutils"
	"io/ioutil"
	"log"
	"os"
	"path"
	"syscall"
	"testing"
)

func TestMain(t *testing.T) {
	g := Goblin(t)
	g.Describe("Main", func() {
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
		})
	})
}
