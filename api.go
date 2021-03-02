package octoprint

import (
	"bytes"
  "net/http"
)

var apiKey string
var host string

// Sets the API Key for the Octoprint Connection
func SetAPIKey(key string) {
	apiKey = key
}
// Sets the Hostname for the Octoprint Connection (needs to be in the format http(s)://<server>/)
func SetHost(hostname string) {
	host = hostname
}

func isAPIKeySet() bool{
	return len(apiKey) > 0
}

func validateHostName()  {

}
// Returns the Hostname
func GetHost() string{
	return host
}
//executes a GET request against the provided Path (using the apiKey as authorization)
func Get(path string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", host+path, nil)
	req.Header.Set("X-Api-Key", apiKey)
	res, _ := client.Do(req)

	return res

}
//executes a POST request against the provided Path with the provided body (using the apiKey as authorization)
func Post(path string, body []byte) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", host+path, bytes.NewBuffer(body))
	req.Header.Set("X-Api-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	res, _ := client.Do(req)
	return res
}
