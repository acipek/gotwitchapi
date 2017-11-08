package gotwitchapi

import (
	"net/url"
	"time"
)

// Follow object
type Follow struct {
	Data []struct {
		FromID     string    `json:"from_id"`
		ToID       string    `json:"to_id"`
		FollowedAt time.Time `json:"followed_at"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// GetFollow Gets information on follow relationships between two Twitch users.
func (client *Client) GetFollow(options map[string]string) (Follow, error) {
	targetURL := baseURL + "/users/follows"
	uv := url.Values{}
	for k, v := range options {
		if v != "" {
			uv.Add(k, v)
		}
	}
	if len(uv) != 0 {
		targetURL += "?" + uv.Encode()
	}

	sr := Follow{}

	err := client.getRequest(targetURL, &sr)

	return sr, err
}
