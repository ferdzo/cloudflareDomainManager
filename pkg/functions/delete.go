package functions

import (
	"cloudflareDomainManager/secrets"
	"fmt"
	"io"
	"net/http"
)

func Delete(secret *secrets.Secret, record_id string) error {
	req, err := http.NewRequest("DELETE", "https://api.cloudflare.com/client/v4/zones/"+secret.Zone_ID+"/dns_records/"+record_id, nil)
	if err != nil {
		return fmt.Errorf("Error creating request")
	}
	req.Header.Set("X-Auth-Email", secret.X_Auth_Email)
	req.Header.Set("X-Auth-Key", secret.X_Auth_Key)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Request failed")
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Printf("failed to delete record: %s", string(bodyBytes))
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}
	if resp.StatusCode == http.StatusNoContent {
		fmt.Println("Record deleted successfully")
	}

	return nil
}
