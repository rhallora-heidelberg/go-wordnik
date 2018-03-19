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

type stringQueryTest struct {
	s, expected string
}

func queryTestBools(t *testing.T, testCases []boolQueryTest, f func(bool) QueryOption) {
	for _, testCase := range testCases {
		q := url.Values{}
		f(testCase.b)(&q)
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

func queryTestStrings(t *testing.T, testCases []stringQueryTest, f func(string) QueryOption) {
	for _, testCase := range testCases {
		q := url.Values{}
		f(testCase.s)(&q)
		if q.Encode() != testCase.expected {
			t.Errorf("For %v got %q, expected: %q", testCase.s, q.Encode(), testCase.expected)
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

var findSenseTests = []stringQueryTest{
	{"a", "findSenseForWord=a"},
	{"b", "findSenseForWord=b"},
	{"a b", "findSenseForWord=a+b"},
}

func TestFindSenseForWord(t *testing.T) {
	queryTestStrings(t, findSenseTests, FindSenseForWord)
}

var includeSourceDictTests = []stringSliceQueryTest{
	{[]string{""}, "includeSourceDictionaries="},
	{[]string{"bird", "orange"}, "includeSourceDictionaries="},
	{[]string{"ahd"}, "includeSourceDictionaries=ahd%2C"},
	{[]string{"ahd", "wiktionary"}, "includeSourceDictionaries=ahd%2Cwiktionary%2C"},
}

func TestIncludeDictionaries(t *testing.T) {
	queryTestStringSlices(t, includeSourceDictTests, IncludeSourceDictionaries)
}

var excludeSourceDictTests = []stringSliceQueryTest{
	{[]string{""}, "excludeSourceDictionaries="},
	{[]string{"bird", "orange"}, "excludeSourceDictionaries="},
	{[]string{"ahd"}, "excludeSourceDictionaries=ahd%2C"},
	{[]string{"ahd", "wiktionary"}, "excludeSourceDictionaries=ahd%2Cwiktionary%2C"},
}

func TestExcludeDictionaries(t *testing.T) {
	queryTestStringSlices(t, excludeSourceDictTests, ExcludeSourceDictionaries)
}

var expandTermsTests = []stringQueryTest{
	{"", ""},
	{"f", ""},
	{"hypernym", "expandTerms=hypernym"},
	{"synonym", "expandTerms=synonym"},
}

func TestExpandTerms(t *testing.T) {
	queryTestStrings(t, expandTermsTests, ExpandTerms)
}

var includeTagsTests = []boolQueryTest{
	{true, "includeTags=true"},
	{false, "includeTags=false"},
}

func TestIncludeTags(t *testing.T) {
	queryTestBools(t, includeTagsTests, IncludeTags)
}

var sortByTests = []stringQueryTest{
	{"", ""},
	{"x", ""},
	{"alpha", "sortBy=alpha"},
	{"count", "sortBy=count"},
	{"length", "sortBy=length"},
}

func TestSortBy(t *testing.T) {
	queryTestStrings(t, sortByTests, SortBy)
}

var sortOrderTests = []stringQueryTest{
	{"", ""},
	{"up", ""},
	{"asc", "sortOrder=asc"},
	{"desc", "sortOrder=desc"},
}

func TestSortOrder(t *testing.T) {
	queryTestStrings(t, sortOrderTests, SortOrder)
}

var hasDictionaryDefTests = []boolQueryTest{
	{true, "hasDictionaryDef=true"},
	{false, "hasDictionaryDef=false"},
}

func TestHasDictionaryDef(t *testing.T) {
	queryTestBools(t, hasDictionaryDefTests, HasDictionaryDef)
}

var useCanonicalTests = []boolQueryTest{
	{true, "useCanonical=true"},
	{false, "useCanonical=false"},
}

func TestUseCanonical(t *testing.T) {
	queryTestBools(t, useCanonicalTests, UseCanonical)
}

var includeSuggestionsTests = []boolQueryTest{
	{true, "includeSuggestions=true"},
	{false, "includeSuggestions=false"},
}

func TestIncludeSuggestions(t *testing.T) {
	queryTestBools(t, includeSuggestionsTests, IncludeSuggestions)
}

var includeDuplicatesTests = []boolQueryTest{
	{true, "includeDuplicates=true"},
	{false, "includeDuplicates=false"},
}

func TestIncludeDuplicates(t *testing.T) {
	queryTestBools(t, includeDuplicatesTests, IncludeDuplicates)
}

var includeRelatedTests = []boolQueryTest{
	{true, "includeRelated=true"},
	{false, "includeRelated=false"},
}

func TestIncludeRelated(t *testing.T) {
	queryTestBools(t, includeRelatedTests, IncludeRelated)
}

var partOfSpeechTests = []stringSliceQueryTest{
	// Normal case
	{[]string{"noun", "interjection"}, "partOfSpeech=noun%2Cinterjection%2C"},

	// Some parts of speech invalid
	{[]string{"orange", "noun", "noun,"}, "partOfSpeech=noun%2C"},

	// All parts of speech invalid
	{[]string{"", "orange", ","}, "partOfSpeech="},

	// Empty input
	{[]string{}, "partOfSpeech="},

	// Repeated input (note: this is allowed by the API, but changes nothing)
	{[]string{"noun", "noun"}, "partOfSpeech=noun%2Cnoun%2C"},
}

func TestPartOfSpeech(t *testing.T) {
	queryTestStringSlices(t, partOfSpeechTests, PartOfSpeech)
}

var sourceDictTests = []stringSliceQueryTest{
	{[]string{""}, "sourceDictionaries="},
	{[]string{"bird", "orange"}, "sourceDictionaries="},
	{[]string{"ahd"}, "sourceDictionaries=ahd%2C"},
	{[]string{"ahd", "wiktionary"}, "sourceDictionaries=ahd%2Cwiktionary%2C"},
}

func TestSourceDictionaries(t *testing.T) {
	queryTestStringSlices(t, sourceDictTests, SourceDictionaries)
}

var relTypesTests = []stringSliceQueryTest{
	{[]string{""}, "relationshipTypes="},
	{[]string{"bird", "orange"}, "relationshipTypes="},
	{[]string{"synonym"}, "relationshipTypes=synonym%2C"},
	{[]string{"synonym", "variant"}, "relationshipTypes=synonym%2Cvariant%2C"},
}

func TestRelTypes(t *testing.T) {
	queryTestStringSlices(t, relTypesTests, RelationshipTypes)
}

var limitrelationshipTests = []int64QueryTest{
	{-10, "limitRelationshipType=-10"},
	{100, "limitRelationshipType=100"},
}

func TestLimitRelationshipType(t *testing.T) {
	queryTestInt64s(t, limitrelationshipTests, LimitRelationshipType)
}

var typeFormatTests = []stringQueryTest{
	{"", ""},
	{"f", ""},
	{"ahd", "typeFormat=ahd"},
	{"IPA", "typeFormat=IPA"},
}

func TestTypeFormat(t *testing.T) {
	queryTestStrings(t, typeFormatTests, TypeFormat)
}

var sourceDictionaryTests = []stringQueryTest{
	{"", ""},
	{"f", ""},
	{"ahd", "sourceDictionary=ahd"},
	{"century", "sourceDictionary=century"},
}

func TestSourceDictionary(t *testing.T) {
	queryTestStrings(t, sourceDictionaryTests, SourceDictionary)
}

var startYearTests = []int64QueryTest{
	{-10, "startYear=-10"},
	{100, "startYear=100"},
}

func TestStartYear(t *testing.T) {
	queryTestInt64s(t, startYearTests, StartYear)
}

var endYearTests = []int64QueryTest{
	{-10, "endYear=-10"},
	{100, "endYear=100"},
}

func TestEndYear(t *testing.T) {
	queryTestInt64s(t, endYearTests, EndYear)
}
