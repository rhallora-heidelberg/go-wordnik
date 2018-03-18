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

// Syllables is a container type for Syllable, for the purpose of
// unmarshalling JSON.
type Syllables struct {
	Syllables []Syllable
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

// Bigrams is a container type for Bigram, for the purpose of
// unmarshalling JSON.
type Bigrams struct {
	Bigrams []Bigram
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

// AudioFiles is a container type for AudioFile, for the purpose of
// unmarshalling JSON.
type AudioFiles struct {
	AudioFiles []AudioFile
}

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

func (c *Client) Examples(word string, queryOptions ...QueryOption) (ExampleSearchResults, error) {
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

	for _, option := range queryOptions {
		option(&q)
	}

	req, err := c.formRequest(rel, q, "GET")
	if err != nil {
		return ExampleSearchResults{}, err
	}

	var results ExampleSearchResults
	err = c.doRequest(req, &results)

	return results, err
}

func (c *Client) Word(word string, queryOptions ...QueryOption) (WordObject, error) {
	if word == "" {
		return WordObject{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "word.json/" + word}

	// Default values
	q := url.Values{
		"useCanonical":       []string{"false"},
		"includeSuggestions": []string{"true"},
	}

	for _, option := range queryOptions {
		option(&q)
	}

	req, err := c.formRequest(rel, q, "GET")
	if err != nil {
		return WordObject{}, err
	}

	var results WordObject
	err = c.doRequest(req, &results)

	return results, err
}

func (c *Client) Definitions(word string, queryOptions ...QueryOption) ([]Definition, error) {
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

	for _, option := range queryOptions {
		option(&q)
	}

	req, err := c.formRequest(rel, q, "GET")
	if err != nil {
		return []Definition{}, err
	}

	var results []Definition
	err = c.doRequest(req, &results)

	return results, err
}
