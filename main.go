package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"

	add "github.com/franela/vault/commands/add"
	get "github.com/franela/vault/commands/get"
	inita "github.com/franela/vault/commands/init"
	recipients "github.com/franela/vault/commands/recipients"
	remove "github.com/franela/vault/commands/remove"
	set "github.com/franela/vault/commands/set"
)

func main() {
	c := cli.NewCLI("vault", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"init":       inita.Factory,
		"set":        set.Factory,
		"get":        get.Factory,
		"recipients": recipients.Factory,
		"add":        add.Factory,
		"remove":     remove.Factory,
		/*
				   TODO:
		          repair
		*/
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
