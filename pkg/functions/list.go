package functions

import (
	"cloudflareDomainManager/secrets"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Meta struct {
	AutoAdded           bool `json:"auto_added"`
	ManagedByApps       bool `json:"managed_by_apps"`
	ManagedByArgoTunnel bool `json:"managed_by_argo_tunnel"`
}

type Record1 struct {
	ID         string   `json:"id"`
	ZoneID     string   `json:"zone_id"`
	ZoneName   string   `json:"zone_name"`
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Content    string   `json:"content"`
	Proxiable  bool     `json:"proxiable"`
	Proxied    bool     `json:"proxied"`
	TTL        int      `json:"ttl"`
	Locked     bool     `json:"locked"`
	Meta       Meta     `json:"meta"`
	Comment    *string  `json:"comment"`
	Tags       []string `json:"tags"`
	CreatedOn  string   `json:"created_on"`
	ModifiedOn string   `json:"modified_on"`
}

type Response struct {
	Result []Record1 `json:"result"`
}

func List(secret *secrets.Secret) error {

	req, err := http.NewRequest("GET", "https://api.cloudflare.com/client/v4/zones/"+secret.Zone_ID+"/dns_records", nil)
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
		fmt.Printf("failed to fetch records: %s", string(bodyBytes))
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	var response Response
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	// Print the parsed records
	for _, record := range response.Result {
		fmt.Printf("ID: %s, Name: %s, Type: %s, Content: %s\n", record.ID, record.Name, record.Type, record.Content)
	}

	return nil

}
