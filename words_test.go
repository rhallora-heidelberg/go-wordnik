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

func TestSearch(t *testing.T) {
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	//TODO: finish tests
	//search with default options
	_, _ = cl.Search("dem")

	//search having set all options
	_, _ = cl.Search("dem", CaseSensitive(true), IncludePartOfSpeech([]string{"noun"}), ExcludePartOfSpeech([]string{"adjective"}), MinCorpusCount(100), MaxCorpusCount(-1), MinDictionaryCount(1), MaxDictionaryCount(-1), MinLength(5), MaxLength(15), Skip(0), Limit(15))
}
