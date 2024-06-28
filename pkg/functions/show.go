package functions

import (
	"cloudflareDomainManager/secrets"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Show(secrets *secrets.Secret) (string, error) {

	req, err := http.NewRequest("GET", "https://api.cloudflare.com/client/v4/zones/"+secrets.Zone_ID+"/dns_records/export", nil)
	if err != nil {
		return "", fmt.Errorf("Error creating request")
	}

	req.Header.Set("X-Auth-Email", secrets.X_Auth_Email)
	req.Header.Set("X-Auth-Key", secrets.X_Auth_Key)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Request failed")
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("failed to fetch records: %s", string(bodyBytes))
		return "", fmt.Errorf("error fetching records, status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}
	return string(bodyBytes), nil
}
