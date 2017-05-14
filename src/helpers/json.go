package helpers

import (
	"encoding/json"
	"net/http"
	"time"
)

var jsonClient = &http.Client{Timeout: 10 * time.Second}

// GetJSON will return
func GetJSON(url string, target interface{}) error {
	r, err := jsonClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
