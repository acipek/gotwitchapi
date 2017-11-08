package gotwitchapi

import (
	"net/url"

	"time"
)

// Stream object
type Stream struct {
	Data []struct {
		ID           string    `json:"id"`
		UserID       string    `json:"user_id"`
		GameID       string    `json:"game_id"`
		CommunityIds []string  `json:"community_ids"`
		Type         string    `json:"type"`
		Title        string    `json:"title"`
		ViewerCount  int       `json:"viewer_count"`
		StartedAt    time.Time `json:"started_at"`
		Language     string    `json:"language"`
		ThumbnailURL string    `json:"thumbnail_url"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// GetStream Gets information about active streams.
func (client *Client) GetStream(options map[string]string) (Stream, error) {
	targetURL := baseURL + "/streams"
	uv := url.Values{}
	for k, v := range options {
		if v != "" {
			uv.Add(k, v)
		}
	}
	if len(uv) != 0 {
		targetURL += "?" + uv.Encode()
	}

	sr := Stream{}

	err := client.getRequest(targetURL, &sr)

	return sr, err
}
