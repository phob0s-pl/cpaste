package pastebin

//
//import (
//	"encoding/json"
//	"io/ioutil"
//	"os"
//	"path/filepath"
//	"testing"
//)
//
//func TestLoadConfigNoDir(t *testing.T) {
//	var (
//		noDir = "/nonexisting/dir/file"
//	)
//
//	config, err := LoadConfig(noDir)
//	if err != nil {
//		t.Errorf("LoadConfig failed: %s", err)
//	}
//
//	if config == nil {
//		t.Fatalf("LoadConfig returned nil config")
//	}
//
//	if config.UserKey != "" {
//		t.Errorf("expected to have UserKey empty")
//	}
//
//	if config.DevKey != "" {
//		t.Errorf("expected to have DevKey empty")
//	}
//}
//
//func TestLoadAPIKeys(t *testing.T) { // nolint: gocyclo
//	var (
//		fileName = "testconf.json"
//	)
//
//	dir, err := ioutil.TempDir("/tmp", "cpastetest")
//	if err != nil {
//		t.Fatalf("TempDir failed: %s", err)
//	}
//	filePath := filepath.Join(dir, fileName)
//
//	defer func() {
//		if err = os.RemoveAll(dir); err != nil {
//			t.Errorf("failed to clean after test: %s", err)
//		}
//	}()
//
//	testConfig := &Config{
//		DevKey:     "key1",
//		UserKey:    "key2",
//		Name:       "somename",
//		Format:     "unformatted",
//		ExpireDate: "1Y",
//		Private:    0,
//	}
//
//	data, err := json.Marshal(testConfig)
//	if err != nil {
//		t.Fatalf("Marshal failed: %s", err)
//	}
//
//	if err = ioutil.WriteFile(filePath, data, 0644); err != nil {
//		t.Fatalf("WriteFile failed: %s", err)
//	}
//
//	config, err := LoadConfig(filePath)
//	if err != nil {
//		t.Fatalf("LoadConfig failed: %s", err)
//	}
//
//	if config == nil {
//		t.Fatalf("LoadConfig returned nil config")
//	}
//
//	if config.UserKey != testConfig.UserKey {
//		t.Errorf("expected to have UserKey %q, got %q", testConfig.UserKey, config.UserKey)
//	}
//
//	if config.DevKey != testConfig.DevKey {
//		t.Errorf("expected to have DevKey %q, got %q", testConfig.DevKey, config.DevKey)
//	}
//
//	if config.Name != "" {
//		t.Errorf("expected to have Name unset after loading")
//	}
//
//	if config.Format != "" {
//		t.Errorf("expected to have Format unset after loading")
//	}
//}
