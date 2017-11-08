package gotwitchapi

import (
	"net/url"
)

// User object
type User struct {
	Data []struct {
		ID              string `json:"id"`
		Login           string `json:"login"`
		DisplayName     string `json:"display_name"`
		Type            string `json:"type"`
		BroadcasterType string `json:"broadcaster_type"`
		Description     string `json:"description"`
		ProfileImageURL string `json:"profile_image_url"`
		OfflineImageURL string `json:"offline_image_url"`
		ViewCount       int    `json:"view_count"`
		Email           string `json:"email"`
	} `json:"data"`
}

// GetUser Gets information about one or more specified Twitch users.
func (client *Client) GetUser(options map[string]string) (User, error) {
	targetURL := baseURL + "/users"

	uv := url.Values{}
	for k, v := range options {
		if v != "" {
			uv.Add(k, v)
		}
	}
	if len(uv) != 0 {
		targetURL += "?" + uv.Encode()
	}

	ur := User{}

	err := client.getRequest(targetURL, &ur)

	return ur, err
}
