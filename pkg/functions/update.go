package functions

import (
	"bytes"
	"cloudflareDomainManager/secrets"
	"encoding/json"
	"fmt"
	"net/http"
)

func Update(secret *secrets.Secret, rec_id string, record Record) {
	// Update function
	// PUT /zones/:zone_identifier/dns_records/:identifier

	// Create a new request
	payloadBytes, err := json.Marshal(record)
	if err != nil {
		fmt.Errorf("Error marshalling data")
	}
	body := bytes.NewReader(payloadBytes)

	req, _ := http.NewRequest("PUT", "https://api.cloudflare.com/client/v4/zones/"+secret.Zone_ID+"/dns_records/"+rec_id, body)
	req.Header.Set("X-Auth-Email", secret.X_Auth_Email)
	req.Header.Set("X-Auth-Key", secret.X_Auth_Key)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("Error requesting")

	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("Error updating record, status code: %d", resp.StatusCode)
	}

	fmt.Println("Record successfully updated")

}
