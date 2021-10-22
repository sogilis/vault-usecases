package vault_test

import (
	"os"
	"testing"

	"vault-usecase/vault"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rootClient *vault.VaultClient

func beforeAll(t *testing.T) {
	vaultRootToken := os.Getenv("VAULT_ROOT_TK")
	require.NotNil(t, vaultRootToken, "cannot retrienve mandatory root token")

	var err error
	rootClient, err = vault.NewVaultClientFromToken(nil, vaultRootToken)
	require.Nil(t, err, "Got error while connecting the root vault client")
}

func TestValidNewUserAuth(t *testing.T) {
	//Given
	beforeAll(t)

	//When
	err := rootClient.CreateNewUserAuth("toto", "007")
	require.Nil(t, err)

	//Then
	list, err := rootClient.ListUsers()

	require.Nil(t, err)
	assert.Equal(t, []string{"toto"}, list)

	//After
	err = rootClient.DeleteUser("toto")
	assert.Nil(t, err, "failed to deleted user tata at end of test")
}
