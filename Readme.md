# Vault use cases
Gathers ready to use use cases around vault solution

## Prerequisite
* Golang >= 1.16.6
* Vault >= v1.8.4


```
go build
./vault-usecase
```

## Vault admin commands
This section describes vault commands that are intended to be run by an administrator user.

### Setup
Starts vault in dev mode.

```
vault server --dev
```

Enable userpass authentication feature of vault.
```
vault auth enable userpass
```

**TODO:** add custom policy

List all users

```
vault list /auth/userpass/users/
```

### Testing vault locally
Set minimal configuration in local env.

```
export VAULT_ADDR='http://127.0.0.1:8200'
export VAULT_TOKEN="s.LdrzTRqiI7CIZpcSCCSTWpI4"
```

## Resource
* [Vault homepage](https://www.hashicorp.com/products/vault)
