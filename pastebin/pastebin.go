package pastebin

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	// pasteURL is used for submitting pastes
	pasteURL = "https://pastebin.com/api/api_post.php"
	// pasteURL is used for acquiring session keys
	loginURL = "https://pastebin.com/api/api_login.php"
	// actionNewPaste is action where user POSTs his paste to the server
	actionNewPaste = "paste"
	// actionListPaste is action where user gets list of hist pastes from the server
	actionListPaste = "list"
	// actionDeletePaste removes paste with give api_paste_key
	actionDeletePaste = "delete"
	// userKeyLen is length in bytes of user key
	userKeyLen = 32
)

// form fields
const (
	apiOption          = "api_option"
	apiUserKey         = "api_user_key"
	apiPastePrivate    = "api_paste_private"
	apiPasteName       = "api_paste_name"
	apiPasteExpireDate = "api_paste_expire_date"
	apiPasteFormat     = "api_paste_format"
	apiDevKey          = "api_dev_key"
	apiPasteCode       = "api_paste_code"
	apiUserName        = "api_user_name"
	apiUserPassword    = "api_user_password"
	apiResultsLimit    = "api_results_limit"
	apiPasteKey        = "api_paste_key"
)

const (
	// PastePublic makes public paste
	PastePublic = 0
	// PasteUnlisted make paste unlisted
	PasteUnlisted = 1
	// PastePrivate makes paste private
	PastePrivate = 2
)

// Client handles operations on pastebin.com
type Client struct {
	keys *Keys
}

// Paste contains paste metadata like format or name
type Paste struct {
	XMLName xml.Name `xml:"paste"`
	// Key is paste unique identifier used to
	Key string `xml:"paste_key"`
	// Date is unix time describing when paste was published
	Date  int64  `xml:"paste_date"`
	Title string `xml:"paste_title"`
	Size  uint   `xml:"paste_size"`
	// ExpireDate is unix date when paste will be removed from server
	ExpireDate int64 `xml:"paste_expire_date"`
	Private    uint8 `xml:"paste_private"`
	// FormatLong is long, descriptive format of the paste
	// For more details see: https://pastebin.com/api#5
	FormatLong string `xml:"paste_format_long"`
	// FormatLong is short format of the paste
	FormatShort string `xml:"paste_format_short"`
	// URL is address where the paste is accessible
	URL  string `xml:"paste_url"`
	Hits int64  `xml:"paste_hits"`
	// Expire is period
	Expire string `xml:"-"`
}

// NewClient creates new client
func NewClient(apiKeysPath string) (client *Client, err error) {
	client = &Client{}
	client.keys, err = readKeys(apiKeysPath)
	return client, err
}

// Publish sends data to pastebin.com server with options specified in paste
// It returns url where paste is hosted. In case something goes wrong,
// error is returned.
func (c *Client) Publish(paste *Paste, data []byte) (string, error) {
	if !c.keys.present() {
		return "", errKeysNotConfigured
	}

	form := createPasteForm(c.keys, paste, data)
	pasteURL, err := post(pasteURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}

	return pasteURL, nil
}

// List returns list of pastes of current user
func (c *Client) List(limit int) ([]Paste, error) {
	if !c.keys.present() {
		return nil, errKeysNotConfigured
	}

	form := createListForm(c.keys, limit)
	pasteList, err := post(pasteURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	list, err := parseListXML(pasteList)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// Delete removes paste with given key from the server
func (c *Client) Delete(pasteKey string) (string, error) {
	if !c.keys.present() {
		return "", errKeysNotConfigured
	}

	form := createDeletePasteForm(c.keys, pasteKey)
	return post(pasteURL, strings.NewReader(form.Encode()))
}

func parseListXML(rawList string) (list []Paste, err error) {
	decoder := xml.NewDecoder(bytes.NewBufferString(rawList))
	for {
		var single Paste
		err := decoder.Decode(&single)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		list = append(list, single)
	}
	return list, nil
}

// RequestUserKey requests key from pastebin.com
// The dev key and user key are stored in path in json format
func RequestUserKey(credentials *Credentials, path string) error {
	form := createKeyRequestForm(credentials)
	key, err := post(loginURL, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	if len(key) != userKeyLen {
		return fmt.Errorf("invalid user key received: %s", key)
	}

	keys := Keys{
		DevKey:  credentials.DevKey,
		UserKey: key,
	}

	return keys.save(path)
}

// createRequest creates POST request which can be sent to server
func createRequest(postURL string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", postURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func createPasteForm(keys *Keys, paste *Paste, code []byte) url.Values {
	form := url.Values{}
	form.Add(apiDevKey, keys.DevKey)
	form.Add(apiOption, actionNewPaste)
	form.Add(apiUserKey, keys.UserKey)
	form.Add(apiPasteName, paste.Title)
	form.Add(apiPasteExpireDate, paste.Expire)
	form.Add(apiPasteFormat, paste.FormatShort)
	form.Add(apiPasteCode, string(code))
	form.Add(apiPastePrivate, fmt.Sprintf("%d", paste.Private))
	return form
}

func createKeyRequestForm(c *Credentials) url.Values {
	form := url.Values{}
	form.Add(apiDevKey, c.DevKey)
	form.Add(apiUserName, c.User)
	form.Add(apiUserPassword, c.Password)
	return form
}

func createListForm(keys *Keys, limit int) url.Values {
	form := url.Values{}
	form.Add(apiDevKey, keys.DevKey)
	form.Add(apiOption, actionListPaste)
	form.Add(apiUserKey, keys.UserKey)
	form.Add(apiResultsLimit, strconv.Itoa(limit))
	return form
}

func createDeletePasteForm(keys *Keys, pasteKey string) url.Values {
	form := url.Values{}
	form.Add(apiDevKey, keys.DevKey)
	form.Add(apiOption, actionDeletePaste)
	form.Add(apiUserKey, keys.UserKey)
	form.Add(apiPasteKey, pasteKey)
	return form
}

// post creates proper POST request and sends to provided URL
// it returns content of response body
func post(postURL string, body io.Reader) (string, error) {
	request, err := http.NewRequest("POST", postURL, body)
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with code: %d", response.StatusCode)
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(content), err
}
