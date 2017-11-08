package gotwitchapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://api.twitch.tv/helix"

// Client is a struct for twitch client.
type Client struct {
	HTTPClient *http.Client
	ClientID   string
}

// NewClient creates a new twitch client for http requests.
func NewClient(clientID string) Client {
	return Client{
		HTTPClient: &http.Client{},
		ClientID:   clientID,
	}
}

func (client *Client) getRequest(url string, out interface{}) error {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("could not create request")
	}

	req.Header.Set("Client-ID", client.ClientID)
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %v", res.Status)
	}

	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}

	return nil
}
