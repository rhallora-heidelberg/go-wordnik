package wordnik

import (
	"net/url"
	"testing"
)

type boolQueryTest struct {
	b        bool
	expected string
}

type stringSliceQueryTest struct {
	strings  []string
	expected string
}

type int64QueryTest struct {
	n        int64
	expected string
}

func queryTestBools(t *testing.T, testCases []boolQueryTest, f func(bool) QueryOption) {
	for _, testCase := range testCases {
		q := url.Values{}
		CaseSensitive(testCase.b)(&q)
		if q.Encode() != testCase.expected {
			t.Errorf("For %v got %q, expected: %q", testCase.b, q.Encode(), testCase.expected)
		}
	}
}

func queryTestStringSlices(t *testing.T, testCases []stringSliceQueryTest, f func([]string) QueryOption) {
	for _, testCase := range testCases {
		q := url.Values{}
		f(testCase.strings)(&q)
		if q.Encode() != testCase.expected {
			t.Errorf("For %v got %q, expected: %q", testCase.strings, q.Encode(), testCase.expected)
		}
	}
}

func queryTestInt64s(t *testing.T, testCases []int64QueryTest, f func(int64) QueryOption) {
	for _, testCase := range testCases {
		q := url.Values{}
		f(testCase.n)(&q)
		if q.Encode() != testCase.expected {
			t.Errorf("For %v got %q, expected: %q", testCase.n, q.Encode(), testCase.expected)
		}
	}
}

var caseSensitiveTests = []boolQueryTest{
	{true, "caseSensitive=true"},
	{false, "caseSensitive=false"},
}

func TestCaseSensitive(t *testing.T) {
	queryTestBools(t, caseSensitiveTests, CaseSensitive)
}

var includePartOfSpeechTests = []stringSliceQueryTest{
	// Normal case
	{[]string{"noun", "interjection"}, "includePartOfSpeech=noun%2Cinterjection%2C"},

	// Some parts of speech invalid
	{[]string{"orange", "noun", "noun,"}, "includePartOfSpeech=noun%2C"},

	// All parts of speech invalid
	{[]string{"", "orange", ","}, "includePartOfSpeech="},

	// Empty input
	{[]string{}, "includePartOfSpeech="},

	// Repeated input (note: this is allowed by the API, but changes nothing)
	{[]string{"noun", "noun"}, "includePartOfSpeech=noun%2Cnoun%2C"},
}

func TestIncludePartOfSpeech(t *testing.T) {
	queryTestStringSlices(t, includePartOfSpeechTests, IncludePartOfSpeech)
}

var excludePartOfSpeechTests = []stringSliceQueryTest{
	// Normal case
	{[]string{"noun", "interjection"}, "excludePartOfSpeech=noun%2Cinterjection%2C"},

	// Some parts of speech invalid
	{[]string{"orange", "noun", "noun,"}, "excludePartOfSpeech=noun%2C"},

	// All parts of speech invalid
	{[]string{"", "orange", ","}, "excludePartOfSpeech="},

	// Empty input
	{[]string{}, "excludePartOfSpeech="},

	// Repeated input (note: this is allowed by the API, but changes nothing)
	{[]string{"noun", "noun"}, "excludePartOfSpeech=noun%2Cnoun%2C"},
}

func TestExcludePartOfSpeech(t *testing.T) {
	queryTestStringSlices(t, excludePartOfSpeechTests, ExcludePartOfSpeech)
}

var minCorpusCountTests = []int64QueryTest{
	{0, "minCorpusCount=0"},
	{100, "minCorpusCount=100"},
}

func TestMinCorpusCount(t *testing.T) {
	queryTestInt64s(t, minCorpusCountTests, MinCorpusCount)
}

var maxCorpusCountTests = []int64QueryTest{
	{-10, "maxCorpusCount=-10"},
	{100, "maxCorpusCount=100"},
}

func TestMaxCorpusCount(t *testing.T) {
	queryTestInt64s(t, maxCorpusCountTests, MaxCorpusCount)
}

var maxDictionaryCountTests = []int64QueryTest{
	{-10, "maxDictionaryCount=-10"},
	{100, "maxDictionaryCount=100"},
}

func TestMaxDictionaryCount(t *testing.T) {
	queryTestInt64s(t, maxDictionaryCountTests, MaxDictionaryCount)
}

var minDictionaryCountTests = []int64QueryTest{
	{-10, "minDictionaryCount=-10"},
	{100, "minDictionaryCount=100"},
}

func TestMinDictionaryCount(t *testing.T) {
	queryTestInt64s(t, minDictionaryCountTests, MinDictionaryCount)
}

var minLengthTests = []int64QueryTest{
	{-33, "minLength=-33"},
	{101, "minLength=101"},
}

func TestMinLength(t *testing.T) {
	queryTestInt64s(t, minLengthTests, MinLength)
}

var maxLengthTests = []int64QueryTest{
	{-33, "maxLength=-33"},
	{101, "maxLength=101"},
}

func TestMaxLength(t *testing.T) {
	queryTestInt64s(t, maxLengthTests, MaxLength)
}

var skipTests = []int64QueryTest{
	{-10, "skip=-10"},
	{100, "skip=100"},
}

func TestSkip(t *testing.T) {
	queryTestInt64s(t, skipTests, Skip)
}

var limitTests = []int64QueryTest{
	{-10, "limit=-10"},
	{100, "limit=100"},
}

func TestLimit(t *testing.T) {
	queryTestInt64s(t, limitTests, Limit)
}
