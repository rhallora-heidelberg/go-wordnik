package wordnik

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	base = "https://api.wordnik.com/v4/"
)

// Client is an http.Client wrapper which stores an API key and base url.
type Client struct {
	apiKey  string
	baseURL *url.URL
	client  *http.Client
}

// NewClient creates a Client with the specified API key. The http.Client
// component is configured with a 10-second timeout.
func NewClient(key string) *Client {
	baseURL, err := url.Parse(base)
	if err != nil {
		panic(err)
	}

	return &Client{key, baseURL, &http.Client{Timeout: time.Second * 10}}
}

func (c *Client) formRequest(relativePath *url.URL, vals url.Values, method string) (*http.Request, error) {
	u := c.baseURL.ResolveReference(relativePath)
	u.RawQuery = vals.Encode()
	request, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return request, err
	}

	request.Header["api_key"] = []string{c.apiKey}
	request.Header["Content-type"] = []string{"application/json"}
	return request, nil
}

func (c *Client) doRequest(req *http.Request, dst interface{}) error {
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(dst)
}

// basicGetRequest is a helper method which makes most of the GET requests
// endpoints simpler.
func (c *Client) basicGetRequest(rel *url.URL, vals url.Values, dst interface{}, options ...QueryOption) error {
	for _, option := range options {
		option(&vals)
	}

	req, err := c.formRequest(rel, vals, "GET")
	if err != nil {
		return err
	}

	err = c.doRequest(req, dst)

	return err
}
