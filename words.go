package wordnik

import (
	"errors"
	"net/url"
)

// WordSearchResult as defined by the Wordnik API.
type WordSearchResult struct {
	Count      int64   `json:"count"`
	Lexicality float64 `json:"lexicality"`
	Word       string  `json:"word"`
}

// WordSearchResults as defined by the Wordnik API.
type WordSearchResults struct {
	SearchResults []WordSearchResult `json:"searchResults"`
	TotalResults  int64              `json:"totalResults"`
}

// WordOfTheDay as defined by the Wordnik API.
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

// ContentProvider as defined by the Wordnik API.
type ContentProvider struct {
	ID   int64  `json:"is"`
	Name string `json:"name"`
}

// SimpleDefinition as defined by the Wordnik API.
type SimpleDefinition struct {
	Text         string `json:"text"`
	Source       string `json:"source"`
	Note         string `json:"note"`
	PartOfSpeech string `json:"partOfSpeech"`
}

// SimpleExample as defined by the Wordnik API.
type SimpleExample struct {
	ID    int64  `json:"is"`
	Title string `json:"title"`
	Text  string `json:"text"`
	URL   string `json:"url"`
}

// DefinitionSearchResults as defined by the Wordnik API.
type DefinitionSearchResults struct {
	Results      []Definition `json:"results"`
	TotalResults int64        `json:"totalResults"`
}

// Definition as defined by the Wordnik API.
type Definition struct {
	ExtendedText     string         `json:"extendedText"`
	Text             string         `json:"text"`
	SourceDictionary string         `json:"sourceDictionary"`
	Citations        []Citation     `json:"citations"`
	Labels           []Label        `json:"labels"`
	Score            float64        `json:"score"` //'NaN' will be zero-valued
	ExampleUses      []ExampleUsage `json:"exampleUses"`
	AttributionURL   string         `json:"attributionUrl"`
	SeqString        string         `json:"seqString"`
	AttributionText  string         `json:"attributionText"`
	RelatedWords     []RelatedWord  `json:"relatedWords"`
	Sequence         string         `json:"sequence"`
	Word             string         `json:"word"`
	Notes            []Note         `json:"notes"`
	TextProns        []TextPron     `json:"textProns"`
	PartOfSpeech     string         `json:"partOfSpeech"`
}

// Citation as defined by the Wordnik API.
type Citation struct {
	Cite   string `json:"cite"`
	Source string `json:"source"`
}

// Label as defined by the Wordnik API.
type Label struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// ExampleUsage as defined by the Wordnik API.
type ExampleUsage struct {
	Text string `json:"text"`
}

// RelatedWord as defined by the Wordnik API.
type RelatedWord struct {
	Label1           string   `json:"label1"`
	RelationshipType string   `json:"relationshipType"`
	Label2           string   `json:"label2"`
	Label3           string   `json:"label3"`
	Words            []string `json:"words"`
	Gram             string   `json:"gram"`
	Label4           string   `json:"label4"`
}

// Note as defined by the Wordnik API.
type Note struct {
	NoteType  string   `json:"noteType"`
	AppliesTo []string `json:"appliesTo"`
	Value     string   `json:"value"`
	Pos       int64    `json:"pos"`
}

// TextPron as defined by the Wordnik API.
type TextPron struct {
	Raw     string `json:"raw"`
	Seq     int64  `json:"seq"`
	RawType string `json:"rawType"`
}

// WordObjects is a container type for WordObject, for the purpose of
// unmarshalling JSON.
type WordObjects struct {
	Words []WordObject
}

// GetWordOfTheDay returns the word of the day for a given date string in the
// format "yyyy-MM-dd".
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

// Search returns the results of a word search. Returns an error for empty input,
// but other 'incorrect' parameters are left to the APIs discretion. Configured
// with QueryOption functions, which ensure basic parameter vailidity.
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

	var results WordSearchResults
	err := c.basicGetRequest(rel, q, &results, queryOptions...)

	return results, err
}

// ReverseDictionary returns the result of a reverse dictionary search. Returns
// an error for empty input, but other 'incorrect' parameters are left to the
// APIs discretion. Configured with QueryOption functions, which ensure basic
// parameter vailidity. See Wordnik docs for appropriate parameters:
// http://developer.wordnik.com/docs.html#!/words/reverseDictionary_get_2
func (c *Client) ReverseDictionary(query string, queryOptions ...QueryOption) (DefinitionSearchResults, error) {
	if query == "" {
		return DefinitionSearchResults{}, errors.New("empty query string not allowed")
	}

	rel := &url.URL{Path: "words.json/reverseDictionary"}

	// Default Query Values
	q := url.Values{
		"query":          []string{query},
		"minCorpusCount": []string{"5"},
		"maxCorpusCount": []string{"-1"},
		"minLength":      []string{"1"},
		"maxLength":      []string{"-1"},
		"includeTags":    []string{"false"},
		"skip":           []string{"0"},
		"limit":          []string{"10"},
	}

	for _, option := range queryOptions {
		option(&q)
	}

	req, err := c.formRequest(rel, q, "GET")
	if err != nil {
		return DefinitionSearchResults{}, err
	}

	var results DefinitionSearchResults
	err = c.doRequest(req, &results)

	if err != nil && err.Error() == "json: cannot unmarshal string into Go struct field Definition.score of type float64" {
		// This error can be ignored, as it means that 'NaN' was present and will
		// be zero-valued, as we would have done anyway.
		err = nil
	}

	return results, err
}

// RandomWords returns random words as a WordObjects struct, with optional
// constraints. Configured with QueryOption functions, which ensure basic
// parameter vailidity. See Wordnik docs for appropriate parameters:
// http://developer.wordnik.com/docs.html#!/words/getRandomWords_get_3
func (c *Client) RandomWords(queryOptions ...QueryOption) (WordObjects, error) {
	rel := &url.URL{Path: "words.json/randomWord"}

	// Default values
	q := url.Values{
		"hasDictionaryDef":   []string{"false"},
		"minCorpusCount":     []string{"0"},
		"maxCorpusCount":     []string{"-1"},
		"minDictionaryCount": []string{"1"},
		"maxDictionaryCount": []string{"-1"},
		"minLength":          []string{"5"},
		"maxLength":          []string{"-1"},
		"limit":              []string{"10"},
	}

	var results WordObjects
	err := c.basicGetRequest(rel, q, &results, queryOptions...)

	return results, err
}

// RandomWord returns a random word as a WordObject, with optional constraints.
// Configured with QueryOption functions, which ensure basic parameter
// vailidity. See Wordnik docs for appropriate parameters:
// http://developer.wordnik.com/docs.html#!/words/getRandomWord_get_4
func (c *Client) RandomWord(queryOptions ...QueryOption) (WordObject, error) {
	rel := &url.URL{Path: "words.json/randomWord"}

	// Default values
	q := url.Values{
		"hasDictionaryDef":   []string{"false"},
		"minCorpusCount":     []string{"0"},
		"maxCorpusCount":     []string{"-1"},
		"minDictionaryCount": []string{"1"},
		"maxDictionaryCount": []string{"-1"},
		"minLength":          []string{"5"},
		"maxLength":          []string{"-1"},
	}

	var results WordObject
	err := c.basicGetRequest(rel, q, &results, queryOptions...)

	return results, err
}
