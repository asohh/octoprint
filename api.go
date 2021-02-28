package octoprint

import (
	"bytes"
	"fmt"
  "net/http"
  "io/ioutil"
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
func Get(path string, key string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", host+path, nil)
	fmt.Println(err)
	// req.Header.Set("X-Api-Key", key)
	fmt.Println(key)
	res, _ := client.Do(req)
	fmt.Println(res)
	if res.StatusCode == http.StatusOK {

		bodyBytes, _ := ioutil.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
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
