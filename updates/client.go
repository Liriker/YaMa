package updates

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Liriker/YaMa/types"
	"io"
	"net/http"
)

const (
	updateUrl = "https://botapi.messenger.yandex.net/bot/v1/messages/getUpdates/"
)

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

func (cl *Client) GetUpdates(limit, offset int64) ([]types.Update, int64, error) {
	reqBody := updateRequest{
		Limit:  limit,
		Offset: offset,
	}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequest(http.MethodPost, updateUrl, bytes.NewBuffer(data))
	if err != nil {
		return nil, 0, err
	}
	req.Header = cl.headers

	resp, err := cl.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	response := updateResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil, 0, errors.New(string(body))
	}

	if !response.Ok {
		return nil, 0, errors.New(response.Description)
	}
	if len(response.Updates) > 0 {
		offset = response.Updates[len(response.Updates)-1].UpdateID + 1
	}
	return response.Updates, offset, nil

}

func (cl *Client) SetWebhook(url string) (bool, string, error) {
	reqBody := webhookRequest{
		WebhookUrl: url,
	}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return false, "", err
	}
	req, err := http.NewRequest(http.MethodPost, updateUrl, bytes.NewBuffer(data))
	if err != nil {
		return false, "", err
	}
	req.Header = cl.headers
	resp, err := cl.client.Do(req)
	if err != nil {
		return false, "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, "", err
	}

	response := webhookResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, "", errors.New(string(body))
	}
	return response.Ok, response.Id, nil
}
