package gotwitchapi

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStream(t *testing.T) {
	clientID := os.Getenv("TWITCH_CLIENT_ID")
	g := NewClient(clientID)

	opt := map[string]string{"first": "2"}
	u, err := g.GetStream(opt)
	if err != nil {
		t.Errorf("error not nil: %v", err)
	}

	b, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
