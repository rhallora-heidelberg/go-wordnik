package wordnik

import (
	"testing"
)

// tests GetWordList, UpdateWordList, and DeleteWordList
func TestWordListFuncs(t *testing.T) {
	t.Parallel()

	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	auth, err := cl.getTestAuth()
	if err != nil {
		t.Fatal(err)
	}

	_, err = cl.GetWordList("", "")
	if err == nil {
		t.Error("expected error for empty string input to GetWordList")
	}

	err = cl.UpdateWordList("", "", WordList{})
	if err == nil {
		t.Error("expected error for empty string input to UpdateWordList")
	}

	err = cl.DeleteWordList("", "")
	if err == nil {
		t.Error("expected error for empty string input to DeleteWordList")
	}

	testList := WordList{
		Name: "WordListFuncTest",
		Type: "PRIVATE",
	}
	res, err := cl.CreateWordList(auth.Token, testList)
	if err != nil {
		t.Fatal("unexpected error while POSTing wordList: " + err.Error())
	}

	res.Description = "updated"
	err = cl.UpdateWordList(auth.Token, res.Permalink, res)
	if err != nil {
		t.Fatal("unexpected error while updating wordList: " + err.Error())
	}

	res, err = cl.GetWordList(auth.Token, res.Permalink)
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	} else if res.Description != "updated" {
		t.Error("failed to verify updated WordList")
	}

	err = cl.DeleteWordList(auth.Token, res.Permalink)
	if err != nil {
		t.Error("unexpected error in DeleteWordList: " + err.Error())
	}
}

// Tests DeleteWordsFromWordList, GetWordListWords, AddWordsToWordList
func TestWordListWordsFuncs(t *testing.T) {
	t.Parallel()

	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	auth, err := cl.getTestAuth()
	if err != nil {
		t.Fatal(err)
	}

	err = cl.AddWordsToWordList("", "", []string{"lamp", "speaker"})
	if err == nil {
		t.Error("expected error for empty string input to AddWordsToWordList")
	}

	_, err = cl.GetWordListWords("", "")
	if err == nil {
		t.Error("expected error for empty string input to GetWordListWords")
	}

	err = cl.DeleteWordsFromWordList("", "", []string{"dock", "glass"})
	if err == nil {
		t.Error("expected error for empty string input to DeleteWordsFromWordList")
	}

	testList := WordList{
		Name: "WordListWordsFuncTest",
		Type: "PRIVATE",
	}
	res, err := cl.CreateWordList(auth.Token, testList)
	if err != nil {
		t.Fatal("unexpected error while POSTing wordList: " + err.Error())
	}

	testWords := []string{"lamp", "speaker"}
	err = cl.AddWordsToWordList(auth.Token, res.Permalink, testWords)
	if err != nil {
		t.Error("unexpected error in AddWordsToWordList: " + err.Error())
	}

	words, err := cl.GetWordListWords(auth.Token, res.Permalink)
	if err != nil {
		t.Error("unexpected error in GetWordListWords: " + err.Error())
	} else if len(words) != 2 {
		t.Error("expected wordList to have two entries")
	}

	err = cl.DeleteWordsFromWordList(auth.Token, res.Permalink, testWords)
	if err != nil {
		t.Error("unexpected error in GetWordListWords: " + err.Error())
	}

	words, err = cl.GetWordListWords(auth.Token, res.Permalink)
	if err != nil {
		t.Error("unexpected error in GetWordListWords: " + err.Error())
	} else if len(words) != 0 {
		t.Error("failed to delete words using DeleteWordsFromWordList")
	}

	err = cl.DeleteWordList(auth.Token, res.Permalink)
	if err != nil {
		t.Error("unexpected error in DeleteWordList: " + err.Error())
	}
}
