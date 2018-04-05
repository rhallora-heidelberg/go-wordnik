package wordnik

import (
	"testing"
)

//NOTE: also dependent on WordListDELETE
func TestCreateWordList(t *testing.T) {
	t.Parallel()

	cl := getClient(t)
	auth, err := cl.getTestAuth(t)
	if err != nil {
		t.Fatal(err)
	}

	_, err = cl.CreateWordList("", WordList{})
	if err == nil {
		t.Error("expected error for empty string input")
	}

	wl := WordList{
		Name: "CreateWordList-test",
		Type: "PRIVATE",
	}

	res, err := cl.CreateWordList(auth.Token, wl)
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	} else if res.Name != "CreateWordList-test" {
		t.Error("expected returned wordList name to match given name")
	} else {
		// clean-up
		err = cl.DeleteWordList(auth.Token, res.Permalink)
		if err != nil {
			t.Error("error while cleaning up via WordListDELETE: " + err.Error())
		}
	}
}
