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
func beforeAll(t *testing.T) {
	vaultRootToken := os.Getenv("VAULT_ROOT_TK")
	require.NotNil(t, vaultRootToken, "cannot retrieve mandatory root token")

	var err error
	rootClient, err = root.NewRootClient(nil, vaultRootToken)
	require.Nil(t, err, "Got error while connecting the root vault client")
}

func TestValidNewUserAuth(t *testing.T) {
	beforeAll(t)
	t.Run("Enable, then disable, userpass auth with root token", enableAndDisableUserpassAuthTest)
}

func enableAndDisableUserpassAuthTest(t *testing.T) {
	//Given

	//When
	err := rootClient.EnableUserpassAuth()
	require.Nil(t, err)

	err = rootClient.DisableUserpassAuth()
	assert.Nil(t, err)
}
