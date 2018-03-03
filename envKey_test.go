package wordnik

import (
	"errors"
	"os"
)

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
