package logdna

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const URL = "https://api.logdna.com/v2/"

func (client *LogdnaClient) basicAuth() string {
	auth := client.ServiceKey + ":"
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (client *LogdnaClient) signRequest(method string, path string, body []byte) *http.Request {
	req, _ := http.NewRequest(method, URL+path, bytes.NewBuffer(body))
	req.Header.Set("Accept", "application/json")
	req.Header.Add("Authorization", "Basic "+client.basicAuth())
	return req
}

func (client *LogdnaClient) _get(path string, body []byte) (*http.Response, error) {
	preparedRequest := client.signRequest("GET", path, body)
	resp, err := client.Client.Do(preparedRequest)
	return resp, err
}

// func (client *LogdnaClient) _post(path string, body []byte) (*http.Response, error) {
// 	preparedRequest := client.signRequest("POST", path, body)
// 	resp, err := client.Client.Do(preparedRequest)
// 	return resp, err
// }

// func (client *LogdnaClient) _delete(path string, body []byte) (*http.Response, error) {
// 	preparedRequest := client.signRequest("DELETE", path, body)
// 	resp, err := client.Client.Do(preparedRequest)
// 	return resp, err
// }

func processResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error processing response: %v", err)
		return err
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		log.Printf("Error processing response: %v", err)
		return err
	}
	return nil
}
