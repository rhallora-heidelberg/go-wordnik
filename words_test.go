package wordnik

import (
	"testing"
)

var wotdTests = []struct {
	date, expectedWord string
	errExpected        bool
}{
	//Normal cases
	{"2017-02-03", "lairage", false},
	{"2016-03-03", "jacklight", false},
	{"2017-04-10", "farouche", false},
	{"2017-05-11", "alegar", false},
	{"2017-06-12", "carnassial", false},

	//Malformed dates
	{"200006-12", "", true},
	{"20170612", "", true},
	{"201706--12", "", true},
	{"2017--12", "", true},

	// Note: "impossible" dates such as "20000-06-12", "1000-03-12",
	// "2017-13-04" etc. will be allowed, letting the service dictate how
	// this is handled.

	//Empty string
	{"", "", true},
}

func TestGetWordOfTheDay(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	for _, testCase := range wotdTests {
		testResult, err := cl.GetWordOfTheDay(testCase.date)
		if err != nil && !testCase.errExpected {
			t.Errorf("Date %q: Unexpected error: %v", testCase.date, err)
		} else if testResult.Word != testCase.expectedWord {
			t.Errorf("For date %q got word %q; expected %q. Dump:\n %q", testCase.date, testResult.Word, testCase.expectedWord, testResult)
		}
	}
}

var SearchWordsTestCases = []struct {
	query        string
	options      []QueryOption
	expectResult bool
}{
	{"dem", []QueryOption{}, true},
	{"dem", []QueryOption{CaseSensitive(true), IncludePartOfSpeech([]string{"noun"}), ExcludePartOfSpeech([]string{"adjective"}), MinCorpusCount(100), MaxCorpusCount(-1), MinDictionaryCount(1), MaxDictionaryCount(-1), MinLength(5), MaxLength(15), Skip(0), Limit(15)}, false},
}

func TestSearchWords(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)

	_, err = cl.SearchWords("")
	if err == nil {
		t.Error("Expected error for empty query")
	}

	for _, testCase := range SearchWordsTestCases {
		res, err := cl.SearchWords(testCase.query, testCase.options...)
		if err != nil {
			t.Error("Unexpected error")
		}

		if res.TotalResults == 0 && testCase.expectResult {
			t.Errorf("Expected at least one result for query '%v'", testCase.query)
		}

		if res.TotalResults > 0 && res.SearchResults[0].Word != testCase.query {
			t.Errorf("Query '%v' not returned as first result", testCase.query)
		}

	}
}

var RevDictTestCases = []struct {
	query        string
	options      []QueryOption
	expectResult bool
	expectedWord string
}{
	{"having the bad qualities of a dog", []QueryOption{}, true, "doggish"},
	{"badly", []QueryOption{FindSenseForWord("incond"), IncludeSourceDictionaries([]string{"ahd"}), ExcludeSourceDictionaries([]string{"webster"}), ExpandTerms("bad input"), IncludeTags(false), SortBy("alpha")}, true, "incondite"},
	{"the quality of not having a result due to excessive specificity", []QueryOption{IncludeSourceDictionaries([]string{"webster"}), FindSenseForWord("excesspecifidocious")}, false, ""},
}

func TestRevDict(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)

	_, err = cl.ReverseDictionary("")
	if err == nil {
		t.Error("Expected error for empty query")
	}

	for _, testCase := range RevDictTestCases {
		res, err := cl.ReverseDictionary(testCase.query, testCase.options...)
		if err != nil {
			t.Error("Unexpected error")
		}

		if res.TotalResults == 0 && testCase.expectResult {
			t.Errorf("Expected at least one result for query '%v'", testCase.query)
		}

		if res.TotalResults > 0 && !testCase.expectResult {
			t.Errorf("Did not expect results for query '%v'", testCase.query)
		}

		if res.TotalResults > 0 && res.Results[0].Word != testCase.expectedWord {
			t.Errorf("Query '%v' did not return '%v' as first result", testCase.query, testCase.expectedWord)
		}
	}
}

func TestRandomWord(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)

	// Expect no result
	res, _ := cl.RandomWord(MinCorpusCount(2), MaxCorpusCount(1))
	if res.Word != "" {
		t.Error("Expected no result")
	}

	// Expect both to return a result of appropriate length
	res, err = cl.RandomWord(MinLength(5), MaxLength(5))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if len(res.Word) != 5 {
		t.Error("expected a result of length 5")
	}

	res, err = cl.RandomWord(MinLength(6), MaxLength(6))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if len(res.Word) != 6 {
		t.Error("expected a result of length 6")
	}
}

func TestRandomWords(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)

	// Expect no result
	res, err := cl.RandomWords(MinCorpusCount(2), MaxCorpusCount(1))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(res.Words) != 0 {
		t.Error("Expected no result")
	}

	// Expect several words of appropriate length
	res, err = cl.RandomWords(MinLength(5), MaxLength(5), Limit(5))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	for _, resWord := range res.Words {
		if len(resWord.Word) != 5 {
			t.Error("expected a result of length 5")
		}
	}
}
