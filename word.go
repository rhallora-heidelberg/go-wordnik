package wordnik

import (
	"errors"
	"net/url"
)

// The WordObject as defined by the Wordnik API.
type WordObject struct {
	ID            int64    `json:"id"`
	Word          string   `json:"word"`
	OriginalWord  string   `json:"originalWord"`
	Suggestions   []string `json:"suggestions"`
	CanonicalForm string   `json:"canonicalForm"`
	Vulgar        string   `json:"vulgar"`
}

// ExampleSearchResults as defined by the Wordnik API.
type ExampleSearchResults struct {
	Facets   []Facet   `json:"facets"`
	Examples []Example `json:"examples"`
}

// Facet as defined by the Wordnik API.
type Facet struct {
	FacetValues []FacetValue `json:"facetValues"`
	Name        string       `json:"name"`
}

// FacetValue as defined by the Wordnik API.
type FacetValue struct {
	Count int64  `json:"count"`
	Value string `json:"value"`
}

// Example as defined by the Wordnik API.
type Example struct {
	ID         int64           `json:"id"`
	ExampleID  int64           `json:"exampleId"`
	Title      string          `json:"title"`
	Text       string          `json:"text"`
	Score      ScoredWord      `json:"score"`
	Sentence   Sentence        `json:"sentence"`
	Word       string          `json:"word"`
	Provider   ContentProvider `json:"provider"`
	Year       int64           `json:"year"`
	Rating     float64         `json:"rating"`
	DocumentID int64           `json:"documentId"`
	URL        string          `json:"url"`
}

// Sentence as defined by the Wordnik API.
type Sentence struct {
	HasScoredWords     bool         `json:"hasScoredWords"`
	ID                 int64        `json:"id"`
	ScoredWords        []ScoredWord `json:"scoredWords"`
	Display            string       `json:"display"`
	Rating             int64        `json:"rating"`
	DocumentMetadataID int64        `json:"documentMetadataId"`
}

// ScoredWord as defined by the Wordnik API.
type ScoredWord struct {
	Position      string `json:"position"`
	ID            string `json:"id"`
	DocTermCount  string `json:"docTermCount"`
	Lemma         string `json:"lemma"`
	WordType      string `json:"wordType"`
	Score         string `json:"score"`
	SentenceID    string `json:"sentenceId"`
	Word          string `json:"word"`
	Stopword      string `json:"stopword"`
	BaseWordScore string `json:"baseWordScore"`
	PartOfSpeech  string `json:"partOfSpeech"`
}

// Syllable as defined by the Wordnik API.
type Syllable struct {
	Text string `json:"text"`
	Seq  int64  `json:"seq"`
	Type string `json:"type"`
}

// FrequencySummary as defined by the Wordnik API.
type FrequencySummary struct {
	UnknownYearCount int64       `json:"unknownYearCount"`
	TotalCount       int64       `json:"totalCount"`
	FrequencyString  string      `json:"frequencyString"`
	Word             string      `json:"word"`
	Frequency        []Frequency `json:"frequency"`
}

// Frequency as defined by the Wordnik API.
type Frequency struct {
	Count int64 `json:"count"`
	Year  int64 `json:"year"`
}

// Bigram as defined by the Wordnik API.
type Bigram struct {
	Count int64   `json:"count"`
	Gram2 string  `json:"gram2"`
	Gram1 string  `json:"gram1"`
	Wlmi  float64 `json:"wlmi"`
	Mi    float64 `json:"mi"`
}

// EtymologiesResponse is provided for convenience, since results are returned
// as a list of strings with xml formatting.
type EtymologiesResponse []string

// AudioFile as defined by the Wordnik API. Note from the docs:
// The metadata includes a time-expiring fileUrl which allows reading the audio
// file directly from the API. Currently only audio pronunciations from the
// American Heritage Dictionary in mp3 format are supported.
type AudioFile struct {
	AttributionURL      string  `json:"attributionUrl"`
	CommentCount        int64   `json:"commentCount"`
	VoteCount           int64   `json:"voteCount"`
	FileURL             string  `json:"fileUrl"`
	AudioType           string  `json:"audioType"`
	ID                  int64   `json:"id"`
	Duration            float64 `json:"duration"`
	AttributionText     string  `json:"attributionText"`
	CreatedBy           string  `json:"createdBy"`
	Description         string  `json:"description"`
	CreatedAt           string  `json:"createdAt"`
	VoteWeightedAverage float64 `json:"voteWeightedAverage"`
	VoteAverage         float64 `json:"voteAverage"`
	Word                string  `json:"word"`
}

