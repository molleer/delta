package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GammaGet(endpoint string, token string, body interface{}) error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s", gamma_url, endpoint), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(text, body)
}