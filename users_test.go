package gotwitchapi

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestUser(t *testing.T) {
	clientID := os.Getenv("TWITCH_CLIENT_ID")
	g := NewClient(clientID)

	opt := map[string]string{"id": "140746824"}

	u, err := g.GetUser(opt)
	if err != nil {
		t.Errorf("error not nil: %v", err)
	}

	b, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
