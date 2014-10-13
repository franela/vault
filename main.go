package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"

	inita "github.com/franela/vault/commands/init"
)

func main() {
	c := cli.NewCLI("vault", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"init": inita.Factory,
		/*
		   "set": addCommandFactory,
		   "get": showCommandFactory,
		   "del": delCommandFactory,
		   "list": lsCommandFactory
		*/
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
