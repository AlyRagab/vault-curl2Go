# vault-curl2Go
Creating Hashicorp-Vault policies using Go client

### First run Vault and set the proper VAULT_ADDR variable:
```
docker run -d --name vault --cap-add=IPC_LOCK -e 'VAULT_DEV_ROOT_TOKEN_ID=tokenroot' -e 'VAULT_DEV_LISTEN_ADDRESS=0.0.0.0:8200' -p 8200:8200 vault

export VAULT_ADDR='http://127.0.0.1:8200'

vault login
```