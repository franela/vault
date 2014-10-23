package main

import (
	"flag"
	"github.com/mitchellh/cli"
	"log"
	"os"

	add "github.com/franela/vault/commands/add"
	get "github.com/franela/vault/commands/get"
	inita "github.com/franela/vault/commands/init"
	recipients "github.com/franela/vault/commands/recipients"
	remove "github.com/franela/vault/commands/remove"
	repair "github.com/franela/vault/commands/repair"
	set "github.com/franela/vault/commands/set"
)

func main() {
	c := initializeCli(os.Args[1:])
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
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
