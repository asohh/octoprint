package octoprint

import (
	"bytes"
  "net/http"
)

var apiKey string
var host string

func SetAPIKey(key string) {
	apiKey = key
}
func IsAPIKeySet() bool{
	return len(apiKey) > 0
}
func SetHost(hostname string) {
	host = hostname
}
func GetHost() string{
	return host
}
func Get(path string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", host+path, nil)
	req.Header.Set("X-Api-Key", apiKey)
	res, _ := client.Do(req)

	return res

}

func Post(path string, body []byte) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", host+path, bytes.NewBuffer(body))
	req.Header.Set("X-Api-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	res, _ := client.Do(req)
	return res
}
