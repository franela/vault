package vault

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func NewRecipient(recipient string) VaultRecipient {
	recipientFingerprint := strings.Split(recipient, ":")[0]
	recipientName := strings.Split(recipient, ":")[1]

	return VaultRecipient{Fingerprint: recipientFingerprint, Name: recipientName}
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

type VaultRecipient struct {
	Name        string
	Fingerprint string
}

func (v VaultRecipient) ToString() string {
	return v.Fingerprint + ":" + v.Name
}

type Vaultfile struct {
	Recipients []VaultRecipient
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
