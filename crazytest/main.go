// +build !windows

package main


import (
  "os"
  "fmt"
)

func main() {
  if len(os.Getenv("GNUPGHOME")) > 0 {
    fmt.Println(os.Getenv("GNUPGHOME"))
  } else {
    fmt.Println("~/.gnupg")
  }
}
