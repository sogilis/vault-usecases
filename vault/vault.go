package vault

import (
	"fmt"

	vault "github.com/hashicorp/vault/api"
)

const DEFAULT_VAULT_POLICY = "admins"

type VaultClient struct {
	c *vault.Client
}

//NewVaultClient creates a vault client and checks its connecatbility
func NewVaultClientFromToken(c *vault.Config, token string) (*VaultClient, error) {
	client, err := vault.NewClient(c)
	if err != nil {
		return nil, err
	}

	client.SetToken(token)

	_, err = client.Logical().Read("kv-v2/data/")
	if err != nil {
		return nil, err
	}

	return &VaultClient{c: client}, nil
}

// CreateNewUserAuth creates a new user authentication in vault.
// -userID will be the user identifier
// -pass will be the password use to authenticate to vault
// NOTE: the default policy will be applied to the vault user
// See DEFAULT_VAULT_POLICY
func (vc *VaultClient) CreateNewUserAuth(userID, pass string) error {
	params := map[string]interface{}{
		"password": pass,
		"policies": DEFAULT_VAULT_POLICY,
	}

	path := fmt.Sprintf("auth/userpass/users/%v", userID)
	_, err := vc.c.Logical().Write(path, params)
	if err != nil {
		fmt.Printf("ERROR: cannot create user auth in vault : %v", err)
		return err
	}

	return nil
}

func (vc *VaultClient) ListUsers() ([]string, error) {
	userList, err := vc.c.Logical().List("auth/userpass/users/")
	if err != nil {
		fmt.Printf("ERROR: cannot read user list in vault : %v", err)
		return nil, err
	}

	var data = userList.Data["keys"].([]interface{})
	return parseSecret(data), nil
}

func (vc *VaultClient) DeleteUser(userID string) error {
	path := fmt.Sprintf("auth/userpass/users/%v", userID)
	_, err := vc.c.Logical().Delete(path)
	if err != nil {
		fmt.Printf("ERROR: cannot deleted user %v in vault : %v", userID, err)
		return err
	}
	return nil
}

func parseSecret(data []interface{}) []string {
	var stringified = make([]string, 0, 2)
	for _, i := range data {
		var s = i.(string)
		t := s
		stringified = append(stringified, t)
	}
	return stringified
}
