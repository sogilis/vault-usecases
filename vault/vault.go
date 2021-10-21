package vault

import (
	"fmt"

	vault "github.com/hashicorp/vault/api"
)

func Login() {
	config := vault.DefaultConfig()
	fmt.Printf("Vault config = %v", config)
	// vault.NewClient()
}
