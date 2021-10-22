package root_test

import (
	"os"
	"testing"

	"vault-usecase/vault/root"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rootClient root.RootClient

//beforeAll retrieves and sets the root vault client
func setRootClientWithRootToken(t *testing.T) {
	vaultRootToken := os.Getenv("VAULT_ROOT_TK")
	require.NotNil(t, vaultRootToken, "cannot retrieve mandatory root token")

	var err error
	rootClient, err = root.NewRootClient(nil, vaultRootToken)
	require.Nil(t, err, "Got error while connecting the root vault client")
}

func TestValidNewUserAuth(t *testing.T) {
	setRootClientWithRootToken(t)

	t.Run("Enable, then disable, userpass auth with root token", enableAndDisableUserpassAuthWithRootTokenTest)
	t.Run("Create and deletes usergen policy with root token", createAndDeleteUsegenPolicyTest)
}

func enableAndDisableUserpassAuthWithRootTokenTest(t *testing.T) {
	//Given

	//When
	err := rootClient.EnableUserpassAuth()
	require.Nil(t, err)

	err = rootClient.DisableUserpassAuth()
	assert.Nil(t, err)
}

func createAndDeleteUsegenPolicyTest(t *testing.T) {
	//Given

	//When
	err := rootClient.CreateUserGeneratorPolicy()
	require.Nil(t, err)

	err = rootClient.DeleteUserGeneratorPolicy()
	assert.Nil(t, err)
}
