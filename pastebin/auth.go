package pastebin

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Keys contains key necessary for authentication in order to perform operations on
// pastebin.com
type Keys struct {
	// DevKey your api_developer_key
	DevKey string `json:"api_dev_key"`
	// User key is key obtained from server. Only one key can be valid
	// at one time
	UserKey string `json:"api_user_key"`
}

// Credentials is set of data required to obtain user key from pastebin
type Credentials struct {
	User     string
	Password string
	DevKey   string
}

// readKeys reads key from file
// If there is no file empty structure is returned
func readKeys(path string) (*Keys, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &Keys{}, nil
	} else if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(path) // nolint: gosec
	if err != nil {
		return nil, err
	}

	keys := &Keys{}
	if err := json.Unmarshal(content, keys); err != nil {
		return nil, err
	}

	return keys, nil
}

func (k *Keys) save(path string) error {
	data, err := json.Marshal(k)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, 0600)
}

// present returns true if both keys are present
func (k *Keys) present() bool {
	return k.UserKey != "" && k.DevKey != ""
}
