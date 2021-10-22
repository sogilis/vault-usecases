package root

import (
	"fmt"

	vault "github.com/hashicorp/vault/api"
)

type RootClient interface {
	CreateAuthUserPolicy() error
	EnableUserpassAuth() error
}

type RootVaultClient struct {
	c *vault.Client
}

//NewRootClient Creates a root client for vault and check it connecatbility
func NewRootClient(c *vault.Config, token string) (*RootVaultClient, error) {
	client, err := vault.NewClient(c)
	if err != nil {
		return nil, err
	}

	client.SetToken(token)

	_, err = client.Logical().Read("kv-v2/data/")
	if err != nil {
		return nil, err
	}

	return &RootVaultClient{c: client}, nil
}

func (rvc *RootVaultClient) CreateAuthUserPolicy() error {
	return nil
}

func (rvc *RootVaultClient) EnableUserpassAuth() error {
	params := map[string]interface{}{
		"path":        "userpass",
		"description": "User/Password auth",
		"type":        "userpass",
	}

	_, err := rvc.c.Logical().Write("/sys/auth/userpass", params)
	if err != nil {
		fmt.Printf("ERROR: cannot create user auth in vault : %v", err)
		return err
	}
	return nil
}
