package wordnik

import (
	"errors"
	"os"
	"sync"
	"testing"
)

var (
	cl   *Client
	once sync.Once
)

// Helper function for testing which either initializes a Client or returns
// the already-existing Client
func getClient(t *testing.T) *Client {
	once.Do(func() {
		testAPIKey, err := getEnvKey()
		if err != nil {
			t.Fatal(err)
		}

		cl = NewClient(testAPIKey)
	})
	return cl
}

// Helper function for testing which retrieves an API key from
// the environment variable WORDNIK_API_KEY. Returns an error if
// no key is found.
func getEnvKey() (string, error) {
	keyVar := os.Getenv("WORDNIK_API_KEY")
	if keyVar != "" {
		return keyVar, nil
	}

	return keyVar, errors.New("environment variable WORDNIK_API_KEY not set")
}

type testUser struct {
	user, pass string
}

// Helper function for testing which retrieves a test username and password from
// the environment variables WORDNIK_TEST_USER and WORDNIK_TEST_PASS. Returns an
// error if no key is found.
func getEnvUserPass() (testUser, error) {
	user := os.Getenv("WORDNIK_TEST_USER")
	pass := os.Getenv("WORDNIK_TEST_PASS")
	if user == "" || pass == "" {
		err := errors.New("environment variable(s) WORDNIK_TEST_USER and/or WORDNIK_TEST_PASS not set")
		return testUser{user, pass}, err
	}

	return testUser{user, pass}, nil
}
