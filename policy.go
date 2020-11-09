package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/hashicorp/vault/api"
	log "github.com/sirupsen/logrus"
)

var (
	token          = os.Getenv("VAULT_TOKEN")
	vaultAddr      = os.Getenv("VAULT_ADDR")
	policyEndPoint = vaultAddr + "/v1/sys/policies/acl/"
)

// Payload is the policy content
type Payload struct {
	Name         string   `json:"name"`
	Path         string   `json:"path"`
	Capabilities []string `json:"capabilities"`
}

func putPolicy() {
	policy := Payload{
		Name:         "dev",
		Path:         "*",
		Capabilities: []string{"create", "read", "list"},
	}
	data, err := json.Marshal(policy)
	if err != nil {
		log.Error(err)
	}
	body := bytes.NewReader(data)
	req, err := http.NewRequest("PUT", policyEndPoint, body)
	if err != nil {
		log.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Vault-Token", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()
	fmt.Println("Policy is Created ...")
}
func main() {
	config := &api.Config{
		Address: vaultAddr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		log.Error(err)
		return
	}
	client.SetToken(token)
	log.Info("Client is Authenticated Successfully ..")

	putPolicy()
}
