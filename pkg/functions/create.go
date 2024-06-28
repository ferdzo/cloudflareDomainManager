package functions

import (
	"bytes"
	"cloudflareDomainManager/secrets"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Record struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
	Proxied bool   `json:"proxied"`
}

func Create(secrets *secrets.Secret, record *Record) error {

	// Create a new record
	// POST /zones/:zone_identifier/dns_records

	data := record
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Error marshalling data")
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.cloudflare.com/client/v4/zones/"+secrets.Zone_ID+"/dns_records", body)

	if err != nil {
		// handle err
	}
	req.Header.Set("X-Auth-Email", secrets.X_Auth_Email)
	req.Header.Set("X-Auth-Key", secrets.X_Auth_Key)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error requesting")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("failed to create record: %s", string(bodyBytes))
		return fmt.Errorf("error creating record, status code: %d", resp.StatusCode)
	}

	log.Println("Record successfully created")

	return nil
}
