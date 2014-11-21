package vault

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func NewRecipient(recipient string) (*VaultRecipient, error) {
	args := strings.Split(recipient, ":")

	if len(args) != 2 {
		return nil, fmt.Errorf("\nInvalid format \"%s\"\nRecipient should be in the form of fingerprint:name\n", recipient)
	}

	if len(args[0]) == 0 || len(args[1]) == 0 {
		return nil, fmt.Errorf("\nInvalid format \"%s\"\nRecipient should be in the form of fingerprint:name\n", recipient)
	}

	recipientFingerprint := strings.Split(recipient, ":")[0]

	if hexFingerprint, err := hex.DecodeString(recipientFingerprint); err != nil {
		return nil, fmt.Errorf("\nSupplied fingerprint \"%s\" does not have the correct format\n", recipientFingerprint)
	} else {

		if len(hexFingerprint) != 16 && len(hexFingerprint) != 20 {
			return nil, fmt.Errorf("\nSupplied fingerprint \"%s\" does not have the correct size\n", hexFingerprint)
		}

	}

	recipientName := strings.Split(recipient, ":")[1]

	return &VaultRecipient{Fingerprint: recipientFingerprint, Name: recipientName}, nil
}

func LoadVaultfile() (*Vaultfile, error) {
	v := &Vaultfile{}
	content, err := ioutil.ReadFile(path.Join(GetHomeDir(), "Vaultfile"))
	if err != nil {
		return v, nil
	}
	if err := json.Unmarshal(content, v); err != nil {
		return nil, fmt.Errorf("\nCouldn't read Vaultfile.\n")
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
