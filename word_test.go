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

// func TestExamples(t *testing.T) {
// 	t.Parallel()
//   testAPIKey, err := getEnvKey()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	cl := NewClient(testAPIKey)
// }
//
// func TestExamples(t *testing.T) {
// 	t.Parallel()
//   testAPIKey, err := getEnvKey()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	cl := NewClient(testAPIKey)
// }
//
// func TestExamples(t *testing.T) {
// 	t.Parallel()
//   testAPIKey, err := getEnvKey()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	cl := NewClient(testAPIKey)
// }
//
// func TestExamples(t *testing.T) {
// 	t.Parallel()
//   testAPIKey, err := getEnvKey()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	cl := NewClient(testAPIKey)
// }
//
// func TestExamples(t *testing.T) {
// 	t.Parallel()
//   testAPIKey, err := getEnvKey()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	cl := NewClient(testAPIKey)
// }
//
// func TestExamples(t *testing.T) {
// 	t.Parallel()
//   testAPIKey, err := getEnvKey()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	cl := NewClient(testAPIKey)
// }
//
// func TestExamples(t *testing.T) {
// 	t.Parallel()
//   testAPIKey, err := getEnvKey()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	cl := NewClient(testAPIKey)
// }
//
// func TestExamples(t *testing.T) {
// 	t.Parallel()
//   testAPIKey, err := getEnvKey()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	cl := NewClient(testAPIKey)
// }
