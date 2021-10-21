package main

import (
	"fmt"
	"vault-usecase/vault"
)

func main() {
	//TODO: provide better config
	_, err := vault.NewVaultClient(nil)
	if err != nil {
		fmt.Errorf("unable to initialize Vault client: %w", err)
	}
}
