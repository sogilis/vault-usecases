package main

import (
	"fmt"
	"vault-usecase/vault"
)

func main() {
	err := vault.Login()
	if err != nil {
		fmt.Errorf("unable to initialize Vault client: %w", err)
	}
}
