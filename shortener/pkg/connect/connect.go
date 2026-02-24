package connect

import (
	"net/http"
	"time"
)

// global http client
var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: 2 * time.Second,
}

// Get checks if the url is valid by sending a GET request to the url, if the response status code is 200, return true, otherwise return false
func Get(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}