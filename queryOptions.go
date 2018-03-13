package wordnik

import (
	"net/url"
	"strconv"
	"strings"
)

var (
	validPartOfSpeech = map[string]bool{
		"noun":                  true,
		"adjective":             true,
		"verb":                  true,
		"adverb":                true,
		"interjection":          true,
		"pronoun":               true,
		"preposition":           true,
		"abbreviation":          true,
		"affix":                 true,
		"article":               true,
		"auxiliary-verb":        true,
		"conjunction":           true,
		"definite-article":      true,
		"family-name":           true,
		"given-name":            true,
		"idiom":                 true,
		"imperative":            true,
		"noun-plural":           true,
		"noun-posessive":        true,
		"past-participle":       true,
		"phrasal-prefix":        true,
		"proper-noun":           true,
		"proper-noun-plural":    true,
		"proper-noun-posessive": true,
		"suffix":                true,
		"verb-intransitive":     true,
		"verb-transitive":       true,
	}

	validSourceDictionaries = map[string]bool{
		"ahd":        true,
		"century":    true,
		"wiktionary": true,
		"webster":    true,
		"wordnet":    true,
	}

	validSortCriteria = map[string]bool{
		"alpha":  true,
		"count":  true,
		"length": true,
	}

	validExpandTerms = map[string]bool{
		"synonym":  true,
		"hypernym": true,
	}

	validSortOrder = map[string]bool{
		"asc":  true,
		"desc": true,
	}
)

// QueryOption functions return functions which modify optional query parameters
// by acting on url.Values pointers.
type QueryOption func(*url.Values)

func buildCommaSepQuery(items []string, validityMap map[string]bool) string {
	var builder strings.Builder

	for _, item := range items {
		if validityMap[item] {
			builder.WriteString(item)
			builder.WriteString(",")
		}
	}
	return builder.String()
}

// CaseSensitive sets the caseSensitive parameter based on boolean input.
func CaseSensitive(b bool) QueryOption {
	return func(q *url.Values) {
		q.Set("caseSensitive", strconv.FormatBool(b))
	}
}

// IncludePartOfSpeech sets the includePartOfSpeech parameter based on string
// slice input.
func IncludePartOfSpeech(parts []string) QueryOption {
	return func(q *url.Values) {
		value := buildCommaSepQuery(parts, validPartOfSpeech)
		q.Set("includePartOfSpeech", value)
	}
}

// ExcludePartOfSpeech sets the excludePartOfSpeech parameter based on string
// slice input.
func ExcludePartOfSpeech(parts []string) QueryOption {
	return func(q *url.Values) {
		value := buildCommaSepQuery(parts, validPartOfSpeech)
		q.Set("excludePartOfSpeech", value)
	}
}

// MinCorpusCount sets the minCorpusCount parameter based on integer input.
func MinCorpusCount(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("minCorpusCount", strconv.FormatInt(n, 10))
	}
}

// MaxCorpusCount sets the maxCorpusCount parameter based on integer input.
func MaxCorpusCount(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("maxCorpusCount", strconv.FormatInt(n, 10))
	}
}

// MinDictionaryCount sets the minDictionaryCount parameter based on integer input.
func MinDictionaryCount(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("minDictionaryCount", strconv.FormatInt(n, 10))
	}
}

// MaxDictionaryCount sets the maxDictionaryCount parameter based on integer input.
func MaxDictionaryCount(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("maxDictionaryCount", strconv.FormatInt(n, 10))
	}
}

// MinLength sets the minLength parameter based on integer input.
func MinLength(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("minLength", strconv.FormatInt(n, 10))
	}
}

// MaxLength sets the maxLength parameter based on integer input.
func MaxLength(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("maxLength", strconv.FormatInt(n, 10))
	}
}

// Skip sets the skip parameter based on integer input.
func Skip(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("skip", strconv.FormatInt(n, 10))
	}
}

// Limit sets the limit parameter based on integer input.
func Limit(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("limit", strconv.FormatInt(n, 10))
	}
}

// FindSenseForWord sets the findSenseForWord parameter based on string input.
func FindSenseForWord(sense string) QueryOption {
	return func(q *url.Values) {
		q.Set("findSenseForWord", sense)
	}
}

// IncludeSourceDictionaries sets the includeSourceDictionaries parameter based on string slice input.
func IncludeSourceDictionaries(dicts []string) QueryOption {
	return func(q *url.Values) {
		value := buildCommaSepQuery(dicts, validSourceDictionaries)
		q.Set("includeSourceDictionaries", value)
	}
}

// ExcludeSourceDictionaries sets the excludeSourceDictionaries parameter based on string slice input.
func ExcludeSourceDictionaries(dicts []string) QueryOption {
	return func(q *url.Values) {
		value := buildCommaSepQuery(dicts, validSourceDictionaries)
		q.Set("excludeSourceDictionaries", value)
	}
}

// ExpandTerms sets the expandTerms parameter based on string input.
func ExpandTerms(term string) QueryOption {
	return func(q *url.Values) {
		if validExpandTerms[term] {
			q.Set("expandTerms", term)
		}
	}
}

// IncludeTags sets the includeTags parameter based on boolean input.
// Controls whether a closed set of XML tags should be returned in response.
func IncludeTags(b bool) QueryOption {
	return func(q *url.Values) {
		q.Set("includeTags", strconv.FormatBool(b))
	}
}

// SortBy sets the sortBy parameter based on string input.
func SortBy(sortCriteria string) QueryOption {
	return func(q *url.Values) {
		if validSortCriteria[sortCriteria] {
			q.Set("sortBy", sortCriteria)
		}
	}
}

// SortOrder sets the sortOrder parameter based on string input.
func SortOrder(direction string) QueryOption {
	return func(q *url.Values) {
		if validSortOrder[direction] {
			q.Set("sortOrder", direction)
		}
	}
}
