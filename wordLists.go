package wordnik

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
)

// CreateWordList attempts to create a word list for a given account. Returns the
// list as a WordList object if successful.
func (c *Client) CreateWordList(authToken string, list WordList) (WordList, error) {
	if authToken == "" {
		return WordList{}, errors.New("empty auth token not allowed")
	}

	rel := &url.URL{Path: "wordLists.json"}
	marshalledList, err := json.Marshal(list)
	if err != nil {
		return WordList{}, err
	}

	body := bytes.NewBuffer(marshalledList)
	req, err := c.formRequest(rel, url.Values{}, "POST", body)
	if err != nil {
		return WordList{}, err
	}

	req.Header["auth_token"] = []string{authToken}

	var results WordList
	err = c.doRequest(req, &results)

	return results, err
}
