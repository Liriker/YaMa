package chats

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Liriker/YaMa/types"
	"net/http"
)

const (
	createUrl   = "https://botapi.messenger.yandex.net/bot/v1/chats/create/"
	updateUrl   = "https://botapi.messenger.yandex.net/bot/v1/chats/updateMembers/"
	userLinkUrl = "https://botapi.messenger.yandex.net/bot/v1/users/getUserLink/"
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

// Create - The method allows you to create a chat or channel, add its description and icon, assign administrators, add participants (for the chat) or subscribers (for the channel).
// Result of this method is string with Chat ID.
// A bot can create a chat (channel) only with members of the organization to which it belongs.
// All created chats (channels) belong to the organization that owns the bot.
// The bot becomes the administrator of the created chat (channel).
// The bot cannot add a participant to the chat for whom this is prohibited by the privacy settings.
func (c *Client) Create(chat types.NewChat) (string, error) {
	data, err := json.Marshal(chat)
	if err != nil {
		return "", err
	}
	body := bytes.NewBuffer(data)

	req, err := http.NewRequest(http.MethodPost, createUrl, body)
	if err != nil {
		return "", err
	}

	req.Header = c.headers
	respData, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	resp := response{}
	err = json.NewDecoder(respData.Body).Decode(&resp)
	if err != nil {
		return "", errors.New(fmt.Sprintf("%v", respData))
	}
	if !resp.Ok {
		return "", errors.New(fmt.Sprint(resp.Description))
	}
	if respData.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprint(respData))
	}
	return resp.ChatID, nil

}

// Update - The method allows you to add and remove participants to the chat, add and remove subscribers to the channel, as well as appoint chat or channel administrators.
// Each user in the request must be unique, otherwise an error will be returned.
// На момент написания почему-то запрос, соответствующий документации выдаёт ошибку invalid_request, что поле "login" является обязательным, хотя оно есть.
// TODO - проверить отправку запроса
func (c *Client) Update(update *types.ChatUpdate) error {
	data, err := json.Marshal(update)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest(http.MethodPost, updateUrl, body)
	if err != nil {
		return err
	}
	req.Header = c.headers

	respData, err := c.client.Do(req)
	if err != nil {
		return err
	}
	resp := response{}
	err = json.NewDecoder(respData.Body).Decode(&resp)
	if err != nil {
		return errors.New(fmt.Sprint(respData))
	}
	if !resp.Ok {
		return errors.New(fmt.Sprint(resp.Description))
	}
	return nil
}

func (c *Client) GetUserLinks(user types.User) (*UserLinkResponse, error) {
	data, err := json.Marshal(user.Login)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest(http.MethodGet, userLinkUrl, body)
	if err != nil {
		return nil, err
	}
	headers := c.headers.Clone()
	headers.Del("Content-Type")
	req.Header = headers
	respData, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	resp := UserLinkResponse{}
	err = json.NewDecoder(respData.Body).Decode(&resp)
	if err != nil {
		return nil, errors.New(fmt.Sprint(respData))
	}
	if !resp.Ok {
		response := response{}
		err := json.NewDecoder(respData.Body).Decode(&response)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%v", respData))
		}
		return nil, errors.New(fmt.Sprintf("%v", response))
	}
	return &resp, nil
}
