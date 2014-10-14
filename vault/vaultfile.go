package vault

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Vaultfile struct {
	Recipients []string
}

func LoadVaultfile() (*Vaultfile, error) {
	wd, err := GetHomeDir()
	if err != nil {
		return nil, err
	}

	v := &Vaultfile{}
	content, err := ioutil.ReadFile(wd + "/Vaultfile")
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
	wd, err := GetHomeDir()
	if err != nil {
		return err
	}

	js, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err2 := ioutil.WriteFile(wd+"/Vaultfile", js, 0644)
	if err2 != nil {
		return err2
	}
	return nil
}
