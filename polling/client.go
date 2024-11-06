package polling

import "net/http"

type Client struct {
	client  *http.Client
	headers http.Header
}

func NewClient(cl *http.Client, h http.Header) *Client {
	return &Client{
		client:  cl,
		headers: h,
	}
}
