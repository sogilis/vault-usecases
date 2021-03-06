package root

import (
	"encoding/base64"
	"fmt"

	vault "github.com/hashicorp/vault/api"
)

type RootClient interface {
	CreateUserGeneratorPolicy() error
	DeleteUserGeneratorPolicy() error
	DisableUserpassAuth() error
	EnableUserpassAuth() error
	GenerateUsergenToken() (string, error)
	RevokeUsergenToken(string) error
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

func (rvc *RootVaultClient) CreateUserGeneratorPolicy() error {
	userGenPolicy := `# Create userpass auth
path "auth/userpass/users" {
    capabilities = ["create"]
}`
	encodedPolicy := base64.StdEncoding.EncodeToString([]byte(userGenPolicy))

	params := map[string]interface{}{
		"name":   "usergen",
		"policy": encodedPolicy,
	}

	_, err := rvc.c.Logical().Write("/sys/policy/usergen", params)
	if err != nil {
		fmt.Printf("ERROR: cannot write usergen policy : %v", err)
		return err
	}
	return nil
}

func (rvc *RootVaultClient) DeleteUserGeneratorPolicy() error {
	_, err := rvc.c.Logical().Delete("/sys/policy/usergen")
	if err != nil {
		fmt.Printf("ERROR: cannot delete usergen policy : %v", err)
		return err
	}
	return nil
}

func (rvc *RootVaultClient) DisableUserpassAuth() error {
	_, err := rvc.c.Logical().Delete("/sys/auth/userpass")
	if err != nil {
		fmt.Printf("ERROR: cannot disable userpass auth : %v", err)
		return err
	}
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
		fmt.Printf("ERROR: cannot enable userpass auth : %v", err)
		return err
	}
	return nil
}

func (rvc *RootVaultClient) GenerateUsergenToken() (string, error) {
	params := map[string]interface{}{
		"role_name":    "usergen",
		"policies":     []string{"usergen"},
		"type":         "service",
		"display_name": "usergen",
	}

	secret, err := rvc.c.Logical().Write("/auth/token/create", params)
	if err != nil {
		fmt.Printf("ERROR: cannot create usergen token : %v", err)
		return "", err
	}
	return extractToken(secret)
}

func (rvc *RootVaultClient) RevokeUsergenToken(token string) error {
	params := map[string]interface{}{
		"token": token,
	}
	_, err := rvc.c.Logical().Write("/auth/token/revoke", params)
	if err != nil {
		fmt.Printf("ERROR: cannot revoke token %v : %v", token, err)
		return err
	}
	return nil
}

func extractToken(secret *vault.Secret) (string, error) {
	if secret.Auth == nil {
		return "", fmt.Errorf("No Auth provided in response")
	}
	if secret.Auth.ClientToken == "" {
		return "", fmt.Errorf("No client token provided")
	}

	return secret.Auth.ClientToken, nil
}
