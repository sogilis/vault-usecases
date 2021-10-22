package vault_test

import (
	"testing"

	"vault-usecase/vault"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVaultClient(t *testing.T) {
	//Given
	_, err := vault.NewVaultClient(nil)

	//Then
	assert.Nil(t, err)
}

func TestValidNewUserAuth(t *testing.T) {
	//Given
	vaultClient, err := vault.NewVaultClient(nil)
	require.Nil(t, err)

	//When
	err = vaultClient.CreateNewUserAuth("toto", "007")
	require.Nil(t, err)

	//Then
	list, err := vaultClient.ListUsers()

	require.Nil(t, err)
	assert.Equal(t, []string{"toto"}, list)

	//After
	err = vaultClient.DeleteUser("toto")
	assert.Nil(t, err, "failed to deleted user tata at end of test")
}
