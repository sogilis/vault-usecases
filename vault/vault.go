package vault

import (
	vault "github.com/hashicorp/vault/api"
)

type VaultClient struct {
	c *vault.Client
}

//NewVaultClient creates a vault client and checks its connecatbility
func NewVaultClient(c *vault.Config) (*VaultClient, error) {
	client, err := vault.NewClient(c)
	if err != nil {
		return nil, err
	}

	_, err = client.Logical().Read("kv-v2/data/")
	if err != nil {
		return nil, err
	}

	return &VaultClient{c: client}, nil
}
