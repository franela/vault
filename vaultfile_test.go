package main

import (
  "testing"
  "io/ioutil"
  "os"
  "encoding/json"
  "reflect"
)

func TestSave(t *testing.T) {
  defer func() {
    os.Remove("Vaultfile")
  }()

  v := &Vaultfile{}
  v.Recipients = []string {"a@a.com"}
  v.Save()

  content, err := ioutil.ReadFile("Vaultfile")
  if err != nil {
    t.Error("Vaultfile should exist")
  }

  v2 := &Vaultfile{}
  er := json.Unmarshal(content, v2)
  if er != nil {
    t.Error("Vaultfile is not a valid")
  }

  if !reflect.DeepEqual(v, v2) {
    t.Errorf("Vaultfile recipients are not the expected ones")
  }
}

func TestLoadExisting(t *testing.T) {
  defer func() {
    os.Remove("Vaultfile")
  }()

  v := &Vaultfile{}
  v.Recipients = []string {"a@a.com"}
  v.Save()

  v2, err := LoadVaultfile()

  if err != nil {
    t.Error(err)
  }

  if !reflect.DeepEqual(v, v2) {
    t.Errorf("Vaultfile recipients are not the expected ones")
  }
}

func TestLoadUnexisting(t *testing.T) {
  defer func() {
    os.Remove("Vaultfile")
  }()

  v, err := LoadVaultfile()

  if err != nil {
    t.Error(err)
  }

  if !reflect.DeepEqual(v, &Vaultfile{}) {
    t.Errorf("Vaultfile recipients are not the expected ones")
  }
}
