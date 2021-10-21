package vault_test

import (
	"testing"

	"vault-usecase/vault"

	"github.com/stretchr/testify/assert"
)

func TestNewVaultClient(t *testing.T) {
	_, err := vault.NewVaultClient(nil)
	assert.Nil(t, err)
}
