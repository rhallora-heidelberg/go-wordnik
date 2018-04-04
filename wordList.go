package wordnik

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
)

// WordList as defined by the Wordnik API.
type WordList struct {
	ID                int64  `json:"id,omitempty"`
	Permalink         string `json:"permalink,omitempty"`
	Name              string `json:"name,omitempty"`
	CreatedAt         string `json:"createdAt,omitempty"`
	UpdatedAt         string `json:"updatedAt,omitempty"`
	LastActivityAt    string `json:"lastActivityAt,omitempty"`
	Username          string `json:"username,omitempty"`
	UserID            int64  `json:"userId,omitempty"`
	Description       string `json:"description,omitempty"`
	NumberWordsInList int64  `json:"numberWordsInList,omitempty"`
	Type              string `json:"type,omitempty"`
}

// stringValue is provided for convenient JSON marshalling in PostWordListWords.
type stringValue struct {
	Word string `json:"word,omitempty"`
}

// WordListWord as defined by the Wordnik API.
type WordListWord struct {
	ID                   int64  `json:"id"`
	Word                 string `json:"word"`
	Username             string `json:"username"`
	UserID               int64  `json:"userId"`
	CreatedAt            string `json:"createdAt"`
	NumberCommentsOnWord int64  `json:"numberCommentsOnWord"`
	NumberLists          int64  `json:"numberLists"`
}

// DeleteWordList deletes a WordList for a given user.
func (c *Client) DeleteWordList(authToken, permalink string) error {
	if authToken == "" || permalink == "" {
		return errors.New("empty auth token  or permalink not allowed")
	}

	rel := &url.URL{Path: "wordList.json/" + permalink}

	req, err := c.formRequest(rel, url.Values{}, "DELETE")
	if err != nil {
		return err
	}

	req.Header["auth_token"] = []string{authToken}

	err = c.doRequest(req, nil)

	return err
}

// UpdateWordList updates a WordList for a given user. Note that this refers to
// the properties of the WordList itself, not to adding or deleting words from
// the list.
func (c *Client) UpdateWordList(authToken, permalink string, wList WordList) error {
	if authToken == "" || permalink == "" {
		return errors.New("empty auth token  or permalink not allowed")
	}

	rel := &url.URL{Path: "wordList.json/" + permalink}
	marshalledList, err := json.Marshal(wList)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(marshalledList)
	req, err := c.formRequest(rel, url.Values{}, "PUT", body)
	if err != nil {
		return err
	}

	req.Header["auth_token"] = []string{authToken}

	err = c.doRequest(req, nil)

	return err
}

// GetWordList retrieves a WordList given it's permalink.
func (c *Client) GetWordList(authToken, permalink string) (WordList, error) {
	if authToken == "" || permalink == "" {
		return WordList{}, errors.New("empty auth token  or permalink not allowed")
	}

	rel := &url.URL{Path: "wordList.json/" + permalink}

	req, err := c.formRequest(rel, url.Values{}, "GET")
	if err != nil {
		return WordList{}, err
	}

	req.Header["auth_token"] = []string{authToken}

	var results WordList
	err = c.doRequest(req, &results)

	return results, err
}

// AddWordsToWordList adds words to a WordList.
func (c *Client) AddWordsToWordList(authToken, permalink string, words []string) error {
	if authToken == "" || permalink == "" {
		return errors.New("empty auth token  or permalink not allowed")
	}

	rel := &url.URL{Path: "wordList.json/" + permalink + "/words"}

	//convert words into []stringValue for ease of marshalling
	packedWords := make([]stringValue, len(words))
	for i, word := range words {
		packedWords[i].Word = word
	}

	marshalledWords, err := json.Marshal(packedWords)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(marshalledWords)
	req, err := c.formRequest(rel, url.Values{}, "POST", body)
	if err != nil {
		return err
	}

	req.Header["auth_token"] = []string{authToken}

	return c.doRequest(req, nil)
}

// GetWordListWords retrieves words from a WordList. Note that this may not be
// all of the words in the list, as determined by the "skip" and "limit" options.
func (c *Client) GetWordListWords(authToken, permalink string, options ...QueryOption) ([]WordListWord, error) {
	if authToken == "" || permalink == "" {
		return []WordListWord{}, errors.New("empty auth token  or permalink not allowed")
	}

	rel := &url.URL{Path: "wordList.json/" + permalink + "/words"}

	// Default Query Values
	q := url.Values{
		"sortBy":    []string{"createDate"},
		"sortOrder": []string{"asc"},
		"skip":      []string{"0"},
		"limit":     []string{"100"},
	}

	for _, option := range options {
		option(&q)
	}

	req, err := c.formRequest(rel, q, "GET")
	if err != nil {
		return []WordListWord{}, err
	}

	req.Header["auth_token"] = []string{authToken}

	var results []WordListWord
	err = c.doRequest(req, &results)

	return results, err
}

// DeleteWordsFromWordList deletes specific words from a WordList if they are
// present.
func (c *Client) DeleteWordsFromWordList(authToken, permalink string, words []string) error {
	if authToken == "" || permalink == "" {
		return errors.New("empty auth token  or permalink not allowed")
	}

	rel := &url.URL{Path: "wordList.json/" + permalink + "/deleteWords"}

	//convert words into []stringValue for ease of marshalling
	packedWords := make([]stringValue, len(words))
	for i, word := range words {
		packedWords[i].Word = word
	}

	marshalledWords, err := json.Marshal(packedWords)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(marshalledWords)
	req, err := c.formRequest(rel, url.Values{}, "POST", body)
	if err != nil {
		return err
	}

	req.Header["auth_token"] = []string{authToken}

	return c.doRequest(req, nil)
}
