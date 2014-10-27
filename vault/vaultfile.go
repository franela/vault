package vault

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
)

type Vaultfile struct {
	Recipients []VaultRecipient
}

type VaultRecipient struct {
	Name        string
	Fingerprint string
}

func LoadVaultfile() (*Vaultfile, error) {
	v := &Vaultfile{}
	content, err := ioutil.ReadFile(path.Join(GetHomeDir(), "Vaultfile"))
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

	err2 := ioutil.WriteFile(path.Join(GetHomeDir(), "Vaultfile"), js, 0644)
	if err2 != nil {
		return err2
	}
	return nil
}
