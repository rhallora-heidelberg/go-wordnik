package wordnik

import (
	"testing"
)

func TestAuthenticateGET(t *testing.T) {
	t.Parallel()
	tUser, err := getEnvUserPass()
	if err != nil {
		t.Fatal(err)
	}

	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	_, err = cl.AuthenticateGET("", "")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.AuthenticateGET(tUser.user, tUser.pass)
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	} else if res.UserID == 0 {
		t.Error("expected non-zero user ID")
	}
}

func TestAuthenticatePOST(t *testing.T) {
	t.Parallel()
	tUser, err := getEnvUserPass()
	if err != nil {
		t.Fatal(err)
	}

	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	_, err = cl.AuthenticatePOST("", "")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.AuthenticatePOST(tUser.user, tUser.pass)
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	} else if res.UserID == 0 {
		t.Error("expected non-zero user ID")
	}
}

// Helper function for testing which attempts to retrieve an AuthenticationToken
// for username and password set by WORDNIK_TEST_USER and WORDNIK_TEST_PASS.
func (c *Client) getTestAuth() (AuthenticationToken, error) {
	tUser, err := getEnvUserPass()
	if err != nil {
		return AuthenticationToken{}, err
	}

	testAPIKey, err := getEnvKey()
	if err != nil {
		return AuthenticationToken{}, err
	}

	cl := NewClient(testAPIKey)
	return cl.AuthenticateGET(tUser.user, tUser.pass)
}

func TestGetAPITokenStatus(t *testing.T) {
	t.Parallel()

	testAPIKey, err := getEnvKey()
	if err != nil {
		t.Fatal(err)
	}

	cl := NewClient(testAPIKey)
	res, err := cl.GetAPITokenStatus()
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	} else if res.Token == "" {
		t.Error("expected response with non-empty token")
	}

}

func TestGetUser(t *testing.T) {
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

	_, err = cl.GetUser("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.GetUser(auth.Token)
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	} else if res.ID == 0 {
		t.Error("expected response with non-zero ID")
	}
}

// NOTE: expects that test account has at least one wordlist associated with it.
// This is so that we can test for meaningful return values while minimizing
// test interdependency.
func TestGetWordListsForUser(t *testing.T) {
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

	_, err = cl.GetWordListsForUser("")
	if err == nil {
		t.Error("expected error for empty string input")
	}

	res, err := cl.GetWordListsForUser(auth.Token)
	if err != nil {
		t.Error("unexpected error: " + err.Error())
	} else if len(res) == 0 {
		t.Error("expected at least one value in result")
	}
}
