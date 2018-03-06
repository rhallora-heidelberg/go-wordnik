package wordnik

import (
	"net/url"
	"testing"
)

func TestNewClient(t *testing.T) {
	cl := NewClient("abc")
	if cl.apiKey != "abc" {
		t.Error("API key failed to embed into Client struct")
	}

	_, err := url.Parse(base)
	if err != nil {
		t.Fatal("Base url constant 'base' failed to parse")
	}
}

func TestFormRequest(t *testing.T) {
	cl := NewClient("abc")
	_, err := cl.formRequest(cl.baseURL, url.Values{}, "bad method")
	if err == nil {
		t.Error("Expected error for invalid method")
	}
}
