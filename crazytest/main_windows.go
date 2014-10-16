package main

import (
  "github.com/luisiturrios/gowin"
  "fmt"
)

func main() {
  val, err := gowin.GetReg("HKCU", `Software\GNU\GnuPG`, "HomeDir")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(val)
}
