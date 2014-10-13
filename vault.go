package main

import (
  "os"
  "log"

  "github.com/mitchellh/cli"
)

func main() {
  c := cli.NewCLI("vault", "0.0.1")
  c.Args = os.Args[1:]
  c.Commands = map[string]cli.CommandFactory{
    "init": initCommandFactory,
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
