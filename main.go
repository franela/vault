package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"

	get "github.com/franela/vault/commands/get"
	inita "github.com/franela/vault/commands/init"
	set "github.com/franela/vault/commands/set"
)

func main() {
	c := cli.NewCLI("vault", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"init": inita.Factory,
		"set":  set.Factory,
		"get":  get.Factory,
		/*
      TODO:
        recipients
        add
        remove
		*/
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
