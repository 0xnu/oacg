package oacg

import (
	"fmt"
	"net/http"
)

func Authentication(apiKey string, org string) error {
	url := "https://api.openai.com/v1/organizations/" + org + "/environments"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Authentication failed with status code %d", resp.StatusCode)
	}

	return nil
}
