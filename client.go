package wordnik

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	base = "http://api.wordnik.com/v4/"
)

type Client struct {
	apiKey  string
	baseURL *url.URL
	client  *http.Client
}

func NewClient(key string) Client {
	baseURL, err := url.Parse(base)
	if err != nil {
		panic(err)
	}

	return Client{key, baseURL, &http.Client{Timeout: time.Second * 10}}
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