// GetExamples returns examples for a given word, with optional constraints.
// Configured with QueryOption functions, which ensure basic parameter
// vailidity.
func (c *Client) GetExamples(word string, queryOptions ...QueryOption) (ExampleSearchResults, error) {
	if word == "" {
		return ExampleSearchResults{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "word.json/" + word + "/examples"}

	// Default values
	q := url.Values{
		"includeDuplicates": []string{"false"},
		"useCanonical":      []string{"false"},
		"skip":              []string{"0"},
		"limit":             []string{"5"},
	}

	var results ExampleSearchResults
	err := c.basicGetRequest(rel, q, &results, queryOptions...)

	return results, err
}

// GetWord returns a WordObject for a given word, with optional constraints.
// Configured with QueryOption functions, which ensure basic parameter
// vailidity.
func (c *Client) GetWord(word string, queryOptions ...QueryOption) (WordObject, error) {
	if word == "" {
		return WordObject{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "word.json/" + word}

	// Default values
	q := url.Values{
		"useCanonical":       []string{"false"},
		"includeSuggestions": []string{"true"},
	}

	var results WordObject
	err := c.basicGetRequest(rel, q, &results, queryOptions...)

	return results, err
}

// GetDefinitions returns definitions for a given word, with optional constraints.
// Configured with QueryOption functions, which ensure basic parameter
// vailidity.
func (c *Client) GetDefinitions(word string, queryOptions ...QueryOption) ([]Definition, error) {
	if word == "" {
		return []Definition{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "word.json/" + word + "/definitions"}

	// Default values
	q := url.Values{
		"limit":          []string{"200"},
		"includeRelated": []string{"false"},
		"useCanonical":   []string{"false"},
		"includeTags":    []string{"false"},
	}

	var results []Definition
	err := c.basicGetRequest(rel, q, &results, queryOptions...)

	return results, err
}

// TopExample returns the top example for a given word, with optional constraints.
// Configured with QueryOption functions, which ensure basic parameter
// vailidity.
func (c *Client) TopExample(word string, options ...QueryOption) (Example, error) {
	if word == "" {
		return Example{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "word.json/" + word + "/topExample"}

	// Default values
	q := url.Values{"useCanonical": []string{"false"}}

	var results Example
	err := c.basicGetRequest(rel, q, &results, options...)

	return results, err
}

// GetRelatedWords returns related words for a given word, with optional constraints.
// Configured with QueryOption functions, which ensure basic parameter
// vailidity.
func (c *Client) GetRelatedWords(word string, queryOptions ...QueryOption) ([]RelatedWord, error) {
	if word == "" {
		return []RelatedWord{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "word.json/" + word + "/relatedWords"}

	// Default values
	q := url.Values{
		"useCanonical":           []string{"false"},
		"limitRelationshipTypes": []string{"10"},
	}

	var results []RelatedWord
	err := c.basicGetRequest(rel, q, &results, queryOptions...)

	return results, err
}

// Pronunciations returns pronunciations for a given word, with optional constraints.
// Configured with QueryOption functions, which ensure basic parameter
// vailidity.
func (c *Client) Pronunciations(word string, queryOptions ...QueryOption) ([]TextPron, error) {
	if word == "" {
		return []TextPron{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "word.json/" + word + "/pronunciations"}

	// Default values
	q := url.Values{
		"limit":        []string{"50"},
		"useCanonical": []string{"false"},
	}

	var results []TextPron
	err := c.basicGetRequest(rel, q, &results, queryOptions...)

	return results, err
}

// Hyphenation returns hyphenated portions of a given word, with optional constraints.
// Configured with QueryOption functions, which ensure basic parameter
// vailidity.
func (c *Client) Hyphenation(word string, queryOptions ...QueryOption) ([]Syllable, error) {
	if word == "" {
		return []Syllable{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "word.json/" + word + "/hyphenation"}

	// Default values
	q := url.Values{
		"limit":        []string{"50"},
		"useCanonical": []string{"false"},
	}

	var results []Syllable
	err := c.basicGetRequest(rel, q, &results, queryOptions...)

	return results, err
}

// GetWordFrequency returns a frequency summary for a given word, with optional constraints.
// Configured with QueryOption functions, which ensure basic parameter
// vailidity.
func (c *Client) GetWordFrequency(word string, queryOptions ...QueryOption) (FrequencySummary, error) {
	if word == "" {
		return FrequencySummary{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "word.json/" + word + "/frequency"}

	// Default values
	q := url.Values{
		"useCanonical": []string{"false"},
	}

	var results FrequencySummary
	err := c.basicGetRequest(rel, q, &results, queryOptions...)

	return results, err
}

// GetPhrases returns two-word phrases for a given word, with optional constraints.
// Configured with QueryOption functions, which ensure basic parameter
// vailidity.
func (c *Client) GetPhrases(word string, queryOptions ...QueryOption) ([]Bigram, error) {
	if word == "" {
		return []Bigram{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "word.json/" + word + "/phrases"}

	// Default values
	q := url.Values{
		"limit":        []string{"5"},
		"wlmi":         []string{"0"},
		"useCanonical": []string{"false"},
	}

	var results []Bigram
	err := c.basicGetRequest(rel, q, &results, queryOptions...)

	return results, err
}

// GetEtymologies returns etymologies for a given word, with optional constraints.
// Configured with QueryOption functions, which ensure basic parameter
// vailidity.
func (c *Client) GetEtymologies(word string, queryOptions ...QueryOption) (EtymologiesResponse, error) {
	if word == "" {
		return EtymologiesResponse{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "word.json/" + word + "/etymologies"}

	// Default values
	q := url.Values{"useCanonical": []string{"false"}}

	var results EtymologiesResponse
	err := c.basicGetRequest(rel, q, &results, queryOptions...)

	return results, err
}

// GetAudio returns a link to pronunciations for a given word, with optional constraints.
// Configured with QueryOption functions, which ensure basic parameter
// vailidity.
func (c *Client) GetAudio(word string, queryOptions ...QueryOption) ([]AudioFile, error) {
	if word == "" {
		return []AudioFile{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "word.json/" + word + "/audio"}

	// Default values
	q := url.Values{
		"limit":        []string{"50"},
		"useCanonical": []string{"false"},
	}

	var results []AudioFile
	err := c.basicGetRequest(rel, q, &results, queryOptions...)

	return results, err
}
