package functions

import (
	"bytes"
	"cloudflareDomainManager/secrets"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Record struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
	Proxied bool   `json:"proxied"`
}

func Create(secrets *secrets.Secret, record *Record) {

	data := record
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
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
		log.Fatal("FATAL ERROR")
		// handle err
	}
	if resp.StatusCode != 200 {
		io.Copy(os.Stdout, resp.Body)
		log.Fatal("Error creating record")
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println("Record successfully created")
	}

}
