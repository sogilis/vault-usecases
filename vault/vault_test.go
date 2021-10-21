package vault_test

import (
	"testing"

	"vault-usecase/vault"

	"github.com/stretchr/testify/assert"
)

func TestVaultLogin(t *testing.T) {
	err := vault.Login()
	assert.Nil(t, err)
}
