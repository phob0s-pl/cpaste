package pastebin

import "fmt"

var (
	// errKeysNotConfigured is returned when at least one key is not present
	errKeysNotConfigured = fmt.Errorf("keys not configured")
)
