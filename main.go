package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/mitchellh/cli"

	add "github.com/franela/vault/commands/add"
	cmdimport "github.com/franela/vault/commands/cmdimport"
	get "github.com/franela/vault/commands/get"
	inita "github.com/franela/vault/commands/init"
	recipients "github.com/franela/vault/commands/recipients"
	remove "github.com/franela/vault/commands/remove"
	repair "github.com/franela/vault/commands/repair"
	set "github.com/franela/vault/commands/set"
	"github.com/franela/vault/validator"
)

var commandValidator validator.Validator

func init() {
	commandValidator = &validator.CommandValidator{}
}

func main() {

	if !isGPGInstalled() {
		log.Println("Could not find GPG in your PATH. Please make sure it is installed and in your PATH. You may download GPG from https://www.gnupg.org/download/ ")
		os.Exit(3)
	}

	c := initializeCli(os.Args[1:])

	exitStatus, err := runCommand(c)

	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}

func isGPGInstalled() bool {
	_, err := exec.LookPath("gpg")

	if err != nil {
		return false
	}

	return true
}

func runCommand(cli *cli.CLI) (int, error) {

	exitStatus, err := cli.Run()

	if exitStatus != 0 {
		commandValidator.Validate()
	}

	return exitStatus, err
}

func initializeCli(args []string) *cli.CLI {
	c := cli.NewCLI("vault", "0.0.1")
	c.Args = args

	devNull, _ := os.Open(os.DevNull)
	log.SetOutput(devNull)

	if !c.IsVersion() && !c.IsHelp() {
		initFlags := flag.NewFlagSet("verbose", flag.ContinueOnError)
		var verbose = initFlags.Bool("verbose", false, "Logs verbose information to stderr")
		initFlags.Parse(args)
		if *verbose {
			log.SetOutput(os.Stderr)
		}
	}

	c.Commands = map[string]cli.CommandFactory{
		"init":       inita.Factory,
		"set":        set.Factory,
		"get":        get.Factory,
		"recipients": recipients.Factory,
		"add":        add.Factory,
		"remove":     remove.Factory,
		"repair":     repair.Factory,
		"import":     cmdimport.Factory,
	}

	return c
}
