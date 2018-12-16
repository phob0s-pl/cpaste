package pastebin

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadKeys(t *testing.T) {
	var (
		tempFile = "/tmp/test_file_cpaste_load.json"
	)

	raw := `{"api_dev_key":"dkey","api_user_key":"ukey"}`

	if err := ioutil.WriteFile(tempFile, []byte(raw), 0600); err != nil {
		t.Fatalf("failed to create file for test: %s", err)
	}

	defer func() {
		if err := os.Remove(tempFile); err != nil {
			t.Fatalf("failed to clean file after test: %s", err)
		}
	}()

	content, err := ioutil.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("failed to read file: %s", err)
	}

	var keys Keys
	if err := json.Unmarshal(content, &keys); err != nil {
		t.Fatalf("failed to unserialize: %s", err)
	}

	if keys.UserKey != "ukey" {
		t.Errorf("UserKey expected %q, got %q", "ukey", keys.UserKey)
	}

	if keys.DevKey != "dkey" {
		t.Errorf("DevKey expected %q, got %q", "dkey", keys.DevKey)
	}
}
