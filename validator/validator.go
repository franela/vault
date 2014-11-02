package validator

import (
	"github.com/franela/vault/gpg"
	"github.com/franela/vault/ui"
)

var noKeyMessage = `
Seems like you haven't set a private gpg key yet, use gpg --gen-key to generate one before continuing

`

type Validator interface {
	Validate()
}

type CommandValidator struct {
}

func (*CommandValidator) Validate() {

	secretKeys, _ := gpg.GetSecretKeysCount()

	if secretKeys == 0 {
		ui.Printf(noKeyMessage)
	}
}
