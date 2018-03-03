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
)

type QueryOption func(*url.Values)

func CaseSensitive(b bool) QueryOption {
	return func(q *url.Values) {
		q.Set("caseSensitive", strconv.FormatBool(b))
	}
}

func IncludePartOfSpeech(parts []string) QueryOption {
	return func(q *url.Values) {
		var builder strings.Builder

		for _, part := range parts {
			if validPartOfSpeech[part] {
				builder.WriteString(part)
				builder.WriteString(",")
			}
		}
		q.Set("includePartOfSpeech", builder.String())
	}
}

func ExcludePartOfSpeech(parts []string) QueryOption {
	return func(q *url.Values) {
		var builder strings.Builder

		for _, part := range parts {
			if validPartOfSpeech[part] {
				builder.WriteString(part)
				builder.WriteString(",")
			}
		}
		q.Set("excludePartOfSpeech", builder.String())
	}
}

func MinCorpusCount(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("minCorpusCount", strconv.FormatInt(n, 10))
	}
}

func MaxCorpusCount(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("maxCorpusCount", strconv.FormatInt(n, 10))
	}
}

func MinDictionaryCount(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("minDictionaryCount", strconv.FormatInt(n, 10))
	}
}

func MaxDictionaryCount(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("maxDictionaryCount", strconv.FormatInt(n, 10))
	}
}

func MinLength(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("minLength", strconv.FormatInt(n, 10))
	}
}

func MaxLength(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("maxLength", strconv.FormatInt(n, 10))
	}
}

func Skip(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("skip", strconv.FormatInt(n, 10))
	}
}

func Limit(n int64) QueryOption {
	return func(q *url.Values) {
		q.Set("limit", strconv.FormatInt(n, 10))
	}
}
