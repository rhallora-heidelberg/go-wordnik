package wordnik

import (
	"testing"
)

func TestExamples(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	_, err = cl.Examples("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.Examples("recalcitrant", Limit(20))
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

func TestWord(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	_, err = cl.Word("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.Word("thought")
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	}

	if res.CanonicalForm != "thought" {
		t.Error("expected query for 'thought' to return wordObject")
	}

	res, err = cl.Word("cats", UseCanonical(true))
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	}

	if res.CanonicalForm != "cat" {
		t.Error("expected 'cats' to return canonicalForm 'cat'")
	}
}

func TestDefinitions(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	_, err = cl.Definitions("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.Definitions("potato")
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	}

	if len(res) == 0 {
		t.Error("expected at least one result")
	}

	res, err = cl.Definitions("potato", SourceDictionaries([]string{"all"}))
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	}

	if len(res) < 10 {
		t.Error("expected at least 10 results")
	}
}

func TestTopExample(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	_, err = cl.TopExample("")
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
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	_, err = cl.RelatedWords("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.RelatedWords("mad", RelationshipTypes([]string{"synonym", "variant"}), LimitRelationshipType(1))
	if err != nil {
		t.Error("unexpected error")
	}

	if len(res) != 2 {
		t.Error("expected result of length 2")
	}
}

func TestPronunciations(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	_, err = cl.Pronunciations("")
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
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	_, err = cl.Hyphenation("")
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

func TestFrequency(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	_, err = cl.Frequency("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.Frequency("orange", StartYear(2000), EndYear(2009))
	if err != nil {
		t.Error("unexpected error")
	}

	if len(res.Frequency) != 10 {
		t.Error("expected result of length 10")
	}
}

func TestPhrases(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	_, err = cl.Phrases("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.Phrases("orange")
	if err != nil {
		t.Error("unexpected error")
	}

	if len(res) < 2 {
		t.Error("expected result of length 3 or more")
	}
}

func TestEtymologies(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	_, err = cl.Etymologies("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.Etymologies("orange")
	if err != nil {
		t.Error("unexpected error")
	}

	if len(res) < 1 {
		t.Error("expected result of length 1 or more")
	}
}

func TestAudio(t *testing.T) {
	t.Parallel()
	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	_, err = cl.Audio("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.Audio("likely")
	if err != nil {
		t.Error("unexpected error")
	}

	if len(res) < 2 {
		t.Error("expected result of length 2 or more")
	}
}
