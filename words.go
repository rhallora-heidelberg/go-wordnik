package wordnik

import (
	"errors"
	"net/url"
)

type WordSearchResult struct {
	Count      int64   `json:"count"`
	Lexicality float64 `json:"lexicality"`
	Word       string  `json:"word"`
}

type WordSearchResults struct {
	SearchResults []WordSearchResult `json:"searchResults"`
	TotalResults  int64              `json:"totalResults"`
}

type WordOfTheDay struct {
	ID              int64              `json:"id"`
	ParentID        string             `json:"parentId"`
	Category        string             `json:"category"`
	CreatedBy       string             `json:"createdBy"`
	CreatedAt       string             `json:"createdAt"`
	ContentProvider ContentProvider    `json:"contentProvider"`
	HTMLExtra       string             `json:"htmlExtra"`
	Word            string             `json:"word"`
	Definitions     []SimpleDefinition `json:"definitions"`
	Examples        []SimpleExample    `json:"examples"`
	Note            string             `json:"note"`
	PublishDate     string             `json:"publishDate"`
}

type ContentProvider struct {
	ID   int64  `json:"is"`
	Name string `json:"name"`
}

type SimpleDefinition struct {
	Text         string `json:"text"`
	Source       string `json:"source"`
	Note         string `json:"note"`
	PartOfSpeech string `json:"partOfSpeech"`
}

type SimpleExample struct {
	ID    int64  `json:"is"`
	Title string `json:"title"`
	Text  string `json:"text"`
	URL   string `json:"url"`
}

func (c *Client) GetWordOfTheDay(dateString string) (WordOfTheDay, error) {
	rel := &url.URL{Path: "words.json/wordOfTheDay"}

	q := url.Values{}
	q.Set("date", dateString)

	req, err := c.formRequest(rel, q, "GET")
	if err != nil {
		return WordOfTheDay{}, err
	}

	var wotd WordOfTheDay
	err = c.doRequest(req, &wotd)
	if err != nil {
		return WordOfTheDay{}, err
	}

	return wotd, nil
}

func (c *Client) Search(query string, queryOptions ...QueryOption) (WordSearchResults, error) {
	if query == "" {
		return WordSearchResults{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "words.json/search/" + query}

	// Default Query Values
	q := url.Values{
		"caseSensitive":      []string{"true"},
		"minCorpusCount":     []string{"5"},
		"maxCorpusCount":     []string{"-1"},
		"minDictionaryCount": []string{"1"},
		"maxDictionaryCount": []string{"-1"},
		"minLength":          []string{"1"},
		"maxLength":          []string{"-1"},
		"skip":               []string{"0"},
		"limit":              []string{"10"},
	}

	for _, option := range queryOptions {
		option(&q)
	}

	req, err := c.formRequest(rel, q, "GET")
	if err != nil {
		return WordSearchResults{}, err
	}

	var results WordSearchResults
	err = c.doRequest(req, &results)
	if err != nil {
		return WordSearchResults{}, err
	}

	return results, nil
}
