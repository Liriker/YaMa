package transport

import (
	"github.com/Liriker/YaMa/chats"
	"github.com/Liriker/YaMa/messages"
	"github.com/Liriker/YaMa/polling"
	"github.com/Liriker/YaMa/updates"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	headers    http.Header
	Chats      *chats.Client
	Messages   *messages.Client
	Polling    *polling.Client
	Updates    *updates.Client
}

func NewClient(token string) *Client {
	client := http.DefaultClient
	headers := http.Header{}
	headers.Add("Authorization", "OAuth "+token)
	headers.Add("Content-Type", "application/json")

	return &Client{
		httpClient: client,
		headers:    headers,
		Chats:      chats.NewClient(client, headers),
		Messages:   messages.NewClient(client, headers),
		Polling:    polling.NewClient(client, headers),
		Updates:    updates.NewClient(client, headers),
	}
}
