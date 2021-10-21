package vault

import (
	vault "github.com/hashicorp/vault/api"
)

// Login checks that vault client can login
//TODO: refactor into NewVaultClient
func Login() error {
	//TODO: read config from env var or file
	client, err := vault.NewClient(nil)
	_, err = client.Logical().Read("kv-v2/data/")
	return err
}
