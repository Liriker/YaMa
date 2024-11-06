package YaMa

import (
	"errors"
	"github.com/Liriker/YaMa/transport"
)

func NewClient(token string) (*transport.Client, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}
	return transport.NewClient(token), nil
}
