package wordnik

import (
	"bytes"
	"errors"
	"net/url"
)

// AuthenticationToken as defined by the Wordnik API. Needed for user-specific
// requests.
type AuthenticationToken struct {
	Token         string `json:"token"`
	UserID        int64  `json:"userId"`
	UserSignature string `json:"userSignature"`
}

// APITokenStatus as defined by the Wordnik API.
type APITokenStatus struct {
	Valid           bool   `json:"valid"`
	Token           string `json:"token"`
	ResetsInMillis  int64  `json:"resetsInMillis"`
	RemainingCalls  int64  `json:"remainingCalls"`
	ExpiresInMillis int64  `json:"expiresInMillis"`
	TotalRequests   int64  `json:"totalRequests"`
}

// User as defined by the Wordnik API.
type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Status      int64  `json:"status"`
	FaceBookID  string `json:"faceBookId"`
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
	Password    string `json:"password"`
}

// AuthenticateGET returns an AuthenticationToken object for a given user.
func (c *Client) AuthenticateGET(user, pass string) (AuthenticationToken, error) {
	if user == "" || pass == "" {
		return AuthenticationToken{}, errors.New("empty username/password not allowed")
	}

	rel := &url.URL{Path: "account.json/authenticate/" + user}

	// Default values
	q := url.Values{
		"password": []string{pass},
	}

	var results AuthenticationToken
	err := c.basicGetRequest(rel, q, &results)

	return results, err
}

// AuthenticatePOST returns an AuthenticationToken object for a given user.
func (c *Client) AuthenticatePOST(user, pass string) (AuthenticationToken, error) {
	if user == "" || pass == "" {
		return AuthenticationToken{}, errors.New("empty username/password not allowed")
	}

	rel := &url.URL{Path: "account.json/authenticate/" + user}
	body := bytes.NewBufferString(pass)
	req, err := c.formRequest(rel, url.Values{}, "POST", body)
	if err != nil {
		return AuthenticationToken{}, err
	}

	var results AuthenticationToken
	err = c.doRequest(req, &results)

	return results, err
}

// GetAPITokenStatus returns an APITokenStatus object for a given API key.
func (c *Client) GetAPITokenStatus() (APITokenStatus, error) {
	rel := &url.URL{Path: "account.json/apiTokenStatus"}

	var results APITokenStatus
	err := c.basicGetRequest(rel, url.Values{}, &results)

	return results, err
}

// GetUser returns a User object for a given authorization token.
func (c *Client) GetUser(authToken string) (User, error) {
	if authToken == "" {
		return User{}, errors.New("empty auth token not allowed")
	}

	rel := &url.URL{Path: "account.json/user"}

	req, err := c.formRequest(rel, url.Values{}, "GET")
	if err != nil {
		return User{}, err
	}

	req.Header["auth_token"] = []string{authToken}

	var results User
	err = c.doRequest(req, &results)

	return results, err
}

// GetWordListsForUser returns a slice of WordList objects for a given account.
func (c *Client) GetWordListsForUser(authToken string, options ...QueryOption) ([]WordList, error) {
	if authToken == "" {
		return []WordList{}, errors.New("empty auth token not allowed")
	}

	rel := &url.URL{Path: "account.json/wordLists"}

	// Default values
	q := url.Values{
		"skip":  []string{"0"},
		"limit": []string{"50"},
	}

	for _, option := range options {
		option(&q)
	}

	req, err := c.formRequest(rel, q, "GET")
	if err != nil {
		return []WordList{}, err
	}

	req.Header["auth_token"] = []string{authToken}
	var results []WordList
	err = c.doRequest(req, &results)

	return results, err
}
