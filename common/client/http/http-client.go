package http

import "net/http"

type Client struct {
	endpoint string
	httpClient *http.Client
	defaultHeaders map[string]string
}

func NewClient(endpoint string) *Client {
	c := &Client{
		endpoint:       endpoint,
		httpClient:     http.DefaultClient,
		defaultHeaders: make(map[string]string),
	}

	return c
}

// AddDefaultHeader - add default header to every request
func (c *Client) AddDefaultHeader(key string, value string) {
	c.defaultHeaders[key] = value
}

