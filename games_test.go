package gotwitchapi

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestGames(t *testing.T) {
	clientID := os.Getenv("TWITCH_CLIENT_ID")
	g := NewClient(clientID)

	opt := map[string]string{"id": "493057"}

	games, err := g.GetGames(opt)
	if err != nil {
		t.Errorf("error not nil: %v", err)
	}

	b, err := json.MarshalIndent(games, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
