package wordnik

import (
	"bytes"
	"net/url"
	"strconv"
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
		"all":        true,
		"ahd":        true,
		"century":    true,
		"wiktionary": true,
		"webster":    true,
		"wordnet":    true,
	}

	validSortCriteria = map[string]bool{
		"alpha":      true,
		"count":      true,
		"length":     true,
		"createDate": true, //only applies to WordListWord queries
	}

	validExpandTerms = map[string]bool{
		"synonym":  true,
		"hypernym": true,
	}

	validSortOrder = map[string]bool{
		"asc":  true,
		"desc": true,
	}

	validRelationshipTypes = map[string]bool{
		"synonym":         true,
		"antonym":         true,
		"variant":         true,
		"equivalent":      true,
		"cross-reference": true,
		"related-word":    true,
		"rhyme":           true,
		"form":            true,
		"etymologically-related-term": true,
		"hypernym":                    true,
		"hyponym":                     true,
		"inflected-form":              true,
		"primary":                     true,
		"same-context":                true,
		"verb-form":                   true,
		"verb-stem":                   true,
	}

	validTypeFormat = map[string]bool{
		"ahd":               true,
		"arpabet":           true,
		"gcide-diacritical": true,
		"IPA":               true,
	}
)

// QueryOption functions return functions which modify optional query parameters
// by acting on url.Values pointers.
type QueryOption func(*url.Values)

func buildCommaSepQuery(items []string, validityMap map[string]bool) string {
	var buffer bytes.Buffer

	for _, item := range items {
		if validityMap[item] {
			buffer.WriteString(item)
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

// CaseSensitive sets the caseSensitive parameter based on boolean input.
func CaseSensitive(b bool) QueryOption {
	return func(q *url.Values) {
		q.Set("caseSensitive", strconv.FormatBool(b))
	}
}

// IncludePartOfSpeech sets the includePartOfSpeech parameter based on variadic
// string input.
func IncludePartOfSpeech(parts ...string) QueryOption {
	return func(q *url.Values) {
		value := buildCommaSepQuery(parts, validPartOfSpeech)
		q.Set("includePartOfSpeech", value)
	}
}

// ExcludePartOfSpeech sets the excludePartOfSpeech parameter based on variadic
// string input.
func ExcludePartOfSpeech(parts ...string) QueryOption {
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

// IncludeSourceDictionaries sets the includeSourceDictionaries parameter based
// on variadic string input.
func IncludeSourceDictionaries(dicts ...string) QueryOption {
	return func(q *url.Values) {
		value := buildCommaSepQuery(dicts, validSourceDictionaries)
		q.Set("includeSourceDictionaries", value)
	}
}

// ExcludeSourceDictionaries sets the excludeSourceDictionaries parameter based
// variadic string input.
func ExcludeSourceDictionaries(dicts ...string) QueryOption {
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

// HasDictionaryDef sets the hasDictionaryDef parameter based on boolean input.
func HasDictionaryDef(b bool) QueryOption {
	return func(q *url.Values) {
		q.Set("hasDictionaryDef", strconv.FormatBool(b))
	}
}

// UseCanonical sets the useCanonical parameter based on boolean input.
func UseCanonical(b bool) QueryOption {
	return func(q *url.Values) {
		q.Set("useCanonical", strconv.FormatBool(b))
	}
}

// IncludeSuggestions sets the includeSuggestions parameter based on boolean input.
func IncludeSuggestions(b bool) QueryOption {
	return func(q *url.Values) {
		q.Set("includeSuggestions", strconv.FormatBool(b))
	}
}

// IncludeDuplicates sets the includeDuplicates parameter based on boolean input.
func IncludeDuplicates(b bool) QueryOption {
	return func(q *url.Values) {
		q.Set("includeDuplicates", strconv.FormatBool(b))
	}
}

// IncludeRelated sets the includeRelated parameter based on boolean input.
func IncludeRelated(b bool) QueryOption {
	return func(q *url.Values) {
		q.Set("includeRelated", strconv.FormatBool(b))
	}
}

// PartOfSpeech sets the partOfSpeech parameter based on variadic string input.
func PartOfSpeech(parts ...string) QueryOption {
	return func(q *url.Values) {
		value := buildCommaSepQuery(parts, validPartOfSpeech)
		q.Set("partOfSpeech", value)
	}
}

// SourceDictionaries sets the sourceDictionaries parameter based on variadic string
// input. Differs notably in effect from "includeSourceDictionaries" when used
// in the context of Definitions. According to the API:  Source dictionary to
// return definitions from. If 'all' is received, results are returned from all
// sources. If multiple values are received (e.g. 'century,wiktionary'), results
//  are returned from the first specified dictionary that has definitions. If
// left blank, results are returned from the first dictionary that has
// definitions. By default, dictionaries are searched in this order: ahd,
// wiktionary, webster, century, wordnet
func SourceDictionaries(dicts ...string) QueryOption {
	return func(q *url.Values) {
		value := buildCommaSepQuery(dicts, validSourceDictionaries)
		q.Set("sourceDictionaries", value)
	}
}

// RelationshipTypes sets the relationshipTypes parameter based on variadic string
// input. This parameter works in conjunction with limitRelationshipType,
// in that this list of relationship types is allowed but each type is limited
// in how many examples it returns by limitRelationshipType.
func RelationshipTypes(types ...string) QueryOption {
	return func(q *url.Values) {
		value := buildCommaSepQuery(types, validRelationshipTypes)
		q.Set("relationshipTypes", value)
	}
}

// LimitRelationshipType sets the limitRelationshipType parameter based on
// integer input. This parameter works in conjunction with relationshipTypes,
// in that the list of relationship types is allowed but each type is limited
// in how many examples it returns by limitRelationshipType.
func LimitRelationshipType(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("limitRelationshipType", strconv.FormatInt(n, 10))
	}
}

// TypeFormat sets the typeFormat parameter based on string input.
func TypeFormat(format string) QueryOption {
	return func(q *url.Values) {
		if validTypeFormat[format] {
			q.Set("typeFormat", format)
		}
	}
}

// SourceDictionary sets the sourceDictionary parameter based on string input.
func SourceDictionary(format string) QueryOption {
	return func(q *url.Values) {
		if validSourceDictionaries[format] {
			q.Set("sourceDictionary", format)
		}
	}
}

// StartYear sets the startYear parameter based on integer input.
func StartYear(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("startYear", strconv.FormatInt(n, 10))
	}
}

// EndYear sets the endYear parameter based on integer input.
func EndYear(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("endYear", strconv.FormatInt(n, 10))
	}
}
