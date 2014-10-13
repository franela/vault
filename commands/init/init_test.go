package init

import (
	"github.com/franela/vault/vault"
	"os"
	"reflect"
	"testing"
)

func TestVaultfileCreationWithRecipients(t *testing.T) {
	defer func() {
		os.Remove("Vaultfile")
	}()

	desiredRecipients := []string{"a@a.com", "b@b.com"}
	c, _ := Factory()
	exitCode := c.Run(desiredRecipients)

	v, err := vault.LoadVaultfile()

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(v.Recipients, desiredRecipients) {
		t.Errorf("Vaultfile recipients are not the expected ones")
	}

	if exitCode != 0 {
		t.Error("Expected exist code 0")
	}
}
