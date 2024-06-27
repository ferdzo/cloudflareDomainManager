package functions

import (
	"cloudflareDomainManager/secrets"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Show(secrets *secrets.Secret) string {

	req, err := http.NewRequest("GET", "https://api.cloudflare.com/client/v4/zones/"+secrets.Zone_ID+"/dns_records/export", nil)
	if err != nil {
		log.Fatalf("Error requesting")
	}
	req.Header.Set("X-Auth-Email", secrets.X_Auth_Email)
	req.Header.Set("X-Auth-Key", secrets.X_Auth_Key)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stdout, "FATAL ERROR")
	}
	io.Copy(os.Stdout, resp.Body) // this line.

	bodyBytes, _ := io.ReadAll(resp.Body)
	return string(bodyBytes)
}
