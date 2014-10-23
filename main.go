package main

import (
	"flag"
	"github.com/mitchellh/cli"
	"log"
	"os"
	"os/exec"

	add "github.com/franela/vault/commands/add"
	get "github.com/franela/vault/commands/get"
	inita "github.com/franela/vault/commands/init"
	recipients "github.com/franela/vault/commands/recipients"
	remove "github.com/franela/vault/commands/remove"
	repair "github.com/franela/vault/commands/repair"
	set "github.com/franela/vault/commands/set"
)

func main() {

	if !isGPGInstalled() {
		log.Println("Could not find GPG in your PATH. Please make sure it is installed and in your PATH.")
		os.Exit(3)
	}

	c := initializeCli(os.Args[1:])
	exitStatus, err := c.Run()
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

func initializeCli(args []string) *cli.CLI {
	initFlags := flag.NewFlagSet("init", flag.ContinueOnError)
	var verbose = initFlags.Bool("verbose", false, "Logs verbose information to stderr")
	initFlags.Parse(args)

	if *verbose {
		log.SetOutput(os.Stderr)
	} else {
		devNull, _ := os.Open(os.DevNull)
		log.SetOutput(devNull)
	}

	c := cli.NewCLI("vault", "0.0.1")
	c.Args = args
	c.Commands = map[string]cli.CommandFactory{
		"init":       inita.Factory,
		"set":        set.Factory,
		"get":        get.Factory,
		"recipients": recipients.Factory,
		"add":        add.Factory,
		"remove":     remove.Factory,
		"repair":     repair.Factory,
	}
	return c
}
