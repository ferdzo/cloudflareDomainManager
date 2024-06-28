package functions

import (
	"bytes"
	"cloudflareDomainManager/secrets"
	"encoding/json"
	"fmt"
	"net/http"
)

func Update(secret *secrets.Secret, recId string, record Record) error {
	// Update function
	// PUT /zones/:zone_identifier/dns_records/:identifier

	// Create a new request
	payloadBytes, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("error marshalling data")
	}
	body := bytes.NewReader(payloadBytes)

	req, _ := http.NewRequest("PUT", "https://api.cloudflare.com/client/v4/zones/"+secret.Zone_ID+"/dns_records/"+recId, body)
	req.Header.Set("X-Auth-Email", secret.X_Auth_Email)
	req.Header.Set("X-Auth-Key", secret.X_Auth_Key)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error requesting")

	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error updating record, status code: %d", resp.StatusCode)
	}

	fmt.Println("Record successfully updated")

	return nil
}
