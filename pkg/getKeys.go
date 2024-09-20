package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"myvpn/utils"
	"log"
	"crypto/tls"
)

func GetAccessKeys() (utils.AccessKeyList, error) {
	
	config, err := utils.LoadConfig(".")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }

	url := fmt.Sprintf("%s/access-keys", config.ManagerAccessApiUrl)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return utils.AccessKeyList{}, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return utils.AccessKeyList{}, fmt.Errorf("error making POST request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return utils.AccessKeyList{}, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return utils.AccessKeyList{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result utils.AccessKeyList

	err = json.Unmarshal(body, &result)
	if err != nil {
		return utils.AccessKeyList{}, fmt.Errorf("error unmarshaling response JSON: %w", err)
	}

	return result, nil
}