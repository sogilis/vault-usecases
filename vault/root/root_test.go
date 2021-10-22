package root_test

import (
	"os"
	"testing"

	"vault-usecase/vault/root"

	"github.com/stretchr/testify/require"
)

var rootClient root.RootClient

func beforeAll(t *testing.T) {
	vaultRootToken := os.Getenv("VAULT_ROOT_TK")
	require.NotNil(t, vaultRootToken, "cannot retrienve mandatory root token")

	var err error
	rootClient, err = root.NewRootClient(nil, vaultRootToken)
	require.Nil(t, err, "Got error while connecting the root vault client")
}

func TestValidNewUserAuth(t *testing.T) {
	//Given
	beforeAll(t)

	//When
	err := rootClient.EnableUserpassAuth()
	require.Nil(t, err)

	err = rootClient.CreateAuthUserPolicy()
	require.Nil(t, err)
}
