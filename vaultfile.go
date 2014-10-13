package main

import (
  "io/ioutil"
  "encoding/json"
  "fmt"
)

type Vaultfile struct {
  Recipients []string
}

func LoadVaultfile() (*Vaultfile, error) {
  v := &Vaultfile{}
  content, err := ioutil.ReadFile("Vaultfile")
  if err != nil {
    return v, nil
  }
  if err := json.Unmarshal(content, v); err != nil {
    return nil, fmt.Errorf("Couldn't read Vaultfile.")
  } else {
    return v, nil
  }
}

func (v Vaultfile) Save() error {
  js, err := json.Marshal(v)
  if err != nil {
    return err
  }

  err2 := ioutil.WriteFile("Vaultfile", js, 0644)
  if err2 != nil {
    return err2
  }
  return nil
}
