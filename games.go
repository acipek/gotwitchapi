package gotwitchapi

import (
	"net/url"
)

// Game object
type Game struct {
	Data []struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		BoxArtURL string `json:"box_art_url"`
	} `json:"data"`
}

// GetGames Gets game information by game ID or name.
func (client *Client) GetGames(options map[string]string) (Game, error) {
	targetURL := baseURL + "/games"

	uv := url.Values{}
	for k, v := range options {
		if v != "" {
			uv.Add(k, v)
		}
	}
	if len(uv) != 0 {
		targetURL += "?" + uv.Encode()
	}

	gr := Game{}

	err := client.getRequest(targetURL, &gr)

	return gr, err
}
