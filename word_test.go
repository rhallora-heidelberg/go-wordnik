package wordnik

import (
	"testing"
)

func TestGetExamples(t *testing.T) {
	t.Parallel()
	cl := getClient(t)
	_, err := cl.GetExamples("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.GetExamples("recalcitrant", Limit(20))
	if err != nil {
		t.Error("unexpected error:" + err.Error())
	}

	found := false
	for _, example := range res.Examples {
		if example.ExampleID == 635971537 {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected example #635971537 in results for 'recalcitrant'")
	}
}

func TestGetWord(t *testing.T) {
	t.Parallel()
	cl := getClient(t)
	_, err := cl.GetWord("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.GetWord("thought")
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	}

	if res.CanonicalForm != "thought" {
		t.Error("expected query for 'thought' to return wordObject")
	}

	res, err = cl.GetWord("cats", UseCanonical(true))
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	}

	if res.CanonicalForm != "cat" {
		t.Error("expected 'cats' to return canonicalForm 'cat'")
	}
}

func TestGetDefinitions(t *testing.T) {
	t.Parallel()
	cl := getClient(t)
	_, err := cl.GetDefinitions("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.GetDefinitions("potato")
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	}

	if len(res) == 0 {
		t.Error("expected at least one result")
	}

	res, err = cl.GetDefinitions("potato", SourceDictionaries("all"))
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	}

	if len(res) < 10 {
		t.Error("expected at least 10 results")
	}
}

func TestTopExample(t *testing.T) {
	t.Parallel()
	cl := getClient(t)
	_, err := cl.TopExample("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.TopExample("potato", UseCanonical(false))
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	}

	if res.Word != "potato" {
		t.Error("expected result to contain 'potato' as word")
	}
}

func TestRelationshipTypes(t *testing.T) {
	t.Parallel()
	cl := getClient(t)
	_, err := cl.GetRelatedWords("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.GetRelatedWords("mad", RelationshipTypes("synonym", "variant"), LimitRelationshipType(1))
	if err != nil {
		t.Error("unexpected error")
	}

	if len(res) != 2 {
		t.Error("expected result of length 2")
	}
}

func TestPronunciations(t *testing.T) {
	t.Parallel()
	cl := getClient(t)
	_, err := cl.Pronunciations("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.Pronunciations("tomato")
	if err != nil {
		t.Error("unexpected error")
	}

	if len(res) < 2 {
		t.Error("expected result of length 2 or greater")
	}
}

func TestHyphenation(t *testing.T) {
	t.Parallel()
	cl := getClient(t)
	_, err := cl.Hyphenation("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.Hyphenation("orange")
	if err != nil {
		t.Error("unexpected error")
	}

	if len(res) != 2 {
		t.Error("expected result of length 2")
	}
}

func TestGetWordFrequency(t *testing.T) {
	t.Parallel()
	cl := getClient(t)
	_, err := cl.GetWordFrequency("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.GetWordFrequency("orange", StartYear(2000), EndYear(2009))
	if err != nil {
		t.Error("unexpected error")
	}

	if len(res.Frequency) != 10 {
		t.Error("expected result of length 10")
	}
}

func TestGetPhrases(t *testing.T) {
	t.Parallel()
	cl := getClient(t)
	_, err := cl.GetPhrases("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.GetPhrases("orange")
	if err != nil {
		t.Error("unexpected error")
	}

	if len(res) < 2 {
		t.Error("expected result of length 3 or more")
	}
}

func TestGetEtymologies(t *testing.T) {
	t.Parallel()
	cl := getClient(t)
	_, err := cl.GetEtymologies("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.GetEtymologies("orange")
	if err != nil {
		t.Error("unexpected error")
	}

	if len(res) < 1 {
		t.Error("expected result of length 1 or more")
	}
}

func TestGetAudio(t *testing.T) {
	t.Parallel()
	cl := getClient(t)
	_, err := cl.GetAudio("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.GetAudio("likely")
	if err != nil {
		t.Error("unexpected error")
	}

	if len(res) < 2 {
		t.Error("expected result of length 2 or more")
	}
}
